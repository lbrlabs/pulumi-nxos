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

// This data source can read the default QoS policy map match class map configuration.
//
// - API Documentation: [ipqosMatchCMap](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:MatchCMap/)
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
//			_, err := nxos.LookupDefaultQosPolicyMapMatchClassMap(ctx, &nxos.LookupDefaultQosPolicyMapMatchClassMapArgs{
//				Name:          "Voice",
//				PolicyMapName: "PM1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDefaultQosPolicyMapMatchClassMap(ctx *pulumi.Context, args *LookupDefaultQosPolicyMapMatchClassMapArgs, opts ...pulumi.InvokeOption) (*LookupDefaultQosPolicyMapMatchClassMapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDefaultQosPolicyMapMatchClassMapResult
	err := ctx.Invoke("nxos:index/getDefaultQosPolicyMapMatchClassMap:getDefaultQosPolicyMapMatchClassMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDefaultQosPolicyMapMatchClassMap.
type LookupDefaultQosPolicyMapMatchClassMapArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Class map name.
	Name string `pulumi:"name"`
	// Policy map name.
	PolicyMapName string `pulumi:"policyMapName"`
}

// A collection of values returned by getDefaultQosPolicyMapMatchClassMap.
type LookupDefaultQosPolicyMapMatchClassMapResult struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// Class map name.
	Name string `pulumi:"name"`
	// Policy map name.
	PolicyMapName string `pulumi:"policyMapName"`
}

func LookupDefaultQosPolicyMapMatchClassMapOutput(ctx *pulumi.Context, args LookupDefaultQosPolicyMapMatchClassMapOutputArgs, opts ...pulumi.InvokeOption) LookupDefaultQosPolicyMapMatchClassMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDefaultQosPolicyMapMatchClassMapResult, error) {
			args := v.(LookupDefaultQosPolicyMapMatchClassMapArgs)
			r, err := LookupDefaultQosPolicyMapMatchClassMap(ctx, &args, opts...)
			var s LookupDefaultQosPolicyMapMatchClassMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDefaultQosPolicyMapMatchClassMapResultOutput)
}

// A collection of arguments for invoking getDefaultQosPolicyMapMatchClassMap.
type LookupDefaultQosPolicyMapMatchClassMapOutputArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// Class map name.
	Name pulumi.StringInput `pulumi:"name"`
	// Policy map name.
	PolicyMapName pulumi.StringInput `pulumi:"policyMapName"`
}

func (LookupDefaultQosPolicyMapMatchClassMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultQosPolicyMapMatchClassMapArgs)(nil)).Elem()
}

// A collection of values returned by getDefaultQosPolicyMapMatchClassMap.
type LookupDefaultQosPolicyMapMatchClassMapResultOutput struct{ *pulumi.OutputState }

func (LookupDefaultQosPolicyMapMatchClassMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultQosPolicyMapMatchClassMapResult)(nil)).Elem()
}

func (o LookupDefaultQosPolicyMapMatchClassMapResultOutput) ToLookupDefaultQosPolicyMapMatchClassMapResultOutput() LookupDefaultQosPolicyMapMatchClassMapResultOutput {
	return o
}

func (o LookupDefaultQosPolicyMapMatchClassMapResultOutput) ToLookupDefaultQosPolicyMapMatchClassMapResultOutputWithContext(ctx context.Context) LookupDefaultQosPolicyMapMatchClassMapResultOutput {
	return o
}

func (o LookupDefaultQosPolicyMapMatchClassMapResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDefaultQosPolicyMapMatchClassMapResult] {
	return pulumix.Output[LookupDefaultQosPolicyMapMatchClassMapResult]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o LookupDefaultQosPolicyMapMatchClassMapResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDefaultQosPolicyMapMatchClassMapResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// The distinguished name of the object.
func (o LookupDefaultQosPolicyMapMatchClassMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosPolicyMapMatchClassMapResult) string { return v.Id }).(pulumi.StringOutput)
}

// Class map name.
func (o LookupDefaultQosPolicyMapMatchClassMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosPolicyMapMatchClassMapResult) string { return v.Name }).(pulumi.StringOutput)
}

// Policy map name.
func (o LookupDefaultQosPolicyMapMatchClassMapResultOutput) PolicyMapName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosPolicyMapMatchClassMapResult) string { return v.PolicyMapName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDefaultQosPolicyMapMatchClassMapResultOutput{})
}