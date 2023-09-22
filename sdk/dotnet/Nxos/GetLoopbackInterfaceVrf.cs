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
    public static class GetLoopbackInterfaceVrf
    {
        /// <summary>
        /// This data source can read a loopback interface VRF association.
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
        ///     var example = Nxos.GetLoopbackInterfaceVrf.Invoke(new()
        ///     {
        ///         InterfaceId = "lo123",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLoopbackInterfaceVrfResult> InvokeAsync(GetLoopbackInterfaceVrfArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLoopbackInterfaceVrfResult>("nxos:index/getLoopbackInterfaceVrf:getLoopbackInterfaceVrf", args ?? new GetLoopbackInterfaceVrfArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read a loopback interface VRF association.
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
        ///     var example = Nxos.GetLoopbackInterfaceVrf.Invoke(new()
        ///     {
        ///         InterfaceId = "lo123",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetLoopbackInterfaceVrfResult> Invoke(GetLoopbackInterfaceVrfInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoopbackInterfaceVrfResult>("nxos:index/getLoopbackInterfaceVrf:getLoopbackInterfaceVrf", args ?? new GetLoopbackInterfaceVrfInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoopbackInterfaceVrfArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `lo123`.
        /// </summary>
        [Input("interfaceId", required: true)]
        public string InterfaceId { get; set; } = null!;

        public GetLoopbackInterfaceVrfArgs()
        {
        }
        public static new GetLoopbackInterfaceVrfArgs Empty => new GetLoopbackInterfaceVrfArgs();
    }

    public sealed class GetLoopbackInterfaceVrfInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `lo123`.
        /// </summary>
        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        public GetLoopbackInterfaceVrfInvokeArgs()
        {
        }
        public static new GetLoopbackInterfaceVrfInvokeArgs Empty => new GetLoopbackInterfaceVrfInvokeArgs();
    }


    [OutputType]
    public sealed class GetLoopbackInterfaceVrfResult
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
        /// Must match first field in the output of `show intf brief`. Example: `lo123`.
        /// </summary>
        public readonly string InterfaceId;
        /// <summary>
        /// DN of VRF. For example: `sys/inst-VRF1`.
        /// </summary>
        public readonly string VrfDn;

        [OutputConstructor]
        private GetLoopbackInterfaceVrfResult(
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
