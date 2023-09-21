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
    [NxosResourceType("nxos:nxos/evpnVniRouteTarget:EvpnVniRouteTarget")]
    public partial class EvpnVniRouteTarget : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Route Target direction. - Choices: `import`, `export`
        /// </summary>
        [Output("direction")]
        public Output<string> Direction { get; private set; } = null!;

        /// <summary>
        /// Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
        /// </summary>
        [Output("encap")]
        public Output<string> Encap { get; private set; } = null!;

        /// <summary>
        /// Route Target in NX-OS DME format.
        /// </summary>
        [Output("routeTarget")]
        public Output<string> RouteTarget { get; private set; } = null!;


        /// <summary>
        /// Create a EvpnVniRouteTarget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EvpnVniRouteTarget(string name, EvpnVniRouteTargetArgs args, CustomResourceOptions? options = null)
            : base("nxos:nxos/evpnVniRouteTarget:EvpnVniRouteTarget", name, args ?? new EvpnVniRouteTargetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EvpnVniRouteTarget(string name, Input<string> id, EvpnVniRouteTargetState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/evpnVniRouteTarget:EvpnVniRouteTarget", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EvpnVniRouteTarget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EvpnVniRouteTarget Get(string name, Input<string> id, EvpnVniRouteTargetState? state = null, CustomResourceOptions? options = null)
        {
            return new EvpnVniRouteTarget(name, id, state, options);
        }
    }

    public sealed class EvpnVniRouteTargetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Route Target direction. - Choices: `import`, `export`
        /// </summary>
        [Input("direction", required: true)]
        public Input<string> Direction { get; set; } = null!;

        /// <summary>
        /// Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
        /// </summary>
        [Input("encap", required: true)]
        public Input<string> Encap { get; set; } = null!;

        /// <summary>
        /// Route Target in NX-OS DME format.
        /// </summary>
        [Input("routeTarget", required: true)]
        public Input<string> RouteTarget { get; set; } = null!;

        public EvpnVniRouteTargetArgs()
        {
        }
        public static new EvpnVniRouteTargetArgs Empty => new EvpnVniRouteTargetArgs();
    }

    public sealed class EvpnVniRouteTargetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Route Target direction. - Choices: `import`, `export`
        /// </summary>
        [Input("direction")]
        public Input<string>? Direction { get; set; }

        /// <summary>
        /// Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
        /// </summary>
        [Input("encap")]
        public Input<string>? Encap { get; set; }

        /// <summary>
        /// Route Target in NX-OS DME format.
        /// </summary>
        [Input("routeTarget")]
        public Input<string>? RouteTarget { get; set; }

        public EvpnVniRouteTargetState()
        {
        }
        public static new EvpnVniRouteTargetState Empty => new EvpnVniRouteTargetState();
    }
}
