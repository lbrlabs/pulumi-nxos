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
    public static class GetPimInterface
    {
        /// <summary>
        /// This data source can read the PIM interface configuration.
        /// 
        /// - API Documentation: [pimIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:If/)
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
        ///     var example = Nxos.GetPimInterface.Invoke(new()
        ///     {
        ///         InterfaceId = "eth1/10",
        ///         VrfName = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPimInterfaceResult> InvokeAsync(GetPimInterfaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPimInterfaceResult>("nxos:index/getPimInterface:getPimInterface", args ?? new GetPimInterfaceArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the PIM interface configuration.
        /// 
        /// - API Documentation: [pimIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:If/)
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
        ///     var example = Nxos.GetPimInterface.Invoke(new()
        ///     {
        ///         InterfaceId = "eth1/10",
        ///         VrfName = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPimInterfaceResult> Invoke(GetPimInterfaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPimInterfaceResult>("nxos:index/getPimInterface:getPimInterface", args ?? new GetPimInterfaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPimInterfaceArgs : global::Pulumi.InvokeArgs
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

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrfName", required: true)]
        public string VrfName { get; set; } = null!;

        public GetPimInterfaceArgs()
        {
        }
        public static new GetPimInterfaceArgs Empty => new GetPimInterfaceArgs();
    }

    public sealed class GetPimInterfaceInvokeArgs : global::Pulumi.InvokeArgs
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

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrfName", required: true)]
        public Input<string> VrfName { get; set; } = null!;

        public GetPimInterfaceInvokeArgs()
        {
        }
        public static new GetPimInterfaceInvokeArgs Empty => new GetPimInterfaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetPimInterfaceResult
    {
        /// <summary>
        /// Administrative state.
        /// </summary>
        public readonly string AdminState;
        /// <summary>
        /// BFD.
        /// </summary>
        public readonly string Bfd;
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// Designated Router priority level.
        /// </summary>
        public readonly int DrPriority;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        public readonly string InterfaceId;
        /// <summary>
        /// Passive interface.
        /// </summary>
        public readonly bool Passive;
        /// <summary>
        /// Sparse mode.
        /// </summary>
        public readonly bool SparseMode;
        /// <summary>
        /// VRF name.
        /// </summary>
        public readonly string VrfName;

        [OutputConstructor]
        private GetPimInterfaceResult(
            string adminState,

            string bfd,

            string? device,

            int drPriority,

            string id,

            string interfaceId,

            bool passive,

            bool sparseMode,

            string vrfName)
        {
            AdminState = adminState;
            Bfd = bfd;
            Device = device;
            DrPriority = drPriority;
            Id = id;
            InterfaceId = interfaceId;
            Passive = passive;
            SparseMode = sparseMode;
            VrfName = vrfName;
        }
    }
}