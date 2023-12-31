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
    public static class GetRouteMapRuleEntry
    {
        /// <summary>
        /// This data source can read a Route-Map Rule Entry configuration.
        /// 
        /// - API Documentation: [rtmapEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:Entry/)
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
        ///     var example = Nxos.GetRouteMapRuleEntry.Invoke(new()
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
        public static Task<GetRouteMapRuleEntryResult> InvokeAsync(GetRouteMapRuleEntryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRouteMapRuleEntryResult>("nxos:index/getRouteMapRuleEntry:getRouteMapRuleEntry", args ?? new GetRouteMapRuleEntryArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read a Route-Map Rule Entry configuration.
        /// 
        /// - API Documentation: [rtmapEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:Entry/)
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
        ///     var example = Nxos.GetRouteMapRuleEntry.Invoke(new()
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
        public static Output<GetRouteMapRuleEntryResult> Invoke(GetRouteMapRuleEntryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRouteMapRuleEntryResult>("nxos:index/getRouteMapRuleEntry:getRouteMapRuleEntry", args ?? new GetRouteMapRuleEntryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRouteMapRuleEntryArgs : global::Pulumi.InvokeArgs
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

        public GetRouteMapRuleEntryArgs()
        {
        }
        public static new GetRouteMapRuleEntryArgs Empty => new GetRouteMapRuleEntryArgs();
    }

    public sealed class GetRouteMapRuleEntryInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetRouteMapRuleEntryInvokeArgs()
        {
        }
        public static new GetRouteMapRuleEntryInvokeArgs Empty => new GetRouteMapRuleEntryInvokeArgs();
    }


    [OutputType]
    public sealed class GetRouteMapRuleEntryResult
    {
        /// <summary>
        /// Route-Map Rule Entry action.
        /// </summary>
        public readonly string Action;
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
        /// Route Map rule name.
        /// </summary>
        public readonly string RuleName;

        [OutputConstructor]
        private GetRouteMapRuleEntryResult(
            string action,

            string? device,

            string id,

            int order,

            string ruleName)
        {
            Action = action;
            Device = device;
            Id = id;
            Order = order;
            RuleName = ruleName;
        }
    }
}
