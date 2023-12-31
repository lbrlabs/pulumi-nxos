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
    public static class GetIpv4PrefixListRuleEntry
    {
        /// <summary>
        /// This data source can read an IPv4 Prefix List entry configuration.
        /// 
        /// - API Documentation: [rtpfxEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtpfx:Entry/)
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
        ///     var example = Nxos.GetIpv4PrefixListRuleEntry.Invoke(new()
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
        public static Task<GetIpv4PrefixListRuleEntryResult> InvokeAsync(GetIpv4PrefixListRuleEntryArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpv4PrefixListRuleEntryResult>("nxos:index/getIpv4PrefixListRuleEntry:getIpv4PrefixListRuleEntry", args ?? new GetIpv4PrefixListRuleEntryArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read an IPv4 Prefix List entry configuration.
        /// 
        /// - API Documentation: [rtpfxEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtpfx:Entry/)
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
        ///     var example = Nxos.GetIpv4PrefixListRuleEntry.Invoke(new()
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
        public static Output<GetIpv4PrefixListRuleEntryResult> Invoke(GetIpv4PrefixListRuleEntryInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpv4PrefixListRuleEntryResult>("nxos:index/getIpv4PrefixListRuleEntry:getIpv4PrefixListRuleEntry", args ?? new GetIpv4PrefixListRuleEntryInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpv4PrefixListRuleEntryArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// IPv4 Prefix List Rule Entry order.
        /// </summary>
        [Input("order", required: true)]
        public int Order { get; set; }

        /// <summary>
        /// IPv4 Prefix List Rule name.
        /// </summary>
        [Input("ruleName", required: true)]
        public string RuleName { get; set; } = null!;

        public GetIpv4PrefixListRuleEntryArgs()
        {
        }
        public static new GetIpv4PrefixListRuleEntryArgs Empty => new GetIpv4PrefixListRuleEntryArgs();
    }

    public sealed class GetIpv4PrefixListRuleEntryInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// IPv4 Prefix List Rule Entry order.
        /// </summary>
        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        /// <summary>
        /// IPv4 Prefix List Rule name.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        public GetIpv4PrefixListRuleEntryInvokeArgs()
        {
        }
        public static new GetIpv4PrefixListRuleEntryInvokeArgs Empty => new GetIpv4PrefixListRuleEntryInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpv4PrefixListRuleEntryResult
    {
        /// <summary>
        /// IPv4 Prefix List Rule Entry action.
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// IPv4 Prefix List Rule Entry criteria.
        /// </summary>
        public readonly string Criteria;
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// IPv4 Prefix List Rule Entry start range.
        /// </summary>
        public readonly int FromRange;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// IPv4 Prefix List Rule Entry order.
        /// </summary>
        public readonly int Order;
        /// <summary>
        /// IPv4 Prefix List Rule Entry prefix.
        /// </summary>
        public readonly string Prefix;
        /// <summary>
        /// IPv4 Prefix List Rule name.
        /// </summary>
        public readonly string RuleName;
        /// <summary>
        /// IPv4 Prefix List Rule Entry end range.
        /// </summary>
        public readonly int ToRange;

        [OutputConstructor]
        private GetIpv4PrefixListRuleEntryResult(
            string action,

            string criteria,

            string? device,

            int fromRange,

            string id,

            int order,

            string prefix,

            string ruleName,

            int toRange)
        {
            Action = action;
            Criteria = criteria;
            Device = device;
            FromRange = fromRange;
            Id = id;
            Order = order;
            Prefix = prefix;
            RuleName = ruleName;
            ToRange = toRange;
        }
    }
}
