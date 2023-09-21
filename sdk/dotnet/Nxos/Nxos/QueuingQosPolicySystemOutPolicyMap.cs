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
    [NxosResourceType("nxos:nxos/queuingQosPolicySystemOutPolicyMap:QueuingQosPolicySystemOutPolicyMap")]
    public partial class QueuingQosPolicySystemOutPolicyMap : global::Pulumi.CustomResource
    {
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
        /// Create a QueuingQosPolicySystemOutPolicyMap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QueuingQosPolicySystemOutPolicyMap(string name, QueuingQosPolicySystemOutPolicyMapArgs args, CustomResourceOptions? options = null)
            : base("nxos:nxos/queuingQosPolicySystemOutPolicyMap:QueuingQosPolicySystemOutPolicyMap", name, args ?? new QueuingQosPolicySystemOutPolicyMapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private QueuingQosPolicySystemOutPolicyMap(string name, Input<string> id, QueuingQosPolicySystemOutPolicyMapState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/queuingQosPolicySystemOutPolicyMap:QueuingQosPolicySystemOutPolicyMap", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing QueuingQosPolicySystemOutPolicyMap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QueuingQosPolicySystemOutPolicyMap Get(string name, Input<string> id, QueuingQosPolicySystemOutPolicyMapState? state = null, CustomResourceOptions? options = null)
        {
            return new QueuingQosPolicySystemOutPolicyMap(name, id, state, options);
        }
    }

    public sealed class QueuingQosPolicySystemOutPolicyMapArgs : global::Pulumi.ResourceArgs
    {
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

        public QueuingQosPolicySystemOutPolicyMapArgs()
        {
        }
        public static new QueuingQosPolicySystemOutPolicyMapArgs Empty => new QueuingQosPolicySystemOutPolicyMapArgs();
    }

    public sealed class QueuingQosPolicySystemOutPolicyMapState : global::Pulumi.ResourceArgs
    {
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

        public QueuingQosPolicySystemOutPolicyMapState()
        {
        }
        public static new QueuingQosPolicySystemOutPolicyMapState Empty => new QueuingQosPolicySystemOutPolicyMapState();
    }
}