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

// This data source can read IPv4 Access List Entries.
//
// - API Documentation: [ipv4aclACE](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Security%20and%20Policing/ipv4acl:ACE/)
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
//			_, err := nxos.LookupIpv4AccessListEntry(ctx, &nxos.LookupIpv4AccessListEntryArgs{
//				Name:           "ACL1",
//				SequenceNumber: 10,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIpv4AccessListEntry(ctx *pulumi.Context, args *LookupIpv4AccessListEntryArgs, opts ...pulumi.InvokeOption) (*LookupIpv4AccessListEntryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpv4AccessListEntryResult
	err := ctx.Invoke("nxos:index/getIpv4AccessListEntry:getIpv4AccessListEntry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpv4AccessListEntry.
type LookupIpv4AccessListEntryArgs struct {
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Access list name.
	Name string `pulumi:"name"`
	// Sequence number.
	SequenceNumber int `pulumi:"sequenceNumber"`
}

// A collection of values returned by getIpv4AccessListEntry.
type LookupIpv4AccessListEntryResult struct {
	// Match TCP ACK flag.
	Ack bool `pulumi:"ack"`
	// Action.
	Action string `pulumi:"action"`
	// Destination address group.
	DestinationAddressGroup string `pulumi:"destinationAddressGroup"`
	// First destination port number or name.
	DestinationPort1 string `pulumi:"destinationPort1"`
	// Second destination port number or name.
	DestinationPort2 string `pulumi:"destinationPort2"`
	// Destination port group.
	DestinationPortGroup string `pulumi:"destinationPortGroup"`
	// Destination port mask number or name.
	DestinationPortMask string `pulumi:"destinationPortMask"`
	// Destination port operator.
	DestinationPortOperator string `pulumi:"destinationPortOperator"`
	// Destination prefix.
	DestinationPrefix string `pulumi:"destinationPrefix"`
	// Destination prefix length.
	DestinationPrefixLength string `pulumi:"destinationPrefixLength"`
	// Destination prefix mask.
	DestinationPrefixMask string `pulumi:"destinationPrefixMask"`
	// A device name from the provider configuration.
	Device *string `pulumi:"device"`
	// Match DSCP.
	Dscp int `pulumi:"dscp"`
	// Match TCP EST flag.
	Est bool `pulumi:"est"`
	// Match TCP FIN flag.
	Fin bool `pulumi:"fin"`
	// Match non-initial fragment.
	Fragment bool `pulumi:"fragment"`
	// HTTP option method.
	HttpOptionType string `pulumi:"httpOptionType"`
	// ICMP code.
	IcmpCode int `pulumi:"icmpCode"`
	// ICMP type.
	IcmpType int `pulumi:"icmpType"`
	// The distinguished name of the object.
	Id string `pulumi:"id"`
	// Log matches against ACL entry.
	Logging bool `pulumi:"logging"`
	// Access list name.
	Name string `pulumi:"name"`
	// First packet length. Either `invalid` or a number between 19 and 9210.
	PacketLength1 string `pulumi:"packetLength1"`
	// Second packet length. Either `invalid` or a number between 19 and 9210.
	PacketLength2 string `pulumi:"packetLength2"`
	// Packet length operator.
	PacketLengthOperator string `pulumi:"packetLengthOperator"`
	// Precedence. Either `unspecified` or a number between 0 and 7.
	Precedence string `pulumi:"precedence"`
	// Protocol name or number.
	Protocol string `pulumi:"protocol"`
	// Protocol mask name or number.
	ProtocolMask string `pulumi:"protocolMask"`
	// Match TCP PSH flag.
	Psh bool `pulumi:"psh"`
	// Redirect action.
	Redirect string `pulumi:"redirect"`
	// ACL comment.
	Remark string `pulumi:"remark"`
	// Match TCP REV flag.
	Rev bool `pulumi:"rev"`
	// Match TCP RST flag.
	Rst bool `pulumi:"rst"`
	// Sequence number.
	SequenceNumber int `pulumi:"sequenceNumber"`
	// Source address group.
	SourceAddressGroup string `pulumi:"sourceAddressGroup"`
	// First source port name or number.
	SourcePort1 string `pulumi:"sourcePort1"`
	// Second source port name or number.
	SourcePort2 string `pulumi:"sourcePort2"`
	// Source port group.
	SourcePortGroup string `pulumi:"sourcePortGroup"`
	// Source port mask name or number.
	SourcePortMask string `pulumi:"sourcePortMask"`
	// Source port operator.
	SourcePortOperator string `pulumi:"sourcePortOperator"`
	// Source prefix.
	SourcePrefix string `pulumi:"sourcePrefix"`
	// Source prefix length.
	SourcePrefixLength string `pulumi:"sourcePrefixLength"`
	// Source prefix mask.
	SourcePrefixMask string `pulumi:"sourcePrefixMask"`
	// Match TCP SYN flag.
	Syn bool `pulumi:"syn"`
	// Time range name.
	TimeRange string `pulumi:"timeRange"`
	// TTL.
	Ttl int `pulumi:"ttl"`
	// Match TCP URG flag.
	Urg bool `pulumi:"urg"`
	// VLAN ID.
	Vlan int `pulumi:"vlan"`
	// NVE VNI ID. Either `invalid` or a number between 0 and 16777216.
	Vni string `pulumi:"vni"`
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
	// A device name from the provider configuration.
	Device pulumi.StringPtrInput `pulumi:"device"`
	// Access list name.
	Name pulumi.StringInput `pulumi:"name"`
	// Sequence number.
	SequenceNumber pulumi.IntInput `pulumi:"sequenceNumber"`
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

// Match TCP ACK flag.
func (o LookupIpv4AccessListEntryResultOutput) Ack() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Ack }).(pulumi.BoolOutput)
}

