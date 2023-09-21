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

func LookupBgpPeerTemplateMaxPrefix(ctx *pulumi.Context, args *LookupBgpPeerTemplateMaxPrefixArgs, opts ...pulumi.InvokeOption) (*LookupBgpPeerTemplateMaxPrefixResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBgpPeerTemplateMaxPrefixResult
	err := ctx.Invoke("nxos:nxos/getBgpPeerTemplateMaxPrefix:getBgpPeerTemplateMaxPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBgpPeerTemplateMaxPrefix.
type LookupBgpPeerTemplateMaxPrefixArgs struct {
	AddressFamily string  `pulumi:"addressFamily"`
	Asn           string  `pulumi:"asn"`
	Device        *string `pulumi:"device"`
	TemplateName  string  `pulumi:"templateName"`
}

// A collection of values returned by getBgpPeerTemplateMaxPrefix.
type LookupBgpPeerTemplateMaxPrefixResult struct {
	Action        string  `pulumi:"action"`
	AddressFamily string  `pulumi:"addressFamily"`
	Asn           string  `pulumi:"asn"`
	Device        *string `pulumi:"device"`
	Id            string  `pulumi:"id"`
	MaximumPrefix int     `pulumi:"maximumPrefix"`
	RestartTime   int     `pulumi:"restartTime"`
	TemplateName  string  `pulumi:"templateName"`
	Threshold     int     `pulumi:"threshold"`
}

func LookupBgpPeerTemplateMaxPrefixOutput(ctx *pulumi.Context, args LookupBgpPeerTemplateMaxPrefixOutputArgs, opts ...pulumi.InvokeOption) LookupBgpPeerTemplateMaxPrefixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBgpPeerTemplateMaxPrefixResult, error) {
			args := v.(LookupBgpPeerTemplateMaxPrefixArgs)
			r, err := LookupBgpPeerTemplateMaxPrefix(ctx, &args, opts...)
			var s LookupBgpPeerTemplateMaxPrefixResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBgpPeerTemplateMaxPrefixResultOutput)
}

// A collection of arguments for invoking getBgpPeerTemplateMaxPrefix.
type LookupBgpPeerTemplateMaxPrefixOutputArgs struct {
	AddressFamily pulumi.StringInput    `pulumi:"addressFamily"`
	Asn           pulumi.StringInput    `pulumi:"asn"`
	Device        pulumi.StringPtrInput `pulumi:"device"`
	TemplateName  pulumi.StringInput    `pulumi:"templateName"`
}

func (LookupBgpPeerTemplateMaxPrefixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpPeerTemplateMaxPrefixArgs)(nil)).Elem()
}

// A collection of values returned by getBgpPeerTemplateMaxPrefix.
type LookupBgpPeerTemplateMaxPrefixResultOutput struct{ *pulumi.OutputState }

func (LookupBgpPeerTemplateMaxPrefixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpPeerTemplateMaxPrefixResult)(nil)).Elem()
}

func (o LookupBgpPeerTemplateMaxPrefixResultOutput) ToLookupBgpPeerTemplateMaxPrefixResultOutput() LookupBgpPeerTemplateMaxPrefixResultOutput {
	return o
}

func (o LookupBgpPeerTemplateMaxPrefixResultOutput) ToLookupBgpPeerTemplateMaxPrefixResultOutputWithContext(ctx context.Context) LookupBgpPeerTemplateMaxPrefixResultOutput {
	return o
}

func (o LookupBgpPeerTemplateMaxPrefixResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBgpPeerTemplateMaxPrefixResult] {
	return pulumix.Output[LookupBgpPeerTemplateMaxPrefixResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupBgpPeerTemplateMaxPrefixResultOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateMaxPrefixResult) string { return v.Action }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateMaxPrefixResultOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateMaxPrefixResult) string { return v.AddressFamily }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateMaxPrefixResultOutput) Asn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateMaxPrefixResult) string { return v.Asn }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateMaxPrefixResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateMaxPrefixResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupBgpPeerTemplateMaxPrefixResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateMaxPrefixResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateMaxPrefixResultOutput) MaximumPrefix() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateMaxPrefixResult) int { return v.MaximumPrefix }).(pulumi.IntOutput)
}

func (o LookupBgpPeerTemplateMaxPrefixResultOutput) RestartTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateMaxPrefixResult) int { return v.RestartTime }).(pulumi.IntOutput)
}

func (o LookupBgpPeerTemplateMaxPrefixResultOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateMaxPrefixResult) string { return v.TemplateName }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateMaxPrefixResultOutput) Threshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateMaxPrefixResult) int { return v.Threshold }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBgpPeerTemplateMaxPrefixResultOutput{})
}
