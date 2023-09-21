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
    public static class GetIsisVrf
    {
        public static Task<GetIsisVrfResult> InvokeAsync(GetIsisVrfArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIsisVrfResult>("nxos:nxos/getIsisVrf:getIsisVrf", args ?? new GetIsisVrfArgs(), options.WithDefaults());

        public static Output<GetIsisVrfResult> Invoke(GetIsisVrfInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIsisVrfResult>("nxos:nxos/getIsisVrf:getIsisVrf", args ?? new GetIsisVrfInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIsisVrfArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        [Input("instanceName", required: true)]
        public string InstanceName { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetIsisVrfArgs()
        {
        }
        public static new GetIsisVrfArgs Empty => new GetIsisVrfArgs();
    }

    public sealed class GetIsisVrfInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetIsisVrfInvokeArgs()
        {
        }
        public static new GetIsisVrfInvokeArgs Empty => new GetIsisVrfInvokeArgs();
    }


    [OutputType]
    public sealed class GetIsisVrfResult
    {
        public readonly string AdminState;
        public readonly bool AuthenticationCheckL1;
        public readonly bool AuthenticationCheckL2;
        public readonly string AuthenticationKeyL1;
        public readonly string AuthenticationKeyL2;
        public readonly string AuthenticationTypeL1;
        public readonly string AuthenticationTypeL2;
        public readonly int BandwidthReference;
        public readonly string BanwidthReferenceUnit;
        public readonly string? Device;
        public readonly string Id;
        public readonly string InstanceName;
        public readonly string IsType;
        public readonly string MetricType;
        public readonly int Mtu;
        public readonly string Name;
        public readonly string Net;
        public readonly string PassiveDefault;

        [OutputConstructor]
        private GetIsisVrfResult(
            string adminState,

            bool authenticationCheckL1,

            bool authenticationCheckL2,

            string authenticationKeyL1,

            string authenticationKeyL2,

            string authenticationTypeL1,

            string authenticationTypeL2,

            int bandwidthReference,

            string banwidthReferenceUnit,

            string? device,

            string id,

            string instanceName,

            string isType,

            string metricType,

            int mtu,

            string name,

            string net,

            string passiveDefault)
        {
            AdminState = adminState;
            AuthenticationCheckL1 = authenticationCheckL1;
            AuthenticationCheckL2 = authenticationCheckL2;
            AuthenticationKeyL1 = authenticationKeyL1;
            AuthenticationKeyL2 = authenticationKeyL2;
            AuthenticationTypeL1 = authenticationTypeL1;
            AuthenticationTypeL2 = authenticationTypeL2;
            BandwidthReference = bandwidthReference;
            BanwidthReferenceUnit = banwidthReferenceUnit;
            Device = device;
            Id = id;
            InstanceName = instanceName;
            IsType = isType;
            MetricType = metricType;
            Mtu = mtu;
            Name = name;
            Net = net;
            PassiveDefault = passiveDefault;
        }
    }
}
