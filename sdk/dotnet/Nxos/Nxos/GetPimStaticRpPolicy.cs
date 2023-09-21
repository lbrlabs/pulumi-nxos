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
    public static class GetPimStaticRpPolicy
    {
        public static Task<GetPimStaticRpPolicyResult> InvokeAsync(GetPimStaticRpPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPimStaticRpPolicyResult>("nxos:nxos/getPimStaticRpPolicy:getPimStaticRpPolicy", args ?? new GetPimStaticRpPolicyArgs(), options.WithDefaults());

        public static Output<GetPimStaticRpPolicyResult> Invoke(GetPimStaticRpPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPimStaticRpPolicyResult>("nxos:nxos/getPimStaticRpPolicy:getPimStaticRpPolicy", args ?? new GetPimStaticRpPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPimStaticRpPolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        [Input("vrfName", required: true)]
        public string VrfName { get; set; } = null!;

        public GetPimStaticRpPolicyArgs()
        {
        }
        public static new GetPimStaticRpPolicyArgs Empty => new GetPimStaticRpPolicyArgs();
    }

    public sealed class GetPimStaticRpPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("vrfName", required: true)]
        public Input<string> VrfName { get; set; } = null!;

        public GetPimStaticRpPolicyInvokeArgs()
        {
        }
        public static new GetPimStaticRpPolicyInvokeArgs Empty => new GetPimStaticRpPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetPimStaticRpPolicyResult
    {
        public readonly string? Device;
        public readonly string Id;
        public readonly string Name;
        public readonly string VrfName;

        [OutputConstructor]
        private GetPimStaticRpPolicyResult(
            string? device,

            string id,

            string name,

            string vrfName)
        {
            Device = device;
            Id = id;
            Name = name;
            VrfName = vrfName;
        }
    }
}
