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
    [NxosResourceType("nxos:nxos/featureNvOverlay:FeatureNvOverlay")]
    public partial class FeatureNvOverlay : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Administrative state. - Choices: `enabled`, `disabled`
        /// </summary>
        [Output("adminState")]
        public Output<string> AdminState { get; private set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;


        /// <summary>
        /// Create a FeatureNvOverlay resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FeatureNvOverlay(string name, FeatureNvOverlayArgs args, CustomResourceOptions? options = null)
            : base("nxos:nxos/featureNvOverlay:FeatureNvOverlay", name, args ?? new FeatureNvOverlayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FeatureNvOverlay(string name, Input<string> id, FeatureNvOverlayState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/featureNvOverlay:FeatureNvOverlay", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FeatureNvOverlay resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FeatureNvOverlay Get(string name, Input<string> id, FeatureNvOverlayState? state = null, CustomResourceOptions? options = null)
        {
            return new FeatureNvOverlay(name, id, state, options);
        }
    }

    public sealed class FeatureNvOverlayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative state. - Choices: `enabled`, `disabled`
        /// </summary>
        [Input("adminState", required: true)]
        public Input<string> AdminState { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        public FeatureNvOverlayArgs()
        {
        }
        public static new FeatureNvOverlayArgs Empty => new FeatureNvOverlayArgs();
    }

    public sealed class FeatureNvOverlayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative state. - Choices: `enabled`, `disabled`
        /// </summary>
        [Input("adminState")]
        public Input<string>? AdminState { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        public FeatureNvOverlayState()
        {
        }
        public static new FeatureNvOverlayState Empty => new FeatureNvOverlayState();
    }
}
