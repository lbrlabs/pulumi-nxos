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
    public static class GetIpv4AccessListPolicyEgressInterface
    {
        /// <summary>
        /// This data source can read an IPv4 Access List Policy Egress Interface.
        /// 
        /// - API Documentation: [aclIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Security%20and%20Policing/acl:If/)
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
        ///     var example = Nxos.GetIpv4AccessListPolicyEgressInterface.Invoke(new()
        ///     {
        ///         InterfaceId = "eth1/10",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIpv4AccessListPolicyEgressInterfaceResult> InvokeAsync(GetIpv4AccessListPolicyEgressInterfaceArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpv4AccessListPolicyEgressInterfaceResult>("nxos:index/getIpv4AccessListPolicyEgressInterface:getIpv4AccessListPolicyEgressInterface", args ?? new GetIpv4AccessListPolicyEgressInterfaceArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read an IPv4 Access List Policy Egress Interface.
        /// 
        /// - API Documentation: [aclIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Security%20and%20Policing/acl:If/)
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
        ///     var example = Nxos.GetIpv4AccessListPolicyEgressInterface.Invoke(new()
        ///     {
        ///         InterfaceId = "eth1/10",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIpv4AccessListPolicyEgressInterfaceResult> Invoke(GetIpv4AccessListPolicyEgressInterfaceInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpv4AccessListPolicyEgressInterfaceResult>("nxos:index/getIpv4AccessListPolicyEgressInterface:getIpv4AccessListPolicyEgressInterface", args ?? new GetIpv4AccessListPolicyEgressInterfaceInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpv4AccessListPolicyEgressInterfaceArgs : global::Pulumi.InvokeArgs
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

        public GetIpv4AccessListPolicyEgressInterfaceArgs()
        {
        }
        public static new GetIpv4AccessListPolicyEgressInterfaceArgs Empty => new GetIpv4AccessListPolicyEgressInterfaceArgs();
    }

    public sealed class GetIpv4AccessListPolicyEgressInterfaceInvokeArgs : global::Pulumi.InvokeArgs
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

        public GetIpv4AccessListPolicyEgressInterfaceInvokeArgs()
        {
        }
        public static new GetIpv4AccessListPolicyEgressInterfaceInvokeArgs Empty => new GetIpv4AccessListPolicyEgressInterfaceInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpv4AccessListPolicyEgressInterfaceResult
    {
        /// <summary>
        /// Access list name.
        /// </summary>
        public readonly string AccessListName;
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
        private GetIpv4AccessListPolicyEgressInterfaceResult(
            string accessListName,

            string? device,

            string id,

            string interfaceId)
        {
            AccessListName = accessListName;
            Device = device;
            Id = id;
            InterfaceId = interfaceId;
        }
    }
}
