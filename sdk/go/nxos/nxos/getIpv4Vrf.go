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

func LookupIpv4Vrf(ctx *pulumi.Context, args *LookupIpv4VrfArgs, opts ...pulumi.InvokeOption) (*LookupIpv4VrfResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpv4VrfResult
	err := ctx.Invoke("nxos:nxos/getIpv4Vrf:getIpv4Vrf", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpv4Vrf.
type LookupIpv4VrfArgs struct {
	Device *string `pulumi:"device"`
	Name   string  `pulumi:"name"`
}

// A collection of values returned by getIpv4Vrf.
type LookupIpv4VrfResult struct {
	Device *string `pulumi:"device"`
	Id     string  `pulumi:"id"`
	Name   string  `pulumi:"name"`
}

func LookupIpv4VrfOutput(ctx *pulumi.Context, args LookupIpv4VrfOutputArgs, opts ...pulumi.InvokeOption) LookupIpv4VrfResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpv4VrfResult, error) {
			args := v.(LookupIpv4VrfArgs)
			r, err := LookupIpv4Vrf(ctx, &args, opts...)
			var s LookupIpv4VrfResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpv4VrfResultOutput)
}

// A collection of arguments for invoking getIpv4Vrf.
type LookupIpv4VrfOutputArgs struct {
	Device pulumi.StringPtrInput `pulumi:"device"`
	Name   pulumi.StringInput    `pulumi:"name"`
}

func (LookupIpv4VrfOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpv4VrfArgs)(nil)).Elem()
}

// A collection of values returned by getIpv4Vrf.
type LookupIpv4VrfResultOutput struct{ *pulumi.OutputState }

func (LookupIpv4VrfResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpv4VrfResult)(nil)).Elem()
}

func (o LookupIpv4VrfResultOutput) ToLookupIpv4VrfResultOutput() LookupIpv4VrfResultOutput {
	return o
}

func (o LookupIpv4VrfResultOutput) ToLookupIpv4VrfResultOutputWithContext(ctx context.Context) LookupIpv4VrfResultOutput {
	return o
}

func (o LookupIpv4VrfResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIpv4VrfResult] {
	return pulumix.Output[LookupIpv4VrfResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupIpv4VrfResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpv4VrfResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupIpv4VrfResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4VrfResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIpv4VrfResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4VrfResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpv4VrfResultOutput{})
}
