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

func LookupIpv4AccessListEntry(ctx *pulumi.Context, args *LookupIpv4AccessListEntryArgs, opts ...pulumi.InvokeOption) (*LookupIpv4AccessListEntryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpv4AccessListEntryResult
	err := ctx.Invoke("nxos:nxos/getIpv4AccessListEntry:getIpv4AccessListEntry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpv4AccessListEntry.
type LookupIpv4AccessListEntryArgs struct {
	Device         *string `pulumi:"device"`
	Name           string  `pulumi:"name"`
	SequenceNumber int     `pulumi:"sequenceNumber"`
}

// A collection of values returned by getIpv4AccessListEntry.
type LookupIpv4AccessListEntryResult struct {
	Ack                     bool    `pulumi:"ack"`
	Action                  string  `pulumi:"action"`
	DestinationAddressGroup string  `pulumi:"destinationAddressGroup"`
	DestinationPort1        string  `pulumi:"destinationPort1"`
	DestinationPort2        string  `pulumi:"destinationPort2"`
	DestinationPortGroup    string  `pulumi:"destinationPortGroup"`
	DestinationPortMask     string  `pulumi:"destinationPortMask"`
	DestinationPortOperator string  `pulumi:"destinationPortOperator"`
	DestinationPrefix       string  `pulumi:"destinationPrefix"`
	DestinationPrefixLength string  `pulumi:"destinationPrefixLength"`
	DestinationPrefixMask   string  `pulumi:"destinationPrefixMask"`
	Device                  *string `pulumi:"device"`
	Dscp                    int     `pulumi:"dscp"`
	Est                     bool    `pulumi:"est"`
	Fin                     bool    `pulumi:"fin"`
	Fragment                bool    `pulumi:"fragment"`
	HttpOptionType          string  `pulumi:"httpOptionType"`
	IcmpCode                int     `pulumi:"icmpCode"`
	IcmpType                int     `pulumi:"icmpType"`
	Id                      string  `pulumi:"id"`
	Logging                 bool    `pulumi:"logging"`
	Name                    string  `pulumi:"name"`
	PacketLength1           string  `pulumi:"packetLength1"`
	PacketLength2           string  `pulumi:"packetLength2"`
	PacketLengthOperator    string  `pulumi:"packetLengthOperator"`
	Precedence              string  `pulumi:"precedence"`
	Protocol                string  `pulumi:"protocol"`
	ProtocolMask            string  `pulumi:"protocolMask"`
	Psh                     bool    `pulumi:"psh"`
	Redirect                string  `pulumi:"redirect"`
	Remark                  string  `pulumi:"remark"`
	Rev                     bool    `pulumi:"rev"`
	Rst                     bool    `pulumi:"rst"`
	SequenceNumber          int     `pulumi:"sequenceNumber"`
	SourceAddressGroup      string  `pulumi:"sourceAddressGroup"`
	SourcePort1             string  `pulumi:"sourcePort1"`
	SourcePort2             string  `pulumi:"sourcePort2"`
	SourcePortGroup         string  `pulumi:"sourcePortGroup"`
	SourcePortMask          string  `pulumi:"sourcePortMask"`
	SourcePortOperator      string  `pulumi:"sourcePortOperator"`
	SourcePrefix            string  `pulumi:"sourcePrefix"`
	SourcePrefixLength      string  `pulumi:"sourcePrefixLength"`
	SourcePrefixMask        string  `pulumi:"sourcePrefixMask"`
	Syn                     bool    `pulumi:"syn"`
	TimeRange               string  `pulumi:"timeRange"`
	Ttl                     int     `pulumi:"ttl"`
	Urg                     bool    `pulumi:"urg"`
	Vlan                    int     `pulumi:"vlan"`
	Vni                     string  `pulumi:"vni"`
}

func LookupIpv4AccessListEntryOutput(ctx *pulumi.Context, args LookupIpv4AccessListEntryOutputArgs, opts ...pulumi.InvokeOption) LookupIpv4AccessListEntryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIpv4AccessListEntryResult, error) {
			args := v.(LookupIpv4AccessListEntryArgs)
			r, err := LookupIpv4AccessListEntry(ctx, &args, opts...)
			var s LookupIpv4AccessListEntryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIpv4AccessListEntryResultOutput)
}

