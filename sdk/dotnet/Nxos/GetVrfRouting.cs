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
    public static class GetVrfRouting
    {
        /// <summary>
        /// This data source can read a VRF Route Distinguisher and VRF VNI.
        /// 
        /// - API Documentation: [rtctrlDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:Dom/)
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
        ///     var example = Nxos.GetVrfRouting.Invoke(new()
        ///     {
        ///         Vrf = "VRF1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVrfRoutingResult> InvokeAsync(GetVrfRoutingArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVrfRoutingResult>("nxos:index/getVrfRouting:getVrfRouting", args ?? new GetVrfRoutingArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read a VRF Route Distinguisher and VRF VNI.
        /// 
        /// - API Documentation: [rtctrlDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:Dom/)
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
        ///     var example = Nxos.GetVrfRouting.Invoke(new()
        ///     {
        ///         Vrf = "VRF1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVrfRoutingResult> Invoke(GetVrfRoutingInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVrfRoutingResult>("nxos:index/getVrfRouting:getVrfRouting", args ?? new GetVrfRoutingInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVrfRoutingArgs : global::Pulumi.InvokeArgs
    {
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

        public GetVrfRoutingArgs()
        {
        }
        public static new GetVrfRoutingArgs Empty => new GetVrfRoutingArgs();
    }

    public sealed class GetVrfRoutingInvokeArgs : global::Pulumi.InvokeArgs
    {
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

        public GetVrfRoutingInvokeArgs()
        {
        }
        public static new GetVrfRoutingInvokeArgs Empty => new GetVrfRoutingInvokeArgs();
    }


    [OutputType]
    public sealed class GetVrfRoutingResult
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
        /// Route Distinguisher value in NX-OS DME format.
        /// </summary>
        public readonly string RouteDistinguisher;
        /// <summary>
        /// VRF name.
        /// </summary>
        public readonly string Vrf;

        [OutputConstructor]
        private GetVrfRoutingResult(
            string? device,

            string id,

            string routeDistinguisher,

            string vrf)
        {
            Device = device;
            Id = id;
            RouteDistinguisher = routeDistinguisher;
            Vrf = vrf;
        }
    }
}
