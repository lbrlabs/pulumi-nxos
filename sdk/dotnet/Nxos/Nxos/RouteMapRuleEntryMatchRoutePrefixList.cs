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
    [NxosResourceType("nxos:nxos/routeMapRuleEntryMatchRoutePrefixList:RouteMapRuleEntryMatchRoutePrefixList")]
    public partial class RouteMapRuleEntryMatchRoutePrefixList : global::Pulumi.CustomResource
    {
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
        /// DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
        /// </summary>
        [Output("prefixListDn")]
        public Output<string> PrefixListDn { get; private set; } = null!;

        /// <summary>
        /// Route Map rule name.
        /// </summary>
        [Output("ruleName")]
        public Output<string> RuleName { get; private set; } = null!;


        /// <summary>
        /// Create a RouteMapRuleEntryMatchRoutePrefixList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteMapRuleEntryMatchRoutePrefixList(string name, RouteMapRuleEntryMatchRoutePrefixListArgs args, CustomResourceOptions? options = null)
            : base("nxos:nxos/routeMapRuleEntryMatchRoutePrefixList:RouteMapRuleEntryMatchRoutePrefixList", name, args ?? new RouteMapRuleEntryMatchRoutePrefixListArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteMapRuleEntryMatchRoutePrefixList(string name, Input<string> id, RouteMapRuleEntryMatchRoutePrefixListState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/routeMapRuleEntryMatchRoutePrefixList:RouteMapRuleEntryMatchRoutePrefixList", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouteMapRuleEntryMatchRoutePrefixList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteMapRuleEntryMatchRoutePrefixList Get(string name, Input<string> id, RouteMapRuleEntryMatchRoutePrefixListState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteMapRuleEntryMatchRoutePrefixList(name, id, state, options);
        }
    }

    public sealed class RouteMapRuleEntryMatchRoutePrefixListArgs : global::Pulumi.ResourceArgs
    {
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
        /// DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
        /// </summary>
        [Input("prefixListDn", required: true)]
        public Input<string> PrefixListDn { get; set; } = null!;

        /// <summary>
        /// Route Map rule name.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        public RouteMapRuleEntryMatchRoutePrefixListArgs()
        {
        }
        public static new RouteMapRuleEntryMatchRoutePrefixListArgs Empty => new RouteMapRuleEntryMatchRoutePrefixListArgs();
    }

    public sealed class RouteMapRuleEntryMatchRoutePrefixListState : global::Pulumi.ResourceArgs
    {
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
        /// DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
        /// </summary>
        [Input("prefixListDn")]
        public Input<string>? PrefixListDn { get; set; }

        /// <summary>
        /// Route Map rule name.
        /// </summary>
        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        public RouteMapRuleEntryMatchRoutePrefixListState()
        {
        }
        public static new RouteMapRuleEntryMatchRoutePrefixListState Empty => new RouteMapRuleEntryMatchRoutePrefixListState();
    }
}
