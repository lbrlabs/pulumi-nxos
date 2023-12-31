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
    public static class GetHmmInterface
    {
        /// <summary>
        /// This data source can read the HMM Fabric Forwarding mode information related to an Interface.
        /// 
        /// - API Documentation: [hmmFwdIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Host%20Mobility/hmm:FwdIf/)
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
        ///     var example = Nxos.GetHmmInterface.Invoke(new()
        ///     {
        ///         InterfaceId = "vlan10",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetHmmInterfaceResult> InvokeAsync(GetHmmInterfaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetHmmInterfaceResult>("nxos:index/getHmmInterface:getHmmInterface", args ?? new GetHmmInterfaceArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the HMM Fabric Forwarding mode information related to an Interface.
        /// 
        /// - API Documentation: [hmmFwdIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Host%20Mobility/hmm:FwdIf/)
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
        ///     var example = Nxos.GetHmmInterface.Invoke(new()
        ///     {
        ///         InterfaceId = "vlan10",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetHmmInterfaceResult> Invoke(GetHmmInterfaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetHmmInterfaceResult>("nxos:index/getHmmInterface:getHmmInterface", args ?? new GetHmmInterfaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetHmmInterfaceArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `vlan10`.
        /// </summary>
        [Input("interfaceId", required: true)]
        public string InterfaceId { get; set; } = null!;

        public GetHmmInterfaceArgs()
        {
        }
        public static new GetHmmInterfaceArgs Empty => new GetHmmInterfaceArgs();
    }

    public sealed class GetHmmInterfaceInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `vlan10`.
        /// </summary>
        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        public GetHmmInterfaceInvokeArgs()
        {
        }
        public static new GetHmmInterfaceInvokeArgs Empty => new GetHmmInterfaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetHmmInterfaceResult
    {
        /// <summary>
        /// Administrative state.
        /// </summary>
        public readonly string AdminState;
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `vlan10`.
        /// </summary>
        public readonly string InterfaceId;
        /// <summary>
        /// HMM Fabric Forwarding mode information for the interface.
        /// </summary>
        public readonly string Mode;

        [OutputConstructor]
        private GetHmmInterfaceResult(
            string adminState,

            string? device,

            string id,

            string interfaceId,

            string mode)
        {
            AdminState = adminState;
            Device = device;
            Id = id;
            InterfaceId = interfaceId;
            Mode = mode;
        }
    }
}
