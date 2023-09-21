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

func LookupQueuingQosPolicySystemOutPolicyMap(ctx *pulumi.Context, args *LookupQueuingQosPolicySystemOutPolicyMapArgs, opts ...pulumi.InvokeOption) (*LookupQueuingQosPolicySystemOutPolicyMapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupQueuingQosPolicySystemOutPolicyMapResult
	err := ctx.Invoke("nxos:nxos/getQueuingQosPolicySystemOutPolicyMap:getQueuingQosPolicySystemOutPolicyMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getQueuingQosPolicySystemOutPolicyMap.
type LookupQueuingQosPolicySystemOutPolicyMapArgs struct {
	Device *string `pulumi:"device"`
}

// A collection of values returned by getQueuingQosPolicySystemOutPolicyMap.
type LookupQueuingQosPolicySystemOutPolicyMapResult struct {
	Device        *string `pulumi:"device"`
	Id            string  `pulumi:"id"`
	PolicyMapName string  `pulumi:"policyMapName"`
}

func LookupQueuingQosPolicySystemOutPolicyMapOutput(ctx *pulumi.Context, args LookupQueuingQosPolicySystemOutPolicyMapOutputArgs, opts ...pulumi.InvokeOption) LookupQueuingQosPolicySystemOutPolicyMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueuingQosPolicySystemOutPolicyMapResult, error) {
			args := v.(LookupQueuingQosPolicySystemOutPolicyMapArgs)
			r, err := LookupQueuingQosPolicySystemOutPolicyMap(ctx, &args, opts...)
			var s LookupQueuingQosPolicySystemOutPolicyMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueuingQosPolicySystemOutPolicyMapResultOutput)
}

// A collection of arguments for invoking getQueuingQosPolicySystemOutPolicyMap.
type LookupQueuingQosPolicySystemOutPolicyMapOutputArgs struct {
	Device pulumi.StringPtrInput `pulumi:"device"`
}

func (LookupQueuingQosPolicySystemOutPolicyMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueuingQosPolicySystemOutPolicyMapArgs)(nil)).Elem()
}

// A collection of values returned by getQueuingQosPolicySystemOutPolicyMap.
type LookupQueuingQosPolicySystemOutPolicyMapResultOutput struct{ *pulumi.OutputState }

func (LookupQueuingQosPolicySystemOutPolicyMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueuingQosPolicySystemOutPolicyMapResult)(nil)).Elem()
}

func (o LookupQueuingQosPolicySystemOutPolicyMapResultOutput) ToLookupQueuingQosPolicySystemOutPolicyMapResultOutput() LookupQueuingQosPolicySystemOutPolicyMapResultOutput {
	return o
}

func (o LookupQueuingQosPolicySystemOutPolicyMapResultOutput) ToLookupQueuingQosPolicySystemOutPolicyMapResultOutputWithContext(ctx context.Context) LookupQueuingQosPolicySystemOutPolicyMapResultOutput {
	return o
}

func (o LookupQueuingQosPolicySystemOutPolicyMapResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupQueuingQosPolicySystemOutPolicyMapResult] {
	return pulumix.Output[LookupQueuingQosPolicySystemOutPolicyMapResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupQueuingQosPolicySystemOutPolicyMapResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueuingQosPolicySystemOutPolicyMapResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupQueuingQosPolicySystemOutPolicyMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueuingQosPolicySystemOutPolicyMapResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupQueuingQosPolicySystemOutPolicyMapResultOutput) PolicyMapName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueuingQosPolicySystemOutPolicyMapResult) string { return v.PolicyMapName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueuingQosPolicySystemOutPolicyMapResultOutput{})
}
