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
    public static class GetQueuingQosPolicySystemOutPolicyMap
    {
        /// <summary>
        /// This data source can read the queuing QoS policy system out policy map configuration.
        /// 
        /// - API Documentation: [ipqosInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:Inst/)
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
        ///     var example = Nxos.GetQueuingQosPolicySystemOutPolicyMap.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetQueuingQosPolicySystemOutPolicyMapResult> InvokeAsync(GetQueuingQosPolicySystemOutPolicyMapArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetQueuingQosPolicySystemOutPolicyMapResult>("nxos:index/getQueuingQosPolicySystemOutPolicyMap:getQueuingQosPolicySystemOutPolicyMap", args ?? new GetQueuingQosPolicySystemOutPolicyMapArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the queuing QoS policy system out policy map configuration.
        /// 
        /// - API Documentation: [ipqosInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:Inst/)
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
        ///     var example = Nxos.GetQueuingQosPolicySystemOutPolicyMap.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetQueuingQosPolicySystemOutPolicyMapResult> Invoke(GetQueuingQosPolicySystemOutPolicyMapInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetQueuingQosPolicySystemOutPolicyMapResult>("nxos:index/getQueuingQosPolicySystemOutPolicyMap:getQueuingQosPolicySystemOutPolicyMap", args ?? new GetQueuingQosPolicySystemOutPolicyMapInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetQueuingQosPolicySystemOutPolicyMapArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        public GetQueuingQosPolicySystemOutPolicyMapArgs()
        {
        }
        public static new GetQueuingQosPolicySystemOutPolicyMapArgs Empty => new GetQueuingQosPolicySystemOutPolicyMapArgs();
    }

    public sealed class GetQueuingQosPolicySystemOutPolicyMapInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        public GetQueuingQosPolicySystemOutPolicyMapInvokeArgs()
        {
        }
        public static new GetQueuingQosPolicySystemOutPolicyMapInvokeArgs Empty => new GetQueuingQosPolicySystemOutPolicyMapInvokeArgs();
    }


    [OutputType]
    public sealed class GetQueuingQosPolicySystemOutPolicyMapResult
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
        /// Policy map name.
        /// </summary>
        public readonly string PolicyMapName;

        [OutputConstructor]
        private GetQueuingQosPolicySystemOutPolicyMapResult(
            string? device,

            string id,

            string policyMapName)
        {
            Device = device;
            Id = id;
            PolicyMapName = policyMapName;
        }
    }
}