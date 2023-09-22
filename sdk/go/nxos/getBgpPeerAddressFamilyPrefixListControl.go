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

// This data source can read the BGP peer address family prefix list control configuration.
//
// - API Documentation: [bgpPfxCtrlP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:PfxCtrlP/)
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
//			_, err := nxos.LookupBgpPeerAddressFamilyPrefixListControl(ctx, &nxos.LookupBgpPeerAddressFamilyPrefixListControlArgs{
//				Address:       "192.168.0.1",
//				AddressFamily: "ipv4-ucast",
//				Asn:           "65001",
//				Direction:     "in",
//				Vrf:           "default",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupBgpPeerAddressFamilyPrefixListControl(ctx *pulumi.Context, args *LookupBgpPeerAddressFamilyPrefixListControlArgs, opts ...pulumi.InvokeOption) (*LookupBgpPeerAddressFamilyPrefixListControlResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBgpPeerAddressFamilyPrefixListControlResult
	err := ctx.Invoke("nxos:index/getBgpPeerAddressFamilyPrefixListControl:getBgpPeerAddressFamilyPrefixListControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBgpPeerAddressFamilyPrefixListControl.
type LookupBgpPeerAddressFamilyPrefixListControlArgs struct {
	// Peer address.
	Address string `pulumi:"address"`
	// Address Family.
	AddressFamily string `pulumi:"addressFamily"`
	// Autonomous system number.
	Asn string `pulumi:"asn"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Route Control direction.
	Direction string `pulumi:"direction"`
	// VRF name.
	Vrf string `pulumi:"vrf"`
}

// A collection of values returned by getBgpPeerAddressFamilyPrefixListControl.
type LookupBgpPeerAddressFamilyPrefixListControlResult struct {
	// Peer address.
	Address string `pulumi:"address"`
	// Address Family.
	AddressFamily string `pulumi:"addressFamily"`
	// Autonomous system number.
	Asn string `pulumi:"asn"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Route Control direction.
	Direction string `pulumi:"direction"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// Route Control Prefix-List name.
	List string `pulumi:"list"`
	// VRF name.
	Vrf string `pulumi:"vrf"`
}

func LookupBgpPeerAddressFamilyPrefixListControlOutput(ctx *pulumi.Context, args LookupBgpPeerAddressFamilyPrefixListControlOutputArgs, opts ...pulumi.InvokeOption) LookupBgpPeerAddressFamilyPrefixListControlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBgpPeerAddressFamilyPrefixListControlResult, error) {
			args := v.(LookupBgpPeerAddressFamilyPrefixListControlArgs)
			r, err := LookupBgpPeerAddressFamilyPrefixListControl(ctx, &args, opts...)
			var s LookupBgpPeerAddressFamilyPrefixListControlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBgpPeerAddressFamilyPrefixListControlResultOutput)
}

// A collection of arguments for invoking getBgpPeerAddressFamilyPrefixListControl.
type LookupBgpPeerAddressFamilyPrefixListControlOutputArgs struct {
	// Peer address.
	Address pulumi.StringInput `pulumi:"address"`
	// Address Family.
	AddressFamily pulumi.StringInput `pulumi:"addressFamily"`
	// Autonomous system number.
	Asn pulumi.StringInput `pulumi:"asn"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// Route Control direction.
	Direction pulumi.StringInput `pulumi:"direction"`
	// VRF name.
	Vrf pulumi.StringInput `pulumi:"vrf"`
}

func (LookupBgpPeerAddressFamilyPrefixListControlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpPeerAddressFamilyPrefixListControlArgs)(nil)).Elem()
}

// A collection of values returned by getBgpPeerAddressFamilyPrefixListControl.
type LookupBgpPeerAddressFamilyPrefixListControlResultOutput struct{ *pulumi.OutputState }

func (LookupBgpPeerAddressFamilyPrefixListControlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpPeerAddressFamilyPrefixListControlResult)(nil)).Elem()
}

func (o LookupBgpPeerAddressFamilyPrefixListControlResultOutput) ToLookupBgpPeerAddressFamilyPrefixListControlResultOutput() LookupBgpPeerAddressFamilyPrefixListControlResultOutput {
	return o
}

func (o LookupBgpPeerAddressFamilyPrefixListControlResultOutput) ToLookupBgpPeerAddressFamilyPrefixListControlResultOutputWithContext(ctx context.Context) LookupBgpPeerAddressFamilyPrefixListControlResultOutput {
	return o
}

func (o LookupBgpPeerAddressFamilyPrefixListControlResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBgpPeerAddressFamilyPrefixListControlResult] {
	return pulumix.Output[LookupBgpPeerAddressFamilyPrefixListControlResult]{
		OutputState: o.OutputState,
	}
}

// Peer address.
func (o LookupBgpPeerAddressFamilyPrefixListControlResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerAddressFamilyPrefixListControlResult) string { return v.Address }).(pulumi.StringOutput)
}

// Address Family.
func (o LookupBgpPeerAddressFamilyPrefixListControlResultOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerAddressFamilyPrefixListControlResult) string { return v.AddressFamily }).(pulumi.StringOutput)
}

// Autonomous system number.
func (o LookupBgpPeerAddressFamilyPrefixListControlResultOutput) Asn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerAddressFamilyPrefixListControlResult) string { return v.Asn }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o LookupBgpPeerAddressFamilyPrefixListControlResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBgpPeerAddressFamilyPrefixListControlResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// Route Control direction.
func (o LookupBgpPeerAddressFamilyPrefixListControlResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerAddressFamilyPrefixListControlResult) string { return v.Direction }).(pulumi.StringOutput)
}

// The distinguished name of the object.
func (o LookupBgpPeerAddressFamilyPrefixListControlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerAddressFamilyPrefixListControlResult) string { return v.Id }).(pulumi.StringOutput)
}

// Route Control Prefix-List name.
func (o LookupBgpPeerAddressFamilyPrefixListControlResultOutput) List() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerAddressFamilyPrefixListControlResult) string { return v.List }).(pulumi.StringOutput)
}

// VRF name.
func (o LookupBgpPeerAddressFamilyPrefixListControlResultOutput) Vrf() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerAddressFamilyPrefixListControlResult) string { return v.Vrf }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBgpPeerAddressFamilyPrefixListControlResultOutput{})
}