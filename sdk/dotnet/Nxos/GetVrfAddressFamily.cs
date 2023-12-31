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
    public static class GetVrfAddressFamily
    {
        /// <summary>
        /// This data source can read a VRF Address Family.
        /// 
        /// - API Documentation: [rtctrlDomAf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:DomAf/)
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
        ///     var example = Nxos.GetVrfAddressFamily.Invoke(new()
        ///     {
        ///         AddressFamily = "ipv4-ucast",
        ///         Vrf = "VRF1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVrfAddressFamilyResult> InvokeAsync(GetVrfAddressFamilyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVrfAddressFamilyResult>("nxos:index/getVrfAddressFamily:getVrfAddressFamily", args ?? new GetVrfAddressFamilyArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read a VRF Address Family.
        /// 
        /// - API Documentation: [rtctrlDomAf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:DomAf/)
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
        ///     var example = Nxos.GetVrfAddressFamily.Invoke(new()
        ///     {
        ///         AddressFamily = "ipv4-ucast",
        ///         Vrf = "VRF1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVrfAddressFamilyResult> Invoke(GetVrfAddressFamilyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVrfAddressFamilyResult>("nxos:index/getVrfAddressFamily:getVrfAddressFamily", args ?? new GetVrfAddressFamilyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVrfAddressFamilyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Address family.
        /// </summary>
        [Input("addressFamily", required: true)]
        public string AddressFamily { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf", required: true)]
        public string Vrf { get; set; } = null!;

        public GetVrfAddressFamilyArgs()
        {
        }
        public static new GetVrfAddressFamilyArgs Empty => new GetVrfAddressFamilyArgs();
    }

    public sealed class GetVrfAddressFamilyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Address family.
        /// </summary>
        [Input("addressFamily", required: true)]
        public Input<string> AddressFamily { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf", required: true)]
        public Input<string> Vrf { get; set; } = null!;

        public GetVrfAddressFamilyInvokeArgs()
        {
        }
        public static new GetVrfAddressFamilyInvokeArgs Empty => new GetVrfAddressFamilyInvokeArgs();
    }


    [OutputType]
    public sealed class GetVrfAddressFamilyResult
    {
        /// <summary>
        /// Address family.
        /// </summary>
        public readonly string AddressFamily;
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// VRF name.
        /// </summary>
        public readonly string Vrf;

        [OutputConstructor]
        private GetVrfAddressFamilyResult(
            string addressFamily,

            string? device,

            string id,

            string vrf)
        {
            AddressFamily = addressFamily;
            Device = device;
            Id = id;
            Vrf = vrf;
        }
    }
}
