// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package charmrevisionupdater

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/juju/clock"
	"github.com/juju/collections/set"
	"github.com/juju/errors"

	apiservererrors "github.com/juju/juju/apiserver/errors"
	charmmetrics "github.com/juju/juju/core/charm/metrics"
	corelogger "github.com/juju/juju/core/logger"
	"github.com/juju/juju/core/objectstore"
	"github.com/juju/juju/environs/config"
	"github.com/juju/juju/internal/charm"
	"github.com/juju/juju/internal/charm/resource"
	"github.com/juju/juju/internal/charmhub"
	"github.com/juju/juju/rpc/params"
	"github.com/juju/juju/version"
)

// CharmRevisionUpdater defines the methods on the charmrevisionupdater API end point.
type CharmRevisionUpdater interface {
	UpdateLatestRevisions(ctx context.Context) (params.ErrorResult, error)
}

// CharmRevisionUpdaterAPI implements the CharmRevisionUpdater interface and is the concrete
// implementation of the api end point.
type CharmRevisionUpdaterAPI struct {
	state State
	store objectstore.ObjectStore
	clock clock.Clock

	newCharmhubClient newCharmhubClientFunc
	logger            corelogger.Logger
}

type newCharmhubClientFunc func(st State) (CharmhubRefreshClient, error)

var _ CharmRevisionUpdater = (*CharmRevisionUpdaterAPI)(nil)

// NewCharmRevisionUpdaterAPIState creates a new charmrevisionupdater API
// with a State interface directly (mainly for use in tests).
func NewCharmRevisionUpdaterAPIState(
	state State,
	store objectstore.ObjectStore,
	clock clock.Clock,
	newCharmhubClient newCharmhubClientFunc,
	logger corelogger.Logger,
) (*CharmRevisionUpdaterAPI, error) {
	return &CharmRevisionUpdaterAPI{
		state:             state,
		store:             store,
		clock:             clock,
		newCharmhubClient: newCharmhubClient,
		logger:            logger,
	}, nil
}

// UpdateLatestRevisions retrieves the latest revision information from the charm store for all deployed charms
// and records this information in state.
func (api *CharmRevisionUpdaterAPI) UpdateLatestRevisions(ctx context.Context) (params.ErrorResult, error) {
	if err := api.updateLatestRevisions(ctx); err != nil {
		return params.ErrorResult{Error: apiservererrors.ServerError(err)}, nil
	}
	return params.ErrorResult{}, nil
}

func (api *CharmRevisionUpdaterAPI) updateLatestRevisions(ctx context.Context) error {
	// Look up the information for all the deployed charms. This is the
	// "expensive" part.
	latest, err := api.retrieveLatestCharmInfo(ctx)
	if err != nil {
		return errors.Trace(err)
	}

	// Process the resulting info for each charm.
	resources := api.state.Resources(api.store)
	for _, info := range latest {
		// First, add a charm placeholder to the model for each.
		if err := api.state.AddCharmPlaceholder(info.url); err != nil {
			return errors.Trace(err)
		}

		// Then save the resources
		err := resources.SetCharmStoreResources(info.appID, info.resources, info.timestamp)
		if err != nil {
			return errors.Trace(err)
		}
	}

	return nil
}

type latestCharmInfo struct {
	url       *charm.URL
	timestamp time.Time
	revision  int
	resources []resource.Resource
	appID     string
}

type appInfo struct {
	id       string
	charmURL *charm.URL
}

