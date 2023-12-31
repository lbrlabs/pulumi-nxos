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
    public static class GetFeatureVnSegment
    {
        /// <summary>
        /// This data source can read the VN Segment feature configuration.
        /// 
        /// - API Documentation: [fmVnSegment](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:VnSegment/)
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
        ///     var example = Nxos.GetFeatureVnSegment.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetFeatureVnSegmentResult> InvokeAsync(GetFeatureVnSegmentArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetFeatureVnSegmentResult>("nxos:index/getFeatureVnSegment:getFeatureVnSegment", args ?? new GetFeatureVnSegmentArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the VN Segment feature configuration.
        /// 
        /// - API Documentation: [fmVnSegment](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:VnSegment/)
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
        ///     var example = Nxos.GetFeatureVnSegment.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetFeatureVnSegmentResult> Invoke(GetFeatureVnSegmentInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetFeatureVnSegmentResult>("nxos:index/getFeatureVnSegment:getFeatureVnSegment", args ?? new GetFeatureVnSegmentInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetFeatureVnSegmentArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        public GetFeatureVnSegmentArgs()
        {
        }
        public static new GetFeatureVnSegmentArgs Empty => new GetFeatureVnSegmentArgs();
    }

    public sealed class GetFeatureVnSegmentInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        public GetFeatureVnSegmentInvokeArgs()
        {
        }
        public static new GetFeatureVnSegmentInvokeArgs Empty => new GetFeatureVnSegmentInvokeArgs();
    }


    [OutputType]
    public sealed class GetFeatureVnSegmentResult
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
        private GetFeatureVnSegmentResult(
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
