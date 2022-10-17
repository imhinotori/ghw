package driver

import (
	"github.com/imhinotori/ghw/pkg/context"
	"github.com/imhinotori/ghw/pkg/pci"
	"github.com/imhinotori/ghw/pkg/topology"
)

type Driver struct {
	// the PCI address where the graphics card can be found
	Address string `json:"address"`
	// The "index" of the card on the bus (generally not useful information,
	// but might as well include it)
	Index int `json:"index"`
	// pointer to a PCIDevice struct that describes the vendor and product
	// model, etc
	// TODO(jaypipes): Rename this field to PCI, instead of DeviceInfo
	DeviceInfo *pci.Device `json:"pci"`
	// Topology node that the graphics card is affined to. Will be nil if the
	// architecture is not NUMA.
	Node *topology.Node `json:"node,omitempty"`
}

type Info struct {
	ctx     *context.Context
	Drivers []*Driver `json:"drivers"`
}
