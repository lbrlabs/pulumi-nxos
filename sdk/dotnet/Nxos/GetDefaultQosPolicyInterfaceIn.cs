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
    public static class GetDefaultQosPolicyInterfaceIn
    {
        /// <summary>
        /// This data source can read the default QoS policy interface in configuration.
        /// 
        /// - API Documentation: [ipqosIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:If/)
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
        ///     var example = Nxos.GetDefaultQosPolicyInterfaceIn.Invoke(new()
        ///     {
        ///         InterfaceId = "eth1/10",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDefaultQosPolicyInterfaceInResult> InvokeAsync(GetDefaultQosPolicyInterfaceInArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDefaultQosPolicyInterfaceInResult>("nxos:index/getDefaultQosPolicyInterfaceIn:getDefaultQosPolicyInterfaceIn", args ?? new GetDefaultQosPolicyInterfaceInArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the default QoS policy interface in configuration.
        /// 
        /// - API Documentation: [ipqosIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:If/)
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
        ///     var example = Nxos.GetDefaultQosPolicyInterfaceIn.Invoke(new()
        ///     {
        ///         InterfaceId = "eth1/10",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDefaultQosPolicyInterfaceInResult> Invoke(GetDefaultQosPolicyInterfaceInInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDefaultQosPolicyInterfaceInResult>("nxos:index/getDefaultQosPolicyInterfaceIn:getDefaultQosPolicyInterfaceIn", args ?? new GetDefaultQosPolicyInterfaceInInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDefaultQosPolicyInterfaceInArgs : global::Pulumi.InvokeArgs
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

        public GetDefaultQosPolicyInterfaceInArgs()
        {
        }
        public static new GetDefaultQosPolicyInterfaceInArgs Empty => new GetDefaultQosPolicyInterfaceInArgs();
    }

    public sealed class GetDefaultQosPolicyInterfaceInInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetDefaultQosPolicyInterfaceInInvokeArgs()
        {
        }
        public static new GetDefaultQosPolicyInterfaceInInvokeArgs Empty => new GetDefaultQosPolicyInterfaceInInvokeArgs();
    }


    [OutputType]
    public sealed class GetDefaultQosPolicyInterfaceInResult
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
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        public readonly string InterfaceId;

        [OutputConstructor]
        private GetDefaultQosPolicyInterfaceInResult(
            string? device,

            string id,

            string interfaceId)
        {
            Device = device;
            Id = id;
            InterfaceId = interfaceId;
        }
    }
}
