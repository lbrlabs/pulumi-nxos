// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nxos

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-nxos/sdk/go/nxos/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// This data source can read the BGP domain (VRF) graceful restart configuration.
//
// - API Documentation: [bgpGr](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Gr/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-nxos/sdk/go/nxos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := nxos.LookupBgpGracefulRestart(ctx, &nxos.LookupBgpGracefulRestartArgs{
//				Asn: "65001",
//				Vrf: "default",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupBgpGracefulRestart(ctx *pulumi.Context, args *LookupBgpGracefulRestartArgs, opts ...pulumi.InvokeOption) (*LookupBgpGracefulRestartResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBgpGracefulRestartResult
	err := ctx.Invoke("nxos:index/getBgpGracefulRestart:getBgpGracefulRestart", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBgpGracefulRestart.
type LookupBgpGracefulRestartArgs struct {
	// Autonomous system number.
	Asn string `pulumi:"asn"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// VRF name.
	Vrf string `pulumi:"vrf"`
}

// A collection of values returned by getBgpGracefulRestart.
type LookupBgpGracefulRestartResult struct {
	// Autonomous system number.
	Asn string `pulumi:"asn"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// The graceful restart interval.
	RestartInterval int `pulumi:"restartInterval"`
	// The stale interval for routes advertised by the BGP peer.
	StaleInterval int `pulumi:"staleInterval"`
	// VRF name.
	Vrf string `pulumi:"vrf"`
}

func LookupBgpGracefulRestartOutput(ctx *pulumi.Context, args LookupBgpGracefulRestartOutputArgs, opts ...pulumi.InvokeOption) LookupBgpGracefulRestartResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBgpGracefulRestartResult, error) {
			args := v.(LookupBgpGracefulRestartArgs)
			r, err := LookupBgpGracefulRestart(ctx, &args, opts...)
			var s LookupBgpGracefulRestartResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBgpGracefulRestartResultOutput)
}

// A collection of arguments for invoking getBgpGracefulRestart.
type LookupBgpGracefulRestartOutputArgs struct {
	// Autonomous system number.
	Asn pulumi.StringInput `pulumi:"asn"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// VRF name.
	Vrf pulumi.StringInput `pulumi:"vrf"`
}

func (LookupBgpGracefulRestartOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpGracefulRestartArgs)(nil)).Elem()
}

// A collection of values returned by getBgpGracefulRestart.
type LookupBgpGracefulRestartResultOutput struct{ *pulumi.OutputState }

func (LookupBgpGracefulRestartResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpGracefulRestartResult)(nil)).Elem()
}

func (o LookupBgpGracefulRestartResultOutput) ToLookupBgpGracefulRestartResultOutput() LookupBgpGracefulRestartResultOutput {
	return o
}

func (o LookupBgpGracefulRestartResultOutput) ToLookupBgpGracefulRestartResultOutputWithContext(ctx context.Context) LookupBgpGracefulRestartResultOutput {
	return o
}

func (o LookupBgpGracefulRestartResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBgpGracefulRestartResult] {
	return pulumix.Output[LookupBgpGracefulRestartResult]{
		OutputState: o.OutputState,
	}
}

// Autonomous system number.
func (o LookupBgpGracefulRestartResultOutput) Asn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpGracefulRestartResult) string { return v.Asn }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o LookupBgpGracefulRestartResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBgpGracefulRestartResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// The distinguished name of the object.
func (o LookupBgpGracefulRestartResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpGracefulRestartResult) string { return v.Id }).(pulumi.StringOutput)
}

// The graceful restart interval.
func (o LookupBgpGracefulRestartResultOutput) RestartInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpGracefulRestartResult) int { return v.RestartInterval }).(pulumi.IntOutput)
}

// The stale interval for routes advertised by the BGP peer.
func (o LookupBgpGracefulRestartResultOutput) StaleInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpGracefulRestartResult) int { return v.StaleInterval }).(pulumi.IntOutput)
}

// VRF name.
func (o LookupBgpGracefulRestartResultOutput) Vrf() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpGracefulRestartResult) string { return v.Vrf }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBgpGracefulRestartResultOutput{})
}
