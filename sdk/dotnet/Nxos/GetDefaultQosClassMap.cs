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
    public static class GetDefaultQosClassMap
    {
        /// <summary>
        /// This data source can read the default QoS class map configuration.
        /// 
        /// - API Documentation: [ipqosCMapInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:CMapInst/)
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
        ///     var example = Nxos.GetDefaultQosClassMap.Invoke(new()
        ///     {
        ///         Name = "Voice",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDefaultQosClassMapResult> InvokeAsync(GetDefaultQosClassMapArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDefaultQosClassMapResult>("nxos:index/getDefaultQosClassMap:getDefaultQosClassMap", args ?? new GetDefaultQosClassMapArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the default QoS class map configuration.
        /// 
        /// - API Documentation: [ipqosCMapInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:CMapInst/)
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
        ///     var example = Nxos.GetDefaultQosClassMap.Invoke(new()
        ///     {
        ///         Name = "Voice",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDefaultQosClassMapResult> Invoke(GetDefaultQosClassMapInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDefaultQosClassMapResult>("nxos:index/getDefaultQosClassMap:getDefaultQosClassMap", args ?? new GetDefaultQosClassMapInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDefaultQosClassMapArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// Class map name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetDefaultQosClassMapArgs()
        {
        }
        public static new GetDefaultQosClassMapArgs Empty => new GetDefaultQosClassMapArgs();
    }

    public sealed class GetDefaultQosClassMapInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Class map name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetDefaultQosClassMapInvokeArgs()
        {
        }
        public static new GetDefaultQosClassMapInvokeArgs Empty => new GetDefaultQosClassMapInvokeArgs();
    }


    [OutputType]
    public sealed class GetDefaultQosClassMapResult
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
        /// Class map name.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetDefaultQosClassMapResult(
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
