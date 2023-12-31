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
    /// This resource can manage the default QoS policy interface in configuration.
    /// 
    /// - API Documentation: [ipqosIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:If/)
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
    ///     var example = new Nxos.DefaultQosPolicyInterfaceIn("example", new()
    ///     {
    ///         InterfaceId = "eth1/10",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import nxos:index/defaultQosPolicyInterfaceIn:DefaultQosPolicyInterfaceIn example "sys/ipqos/dflt/policy/in/intf-[eth1/10]"
    /// ```
    /// </summary>
    [NxosResourceType("nxos:index/defaultQosPolicyInterfaceIn:DefaultQosPolicyInterfaceIn")]
    public partial class DefaultQosPolicyInterfaceIn : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        [Output("interfaceId")]
        public Output<string> InterfaceId { get; private set; } = null!;


        /// <summary>
        /// Create a DefaultQosPolicyInterfaceIn resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DefaultQosPolicyInterfaceIn(string name, DefaultQosPolicyInterfaceInArgs args, CustomResourceOptions? options = null)
            : base("nxos:index/defaultQosPolicyInterfaceIn:DefaultQosPolicyInterfaceIn", name, args ?? new DefaultQosPolicyInterfaceInArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DefaultQosPolicyInterfaceIn(string name, Input<string> id, DefaultQosPolicyInterfaceInState? state = null, CustomResourceOptions? options = null)
            : base("nxos:index/defaultQosPolicyInterfaceIn:DefaultQosPolicyInterfaceIn", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DefaultQosPolicyInterfaceIn resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DefaultQosPolicyInterfaceIn Get(string name, Input<string> id, DefaultQosPolicyInterfaceInState? state = null, CustomResourceOptions? options = null)
        {
            return new DefaultQosPolicyInterfaceIn(name, id, state, options);
        }
    }

    public sealed class DefaultQosPolicyInterfaceInArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        public DefaultQosPolicyInterfaceInArgs()
        {
        }
        public static new DefaultQosPolicyInterfaceInArgs Empty => new DefaultQosPolicyInterfaceInArgs();
    }

    public sealed class DefaultQosPolicyInterfaceInState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        [Input("interfaceId")]
        public Input<string>? InterfaceId { get; set; }

        public DefaultQosPolicyInterfaceInState()
        {
        }
        public static new DefaultQosPolicyInterfaceInState Empty => new DefaultQosPolicyInterfaceInState();
    }
}
