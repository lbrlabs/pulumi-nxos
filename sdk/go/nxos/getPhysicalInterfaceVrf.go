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

// This data source can read a physical interface VRF association.
//
// - API Documentation: [nwRtVrfMbr](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/nw:RtVrfMbr/)
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
//			_, err := nxos.LookupPhysicalInterfaceVrf(ctx, &nxos.LookupPhysicalInterfaceVrfArgs{
//				InterfaceId: "eth1/10",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPhysicalInterfaceVrf(ctx *pulumi.Context, args *LookupPhysicalInterfaceVrfArgs, opts ...pulumi.InvokeOption) (*LookupPhysicalInterfaceVrfResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPhysicalInterfaceVrfResult
	err := ctx.Invoke("nxos:index/getPhysicalInterfaceVrf:getPhysicalInterfaceVrf", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPhysicalInterfaceVrf.
type LookupPhysicalInterfaceVrfArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId string `pulumi:"interfaceId"`
}

// A collection of values returned by getPhysicalInterfaceVrf.
type LookupPhysicalInterfaceVrfResult struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId string `pulumi:"interfaceId"`
	// DN of VRF. For example: `sys/inst-VRF1`.
	VrfDn string `pulumi:"vrfDn"`
}

func LookupPhysicalInterfaceVrfOutput(ctx *pulumi.Context, args LookupPhysicalInterfaceVrfOutputArgs, opts ...pulumi.InvokeOption) LookupPhysicalInterfaceVrfResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPhysicalInterfaceVrfResult, error) {
			args := v.(LookupPhysicalInterfaceVrfArgs)
			r, err := LookupPhysicalInterfaceVrf(ctx, &args, opts...)
			var s LookupPhysicalInterfaceVrfResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPhysicalInterfaceVrfResultOutput)
}

// A collection of arguments for invoking getPhysicalInterfaceVrf.
type LookupPhysicalInterfaceVrfOutputArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
	InterfaceId pulumi.StringInput `pulumi:"interfaceId"`
}

func (LookupPhysicalInterfaceVrfOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPhysicalInterfaceVrfArgs)(nil)).Elem()
}

// A collection of values returned by getPhysicalInterfaceVrf.
type LookupPhysicalInterfaceVrfResultOutput struct{ *pulumi.OutputState }

func (LookupPhysicalInterfaceVrfResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPhysicalInterfaceVrfResult)(nil)).Elem()
}

func (o LookupPhysicalInterfaceVrfResultOutput) ToLookupPhysicalInterfaceVrfResultOutput() LookupPhysicalInterfaceVrfResultOutput {
	return o
}

func (o LookupPhysicalInterfaceVrfResultOutput) ToLookupPhysicalInterfaceVrfResultOutputWithContext(ctx context.Context) LookupPhysicalInterfaceVrfResultOutput {
	return o
}

func (o LookupPhysicalInterfaceVrfResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPhysicalInterfaceVrfResult] {
	return pulumix.Output[LookupPhysicalInterfaceVrfResult]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o LookupPhysicalInterfaceVrfResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPhysicalInterfaceVrfResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// The distinguished name of the object.
func (o LookupPhysicalInterfaceVrfResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPhysicalInterfaceVrfResult) string { return v.Id }).(pulumi.StringOutput)
}

// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
func (o LookupPhysicalInterfaceVrfResultOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPhysicalInterfaceVrfResult) string { return v.InterfaceId }).(pulumi.StringOutput)
}

// DN of VRF. For example: `sys/inst-VRF1`.
func (o LookupPhysicalInterfaceVrfResultOutput) VrfDn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPhysicalInterfaceVrfResult) string { return v.VrfDn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPhysicalInterfaceVrfResultOutput{})
}