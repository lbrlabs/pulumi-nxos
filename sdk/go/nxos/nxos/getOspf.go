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

func LookupOspf(ctx *pulumi.Context, args *LookupOspfArgs, opts ...pulumi.InvokeOption) (*LookupOspfResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOspfResult
	err := ctx.Invoke("nxos:nxos/getOspf:getOspf", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOspf.
type LookupOspfArgs struct {
	Device *string `pulumi:"device"`
}

// A collection of values returned by getOspf.
type LookupOspfResult struct {
	AdminState string  `pulumi:"adminState"`
	Device     *string `pulumi:"device"`
	Id         string  `pulumi:"id"`
}

func LookupOspfOutput(ctx *pulumi.Context, args LookupOspfOutputArgs, opts ...pulumi.InvokeOption) LookupOspfResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOspfResult, error) {
			args := v.(LookupOspfArgs)
			r, err := LookupOspf(ctx, &args, opts...)
			var s LookupOspfResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOspfResultOutput)
}

// A collection of arguments for invoking getOspf.
type LookupOspfOutputArgs struct {
	Device pulumi.StringPtrInput `pulumi:"device"`
}

func (LookupOspfOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOspfArgs)(nil)).Elem()
}

// A collection of values returned by getOspf.
type LookupOspfResultOutput struct{ *pulumi.OutputState }

func (LookupOspfResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOspfResult)(nil)).Elem()
}

func (o LookupOspfResultOutput) ToLookupOspfResultOutput() LookupOspfResultOutput {
	return o
}

func (o LookupOspfResultOutput) ToLookupOspfResultOutputWithContext(ctx context.Context) LookupOspfResultOutput {
	return o
}

func (o LookupOspfResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOspfResult] {
	return pulumix.Output[LookupOspfResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupOspfResultOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.AdminState }).(pulumi.StringOutput)
}

func (o LookupOspfResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOspfResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupOspfResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOspfResultOutput{})
}