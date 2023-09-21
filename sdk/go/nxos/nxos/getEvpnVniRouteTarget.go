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

func LookupEvpnVniRouteTarget(ctx *pulumi.Context, args *LookupEvpnVniRouteTargetArgs, opts ...pulumi.InvokeOption) (*LookupEvpnVniRouteTargetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEvpnVniRouteTargetResult
	err := ctx.Invoke("nxos:nxos/getEvpnVniRouteTarget:getEvpnVniRouteTarget", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEvpnVniRouteTarget.
type LookupEvpnVniRouteTargetArgs struct {
	Device      *string `pulumi:"device"`
	Direction   string  `pulumi:"direction"`
	Encap       string  `pulumi:"encap"`
	RouteTarget string  `pulumi:"routeTarget"`
}

// A collection of values returned by getEvpnVniRouteTarget.
type LookupEvpnVniRouteTargetResult struct {
	Device      *string `pulumi:"device"`
	Direction   string  `pulumi:"direction"`
	Encap       string  `pulumi:"encap"`
	Id          string  `pulumi:"id"`
	RouteTarget string  `pulumi:"routeTarget"`
}

func LookupEvpnVniRouteTargetOutput(ctx *pulumi.Context, args LookupEvpnVniRouteTargetOutputArgs, opts ...pulumi.InvokeOption) LookupEvpnVniRouteTargetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEvpnVniRouteTargetResult, error) {
			args := v.(LookupEvpnVniRouteTargetArgs)
			r, err := LookupEvpnVniRouteTarget(ctx, &args, opts...)
			var s LookupEvpnVniRouteTargetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEvpnVniRouteTargetResultOutput)
}

// A collection of arguments for invoking getEvpnVniRouteTarget.
type LookupEvpnVniRouteTargetOutputArgs struct {
	Device      pulumi.StringPtrInput `pulumi:"device"`
	Direction   pulumi.StringInput    `pulumi:"direction"`
	Encap       pulumi.StringInput    `pulumi:"encap"`
	RouteTarget pulumi.StringInput    `pulumi:"routeTarget"`
}

func (LookupEvpnVniRouteTargetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEvpnVniRouteTargetArgs)(nil)).Elem()
}

// A collection of values returned by getEvpnVniRouteTarget.
type LookupEvpnVniRouteTargetResultOutput struct{ *pulumi.OutputState }

func (LookupEvpnVniRouteTargetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEvpnVniRouteTargetResult)(nil)).Elem()
}

func (o LookupEvpnVniRouteTargetResultOutput) ToLookupEvpnVniRouteTargetResultOutput() LookupEvpnVniRouteTargetResultOutput {
	return o
}

func (o LookupEvpnVniRouteTargetResultOutput) ToLookupEvpnVniRouteTargetResultOutputWithContext(ctx context.Context) LookupEvpnVniRouteTargetResultOutput {
	return o
}

func (o LookupEvpnVniRouteTargetResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupEvpnVniRouteTargetResult] {
	return pulumix.Output[LookupEvpnVniRouteTargetResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupEvpnVniRouteTargetResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEvpnVniRouteTargetResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupEvpnVniRouteTargetResultOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEvpnVniRouteTargetResult) string { return v.Direction }).(pulumi.StringOutput)
}

func (o LookupEvpnVniRouteTargetResultOutput) Encap() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEvpnVniRouteTargetResult) string { return v.Encap }).(pulumi.StringOutput)
}

func (o LookupEvpnVniRouteTargetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEvpnVniRouteTargetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEvpnVniRouteTargetResultOutput) RouteTarget() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEvpnVniRouteTargetResult) string { return v.RouteTarget }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEvpnVniRouteTargetResultOutput{})
}