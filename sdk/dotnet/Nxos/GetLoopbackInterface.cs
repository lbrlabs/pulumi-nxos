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
    public static class GetLoopbackInterface
    {
        /// <summary>
        /// This data source can read a loopback interface.
        /// 
        /// - API Documentation: [l3LbRtdIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/l3:LbRtdIf/)
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
        ///     var example = Nxos.GetLoopbackInterface.Invoke(new()
        ///     {
        ///         InterfaceId = "lo123",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLoopbackInterfaceResult> InvokeAsync(GetLoopbackInterfaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLoopbackInterfaceResult>("nxos:index/getLoopbackInterface:getLoopbackInterface", args ?? new GetLoopbackInterfaceArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read a loopback interface.
        /// 
        /// - API Documentation: [l3LbRtdIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/l3:LbRtdIf/)
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
        ///     var example = Nxos.GetLoopbackInterface.Invoke(new()
        ///     {
        ///         InterfaceId = "lo123",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetLoopbackInterfaceResult> Invoke(GetLoopbackInterfaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLoopbackInterfaceResult>("nxos:index/getLoopbackInterface:getLoopbackInterface", args ?? new GetLoopbackInterfaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoopbackInterfaceArgs : global::Pulumi.InvokeArgs
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

        public GetLoopbackInterfaceArgs()
        {
        }
        public static new GetLoopbackInterfaceArgs Empty => new GetLoopbackInterfaceArgs();
    }

    public sealed class GetLoopbackInterfaceInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetLoopbackInterfaceInvokeArgs()
        {
        }
        public static new GetLoopbackInterfaceInvokeArgs Empty => new GetLoopbackInterfaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetLoopbackInterfaceResult
    {
        /// <summary>
        /// Administrative state.
        /// </summary>
        public readonly string AdminState;
        /// <summary>
        /// Interface description.
        /// </summary>
        public readonly string Description;
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
