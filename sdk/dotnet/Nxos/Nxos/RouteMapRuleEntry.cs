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
    [NxosResourceType("nxos:nxos/routeMapRuleEntry:RouteMapRuleEntry")]
    public partial class RouteMapRuleEntry : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Route-Map Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
        /// </summary>
        [Output("action")]
        public Output<string> Action { get; private set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Route-Map Rule Entry order. - Range: `0`-`65535`
        /// </summary>
        [Output("order")]
        public Output<int> Order { get; private set; } = null!;

        /// <summary>
        /// Route Map rule name.
        /// </summary>
        [Output("ruleName")]
        public Output<string> RuleName { get; private set; } = null!;


        /// <summary>
        /// Create a RouteMapRuleEntry resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteMapRuleEntry(string name, RouteMapRuleEntryArgs args, CustomResourceOptions? options = null)
            : base("nxos:nxos/routeMapRuleEntry:RouteMapRuleEntry", name, args ?? new RouteMapRuleEntryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteMapRuleEntry(string name, Input<string> id, RouteMapRuleEntryState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/routeMapRuleEntry:RouteMapRuleEntry", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouteMapRuleEntry resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteMapRuleEntry Get(string name, Input<string> id, RouteMapRuleEntryState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteMapRuleEntry(name, id, state, options);
        }
    }

    public sealed class RouteMapRuleEntryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Route-Map Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Route-Map Rule Entry order. - Range: `0`-`65535`
        /// </summary>
        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        /// <summary>
        /// Route Map rule name.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        public RouteMapRuleEntryArgs()
        {
        }
        public static new RouteMapRuleEntryArgs Empty => new RouteMapRuleEntryArgs();
    }

    public sealed class RouteMapRuleEntryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Route-Map Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Route-Map Rule Entry order. - Range: `0`-`65535`
        /// </summary>
        [Input("order")]
        public Input<int>? Order { get; set; }

        /// <summary>
        /// Route Map rule name.
        /// </summary>
        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        public RouteMapRuleEntryState()
        {
        }
        public static new RouteMapRuleEntryState Empty => new RouteMapRuleEntryState();
    }
}
