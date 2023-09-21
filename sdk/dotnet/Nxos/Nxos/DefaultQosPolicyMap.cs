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
    [NxosResourceType("nxos:nxos/defaultQosPolicyMap:DefaultQosPolicyMap")]
    public partial class DefaultQosPolicyMap : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
        /// </summary>
        [Output("matchType")]
        public Output<string> MatchType { get; private set; } = null!;

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a DefaultQosPolicyMap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DefaultQosPolicyMap(string name, DefaultQosPolicyMapArgs? args = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/defaultQosPolicyMap:DefaultQosPolicyMap", name, args ?? new DefaultQosPolicyMapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DefaultQosPolicyMap(string name, Input<string> id, DefaultQosPolicyMapState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/defaultQosPolicyMap:DefaultQosPolicyMap", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DefaultQosPolicyMap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DefaultQosPolicyMap Get(string name, Input<string> id, DefaultQosPolicyMapState? state = null, CustomResourceOptions? options = null)
        {
            return new DefaultQosPolicyMap(name, id, state, options);
        }
    }

    public sealed class DefaultQosPolicyMapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
        /// </summary>
        [Input("matchType")]
        public Input<string>? MatchType { get; set; }

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DefaultQosPolicyMapArgs()
        {
        }
        public static new DefaultQosPolicyMapArgs Empty => new DefaultQosPolicyMapArgs();
    }

    public sealed class DefaultQosPolicyMapState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
        /// </summary>
        [Input("matchType")]
        public Input<string>? MatchType { get; set; }

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DefaultQosPolicyMapState()
        {
        }
        public static new DefaultQosPolicyMapState Empty => new DefaultQosPolicyMapState();
    }
}
