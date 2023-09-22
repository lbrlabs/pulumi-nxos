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
    public static class GetSviInterfaceVrf
    {
        /// <summary>
        /// This data source can read an SVI interface VRF association.
        /// 
        /// - API Documentation: [nwRtVrfMbr](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/nw:RtVrfMbr/)
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
        ///     var example = Nxos.GetSviInterfaceVrf.Invoke(new()
        ///     {
        ///         InterfaceId = "vlan293",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSviInterfaceVrfResult> InvokeAsync(GetSviInterfaceVrfArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSviInterfaceVrfResult>("nxos:index/getSviInterfaceVrf:getSviInterfaceVrf", args ?? new GetSviInterfaceVrfArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read an SVI interface VRF association.
        /// 
        /// - API Documentation: [nwRtVrfMbr](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/nw:RtVrfMbr/)
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
        ///     var example = Nxos.GetSviInterfaceVrf.Invoke(new()
        ///     {
        ///         InterfaceId = "vlan293",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSviInterfaceVrfResult> Invoke(GetSviInterfaceVrfInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSviInterfaceVrfResult>("nxos:index/getSviInterfaceVrf:getSviInterfaceVrf", args ?? new GetSviInterfaceVrfInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSviInterfaceVrfArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `vlan100`.
        /// </summary>
        [Input("interfaceId", required: true)]
        public string InterfaceId { get; set; } = null!;

        public GetSviInterfaceVrfArgs()
        {
        }
        public static new GetSviInterfaceVrfArgs Empty => new GetSviInterfaceVrfArgs();
    }

    public sealed class GetSviInterfaceVrfInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `vlan100`.
        /// </summary>
        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        public GetSviInterfaceVrfInvokeArgs()
        {
        }
        public static new GetSviInterfaceVrfInvokeArgs Empty => new GetSviInterfaceVrfInvokeArgs();
    }


    [OutputType]
    public sealed class GetSviInterfaceVrfResult
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
        /// Must match first field in the output of `show intf brief`. Example: `vlan100`.
        /// </summary>
        public readonly string InterfaceId;
        /// <summary>
        /// DN of VRF. For example: `sys/inst-VRF1`.
        /// </summary>
        public readonly string VrfDn;

        [OutputConstructor]
        private GetSviInterfaceVrfResult(
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
