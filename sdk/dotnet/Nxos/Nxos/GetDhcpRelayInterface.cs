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
    public static class GetDhcpRelayInterface
    {
        public static Task<GetDhcpRelayInterfaceResult> InvokeAsync(GetDhcpRelayInterfaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDhcpRelayInterfaceResult>("nxos:nxos/getDhcpRelayInterface:getDhcpRelayInterface", args ?? new GetDhcpRelayInterfaceArgs(), options.WithDefaults());

        public static Output<GetDhcpRelayInterfaceResult> Invoke(GetDhcpRelayInterfaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDhcpRelayInterfaceResult>("nxos:nxos/getDhcpRelayInterface:getDhcpRelayInterface", args ?? new GetDhcpRelayInterfaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDhcpRelayInterfaceArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public string? Device { get; set; }

        [Input("interfaceId", required: true)]
        public string InterfaceId { get; set; } = null!;

        public GetDhcpRelayInterfaceArgs()
        {
        }
        public static new GetDhcpRelayInterfaceArgs Empty => new GetDhcpRelayInterfaceArgs();
    }

    public sealed class GetDhcpRelayInterfaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        public GetDhcpRelayInterfaceInvokeArgs()
        {
        }
        public static new GetDhcpRelayInterfaceInvokeArgs Empty => new GetDhcpRelayInterfaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetDhcpRelayInterfaceResult
    {
        public readonly string? Device;
        public readonly string Id;
        public readonly string InterfaceId;

        [OutputConstructor]
        private GetDhcpRelayInterfaceResult(
            string? device,

            string id,

            string interfaceId)
        {
            Device = device;
            Id = id;
            InterfaceId = interfaceId;
        }
    }
}