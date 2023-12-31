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

// This data source can read the BGP peer template address family configuration.
//
// - API Documentation: [bgpPeerAf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:PeerAf/)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-nxos/sdk/go/nxos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := nxos.LookupBgpPeerTemplateAddressFamily(ctx, &nxos.LookupBgpPeerTemplateAddressFamilyArgs{
//				AddressFamily: "ipv4-ucast",
//				Asn:           "65001",
//				TemplateName:  "SPINE-PEERS",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupBgpPeerTemplateAddressFamily(ctx *pulumi.Context, args *LookupBgpPeerTemplateAddressFamilyArgs, opts ...pulumi.InvokeOption) (*LookupBgpPeerTemplateAddressFamilyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBgpPeerTemplateAddressFamilyResult
	err := ctx.Invoke("nxos:index/getBgpPeerTemplateAddressFamily:getBgpPeerTemplateAddressFamily", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBgpPeerTemplateAddressFamily.
type LookupBgpPeerTemplateAddressFamilyArgs struct {
	AddressFamily string  `pulumi:"addressFamily"`
	Asn           string  `pulumi:"asn"`
	Device        *string `pulumi:"device"`
	TemplateName  string  `pulumi:"templateName"`
}

// A collection of values returned by getBgpPeerTemplateAddressFamily.
type LookupBgpPeerTemplateAddressFamilyResult struct {
	AddressFamily         string  `pulumi:"addressFamily"`
	Asn                   string  `pulumi:"asn"`
	Control               string  `pulumi:"control"`
	Device                *string `pulumi:"device"`
	Id                    string  `pulumi:"id"`
	SendCommunityExtended string  `pulumi:"sendCommunityExtended"`
	SendCommunityStandard string  `pulumi:"sendCommunityStandard"`
	TemplateName          string  `pulumi:"templateName"`
}

func LookupBgpPeerTemplateAddressFamilyOutput(ctx *pulumi.Context, args LookupBgpPeerTemplateAddressFamilyOutputArgs, opts ...pulumi.InvokeOption) LookupBgpPeerTemplateAddressFamilyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBgpPeerTemplateAddressFamilyResult, error) {
			args := v.(LookupBgpPeerTemplateAddressFamilyArgs)
			r, err := LookupBgpPeerTemplateAddressFamily(ctx, &args, opts...)
			var s LookupBgpPeerTemplateAddressFamilyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBgpPeerTemplateAddressFamilyResultOutput)
}

// A collection of arguments for invoking getBgpPeerTemplateAddressFamily.
type LookupBgpPeerTemplateAddressFamilyOutputArgs struct {
	AddressFamily pulumi.StringInput    `pulumi:"addressFamily"`
	Asn           pulumi.StringInput    `pulumi:"asn"`
	Device        pulumi.StringPtrInput `pulumi:"device"`
	TemplateName  pulumi.StringInput    `pulumi:"templateName"`
}

func (LookupBgpPeerTemplateAddressFamilyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpPeerTemplateAddressFamilyArgs)(nil)).Elem()
}

// A collection of values returned by getBgpPeerTemplateAddressFamily.
type LookupBgpPeerTemplateAddressFamilyResultOutput struct{ *pulumi.OutputState }

func (LookupBgpPeerTemplateAddressFamilyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBgpPeerTemplateAddressFamilyResult)(nil)).Elem()
}

func (o LookupBgpPeerTemplateAddressFamilyResultOutput) ToLookupBgpPeerTemplateAddressFamilyResultOutput() LookupBgpPeerTemplateAddressFamilyResultOutput {
	return o
}

func (o LookupBgpPeerTemplateAddressFamilyResultOutput) ToLookupBgpPeerTemplateAddressFamilyResultOutputWithContext(ctx context.Context) LookupBgpPeerTemplateAddressFamilyResultOutput {
	return o
}

func (o LookupBgpPeerTemplateAddressFamilyResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBgpPeerTemplateAddressFamilyResult] {
	return pulumix.Output[LookupBgpPeerTemplateAddressFamilyResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupBgpPeerTemplateAddressFamilyResultOutput) AddressFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateAddressFamilyResult) string { return v.AddressFamily }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateAddressFamilyResultOutput) Asn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateAddressFamilyResult) string { return v.Asn }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateAddressFamilyResultOutput) Control() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateAddressFamilyResult) string { return v.Control }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateAddressFamilyResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateAddressFamilyResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupBgpPeerTemplateAddressFamilyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateAddressFamilyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateAddressFamilyResultOutput) SendCommunityExtended() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateAddressFamilyResult) string { return v.SendCommunityExtended }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateAddressFamilyResultOutput) SendCommunityStandard() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateAddressFamilyResult) string { return v.SendCommunityStandard }).(pulumi.StringOutput)
}

func (o LookupBgpPeerTemplateAddressFamilyResultOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBgpPeerTemplateAddressFamilyResult) string { return v.TemplateName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBgpPeerTemplateAddressFamilyResultOutput{})
}
