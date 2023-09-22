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
    /// This resource can manage the PTP feature configuration.
    /// 
    /// - API Documentation: [fmPtp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Ptp/)
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
    ///     var example = new Nxos.FeaturePtp("example", new()
    ///     {
    ///         AdminState = "enabled",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import nxos:index/featurePtp:FeaturePtp example "sys/fm/ptp"
    /// ```
    /// </summary>
    [NxosResourceType("nxos:index/featurePtp:FeaturePtp")]
    public partial class FeaturePtp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Administrative state. - Choices: `enabled`, `disabled`
        /// </summary>
        [Output("adminState")]
        public Output<string> AdminState { get; private set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;


        /// <summary>
        /// Create a FeaturePtp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FeaturePtp(string name, FeaturePtpArgs args, CustomResourceOptions? options = null)
            : base("nxos:index/featurePtp:FeaturePtp", name, args ?? new FeaturePtpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FeaturePtp(string name, Input<string> id, FeaturePtpState? state = null, CustomResourceOptions? options = null)
            : base("nxos:index/featurePtp:FeaturePtp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FeaturePtp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FeaturePtp Get(string name, Input<string> id, FeaturePtpState? state = null, CustomResourceOptions? options = null)
        {
            return new FeaturePtp(name, id, state, options);
        }
    }

    public sealed class FeaturePtpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative state. - Choices: `enabled`, `disabled`
        /// </summary>
        [Input("adminState", required: true)]
        public Input<string> AdminState { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        public FeaturePtpArgs()
        {
        }
        public static new FeaturePtpArgs Empty => new FeaturePtpArgs();
    }

    public sealed class FeaturePtpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative state. - Choices: `enabled`, `disabled`
        /// </summary>
        [Input("adminState")]
        public Input<string>? AdminState { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        public FeaturePtpState()
        {
        }
        public static new FeaturePtpState Empty => new FeaturePtpState();
    }
}