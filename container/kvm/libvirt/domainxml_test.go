package libvirt_test

import (
	"encoding/xml"

	gc "gopkg.in/check.v1"

	. "github.com/juju/juju/container/kvm/libvirt"
	jc "github.com/juju/testing/checkers"
	"github.com/pkg/errors"
)

// gocheck boilerplate.
type domainXMLSuite struct{}

var _ = gc.Suite(domainXMLSuite{})

var wantDomainStr = `
<domain type="kvm">
    <os>
        <type>hvm</type>
    </os>
    <features>
        <acpi></acpi>
        <apic></apic>
        <pae></pae>
    </features>
    <devices>
        <controller type="usb" index="0">
            <address type="pci" domain="0x0000" bus="0x00" slot="0x01" function="0x2"></address>
        </controller>
        <controller type="pci" index="0" model="pci-root"></controller>
        <console type="stdio">
            <target type="serial" port="0"></target>
        </console>
        <input type="mouse" bus="ps2"></input>
        <input type="keyboard" bus="ps2"></input>
        <graphics type="vnc" port="-1" autoport="yes" listen="127.0.0.1">
            <listen type="address" address="127.0.0.1"></listen>
        </graphics>
        <video>
            <model type="cirrus" vram="9216" heads="1"></model>
            <address type="pci" domain="0x0000" bus="0x00" slot="0x02" function="0x0"></address>
        </video>
        <interface type="bridge">
            <mac address="00:00:00:00:00:00"></mac>
            <model type="virtio"></model>
            <source bridge="parent-dev"></source>
            <guest dev="device-name"></guest>
        </interface>
        <disk device="disk" type="file">
            <driver type="qcow2" name="qemu"></driver>
            <source file="/some/path"></source>
            <target dev="vda"></target>
        </disk>
        <disk device="disk" type="file">
            <driver type="raw" name="qemu"></driver>
            <source file="/another/path"></source>
            <target dev="vdb"></target>
        </disk>
    </devices>
    <name>juju-someid</name>
    <vcpu>2</vcpu>
    <currentMemory unit="MiB">1024</currentMemory>
    <memory unit="MiB">1024</memory>
</domain>`[1:]

func (domainXMLSuite) TestNewDomain(c *gc.C) {
	ifaces := []InterfaceInfo{
		dummyInterface{
			mac:    "00:00:00:00:00:00",
			parent: "parent-dev",
			name:   "device-name"}}
	disks := []DiskInfo{
		dummyDisk{driver: "qcow2", source: "/some/path"},
		dummyDisk{driver: "raw", source: "/another/path"},
	}
	params := dummyParams{ifaceInfo: ifaces, diskInfo: disks, memory: 1024, cpuCores: 2, hostname: "juju-someid"}
	d, err := NewDomain(params)
	c.Assert(err, jc.ErrorIsNil)
	ml, err := xml.MarshalIndent(&d, "", "    ")
	c.Check(err, jc.ErrorIsNil)
	c.Assert(string(ml), jc.DeepEquals, wantDomainStr)
}

func (domainXMLSuite) TestNewDomainError(c *gc.C) {
	d, err := NewDomain(dummyParams{err: errors.Errorf("boom")})
	c.Check(d, jc.DeepEquals, Domain{})
	c.Check(err, gc.ErrorMatches, "boom")
}

type dummyParams struct {
	err       error
	cpuCores  uint64
	diskInfo  []DiskInfo
	hostname  string
	ifaceInfo []InterfaceInfo
	memory    uint64
}

func (p dummyParams) DiskInfo() []DiskInfo        { return p.diskInfo }
func (p dummyParams) Interfaces() []InterfaceInfo { return p.ifaceInfo }
func (p dummyParams) Hostname() string            { return p.hostname }
func (p dummyParams) CPUCores() uint64            { return p.cpuCores }
func (p dummyParams) Memory() uint64              { return p.memory }
func (p dummyParams) Validate() error             { return p.err }

type dummyDisk struct {
	source string
	driver string
}

func (d dummyDisk) Source() string { return d.source }
func (d dummyDisk) Driver() string { return d.driver }

type dummyInterface struct {
	mac, parent, name string
}

func (i dummyInterface) MAC() string              { return i.mac }
func (i dummyInterface) ParentDeviceName() string { return i.parent }
func (i dummyInterface) DeviceName() string       { return i.name }