// A collection of arguments for invoking getIpv4AccessListEntry.
type LookupIpv4AccessListEntryOutputArgs struct {
	Device         pulumi.StringPtrInput `pulumi:"device"`
	Name           pulumi.StringInput    `pulumi:"name"`
	SequenceNumber pulumi.IntInput       `pulumi:"sequenceNumber"`
}

func (LookupIpv4AccessListEntryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpv4AccessListEntryArgs)(nil)).Elem()
}

// A collection of values returned by getIpv4AccessListEntry.
type LookupIpv4AccessListEntryResultOutput struct{ *pulumi.OutputState }

func (LookupIpv4AccessListEntryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpv4AccessListEntryResult)(nil)).Elem()
}

func (o LookupIpv4AccessListEntryResultOutput) ToLookupIpv4AccessListEntryResultOutput() LookupIpv4AccessListEntryResultOutput {
	return o
}

func (o LookupIpv4AccessListEntryResultOutput) ToLookupIpv4AccessListEntryResultOutputWithContext(ctx context.Context) LookupIpv4AccessListEntryResultOutput {
	return o
}

func (o LookupIpv4AccessListEntryResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupIpv4AccessListEntryResult] {
	return pulumix.Output[LookupIpv4AccessListEntryResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupIpv4AccessListEntryResultOutput) Ack() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Ack }).(pulumi.BoolOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Action }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) DestinationAddressGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationAddressGroup }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) DestinationPort1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPort1 }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) DestinationPort2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPort2 }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) DestinationPortGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPortGroup }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) DestinationPortMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPortMask }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) DestinationPortOperator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPortOperator }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) DestinationPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPrefix }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) DestinationPrefixLength() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPrefixLength }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) DestinationPrefixMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPrefixMask }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Dscp() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) int { return v.Dscp }).(pulumi.IntOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Est() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Est }).(pulumi.BoolOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Fin() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Fin }).(pulumi.BoolOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Fragment() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Fragment }).(pulumi.BoolOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) HttpOptionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.HttpOptionType }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) IcmpCode() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) int { return v.IcmpCode }).(pulumi.IntOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) IcmpType() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) int { return v.IcmpType }).(pulumi.IntOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Logging() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Logging }).(pulumi.BoolOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) PacketLength1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.PacketLength1 }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) PacketLength2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.PacketLength2 }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) PacketLengthOperator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.PacketLengthOperator }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Precedence() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Precedence }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) ProtocolMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.ProtocolMask }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Psh() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Psh }).(pulumi.BoolOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Redirect() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Redirect }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Remark() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Remark }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Rev() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Rev }).(pulumi.BoolOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Rst() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Rst }).(pulumi.BoolOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) SequenceNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) int { return v.SequenceNumber }).(pulumi.IntOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) SourceAddressGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourceAddressGroup }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) SourcePort1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePort1 }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) SourcePort2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePort2 }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) SourcePortGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePortGroup }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) SourcePortMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePortMask }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) SourcePortOperator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePortOperator }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) SourcePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePrefix }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) SourcePrefixLength() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePrefixLength }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) SourcePrefixMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePrefixMask }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Syn() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Syn }).(pulumi.BoolOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) TimeRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.TimeRange }).(pulumi.StringOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) int { return v.Ttl }).(pulumi.IntOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Urg() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Urg }).(pulumi.BoolOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Vlan() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) int { return v.Vlan }).(pulumi.IntOutput)
}

func (o LookupIpv4AccessListEntryResultOutput) Vni() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Vni }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpv4AccessListEntryResultOutput{})
}
