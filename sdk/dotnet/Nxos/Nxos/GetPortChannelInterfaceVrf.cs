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
    public static class GetPortChannelInterfaceVrf
    {
        public static Task<GetPortChannelInterfaceVrfResult> InvokeAsync(GetPortChannelInterfaceVrfArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPortChannelInterfaceVrfResult>("nxos:nxos/getPortChannelInterfaceVrf:getPortChannelInterfaceVrf", args ?? new GetPortChannelInterfaceVrfArgs(), options.WithDefaults());

        public static Output<GetPortChannelInterfaceVrfResult> Invoke(GetPortChannelInterfaceVrfInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPortChannelInterfaceVrfResult>("nxos:nxos/getPortChannelInterfaceVrf:getPortChannelInterfaceVrf", args ?? new GetPortChannelInterfaceVrfInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPortChannelInterfaceVrfArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        [Input("interfaceId", required: true)]
        public string InterfaceId { get; set; } = null!;

        public GetPortChannelInterfaceVrfArgs()
        {
        }
        public static new GetPortChannelInterfaceVrfArgs Empty => new GetPortChannelInterfaceVrfArgs();
    }

    public sealed class GetPortChannelInterfaceVrfInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        public GetPortChannelInterfaceVrfInvokeArgs()
        {
        }
        public static new GetPortChannelInterfaceVrfInvokeArgs Empty => new GetPortChannelInterfaceVrfInvokeArgs();
    }


    [OutputType]
    public sealed class GetPortChannelInterfaceVrfResult
    {
        public readonly string? Device;
        public readonly string Id;
        public readonly string InterfaceId;
        public readonly string VrfDn;

        [OutputConstructor]
        private GetPortChannelInterfaceVrfResult(
            string? device,

            string id,

            string interfaceId,

            string vrfDn)
        {
            Device = device;
            Id = id;
            InterfaceId = interfaceId;
            VrfDn = vrfDn;
        }
    }
}
