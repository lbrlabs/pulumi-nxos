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
    [NxosResourceType("nxos:nxos/ospfVrf:OspfVrf")]
    public partial class OspfVrf : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        /// </summary>
        [Output("adminState")]
        public Output<string> AdminState { get; private set; } = null!;

        /// <summary>
        /// Bandwidth reference value. - Range: `0`-`4294967295` - Default value: `40000`
        /// </summary>
        [Output("bandwidthReference")]
        public Output<int> BandwidthReference { get; private set; } = null!;

        /// <summary>
        /// Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
        /// </summary>
        [Output("banwidthReferenceUnit")]
        public Output<string> BanwidthReferenceUnit { get; private set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Administrative distance preference. - Range: `1`-`255` - Default value: `110`
        /// </summary>
        [Output("distance")]
        public Output<int> Distance { get; private set; } = null!;

        /// <summary>
        /// OSPF instance name.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Router ID. - Default value: `0.0.0.0`
        /// </summary>
        [Output("routerId")]
        public Output<string> RouterId { get; private set; } = null!;


        /// <summary>
        /// Create a OspfVrf resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public OspfVrf(string name, OspfVrfArgs args, CustomResourceOptions? options = null)
            : base("nxos:nxos/ospfVrf:OspfVrf", name, args ?? new OspfVrfArgs(), MakeResourceOptions(options, ""))
        {
        }

        private OspfVrf(string name, Input<string> id, OspfVrfState? state = null, CustomResourceOptions? options = null)
            : base("nxos:nxos/ospfVrf:OspfVrf", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing OspfVrf resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static OspfVrf Get(string name, Input<string> id, OspfVrfState? state = null, CustomResourceOptions? options = null)
        {
            return new OspfVrf(name, id, state, options);
        }
    }

    public sealed class OspfVrfArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        /// </summary>
        [Input("adminState")]
        public Input<string>? AdminState { get; set; }

        /// <summary>
        /// Bandwidth reference value. - Range: `0`-`4294967295` - Default value: `40000`
        /// </summary>
        [Input("bandwidthReference")]
        public Input<int>? BandwidthReference { get; set; }

        /// <summary>
        /// Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
        /// </summary>
        [Input("banwidthReferenceUnit")]
        public Input<string>? BanwidthReferenceUnit { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Administrative distance preference. - Range: `1`-`255` - Default value: `110`
        /// </summary>
        [Input("distance")]
        public Input<int>? Distance { get; set; }

        /// <summary>
        /// OSPF instance name.
        /// </summary>
        [Input("instanceName", required: true)]
        public Input<string> InstanceName { get; set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Router ID. - Default value: `0.0.0.0`
        /// </summary>
        [Input("routerId")]
        public Input<string>? RouterId { get; set; }

        public OspfVrfArgs()
        {
        }
        public static new OspfVrfArgs Empty => new OspfVrfArgs();
    }

    public sealed class OspfVrfState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        /// </summary>
        [Input("adminState")]
        public Input<string>? AdminState { get; set; }

        /// <summary>
        /// Bandwidth reference value. - Range: `0`-`4294967295` - Default value: `40000`
        /// </summary>
        [Input("bandwidthReference")]
        public Input<int>? BandwidthReference { get; set; }

        /// <summary>
        /// Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
        /// </summary>
        [Input("banwidthReferenceUnit")]
        public Input<string>? BanwidthReferenceUnit { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Administrative distance preference. - Range: `1`-`255` - Default value: `110`
        /// </summary>
        [Input("distance")]
        public Input<int>? Distance { get; set; }

        /// <summary>
        /// OSPF instance name.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Router ID. - Default value: `0.0.0.0`
        /// </summary>
        [Input("routerId")]
        public Input<string>? RouterId { get; set; }

        public OspfVrfState()
        {
        }
        public static new OspfVrfState Empty => new OspfVrfState();
    }
}
