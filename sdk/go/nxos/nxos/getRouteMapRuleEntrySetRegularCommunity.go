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

func LookupRouteMapRuleEntrySetRegularCommunity(ctx *pulumi.Context, args *LookupRouteMapRuleEntrySetRegularCommunityArgs, opts ...pulumi.InvokeOption) (*LookupRouteMapRuleEntrySetRegularCommunityResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRouteMapRuleEntrySetRegularCommunityResult
	err := ctx.Invoke("nxos:nxos/getRouteMapRuleEntrySetRegularCommunity:getRouteMapRuleEntrySetRegularCommunity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteMapRuleEntrySetRegularCommunity.
type LookupRouteMapRuleEntrySetRegularCommunityArgs struct {
	Device   *string `pulumi:"device"`
	Order    int     `pulumi:"order"`
	RuleName string  `pulumi:"ruleName"`
}

// A collection of values returned by getRouteMapRuleEntrySetRegularCommunity.
type LookupRouteMapRuleEntrySetRegularCommunityResult struct {
	Additive    string  `pulumi:"additive"`
	Device      *string `pulumi:"device"`
	Id          string  `pulumi:"id"`
	NoCommunity string  `pulumi:"noCommunity"`
	Order       int     `pulumi:"order"`
	RuleName    string  `pulumi:"ruleName"`
	SetCriteria string  `pulumi:"setCriteria"`
}

func LookupRouteMapRuleEntrySetRegularCommunityOutput(ctx *pulumi.Context, args LookupRouteMapRuleEntrySetRegularCommunityOutputArgs, opts ...pulumi.InvokeOption) LookupRouteMapRuleEntrySetRegularCommunityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteMapRuleEntrySetRegularCommunityResult, error) {
			args := v.(LookupRouteMapRuleEntrySetRegularCommunityArgs)
			r, err := LookupRouteMapRuleEntrySetRegularCommunity(ctx, &args, opts...)
			var s LookupRouteMapRuleEntrySetRegularCommunityResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteMapRuleEntrySetRegularCommunityResultOutput)
}

// A collection of arguments for invoking getRouteMapRuleEntrySetRegularCommunity.
type LookupRouteMapRuleEntrySetRegularCommunityOutputArgs struct {
	Device   pulumi.StringPtrInput `pulumi:"device"`
	Order    pulumi.IntInput       `pulumi:"order"`
	RuleName pulumi.StringInput    `pulumi:"ruleName"`
}

func (LookupRouteMapRuleEntrySetRegularCommunityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteMapRuleEntrySetRegularCommunityArgs)(nil)).Elem()
}

// A collection of values returned by getRouteMapRuleEntrySetRegularCommunity.
type LookupRouteMapRuleEntrySetRegularCommunityResultOutput struct{ *pulumi.OutputState }

func (LookupRouteMapRuleEntrySetRegularCommunityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteMapRuleEntrySetRegularCommunityResult)(nil)).Elem()
}

func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) ToLookupRouteMapRuleEntrySetRegularCommunityResultOutput() LookupRouteMapRuleEntrySetRegularCommunityResultOutput {
	return o
}

func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) ToLookupRouteMapRuleEntrySetRegularCommunityResultOutputWithContext(ctx context.Context) LookupRouteMapRuleEntrySetRegularCommunityResultOutput {
	return o
}

func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRouteMapRuleEntrySetRegularCommunityResult] {
	return pulumix.Output[LookupRouteMapRuleEntrySetRegularCommunityResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) Additive() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapRuleEntrySetRegularCommunityResult) string { return v.Additive }).(pulumi.StringOutput)
}

func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteMapRuleEntrySetRegularCommunityResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapRuleEntrySetRegularCommunityResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) NoCommunity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapRuleEntrySetRegularCommunityResult) string { return v.NoCommunity }).(pulumi.StringOutput)
}

func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouteMapRuleEntrySetRegularCommunityResult) int { return v.Order }).(pulumi.IntOutput)
}

func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapRuleEntrySetRegularCommunityResult) string { return v.RuleName }).(pulumi.StringOutput)
}

func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) SetCriteria() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapRuleEntrySetRegularCommunityResult) string { return v.SetCriteria }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteMapRuleEntrySetRegularCommunityResultOutput{})
}
