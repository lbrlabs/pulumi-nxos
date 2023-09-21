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
    [NxosResourceType("nxos:nxos/sviInterface:SviInterface")]
    public partial class SviInterface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Administrative port state. - Choices: `up`, `down` - Default value: `up`
        /// </summary>
        [Output("adminState")]
        public Output<string> AdminState { get; private set; } = null!;

        /// <summary>
        /// Specifies the administrative port bandwidth. - Range: `1`-`400000000` - Default value: `1000000`
        /// </summary>
        [Output("bandwidth")]
        public Output<int> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
        /// </summary>
        [Output("delay")]
        public Output<int> Delay { get; private set; } = null!;

        /// <summary>
        /// Interface description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `vlan100`.
        /// </summary>
        [Output("interfaceId")]
        public Output<string> InterfaceId { get; private set; } = null!;

        /// <summary>
        /// The administrative port medium type. - Choices: `bcast`, `p2p` - Default value: `bcast`
        /// </summary>
        [Output("medium")]
        public Output<string> Medium { get; private set; } = null!;

        /// <summary>
        /// Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
        /// </summary>
        [Output("mtu")]
        public Output<int> Mtu { get; private set; } = null!;


        /// <summary>
        /// Create a SviInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SviInterface(string name, SviInterfaceArgs args, CustomResourceOptions? options = null)
            : base("nxos:nxos/sviInterface:SviInterface", name, args ?? new SviInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SviInterface(string name, Input<string> id, SviInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/sviInterface:SviInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SviInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SviInterface Get(string name, Input<string> id, SviInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new SviInterface(name, id, state, options);
        }
    }

    public sealed class SviInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative port state. - Choices: `up`, `down` - Default value: `up`
        /// </summary>
        [Input("adminState")]
        public Input<string>? AdminState { get; set; }

        /// <summary>
        /// Specifies the administrative port bandwidth. - Range: `1`-`400000000` - Default value: `1000000`
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
        /// </summary>
        [Input("delay")]
        public Input<int>? Delay { get; set; }

        /// <summary>
        /// Interface description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `vlan100`.
        /// </summary>
        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        /// <summary>
        /// The administrative port medium type. - Choices: `bcast`, `p2p` - Default value: `bcast`
        /// </summary>
        [Input("medium")]
        public Input<string>? Medium { get; set; }

        /// <summary>
        /// Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        public SviInterfaceArgs()
        {
        }
        public static new SviInterfaceArgs Empty => new SviInterfaceArgs();
    }

    public sealed class SviInterfaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative port state. - Choices: `up`, `down` - Default value: `up`
        /// </summary>
        [Input("adminState")]
        public Input<string>? AdminState { get; set; }

        /// <summary>
        /// Specifies the administrative port bandwidth. - Range: `1`-`400000000` - Default value: `1000000`
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
        /// </summary>
        [Input("delay")]
        public Input<int>? Delay { get; set; }

        /// <summary>
        /// Interface description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `vlan100`.
        /// </summary>
        [Input("interfaceId")]
        public Input<string>? InterfaceId { get; set; }

        /// <summary>
        /// The administrative port medium type. - Choices: `bcast`, `p2p` - Default value: `bcast`
        /// </summary>
        [Input("medium")]
        public Input<string>? Medium { get; set; }

        /// <summary>
        /// Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
        /// </summary>
        [Input("mtu")]
        public Input<int>? Mtu { get; set; }

        public SviInterfaceState()
        {
        }
        public static new SviInterfaceState Empty => new SviInterfaceState();
    }
}
