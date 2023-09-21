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

func LookupPimStaticRpGroupList(ctx *pulumi.Context, args *LookupPimStaticRpGroupListArgs, opts ...pulumi.InvokeOption) (*LookupPimStaticRpGroupListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPimStaticRpGroupListResult
	err := ctx.Invoke("nxos:nxos/getPimStaticRpGroupList:getPimStaticRpGroupList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPimStaticRpGroupList.
type LookupPimStaticRpGroupListArgs struct {
	Address   string  `pulumi:"address"`
	Device    *string `pulumi:"device"`
	RpAddress string  `pulumi:"rpAddress"`
	VrfName   string  `pulumi:"vrfName"`
}

// A collection of values returned by getPimStaticRpGroupList.
type LookupPimStaticRpGroupListResult struct {
	Address   string  `pulumi:"address"`
	Bidir     bool    `pulumi:"bidir"`
	Device    *string `pulumi:"device"`
	Id        string  `pulumi:"id"`
	Override  bool    `pulumi:"override"`
	RpAddress string  `pulumi:"rpAddress"`
	VrfName   string  `pulumi:"vrfName"`
}

func LookupPimStaticRpGroupListOutput(ctx *pulumi.Context, args LookupPimStaticRpGroupListOutputArgs, opts ...pulumi.InvokeOption) LookupPimStaticRpGroupListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPimStaticRpGroupListResult, error) {
			args := v.(LookupPimStaticRpGroupListArgs)
			r, err := LookupPimStaticRpGroupList(ctx, &args, opts...)
			var s LookupPimStaticRpGroupListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPimStaticRpGroupListResultOutput)
}

// A collection of arguments for invoking getPimStaticRpGroupList.
type LookupPimStaticRpGroupListOutputArgs struct {
	Address   pulumi.StringInput    `pulumi:"address"`
	Device    pulumi.StringPtrInput `pulumi:"device"`
	RpAddress pulumi.StringInput    `pulumi:"rpAddress"`
	VrfName   pulumi.StringInput    `pulumi:"vrfName"`
}

func (LookupPimStaticRpGroupListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPimStaticRpGroupListArgs)(nil)).Elem()
}

// A collection of values returned by getPimStaticRpGroupList.
type LookupPimStaticRpGroupListResultOutput struct{ *pulumi.OutputState }

func (LookupPimStaticRpGroupListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPimStaticRpGroupListResult)(nil)).Elem()
}

func (o LookupPimStaticRpGroupListResultOutput) ToLookupPimStaticRpGroupListResultOutput() LookupPimStaticRpGroupListResultOutput {
	return o
}

func (o LookupPimStaticRpGroupListResultOutput) ToLookupPimStaticRpGroupListResultOutputWithContext(ctx context.Context) LookupPimStaticRpGroupListResultOutput {
	return o
}

func (o LookupPimStaticRpGroupListResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPimStaticRpGroupListResult] {
	return pulumix.Output[LookupPimStaticRpGroupListResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupPimStaticRpGroupListResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimStaticRpGroupListResult) string { return v.Address }).(pulumi.StringOutput)
}

func (o LookupPimStaticRpGroupListResultOutput) Bidir() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPimStaticRpGroupListResult) bool { return v.Bidir }).(pulumi.BoolOutput)
}

func (o LookupPimStaticRpGroupListResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPimStaticRpGroupListResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupPimStaticRpGroupListResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimStaticRpGroupListResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPimStaticRpGroupListResultOutput) Override() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPimStaticRpGroupListResult) bool { return v.Override }).(pulumi.BoolOutput)
}

func (o LookupPimStaticRpGroupListResultOutput) RpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimStaticRpGroupListResult) string { return v.RpAddress }).(pulumi.StringOutput)
}

func (o LookupPimStaticRpGroupListResultOutput) VrfName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimStaticRpGroupListResult) string { return v.VrfName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPimStaticRpGroupListResultOutput{})
}
