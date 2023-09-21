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
    public static class GetFeaturePtp
    {
        public static Task<GetFeaturePtpResult> InvokeAsync(GetFeaturePtpArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFeaturePtpResult>("nxos:nxos/getFeaturePtp:getFeaturePtp", args ?? new GetFeaturePtpArgs(), options.WithDefaults());

        public static Output<GetFeaturePtpResult> Invoke(GetFeaturePtpInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFeaturePtpResult>("nxos:nxos/getFeaturePtp:getFeaturePtp", args ?? new GetFeaturePtpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFeaturePtpArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        public GetFeaturePtpArgs()
        {
        }
        public static new GetFeaturePtpArgs Empty => new GetFeaturePtpArgs();
    }

    public sealed class GetFeaturePtpInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        public GetFeaturePtpInvokeArgs()
        {
        }
        public static new GetFeaturePtpInvokeArgs Empty => new GetFeaturePtpInvokeArgs();
    }


    [OutputType]
    public sealed class GetFeaturePtpResult
    {
        public readonly string AdminState;
        public readonly string? Device;
        public readonly string Id;

        [OutputConstructor]
        private GetFeaturePtpResult(
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
