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
    [NxosResourceType("nxos:nxos/featureVpc:FeatureVpc")]
    public partial class FeatureVpc : global::Pulumi.CustomResource
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
        /// Create a FeatureVpc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FeatureVpc(string name, FeatureVpcArgs args, CustomResourceOptions? options = null)
            : base("nxos:nxos/featureVpc:FeatureVpc", name, args ?? new FeatureVpcArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FeatureVpc(string name, Input<string> id, FeatureVpcState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/featureVpc:FeatureVpc", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FeatureVpc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FeatureVpc Get(string name, Input<string> id, FeatureVpcState? state = null, CustomResourceOptions? options = null)
        {
            return new FeatureVpc(name, id, state, options);
        }
    }

    public sealed class FeatureVpcArgs : global::Pulumi.ResourceArgs
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

        public FeatureVpcArgs()
        {
        }
        public static new FeatureVpcArgs Empty => new FeatureVpcArgs();
    }

    public sealed class FeatureVpcState : global::Pulumi.ResourceArgs
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

        public FeatureVpcState()
        {
        }
        public static new FeatureVpcState Empty => new FeatureVpcState();
    }
}