// retrieveLatestCharmInfo looks up the charm store to return the charm URLs for the
// latest revision of the deployed charms.
func (api *CharmRevisionUpdaterAPI) retrieveLatestCharmInfo(ctx context.Context) ([]latestCharmInfo, error) {
	applications, err := api.state.AllApplications()
	if err != nil {
		return nil, errors.Trace(err)
	}
	model, err := api.state.Model()
	if err != nil {
		return nil, errors.Trace(err)
	}
	cfg, err := model.Config()
	if err != nil {
		return nil, errors.Trace(err)
	}
	telemetry := cfg.Telemetry()

	// If there are no applications, exit now, check telemetry for additional work.
	if len(applications) == 0 {
		if !telemetry {
			return nil, nil
		}
		return nil, errors.Trace(api.sendEmptyModelMetrics(ctx))
	}

	var (
		charmhubIDs  []charmhubID
		charmhubApps []appInfo
	)
	for _, application := range applications {
		cURL, _ := application.CharmURL()
		curl, err := charm.ParseURL(*cURL)
		if err != nil {
			return nil, errors.Trace(err)
		}
		switch {
		case charm.Local.Matches(curl.Schema):
			continue

		case charm.CharmHub.Matches(curl.Schema):
			origin := application.CharmOrigin()
			if origin == nil {
				// If this fails, we have big problems, so make this Errorf
				api.logger.Errorf("charm %s has no origin, skipping", curl)
				continue
			}
			if origin.ID == "" || origin.Revision == nil || origin.Channel == nil || origin.Platform == nil {
				api.logger.Errorf("charm %s has missing id(%s), revision (%p), channel (%p), or platform (%p), skipping",
					curl, origin.Revision, origin.Channel, origin.Platform)
				continue
			}
			channel, err := charm.MakeChannel(origin.Channel.Track, origin.Channel.Risk, origin.Channel.Branch)
			if err != nil {
				return nil, errors.Trace(err)
			}
			cid := charmhubID{
				id:          origin.ID,
				revision:    *origin.Revision,
				channel:     channel.String(),
				osType:      strings.ToLower(origin.Platform.OS), // charmhub API requires lowercase OS key
				osChannel:   origin.Platform.Channel,
				arch:        origin.Platform.Architecture,
				instanceKey: charmhub.CreateInstanceKey(application.ApplicationTag(), model.ModelTag()),
			}
			if telemetry {
				cid.metrics = map[charmmetrics.MetricKey]string{
					charmmetrics.NumUnits: strconv.Itoa(application.UnitCount()),
				}
			}
			charmhubIDs = append(charmhubIDs, cid)
			charmhubApps = append(charmhubApps, appInfo{
				id:       application.ApplicationTag().Id(),
				charmURL: curl,
			})

		default:
			return nil, errors.NotValidf("charm schema %q", curl.Schema)
		}
	}

	var (
		charmhubErr error

		latest []latestCharmInfo
	)
	if len(charmhubIDs) > 0 {
		if telemetry {
			charmhubIDs, err = api.addMetricsToCharmhubInfos(charmhubIDs, charmhubApps)
			if err != nil {
				// It's fine to error out here, as this is a state backed
				// request and should be transitive.
				return nil, errors.Trace(err)
			}
		}
		hubLatest, err := api.fetchCharmhubInfos(ctx, cfg, charmhubIDs, charmhubApps)
		if err != nil {
			charmhubErr = err
		} else {
			latest = append(latest, hubLatest...)
		}
	}

	if charmhubErr != nil {
		return nil, errors.Errorf("charmhub: %v", charmhubErr)
	}

	return latest, nil
}

func (api *CharmRevisionUpdaterAPI) addMetricsToCharmhubInfos(ids []charmhubID, appInfos []appInfo) ([]charmhubID, error) {
	relationKeys := api.state.AliveRelationKeys()
	if len(relationKeys) == 0 {
		return ids, nil
	}
	for k, v := range convertRelations(appInfos, relationKeys) {
		for i, app := range appInfos {
			if app.id != k {
				continue
			}
			if ids[i].metrics == nil {
				ids[i].metrics = map[charmmetrics.MetricKey]string{}
			}
			ids[i].metrics[charmmetrics.Relations] = strings.Join(v, ",")
		}
	}
	return ids, nil
}

// sendEmptyModelMetrics sends the controller and model metrics
// for an empty model.  This is highly likely for juju 2.9.  A
// controller charm was introduced for juju 3.0.  Allows us to get
// the controller version etc.
func (api *CharmRevisionUpdaterAPI) sendEmptyModelMetrics(ctx context.Context) error {
	requestMetrics, err := charmhubRequestMetadata(api.state)
	if err != nil {
		return errors.Trace(err)
	}
	client, err := api.newCharmhubClient(api.state)
	if err != nil {
		return errors.Trace(err)
	}
	ctx, cancel := context.WithTimeout(ctx, charmhub.RefreshTimeout)
	defer cancel()
	return errors.Trace(client.RefreshWithMetricsOnly(ctx, requestMetrics))
}

