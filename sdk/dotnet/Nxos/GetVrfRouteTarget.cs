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
    public static class GetVrfRouteTarget
    {
        /// <summary>
        /// This data source can read a VRF Route Target Entry.
        /// 
        /// - API Documentation: [rtctrlRttEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:RttEntry/)
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
        ///     var example = Nxos.GetVrfRouteTarget.Invoke(new()
        ///     {
        ///         AddressFamily = "ipv4-ucast",
        ///         Direction = "import",
        ///         RouteTarget = "route-target:as2-nn2:2:2",
        ///         RouteTargetAddressFamily = "ipv4-ucast",
        ///         Vrf = "VRF1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVrfRouteTargetResult> InvokeAsync(GetVrfRouteTargetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVrfRouteTargetResult>("nxos:index/getVrfRouteTarget:getVrfRouteTarget", args ?? new GetVrfRouteTargetArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read a VRF Route Target Entry.
        /// 
        /// - API Documentation: [rtctrlRttEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:RttEntry/)
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
        ///     var example = Nxos.GetVrfRouteTarget.Invoke(new()
        ///     {
        ///         AddressFamily = "ipv4-ucast",
        ///         Direction = "import",
        ///         RouteTarget = "route-target:as2-nn2:2:2",
        ///         RouteTargetAddressFamily = "ipv4-ucast",
        ///         Vrf = "VRF1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVrfRouteTargetResult> Invoke(GetVrfRouteTargetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVrfRouteTargetResult>("nxos:index/getVrfRouteTarget:getVrfRouteTarget", args ?? new GetVrfRouteTargetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVrfRouteTargetArgs : global::Pulumi.InvokeArgs
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
        /// Route Target direction.
        /// </summary>
        [Input("direction", required: true)]
        public string Direction { get; set; } = null!;

        /// <summary>
        /// Route Target in NX-OS DME format.
        /// </summary>
        [Input("routeTarget", required: true)]
        public string RouteTarget { get; set; } = null!;

        /// <summary>
        /// Route Target Address Family.
        /// </summary>
        [Input("routeTargetAddressFamily", required: true)]
        public string RouteTargetAddressFamily { get; set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf", required: true)]
        public string Vrf { get; set; } = null!;

        public GetVrfRouteTargetArgs()
        {
        }
        public static new GetVrfRouteTargetArgs Empty => new GetVrfRouteTargetArgs();
    }

    public sealed class GetVrfRouteTargetInvokeArgs : global::Pulumi.InvokeArgs
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
        /// Route Target direction.
        /// </summary>
        [Input("direction", required: true)]
        public Input<string> Direction { get; set; } = null!;

        /// <summary>
        /// Route Target in NX-OS DME format.
        /// </summary>
        [Input("routeTarget", required: true)]
        public Input<string> RouteTarget { get; set; } = null!;

        /// <summary>
        /// Route Target Address Family.
        /// </summary>
        [Input("routeTargetAddressFamily", required: true)]
        public Input<string> RouteTargetAddressFamily { get; set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf", required: true)]
        public Input<string> Vrf { get; set; } = null!;

        public GetVrfRouteTargetInvokeArgs()
        {
        }
        public static new GetVrfRouteTargetInvokeArgs Empty => new GetVrfRouteTargetInvokeArgs();
    }


    [OutputType]
    public sealed class GetVrfRouteTargetResult
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
        /// Route Target direction.
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Route Target in NX-OS DME format.
        /// </summary>
        public readonly string RouteTarget;
        /// <summary>
        /// Route Target Address Family.
        /// </summary>
        public readonly string RouteTargetAddressFamily;
        /// <summary>
        /// VRF name.
        /// </summary>
        public readonly string Vrf;

        [OutputConstructor]
        private GetVrfRouteTargetResult(
            string addressFamily,

            string? device,

            string direction,

            string id,

            string routeTarget,

            string routeTargetAddressFamily,

            string vrf)
        {
            AddressFamily = addressFamily;
            Device = device;
            Direction = direction;
            Id = id;
            RouteTarget = routeTarget;
            RouteTargetAddressFamily = routeTargetAddressFamily;
            Vrf = vrf;
        }
    }
}
