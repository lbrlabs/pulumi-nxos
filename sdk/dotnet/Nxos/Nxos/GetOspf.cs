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
    public static class GetOspf
    {
        public static Task<GetOspfResult> InvokeAsync(GetOspfArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOspfResult>("nxos:nxos/getOspf:getOspf", args ?? new GetOspfArgs(), options.WithDefaults());

        public static Output<GetOspfResult> Invoke(GetOspfInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOspfResult>("nxos:nxos/getOspf:getOspf", args ?? new GetOspfInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOspfArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        public GetOspfArgs()
        {
        }
        public static new GetOspfArgs Empty => new GetOspfArgs();
    }

    public sealed class GetOspfInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        public GetOspfInvokeArgs()
        {
        }
        public static new GetOspfInvokeArgs Empty => new GetOspfInvokeArgs();
    }


    [OutputType]
    public sealed class GetOspfResult
    {
        public readonly string AdminState;
        public readonly string? Device;
        public readonly string Id;

        [OutputConstructor]
        private GetOspfResult(
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
