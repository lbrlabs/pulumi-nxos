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

// This data source can read the default QoS class map DSCP configuration.
//
// - API Documentation: [ipqosDscp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:Dscp/)
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
//			_, err := nxos.LookupDefaultQosClassMapDscp(ctx, &nxos.LookupDefaultQosClassMapDscpArgs{
//				ClassMapName: "Voice",
//				Value:        "ef",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDefaultQosClassMapDscp(ctx *pulumi.Context, args *LookupDefaultQosClassMapDscpArgs, opts ...pulumi.InvokeOption) (*LookupDefaultQosClassMapDscpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDefaultQosClassMapDscpResult
	err := ctx.Invoke("nxos:index/getDefaultQosClassMapDscp:getDefaultQosClassMapDscp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDefaultQosClassMapDscp.
type LookupDefaultQosClassMapDscpArgs struct {
	// Class map name.
	ClassMapName string `pulumi:"classMapName"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// DSCP value.
	Value string `pulumi:"value"`
}

// A collection of values returned by getDefaultQosClassMapDscp.
type LookupDefaultQosClassMapDscpResult struct {
	// Class map name.
	ClassMapName string `pulumi:"classMapName"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// DSCP value.
	Value string `pulumi:"value"`
}

func LookupDefaultQosClassMapDscpOutput(ctx *pulumi.Context, args LookupDefaultQosClassMapDscpOutputArgs, opts ...pulumi.InvokeOption) LookupDefaultQosClassMapDscpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDefaultQosClassMapDscpResult, error) {
			args := v.(LookupDefaultQosClassMapDscpArgs)
			r, err := LookupDefaultQosClassMapDscp(ctx, &args, opts...)
			var s LookupDefaultQosClassMapDscpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDefaultQosClassMapDscpResultOutput)
}

// A collection of arguments for invoking getDefaultQosClassMapDscp.
type LookupDefaultQosClassMapDscpOutputArgs struct {
	// Class map name.
	ClassMapName pulumi.StringInput `pulumi:"classMapName"`
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// DSCP value.
	Value pulumi.StringInput `pulumi:"value"`
}

func (LookupDefaultQosClassMapDscpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultQosClassMapDscpArgs)(nil)).Elem()
}

// A collection of values returned by getDefaultQosClassMapDscp.
type LookupDefaultQosClassMapDscpResultOutput struct{ *pulumi.OutputState }

func (LookupDefaultQosClassMapDscpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDefaultQosClassMapDscpResult)(nil)).Elem()
}

func (o LookupDefaultQosClassMapDscpResultOutput) ToLookupDefaultQosClassMapDscpResultOutput() LookupDefaultQosClassMapDscpResultOutput {
	return o
}

func (o LookupDefaultQosClassMapDscpResultOutput) ToLookupDefaultQosClassMapDscpResultOutputWithContext(ctx context.Context) LookupDefaultQosClassMapDscpResultOutput {
	return o
}

func (o LookupDefaultQosClassMapDscpResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDefaultQosClassMapDscpResult] {
	return pulumix.Output[LookupDefaultQosClassMapDscpResult]{
		OutputState: o.OutputState,
	}
}

// Class map name.
func (o LookupDefaultQosClassMapDscpResultOutput) ClassMapName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosClassMapDscpResult) string { return v.ClassMapName }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o LookupDefaultQosClassMapDscpResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDefaultQosClassMapDscpResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// The distinguished name of the object.
func (o LookupDefaultQosClassMapDscpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosClassMapDscpResult) string { return v.Id }).(pulumi.StringOutput)
}

// DSCP value.
func (o LookupDefaultQosClassMapDscpResultOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDefaultQosClassMapDscpResult) string { return v.Value }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDefaultQosClassMapDscpResultOutput{})
}
