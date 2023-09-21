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

func LookupSpanningTreeInterface(ctx *pulumi.Context, args *LookupSpanningTreeInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupSpanningTreeInterfaceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSpanningTreeInterfaceResult
	err := ctx.Invoke("nxos:nxos/getSpanningTreeInterface:getSpanningTreeInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSpanningTreeInterface.
type LookupSpanningTreeInterfaceArgs struct {
	Device      *string `pulumi:"device"`
	InterfaceId string  `pulumi:"interfaceId"`
}

// A collection of values returned by getSpanningTreeInterface.
type LookupSpanningTreeInterfaceResult struct {
	AdminState  string  `pulumi:"adminState"`
	BpduFilter  string  `pulumi:"bpduFilter"`
	BpduGuard   string  `pulumi:"bpduGuard"`
	Cost        int     `pulumi:"cost"`
	Device      *string `pulumi:"device"`
	Guard       string  `pulumi:"guard"`
	Id          string  `pulumi:"id"`
	InterfaceId string  `pulumi:"interfaceId"`
	LinkType    string  `pulumi:"linkType"`
	Mode        string  `pulumi:"mode"`
	Priority    int     `pulumi:"priority"`
}

func LookupSpanningTreeInterfaceOutput(ctx *pulumi.Context, args LookupSpanningTreeInterfaceOutputArgs, opts ...pulumi.InvokeOption) LookupSpanningTreeInterfaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSpanningTreeInterfaceResult, error) {
			args := v.(LookupSpanningTreeInterfaceArgs)
			r, err := LookupSpanningTreeInterface(ctx, &args, opts...)
			var s LookupSpanningTreeInterfaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSpanningTreeInterfaceResultOutput)
}

// A collection of arguments for invoking getSpanningTreeInterface.
type LookupSpanningTreeInterfaceOutputArgs struct {
	Device      pulumi.StringPtrInput `pulumi:"device"`
	InterfaceId pulumi.StringInput    `pulumi:"interfaceId"`
}

func (LookupSpanningTreeInterfaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpanningTreeInterfaceArgs)(nil)).Elem()
}

// A collection of values returned by getSpanningTreeInterface.
type LookupSpanningTreeInterfaceResultOutput struct{ *pulumi.OutputState }

func (LookupSpanningTreeInterfaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSpanningTreeInterfaceResult)(nil)).Elem()
}

func (o LookupSpanningTreeInterfaceResultOutput) ToLookupSpanningTreeInterfaceResultOutput() LookupSpanningTreeInterfaceResultOutput {
	return o
}

func (o LookupSpanningTreeInterfaceResultOutput) ToLookupSpanningTreeInterfaceResultOutputWithContext(ctx context.Context) LookupSpanningTreeInterfaceResultOutput {
	return o
}

func (o LookupSpanningTreeInterfaceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSpanningTreeInterfaceResult] {
	return pulumix.Output[LookupSpanningTreeInterfaceResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupSpanningTreeInterfaceResultOutput) AdminState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpanningTreeInterfaceResult) string { return v.AdminState }).(pulumi.StringOutput)
}

func (o LookupSpanningTreeInterfaceResultOutput) BpduFilter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpanningTreeInterfaceResult) string { return v.BpduFilter }).(pulumi.StringOutput)
}

func (o LookupSpanningTreeInterfaceResultOutput) BpduGuard() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpanningTreeInterfaceResult) string { return v.BpduGuard }).(pulumi.StringOutput)
}

func (o LookupSpanningTreeInterfaceResultOutput) Cost() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSpanningTreeInterfaceResult) int { return v.Cost }).(pulumi.IntOutput)
}

func (o LookupSpanningTreeInterfaceResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSpanningTreeInterfaceResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupSpanningTreeInterfaceResultOutput) Guard() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpanningTreeInterfaceResult) string { return v.Guard }).(pulumi.StringOutput)
}

func (o LookupSpanningTreeInterfaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpanningTreeInterfaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSpanningTreeInterfaceResultOutput) InterfaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpanningTreeInterfaceResult) string { return v.InterfaceId }).(pulumi.StringOutput)
}

func (o LookupSpanningTreeInterfaceResultOutput) LinkType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpanningTreeInterfaceResult) string { return v.LinkType }).(pulumi.StringOutput)
}

func (o LookupSpanningTreeInterfaceResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSpanningTreeInterfaceResult) string { return v.Mode }).(pulumi.StringOutput)
}

func (o LookupSpanningTreeInterfaceResultOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSpanningTreeInterfaceResult) int { return v.Priority }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSpanningTreeInterfaceResultOutput{})
}