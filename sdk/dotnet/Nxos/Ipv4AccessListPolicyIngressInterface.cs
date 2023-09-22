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
    /// This resource can manage an IPv4 Access List Policy Ingress Interface.
    /// 
    /// - API Documentation: [aclIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Security%20and%20Policing/acl:If/)
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
    ///     var example = new Nxos.Ipv4AccessListPolicyIngressInterface("example", new()
    ///     {
    ///         AccessListName = "ACL1",
    ///         InterfaceId = "eth1/10",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import nxos:index/ipv4AccessListPolicyIngressInterface:Ipv4AccessListPolicyIngressInterface example "sys/acl/ipv4/policy/ingress/intf-[eth1/10]"
    /// ```
    /// </summary>
    [NxosResourceType("nxos:index/ipv4AccessListPolicyIngressInterface:Ipv4AccessListPolicyIngressInterface")]
    public partial class Ipv4AccessListPolicyIngressInterface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Access list name.
        /// </summary>
        [Output("accessListName")]
        public Output<string?> AccessListName { get; private set; } = null!;

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
        /// Create a Ipv4AccessListPolicyIngressInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipv4AccessListPolicyIngressInterface(string name, Ipv4AccessListPolicyIngressInterfaceArgs args, CustomResourceOptions? options = null)
            : base("nxos:index/ipv4AccessListPolicyIngressInterface:Ipv4AccessListPolicyIngressInterface", name, args ?? new Ipv4AccessListPolicyIngressInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipv4AccessListPolicyIngressInterface(string name, Input<string> id, Ipv4AccessListPolicyIngressInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("nxos:index/ipv4AccessListPolicyIngressInterface:Ipv4AccessListPolicyIngressInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipv4AccessListPolicyIngressInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipv4AccessListPolicyIngressInterface Get(string name, Input<string> id, Ipv4AccessListPolicyIngressInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipv4AccessListPolicyIngressInterface(name, id, state, options);
        }
    }

    public sealed class Ipv4AccessListPolicyIngressInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access list name.
        /// </summary>
        [Input("accessListName")]
        public Input<string>? AccessListName { get; set; }

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

        public Ipv4AccessListPolicyIngressInterfaceArgs()
        {
        }
        public static new Ipv4AccessListPolicyIngressInterfaceArgs Empty => new Ipv4AccessListPolicyIngressInterfaceArgs();
    }

    public sealed class Ipv4AccessListPolicyIngressInterfaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Access list name.
        /// </summary>
        [Input("accessListName")]
        public Input<string>? AccessListName { get; set; }

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

        public Ipv4AccessListPolicyIngressInterfaceState()
        {
        }
        public static new Ipv4AccessListPolicyIngressInterfaceState Empty => new Ipv4AccessListPolicyIngressInterfaceState();
    }
}