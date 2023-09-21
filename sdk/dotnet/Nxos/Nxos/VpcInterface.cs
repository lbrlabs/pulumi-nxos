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
    [NxosResourceType("nxos:nxos/vpcInterface:VpcInterface")]
    public partial class VpcInterface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Port-channel interface DN.
        /// </summary>
        [Output("portChannelInterfaceDn")]
        public Output<string?> PortChannelInterfaceDn { get; private set; } = null!;

        /// <summary>
        /// The vPC interface identifier. - Range: `1`-`16384`
        /// </summary>
        [Output("vpcInterfaceId")]
        public Output<int> VpcInterfaceId { get; private set; } = null!;


        /// <summary>
        /// Create a VpcInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpcInterface(string name, VpcInterfaceArgs args, CustomResourceOptions? options = null)
            : base("nxos:nxos/vpcInterface:VpcInterface", name, args ?? new VpcInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpcInterface(string name, Input<string> id, VpcInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/vpcInterface:VpcInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VpcInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpcInterface Get(string name, Input<string> id, VpcInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new VpcInterface(name, id, state, options);
        }
    }

    public sealed class VpcInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Port-channel interface DN.
        /// </summary>
        [Input("portChannelInterfaceDn")]
        public Input<string>? PortChannelInterfaceDn { get; set; }

        /// <summary>
        /// The vPC interface identifier. - Range: `1`-`16384`
        /// </summary>
        [Input("vpcInterfaceId", required: true)]
        public Input<int> VpcInterfaceId { get; set; } = null!;

        public VpcInterfaceArgs()
        {
        }
        public static new VpcInterfaceArgs Empty => new VpcInterfaceArgs();
    }

    public sealed class VpcInterfaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Port-channel interface DN.
        /// </summary>
        [Input("portChannelInterfaceDn")]
        public Input<string>? PortChannelInterfaceDn { get; set; }

        /// <summary>
        /// The vPC interface identifier. - Range: `1`-`16384`
        /// </summary>
        [Input("vpcInterfaceId")]
        public Input<int>? VpcInterfaceId { get; set; }

        public VpcInterfaceState()
        {
        }
        public static new VpcInterfaceState Empty => new VpcInterfaceState();
    }
}