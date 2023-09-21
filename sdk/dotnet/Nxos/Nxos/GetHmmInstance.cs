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
    public static class GetHmmInstance
    {
        public static Task<GetHmmInstanceResult> InvokeAsync(GetHmmInstanceArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHmmInstanceResult>("nxos:nxos/getHmmInstance:getHmmInstance", args ?? new GetHmmInstanceArgs(), options.WithDefaults());

        public static Output<GetHmmInstanceResult> Invoke(GetHmmInstanceInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHmmInstanceResult>("nxos:nxos/getHmmInstance:getHmmInstance", args ?? new GetHmmInstanceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHmmInstanceArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        public GetHmmInstanceArgs()
        {
        }
        public static new GetHmmInstanceArgs Empty => new GetHmmInstanceArgs();
    }

    public sealed class GetHmmInstanceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        public GetHmmInstanceInvokeArgs()
        {
        }
        public static new GetHmmInstanceInvokeArgs Empty => new GetHmmInstanceInvokeArgs();
    }


    [OutputType]
    public sealed class GetHmmInstanceResult
    {
        public readonly string AdminState;
        public readonly string AnycastMac;
        public readonly string? Device;
        public readonly string Id;

        [OutputConstructor]
        private GetHmmInstanceResult(
            string adminState,

            string anycastMac,

            string? device,

            string id)
        {
            AdminState = adminState;
            AnycastMac = anycastMac;
            Device = device;
            Id = id;
        }
    }
}
