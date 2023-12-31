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
    public static class GetPimAnycastRpPeer
    {
        /// <summary>
        /// This data source can read the PIM Anycast RP peer configuration.
        /// 
        /// - API Documentation: [pimAcastRPPeer](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:AcastRPPeer/)
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
        ///     var example = Nxos.GetPimAnycastRpPeer.Invoke(new()
        ///     {
        ///         Address = "10.1.1.1/32",
        ///         RpSetAddress = "20.1.1.1/32",
        ///         VrfName = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPimAnycastRpPeerResult> InvokeAsync(GetPimAnycastRpPeerArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPimAnycastRpPeerResult>("nxos:index/getPimAnycastRpPeer:getPimAnycastRpPeer", args ?? new GetPimAnycastRpPeerArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the PIM Anycast RP peer configuration.
        /// 
        /// - API Documentation: [pimAcastRPPeer](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:AcastRPPeer/)
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
        ///     var example = Nxos.GetPimAnycastRpPeer.Invoke(new()
        ///     {
        ///         Address = "10.1.1.1/32",
        ///         RpSetAddress = "20.1.1.1/32",
        ///         VrfName = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPimAnycastRpPeerResult> Invoke(GetPimAnycastRpPeerInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPimAnycastRpPeerResult>("nxos:index/getPimAnycastRpPeer:getPimAnycastRpPeer", args ?? new GetPimAnycastRpPeerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPimAnycastRpPeerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Anycast RP address.
        /// </summary>
        [Input("address", required: true)]
        public string Address { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// RP set address.
        /// </summary>
        [Input("rpSetAddress", required: true)]
        public string RpSetAddress { get; set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrfName", required: true)]
        public string VrfName { get; set; } = null!;

        public GetPimAnycastRpPeerArgs()
        {
        }
        public static new GetPimAnycastRpPeerArgs Empty => new GetPimAnycastRpPeerArgs();
    }

    public sealed class GetPimAnycastRpPeerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Anycast RP address.
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// RP set address.
        /// </summary>
        [Input("rpSetAddress", required: true)]
        public Input<string> RpSetAddress { get; set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrfName", required: true)]
        public Input<string> VrfName { get; set; } = null!;

        public GetPimAnycastRpPeerInvokeArgs()
        {
        }
        public static new GetPimAnycastRpPeerInvokeArgs Empty => new GetPimAnycastRpPeerInvokeArgs();
    }


    [OutputType]
    public sealed class GetPimAnycastRpPeerResult
    {
        /// <summary>
        /// Anycast RP address.
        /// </summary>
        public readonly string Address;
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// RP set address.
        /// </summary>
        public readonly string RpSetAddress;
        /// <summary>
        /// VRF name.
        /// </summary>
        public readonly string VrfName;

        [OutputConstructor]
        private GetPimAnycastRpPeerResult(
            string address,

            string? device,

            string id,

            string rpSetAddress,

            string vrfName)
        {
            Address = address;
            Device = device;
            Id = id;
            RpSetAddress = rpSetAddress;
            VrfName = vrfName;
        }
    }
}