// Action.
func (o LookupIpv4AccessListEntryResultOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Action }).(pulumi.StringOutput)
}

// Destination address group.
func (o LookupIpv4AccessListEntryResultOutput) DestinationAddressGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationAddressGroup }).(pulumi.StringOutput)
}

// First destination port number or name.
func (o LookupIpv4AccessListEntryResultOutput) DestinationPort1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPort1 }).(pulumi.StringOutput)
}

// Second destination port number or name.
func (o LookupIpv4AccessListEntryResultOutput) DestinationPort2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPort2 }).(pulumi.StringOutput)
}

// Destination port group.
func (o LookupIpv4AccessListEntryResultOutput) DestinationPortGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPortGroup }).(pulumi.StringOutput)
}

// Destination port mask number or name.
func (o LookupIpv4AccessListEntryResultOutput) DestinationPortMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPortMask }).(pulumi.StringOutput)
}

// Destination port operator.
func (o LookupIpv4AccessListEntryResultOutput) DestinationPortOperator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPortOperator }).(pulumi.StringOutput)
}

// Destination prefix.
func (o LookupIpv4AccessListEntryResultOutput) DestinationPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPrefix }).(pulumi.StringOutput)
}

// Destination prefix length.
func (o LookupIpv4AccessListEntryResultOutput) DestinationPrefixLength() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPrefixLength }).(pulumi.StringOutput)
}

// Destination prefix mask.
func (o LookupIpv4AccessListEntryResultOutput) DestinationPrefixMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.DestinationPrefixMask }).(pulumi.StringOutput)
}

// A device name from the provider configuration.
func (o LookupIpv4AccessListEntryResultOutput) Device() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) *string { return v.Device }).(pulumi.StringPtrOutput)
}

// Match DSCP.
func (o LookupIpv4AccessListEntryResultOutput) Dscp() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) int { return v.Dscp }).(pulumi.IntOutput)
}

// Match TCP EST flag.
func (o LookupIpv4AccessListEntryResultOutput) Est() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Est }).(pulumi.BoolOutput)
}

// Match TCP FIN flag.
func (o LookupIpv4AccessListEntryResultOutput) Fin() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Fin }).(pulumi.BoolOutput)
}

// Match non-initial fragment.
func (o LookupIpv4AccessListEntryResultOutput) Fragment() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Fragment }).(pulumi.BoolOutput)
}

// HTTP option method.
func (o LookupIpv4AccessListEntryResultOutput) HttpOptionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.HttpOptionType }).(pulumi.StringOutput)
}

// ICMP code.
func (o LookupIpv4AccessListEntryResultOutput) IcmpCode() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) int { return v.IcmpCode }).(pulumi.IntOutput)
}

// ICMP type.
func (o LookupIpv4AccessListEntryResultOutput) IcmpType() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) int { return v.IcmpType }).(pulumi.IntOutput)
}

// The distinguished name of the object.
func (o LookupIpv4AccessListEntryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Id }).(pulumi.StringOutput)
}

// Log matches against ACL entry.
func (o LookupIpv4AccessListEntryResultOutput) Logging() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Logging }).(pulumi.BoolOutput)
}

