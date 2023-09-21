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
    [NxosResourceType("nxos:nxos/queuingQosPolicyMapMatchClassMapRemainingBandwidth:QueuingQosPolicyMapMatchClassMapRemainingBandwidth")]
    public partial class QueuingQosPolicyMapMatchClassMapRemainingBandwidth : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Class map name.
        /// </summary>
        [Output("classMapName")]
        public Output<string> ClassMapName { get; private set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Output("policyMapName")]
        public Output<string> PolicyMapName { get; private set; } = null!;

        /// <summary>
        /// Remaining bandwidth percent. - Range: `0`-`100`
        /// </summary>
        [Output("value")]
        public Output<int> Value { get; private set; } = null!;


        /// <summary>
        /// Create a QueuingQosPolicyMapMatchClassMapRemainingBandwidth resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QueuingQosPolicyMapMatchClassMapRemainingBandwidth(string name, QueuingQosPolicyMapMatchClassMapRemainingBandwidthArgs args, CustomResourceOptions? options = null)
            : base("nxos:nxos/queuingQosPolicyMapMatchClassMapRemainingBandwidth:QueuingQosPolicyMapMatchClassMapRemainingBandwidth", name, args ?? new QueuingQosPolicyMapMatchClassMapRemainingBandwidthArgs(), MakeResourceOptions(options, ""))
        {
        }

        private QueuingQosPolicyMapMatchClassMapRemainingBandwidth(string name, Input<string> id, QueuingQosPolicyMapMatchClassMapRemainingBandwidthState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/queuingQosPolicyMapMatchClassMapRemainingBandwidth:QueuingQosPolicyMapMatchClassMapRemainingBandwidth", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing QueuingQosPolicyMapMatchClassMapRemainingBandwidth resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QueuingQosPolicyMapMatchClassMapRemainingBandwidth Get(string name, Input<string> id, QueuingQosPolicyMapMatchClassMapRemainingBandwidthState? state = null, CustomResourceOptions? options = null)
        {
            return new QueuingQosPolicyMapMatchClassMapRemainingBandwidth(name, id, state, options);
        }
    }

    public sealed class QueuingQosPolicyMapMatchClassMapRemainingBandwidthArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Class map name.
        /// </summary>
        [Input("classMapName", required: true)]
        public Input<string> ClassMapName { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Input("policyMapName", required: true)]
        public Input<string> PolicyMapName { get; set; } = null!;

        /// <summary>
        /// Remaining bandwidth percent. - Range: `0`-`100`
        /// </summary>
        [Input("value", required: true)]
        public Input<int> Value { get; set; } = null!;

        public QueuingQosPolicyMapMatchClassMapRemainingBandwidthArgs()
        {
        }
        public static new QueuingQosPolicyMapMatchClassMapRemainingBandwidthArgs Empty => new QueuingQosPolicyMapMatchClassMapRemainingBandwidthArgs();
    }

    public sealed class QueuingQosPolicyMapMatchClassMapRemainingBandwidthState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Class map name.
        /// </summary>
        [Input("classMapName")]
        public Input<string>? ClassMapName { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Input("policyMapName")]
        public Input<string>? PolicyMapName { get; set; }

        /// <summary>
        /// Remaining bandwidth percent. - Range: `0`-`100`
        /// </summary>
        [Input("value")]
        public Input<int>? Value { get; set; }

        public QueuingQosPolicyMapMatchClassMapRemainingBandwidthState()
        {
        }
        public static new QueuingQosPolicyMapMatchClassMapRemainingBandwidthState Empty => new QueuingQosPolicyMapMatchClassMapRemainingBandwidthState();
    }
}
