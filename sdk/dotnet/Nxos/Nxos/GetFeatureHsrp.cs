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
    public static class GetFeatureHsrp
    {
        public static Task<GetFeatureHsrpResult> InvokeAsync(GetFeatureHsrpArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFeatureHsrpResult>("nxos:nxos/getFeatureHsrp:getFeatureHsrp", args ?? new GetFeatureHsrpArgs(), options.WithDefaults());

        public static Output<GetFeatureHsrpResult> Invoke(GetFeatureHsrpInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFeatureHsrpResult>("nxos:nxos/getFeatureHsrp:getFeatureHsrp", args ?? new GetFeatureHsrpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFeatureHsrpArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        public GetFeatureHsrpArgs()
        {
        }
        public static new GetFeatureHsrpArgs Empty => new GetFeatureHsrpArgs();
    }

    public sealed class GetFeatureHsrpInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        public GetFeatureHsrpInvokeArgs()
        {
        }
        public static new GetFeatureHsrpInvokeArgs Empty => new GetFeatureHsrpInvokeArgs();
    }


    [OutputType]
    public sealed class GetFeatureHsrpResult
    {
        public readonly string AdminState;
        public readonly string? Device;
        public readonly string Id;

        [OutputConstructor]
        private GetFeatureHsrpResult(
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