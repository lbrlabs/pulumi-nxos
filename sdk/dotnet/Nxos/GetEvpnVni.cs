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
    public static class GetEvpnVni
    {
        /// <summary>
        /// This data source can read a EVPN VNI Route Distinguisher.
        /// 
        /// - API Documentation: [rtctrlBDEvi](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:BDEvi/)
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
        ///     var example = Nxos.GetEvpnVni.Invoke(new()
        ///     {
        ///         Encap = "vxlan-123456",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEvpnVniResult> InvokeAsync(GetEvpnVniArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEvpnVniResult>("nxos:index/getEvpnVni:getEvpnVni", args ?? new GetEvpnVniArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read a EVPN VNI Route Distinguisher.
        /// 
        /// - API Documentation: [rtctrlBDEvi](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:BDEvi/)
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
        ///     var example = Nxos.GetEvpnVni.Invoke(new()
        ///     {
        ///         Encap = "vxlan-123456",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEvpnVniResult> Invoke(GetEvpnVniInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEvpnVniResult>("nxos:index/getEvpnVni:getEvpnVni", args ?? new GetEvpnVniInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEvpnVniArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
        /// </summary>
        [Input("encap", required: true)]
        public string Encap { get; set; } = null!;

        public GetEvpnVniArgs()
        {
        }
        public static new GetEvpnVniArgs Empty => new GetEvpnVniArgs();
    }

    public sealed class GetEvpnVniInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
        /// </summary>
        [Input("encap", required: true)]
        public Input<string> Encap { get; set; } = null!;

        public GetEvpnVniInvokeArgs()
        {
        }
        public static new GetEvpnVniInvokeArgs Empty => new GetEvpnVniInvokeArgs();
    }


    [OutputType]
    public sealed class GetEvpnVniResult
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
        /// </summary>
        public readonly string Encap;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Route Distinguisher value in NX-OS DME format.
        /// </summary>
        public readonly string RouteDistinguisher;

        [OutputConstructor]
        private GetEvpnVniResult(
            string? device,

            string encap,

            string id,

            string routeDistinguisher)
        {
            Device = device;
            Encap = encap;
            Id = id;
            RouteDistinguisher = routeDistinguisher;
        }
    }
}
