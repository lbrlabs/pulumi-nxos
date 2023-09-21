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

func LookupQueuingQosPolicyMapMatchClassMapPriority(ctx *pulumi.Context, args *LookupQueuingQosPolicyMapMatchClassMapPriorityArgs, opts ...pulumi.InvokeOption) (*LookupQueuingQosPolicyMapMatchClassMapPriorityResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupQueuingQosPolicyMapMatchClassMapPriorityResult
	err := ctx.Invoke("nxos:nxos/getQueuingQosPolicyMapMatchClassMapPriority:getQueuingQosPolicyMapMatchClassMapPriority", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQueuingQosPolicyMapMatchClassMapPriority.
type LookupQueuingQosPolicyMapMatchClassMapPriorityArgs struct {
	ClassMapName  string  `pulumi:"classMapName"`
	Device        *string `pulumi:"device"`
	PolicyMapName string  `pulumi:"policyMapName"`
}

// A collection of values returned by getQueuingQosPolicyMapMatchClassMapPriority.
type LookupQueuingQosPolicyMapMatchClassMapPriorityResult struct {
	ClassMapName  string  `pulumi:"classMapName"`
	Device        *string `pulumi:"device"`
	Id            string  `pulumi:"id"`
	Level         int     `pulumi:"level"`
	PolicyMapName string  `pulumi:"policyMapName"`
}

func LookupQueuingQosPolicyMapMatchClassMapPriorityOutput(ctx *pulumi.Context, args LookupQueuingQosPolicyMapMatchClassMapPriorityOutputArgs, opts ...pulumi.InvokeOption) LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueuingQosPolicyMapMatchClassMapPriorityResult, error) {
			args := v.(LookupQueuingQosPolicyMapMatchClassMapPriorityArgs)
			r, err := LookupQueuingQosPolicyMapMatchClassMapPriority(ctx, &args, opts...)
			var s LookupQueuingQosPolicyMapMatchClassMapPriorityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput)
}

// A collection of arguments for invoking getQueuingQosPolicyMapMatchClassMapPriority.
type LookupQueuingQosPolicyMapMatchClassMapPriorityOutputArgs struct {
	ClassMapName  pulumi.StringInput    `pulumi:"classMapName"`
	Device        pulumi.StringPtrInput `pulumi:"device"`
	PolicyMapName pulumi.StringInput    `pulumi:"policyMapName"`
}

func (LookupQueuingQosPolicyMapMatchClassMapPriorityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueuingQosPolicyMapMatchClassMapPriorityArgs)(nil)).Elem()
}

// A collection of values returned by getQueuingQosPolicyMapMatchClassMapPriority.
type LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput struct{ *pulumi.OutputState }

func (LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueuingQosPolicyMapMatchClassMapPriorityResult)(nil)).Elem()
}

func (o LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput) ToLookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput() LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput {
	return o
}

func (o LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput) ToLookupQueuingQosPolicyMapMatchClassMapPriorityResultOutputWithContext(ctx context.Context) LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput {
	return o
}

func (o LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupQueuingQosPolicyMapMatchClassMapPriorityResult] {
	return pulumix.Output[LookupQueuingQosPolicyMapMatchClassMapPriorityResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput) ClassMapName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueuingQosPolicyMapMatchClassMapPriorityResult) string { return v.ClassMapName }).(pulumi.StringOutput)
}

func (o LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueuingQosPolicyMapMatchClassMapPriorityResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueuingQosPolicyMapMatchClassMapPriorityResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput) Level() pulumi.IntOutput {
	return o.ApplyT(func(v LookupQueuingQosPolicyMapMatchClassMapPriorityResult) int { return v.Level }).(pulumi.IntOutput)
}

func (o LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput) PolicyMapName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueuingQosPolicyMapMatchClassMapPriorityResult) string { return v.PolicyMapName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueuingQosPolicyMapMatchClassMapPriorityResultOutput{})
}