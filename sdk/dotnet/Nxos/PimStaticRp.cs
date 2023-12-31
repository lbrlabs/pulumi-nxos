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
    /// This resource can manage the PIM Static RP configuration.
    /// 
    /// - API Documentation: [pimStaticRP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:StaticRP/)
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
    ///     var example = new Nxos.PimStaticRp("example", new()
    ///     {
    ///         Address = "1.2.3.4",
    ///         VrfName = "default",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import nxos:index/pimStaticRp:PimStaticRp example "sys/pim/inst/dom-[default]/staticrp/rp-[1.2.3.4]"
    /// ```
    /// </summary>
    [NxosResourceType("nxos:index/pimStaticRp:PimStaticRp")]
    public partial class PimStaticRp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Address.
        /// </summary>
        [Output("address")]
        public Output<string> Address { get; private set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Output("vrfName")]
        public Output<string> VrfName { get; private set; } = null!;


        /// <summary>
        /// Create a PimStaticRp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PimStaticRp(string name, PimStaticRpArgs args, CustomResourceOptions? options = null)
            : base("nxos:index/pimStaticRp:PimStaticRp", name, args ?? new PimStaticRpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PimStaticRp(string name, Input<string> id, PimStaticRpState? state = null, CustomResourceOptions? options = null)
            : base("nxos:index/pimStaticRp:PimStaticRp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PimStaticRp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PimStaticRp Get(string name, Input<string> id, PimStaticRpState? state = null, CustomResourceOptions? options = null)
        {
            return new PimStaticRp(name, id, state, options);
        }
    }

    public sealed class PimStaticRpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address.
        /// </summary>
        [Input("address", required: true)]
        public Input<string> Address { get; set; } = null!;

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

        public PimStaticRpArgs()
        {
        }
        public static new PimStaticRpArgs Empty => new PimStaticRpArgs();
    }

    public sealed class PimStaticRpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address.
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrfName")]
        public Input<string>? VrfName { get; set; }

        public PimStaticRpState()
        {
        }
        public static new PimStaticRpState Empty => new PimStaticRpState();
    }
}
