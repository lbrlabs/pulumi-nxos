package shim

import (
	"github.com/CiscoDevNet/terraform-provider-nxos/internal/provider"
	tfProv "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/lbrlabs/pulumi-nxos/provider/pkg/version"
)

func NewProvider() tfProv.Provider {
	prov := provider.New(version.Version)()
	return prov
}
