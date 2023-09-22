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
    public static class GetIpv4Vrf
    {
        /// <summary>
        /// This data source can read the IPv4 VRF information.
        /// 
        /// - API Documentation: [ipv4Dom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/ipv4:Dom/)
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
        ///     var example = Nxos.GetIpv4Vrf.Invoke(new()
        ///     {
        ///         Name = "VRF1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIpv4VrfResult> InvokeAsync(GetIpv4VrfArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpv4VrfResult>("nxos:index/getIpv4Vrf:getIpv4Vrf", args ?? new GetIpv4VrfArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the IPv4 VRF information.
        /// 
        /// - API Documentation: [ipv4Dom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/ipv4:Dom/)
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
        ///     var example = Nxos.GetIpv4Vrf.Invoke(new()
        ///     {
        ///         Name = "VRF1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIpv4VrfResult> Invoke(GetIpv4VrfInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpv4VrfResult>("nxos:index/getIpv4Vrf:getIpv4Vrf", args ?? new GetIpv4VrfInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpv4VrfArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetIpv4VrfArgs()
        {
        }
        public static new GetIpv4VrfArgs Empty => new GetIpv4VrfArgs();
    }

    public sealed class GetIpv4VrfInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetIpv4VrfInvokeArgs()
        {
        }
        public static new GetIpv4VrfInvokeArgs Empty => new GetIpv4VrfInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpv4VrfResult
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
        /// VRF name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetIpv4VrfResult(
            string? device,

            string id,

            string name)
        {
            Device = device;
            Id = id;
            Name = name;
        }
    }
}
