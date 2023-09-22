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
    public static class GetIpv4StaticRoute
    {
        /// <summary>
        /// This data source can read an IPv4 static route.
        /// 
        /// - API Documentation: [ipv4Route](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/ipv4:Route/)
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
        ///     var example = Nxos.GetIpv4StaticRoute.Invoke(new()
        ///     {
        ///         Prefix = "1.1.1.0/24",
        ///         VrfName = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIpv4StaticRouteResult> InvokeAsync(GetIpv4StaticRouteArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpv4StaticRouteResult>("nxos:index/getIpv4StaticRoute:getIpv4StaticRoute", args ?? new GetIpv4StaticRouteArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read an IPv4 static route.
        /// 
        /// - API Documentation: [ipv4Route](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/ipv4:Route/)
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
        ///     var example = Nxos.GetIpv4StaticRoute.Invoke(new()
        ///     {
        ///         Prefix = "1.1.1.0/24",
        ///         VrfName = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIpv4StaticRouteResult> Invoke(GetIpv4StaticRouteInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpv4StaticRouteResult>("nxos:index/getIpv4StaticRoute:getIpv4StaticRoute", args ?? new GetIpv4StaticRouteInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpv4StaticRouteArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// Prefix.
        /// </summary>
        [Input("prefix", required: true)]
        public string Prefix { get; set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrfName", required: true)]
        public string VrfName { get; set; } = null!;

        public GetIpv4StaticRouteArgs()
        {
        }
        public static new GetIpv4StaticRouteArgs Empty => new GetIpv4StaticRouteArgs();
    }

    public sealed class GetIpv4StaticRouteInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Prefix.
        /// </summary>
        [Input("prefix", required: true)]
        public Input<string> Prefix { get; set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrfName", required: true)]
        public Input<string> VrfName { get; set; } = null!;

        public GetIpv4StaticRouteInvokeArgs()
        {
        }
        public static new GetIpv4StaticRouteInvokeArgs Empty => new GetIpv4StaticRouteInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpv4StaticRouteResult
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
        /// List of next hops.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIpv4StaticRouteNextHopResult> NextHops;
        /// <summary>
        /// Prefix.
        /// </summary>
        public readonly string Prefix;
        /// <summary>
        /// VRF name.
        /// </summary>
        public readonly string VrfName;

        [OutputConstructor]
        private GetIpv4StaticRouteResult(
            string? device,

            string id,

            ImmutableArray<Outputs.GetIpv4StaticRouteNextHopResult> nextHops,

            string prefix,

            string vrfName)
        {
            Device = device;
            Id = id;
            NextHops = nextHops;
            Prefix = prefix;
            VrfName = vrfName;
        }
    }
}
