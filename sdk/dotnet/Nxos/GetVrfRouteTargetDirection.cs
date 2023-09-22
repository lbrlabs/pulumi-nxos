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
    public static class GetVrfRouteTargetDirection
    {
        /// <summary>
        /// This data source can read a VRF Route Target direction.
        /// 
        /// - API Documentation: [rtctrlRttP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:RttP/)
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
        ///     var example = Nxos.GetVrfRouteTargetDirection.Invoke(new()
        ///     {
        ///         AddressFamily = "ipv4-ucast",
        ///         Direction = "import",
        ///         RouteTargetAddressFamily = "ipv4-ucast",
        ///         Vrf = "VRF1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetVrfRouteTargetDirectionResult> InvokeAsync(GetVrfRouteTargetDirectionArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVrfRouteTargetDirectionResult>("nxos:index/getVrfRouteTargetDirection:getVrfRouteTargetDirection", args ?? new GetVrfRouteTargetDirectionArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read a VRF Route Target direction.
        /// 
        /// - API Documentation: [rtctrlRttP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:RttP/)
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
        ///     var example = Nxos.GetVrfRouteTargetDirection.Invoke(new()
        ///     {
        ///         AddressFamily = "ipv4-ucast",
        ///         Direction = "import",
        ///         RouteTargetAddressFamily = "ipv4-ucast",
        ///         Vrf = "VRF1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetVrfRouteTargetDirectionResult> Invoke(GetVrfRouteTargetDirectionInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVrfRouteTargetDirectionResult>("nxos:index/getVrfRouteTargetDirection:getVrfRouteTargetDirection", args ?? new GetVrfRouteTargetDirectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVrfRouteTargetDirectionArgs : global::Pulumi.InvokeArgs
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
        /// Route Target Address Family.
        /// </summary>
        [Input("routeTargetAddressFamily", required: true)]
        public string RouteTargetAddressFamily { get; set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf", required: true)]
        public string Vrf { get; set; } = null!;

        public GetVrfRouteTargetDirectionArgs()
        {
        }
        public static new GetVrfRouteTargetDirectionArgs Empty => new GetVrfRouteTargetDirectionArgs();
    }

    public sealed class GetVrfRouteTargetDirectionInvokeArgs : global::Pulumi.InvokeArgs
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
        /// Route Target Address Family.
        /// </summary>
        [Input("routeTargetAddressFamily", required: true)]
        public Input<string> RouteTargetAddressFamily { get; set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf", required: true)]
        public Input<string> Vrf { get; set; } = null!;

        public GetVrfRouteTargetDirectionInvokeArgs()
        {
        }
        public static new GetVrfRouteTargetDirectionInvokeArgs Empty => new GetVrfRouteTargetDirectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetVrfRouteTargetDirectionResult
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
        /// Route Target Address Family.
        /// </summary>
        public readonly string RouteTargetAddressFamily;
        /// <summary>
        /// VRF name.
        /// </summary>
        public readonly string Vrf;

        [OutputConstructor]
        private GetVrfRouteTargetDirectionResult(
            string addressFamily,

            string? device,

            string direction,

            string id,

            string routeTargetAddressFamily,

            string vrf)
        {
            AddressFamily = addressFamily;
            Device = device;
            Direction = direction;
            Id = id;
            RouteTargetAddressFamily = routeTargetAddressFamily;
            Vrf = vrf;
        }
    }
}
