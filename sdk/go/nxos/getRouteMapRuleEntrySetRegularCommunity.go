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

// This data source can read a Set Community configuration in a Route-Map Rule Entry.
//
// - API Documentation: [rtmapSetRegComm](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:SetRegComm/)
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
//			_, err := nxos.LookupRouteMapRuleEntrySetRegularCommunity(ctx, &nxos.LookupRouteMapRuleEntrySetRegularCommunityArgs{
//				Order:    10,
//				RuleName: "RULE1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRouteMapRuleEntrySetRegularCommunity(ctx *pulumi.Context, args *LookupRouteMapRuleEntrySetRegularCommunityArgs, opts ...pulumi.InvokeOption) (*LookupRouteMapRuleEntrySetRegularCommunityResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRouteMapRuleEntrySetRegularCommunityResult
	err := ctx.Invoke("nxos:index/getRouteMapRuleEntrySetRegularCommunity:getRouteMapRuleEntrySetRegularCommunity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteMapRuleEntrySetRegularCommunity.
type LookupRouteMapRuleEntrySetRegularCommunityArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Route-Map Rule Entry order.
	Order int `pulumi:"order"`
	// Route Map rule name.
	RuleName string `pulumi:"ruleName"`
}

// A collection of values returned by getRouteMapRuleEntrySetRegularCommunity.
type LookupRouteMapRuleEntrySetRegularCommunityResult struct {
	// Option to add to an existing community.
	Additive string `pulumi:"additive"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// Option to have no community attribute.
	NoCommunity string `pulumi:"noCommunity"`
	// Route-Map Rule Entry order.
	Order int `pulumi:"order"`
	// Route Map rule name.
	RuleName string `pulumi:"ruleName"`
	// Operation on the current community.
	SetCriteria string `pulumi:"setCriteria"`
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
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// Route-Map Rule Entry order.
	Order pulumi.IntInput `pulumi:"order"`
	// Route Map rule name.
	RuleName pulumi.StringInput `pulumi:"ruleName"`
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

// Option to add to an existing community.
func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) Additive() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapRuleEntrySetRegularCommunityResult) string { return v.Additive }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteMapRuleEntrySetRegularCommunityResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// The distinguished name of the object.
func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapRuleEntrySetRegularCommunityResult) string { return v.Id }).(pulumi.StringOutput)
}

// Option to have no community attribute.
func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) NoCommunity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapRuleEntrySetRegularCommunityResult) string { return v.NoCommunity }).(pulumi.StringOutput)
}

// Route-Map Rule Entry order.
func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRouteMapRuleEntrySetRegularCommunityResult) int { return v.Order }).(pulumi.IntOutput)
}

// Route Map rule name.
func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) RuleName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapRuleEntrySetRegularCommunityResult) string { return v.RuleName }).(pulumi.StringOutput)
}

// Operation on the current community.
func (o LookupRouteMapRuleEntrySetRegularCommunityResultOutput) SetCriteria() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapRuleEntrySetRegularCommunityResult) string { return v.SetCriteria }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteMapRuleEntrySetRegularCommunityResultOutput{})
}