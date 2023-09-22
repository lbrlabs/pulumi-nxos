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

// This data source can read a Route-Map Rule configuration.
//
// - API Documentation: [rtmapRule](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:Rule/)
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
//			_, err := nxos.LookupRouteMapRule(ctx, &nxos.LookupRouteMapRuleArgs{
//				Name: "RULE1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRouteMapRule(ctx *pulumi.Context, args *LookupRouteMapRuleArgs, opts ...pulumi.InvokeOption) (*LookupRouteMapRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRouteMapRuleResult
	err := ctx.Invoke("nxos:index/getRouteMapRule:getRouteMapRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRouteMapRule.
type LookupRouteMapRuleArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Route-Map Rule name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getRouteMapRule.
type LookupRouteMapRuleResult struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// Route-Map Rule name.
	Name string `pulumi:"name"`
}

func LookupRouteMapRuleOutput(ctx *pulumi.Context, args LookupRouteMapRuleOutputArgs, opts ...pulumi.InvokeOption) LookupRouteMapRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRouteMapRuleResult, error) {
			args := v.(LookupRouteMapRuleArgs)
			r, err := LookupRouteMapRule(ctx, &args, opts...)
			var s LookupRouteMapRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRouteMapRuleResultOutput)
}

// A collection of arguments for invoking getRouteMapRule.
type LookupRouteMapRuleOutputArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// Route-Map Rule name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupRouteMapRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteMapRuleArgs)(nil)).Elem()
}

// A collection of values returned by getRouteMapRule.
type LookupRouteMapRuleResultOutput struct{ *pulumi.OutputState }

func (LookupRouteMapRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRouteMapRuleResult)(nil)).Elem()
}

func (o LookupRouteMapRuleResultOutput) ToLookupRouteMapRuleResultOutput() LookupRouteMapRuleResultOutput {
	return o
}

func (o LookupRouteMapRuleResultOutput) ToLookupRouteMapRuleResultOutputWithContext(ctx context.Context) LookupRouteMapRuleResultOutput {
	return o
}

func (o LookupRouteMapRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupRouteMapRuleResult] {
	return pulumix.Output[LookupRouteMapRuleResult]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o LookupRouteMapRuleResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRouteMapRuleResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// The distinguished name of the object.
func (o LookupRouteMapRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// Route-Map Rule name.
func (o LookupRouteMapRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRouteMapRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRouteMapRuleResultOutput{})
}