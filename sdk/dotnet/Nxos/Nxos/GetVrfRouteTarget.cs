// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Nxos.Nxos
{
    public static class GetVrfRouteTarget
    {
        public static Task<GetVrfRouteTargetResult> InvokeAsync(GetVrfRouteTargetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetVrfRouteTargetResult>("nxos:nxos/getVrfRouteTarget:getVrfRouteTarget", args ?? new GetVrfRouteTargetArgs(), options.WithDefaults());

        public static Output<GetVrfRouteTargetResult> Invoke(GetVrfRouteTargetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetVrfRouteTargetResult>("nxos:nxos/getVrfRouteTarget:getVrfRouteTarget", args ?? new GetVrfRouteTargetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVrfRouteTargetArgs : global::Pulumi.InvokeArgs
    {
        [Input("addressFamily", required: true)]
        public string AddressFamily { get; set; } = null!;

        [Input("device")]
        public string? Device { get; set; }

        [Input("direction", required: true)]
        public string Direction { get; set; } = null!;

        [Input("routeTarget", required: true)]
        public string RouteTarget { get; set; } = null!;

        [Input("routeTargetAddressFamily", required: true)]
        public string RouteTargetAddressFamily { get; set; } = null!;

        [Input("vrf", required: true)]
        public string Vrf { get; set; } = null!;

        public GetVrfRouteTargetArgs()
        {
        }
        public static new GetVrfRouteTargetArgs Empty => new GetVrfRouteTargetArgs();
    }

    public sealed class GetVrfRouteTargetInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("addressFamily", required: true)]
        public Input<string> AddressFamily { get; set; } = null!;

        [Input("device")]
        public Input<string>? Device { get; set; }

        [Input("direction", required: true)]
        public Input<string> Direction { get; set; } = null!;

        [Input("routeTarget", required: true)]
        public Input<string> RouteTarget { get; set; } = null!;

        [Input("routeTargetAddressFamily", required: true)]
        public Input<string> RouteTargetAddressFamily { get; set; } = null!;

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
        public readonly string AddressFamily;
        public readonly string? Device;
        public readonly string Direction;
        public readonly string Id;
        public readonly string RouteTarget;
        public readonly string RouteTargetAddressFamily;
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