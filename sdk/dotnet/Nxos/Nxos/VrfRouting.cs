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
    [NxosResourceType("nxos:nxos/vrfRouting:VrfRouting")]
    public partial class VrfRouting : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Route Distinguisher value in NX-OS DME format. - Default value: `unknown:unknown:0:0`
        /// </summary>
        [Output("routeDistinguisher")]
        public Output<string> RouteDistinguisher { get; private set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Output("vrf")]
        public Output<string> Vrf { get; private set; } = null!;


        /// <summary>
        /// Create a VrfRouting resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VrfRouting(string name, VrfRoutingArgs args, CustomResourceOptions? options = null)
            : base("nxos:nxos/vrfRouting:VrfRouting", name, args ?? new VrfRoutingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VrfRouting(string name, Input<string> id, VrfRoutingState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/vrfRouting:VrfRouting", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VrfRouting resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VrfRouting Get(string name, Input<string> id, VrfRoutingState? state = null, CustomResourceOptions? options = null)
        {
            return new VrfRouting(name, id, state, options);
        }
    }

    public sealed class VrfRoutingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Route Distinguisher value in NX-OS DME format. - Default value: `unknown:unknown:0:0`
        /// </summary>
        [Input("routeDistinguisher")]
        public Input<string>? RouteDistinguisher { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf", required: true)]
        public Input<string> Vrf { get; set; } = null!;

        public VrfRoutingArgs()
        {
        }
        public static new VrfRoutingArgs Empty => new VrfRoutingArgs();
    }

    public sealed class VrfRoutingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Route Distinguisher value in NX-OS DME format. - Default value: `unknown:unknown:0:0`
        /// </summary>
        [Input("routeDistinguisher")]
        public Input<string>? RouteDistinguisher { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf")]
        public Input<string>? Vrf { get; set; }

        public VrfRoutingState()
        {
        }
        public static new VrfRoutingState Empty => new VrfRoutingState();
    }
}