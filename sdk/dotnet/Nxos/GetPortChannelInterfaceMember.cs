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
    public static class GetPortChannelInterfaceMember
    {
        /// <summary>
        /// This data source can read the configuration of a port-channel interface member.
        /// 
        /// - API Documentation: [pcRsMbrIfs](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Interfaces/pc:RsMbrIfs/)
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
        ///     var example = Nxos.GetPortChannelInterfaceMember.Invoke(new()
        ///     {
        ///         InterfaceDn = "sys/intf/phys-[eth1/11]",
        ///         InterfaceId = "po1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPortChannelInterfaceMemberResult> InvokeAsync(GetPortChannelInterfaceMemberArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPortChannelInterfaceMemberResult>("nxos:index/getPortChannelInterfaceMember:getPortChannelInterfaceMember", args ?? new GetPortChannelInterfaceMemberArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the configuration of a port-channel interface member.
        /// 
        /// - API Documentation: [pcRsMbrIfs](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Interfaces/pc:RsMbrIfs/)
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
        ///     var example = Nxos.GetPortChannelInterfaceMember.Invoke(new()
        ///     {
        ///         InterfaceDn = "sys/intf/phys-[eth1/11]",
        ///         InterfaceId = "po1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPortChannelInterfaceMemberResult> Invoke(GetPortChannelInterfaceMemberInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPortChannelInterfaceMemberResult>("nxos:index/getPortChannelInterfaceMember:getPortChannelInterfaceMember", args ?? new GetPortChannelInterfaceMemberInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPortChannelInterfaceMemberArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// DN of interface. For example: `sys/intf/phys-[eth1/1]`.
        /// </summary>
        [Input("interfaceDn", required: true)]
        public string InterfaceDn { get; set; } = null!;

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `po1`.
        /// </summary>
        [Input("interfaceId", required: true)]
        public string InterfaceId { get; set; } = null!;

        public GetPortChannelInterfaceMemberArgs()
        {
        }
        public static new GetPortChannelInterfaceMemberArgs Empty => new GetPortChannelInterfaceMemberArgs();
    }

    public sealed class GetPortChannelInterfaceMemberInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// DN of interface. For example: `sys/intf/phys-[eth1/1]`.
        /// </summary>
        [Input("interfaceDn", required: true)]
        public Input<string> InterfaceDn { get; set; } = null!;

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `po1`.
        /// </summary>
        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        public GetPortChannelInterfaceMemberInvokeArgs()
        {
        }
        public static new GetPortChannelInterfaceMemberInvokeArgs Empty => new GetPortChannelInterfaceMemberInvokeArgs();
    }


    [OutputType]
    public sealed class GetPortChannelInterfaceMemberResult
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
        /// DN of interface. For example: `sys/intf/phys-[eth1/1]`.
        /// </summary>
        public readonly string InterfaceDn;
        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `po1`.
        /// </summary>
        public readonly string InterfaceId;

        [OutputConstructor]
        private GetPortChannelInterfaceMemberResult(
            string? device,

            string id,

            string interfaceDn,

            string interfaceId)
        {
            Device = device;
            Id = id;
            InterfaceDn = interfaceDn;
            InterfaceId = interfaceId;
        }
    }
}
