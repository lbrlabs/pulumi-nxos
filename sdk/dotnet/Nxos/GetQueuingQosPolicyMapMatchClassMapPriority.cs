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
    public static class GetQueuingQosPolicyMapMatchClassMapPriority
    {
        /// <summary>
        /// This data source can read the queuing QoS policy map match class map priority configuration.
        /// 
        /// - API Documentation: [ipqosPriority](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:Priority/)
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
        ///     var example = Nxos.GetQueuingQosPolicyMapMatchClassMapPriority.Invoke(new()
        ///     {
        ///         ClassMapName = "c-out-q1",
        ///         PolicyMapName = "PM1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetQueuingQosPolicyMapMatchClassMapPriorityResult> InvokeAsync(GetQueuingQosPolicyMapMatchClassMapPriorityArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetQueuingQosPolicyMapMatchClassMapPriorityResult>("nxos:index/getQueuingQosPolicyMapMatchClassMapPriority:getQueuingQosPolicyMapMatchClassMapPriority", args ?? new GetQueuingQosPolicyMapMatchClassMapPriorityArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the queuing QoS policy map match class map priority configuration.
        /// 
        /// - API Documentation: [ipqosPriority](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:Priority/)
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
        ///     var example = Nxos.GetQueuingQosPolicyMapMatchClassMapPriority.Invoke(new()
        ///     {
        ///         ClassMapName = "c-out-q1",
        ///         PolicyMapName = "PM1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetQueuingQosPolicyMapMatchClassMapPriorityResult> Invoke(GetQueuingQosPolicyMapMatchClassMapPriorityInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetQueuingQosPolicyMapMatchClassMapPriorityResult>("nxos:index/getQueuingQosPolicyMapMatchClassMapPriority:getQueuingQosPolicyMapMatchClassMapPriority", args ?? new GetQueuingQosPolicyMapMatchClassMapPriorityInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetQueuingQosPolicyMapMatchClassMapPriorityArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Class map name.
        /// </summary>
        [Input("classMapName", required: true)]
        public string ClassMapName { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Input("policyMapName", required: true)]
        public string PolicyMapName { get; set; } = null!;

        public GetQueuingQosPolicyMapMatchClassMapPriorityArgs()
        {
        }
        public static new GetQueuingQosPolicyMapMatchClassMapPriorityArgs Empty => new GetQueuingQosPolicyMapMatchClassMapPriorityArgs();
    }

    public sealed class GetQueuingQosPolicyMapMatchClassMapPriorityInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Class map name.
        /// </summary>
        [Input("classMapName", required: true)]
        public Input<string> ClassMapName { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Input("policyMapName", required: true)]
        public Input<string> PolicyMapName { get; set; } = null!;

        public GetQueuingQosPolicyMapMatchClassMapPriorityInvokeArgs()
        {
        }
        public static new GetQueuingQosPolicyMapMatchClassMapPriorityInvokeArgs Empty => new GetQueuingQosPolicyMapMatchClassMapPriorityInvokeArgs();
    }


    [OutputType]
    public sealed class GetQueuingQosPolicyMapMatchClassMapPriorityResult
    {
        /// <summary>
        /// Class map name.
        /// </summary>
        public readonly string ClassMapName;
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Priority level.
        /// </summary>
        public readonly int Level;
        /// <summary>
        /// Policy map name.
        /// </summary>
        public readonly string PolicyMapName;

        [OutputConstructor]
        private GetQueuingQosPolicyMapMatchClassMapPriorityResult(
            string classMapName,

            string? device,

            string id,

            int level,

            string policyMapName)
        {
            ClassMapName = classMapName;
            Device = device;
            Id = id;
            Level = level;
            PolicyMapName = policyMapName;
        }
    }
}
