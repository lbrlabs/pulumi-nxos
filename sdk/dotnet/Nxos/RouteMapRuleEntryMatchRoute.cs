// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Nxos
{
    /// <summary>
    /// This resource can manage a Match Route in  Route-Map Rule Entry configuration.
    /// 
    /// - API Documentation: [rtmapMatchRtDst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:MatchRtDst/)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Nxos = Lbrlabs.PulumiPackage.Nxos;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Nxos.RouteMapRuleEntryMatchRoute("example", new()
    ///     {
    ///         Order = 10,
    ///         RuleName = "RULE1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import nxos:index/routeMapRuleEntryMatchRoute:RouteMapRuleEntryMatchRoute example "sys/rpm/rtmap-[RULE1]/ent-[10]/mrtdst"
    /// ```
    /// </summary>
    [NxosResourceType("nxos:index/routeMapRuleEntryMatchRoute:RouteMapRuleEntryMatchRoute")]
    public partial class RouteMapRuleEntryMatchRoute : global::Pulumi.CustomResource
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
        /// Route Map rule name.
        /// </summary>
        [Output("ruleName")]
        public Output<string> RuleName { get; private set; } = null!;


        /// <summary>
        /// Create a RouteMapRuleEntryMatchRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RouteMapRuleEntryMatchRoute(string name, RouteMapRuleEntryMatchRouteArgs args, CustomResourceOptions? options = null)
            : base("nxos:index/routeMapRuleEntryMatchRoute:RouteMapRuleEntryMatchRoute", name, args ?? new RouteMapRuleEntryMatchRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RouteMapRuleEntryMatchRoute(string name, Input<string> id, RouteMapRuleEntryMatchRouteState? state = null, CustomResourceOptions? options = null)
            : base("nxos:index/routeMapRuleEntryMatchRoute:RouteMapRuleEntryMatchRoute", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RouteMapRuleEntryMatchRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RouteMapRuleEntryMatchRoute Get(string name, Input<string> id, RouteMapRuleEntryMatchRouteState? state = null, CustomResourceOptions? options = null)
        {
            return new RouteMapRuleEntryMatchRoute(name, id, state, options);
        }
    }

    public sealed class RouteMapRuleEntryMatchRouteArgs : global::Pulumi.ResourceArgs
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
        /// Route Map rule name.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        public RouteMapRuleEntryMatchRouteArgs()
        {
        }
        public static new RouteMapRuleEntryMatchRouteArgs Empty => new RouteMapRuleEntryMatchRouteArgs();
    }

    public sealed class RouteMapRuleEntryMatchRouteState : global::Pulumi.ResourceArgs
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
        /// Route Map rule name.
        /// </summary>
        [Input("ruleName")]
        public Input<string>? RuleName { get; set; }

        public RouteMapRuleEntryMatchRouteState()
        {
        }
        public static new RouteMapRuleEntryMatchRouteState Empty => new RouteMapRuleEntryMatchRouteState();
    }
}
