// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Nxos
{
    public static class GetDhcpRelayInterface
    {
        /// <summary>
        /// This data source can read a DHCP relay interface.
        /// 
        /// - API Documentation: [dhcpRelayIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/DHCP/dhcp:RelayIf/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Nxos = Pulumi.Nxos;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Nxos.GetDhcpRelayInterface.Invoke(new()
        ///     {
        ///         InterfaceId = "eth1/10",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDhcpRelayInterfaceResult> InvokeAsync(GetDhcpRelayInterfaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDhcpRelayInterfaceResult>("nxos:index/getDhcpRelayInterface:getDhcpRelayInterface", args ?? new GetDhcpRelayInterfaceArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read a DHCP relay interface.
        /// 
        /// - API Documentation: [dhcpRelayIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/DHCP/dhcp:RelayIf/)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Nxos = Pulumi.Nxos;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Nxos.GetDhcpRelayInterface.Invoke(new()
        ///     {
        ///         InterfaceId = "eth1/10",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDhcpRelayInterfaceResult> Invoke(GetDhcpRelayInterfaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDhcpRelayInterfaceResult>("nxos:index/getDhcpRelayInterface:getDhcpRelayInterface", args ?? new GetDhcpRelayInterfaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDhcpRelayInterfaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        [Input("interfaceId", required: true)]
        public string InterfaceId { get; set; } = null!;

        public GetDhcpRelayInterfaceArgs()
        {
        }
        public static new GetDhcpRelayInterfaceArgs Empty => new GetDhcpRelayInterfaceArgs();
    }

    public sealed class GetDhcpRelayInterfaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
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
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
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