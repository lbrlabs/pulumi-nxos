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

func LookupPortChannelInterfaceMember(ctx *pulumi.Context, args *LookupPortChannelInterfaceMemberArgs, opts ...pulumi.InvokeOption) (*LookupPortChannelInterfaceMemberResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPortChannelInterfaceMemberResult
	err := ctx.Invoke("nxos:nxos/getPortChannelInterfaceMember:getPortChannelInterfaceMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPortChannelInterfaceMember.
type LookupPortChannelInterfaceMemberArgs struct {
	Device      *string `pulumi:"device"`
	InterfaceDn string  `pulumi:"interfaceDn"`
	InterfaceId string  `pulumi:"interfaceId"`
}

// A collection of values returned by getPortChannelInterfaceMember.
type LookupPortChannelInterfaceMemberResult struct {
	Device      *string `pulumi:"device"`
	Id          string  `pulumi:"id"`
	InterfaceDn string  `pulumi:"interfaceDn"`
	InterfaceId string  `pulumi:"interfaceId"`
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
	Device      pulumi.StringPtrInput `pulumi:"device"`
	InterfaceDn pulumi.StringInput    `pulumi:"interfaceDn"`
	InterfaceId pulumi.StringInput    `pulumi:"interfaceId"`
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

func (o LookupPortChannelInterfaceMemberResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPortChannelInterfaceMemberResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupPortChannelInterfaceMemberResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortChannelInterfaceMemberResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPortChannelInterfaceMemberResultOutput) InterfaceDn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortChannelInterfaceMemberResult) string { return v.InterfaceDn }).(pulumi.StringOutput)
}

func (o LookupPortChannelInterfaceMemberResultOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPortChannelInterfaceMemberResult) string { return v.InterfaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPortChannelInterfaceMemberResultOutput{})
}