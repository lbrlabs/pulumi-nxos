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

func LookupDefaultQosPolicyInterfaceInPolicyMap(ctx *pulumi.Context, args *LookupDefaultQosPolicyInterfaceInPolicyMapArgs, opts ...pulumi.InvokeOption) (*LookupDefaultQosPolicyInterfaceInPolicyMapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDefaultQosPolicyInterfaceInPolicyMapResult
	err := ctx.Invoke("nxos:nxos/getDefaultQosPolicyInterfaceInPolicyMap:getDefaultQosPolicyInterfaceInPolicyMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDefaultQosPolicyInterfaceInPolicyMap.
type LookupDefaultQosPolicyInterfaceInPolicyMapArgs struct {
	Device      *string `pulumi:"device"`
	InterfaceId string  `pulumi:"interfaceId"`
}

// A collection of values returned by getDefaultQosPolicyInterfaceInPolicyMap.
type LookupDefaultQosPolicyInterfaceInPolicyMapResult struct {
	Device        *string `pulumi:"device"`
	Id            string  `pulumi:"id"`
	InterfaceId   string  `pulumi:"interfaceId"`
	PolicyMapName string  `pulumi:"policyMapName"`
}

func LookupDefaultQosPolicyInterfaceInPolicyMapOutput(ctx *pulumi.Context, args LookupDefaultQosPolicyInterfaceInPolicyMapOutputArgs, opts ...pulumi.InvokeOption) LookupDefaultQosPolicyInterfaceInPolicyMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDefaultQosPolicyInterfaceInPolicyMapResult, error) {
			args := v.(LookupDefaultQosPolicyInterfaceInPolicyMapArgs)
			r, err := LookupDefaultQosPolicyInterfaceInPolicyMap(ctx, &args, opts...)
			var s LookupDefaultQosPolicyInterfaceInPolicyMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDefaultQosPolicyInterfaceInPolicyMapResultOutput)
}

// A collection of arguments for invoking getDefaultQosPolicyInterfaceInPolicyMap.
type LookupDefaultQosPolicyInterfaceInPolicyMapOutputArgs struct {
	Device      pulumi.StringPtrInput `pulumi:"device"`
	InterfaceId pulumi.StringInput    `pulumi:"interfaceId"`
}

func (LookupDefaultQosPolicyInterfaceInPolicyMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultQosPolicyInterfaceInPolicyMapArgs)(nil)).Elem()
}

// A collection of values returned by getDefaultQosPolicyInterfaceInPolicyMap.
type LookupDefaultQosPolicyInterfaceInPolicyMapResultOutput struct{ *pulumi.OutputState }

func (LookupDefaultQosPolicyInterfaceInPolicyMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultQosPolicyInterfaceInPolicyMapResult)(nil)).Elem()
}

func (o LookupDefaultQosPolicyInterfaceInPolicyMapResultOutput) ToLookupDefaultQosPolicyInterfaceInPolicyMapResultOutput() LookupDefaultQosPolicyInterfaceInPolicyMapResultOutput {
	return o
}

func (o LookupDefaultQosPolicyInterfaceInPolicyMapResultOutput) ToLookupDefaultQosPolicyInterfaceInPolicyMapResultOutputWithContext(ctx context.Context) LookupDefaultQosPolicyInterfaceInPolicyMapResultOutput {
	return o
}

func (o LookupDefaultQosPolicyInterfaceInPolicyMapResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDefaultQosPolicyInterfaceInPolicyMapResult] {
	return pulumix.Output[LookupDefaultQosPolicyInterfaceInPolicyMapResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupDefaultQosPolicyInterfaceInPolicyMapResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDefaultQosPolicyInterfaceInPolicyMapResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupDefaultQosPolicyInterfaceInPolicyMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosPolicyInterfaceInPolicyMapResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDefaultQosPolicyInterfaceInPolicyMapResultOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosPolicyInterfaceInPolicyMapResult) string { return v.InterfaceId }).(pulumi.StringOutput)
}

func (o LookupDefaultQosPolicyInterfaceInPolicyMapResultOutput) PolicyMapName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosPolicyInterfaceInPolicyMapResult) string { return v.PolicyMapName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDefaultQosPolicyInterfaceInPolicyMapResultOutput{})
}