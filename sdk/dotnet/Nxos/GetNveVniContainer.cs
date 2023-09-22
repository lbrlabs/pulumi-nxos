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
    public static class GetNveVniContainer
    {
        /// <summary>
        /// This data source can read the configuration of Container object for Virtual Network IDs (VNIs).
        /// 
        /// - API Documentation: [nvoNws](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Network%20Virtualization/nvo:Nws/)
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
        ///     var example = Nxos.GetNveVniContainer.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetNveVniContainerResult> InvokeAsync(GetNveVniContainerArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNveVniContainerResult>("nxos:index/getNveVniContainer:getNveVniContainer", args ?? new GetNveVniContainerArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the configuration of Container object for Virtual Network IDs (VNIs).
        /// 
        /// - API Documentation: [nvoNws](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Network%20Virtualization/nvo:Nws/)
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
        ///     var example = Nxos.GetNveVniContainer.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetNveVniContainerResult> Invoke(GetNveVniContainerInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNveVniContainerResult>("nxos:index/getNveVniContainer:getNveVniContainer", args ?? new GetNveVniContainerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNveVniContainerArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        public GetNveVniContainerArgs()
        {
        }
        public static new GetNveVniContainerArgs Empty => new GetNveVniContainerArgs();
    }

    public sealed class GetNveVniContainerInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        public GetNveVniContainerInvokeArgs()
        {
        }
        public static new GetNveVniContainerInvokeArgs Empty => new GetNveVniContainerInvokeArgs();
    }


    [OutputType]
    public sealed class GetNveVniContainerResult
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetNveVniContainerResult(
            string? device,

            string id)
        {
            Device = device;
            Id = id;
        }
    }
}