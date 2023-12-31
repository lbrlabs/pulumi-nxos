// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Nxos
{
    public static class GetBgpAdvertisedPrefix
    {
        /// <summary>
        /// This data source can read the BGP (VRF) advertised prefix configuration.
        /// 
        /// - API Documentation: [bgpAdvPrefix](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:AdvPrefix/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Nxos = Pulumi.Nxos;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Nxos.GetBgpAdvertisedPrefix.Invoke(new()
        ///     {
        ///         AddressFamily = "ipv4-ucast",
        ///         Asn = "65001",
        ///         Prefix = "192.168.1.0/24",
        ///         Vrf = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBgpAdvertisedPrefixResult> InvokeAsync(GetBgpAdvertisedPrefixArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBgpAdvertisedPrefixResult>("nxos:index/getBgpAdvertisedPrefix:getBgpAdvertisedPrefix", args ?? new GetBgpAdvertisedPrefixArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the BGP (VRF) advertised prefix configuration.
        /// 
        /// - API Documentation: [bgpAdvPrefix](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:AdvPrefix/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Nxos = Pulumi.Nxos;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Nxos.GetBgpAdvertisedPrefix.Invoke(new()
        ///     {
        ///         AddressFamily = "ipv4-ucast",
        ///         Asn = "65001",
        ///         Prefix = "192.168.1.0/24",
        ///         Vrf = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBgpAdvertisedPrefixResult> Invoke(GetBgpAdvertisedPrefixInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBgpAdvertisedPrefixResult>("nxos:index/getBgpAdvertisedPrefix:getBgpAdvertisedPrefix", args ?? new GetBgpAdvertisedPrefixInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBgpAdvertisedPrefixArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Address Family.
        /// </summary>
        [Input("addressFamily", required: true)]
        public string AddressFamily { get; set; } = null!;

        /// <summary>
        /// Autonomous system number.
        /// </summary>
        [Input("asn", required: true)]
        public string Asn { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// IP address of the network or prefix to advertise.
        /// </summary>
        [Input("prefix", required: true)]
        public string Prefix { get; set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf", required: true)]
        public string Vrf { get; set; } = null!;

        public GetBgpAdvertisedPrefixArgs()
        {
        }
        public static new GetBgpAdvertisedPrefixArgs Empty => new GetBgpAdvertisedPrefixArgs();
    }

    public sealed class GetBgpAdvertisedPrefixInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Address Family.
        /// </summary>
        [Input("addressFamily", required: true)]
        public Input<string> AddressFamily { get; set; } = null!;

        /// <summary>
        /// Autonomous system number.
        /// </summary>
        [Input("asn", required: true)]
        public Input<string> Asn { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// IP address of the network or prefix to advertise.
        /// </summary>
        [Input("prefix", required: true)]
        public Input<string> Prefix { get; set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
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
        /// <summary>
        /// Address Family.
        /// </summary>
        public readonly string AddressFamily;
        /// <summary>
        /// Autonomous system number.
        /// </summary>
        public readonly string Asn;
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IP address of the network or prefix to advertise.
        /// </summary>
        public readonly string Prefix;
        /// <summary>
        /// Route map to modify attributes.
        /// </summary>
        public readonly string RouteMap;
        /// <summary>
        /// VRF name.
        /// </summary>
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
