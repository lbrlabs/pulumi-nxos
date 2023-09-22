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
    public static class GetRouteMapRuleEntrySetRegularCommunity
    {
        /// <summary>
        /// This data source can read a Set Community configuration in a Route-Map Rule Entry.
        /// 
        /// - API Documentation: [rtmapSetRegComm](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:SetRegComm/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Nxos = Pulumi.Nxos;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Nxos.GetRouteMapRuleEntrySetRegularCommunity.Invoke(new()
        ///     {
        ///         Order = 10,
        ///         RuleName = "RULE1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRouteMapRuleEntrySetRegularCommunityResult> InvokeAsync(GetRouteMapRuleEntrySetRegularCommunityArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouteMapRuleEntrySetRegularCommunityResult>("nxos:index/getRouteMapRuleEntrySetRegularCommunity:getRouteMapRuleEntrySetRegularCommunity", args ?? new GetRouteMapRuleEntrySetRegularCommunityArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read a Set Community configuration in a Route-Map Rule Entry.
        /// 
        /// - API Documentation: [rtmapSetRegComm](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:SetRegComm/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Nxos = Pulumi.Nxos;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Nxos.GetRouteMapRuleEntrySetRegularCommunity.Invoke(new()
        ///     {
        ///         Order = 10,
        ///         RuleName = "RULE1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRouteMapRuleEntrySetRegularCommunityResult> Invoke(GetRouteMapRuleEntrySetRegularCommunityInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouteMapRuleEntrySetRegularCommunityResult>("nxos:index/getRouteMapRuleEntrySetRegularCommunity:getRouteMapRuleEntrySetRegularCommunity", args ?? new GetRouteMapRuleEntrySetRegularCommunityInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouteMapRuleEntrySetRegularCommunityArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// Route-Map Rule Entry order.
        /// </summary>
        [Input("order", required: true)]
        public int Order { get; set; }

        /// <summary>
        /// Route Map rule name.
        /// </summary>
        [Input("ruleName", required: true)]
        public string RuleName { get; set; } = null!;

        public GetRouteMapRuleEntrySetRegularCommunityArgs()
        {
        }
        public static new GetRouteMapRuleEntrySetRegularCommunityArgs Empty => new GetRouteMapRuleEntrySetRegularCommunityArgs();
    }

    public sealed class GetRouteMapRuleEntrySetRegularCommunityInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Route-Map Rule Entry order.
        /// </summary>
        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        /// <summary>
        /// Route Map rule name.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        public GetRouteMapRuleEntrySetRegularCommunityInvokeArgs()
        {
        }
        public static new GetRouteMapRuleEntrySetRegularCommunityInvokeArgs Empty => new GetRouteMapRuleEntrySetRegularCommunityInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouteMapRuleEntrySetRegularCommunityResult
    {
        /// <summary>
        /// Option to add to an existing community.
        /// </summary>
        public readonly string Additive;
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Option to have no community attribute.
        /// </summary>
        public readonly string NoCommunity;
        /// <summary>
        /// Route-Map Rule Entry order.
        /// </summary>
        public readonly int Order;
        /// <summary>
        /// Route Map rule name.
        /// </summary>
        public readonly string RuleName;
        /// <summary>
        /// Operation on the current community.
        /// </summary>
        public readonly string SetCriteria;

        [OutputConstructor]
        private GetRouteMapRuleEntrySetRegularCommunityResult(
            string additive,

            string? device,

            string id,

            string noCommunity,

            int order,

            string ruleName,

            string setCriteria)
        {
            Additive = additive;
            Device = device;
            Id = id;
            NoCommunity = noCommunity;
            Order = order;
            RuleName = ruleName;
            SetCriteria = setCriteria;
        }
    }
}
