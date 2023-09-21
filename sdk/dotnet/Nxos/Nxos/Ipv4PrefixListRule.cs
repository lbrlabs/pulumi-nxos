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
    [NxosResourceType("nxos:nxos/ipv4PrefixListRule:Ipv4PrefixListRule")]
    public partial class Ipv4PrefixListRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// IPv4 Prefix List Rule name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Ipv4PrefixListRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipv4PrefixListRule(string name, Ipv4PrefixListRuleArgs? args = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/ipv4PrefixListRule:Ipv4PrefixListRule", name, args ?? new Ipv4PrefixListRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipv4PrefixListRule(string name, Input<string> id, Ipv4PrefixListRuleState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/ipv4PrefixListRule:Ipv4PrefixListRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipv4PrefixListRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipv4PrefixListRule Get(string name, Input<string> id, Ipv4PrefixListRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipv4PrefixListRule(name, id, state, options);
        }
    }

    public sealed class Ipv4PrefixListRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// IPv4 Prefix List Rule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Ipv4PrefixListRuleArgs()
        {
        }
        public static new Ipv4PrefixListRuleArgs Empty => new Ipv4PrefixListRuleArgs();
    }

    public sealed class Ipv4PrefixListRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// IPv4 Prefix List Rule name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Ipv4PrefixListRuleState()
        {
        }
        public static new Ipv4PrefixListRuleState Empty => new Ipv4PrefixListRuleState();
    }
}