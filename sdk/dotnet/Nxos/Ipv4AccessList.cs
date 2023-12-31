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
    /// This resource can manage IPv4 Access Lists.
    /// 
    /// - API Documentation: [ipv4aclACL](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Security%20and%20Policing/ipv4acl:ACL/)
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
    ///     var example = new Nxos.Ipv4AccessList("example");
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import nxos:index/ipv4AccessList:Ipv4AccessList example "sys/acl/ipv4/name-[ACL1]"
    /// ```
    /// </summary>
    [NxosResourceType("nxos:index/ipv4AccessList:Ipv4AccessList")]
    public partial class Ipv4AccessList : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Access list name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a Ipv4AccessList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Ipv4AccessList(string name, Ipv4AccessListArgs? args = null, CustomResourceOptions? options = null)
            : base("nxos:index/ipv4AccessList:Ipv4AccessList", name, args ?? new Ipv4AccessListArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Ipv4AccessList(string name, Input<string> id, Ipv4AccessListState? state = null, CustomResourceOptions? options = null)
            : base("nxos:index/ipv4AccessList:Ipv4AccessList", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Ipv4AccessList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Ipv4AccessList Get(string name, Input<string> id, Ipv4AccessListState? state = null, CustomResourceOptions? options = null)
        {
            return new Ipv4AccessList(name, id, state, options);
        }
    }

    public sealed class Ipv4AccessListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Access list name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Ipv4AccessListArgs()
        {
        }
        public static new Ipv4AccessListArgs Empty => new Ipv4AccessListArgs();
    }

    public sealed class Ipv4AccessListState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Access list name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public Ipv4AccessListState()
        {
        }
        public static new Ipv4AccessListState Empty => new Ipv4AccessListState();
    }
}
