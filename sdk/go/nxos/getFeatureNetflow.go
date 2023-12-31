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

// This data source can read the Netflow feature configuration.
//
// - API Documentation: [fmNetflow](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Netflow/)
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
//			_, err := nxos.LookupFeatureNetflow(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFeatureNetflow(ctx *pulumi.Context, args *LookupFeatureNetflowArgs, opts ...pulumi.InvokeOption) (*LookupFeatureNetflowResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFeatureNetflowResult
	err := ctx.Invoke("nxos:index/getFeatureNetflow:getFeatureNetflow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFeatureNetflow.
type LookupFeatureNetflowArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
}

// A collection of values returned by getFeatureNetflow.
type LookupFeatureNetflowResult struct {
	// Administrative state.
	AdminState string `pulumi:"adminState"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
}

func LookupFeatureNetflowOutput(ctx *pulumi.Context, args LookupFeatureNetflowOutputArgs, opts ...pulumi.InvokeOption) LookupFeatureNetflowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFeatureNetflowResult, error) {
			args := v.(LookupFeatureNetflowArgs)
			r, err := LookupFeatureNetflow(ctx, &args, opts...)
			var s LookupFeatureNetflowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFeatureNetflowResultOutput)
}

// A collection of arguments for invoking getFeatureNetflow.
type LookupFeatureNetflowOutputArgs struct {
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
}

func (LookupFeatureNetflowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureNetflowArgs)(nil)).Elem()
}

// A collection of values returned by getFeatureNetflow.
type LookupFeatureNetflowResultOutput struct{ *pulumi.OutputState }

func (LookupFeatureNetflowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureNetflowResult)(nil)).Elem()
}

func (o LookupFeatureNetflowResultOutput) ToLookupFeatureNetflowResultOutput() LookupFeatureNetflowResultOutput {
	return o
}

func (o LookupFeatureNetflowResultOutput) ToLookupFeatureNetflowResultOutputWithContext(ctx context.Context) LookupFeatureNetflowResultOutput {
	return o
}

func (o LookupFeatureNetflowResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFeatureNetflowResult] {
	return pulumix.Output[LookupFeatureNetflowResult]{
		OutputState: o.OutputState,
	}
}

// Administrative state.
func (o LookupFeatureNetflowResultOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureNetflowResult) string { return v.AdminState }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o LookupFeatureNetflowResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFeatureNetflowResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// The distinguished name of the object.
func (o LookupFeatureNetflowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureNetflowResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFeatureNetflowResultOutput{})
}
