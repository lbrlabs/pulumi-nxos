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

// This data source can read a VRF Route Target Entry.
//
// - API Documentation: [rtctrlRttEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:RttEntry/)
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
//			_, err := nxos.LookupVrfRouteTarget(ctx, &nxos.LookupVrfRouteTargetArgs{
//				AddressFamily:            "ipv4-ucast",
//				Direction:                "import",
//				RouteTarget:              "route-target:as2-nn2:2:2",
//				RouteTargetAddressFamily: "ipv4-ucast",
//				Vrf:                      "VRF1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVrfRouteTarget(ctx *pulumi.Context, args *LookupVrfRouteTargetArgs, opts ...pulumi.InvokeOption) (*LookupVrfRouteTargetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVrfRouteTargetResult
	err := ctx.Invoke("nxos:index/getVrfRouteTarget:getVrfRouteTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVrfRouteTarget.
type LookupVrfRouteTargetArgs struct {
	// Address family.
	AddressFamily string `pulumi:"addressFamily"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Route Target direction.
	Direction string `pulumi:"direction"`
	// Route Target in NX-OS DME format.
	RouteTarget string `pulumi:"routeTarget"`
	// Route Target Address Family.
	RouteTargetAddressFamily string `pulumi:"routeTargetAddressFamily"`
	// VRF name.
	Vrf string `pulumi:"vrf"`
}

// A collection of values returned by getVrfRouteTarget.
type LookupVrfRouteTargetResult struct {
	// Address family.
	AddressFamily string `pulumi:"addressFamily"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Route Target direction.
	Direction string `pulumi:"direction"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// Route Target in NX-OS DME format.
	RouteTarget string `pulumi:"routeTarget"`
	// Route Target Address Family.
	RouteTargetAddressFamily string `pulumi:"routeTargetAddressFamily"`
	// VRF name.
	Vrf string `pulumi:"vrf"`
}

func LookupVrfRouteTargetOutput(ctx *pulumi.Context, args LookupVrfRouteTargetOutputArgs, opts ...pulumi.InvokeOption) LookupVrfRouteTargetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVrfRouteTargetResult, error) {
			args := v.(LookupVrfRouteTargetArgs)
			r, err := LookupVrfRouteTarget(ctx, &args, opts...)
			var s LookupVrfRouteTargetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVrfRouteTargetResultOutput)
}

// A collection of arguments for invoking getVrfRouteTarget.
type LookupVrfRouteTargetOutputArgs struct {
	// Address family.
	AddressFamily pulumi.StringInput `pulumi:"addressFamily"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// Route Target direction.
	Direction pulumi.StringInput `pulumi:"direction"`
	// Route Target in NX-OS DME format.
	RouteTarget pulumi.StringInput `pulumi:"routeTarget"`
	// Route Target Address Family.
	RouteTargetAddressFamily pulumi.StringInput `pulumi:"routeTargetAddressFamily"`
	// VRF name.
	Vrf pulumi.StringInput `pulumi:"vrf"`
}

func (LookupVrfRouteTargetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVrfRouteTargetArgs)(nil)).Elem()
}

// A collection of values returned by getVrfRouteTarget.
type LookupVrfRouteTargetResultOutput struct{ *pulumi.OutputState }

func (LookupVrfRouteTargetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVrfRouteTargetResult)(nil)).Elem()
}

func (o LookupVrfRouteTargetResultOutput) ToLookupVrfRouteTargetResultOutput() LookupVrfRouteTargetResultOutput {
	return o
}

func (o LookupVrfRouteTargetResultOutput) ToLookupVrfRouteTargetResultOutputWithContext(ctx context.Context) LookupVrfRouteTargetResultOutput {
	return o
}

func (o LookupVrfRouteTargetResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupVrfRouteTargetResult] {
	return pulumix.Output[LookupVrfRouteTargetResult]{
		OutputState: o.OutputState,
	}
}

// Address family.
func (o LookupVrfRouteTargetResultOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrfRouteTargetResult) string { return v.AddressFamily }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o LookupVrfRouteTargetResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVrfRouteTargetResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// Route Target direction.
func (o LookupVrfRouteTargetResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrfRouteTargetResult) string { return v.Direction }).(pulumi.StringOutput)
}

// The distinguished name of the object.
func (o LookupVrfRouteTargetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrfRouteTargetResult) string { return v.Id }).(pulumi.StringOutput)
}

// Route Target in NX-OS DME format.
func (o LookupVrfRouteTargetResultOutput) RouteTarget() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrfRouteTargetResult) string { return v.RouteTarget }).(pulumi.StringOutput)
}

// Route Target Address Family.
func (o LookupVrfRouteTargetResultOutput) RouteTargetAddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrfRouteTargetResult) string { return v.RouteTargetAddressFamily }).(pulumi.StringOutput)
}

// VRF name.
func (o LookupVrfRouteTargetResultOutput) Vrf() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrfRouteTargetResult) string { return v.Vrf }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVrfRouteTargetResultOutput{})
}