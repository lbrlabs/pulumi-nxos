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
    public static class GetLoopbackInterface
    {
        public static Task<GetLoopbackInterfaceResult> InvokeAsync(GetLoopbackInterfaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLoopbackInterfaceResult>("nxos:nxos/getLoopbackInterface:getLoopbackInterface", args ?? new GetLoopbackInterfaceArgs(), options.WithDefaults());

        public static Output<GetLoopbackInterfaceResult> Invoke(GetLoopbackInterfaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoopbackInterfaceResult>("nxos:nxos/getLoopbackInterface:getLoopbackInterface", args ?? new GetLoopbackInterfaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoopbackInterfaceArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        [Input("interfaceId", required: true)]
        public string InterfaceId { get; set; } = null!;

        public GetLoopbackInterfaceArgs()
        {
        }
        public static new GetLoopbackInterfaceArgs Empty => new GetLoopbackInterfaceArgs();
    }

    public sealed class GetLoopbackInterfaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        public GetLoopbackInterfaceInvokeArgs()
        {
        }
        public static new GetLoopbackInterfaceInvokeArgs Empty => new GetLoopbackInterfaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetLoopbackInterfaceResult
    {
        public readonly string AdminState;
        public readonly string Description;
        public readonly string? Device;
        public readonly string Id;
        public readonly string InterfaceId;

        [OutputConstructor]
        private GetLoopbackInterfaceResult(
            string adminState,

            string description,

            string? device,

            string id,

            string interfaceId)
        {
            AdminState = adminState;
            Description = description;
            Device = device;
            Id = id;
            InterfaceId = interfaceId;
        }
    }
}
