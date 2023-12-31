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

// This data source can read the configuration of a port-channel interface member.
//
// - API Documentation: [pcRsMbrIfs](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Interfaces/pc:RsMbrIfs/)
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
//			_, err := nxos.LookupPortChannelInterfaceMember(ctx, &nxos.LookupPortChannelInterfaceMemberArgs{
//				InterfaceDn: "sys/intf/phys-[eth1/11]",
//				InterfaceId: "po1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPortChannelInterfaceMember(ctx *pulumi.Context, args *LookupPortChannelInterfaceMemberArgs, opts ...pulumi.InvokeOption) (*LookupPortChannelInterfaceMemberResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPortChannelInterfaceMemberResult
	err := ctx.Invoke("nxos:index/getPortChannelInterfaceMember:getPortChannelInterfaceMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPortChannelInterfaceMember.
type LookupPortChannelInterfaceMemberArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// DN of interface. For example: `sys/intf/phys-[eth1/1]`.
	InterfaceDn string `pulumi:"interfaceDn"`
	// Must match first field in the output of `show intf brief`. Example: `po1`.
	InterfaceId string `pulumi:"interfaceId"`
}

// A collection of values returned by getPortChannelInterfaceMember.
type LookupPortChannelInterfaceMemberResult struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// DN of interface. For example: `sys/intf/phys-[eth1/1]`.
	InterfaceDn string `pulumi:"interfaceDn"`
	// Must match first field in the output of `show intf brief`. Example: `po1`.
	InterfaceId string `pulumi:"interfaceId"`
}

func LookupPortChannelInterfaceMemberOutput(ctx *pulumi.Context, args LookupPortChannelInterfaceMemberOutputArgs, opts ...pulumi.InvokeOption) LookupPortChannelInterfaceMemberResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPortChannelInterfaceMemberResult, error) {
			args := v.(LookupPortChannelInterfaceMemberArgs)
			r, err := LookupPortChannelInterfaceMember(ctx, &args, opts...)
			var s LookupPortChannelInterfaceMemberResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPortChannelInterfaceMemberResultOutput)
}

// A collection of arguments for invoking getPortChannelInterfaceMember.
type LookupPortChannelInterfaceMemberOutputArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// DN of interface. For example: `sys/intf/phys-[eth1/1]`.
	InterfaceDn pulumi.StringInput `pulumi:"interfaceDn"`
	// Must match first field in the output of `show intf brief`. Example: `po1`.
	InterfaceId pulumi.StringInput `pulumi:"interfaceId"`
}

func (LookupPortChannelInterfaceMemberOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPortChannelInterfaceMemberArgs)(nil)).Elem()
}

// A collection of values returned by getPortChannelInterfaceMember.
type LookupPortChannelInterfaceMemberResultOutput struct{ *pulumi.OutputState }

func (LookupPortChannelInterfaceMemberResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPortChannelInterfaceMemberResult)(nil)).Elem()
}

func (o LookupPortChannelInterfaceMemberResultOutput) ToLookupPortChannelInterfaceMemberResultOutput() LookupPortChannelInterfaceMemberResultOutput {
	return o
}

func (o LookupPortChannelInterfaceMemberResultOutput) ToLookupPortChannelInterfaceMemberResultOutputWithContext(ctx context.Context) LookupPortChannelInterfaceMemberResultOutput {
	return o
}

func (o LookupPortChannelInterfaceMemberResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPortChannelInterfaceMemberResult] {
	return pulumix.Output[LookupPortChannelInterfaceMemberResult]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o LookupPortChannelInterfaceMemberResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortChannelInterfaceMemberResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// The distinguished name of the object.
func (o LookupPortChannelInterfaceMemberResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortChannelInterfaceMemberResult) string { return v.Id }).(pulumi.StringOutput)
}

// DN of interface. For example: `sys/intf/phys-[eth1/1]`.
func (o LookupPortChannelInterfaceMemberResultOutput) InterfaceDn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortChannelInterfaceMemberResult) string { return v.InterfaceDn }).(pulumi.StringOutput)
}

// Must match first field in the output of `show intf brief`. Example: `po1`.
func (o LookupPortChannelInterfaceMemberResultOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortChannelInterfaceMemberResult) string { return v.InterfaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPortChannelInterfaceMemberResultOutput{})
}
