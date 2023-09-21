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

func LookupFeatureVpc(ctx *pulumi.Context, args *LookupFeatureVpcArgs, opts ...pulumi.InvokeOption) (*LookupFeatureVpcResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFeatureVpcResult
	err := ctx.Invoke("nxos:nxos/getFeatureVpc:getFeatureVpc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFeatureVpc.
type LookupFeatureVpcArgs struct {
	Device *string `pulumi:"device"`
}

// A collection of values returned by getFeatureVpc.
type LookupFeatureVpcResult struct {
	AdminState string  `pulumi:"adminState"`
	Device     *string `pulumi:"device"`
	Id         string  `pulumi:"id"`
}

func LookupFeatureVpcOutput(ctx *pulumi.Context, args LookupFeatureVpcOutputArgs, opts ...pulumi.InvokeOption) LookupFeatureVpcResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFeatureVpcResult, error) {
			args := v.(LookupFeatureVpcArgs)
			r, err := LookupFeatureVpc(ctx, &args, opts...)
			var s LookupFeatureVpcResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFeatureVpcResultOutput)
}

// A collection of arguments for invoking getFeatureVpc.
type LookupFeatureVpcOutputArgs struct {
	Device pulumi.StringPtrInput `pulumi:"device"`
}

func (LookupFeatureVpcOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureVpcArgs)(nil)).Elem()
}

// A collection of values returned by getFeatureVpc.
type LookupFeatureVpcResultOutput struct{ *pulumi.OutputState }

func (LookupFeatureVpcResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFeatureVpcResult)(nil)).Elem()
}

func (o LookupFeatureVpcResultOutput) ToLookupFeatureVpcResultOutput() LookupFeatureVpcResultOutput {
	return o
}

func (o LookupFeatureVpcResultOutput) ToLookupFeatureVpcResultOutputWithContext(ctx context.Context) LookupFeatureVpcResultOutput {
	return o
}

func (o LookupFeatureVpcResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupFeatureVpcResult] {
	return pulumix.Output[LookupFeatureVpcResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupFeatureVpcResultOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureVpcResult) string { return v.AdminState }).(pulumi.StringOutput)
}

func (o LookupFeatureVpcResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFeatureVpcResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupFeatureVpcResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFeatureVpcResult) string { return v.Id }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFeatureVpcResultOutput{})
}