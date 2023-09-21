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
    public static class GetPhysicalInterfaceVrf
    {
        public static Task<GetPhysicalInterfaceVrfResult> InvokeAsync(GetPhysicalInterfaceVrfArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPhysicalInterfaceVrfResult>("nxos:nxos/getPhysicalInterfaceVrf:getPhysicalInterfaceVrf", args ?? new GetPhysicalInterfaceVrfArgs(), options.WithDefaults());

        public static Output<GetPhysicalInterfaceVrfResult> Invoke(GetPhysicalInterfaceVrfInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPhysicalInterfaceVrfResult>("nxos:nxos/getPhysicalInterfaceVrf:getPhysicalInterfaceVrf", args ?? new GetPhysicalInterfaceVrfInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPhysicalInterfaceVrfArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        [Input("interfaceId", required: true)]
        public string InterfaceId { get; set; } = null!;

        public GetPhysicalInterfaceVrfArgs()
        {
        }
        public static new GetPhysicalInterfaceVrfArgs Empty => new GetPhysicalInterfaceVrfArgs();
    }

    public sealed class GetPhysicalInterfaceVrfInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        public GetPhysicalInterfaceVrfInvokeArgs()
        {
        }
        public static new GetPhysicalInterfaceVrfInvokeArgs Empty => new GetPhysicalInterfaceVrfInvokeArgs();
    }


    [OutputType]
    public sealed class GetPhysicalInterfaceVrfResult
    {
        public readonly string? Device;
        public readonly string Id;
        public readonly string InterfaceId;
        public readonly string VrfDn;

        [OutputConstructor]
        private GetPhysicalInterfaceVrfResult(
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