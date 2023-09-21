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

func LookupFeatureNvOverlay(ctx *pulumi.Context, args *LookupFeatureNvOverlayArgs, opts ...pulumi.InvokeOption) (*LookupFeatureNvOverlayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFeatureNvOverlayResult
	err := ctx.Invoke("nxos:nxos/getFeatureNvOverlay:getFeatureNvOverlay", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFeatureNvOverlay.
type LookupFeatureNvOverlayArgs struct {
	Device *string `pulumi:"device"`
}

// A collection of values returned by getFeatureNvOverlay.
type LookupFeatureNvOverlayResult struct {
	AdminState string  `pulumi:"adminState"`
	Device     *string `pulumi:"device"`
	Id         string  `pulumi:"id"`
}

func LookupFeatureNvOverlayOutput(ctx *pulumi.Context, args LookupFeatureNvOverlayOutputArgs, opts ...pulumi.InvokeOption) LookupFeatureNvOverlayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFeatureNvOverlayResult, error) {
			args := v.(LookupFeatureNvOverlayArgs)
			r, err := LookupFeatureNvOverlay(ctx, &args, opts...)
			var s LookupFeatureNvOverlayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFeatureNvOverlayResultOutput)
}

// A collection of arguments for invoking getFeatureNvOverlay.
type LookupFeatureNvOverlayOutputArgs struct {
	Device pulumi.StringPtrInput `pulumi:"device"`
}

func (LookupFeatureNvOverlayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureNvOverlayArgs)(nil)).Elem()
}

// A collection of values returned by getFeatureNvOverlay.
type LookupFeatureNvOverlayResultOutput struct{ *pulumi.OutputState }

func (LookupFeatureNvOverlayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureNvOverlayResult)(nil)).Elem()
}

func (o LookupFeatureNvOverlayResultOutput) ToLookupFeatureNvOverlayResultOutput() LookupFeatureNvOverlayResultOutput {
	return o
}

func (o LookupFeatureNvOverlayResultOutput) ToLookupFeatureNvOverlayResultOutputWithContext(ctx context.Context) LookupFeatureNvOverlayResultOutput {
	return o
}

func (o LookupFeatureNvOverlayResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFeatureNvOverlayResult] {
	return pulumix.Output[LookupFeatureNvOverlayResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupFeatureNvOverlayResultOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureNvOverlayResult) string { return v.AdminState }).(pulumi.StringOutput)
}

func (o LookupFeatureNvOverlayResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFeatureNvOverlayResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupFeatureNvOverlayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureNvOverlayResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFeatureNvOverlayResultOutput{})
}
