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

func LookupFeatureLldp(ctx *pulumi.Context, args *LookupFeatureLldpArgs, opts ...pulumi.InvokeOption) (*LookupFeatureLldpResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFeatureLldpResult
	err := ctx.Invoke("nxos:nxos/getFeatureLldp:getFeatureLldp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFeatureLldp.
type LookupFeatureLldpArgs struct {
	Device *string `pulumi:"device"`
}

// A collection of values returned by getFeatureLldp.
type LookupFeatureLldpResult struct {
	AdminState string  `pulumi:"adminState"`
	Device     *string `pulumi:"device"`
	Id         string  `pulumi:"id"`
}

func LookupFeatureLldpOutput(ctx *pulumi.Context, args LookupFeatureLldpOutputArgs, opts ...pulumi.InvokeOption) LookupFeatureLldpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFeatureLldpResult, error) {
			args := v.(LookupFeatureLldpArgs)
			r, err := LookupFeatureLldp(ctx, &args, opts...)
			var s LookupFeatureLldpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFeatureLldpResultOutput)
}

// A collection of arguments for invoking getFeatureLldp.
type LookupFeatureLldpOutputArgs struct {
	Device pulumi.StringPtrInput `pulumi:"device"`
}

func (LookupFeatureLldpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureLldpArgs)(nil)).Elem()
}

// A collection of values returned by getFeatureLldp.
type LookupFeatureLldpResultOutput struct{ *pulumi.OutputState }

func (LookupFeatureLldpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureLldpResult)(nil)).Elem()
}

func (o LookupFeatureLldpResultOutput) ToLookupFeatureLldpResultOutput() LookupFeatureLldpResultOutput {
	return o
}

func (o LookupFeatureLldpResultOutput) ToLookupFeatureLldpResultOutputWithContext(ctx context.Context) LookupFeatureLldpResultOutput {
	return o
}

func (o LookupFeatureLldpResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFeatureLldpResult] {
	return pulumix.Output[LookupFeatureLldpResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupFeatureLldpResultOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureLldpResult) string { return v.AdminState }).(pulumi.StringOutput)
}

func (o LookupFeatureLldpResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFeatureLldpResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupFeatureLldpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureLldpResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFeatureLldpResultOutput{})
}
