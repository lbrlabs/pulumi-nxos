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

func LookupQueuingQosPolicyMapMatchClassMap(ctx *pulumi.Context, args *LookupQueuingQosPolicyMapMatchClassMapArgs, opts ...pulumi.InvokeOption) (*LookupQueuingQosPolicyMapMatchClassMapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupQueuingQosPolicyMapMatchClassMapResult
	err := ctx.Invoke("nxos:nxos/getQueuingQosPolicyMapMatchClassMap:getQueuingQosPolicyMapMatchClassMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQueuingQosPolicyMapMatchClassMap.
type LookupQueuingQosPolicyMapMatchClassMapArgs struct {
	Device        *string `pulumi:"device"`
	Name          string  `pulumi:"name"`
	PolicyMapName string  `pulumi:"policyMapName"`
}

// A collection of values returned by getQueuingQosPolicyMapMatchClassMap.
type LookupQueuingQosPolicyMapMatchClassMapResult struct {
	Device        *string `pulumi:"device"`
	Id            string  `pulumi:"id"`
	Name          string  `pulumi:"name"`
	PolicyMapName string  `pulumi:"policyMapName"`
}

func LookupQueuingQosPolicyMapMatchClassMapOutput(ctx *pulumi.Context, args LookupQueuingQosPolicyMapMatchClassMapOutputArgs, opts ...pulumi.InvokeOption) LookupQueuingQosPolicyMapMatchClassMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueuingQosPolicyMapMatchClassMapResult, error) {
			args := v.(LookupQueuingQosPolicyMapMatchClassMapArgs)
			r, err := LookupQueuingQosPolicyMapMatchClassMap(ctx, &args, opts...)
			var s LookupQueuingQosPolicyMapMatchClassMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueuingQosPolicyMapMatchClassMapResultOutput)
}

// A collection of arguments for invoking getQueuingQosPolicyMapMatchClassMap.
type LookupQueuingQosPolicyMapMatchClassMapOutputArgs struct {
	Device        pulumi.StringPtrInput `pulumi:"device"`
	Name          pulumi.StringInput    `pulumi:"name"`
	PolicyMapName pulumi.StringInput    `pulumi:"policyMapName"`
}

func (LookupQueuingQosPolicyMapMatchClassMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueuingQosPolicyMapMatchClassMapArgs)(nil)).Elem()
}

// A collection of values returned by getQueuingQosPolicyMapMatchClassMap.
type LookupQueuingQosPolicyMapMatchClassMapResultOutput struct{ *pulumi.OutputState }

func (LookupQueuingQosPolicyMapMatchClassMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueuingQosPolicyMapMatchClassMapResult)(nil)).Elem()
}

func (o LookupQueuingQosPolicyMapMatchClassMapResultOutput) ToLookupQueuingQosPolicyMapMatchClassMapResultOutput() LookupQueuingQosPolicyMapMatchClassMapResultOutput {
	return o
}

func (o LookupQueuingQosPolicyMapMatchClassMapResultOutput) ToLookupQueuingQosPolicyMapMatchClassMapResultOutputWithContext(ctx context.Context) LookupQueuingQosPolicyMapMatchClassMapResultOutput {
	return o
}

func (o LookupQueuingQosPolicyMapMatchClassMapResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupQueuingQosPolicyMapMatchClassMapResult] {
	return pulumix.Output[LookupQueuingQosPolicyMapMatchClassMapResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupQueuingQosPolicyMapMatchClassMapResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueuingQosPolicyMapMatchClassMapResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupQueuingQosPolicyMapMatchClassMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueuingQosPolicyMapMatchClassMapResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupQueuingQosPolicyMapMatchClassMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueuingQosPolicyMapMatchClassMapResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupQueuingQosPolicyMapMatchClassMapResultOutput) PolicyMapName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueuingQosPolicyMapMatchClassMapResult) string { return v.PolicyMapName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueuingQosPolicyMapMatchClassMapResultOutput{})
}