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
    public static class GetFeatureLacp
    {
        public static Task<GetFeatureLacpResult> InvokeAsync(GetFeatureLacpArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFeatureLacpResult>("nxos:nxos/getFeatureLacp:getFeatureLacp", args ?? new GetFeatureLacpArgs(), options.WithDefaults());

        public static Output<GetFeatureLacpResult> Invoke(GetFeatureLacpInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFeatureLacpResult>("nxos:nxos/getFeatureLacp:getFeatureLacp", args ?? new GetFeatureLacpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFeatureLacpArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        public GetFeatureLacpArgs()
        {
        }
        public static new GetFeatureLacpArgs Empty => new GetFeatureLacpArgs();
    }

    public sealed class GetFeatureLacpInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        public GetFeatureLacpInvokeArgs()
        {
        }
        public static new GetFeatureLacpInvokeArgs Empty => new GetFeatureLacpInvokeArgs();
    }


    [OutputType]
    public sealed class GetFeatureLacpResult
    {
        public readonly string AdminState;
        public readonly string? Device;
        public readonly string Id;

        [OutputConstructor]
        private GetFeatureLacpResult(
            string adminState,

            string? device,

            string id)
        {
            AdminState = adminState;
            Device = device;
            Id = id;
        }
    }
}
