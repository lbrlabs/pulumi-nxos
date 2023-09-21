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
    [NxosResourceType("nxos:nxos/bgpRouteControl:BgpRouteControl")]
    public partial class BgpRouteControl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Autonomous system number.
        /// </summary>
        [Output("asn")]
        public Output<string> Asn { get; private set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Enforce First AS For Ebgp. Can be configured only for VRF default. - Choices: `enabled`, `disabled` - Default value:
        /// `enabled`
        /// </summary>
        [Output("enforceFirstAs")]
        public Output<string> EnforceFirstAs { get; private set; } = null!;

        /// <summary>
        /// Accelerate the hardware updates for IP/IPv6 adjacencies for neighbor. Can be configured only for VRF default. - Choices:
        /// `enabled`, `disabled` - Default value: `disabled`
        /// </summary>
        [Output("fibAccelerate")]
        public Output<string> FibAccelerate { get; private set; } = null!;

        /// <summary>
        /// Log Neighbor Changes. - Choices: `enabled`, `disabled` - Default value: `disabled`
        /// </summary>
        [Output("logNeighborChanges")]
        public Output<string> LogNeighborChanges { get; private set; } = null!;

        /// <summary>
        /// Suppress Routes: Advertise only routes that are programmed in hardware to peers. Can be configured only for VRF default.
        /// - Choices: `enabled`, `disabled` - Default value: `enabled`
        /// </summary>
        [Output("suppressRoutes")]
        public Output<string> SuppressRoutes { get; private set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Output("vrf")]
        public Output<string> Vrf { get; private set; } = null!;


        /// <summary>
        /// Create a BgpRouteControl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BgpRouteControl(string name, BgpRouteControlArgs args, CustomResourceOptions? options = null)
            : base("nxos:nxos/bgpRouteControl:BgpRouteControl", name, args ?? new BgpRouteControlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BgpRouteControl(string name, Input<string> id, BgpRouteControlState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/bgpRouteControl:BgpRouteControl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BgpRouteControl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BgpRouteControl Get(string name, Input<string> id, BgpRouteControlState? state = null, CustomResourceOptions? options = null)
        {
            return new BgpRouteControl(name, id, state, options);
        }
    }

    public sealed class BgpRouteControlArgs : global::Pulumi.ResourceArgs
    {
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
        /// Enforce First AS For Ebgp. Can be configured only for VRF default. - Choices: `enabled`, `disabled` - Default value:
        /// `enabled`
        /// </summary>
        [Input("enforceFirstAs")]
        public Input<string>? EnforceFirstAs { get; set; }

        /// <summary>
        /// Accelerate the hardware updates for IP/IPv6 adjacencies for neighbor. Can be configured only for VRF default. - Choices:
        /// `enabled`, `disabled` - Default value: `disabled`
        /// </summary>
        [Input("fibAccelerate")]
        public Input<string>? FibAccelerate { get; set; }

        /// <summary>
        /// Log Neighbor Changes. - Choices: `enabled`, `disabled` - Default value: `disabled`
        /// </summary>
        [Input("logNeighborChanges")]
        public Input<string>? LogNeighborChanges { get; set; }

        /// <summary>
        /// Suppress Routes: Advertise only routes that are programmed in hardware to peers. Can be configured only for VRF default.
        /// - Choices: `enabled`, `disabled` - Default value: `enabled`
        /// </summary>
        [Input("suppressRoutes")]
        public Input<string>? SuppressRoutes { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf", required: true)]
        public Input<string> Vrf { get; set; } = null!;

        public BgpRouteControlArgs()
        {
        }
        public static new BgpRouteControlArgs Empty => new BgpRouteControlArgs();
    }

    public sealed class BgpRouteControlState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Autonomous system number.
        /// </summary>
        [Input("asn")]
        public Input<string>? Asn { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Enforce First AS For Ebgp. Can be configured only for VRF default. - Choices: `enabled`, `disabled` - Default value:
        /// `enabled`
        /// </summary>
        [Input("enforceFirstAs")]
        public Input<string>? EnforceFirstAs { get; set; }

        /// <summary>
        /// Accelerate the hardware updates for IP/IPv6 adjacencies for neighbor. Can be configured only for VRF default. - Choices:
        /// `enabled`, `disabled` - Default value: `disabled`
        /// </summary>
        [Input("fibAccelerate")]
        public Input<string>? FibAccelerate { get; set; }

        /// <summary>
        /// Log Neighbor Changes. - Choices: `enabled`, `disabled` - Default value: `disabled`
        /// </summary>
        [Input("logNeighborChanges")]
        public Input<string>? LogNeighborChanges { get; set; }

        /// <summary>
        /// Suppress Routes: Advertise only routes that are programmed in hardware to peers. Can be configured only for VRF default.
        /// - Choices: `enabled`, `disabled` - Default value: `enabled`
        /// </summary>
        [Input("suppressRoutes")]
        public Input<string>? SuppressRoutes { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf")]
        public Input<string>? Vrf { get; set; }

        public BgpRouteControlState()
        {
        }
        public static new BgpRouteControlState Empty => new BgpRouteControlState();
    }
}
