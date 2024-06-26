check_managed_identity_controller() {
	local name identity_name
	name=${1}
	identity_name=${2}

	cloud="azure"
	if [[ -n ${BOOTSTRAP_REGION:-} ]]; then
		cloud_region="${cloud}/${BOOTSTRAP_REGION}"
	else
		cloud_region="${cloud}"
	fi

	juju bootstrap "${cloud_region}" "${name}" \
		--debug \
		--constraints="instance-role=${identity_name}" 2>&1 | OUTPUT "${file}"

	cred_name=${AZURE_CREDENTIAL_NAME:-credentials}
	cred=$(juju show-credential --controller "${name}" azure "${cred_name}" 2>&1 || true)
	check_contains "$cred" "managed-identity"

	juju switch controller
	juju enable-ha
	wait_for_controller_machines 3
	wait_for_ha 3

	juju add-model test
	juju deploy jameinel-ubuntu-lite
	wait_for "ubuntu-lite" "$(idle_condition "ubuntu-lite")"

	# Takes too long to tear down, so forcibly destroy it
	export KILL_CONTROLLER=true
	destroy_controller "${name}"

}

run_auto_managed_identity() {
	echo

	name="azure-auto-managed-identity"
	file="${TEST_DIR}/test-auto-managed-identity.log"

	check_managed_identity_controller ${name} "auto"
}

run_custom_managed_identity() {
	echo

	# Create the managed identity to use with the controller.
	group="jtest-$(xxd -l 6 -c 32 -p </dev/random)"
	identity_name=jmid
	subscription=$(az account show --query id --output tsv)

	add_clean_func "run_cleanup_azure"
	az group create --name "${group}" --location westus
	echo "${group}" >>"${TEST_DIR}/azure-groups"
	az identity create --resource-group "${group}" --name jmid
	mid=$(az identity show --resource-group "${group}" --name jmid --query principalId --output tsv)
	az role assignment create --assignee-object-id "${mid}" --assignee-principal-type "ServicePrincipal" --role "JujuRoles" --scope "/subscriptions/${subscription}"

	name="azure-custom-managed-identity"
	file="${TEST_DIR}/test-custom-managed-identity.log"

	check_managed_identity_controller ${name} "${group}/${identity_name}"
}

run_cleanup_azure() {
	set +e
	echo "==> Removing resource groups"

	if [[ -f "${TEST_DIR}/azure-groups" ]]; then
		while read -r group; do
			az group delete -y --resource-group "${group}" >>"${TEST_DIR}/azure_cleanup"
		done <"${TEST_DIR}/azure-groups"
	fi
	echo "==> Removed resource groups"
}

test_managed_identity() {
	if [ "$(skip 'test_managed_identity')" ]; then
		echo "==> TEST SKIPPED: managed identity"
		return
	fi

	if [ "$(az account list | jq length)" -lt 1 ]; then
		echo "==> TEST SKIPPED: not logged in to Azure cloud"
		return
	fi

	(
		set_verbosity

		cd .. || exit

		run "run_auto_managed_identity" "$@"
		run "run_custom_managed_identity" "$@"
	)
}
