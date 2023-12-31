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
    /// This resource can manage the BGP domain (VRF) graceful restart configuration.
    /// 
    /// - API Documentation: [bgpGr](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Gr/)
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
    ///     var example = new Nxos.BgpGracefulRestart("example", new()
    ///     {
    ///         Asn = "65001",
    ///         RestartInterval = 240,
    ///         StaleInterval = 1800,
    ///         Vrf = "default",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import nxos:index/bgpGracefulRestart:BgpGracefulRestart example "sys/bgp/inst/dom-[default]/gr"
    /// ```
    /// </summary>
    [NxosResourceType("nxos:index/bgpGracefulRestart:BgpGracefulRestart")]
    public partial class BgpGracefulRestart : global::Pulumi.CustomResource
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
        /// The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
        /// </summary>
        [Output("restartInterval")]
        public Output<int> RestartInterval { get; private set; } = null!;

        /// <summary>
        /// The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
        /// </summary>
        [Output("staleInterval")]
        public Output<int> StaleInterval { get; private set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Output("vrf")]
        public Output<string> Vrf { get; private set; } = null!;


        /// <summary>
        /// Create a BgpGracefulRestart resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BgpGracefulRestart(string name, BgpGracefulRestartArgs args, CustomResourceOptions? options = null)
            : base("nxos:index/bgpGracefulRestart:BgpGracefulRestart", name, args ?? new BgpGracefulRestartArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BgpGracefulRestart(string name, Input<string> id, BgpGracefulRestartState? state = null, CustomResourceOptions? options = null)
            : base("nxos:index/bgpGracefulRestart:BgpGracefulRestart", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BgpGracefulRestart resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BgpGracefulRestart Get(string name, Input<string> id, BgpGracefulRestartState? state = null, CustomResourceOptions? options = null)
        {
            return new BgpGracefulRestart(name, id, state, options);
        }
    }

    public sealed class BgpGracefulRestartArgs : global::Pulumi.ResourceArgs
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
        /// The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
        /// </summary>
        [Input("restartInterval")]
        public Input<int>? RestartInterval { get; set; }

        /// <summary>
        /// The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
        /// </summary>
        [Input("staleInterval")]
        public Input<int>? StaleInterval { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf", required: true)]
        public Input<string> Vrf { get; set; } = null!;

        public BgpGracefulRestartArgs()
        {
        }
        public static new BgpGracefulRestartArgs Empty => new BgpGracefulRestartArgs();
    }

    public sealed class BgpGracefulRestartState : global::Pulumi.ResourceArgs
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
        /// The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
        /// </summary>
        [Input("restartInterval")]
        public Input<int>? RestartInterval { get; set; }

        /// <summary>
        /// The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
        /// </summary>
        [Input("staleInterval")]
        public Input<int>? StaleInterval { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf")]
        public Input<string>? Vrf { get; set; }

        public BgpGracefulRestartState()
        {
        }
        public static new BgpGracefulRestartState Empty => new BgpGracefulRestartState();
    }
}
