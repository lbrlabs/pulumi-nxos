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

// This data source can read an IPv4 Prefix List configuration.
//
// - API Documentation: [rtpfxRuleV4](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtpfx:RuleV4/)
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
//			_, err := nxos.LookupIpv4PrefixListRule(ctx, &nxos.LookupIpv4PrefixListRuleArgs{
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
func LookupIpv4PrefixListRule(ctx *pulumi.Context, args *LookupIpv4PrefixListRuleArgs, opts ...pulumi.InvokeOption) (*LookupIpv4PrefixListRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpv4PrefixListRuleResult
	err := ctx.Invoke("nxos:index/getIpv4PrefixListRule:getIpv4PrefixListRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpv4PrefixListRule.
type LookupIpv4PrefixListRuleArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// IPv4 Prefix List Rule name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getIpv4PrefixListRule.
type LookupIpv4PrefixListRuleResult struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// IPv4 Prefix List Rule name.
	Name string `pulumi:"name"`
}

func LookupIpv4PrefixListRuleOutput(ctx *pulumi.Context, args LookupIpv4PrefixListRuleOutputArgs, opts ...pulumi.InvokeOption) LookupIpv4PrefixListRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpv4PrefixListRuleResult, error) {
			args := v.(LookupIpv4PrefixListRuleArgs)
			r, err := LookupIpv4PrefixListRule(ctx, &args, opts...)
			var s LookupIpv4PrefixListRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpv4PrefixListRuleResultOutput)
}

// A collection of arguments for invoking getIpv4PrefixListRule.
type LookupIpv4PrefixListRuleOutputArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// IPv4 Prefix List Rule name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupIpv4PrefixListRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpv4PrefixListRuleArgs)(nil)).Elem()
}

// A collection of values returned by getIpv4PrefixListRule.
type LookupIpv4PrefixListRuleResultOutput struct{ *pulumi.OutputState }

func (LookupIpv4PrefixListRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpv4PrefixListRuleResult)(nil)).Elem()
}

func (o LookupIpv4PrefixListRuleResultOutput) ToLookupIpv4PrefixListRuleResultOutput() LookupIpv4PrefixListRuleResultOutput {
	return o
}

func (o LookupIpv4PrefixListRuleResultOutput) ToLookupIpv4PrefixListRuleResultOutputWithContext(ctx context.Context) LookupIpv4PrefixListRuleResultOutput {
	return o
}

func (o LookupIpv4PrefixListRuleResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIpv4PrefixListRuleResult] {
	return pulumix.Output[LookupIpv4PrefixListRuleResult]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o LookupIpv4PrefixListRuleResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpv4PrefixListRuleResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// The distinguished name of the object.
func (o LookupIpv4PrefixListRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4PrefixListRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

// IPv4 Prefix List Rule name.
func (o LookupIpv4PrefixListRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4PrefixListRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpv4PrefixListRuleResultOutput{})
}
