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
    /// This resource can manage the default QoS policy map match class map set QoS group configuration.
    /// 
    /// - API Documentation: [ipqosSetQoSGrp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:SetQoSGrp/)
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
    ///     var example = new Nxos.DefaultQosPolicyMapMatchClassMapSetQosGroup("example", new()
    ///     {
    ///         ClassMapName = "Voice",
    ///         PolicyMapName = "PM1",
    ///         QosGroupId = 1,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import nxos:index/defaultQosPolicyMapMatchClassMapSetQosGroup:DefaultQosPolicyMapMatchClassMapSetQosGroup example "sys/ipqos/dflt/p/name-[PM1]/cmap-[Voice]/setGrp"
    /// ```
    /// </summary>
    [NxosResourceType("nxos:index/defaultQosPolicyMapMatchClassMapSetQosGroup:DefaultQosPolicyMapMatchClassMapSetQosGroup")]
    public partial class DefaultQosPolicyMapMatchClassMapSetQosGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Class map name.
        /// </summary>
        [Output("classMapName")]
        public Output<string> ClassMapName { get; private set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Output("policyMapName")]
        public Output<string> PolicyMapName { get; private set; } = null!;

        /// <summary>
        /// QoS group ID. - Range: `0`-`7` - Default value: `0`
        /// </summary>
        [Output("qosGroupId")]
        public Output<int> QosGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a DefaultQosPolicyMapMatchClassMapSetQosGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DefaultQosPolicyMapMatchClassMapSetQosGroup(string name, DefaultQosPolicyMapMatchClassMapSetQosGroupArgs args, CustomResourceOptions? options = null)
            : base("nxos:index/defaultQosPolicyMapMatchClassMapSetQosGroup:DefaultQosPolicyMapMatchClassMapSetQosGroup", name, args ?? new DefaultQosPolicyMapMatchClassMapSetQosGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DefaultQosPolicyMapMatchClassMapSetQosGroup(string name, Input<string> id, DefaultQosPolicyMapMatchClassMapSetQosGroupState? state = null, CustomResourceOptions? options = null)
            : base("nxos:index/defaultQosPolicyMapMatchClassMapSetQosGroup:DefaultQosPolicyMapMatchClassMapSetQosGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DefaultQosPolicyMapMatchClassMapSetQosGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DefaultQosPolicyMapMatchClassMapSetQosGroup Get(string name, Input<string> id, DefaultQosPolicyMapMatchClassMapSetQosGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new DefaultQosPolicyMapMatchClassMapSetQosGroup(name, id, state, options);
        }
    }

    public sealed class DefaultQosPolicyMapMatchClassMapSetQosGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Class map name.
        /// </summary>
        [Input("classMapName", required: true)]
        public Input<string> ClassMapName { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Input("policyMapName", required: true)]
        public Input<string> PolicyMapName { get; set; } = null!;

        /// <summary>
        /// QoS group ID. - Range: `0`-`7` - Default value: `0`
        /// </summary>
        [Input("qosGroupId")]
        public Input<int>? QosGroupId { get; set; }

        public DefaultQosPolicyMapMatchClassMapSetQosGroupArgs()
        {
        }
        public static new DefaultQosPolicyMapMatchClassMapSetQosGroupArgs Empty => new DefaultQosPolicyMapMatchClassMapSetQosGroupArgs();
    }

    public sealed class DefaultQosPolicyMapMatchClassMapSetQosGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Class map name.
        /// </summary>
        [Input("classMapName")]
        public Input<string>? ClassMapName { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Policy map name.
        /// </summary>
        [Input("policyMapName")]
        public Input<string>? PolicyMapName { get; set; }

        /// <summary>
        /// QoS group ID. - Range: `0`-`7` - Default value: `0`
        /// </summary>
        [Input("qosGroupId")]
        public Input<int>? QosGroupId { get; set; }

        public DefaultQosPolicyMapMatchClassMapSetQosGroupState()
        {
        }
        public static new DefaultQosPolicyMapMatchClassMapSetQosGroupState Empty => new DefaultQosPolicyMapMatchClassMapSetQosGroupState();
    }
}
