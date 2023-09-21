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
    public static class GetBgpAdvertisedPrefix
    {
        public static Task<GetBgpAdvertisedPrefixResult> InvokeAsync(GetBgpAdvertisedPrefixArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBgpAdvertisedPrefixResult>("nxos:nxos/getBgpAdvertisedPrefix:getBgpAdvertisedPrefix", args ?? new GetBgpAdvertisedPrefixArgs(), options.WithDefaults());

        public static Output<GetBgpAdvertisedPrefixResult> Invoke(GetBgpAdvertisedPrefixInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBgpAdvertisedPrefixResult>("nxos:nxos/getBgpAdvertisedPrefix:getBgpAdvertisedPrefix", args ?? new GetBgpAdvertisedPrefixInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBgpAdvertisedPrefixArgs : global::Pulumi.InvokeArgs
    {
        [Input("addressFamily", required: true)]
        public string AddressFamily { get; set; } = null!;

        [Input("asn", required: true)]
        public string Asn { get; set; } = null!;

        [Input("device")]
        public string? Device { get; set; }

        [Input("prefix", required: true)]
        public string Prefix { get; set; } = null!;

        [Input("vrf", required: true)]
        public string Vrf { get; set; } = null!;

        public GetBgpAdvertisedPrefixArgs()
        {
        }
        public static new GetBgpAdvertisedPrefixArgs Empty => new GetBgpAdvertisedPrefixArgs();
    }

    public sealed class GetBgpAdvertisedPrefixInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("addressFamily", required: true)]
        public Input<string> AddressFamily { get; set; } = null!;

        [Input("asn", required: true)]
        public Input<string> Asn { get; set; } = null!;

        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("prefix", required: true)]
        public Input<string> Prefix { get; set; } = null!;

        [Input("vrf", required: true)]
        public Input<string> Vrf { get; set; } = null!;

        public GetBgpAdvertisedPrefixInvokeArgs()
        {
        }
        public static new GetBgpAdvertisedPrefixInvokeArgs Empty => new GetBgpAdvertisedPrefixInvokeArgs();
    }


    [OutputType]
    public sealed class GetBgpAdvertisedPrefixResult
    {
        public readonly string AddressFamily;
        public readonly string Asn;
        public readonly string? Device;
        public readonly string Id;
        public readonly string Prefix;
        public readonly string RouteMap;
        public readonly string Vrf;

        [OutputConstructor]
        private GetBgpAdvertisedPrefixResult(
            string addressFamily,

            string asn,

            string? device,

            string id,

            string prefix,

            string routeMap,

            string vrf)
        {
            AddressFamily = addressFamily;
            Asn = asn;
            Device = device;
            Id = id;
            Prefix = prefix;
            RouteMap = routeMap;
            Vrf = vrf;
        }
    }
}