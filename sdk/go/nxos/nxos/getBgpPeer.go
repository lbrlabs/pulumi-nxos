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

func LookupBgpPeer(ctx *pulumi.Context, args *LookupBgpPeerArgs, opts ...pulumi.InvokeOption) (*LookupBgpPeerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBgpPeerResult
	err := ctx.Invoke("nxos:nxos/getBgpPeer:getBgpPeer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBgpPeer.
type LookupBgpPeerArgs struct {
	Address string  `pulumi:"address"`
	Asn     string  `pulumi:"asn"`
	Device  *string `pulumi:"device"`
	Vrf     string  `pulumi:"vrf"`
}

// A collection of values returned by getBgpPeer.
type LookupBgpPeerResult struct {
	Address         string  `pulumi:"address"`
	Asn             string  `pulumi:"asn"`
	Description     string  `pulumi:"description"`
	Device          *string `pulumi:"device"`
	HoldTime        int     `pulumi:"holdTime"`
	Id              string  `pulumi:"id"`
	Keepalive       int     `pulumi:"keepalive"`
	PeerTemplate    string  `pulumi:"peerTemplate"`
	PeerType        string  `pulumi:"peerType"`
	RemoteAsn       string  `pulumi:"remoteAsn"`
	SourceInterface string  `pulumi:"sourceInterface"`
	Vrf             string  `pulumi:"vrf"`
}

func LookupBgpPeerOutput(ctx *pulumi.Context, args LookupBgpPeerOutputArgs, opts ...pulumi.InvokeOption) LookupBgpPeerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBgpPeerResult, error) {
			args := v.(LookupBgpPeerArgs)
			r, err := LookupBgpPeer(ctx, &args, opts...)
			var s LookupBgpPeerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBgpPeerResultOutput)
}

// A collection of arguments for invoking getBgpPeer.
type LookupBgpPeerOutputArgs struct {
	Address pulumi.StringInput    `pulumi:"address"`
	Asn     pulumi.StringInput    `pulumi:"asn"`
	Device  pulumi.StringPtrInput `pulumi:"device"`
	Vrf     pulumi.StringInput    `pulumi:"vrf"`
}

func (LookupBgpPeerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpPeerArgs)(nil)).Elem()
}

// A collection of values returned by getBgpPeer.
type LookupBgpPeerResultOutput struct{ *pulumi.OutputState }

func (LookupBgpPeerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpPeerResult)(nil)).Elem()
}

func (o LookupBgpPeerResultOutput) ToLookupBgpPeerResultOutput() LookupBgpPeerResultOutput {
	return o
}

func (o LookupBgpPeerResultOutput) ToLookupBgpPeerResultOutputWithContext(ctx context.Context) LookupBgpPeerResultOutput {
	return o
}

func (o LookupBgpPeerResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBgpPeerResult] {
	return pulumix.Output[LookupBgpPeerResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupBgpPeerResultOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerResult) string { return v.Address }).(pulumi.StringOutput)
}

func (o LookupBgpPeerResultOutput) Asn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerResult) string { return v.Asn }).(pulumi.StringOutput)
}

func (o LookupBgpPeerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupBgpPeerResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBgpPeerResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupBgpPeerResultOutput) HoldTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpPeerResult) int { return v.HoldTime }).(pulumi.IntOutput)
}

func (o LookupBgpPeerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBgpPeerResultOutput) Keepalive() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBgpPeerResult) int { return v.Keepalive }).(pulumi.IntOutput)
}

func (o LookupBgpPeerResultOutput) PeerTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerResult) string { return v.PeerTemplate }).(pulumi.StringOutput)
}

func (o LookupBgpPeerResultOutput) PeerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerResult) string { return v.PeerType }).(pulumi.StringOutput)
}

func (o LookupBgpPeerResultOutput) RemoteAsn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerResult) string { return v.RemoteAsn }).(pulumi.StringOutput)
}

func (o LookupBgpPeerResultOutput) SourceInterface() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerResult) string { return v.SourceInterface }).(pulumi.StringOutput)
}

func (o LookupBgpPeerResultOutput) Vrf() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerResult) string { return v.Vrf }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBgpPeerResultOutput{})
}