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
    [NxosResourceType("nxos:nxos/queuingQosPolicyMapMatchClassMap:QueuingQosPolicyMapMatchClassMap")]
    public partial class QueuingQosPolicyMapMatchClassMap : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Class map name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Output("policyMapName")]
        public Output<string> PolicyMapName { get; private set; } = null!;


        /// <summary>
        /// Create a QueuingQosPolicyMapMatchClassMap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public QueuingQosPolicyMapMatchClassMap(string name, QueuingQosPolicyMapMatchClassMapArgs args, CustomResourceOptions? options = null)
            : base("nxos:nxos/queuingQosPolicyMapMatchClassMap:QueuingQosPolicyMapMatchClassMap", name, args ?? new QueuingQosPolicyMapMatchClassMapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private QueuingQosPolicyMapMatchClassMap(string name, Input<string> id, QueuingQosPolicyMapMatchClassMapState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/queuingQosPolicyMapMatchClassMap:QueuingQosPolicyMapMatchClassMap", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing QueuingQosPolicyMapMatchClassMap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static QueuingQosPolicyMapMatchClassMap Get(string name, Input<string> id, QueuingQosPolicyMapMatchClassMapState? state = null, CustomResourceOptions? options = null)
        {
            return new QueuingQosPolicyMapMatchClassMap(name, id, state, options);
        }
    }

    public sealed class QueuingQosPolicyMapMatchClassMapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Class map name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Input("policyMapName", required: true)]
        public Input<string> PolicyMapName { get; set; } = null!;

        public QueuingQosPolicyMapMatchClassMapArgs()
        {
        }
        public static new QueuingQosPolicyMapMatchClassMapArgs Empty => new QueuingQosPolicyMapMatchClassMapArgs();
    }

    public sealed class QueuingQosPolicyMapMatchClassMapState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Class map name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Input("policyMapName")]
        public Input<string>? PolicyMapName { get; set; }

        public QueuingQosPolicyMapMatchClassMapState()
        {
        }
        public static new QueuingQosPolicyMapMatchClassMapState Empty => new QueuingQosPolicyMapMatchClassMapState();
    }
}
