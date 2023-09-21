// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Nxos.Nxos
{
    [NxosResourceType("nxos:nxos/pimVrf:PimVrf")]
    public partial class PimVrf : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        /// </summary>
        [Output("adminState")]
        public Output<string> AdminState { get; private set; } = null!;

        /// <summary>
        /// BFD. - Default value: `false`
        /// </summary>
        [Output("bfd")]
        public Output<bool> Bfd { get; private set; } = null!;

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
        /// Create a PimVrf resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PimVrf(string name, PimVrfArgs? args = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/pimVrf:PimVrf", name, args ?? new PimVrfArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PimVrf(string name, Input<string> id, PimVrfState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/pimVrf:PimVrf", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PimVrf resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PimVrf Get(string name, Input<string> id, PimVrfState? state = null, CustomResourceOptions? options = null)
        {
            return new PimVrf(name, id, state, options);
        }
    }

    public sealed class PimVrfArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        /// </summary>
        [Input("adminState")]
        public Input<string>? AdminState { get; set; }

        /// <summary>
        /// BFD. - Default value: `false`
        /// </summary>
        [Input("bfd")]
        public Input<bool>? Bfd { get; set; }

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

        public PimVrfArgs()
        {
        }
        public static new PimVrfArgs Empty => new PimVrfArgs();
    }

    public sealed class PimVrfState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        /// </summary>
        [Input("adminState")]
        public Input<string>? AdminState { get; set; }

        /// <summary>
        /// BFD. - Default value: `false`
        /// </summary>
        [Input("bfd")]
        public Input<bool>? Bfd { get; set; }

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

        public PimVrfState()
        {
        }
        public static new PimVrfState Empty => new PimVrfState();
    }
}
