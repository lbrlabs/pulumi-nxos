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
    public static class GetRouteMapRuleEntryMatchRoutePrefixList
    {
        /// <summary>
        /// This data source can read a Match Route Prefix List in Route-Map Rule Entry configuration.
        /// 
        /// - API Documentation: [rtmapRsRtDstAtt](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:RsRtDstAtt/)
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
        ///     var example = Nxos.GetRouteMapRuleEntryMatchRoutePrefixList.Invoke(new()
        ///     {
        ///         Order = 10,
        ///         PrefixListDn = "sys/rpm/pfxlistv4-[LIST1]",
        ///         RuleName = "RULE1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRouteMapRuleEntryMatchRoutePrefixListResult> InvokeAsync(GetRouteMapRuleEntryMatchRoutePrefixListArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouteMapRuleEntryMatchRoutePrefixListResult>("nxos:index/getRouteMapRuleEntryMatchRoutePrefixList:getRouteMapRuleEntryMatchRoutePrefixList", args ?? new GetRouteMapRuleEntryMatchRoutePrefixListArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read a Match Route Prefix List in Route-Map Rule Entry configuration.
        /// 
        /// - API Documentation: [rtmapRsRtDstAtt](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:RsRtDstAtt/)
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
        ///     var example = Nxos.GetRouteMapRuleEntryMatchRoutePrefixList.Invoke(new()
        ///     {
        ///         Order = 10,
        ///         PrefixListDn = "sys/rpm/pfxlistv4-[LIST1]",
        ///         RuleName = "RULE1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRouteMapRuleEntryMatchRoutePrefixListResult> Invoke(GetRouteMapRuleEntryMatchRoutePrefixListInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouteMapRuleEntryMatchRoutePrefixListResult>("nxos:index/getRouteMapRuleEntryMatchRoutePrefixList:getRouteMapRuleEntryMatchRoutePrefixList", args ?? new GetRouteMapRuleEntryMatchRoutePrefixListInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouteMapRuleEntryMatchRoutePrefixListArgs : global::Pulumi.InvokeArgs
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
        /// DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
        /// </summary>
        [Input("prefixListDn", required: true)]
        public string PrefixListDn { get; set; } = null!;

        /// <summary>
        /// Route Map rule name.
        /// </summary>
        [Input("ruleName", required: true)]
        public string RuleName { get; set; } = null!;

        public GetRouteMapRuleEntryMatchRoutePrefixListArgs()
        {
        }
        public static new GetRouteMapRuleEntryMatchRoutePrefixListArgs Empty => new GetRouteMapRuleEntryMatchRoutePrefixListArgs();
    }

    public sealed class GetRouteMapRuleEntryMatchRoutePrefixListInvokeArgs : global::Pulumi.InvokeArgs
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
        /// DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
        /// </summary>
        [Input("prefixListDn", required: true)]
        public Input<string> PrefixListDn { get; set; } = null!;

        /// <summary>
        /// Route Map rule name.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        public GetRouteMapRuleEntryMatchRoutePrefixListInvokeArgs()
        {
        }
        public static new GetRouteMapRuleEntryMatchRoutePrefixListInvokeArgs Empty => new GetRouteMapRuleEntryMatchRoutePrefixListInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouteMapRuleEntryMatchRoutePrefixListResult
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Route-Map Rule Entry order.
        /// </summary>
        public readonly int Order;
        /// <summary>
        /// DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
        /// </summary>
        public readonly string PrefixListDn;
        /// <summary>
        /// Route Map rule name.
        /// </summary>
        public readonly string RuleName;

        [OutputConstructor]
        private GetRouteMapRuleEntryMatchRoutePrefixListResult(
            string? device,

            string id,

            int order,

            string prefixListDn,

            string ruleName)
        {
            Device = device;
            Id = id;
            Order = order;
            PrefixListDn = prefixListDn;
            RuleName = ruleName;
        }
    }
}
