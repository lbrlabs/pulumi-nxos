package main

import (
	nxos "github.com/lbrlabs/pulumi-nxos/provider"
	"github.com/lbrlabs/pulumi-nxos/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen"
)

func main() {
	// Modify the path to point to the new provider
	tfgen.Main("nxos", version.Version, nxos.Provider())
}
