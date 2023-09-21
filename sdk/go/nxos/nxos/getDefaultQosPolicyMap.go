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

func LookupDefaultQosPolicyMap(ctx *pulumi.Context, args *LookupDefaultQosPolicyMapArgs, opts ...pulumi.InvokeOption) (*LookupDefaultQosPolicyMapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDefaultQosPolicyMapResult
	err := ctx.Invoke("nxos:nxos/getDefaultQosPolicyMap:getDefaultQosPolicyMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDefaultQosPolicyMap.
type LookupDefaultQosPolicyMapArgs struct {
	Device *string `pulumi:"device"`
	Name   string  `pulumi:"name"`
}

// A collection of values returned by getDefaultQosPolicyMap.
type LookupDefaultQosPolicyMapResult struct {
	Device    *string `pulumi:"device"`
	Id        string  `pulumi:"id"`
	MatchType string  `pulumi:"matchType"`
	Name      string  `pulumi:"name"`
}

func LookupDefaultQosPolicyMapOutput(ctx *pulumi.Context, args LookupDefaultQosPolicyMapOutputArgs, opts ...pulumi.InvokeOption) LookupDefaultQosPolicyMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDefaultQosPolicyMapResult, error) {
			args := v.(LookupDefaultQosPolicyMapArgs)
			r, err := LookupDefaultQosPolicyMap(ctx, &args, opts...)
			var s LookupDefaultQosPolicyMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDefaultQosPolicyMapResultOutput)
}

// A collection of arguments for invoking getDefaultQosPolicyMap.
type LookupDefaultQosPolicyMapOutputArgs struct {
	Device pulumi.StringPtrInput `pulumi:"device"`
	Name   pulumi.StringInput    `pulumi:"name"`
}

func (LookupDefaultQosPolicyMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultQosPolicyMapArgs)(nil)).Elem()
}

// A collection of values returned by getDefaultQosPolicyMap.
type LookupDefaultQosPolicyMapResultOutput struct{ *pulumi.OutputState }

func (LookupDefaultQosPolicyMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultQosPolicyMapResult)(nil)).Elem()
}

func (o LookupDefaultQosPolicyMapResultOutput) ToLookupDefaultQosPolicyMapResultOutput() LookupDefaultQosPolicyMapResultOutput {
	return o
}

func (o LookupDefaultQosPolicyMapResultOutput) ToLookupDefaultQosPolicyMapResultOutputWithContext(ctx context.Context) LookupDefaultQosPolicyMapResultOutput {
	return o
}

func (o LookupDefaultQosPolicyMapResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDefaultQosPolicyMapResult] {
	return pulumix.Output[LookupDefaultQosPolicyMapResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupDefaultQosPolicyMapResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDefaultQosPolicyMapResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupDefaultQosPolicyMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosPolicyMapResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDefaultQosPolicyMapResultOutput) MatchType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosPolicyMapResult) string { return v.MatchType }).(pulumi.StringOutput)
}

func (o LookupDefaultQosPolicyMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosPolicyMapResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDefaultQosPolicyMapResultOutput{})
}