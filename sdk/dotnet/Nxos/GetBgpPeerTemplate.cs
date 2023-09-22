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
    public static class GetBgpPeerTemplate
    {
        /// <summary>
        /// This data source can read the BGP peer template configuration.
        /// 
        /// - API Documentation: [bgpPeerCont](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:PeerCont/)
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
        ///     var example = Nxos.GetBgpPeerTemplate.Invoke(new()
        ///     {
        ///         Asn = "65001",
        ///         TemplateName = "SPINE-PEERS",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetBgpPeerTemplateResult> InvokeAsync(GetBgpPeerTemplateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBgpPeerTemplateResult>("nxos:index/getBgpPeerTemplate:getBgpPeerTemplate", args ?? new GetBgpPeerTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// This data source can read the BGP peer template configuration.
        /// 
        /// - API Documentation: [bgpPeerCont](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:PeerCont/)
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
        ///     var example = Nxos.GetBgpPeerTemplate.Invoke(new()
        ///     {
        ///         Asn = "65001",
        ///         TemplateName = "SPINE-PEERS",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetBgpPeerTemplateResult> Invoke(GetBgpPeerTemplateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBgpPeerTemplateResult>("nxos:index/getBgpPeerTemplate:getBgpPeerTemplate", args ?? new GetBgpPeerTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBgpPeerTemplateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Autonomous system number.
        /// </summary>
        [Input("asn", required: true)]
        public string Asn { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public string? Device { get; set; }

        /// <summary>
        /// Peer template name.
        /// </summary>
        [Input("templateName", required: true)]
        public string TemplateName { get; set; } = null!;

        public GetBgpPeerTemplateArgs()
        {
        }
        public static new GetBgpPeerTemplateArgs Empty => new GetBgpPeerTemplateArgs();
    }

    public sealed class GetBgpPeerTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Autonomous system number.
        /// </summary>
        [Input("asn", required: true)]
        public Input<string> Asn { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Peer template name.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        public GetBgpPeerTemplateInvokeArgs()
        {
        }
        public static new GetBgpPeerTemplateInvokeArgs Empty => new GetBgpPeerTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetBgpPeerTemplateResult
    {
        /// <summary>
        /// Autonomous system number.
        /// </summary>
        public readonly string Asn;
        /// <summary>
        /// Peer template description.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        public readonly string? Device;
        /// <summary>
        /// The distinguished name of the object.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Neighbor Fabric Type.
        /// </summary>
        public readonly string PeerType;
        /// <summary>
        /// Peer template autonomous system number.
        /// </summary>
        public readonly string RemoteAsn;
        /// <summary>
        /// Source Interface. Must match first field in the output of `show intf brief`.
        /// </summary>
        public readonly string SourceInterface;
        /// <summary>
        /// Peer template name.
        /// </summary>
        public readonly string TemplateName;

        [OutputConstructor]
        private GetBgpPeerTemplateResult(
            string asn,

            string description,

            string? device,

            string id,

            string peerType,

            string remoteAsn,

            string sourceInterface,

            string templateName)
        {
            Asn = asn;
            Description = description;
            Device = device;
            Id = id;
            PeerType = peerType;
            RemoteAsn = remoteAsn;
            SourceInterface = sourceInterface;
            TemplateName = templateName;
        }
    }
}
