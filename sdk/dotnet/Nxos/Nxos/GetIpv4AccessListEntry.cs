// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Nxos.Nxos
{
    public static class GetIpv4AccessListEntry
    {
        public static Task<GetIpv4AccessListEntryResult> InvokeAsync(GetIpv4AccessListEntryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpv4AccessListEntryResult>("nxos:nxos/getIpv4AccessListEntry:getIpv4AccessListEntry", args ?? new GetIpv4AccessListEntryArgs(), options.WithDefaults());

        public static Output<GetIpv4AccessListEntryResult> Invoke(GetIpv4AccessListEntryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpv4AccessListEntryResult>("nxos:nxos/getIpv4AccessListEntry:getIpv4AccessListEntry", args ?? new GetIpv4AccessListEntryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpv4AccessListEntryArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("sequenceNumber", required: true)]
        public int SequenceNumber { get; set; }

        public GetIpv4AccessListEntryArgs()
        {
        }
        public static new GetIpv4AccessListEntryArgs Empty => new GetIpv4AccessListEntryArgs();
    }

    public sealed class GetIpv4AccessListEntryInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("sequenceNumber", required: true)]
        public Input<int> SequenceNumber { get; set; } = null!;

        public GetIpv4AccessListEntryInvokeArgs()
        {
        }
        public static new GetIpv4AccessListEntryInvokeArgs Empty => new GetIpv4AccessListEntryInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpv4AccessListEntryResult
    {
        public readonly bool Ack;
        public readonly string Action;
        public readonly string DestinationAddressGroup;
        public readonly string DestinationPort1;
        public readonly string DestinationPort2;
        public readonly string DestinationPortGroup;
        public readonly string DestinationPortMask;
        public readonly string DestinationPortOperator;
        public readonly string DestinationPrefix;
        public readonly string DestinationPrefixLength;
        public readonly string DestinationPrefixMask;
        public readonly string? Device;
        public readonly int Dscp;
        public readonly bool Est;
        public readonly bool Fin;
        public readonly bool Fragment;
        public readonly string HttpOptionType;
        public readonly int IcmpCode;
        public readonly int IcmpType;
        public readonly string Id;
        public readonly bool Logging;
        public readonly string Name;
        public readonly string PacketLength1;
        public readonly string PacketLength2;
        public readonly string PacketLengthOperator;
        public readonly string Precedence;
        public readonly string Protocol;
        public readonly string ProtocolMask;
        public readonly bool Psh;
        public readonly string Redirect;
        public readonly string Remark;
        public readonly bool Rev;
        public readonly bool Rst;
        public readonly int SequenceNumber;
        public readonly string SourceAddressGroup;
        public readonly string SourcePort1;
        public readonly string SourcePort2;
        public readonly string SourcePortGroup;
        public readonly string SourcePortMask;
        public readonly string SourcePortOperator;
        public readonly string SourcePrefix;
        public readonly string SourcePrefixLength;
        public readonly string SourcePrefixMask;
        public readonly bool Syn;
        public readonly string TimeRange;
        public readonly int Ttl;
        public readonly bool Urg;
        public readonly int Vlan;
        public readonly string Vni;

        [OutputConstructor]
        private GetIpv4AccessListEntryResult(
            bool ack,

            string action,

            string destinationAddressGroup,

            string destinationPort1,

            string destinationPort2,

            string destinationPortGroup,

            string destinationPortMask,

            string destinationPortOperator,

            string destinationPrefix,

            string destinationPrefixLength,

            string destinationPrefixMask,

            string? device,

            int dscp,

            bool est,

            bool fin,

            bool fragment,

            string httpOptionType,

            int icmpCode,

            int icmpType,

            string id,

            bool logging,

            string name,

            string packetLength1,

            string packetLength2,

            string packetLengthOperator,

            string precedence,

            string protocol,

            string protocolMask,

            bool psh,

            string redirect,

            string remark,

            bool rev,

            bool rst,

            int sequenceNumber,

            string sourceAddressGroup,

            string sourcePort1,

            string sourcePort2,

            string sourcePortGroup,

            string sourcePortMask,

            string sourcePortOperator,

            string sourcePrefix,

            string sourcePrefixLength,

            string sourcePrefixMask,

            bool syn,

            string timeRange,

            int ttl,

            bool urg,

            int vlan,

            string vni)
        {
            Ack = ack;
            Action = action;
            DestinationAddressGroup = destinationAddressGroup;
            DestinationPort1 = destinationPort1;
            DestinationPort2 = destinationPort2;
            DestinationPortGroup = destinationPortGroup;
            DestinationPortMask = destinationPortMask;
            DestinationPortOperator = destinationPortOperator;
            DestinationPrefix = destinationPrefix;
            DestinationPrefixLength = destinationPrefixLength;
            DestinationPrefixMask = destinationPrefixMask;
            Device = device;
            Dscp = dscp;
            Est = est;
            Fin = fin;
            Fragment = fragment;
            HttpOptionType = httpOptionType;
            IcmpCode = icmpCode;
            IcmpType = icmpType;
            Id = id;
            Logging = logging;
            Name = name;
            PacketLength1 = packetLength1;
            PacketLength2 = packetLength2;
            PacketLengthOperator = packetLengthOperator;
            Precedence = precedence;
            Protocol = protocol;
            ProtocolMask = protocolMask;
            Psh = psh;
            Redirect = redirect;
            Remark = remark;
            Rev = rev;
            Rst = rst;
            SequenceNumber = sequenceNumber;
            SourceAddressGroup = sourceAddressGroup;
            SourcePort1 = sourcePort1;
            SourcePort2 = sourcePort2;
            SourcePortGroup = sourcePortGroup;
            SourcePortMask = sourcePortMask;
            SourcePortOperator = sourcePortOperator;
            SourcePrefix = sourcePrefix;
            SourcePrefixLength = sourcePrefixLength;
            SourcePrefixMask = sourcePrefixMask;
            Syn = syn;
            TimeRange = timeRange;
            Ttl = ttl;
            Urg = urg;
            Vlan = vlan;
            Vni = vni;
        }
    }
}