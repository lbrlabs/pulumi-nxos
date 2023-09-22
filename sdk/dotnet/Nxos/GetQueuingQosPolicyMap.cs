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
    public static class GetQueuingQosPolicyMap
    {
        /// <summary>
        /// This data source can read the queuing QoS policy map configuration.
        /// 
        /// - API Documentation: [ipqosPMapInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:PMapInst/)
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
        ///     var example = Nxos.GetQueuingQosPolicyMap.Invoke(new()
        ///     {
        ///         Name = "PM1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetQueuingQosPolicyMapResult> InvokeAsync(GetQueuingQosPolicyMapArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetQueuingQosPolicyMapResult>("nxos:index/getQueuingQosPolicyMap:getQueuingQosPolicyMap", args ?? new GetQueuingQosPolicyMapArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the queuing QoS policy map configuration.
        /// 
        /// - API Documentation: [ipqosPMapInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:PMapInst/)
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
        ///     var example = Nxos.GetQueuingQosPolicyMap.Invoke(new()
        ///     {
        ///         Name = "PM1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetQueuingQosPolicyMapResult> Invoke(GetQueuingQosPolicyMapInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetQueuingQosPolicyMapResult>("nxos:index/getQueuingQosPolicyMap:getQueuingQosPolicyMap", args ?? new GetQueuingQosPolicyMapInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetQueuingQosPolicyMapArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetQueuingQosPolicyMapArgs()
        {
        }
        public static new GetQueuingQosPolicyMapArgs Empty => new GetQueuingQosPolicyMapArgs();
    }

    public sealed class GetQueuingQosPolicyMapInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetQueuingQosPolicyMapInvokeArgs()
        {
        }
        public static new GetQueuingQosPolicyMapInvokeArgs Empty => new GetQueuingQosPolicyMapInvokeArgs();
    }


    [OutputType]
    public sealed class GetQueuingQosPolicyMapResult
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
        /// Match type.
        /// </summary>
        public readonly string MatchType;
        /// <summary>
        /// Policy map name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetQueuingQosPolicyMapResult(
            string? device,

            string id,

            string matchType,

            string name)
        {
            Device = device;
            Id = id;
            MatchType = matchType;
            Name = name;
        }
    }
}
