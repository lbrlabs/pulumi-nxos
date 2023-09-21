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

func LookupBgpPeerTemplate(ctx *pulumi.Context, args *LookupBgpPeerTemplateArgs, opts ...pulumi.InvokeOption) (*LookupBgpPeerTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBgpPeerTemplateResult
	err := ctx.Invoke("nxos:nxos/getBgpPeerTemplate:getBgpPeerTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBgpPeerTemplate.
type LookupBgpPeerTemplateArgs struct {
	Asn          string  `pulumi:"asn"`
	Device       *string `pulumi:"device"`
	TemplateName string  `pulumi:"templateName"`
}

// A collection of values returned by getBgpPeerTemplate.
type LookupBgpPeerTemplateResult struct {
	Asn             string  `pulumi:"asn"`
	Description     string  `pulumi:"description"`
	Device          *string `pulumi:"device"`
	Id              string  `pulumi:"id"`
	PeerType        string  `pulumi:"peerType"`
	RemoteAsn       string  `pulumi:"remoteAsn"`
	SourceInterface string  `pulumi:"sourceInterface"`
	TemplateName    string  `pulumi:"templateName"`
}

func LookupBgpPeerTemplateOutput(ctx *pulumi.Context, args LookupBgpPeerTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupBgpPeerTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBgpPeerTemplateResult, error) {
			args := v.(LookupBgpPeerTemplateArgs)
			r, err := LookupBgpPeerTemplate(ctx, &args, opts...)
			var s LookupBgpPeerTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBgpPeerTemplateResultOutput)
}

// A collection of arguments for invoking getBgpPeerTemplate.
type LookupBgpPeerTemplateOutputArgs struct {
	Asn          pulumi.StringInput    `pulumi:"asn"`
	Device       pulumi.StringPtrInput `pulumi:"device"`
	TemplateName pulumi.StringInput    `pulumi:"templateName"`
}

func (LookupBgpPeerTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpPeerTemplateArgs)(nil)).Elem()
}

// A collection of values returned by getBgpPeerTemplate.
type LookupBgpPeerTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupBgpPeerTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpPeerTemplateResult)(nil)).Elem()
}

func (o LookupBgpPeerTemplateResultOutput) ToLookupBgpPeerTemplateResultOutput() LookupBgpPeerTemplateResultOutput {
	return o
}

func (o LookupBgpPeerTemplateResultOutput) ToLookupBgpPeerTemplateResultOutputWithContext(ctx context.Context) LookupBgpPeerTemplateResultOutput {
	return o
}

func (o LookupBgpPeerTemplateResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBgpPeerTemplateResult] {
	return pulumix.Output[LookupBgpPeerTemplateResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupBgpPeerTemplateResultOutput) Asn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateResult) string { return v.Asn }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupBgpPeerTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateResultOutput) PeerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateResult) string { return v.PeerType }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateResultOutput) RemoteAsn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateResult) string { return v.RemoteAsn }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateResultOutput) SourceInterface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateResult) string { return v.SourceInterface }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateResultOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateResult) string { return v.TemplateName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBgpPeerTemplateResultOutput{})
}
