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

func LookupPimAnycastRp(ctx *pulumi.Context, args *LookupPimAnycastRpArgs, opts ...pulumi.InvokeOption) (*LookupPimAnycastRpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPimAnycastRpResult
	err := ctx.Invoke("nxos:nxos/getPimAnycastRp:getPimAnycastRp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPimAnycastRp.
type LookupPimAnycastRpArgs struct {
	Device  *string `pulumi:"device"`
	VrfName string  `pulumi:"vrfName"`
}

// A collection of values returned by getPimAnycastRp.
type LookupPimAnycastRpResult struct {
	Device          *string `pulumi:"device"`
	Id              string  `pulumi:"id"`
	LocalInterface  string  `pulumi:"localInterface"`
	SourceInterface string  `pulumi:"sourceInterface"`
	VrfName         string  `pulumi:"vrfName"`
}

func LookupPimAnycastRpOutput(ctx *pulumi.Context, args LookupPimAnycastRpOutputArgs, opts ...pulumi.InvokeOption) LookupPimAnycastRpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPimAnycastRpResult, error) {
			args := v.(LookupPimAnycastRpArgs)
			r, err := LookupPimAnycastRp(ctx, &args, opts...)
			var s LookupPimAnycastRpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPimAnycastRpResultOutput)
}

// A collection of arguments for invoking getPimAnycastRp.
type LookupPimAnycastRpOutputArgs struct {
	Device  pulumi.StringPtrInput `pulumi:"device"`
	VrfName pulumi.StringInput    `pulumi:"vrfName"`
}

func (LookupPimAnycastRpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPimAnycastRpArgs)(nil)).Elem()
}

// A collection of values returned by getPimAnycastRp.
type LookupPimAnycastRpResultOutput struct{ *pulumi.OutputState }

func (LookupPimAnycastRpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPimAnycastRpResult)(nil)).Elem()
}

func (o LookupPimAnycastRpResultOutput) ToLookupPimAnycastRpResultOutput() LookupPimAnycastRpResultOutput {
	return o
}

func (o LookupPimAnycastRpResultOutput) ToLookupPimAnycastRpResultOutputWithContext(ctx context.Context) LookupPimAnycastRpResultOutput {
	return o
}

func (o LookupPimAnycastRpResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPimAnycastRpResult] {
	return pulumix.Output[LookupPimAnycastRpResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupPimAnycastRpResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPimAnycastRpResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupPimAnycastRpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimAnycastRpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPimAnycastRpResultOutput) LocalInterface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimAnycastRpResult) string { return v.LocalInterface }).(pulumi.StringOutput)
}

func (o LookupPimAnycastRpResultOutput) SourceInterface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimAnycastRpResult) string { return v.SourceInterface }).(pulumi.StringOutput)
}

func (o LookupPimAnycastRpResultOutput) VrfName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPimAnycastRpResult) string { return v.VrfName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPimAnycastRpResultOutput{})
}