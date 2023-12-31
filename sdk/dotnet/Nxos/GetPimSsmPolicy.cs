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
    public static class GetPimSsmPolicy
    {
        /// <summary>
        /// This data source can read the PIM SSM policy configuration.
        /// 
        /// - API Documentation: [pimSSMPatP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:SSMPatP/)
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
        ///     var example = Nxos.GetPimSsmPolicy.Invoke(new()
        ///     {
        ///         VrfName = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPimSsmPolicyResult> InvokeAsync(GetPimSsmPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPimSsmPolicyResult>("nxos:index/getPimSsmPolicy:getPimSsmPolicy", args ?? new GetPimSsmPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the PIM SSM policy configuration.
        /// 
        /// - API Documentation: [pimSSMPatP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:SSMPatP/)
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
        ///     var example = Nxos.GetPimSsmPolicy.Invoke(new()
        ///     {
        ///         VrfName = "default",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPimSsmPolicyResult> Invoke(GetPimSsmPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPimSsmPolicyResult>("nxos:index/getPimSsmPolicy:getPimSsmPolicy", args ?? new GetPimSsmPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPimSsmPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrfName", required: true)]
        public string VrfName { get; set; } = null!;

        public GetPimSsmPolicyArgs()
        {
        }
        public static new GetPimSsmPolicyArgs Empty => new GetPimSsmPolicyArgs();
    }

    public sealed class GetPimSsmPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrfName", required: true)]
        public Input<string> VrfName { get; set; } = null!;

        public GetPimSsmPolicyInvokeArgs()
        {
        }
        public static new GetPimSsmPolicyInvokeArgs Empty => new GetPimSsmPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetPimSsmPolicyResult
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
        /// Policy name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// VRF name.
        /// </summary>
        public readonly string VrfName;

        [OutputConstructor]
        private GetPimSsmPolicyResult(
            string? device,

            string id,

            string name,

            string vrfName)
        {
            Device = device;
            Id = id;
            Name = name;
            VrfName = vrfName;
        }
    }
}