// convertRelations converts the list of relations by application name to charm name
// for the metrics response.
func convertRelations(appInfos []appInfo, relationKeys []string) map[string][]string {
	// Map application names to its charm name.
	appToCharm := make(map[string]string)
	for _, v := range appInfos {
		if v.charmURL != nil {
			appToCharm[v.id] = v.charmURL.Name
		}
	}

	// relationsByAppName is a map of relation providers, by application name,
	// to a slice of requirers, by application name.
	relationsByAppName := make(map[string][]string)
	for _, key := range relationKeys {
		one, two, use := relationApplicationNames(key)
		if !use {
			continue
		}
		values, _ := relationsByAppName[one]
		values = append(values, two)
		relationsByAppName[one] = values
		rValues, _ := relationsByAppName[two]
		rValues = append(rValues, one)
		relationsByAppName[two] = rValues
	}

	// Put them together to create a map of relations, by
	// application name, to a slice of relations, by charm name.
	relations := make(map[string][]string)
	for appName, appNameRels := range relationsByAppName {
		// It is possible the same charm is deployed more than once with
		// different names.
		c := relations[appName]
		relatedTo := set.NewStrings(c...)
		// Using a set, also ensures there are no duplicates.
		for _, v := range appNameRels {
			relatedTo.Add(appToCharm[v])
		}
		relations[appName] = relatedTo.SortedValues()
	}

	return relations
}

// relationApplicationNames returns the application names in the relation.
// Peer relations are filtered out, not needed for metrics.
// Keys are in the format: "appName:endpoint appName:endpoint"
func relationApplicationNames(str string) (string, string, bool) {
	endpoints := strings.Split(str, " ")
	if len(endpoints) != 2 {
		return "", "", false
	}
	one := strings.Split(endpoints[0], ":")
	if len(one) != 2 {
		return "", "", false
	}
	two := strings.Split(endpoints[1], ":")
	if len(two) != 2 {
		return "", "", false
	}
	return one[0], two[0], true
}

func (api *CharmRevisionUpdaterAPI) fetchCharmhubInfos(ctx context.Context, cfg *config.Config, ids []charmhubID, appInfos []appInfo) ([]latestCharmInfo, error) {
	var requestMetrics map[charmmetrics.MetricKey]map[charmmetrics.MetricKey]string
	if cfg.Telemetry() {
		var err error
		requestMetrics, err = charmhubRequestMetadata(api.state)
		if err != nil {
			return nil, errors.Trace(err)
		}
	}
	client, err := api.newCharmhubClient(api.state)
	if err != nil {
		return nil, errors.Trace(err)
	}
	results, err := charmhubLatestCharmInfo(ctx, client, requestMetrics, ids, api.clock.Now().UTC(), api.logger)
	if err != nil {
		return nil, errors.Trace(err)
	}

	var latest []latestCharmInfo
	for i, result := range results {
		if i >= len(appInfos) {
			api.logger.Errorf("retrieved more results (%d) than charmhub applications (%d)",
				i, len(appInfos))
			break
		}
		if result.error != nil {
			api.logger.Errorf("retrieving charm info for ID %s: %v", ids[i].id, result.error)
			continue
		}
		appInfo := appInfos[i]
		latest = append(latest, latestCharmInfo{
			url:       appInfo.charmURL.WithRevision(result.revision),
			timestamp: result.timestamp,
			revision:  result.revision,
			resources: result.resources,
			appID:     appInfo.id,
		})
	}
	return latest, nil
}

// charmhubRequestMetadata returns a map containing metadata key/value pairs to
// send to the charmhub for tracking metrics.
func charmhubRequestMetadata(st State) (map[charmmetrics.MetricKey]map[charmmetrics.MetricKey]string, error) {
	model, err := st.Model()
	if err != nil {
		return nil, errors.Trace(err)
	}

	metrics, err := model.Metrics()
	if err != nil {
		return nil, errors.Trace(err)
	}

	metadata := map[charmmetrics.MetricKey]map[charmmetrics.MetricKey]string{
		charmmetrics.Controller: {
			charmmetrics.JujuVersion: version.Current.String(),
			charmmetrics.UUID:        metrics.ControllerUUID,
		},
		charmmetrics.Model: {
			charmmetrics.Cloud:           metrics.CloudName,
			charmmetrics.UUID:            metrics.UUID,
			charmmetrics.NumApplications: metrics.ApplicationCount,
			charmmetrics.NumMachines:     metrics.MachineCount,
			charmmetrics.NumUnits:        metrics.UnitCount,
			charmmetrics.Provider:        metrics.Provider,
			charmmetrics.Region:          metrics.CloudRegion,
		},
	}

	return metadata, nil
}
