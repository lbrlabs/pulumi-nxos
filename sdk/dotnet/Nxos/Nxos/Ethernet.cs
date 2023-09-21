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
    [NxosResourceType("nxos:nxos/ethernet:Ethernet")]
    public partial class Ethernet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// System jumbo MTU. - Range: `576`-`9216` - Default value: `9216`
        /// </summary>
        [Output("mtu")]
        public Output<int> Mtu { get; private set; } = null!;


        /// <summary>
        /// Create a Ethernet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ethernet(string name, EthernetArgs? args = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/ethernet:Ethernet", name, args ?? new EthernetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ethernet(string name, Input<string> id, EthernetState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/ethernet:Ethernet", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ethernet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ethernet Get(string name, Input<string> id, EthernetState? state = null, CustomResourceOptions? options = null)
        {
            return new Ethernet(name, id, state, options);
        }
    }

    public sealed class EthernetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// System jumbo MTU. - Range: `576`-`9216` - Default value: `9216`
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        public EthernetArgs()
        {
        }
        public static new EthernetArgs Empty => new EthernetArgs();
    }

    public sealed class EthernetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// System jumbo MTU. - Range: `576`-`9216` - Default value: `9216`
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        public EthernetState()
        {
        }
        public static new EthernetState Empty => new EthernetState();
    }
}
