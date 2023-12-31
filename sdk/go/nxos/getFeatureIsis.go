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

// This data source can read the ISIS feature configuration.
//
// - API Documentation: [fmIsis](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Isis/)
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
//			_, err := nxos.LookupFeatureIsis(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFeatureIsis(ctx *pulumi.Context, args *LookupFeatureIsisArgs, opts ...pulumi.InvokeOption) (*LookupFeatureIsisResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFeatureIsisResult
	err := ctx.Invoke("nxos:index/getFeatureIsis:getFeatureIsis", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFeatureIsis.
type LookupFeatureIsisArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

// A collection of values returned by getFeatureIsis.
type LookupFeatureIsisResult struct {
	// Administrative state.
	AdminState string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
}

func LookupFeatureIsisOutput(ctx *pulumi.Context, args LookupFeatureIsisOutputArgs, opts ...pulumi.InvokeOption) LookupFeatureIsisResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFeatureIsisResult, error) {
			args := v.(LookupFeatureIsisArgs)
			r, err := LookupFeatureIsis(ctx, &args, opts...)
			var s LookupFeatureIsisResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFeatureIsisResultOutput)
}

// A collection of arguments for invoking getFeatureIsis.
type LookupFeatureIsisOutputArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
}

func (LookupFeatureIsisOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureIsisArgs)(nil)).Elem()
}

// A collection of values returned by getFeatureIsis.
type LookupFeatureIsisResultOutput struct{ *pulumi.OutputState }

func (LookupFeatureIsisResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureIsisResult)(nil)).Elem()
}

func (o LookupFeatureIsisResultOutput) ToLookupFeatureIsisResultOutput() LookupFeatureIsisResultOutput {
	return o
}

func (o LookupFeatureIsisResultOutput) ToLookupFeatureIsisResultOutputWithContext(ctx context.Context) LookupFeatureIsisResultOutput {
	return o
}

func (o LookupFeatureIsisResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFeatureIsisResult] {
	return pulumix.Output[LookupFeatureIsisResult]{
		OutputState: o.OutputState,
	}
}

// Administrative state.
func (o LookupFeatureIsisResultOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureIsisResult) string { return v.AdminState }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o LookupFeatureIsisResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFeatureIsisResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// The distinguished name of the object.
func (o LookupFeatureIsisResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureIsisResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFeatureIsisResultOutput{})
}
