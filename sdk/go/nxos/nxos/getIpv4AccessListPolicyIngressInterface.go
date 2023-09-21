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

func LookupIpv4AccessListPolicyIngressInterface(ctx *pulumi.Context, args *LookupIpv4AccessListPolicyIngressInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupIpv4AccessListPolicyIngressInterfaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpv4AccessListPolicyIngressInterfaceResult
	err := ctx.Invoke("nxos:nxos/getIpv4AccessListPolicyIngressInterface:getIpv4AccessListPolicyIngressInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpv4AccessListPolicyIngressInterface.
type LookupIpv4AccessListPolicyIngressInterfaceArgs struct {
	Device      *string `pulumi:"device"`
	InterfaceId string  `pulumi:"interfaceId"`
}

// A collection of values returned by getIpv4AccessListPolicyIngressInterface.
type LookupIpv4AccessListPolicyIngressInterfaceResult struct {
	AccessListName string  `pulumi:"accessListName"`
	Device         *string `pulumi:"device"`
	Id             string  `pulumi:"id"`
	InterfaceId    string  `pulumi:"interfaceId"`
}

func LookupIpv4AccessListPolicyIngressInterfaceOutput(ctx *pulumi.Context, args LookupIpv4AccessListPolicyIngressInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupIpv4AccessListPolicyIngressInterfaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpv4AccessListPolicyIngressInterfaceResult, error) {
			args := v.(LookupIpv4AccessListPolicyIngressInterfaceArgs)
			r, err := LookupIpv4AccessListPolicyIngressInterface(ctx, &args, opts...)
			var s LookupIpv4AccessListPolicyIngressInterfaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpv4AccessListPolicyIngressInterfaceResultOutput)
}

// A collection of arguments for invoking getIpv4AccessListPolicyIngressInterface.
type LookupIpv4AccessListPolicyIngressInterfaceOutputArgs struct {
	Device      pulumi.StringPtrInput `pulumi:"device"`
	InterfaceId pulumi.StringInput    `pulumi:"interfaceId"`
}

func (LookupIpv4AccessListPolicyIngressInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpv4AccessListPolicyIngressInterfaceArgs)(nil)).Elem()
}

// A collection of values returned by getIpv4AccessListPolicyIngressInterface.
type LookupIpv4AccessListPolicyIngressInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupIpv4AccessListPolicyIngressInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpv4AccessListPolicyIngressInterfaceResult)(nil)).Elem()
}

func (o LookupIpv4AccessListPolicyIngressInterfaceResultOutput) ToLookupIpv4AccessListPolicyIngressInterfaceResultOutput() LookupIpv4AccessListPolicyIngressInterfaceResultOutput {
	return o
}

func (o LookupIpv4AccessListPolicyIngressInterfaceResultOutput) ToLookupIpv4AccessListPolicyIngressInterfaceResultOutputWithContext(ctx context.Context) LookupIpv4AccessListPolicyIngressInterfaceResultOutput {
	return o
}

func (o LookupIpv4AccessListPolicyIngressInterfaceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIpv4AccessListPolicyIngressInterfaceResult] {
	return pulumix.Output[LookupIpv4AccessListPolicyIngressInterfaceResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupIpv4AccessListPolicyIngressInterfaceResultOutput) AccessListName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListPolicyIngressInterfaceResult) string { return v.AccessListName }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListPolicyIngressInterfaceResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpv4AccessListPolicyIngressInterfaceResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupIpv4AccessListPolicyIngressInterfaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListPolicyIngressInterfaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListPolicyIngressInterfaceResultOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListPolicyIngressInterfaceResult) string { return v.InterfaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpv4AccessListPolicyIngressInterfaceResultOutput{})
}
