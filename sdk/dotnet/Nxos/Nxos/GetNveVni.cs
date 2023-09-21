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
    public static class GetNveVni
    {
        public static Task<GetNveVniResult> InvokeAsync(GetNveVniArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNveVniResult>("nxos:nxos/getNveVni:getNveVni", args ?? new GetNveVniArgs(), options.WithDefaults());

        public static Output<GetNveVniResult> Invoke(GetNveVniInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNveVniResult>("nxos:nxos/getNveVni:getNveVni", args ?? new GetNveVniInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNveVniArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        [Input("vni", required: true)]
        public int Vni { get; set; }

        public GetNveVniArgs()
        {
        }
        public static new GetNveVniArgs Empty => new GetNveVniArgs();
    }

    public sealed class GetNveVniInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("vni", required: true)]
        public Input<int> Vni { get; set; } = null!;

        public GetNveVniInvokeArgs()
        {
        }
        public static new GetNveVniInvokeArgs Empty => new GetNveVniInvokeArgs();
    }


    [OutputType]
    public sealed class GetNveVniResult
    {
        public readonly bool AssociateVrf;
        public readonly string? Device;
        public readonly string Id;
        public readonly string MulticastGroup;
        public readonly string MultisiteIngressReplication;
        public readonly string SuppressArp;
        public readonly int Vni;

        [OutputConstructor]
        private GetNveVniResult(
            bool associateVrf,

            string? device,

            string id,

            string multicastGroup,

            string multisiteIngressReplication,

            string suppressArp,

            int vni)
        {
            AssociateVrf = associateVrf;
            Device = device;
            Id = id;
            MulticastGroup = multicastGroup;
            MultisiteIngressReplication = multisiteIngressReplication;
            SuppressArp = suppressArp;
            Vni = vni;
        }
    }
}