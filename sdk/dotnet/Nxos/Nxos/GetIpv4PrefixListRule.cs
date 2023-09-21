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
    public static class GetIpv4PrefixListRule
    {
        public static Task<GetIpv4PrefixListRuleResult> InvokeAsync(GetIpv4PrefixListRuleArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpv4PrefixListRuleResult>("nxos:nxos/getIpv4PrefixListRule:getIpv4PrefixListRule", args ?? new GetIpv4PrefixListRuleArgs(), options.WithDefaults());

        public static Output<GetIpv4PrefixListRuleResult> Invoke(GetIpv4PrefixListRuleInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpv4PrefixListRuleResult>("nxos:nxos/getIpv4PrefixListRule:getIpv4PrefixListRule", args ?? new GetIpv4PrefixListRuleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpv4PrefixListRuleArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetIpv4PrefixListRuleArgs()
        {
        }
        public static new GetIpv4PrefixListRuleArgs Empty => new GetIpv4PrefixListRuleArgs();
    }

    public sealed class GetIpv4PrefixListRuleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetIpv4PrefixListRuleInvokeArgs()
        {
        }
        public static new GetIpv4PrefixListRuleInvokeArgs Empty => new GetIpv4PrefixListRuleInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpv4PrefixListRuleResult
    {
        public readonly string? Device;
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetIpv4PrefixListRuleResult(
            string? device,

            string id,

            string name)
        {
            Device = device;
            Id = id;
            Name = name;
        }
    }
}
