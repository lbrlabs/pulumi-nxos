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

// This data source can read the default QoS class map configuration.
//
// - API Documentation: [ipqosCMapInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:CMapInst/)
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
//			_, err := nxos.LookupDefaultQosClassMap(ctx, &nxos.LookupDefaultQosClassMapArgs{
//				Name: "Voice",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDefaultQosClassMap(ctx *pulumi.Context, args *LookupDefaultQosClassMapArgs, opts ...pulumi.InvokeOption) (*LookupDefaultQosClassMapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDefaultQosClassMapResult
	err := ctx.Invoke("nxos:index/getDefaultQosClassMap:getDefaultQosClassMap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDefaultQosClassMap.
type LookupDefaultQosClassMapArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Class map name.
	Name string `pulumi:"name"`
}

// A collection of values returned by getDefaultQosClassMap.
type LookupDefaultQosClassMapResult struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// Match type.
	MatchType string `pulumi:"matchType"`
	// Class map name.
	Name string `pulumi:"name"`
}

func LookupDefaultQosClassMapOutput(ctx *pulumi.Context, args LookupDefaultQosClassMapOutputArgs, opts ...pulumi.InvokeOption) LookupDefaultQosClassMapResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDefaultQosClassMapResult, error) {
			args := v.(LookupDefaultQosClassMapArgs)
			r, err := LookupDefaultQosClassMap(ctx, &args, opts...)
			var s LookupDefaultQosClassMapResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDefaultQosClassMapResultOutput)
}

// A collection of arguments for invoking getDefaultQosClassMap.
type LookupDefaultQosClassMapOutputArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// Class map name.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupDefaultQosClassMapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultQosClassMapArgs)(nil)).Elem()
}

// A collection of values returned by getDefaultQosClassMap.
type LookupDefaultQosClassMapResultOutput struct{ *pulumi.OutputState }

func (LookupDefaultQosClassMapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultQosClassMapResult)(nil)).Elem()
}

func (o LookupDefaultQosClassMapResultOutput) ToLookupDefaultQosClassMapResultOutput() LookupDefaultQosClassMapResultOutput {
	return o
}

func (o LookupDefaultQosClassMapResultOutput) ToLookupDefaultQosClassMapResultOutputWithContext(ctx context.Context) LookupDefaultQosClassMapResultOutput {
	return o
}

func (o LookupDefaultQosClassMapResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDefaultQosClassMapResult] {
	return pulumix.Output[LookupDefaultQosClassMapResult]{
		OutputState: o.OutputState,
	}
}

// A device name from the provider configuration.
func (o LookupDefaultQosClassMapResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDefaultQosClassMapResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// The distinguished name of the object.
func (o LookupDefaultQosClassMapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosClassMapResult) string { return v.Id }).(pulumi.StringOutput)
}

// Match type.
func (o LookupDefaultQosClassMapResultOutput) MatchType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosClassMapResult) string { return v.MatchType }).(pulumi.StringOutput)
}

// Class map name.
func (o LookupDefaultQosClassMapResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosClassMapResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDefaultQosClassMapResultOutput{})
}