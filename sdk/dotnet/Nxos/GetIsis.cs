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
    public static class GetIsis
    {
        /// <summary>
        /// This data source can read the global IS-IS configuration.
        /// 
        /// - API Documentation: [isisEntity](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:Entity/)
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
        ///     var example = Nxos.GetIsis.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIsisResult> InvokeAsync(GetIsisArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIsisResult>("nxos:index/getIsis:getIsis", args ?? new GetIsisArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the global IS-IS configuration.
        /// 
        /// - API Documentation: [isisEntity](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:Entity/)
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
        ///     var example = Nxos.GetIsis.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIsisResult> Invoke(GetIsisInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIsisResult>("nxos:index/getIsis:getIsis", args ?? new GetIsisInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIsisArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        public GetIsisArgs()
        {
        }
        public static new GetIsisArgs Empty => new GetIsisArgs();
    }

    public sealed class GetIsisInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        public GetIsisInvokeArgs()
        {
        }
        public static new GetIsisInvokeArgs Empty => new GetIsisInvokeArgs();
    }


    [OutputType]
    public sealed class GetIsisResult
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
        private GetIsisResult(
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
