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
    /// This resource can manage the PIM Anycast RP configuration.
    /// 
    /// - API Documentation: [pimAcastRPFuncP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:AcastRPFuncP/)
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
    ///     var example = new Nxos.PimAnycastRp("example", new()
    ///     {
    ///         LocalInterface = "eth1/10",
    ///         SourceInterface = "eth1/10",
    ///         VrfName = "default",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import nxos:index/pimAnycastRp:PimAnycastRp example "sys/pim/inst/dom-[default]/acastrpfunc"
    /// ```
    /// </summary>
    [NxosResourceType("nxos:index/pimAnycastRp:PimAnycastRp")]
    public partial class PimAnycastRp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        [Output("localInterface")]
        public Output<string?> LocalInterface { get; private set; } = null!;

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        [Output("sourceInterface")]
        public Output<string?> SourceInterface { get; private set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Output("vrfName")]
        public Output<string> VrfName { get; private set; } = null!;


        /// <summary>
        /// Create a PimAnycastRp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PimAnycastRp(string name, PimAnycastRpArgs args, CustomResourceOptions? options = null)
            : base("nxos:index/pimAnycastRp:PimAnycastRp", name, args ?? new PimAnycastRpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PimAnycastRp(string name, Input<string> id, PimAnycastRpState? state = null, CustomResourceOptions? options = null)
            : base("nxos:index/pimAnycastRp:PimAnycastRp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PimAnycastRp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PimAnycastRp Get(string name, Input<string> id, PimAnycastRpState? state = null, CustomResourceOptions? options = null)
        {
            return new PimAnycastRp(name, id, state, options);
        }
    }

    public sealed class PimAnycastRpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        [Input("localInterface")]
        public Input<string>? LocalInterface { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        [Input("sourceInterface")]
        public Input<string>? SourceInterface { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrfName", required: true)]
        public Input<string> VrfName { get; set; } = null!;

        public PimAnycastRpArgs()
        {
        }
        public static new PimAnycastRpArgs Empty => new PimAnycastRpArgs();
    }

    public sealed class PimAnycastRpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        [Input("localInterface")]
        public Input<string>? LocalInterface { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        [Input("sourceInterface")]
        public Input<string>? SourceInterface { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrfName")]
        public Input<string>? VrfName { get; set; }

        public PimAnycastRpState()
        {
        }
        public static new PimAnycastRpState Empty => new PimAnycastRpState();
    }
}
