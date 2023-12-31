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
    /// <summary>
    /// This resource can manage the BGP domain (VRF) configuration.
    /// 
    /// - API Documentation: [bgpDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Dom/)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Nxos = Lbrlabs.PulumiPackage.Nxos;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new Nxos.BgpVrf("example", new()
    ///     {
    ///         Asn = "65001",
    ///         RouterId = "1.1.1.1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import nxos:index/bgpVrf:BgpVrf example "sys/bgp/inst/dom-[default]"
    /// ```
    /// </summary>
    [NxosResourceType("nxos:index/bgpVrf:BgpVrf")]
    public partial class BgpVrf : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Autonomous system number.
        /// </summary>
        [Output("asn")]
        public Output<string> Asn { get; private set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Router ID.
        /// </summary>
        [Output("routerId")]
        public Output<string?> RouterId { get; private set; } = null!;


        /// <summary>
        /// Create a BgpVrf resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BgpVrf(string name, BgpVrfArgs args, CustomResourceOptions? options = null)
            : base("nxos:index/bgpVrf:BgpVrf", name, args ?? new BgpVrfArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BgpVrf(string name, Input<string> id, BgpVrfState? state = null, CustomResourceOptions? options = null)
            : base("nxos:index/bgpVrf:BgpVrf", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BgpVrf resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BgpVrf Get(string name, Input<string> id, BgpVrfState? state = null, CustomResourceOptions? options = null)
        {
            return new BgpVrf(name, id, state, options);
        }
    }

    public sealed class BgpVrfArgs : global::Pulumi.ResourceArgs
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
        /// VRF name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Router ID.
        /// </summary>
        [Input("routerId")]
        public Input<string>? RouterId { get; set; }

        public BgpVrfArgs()
        {
        }
        public static new BgpVrfArgs Empty => new BgpVrfArgs();
    }

    public sealed class BgpVrfState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Autonomous system number.
        /// </summary>
        [Input("asn")]
        public Input<string>? Asn { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Router ID.
        /// </summary>
        [Input("routerId")]
        public Input<string>? RouterId { get; set; }

        public BgpVrfState()
        {
        }
        public static new BgpVrfState Empty => new BgpVrfState();
    }
}
