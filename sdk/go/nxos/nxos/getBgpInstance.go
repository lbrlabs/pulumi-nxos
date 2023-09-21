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

func LookupBgpInstance(ctx *pulumi.Context, args *LookupBgpInstanceArgs, opts ...pulumi.InvokeOption) (*LookupBgpInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBgpInstanceResult
	err := ctx.Invoke("nxos:nxos/getBgpInstance:getBgpInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBgpInstance.
type LookupBgpInstanceArgs struct {
	Device *string `pulumi:"device"`
}

// A collection of values returned by getBgpInstance.
type LookupBgpInstanceResult struct {
	AdminState            string  `pulumi:"adminState"`
	Asn                   string  `pulumi:"asn"`
	Device                *string `pulumi:"device"`
	EnhancedErrorHandling bool    `pulumi:"enhancedErrorHandling"`
	Id                    string  `pulumi:"id"`
}

func LookupBgpInstanceOutput(ctx *pulumi.Context, args LookupBgpInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupBgpInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBgpInstanceResult, error) {
			args := v.(LookupBgpInstanceArgs)
			r, err := LookupBgpInstance(ctx, &args, opts...)
			var s LookupBgpInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBgpInstanceResultOutput)
}

// A collection of arguments for invoking getBgpInstance.
type LookupBgpInstanceOutputArgs struct {
	Device pulumi.StringPtrInput `pulumi:"device"`
}

func (LookupBgpInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getBgpInstance.
type LookupBgpInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupBgpInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpInstanceResult)(nil)).Elem()
}

func (o LookupBgpInstanceResultOutput) ToLookupBgpInstanceResultOutput() LookupBgpInstanceResultOutput {
	return o
}

func (o LookupBgpInstanceResultOutput) ToLookupBgpInstanceResultOutputWithContext(ctx context.Context) LookupBgpInstanceResultOutput {
	return o
}

func (o LookupBgpInstanceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBgpInstanceResult] {
	return pulumix.Output[LookupBgpInstanceResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupBgpInstanceResultOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpInstanceResult) string { return v.AdminState }).(pulumi.StringOutput)
}

func (o LookupBgpInstanceResultOutput) Asn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpInstanceResult) string { return v.Asn }).(pulumi.StringOutput)
}

func (o LookupBgpInstanceResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBgpInstanceResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupBgpInstanceResultOutput) EnhancedErrorHandling() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBgpInstanceResult) bool { return v.EnhancedErrorHandling }).(pulumi.BoolOutput)
}

func (o LookupBgpInstanceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpInstanceResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBgpInstanceResultOutput{})
}
