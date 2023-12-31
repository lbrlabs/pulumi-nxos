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
    /// <summary>
    /// This resource can manage the BGP (VRF) address family configuration.
    /// 
    /// - API Documentation: [bgpDomAf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:DomAf/)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Nxos = Lbrlabs.PulumiPackage.Nxos;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Nxos.BgpAddressFamily("example", new()
    ///     {
    ///         AddressFamily = "ipv4-ucast",
    ///         Asn = "65001",
    ///         CriticalNexthopTimeout = 1800,
    ///         NonCriticalNexthopTimeout = 1800,
    ///         Vrf = "default",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import nxos:index/bgpAddressFamily:BgpAddressFamily example "sys/bgp/inst/dom-[default]/af-[ipv4-ucast]"
    /// ```
    /// </summary>
    [NxosResourceType("nxos:index/bgpAddressFamily:BgpAddressFamily")]
    public partial class BgpAddressFamily : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
        /// `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
        /// `ipv4-mdt` - Default value: `ipv4-ucast`
        /// </summary>
        [Output("addressFamily")]
        public Output<string> AddressFamily { get; private set; } = null!;

        /// <summary>
        /// Autonomous system number.
        /// </summary>
        [Output("asn")]
        public Output<string> Asn { get; private set; } = null!;

        /// <summary>
        /// The next-hop address tracking delay timer for critical next-hop reachability routes. - Range: `1`-`4294967295` - Default
        /// value: `3000`
        /// </summary>
        [Output("criticalNexthopTimeout")]
        public Output<int> CriticalNexthopTimeout { get; private set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// The next-hop address tracking delay timer for non-critical next-hop reachability routes. - Range: `1`-`4294967295` -
        /// Default value: `10000`
        /// </summary>
        [Output("nonCriticalNexthopTimeout")]
        public Output<int> NonCriticalNexthopTimeout { get; private set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Output("vrf")]
        public Output<string> Vrf { get; private set; } = null!;


        /// <summary>
        /// Create a BgpAddressFamily resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BgpAddressFamily(string name, BgpAddressFamilyArgs args, CustomResourceOptions? options = null)
            : base("nxos:index/bgpAddressFamily:BgpAddressFamily", name, args ?? new BgpAddressFamilyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BgpAddressFamily(string name, Input<string> id, BgpAddressFamilyState? state = null, CustomResourceOptions? options = null)
            : base("nxos:index/bgpAddressFamily:BgpAddressFamily", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BgpAddressFamily resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BgpAddressFamily Get(string name, Input<string> id, BgpAddressFamilyState? state = null, CustomResourceOptions? options = null)
        {
            return new BgpAddressFamily(name, id, state, options);
        }
    }

    public sealed class BgpAddressFamilyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
        /// `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
        /// `ipv4-mdt` - Default value: `ipv4-ucast`
        /// </summary>
        [Input("addressFamily", required: true)]
        public Input<string> AddressFamily { get; set; } = null!;

        /// <summary>
        /// Autonomous system number.
        /// </summary>
        [Input("asn", required: true)]
        public Input<string> Asn { get; set; } = null!;

        /// <summary>
        /// The next-hop address tracking delay timer for critical next-hop reachability routes. - Range: `1`-`4294967295` - Default
        /// value: `3000`
        /// </summary>
        [Input("criticalNexthopTimeout")]
        public Input<int>? CriticalNexthopTimeout { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// The next-hop address tracking delay timer for non-critical next-hop reachability routes. - Range: `1`-`4294967295` -
        /// Default value: `10000`
        /// </summary>
        [Input("nonCriticalNexthopTimeout")]
        public Input<int>? NonCriticalNexthopTimeout { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf", required: true)]
        public Input<string> Vrf { get; set; } = null!;

        public BgpAddressFamilyArgs()
        {
        }
        public static new BgpAddressFamilyArgs Empty => new BgpAddressFamilyArgs();
    }

    public sealed class BgpAddressFamilyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
        /// `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
        /// `ipv4-mdt` - Default value: `ipv4-ucast`
        /// </summary>
        [Input("addressFamily")]
        public Input<string>? AddressFamily { get; set; }

        /// <summary>
        /// Autonomous system number.
        /// </summary>
        [Input("asn")]
        public Input<string>? Asn { get; set; }

        /// <summary>
        /// The next-hop address tracking delay timer for critical next-hop reachability routes. - Range: `1`-`4294967295` - Default
        /// value: `3000`
        /// </summary>
        [Input("criticalNexthopTimeout")]
        public Input<int>? CriticalNexthopTimeout { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// The next-hop address tracking delay timer for non-critical next-hop reachability routes. - Range: `1`-`4294967295` -
        /// Default value: `10000`
        /// </summary>
        [Input("nonCriticalNexthopTimeout")]
        public Input<int>? NonCriticalNexthopTimeout { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf")]
        public Input<string>? Vrf { get; set; }

        public BgpAddressFamilyState()
        {
        }
        public static new BgpAddressFamilyState Empty => new BgpAddressFamilyState();
    }
}
