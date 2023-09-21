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

func LookupOspfArea(ctx *pulumi.Context, args *LookupOspfAreaArgs, opts ...pulumi.InvokeOption) (*LookupOspfAreaResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOspfAreaResult
	err := ctx.Invoke("nxos:nxos/getOspfArea:getOspfArea", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOspfArea.
type LookupOspfAreaArgs struct {
	AreaId       string  `pulumi:"areaId"`
	Device       *string `pulumi:"device"`
	InstanceName string  `pulumi:"instanceName"`
	VrfName      string  `pulumi:"vrfName"`
}

// A collection of values returned by getOspfArea.
type LookupOspfAreaResult struct {
	AreaId             string  `pulumi:"areaId"`
	AuthenticationType string  `pulumi:"authenticationType"`
	Cost               int     `pulumi:"cost"`
	Device             *string `pulumi:"device"`
	Id                 string  `pulumi:"id"`
	InstanceName       string  `pulumi:"instanceName"`
	Type               string  `pulumi:"type"`
	VrfName            string  `pulumi:"vrfName"`
}

func LookupOspfAreaOutput(ctx *pulumi.Context, args LookupOspfAreaOutputArgs, opts ...pulumi.InvokeOption) LookupOspfAreaResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOspfAreaResult, error) {
			args := v.(LookupOspfAreaArgs)
			r, err := LookupOspfArea(ctx, &args, opts...)
			var s LookupOspfAreaResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOspfAreaResultOutput)
}

// A collection of arguments for invoking getOspfArea.
type LookupOspfAreaOutputArgs struct {
	AreaId       pulumi.StringInput    `pulumi:"areaId"`
	Device       pulumi.StringPtrInput `pulumi:"device"`
	InstanceName pulumi.StringInput    `pulumi:"instanceName"`
	VrfName      pulumi.StringInput    `pulumi:"vrfName"`
}

func (LookupOspfAreaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOspfAreaArgs)(nil)).Elem()
}

// A collection of values returned by getOspfArea.
type LookupOspfAreaResultOutput struct{ *pulumi.OutputState }

func (LookupOspfAreaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOspfAreaResult)(nil)).Elem()
}

func (o LookupOspfAreaResultOutput) ToLookupOspfAreaResultOutput() LookupOspfAreaResultOutput {
	return o
}

func (o LookupOspfAreaResultOutput) ToLookupOspfAreaResultOutputWithContext(ctx context.Context) LookupOspfAreaResultOutput {
	return o
}

func (o LookupOspfAreaResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupOspfAreaResult] {
	return pulumix.Output[LookupOspfAreaResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupOspfAreaResultOutput) AreaId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfAreaResult) string { return v.AreaId }).(pulumi.StringOutput)
}

func (o LookupOspfAreaResultOutput) AuthenticationType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfAreaResult) string { return v.AuthenticationType }).(pulumi.StringOutput)
}

func (o LookupOspfAreaResultOutput) Cost() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOspfAreaResult) int { return v.Cost }).(pulumi.IntOutput)
}

func (o LookupOspfAreaResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOspfAreaResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupOspfAreaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfAreaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOspfAreaResultOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfAreaResult) string { return v.InstanceName }).(pulumi.StringOutput)
}

func (o LookupOspfAreaResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfAreaResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupOspfAreaResultOutput) VrfName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOspfAreaResult) string { return v.VrfName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOspfAreaResultOutput{})
}
