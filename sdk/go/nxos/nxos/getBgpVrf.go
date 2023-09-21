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

func LookupBgpVrf(ctx *pulumi.Context, args *LookupBgpVrfArgs, opts ...pulumi.InvokeOption) (*LookupBgpVrfResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBgpVrfResult
	err := ctx.Invoke("nxos:nxos/getBgpVrf:getBgpVrf", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBgpVrf.
type LookupBgpVrfArgs struct {
	Asn    string  `pulumi:"asn"`
	Device *string `pulumi:"device"`
	Name   string  `pulumi:"name"`
}

// A collection of values returned by getBgpVrf.
type LookupBgpVrfResult struct {
	Asn      string  `pulumi:"asn"`
	Device   *string `pulumi:"device"`
	Id       string  `pulumi:"id"`
	Name     string  `pulumi:"name"`
	RouterId string  `pulumi:"routerId"`
}

func LookupBgpVrfOutput(ctx *pulumi.Context, args LookupBgpVrfOutputArgs, opts ...pulumi.InvokeOption) LookupBgpVrfResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBgpVrfResult, error) {
			args := v.(LookupBgpVrfArgs)
			r, err := LookupBgpVrf(ctx, &args, opts...)
			var s LookupBgpVrfResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBgpVrfResultOutput)
}

// A collection of arguments for invoking getBgpVrf.
type LookupBgpVrfOutputArgs struct {
	Asn    pulumi.StringInput    `pulumi:"asn"`
	Device pulumi.StringPtrInput `pulumi:"device"`
	Name   pulumi.StringInput    `pulumi:"name"`
}

func (LookupBgpVrfOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpVrfArgs)(nil)).Elem()
}

// A collection of values returned by getBgpVrf.
type LookupBgpVrfResultOutput struct{ *pulumi.OutputState }

func (LookupBgpVrfResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpVrfResult)(nil)).Elem()
}

func (o LookupBgpVrfResultOutput) ToLookupBgpVrfResultOutput() LookupBgpVrfResultOutput {
	return o
}

func (o LookupBgpVrfResultOutput) ToLookupBgpVrfResultOutputWithContext(ctx context.Context) LookupBgpVrfResultOutput {
	return o
}

func (o LookupBgpVrfResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBgpVrfResult] {
	return pulumix.Output[LookupBgpVrfResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupBgpVrfResultOutput) Asn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpVrfResult) string { return v.Asn }).(pulumi.StringOutput)
}

func (o LookupBgpVrfResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBgpVrfResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupBgpVrfResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpVrfResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBgpVrfResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpVrfResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupBgpVrfResultOutput) RouterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpVrfResult) string { return v.RouterId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBgpVrfResultOutput{})
}
