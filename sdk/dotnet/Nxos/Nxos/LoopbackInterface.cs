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
    [NxosResourceType("nxos:nxos/loopbackInterface:LoopbackInterface")]
    public partial class LoopbackInterface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Administrative state. - Choices: `up`, `down` - Default value: `up`
        /// </summary>
        [Output("adminState")]
        public Output<string> AdminState { get; private set; } = null!;

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
        /// Must match first field in the output of `show intf brief`. Example: `lo123`.
        /// </summary>
        [Output("interfaceId")]
        public Output<string> InterfaceId { get; private set; } = null!;


        /// <summary>
        /// Create a LoopbackInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoopbackInterface(string name, LoopbackInterfaceArgs args, CustomResourceOptions? options = null)
            : base("nxos:nxos/loopbackInterface:LoopbackInterface", name, args ?? new LoopbackInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoopbackInterface(string name, Input<string> id, LoopbackInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/loopbackInterface:LoopbackInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing LoopbackInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoopbackInterface Get(string name, Input<string> id, LoopbackInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new LoopbackInterface(name, id, state, options);
        }
    }

    public sealed class LoopbackInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative state. - Choices: `up`, `down` - Default value: `up`
        /// </summary>
        [Input("adminState")]
        public Input<string>? AdminState { get; set; }

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
        /// Must match first field in the output of `show intf brief`. Example: `lo123`.
        /// </summary>
        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        public LoopbackInterfaceArgs()
        {
        }
        public static new LoopbackInterfaceArgs Empty => new LoopbackInterfaceArgs();
    }

    public sealed class LoopbackInterfaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative state. - Choices: `up`, `down` - Default value: `up`
        /// </summary>
        [Input("adminState")]
        public Input<string>? AdminState { get; set; }

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
        /// Must match first field in the output of `show intf brief`. Example: `lo123`.
        /// </summary>
        [Input("interfaceId")]
        public Input<string>? InterfaceId { get; set; }

        public LoopbackInterfaceState()
        {
        }
        public static new LoopbackInterfaceState Empty => new LoopbackInterfaceState();
    }
}
