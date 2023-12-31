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
    public static class GetFeaturePvlan
    {
        /// <summary>
        /// This data source can read the PVLAN feature configuration.
        /// 
        /// - API Documentation: [fmPvlan](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Pvlan/)
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
        ///     var example = Nxos.GetFeaturePvlan.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFeaturePvlanResult> InvokeAsync(GetFeaturePvlanArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFeaturePvlanResult>("nxos:index/getFeaturePvlan:getFeaturePvlan", args ?? new GetFeaturePvlanArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the PVLAN feature configuration.
        /// 
        /// - API Documentation: [fmPvlan](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Pvlan/)
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
        ///     var example = Nxos.GetFeaturePvlan.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFeaturePvlanResult> Invoke(GetFeaturePvlanInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFeaturePvlanResult>("nxos:index/getFeaturePvlan:getFeaturePvlan", args ?? new GetFeaturePvlanInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFeaturePvlanArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        public GetFeaturePvlanArgs()
        {
        }
        public static new GetFeaturePvlanArgs Empty => new GetFeaturePvlanArgs();
    }

    public sealed class GetFeaturePvlanInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        public GetFeaturePvlanInvokeArgs()
        {
        }
        public static new GetFeaturePvlanInvokeArgs Empty => new GetFeaturePvlanInvokeArgs();
    }


    [OutputType]
    public sealed class GetFeaturePvlanResult
    {
        /// <summary>
        /// Administrative state.
        /// </summary>
        public readonly string AdminState;
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetFeaturePvlanResult(
            string adminState,

            string? device,

            string id)
        {
            AdminState = adminState;
            Device = device;
            Id = id;
        }
    }
}
