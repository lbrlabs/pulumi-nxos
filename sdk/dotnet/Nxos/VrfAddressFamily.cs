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
    /// This resource can manage a VRF Address Family.
    /// 
    /// - API Documentation: [rtctrlDomAf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:DomAf/)
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
    ///     var example = new Nxos.VrfAddressFamily("example", new()
    ///     {
    ///         AddressFamily = "ipv4-ucast",
    ///         Vrf = "VRF1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import nxos:index/vrfAddressFamily:VrfAddressFamily example "sys/inst-[VRF1]/dom-[VRF1]/af-[ipv4-ucast]"
    /// ```
    /// </summary>
    [NxosResourceType("nxos:index/vrfAddressFamily:VrfAddressFamily")]
    public partial class VrfAddressFamily : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
        /// </summary>
        [Output("addressFamily")]
        public Output<string> AddressFamily { get; private set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// VRF name.
        /// </summary>
        [Output("vrf")]
        public Output<string> Vrf { get; private set; } = null!;


        /// <summary>
        /// Create a VrfAddressFamily resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VrfAddressFamily(string name, VrfAddressFamilyArgs args, CustomResourceOptions? options = null)
            : base("nxos:index/vrfAddressFamily:VrfAddressFamily", name, args ?? new VrfAddressFamilyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VrfAddressFamily(string name, Input<string> id, VrfAddressFamilyState? state = null, CustomResourceOptions? options = null)
            : base("nxos:index/vrfAddressFamily:VrfAddressFamily", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VrfAddressFamily resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VrfAddressFamily Get(string name, Input<string> id, VrfAddressFamilyState? state = null, CustomResourceOptions? options = null)
        {
            return new VrfAddressFamily(name, id, state, options);
        }
    }

    public sealed class VrfAddressFamilyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
        /// </summary>
        [Input("addressFamily", required: true)]
        public Input<string> AddressFamily { get; set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf", required: true)]
        public Input<string> Vrf { get; set; } = null!;

        public VrfAddressFamilyArgs()
        {
        }
        public static new VrfAddressFamilyArgs Empty => new VrfAddressFamilyArgs();
    }

    public sealed class VrfAddressFamilyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
        /// </summary>
        [Input("addressFamily")]
        public Input<string>? AddressFamily { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// VRF name.
        /// </summary>
        [Input("vrf")]
        public Input<string>? Vrf { get; set; }

        public VrfAddressFamilyState()
        {
        }
        public static new VrfAddressFamilyState Empty => new VrfAddressFamilyState();
    }
}