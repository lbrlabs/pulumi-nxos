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
    public static class GetFeaturePim
    {
        /// <summary>
        /// This data source can read the PIM feature configuration.
        /// 
        /// - API Documentation: [fmPim](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Pim/)
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
        ///     var example = Nxos.GetFeaturePim.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFeaturePimResult> InvokeAsync(GetFeaturePimArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFeaturePimResult>("nxos:index/getFeaturePim:getFeaturePim", args ?? new GetFeaturePimArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the PIM feature configuration.
        /// 
        /// - API Documentation: [fmPim](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Pim/)
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
        ///     var example = Nxos.GetFeaturePim.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFeaturePimResult> Invoke(GetFeaturePimInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFeaturePimResult>("nxos:index/getFeaturePim:getFeaturePim", args ?? new GetFeaturePimInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFeaturePimArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        public GetFeaturePimArgs()
        {
        }
        public static new GetFeaturePimArgs Empty => new GetFeaturePimArgs();
    }

    public sealed class GetFeaturePimInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        public GetFeaturePimInvokeArgs()
        {
        }
        public static new GetFeaturePimInvokeArgs Empty => new GetFeaturePimInvokeArgs();
    }


    [OutputType]
    public sealed class GetFeaturePimResult
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
        private GetFeaturePimResult(
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