// Access list name.
func (o LookupIpv4AccessListEntryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Name }).(pulumi.StringOutput)
}

// First packet length. Either `invalid` or a number between 19 and 9210.
func (o LookupIpv4AccessListEntryResultOutput) PacketLength1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.PacketLength1 }).(pulumi.StringOutput)
}

// Second packet length. Either `invalid` or a number between 19 and 9210.
func (o LookupIpv4AccessListEntryResultOutput) PacketLength2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.PacketLength2 }).(pulumi.StringOutput)
}

// Packet length operator.
func (o LookupIpv4AccessListEntryResultOutput) PacketLengthOperator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.PacketLengthOperator }).(pulumi.StringOutput)
}

// Precedence. Either `unspecified` or a number between 0 and 7.
func (o LookupIpv4AccessListEntryResultOutput) Precedence() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Precedence }).(pulumi.StringOutput)
}

// Protocol name or number.
func (o LookupIpv4AccessListEntryResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// Protocol mask name or number.
func (o LookupIpv4AccessListEntryResultOutput) ProtocolMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.ProtocolMask }).(pulumi.StringOutput)
}

// Match TCP PSH flag.
func (o LookupIpv4AccessListEntryResultOutput) Psh() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Psh }).(pulumi.BoolOutput)
}

// Redirect action.
func (o LookupIpv4AccessListEntryResultOutput) Redirect() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Redirect }).(pulumi.StringOutput)
}

// ACL comment.
func (o LookupIpv4AccessListEntryResultOutput) Remark() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Remark }).(pulumi.StringOutput)
}

// Match TCP REV flag.
func (o LookupIpv4AccessListEntryResultOutput) Rev() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Rev }).(pulumi.BoolOutput)
}

// Match TCP RST flag.
func (o LookupIpv4AccessListEntryResultOutput) Rst() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Rst }).(pulumi.BoolOutput)
}

// Sequence number.
func (o LookupIpv4AccessListEntryResultOutput) SequenceNumber() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) int { return v.SequenceNumber }).(pulumi.IntOutput)
}

// Source address group.
func (o LookupIpv4AccessListEntryResultOutput) SourceAddressGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourceAddressGroup }).(pulumi.StringOutput)
}

// First source port name or number.
func (o LookupIpv4AccessListEntryResultOutput) SourcePort1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePort1 }).(pulumi.StringOutput)
}

// Second source port name or number.
func (o LookupIpv4AccessListEntryResultOutput) SourcePort2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePort2 }).(pulumi.StringOutput)
}

// Source port group.
func (o LookupIpv4AccessListEntryResultOutput) SourcePortGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePortGroup }).(pulumi.StringOutput)
}

// Source port mask name or number.
func (o LookupIpv4AccessListEntryResultOutput) SourcePortMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePortMask }).(pulumi.StringOutput)
}

// Source port operator.
func (o LookupIpv4AccessListEntryResultOutput) SourcePortOperator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePortOperator }).(pulumi.StringOutput)
}

// Source prefix.
func (o LookupIpv4AccessListEntryResultOutput) SourcePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePrefix }).(pulumi.StringOutput)
}

// Source prefix length.
func (o LookupIpv4AccessListEntryResultOutput) SourcePrefixLength() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePrefixLength }).(pulumi.StringOutput)
}

// Source prefix mask.
func (o LookupIpv4AccessListEntryResultOutput) SourcePrefixMask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.SourcePrefixMask }).(pulumi.StringOutput)
}

// Match TCP SYN flag.
func (o LookupIpv4AccessListEntryResultOutput) Syn() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Syn }).(pulumi.BoolOutput)
}

// Time range name.
func (o LookupIpv4AccessListEntryResultOutput) TimeRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.TimeRange }).(pulumi.StringOutput)
}

// TTL.
func (o LookupIpv4AccessListEntryResultOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) int { return v.Ttl }).(pulumi.IntOutput)
}

// Match TCP URG flag.
func (o LookupIpv4AccessListEntryResultOutput) Urg() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) bool { return v.Urg }).(pulumi.BoolOutput)
}

// VLAN ID.
func (o LookupIpv4AccessListEntryResultOutput) Vlan() pulumi.IntOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) int { return v.Vlan }).(pulumi.IntOutput)
}

// NVE VNI ID. Either `invalid` or a number between 0 and 16777216.
func (o LookupIpv4AccessListEntryResultOutput) Vni() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpv4AccessListEntryResult) string { return v.Vni }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpv4AccessListEntryResultOutput{})
}
