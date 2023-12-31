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
    public static class GetSystem
    {
        /// <summary>
        /// This data source can read the system configuration.
        /// 
        /// - API Documentation: [topSystem](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/System/top:System/)
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
        ///     var example = Nxos.GetSystem.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetSystemResult> InvokeAsync(GetSystemArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSystemResult>("nxos:index/getSystem:getSystem", args ?? new GetSystemArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the system configuration.
        /// 
        /// - API Documentation: [topSystem](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/System/top:System/)
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
        ///     var example = Nxos.GetSystem.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetSystemResult> Invoke(GetSystemInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSystemResult>("nxos:index/getSystem:getSystem", args ?? new GetSystemInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSystemArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        public GetSystemArgs()
        {
        }
        public static new GetSystemArgs Empty => new GetSystemArgs();
    }

    public sealed class GetSystemInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        public GetSystemInvokeArgs()
        {
        }
        public static new GetSystemInvokeArgs Empty => new GetSystemInvokeArgs();
    }


    [OutputType]
    public sealed class GetSystemResult
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
        /// The system name (hostname).
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetSystemResult(
            string? device,

            string id,

            string name)
        {
            Device = device;
            Id = id;
            Name = name;
        }
    }
}
