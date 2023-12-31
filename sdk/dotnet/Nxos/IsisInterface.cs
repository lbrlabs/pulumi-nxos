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
    /// This resource can manage the IS-IS interface configuration.
    /// 
    /// - API Documentation: [isisInternalIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:InternalIf/)
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
    ///     var example = new Nxos.IsisInterface("example", new()
    ///     {
    ///         AuthenticationCheck = false,
    ///         AuthenticationCheckL1 = false,
    ///         AuthenticationCheckL2 = false,
    ///         AuthenticationKey = "",
    ///         AuthenticationKeyL1 = "",
    ///         AuthenticationKeyL2 = "",
    ///         AuthenticationType = "unknown",
    ///         AuthenticationTypeL1 = "unknown",
    ///         AuthenticationTypeL2 = "unknown",
    ///         CircuitType = "l2",
    ///         HelloInterval = 20,
    ///         HelloIntervalL1 = 20,
    ///         HelloIntervalL2 = 20,
    ///         HelloMultiplier = 4,
    ///         HelloMultiplierL1 = 4,
    ///         HelloMultiplierL2 = 4,
    ///         HelloPadding = "never",
    ///         InterfaceId = "eth1/10",
    ///         MetricL1 = 1000,
    ///         MetricL2 = 1000,
    ///         MtuCheck = true,
    ///         MtuCheckL1 = true,
    ///         MtuCheckL2 = true,
    ///         NetworkTypeP2p = "on",
    ///         Passive = "l1",
    ///         PriorityL1 = 80,
    ///         PriorityL2 = 80,
    ///         Vrf = "default",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    ///  $ pulumi import nxos:index/isisInterface:IsisInterface example "sys/isis/if-[eth1/10]"
    /// ```
    /// </summary>
    [NxosResourceType("nxos:index/isisInterface:IsisInterface")]
    public partial class IsisInterface : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Authentication Check for ISIS without specific level. - Default value: `true`
        /// </summary>
        [Output("authenticationCheck")]
        public Output<bool> AuthenticationCheck { get; private set; } = null!;

        /// <summary>
        /// Authentication Check for ISIS on Level-1. - Default value: `true`
        /// </summary>
        [Output("authenticationCheckL1")]
        public Output<bool> AuthenticationCheckL1 { get; private set; } = null!;

        /// <summary>
        /// Authentication Check for ISIS on Level-2. - Default value: `true`
        /// </summary>
        [Output("authenticationCheckL2")]
        public Output<bool> AuthenticationCheckL2 { get; private set; } = null!;

        /// <summary>
        /// Authentication Key for IS-IS without specific level.
        /// </summary>
        [Output("authenticationKey")]
        public Output<string?> AuthenticationKey { get; private set; } = null!;

        /// <summary>
        /// Authentication Key for IS-IS on Level-1.
        /// </summary>
        [Output("authenticationKeyL1")]
        public Output<string?> AuthenticationKeyL1 { get; private set; } = null!;

        /// <summary>
        /// Authentication Key for IS-IS on Level-2.
        /// </summary>
        [Output("authenticationKeyL2")]
        public Output<string?> AuthenticationKeyL2 { get; private set; } = null!;

        /// <summary>
        /// IS-IS Authentication-Type without specific level. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
        /// </summary>
        [Output("authenticationType")]
        public Output<string> AuthenticationType { get; private set; } = null!;

        /// <summary>
        /// IS-IS Authentication-Type for Level-1. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
        /// </summary>
        [Output("authenticationTypeL1")]
        public Output<string> AuthenticationTypeL1 { get; private set; } = null!;

        /// <summary>
        /// IS-IS Authentication-Type for Level-2. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
        /// </summary>
        [Output("authenticationTypeL2")]
        public Output<string> AuthenticationTypeL2 { get; private set; } = null!;

        /// <summary>
        /// Circuit type. - Choices: `l1`, `l2`, `l12` - Default value: `l12`
        /// </summary>
        [Output("circuitType")]
        public Output<string> CircuitType { get; private set; } = null!;

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Output("device")]
        public Output<string?> Device { get; private set; } = null!;

        /// <summary>
        /// Hello interval. - Range: `1`-`65535` - Default value: `10`
        /// </summary>
        [Output("helloInterval")]
        public Output<int> HelloInterval { get; private set; } = null!;

        /// <summary>
        /// Hello interval Level-1. - Range: `1`-`65535` - Default value: `10`
        /// </summary>
        [Output("helloIntervalL1")]
        public Output<int> HelloIntervalL1 { get; private set; } = null!;

        /// <summary>
        /// Hello interval Level-2. - Range: `1`-`65535` - Default value: `10`
        /// </summary>
        [Output("helloIntervalL2")]
        public Output<int> HelloIntervalL2 { get; private set; } = null!;

        /// <summary>
        /// Hello multiplier. - Range: `3`-`1000` - Default value: `3`
        /// </summary>
        [Output("helloMultiplier")]
        public Output<int> HelloMultiplier { get; private set; } = null!;

        /// <summary>
        /// Hello multiplier Level-1. - Range: `3`-`1000` - Default value: `3`
        /// </summary>
        [Output("helloMultiplierL1")]
        public Output<int> HelloMultiplierL1 { get; private set; } = null!;

        /// <summary>
        /// Hello multiplier Level-2. - Range: `3`-`1000` - Default value: `3`
        /// </summary>
        [Output("helloMultiplierL2")]
        public Output<int> HelloMultiplierL2 { get; private set; } = null!;

        /// <summary>
        /// Hello padding. - Choices: `always`, `transient`, `never` - Default value: `always`
        /// </summary>
        [Output("helloPadding")]
        public Output<string> HelloPadding { get; private set; } = null!;

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        [Output("interfaceId")]
        public Output<string> InterfaceId { get; private set; } = null!;

        /// <summary>
        /// Interface metric Level-1. - Range: `0`-`16777216` - Default value: `16777216`
        /// </summary>
        [Output("metricL1")]
        public Output<int> MetricL1 { get; private set; } = null!;

        /// <summary>
        /// Interface metric Level-2. - Range: `0`-`16777216` - Default value: `16777216`
        /// </summary>
        [Output("metricL2")]
        public Output<int> MetricL2 { get; private set; } = null!;

        /// <summary>
        /// MTU Check for IS-IS without specific level. - Default value: `false`
        /// </summary>
        [Output("mtuCheck")]
        public Output<bool> MtuCheck { get; private set; } = null!;

        /// <summary>
        /// MTU Check for IS-IS on Level-1. - Default value: `false`
        /// </summary>
        [Output("mtuCheckL1")]
        public Output<bool> MtuCheckL1 { get; private set; } = null!;

        /// <summary>
        /// MTU Check for IS-IS on Level-2. - Default value: `false`
        /// </summary>
        [Output("mtuCheckL2")]
        public Output<bool> MtuCheckL2 { get; private set; } = null!;

        /// <summary>
        /// Enabling Point-to-Point Network Type on IS-IS Interface. - Choices: `off`, `on`, `useAllISMac` - Default value: `off`
        /// </summary>
        [Output("networkTypeP2p")]
        public Output<string> NetworkTypeP2p { get; private set; } = null!;

        /// <summary>
        /// IS-IS Passive Interface Info. - Choices: `l1`, `l2`, `l12`, `noL1`, `noL2`, `noL12`, `inheritDef` - Default value:
        /// `inheritDef`
        /// </summary>
        [Output("passive")]
        public Output<string> Passive { get; private set; } = null!;

        /// <summary>
        /// Circuit priority. - Range: `0`-`127` - Default value: `64`
        /// </summary>
        [Output("priorityL1")]
        public Output<int> PriorityL1 { get; private set; } = null!;

        /// <summary>
        /// Circuit priority. - Range: `0`-`127` - Default value: `64`
        /// </summary>
        [Output("priorityL2")]
        public Output<int> PriorityL2 { get; private set; } = null!;

        /// <summary>
        /// VRF.
        /// </summary>
        [Output("vrf")]
        public Output<string?> Vrf { get; private set; } = null!;


        /// <summary>
        /// Create a IsisInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IsisInterface(string name, IsisInterfaceArgs args, CustomResourceOptions? options = null)
            : base("nxos:index/isisInterface:IsisInterface", name, args ?? new IsisInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IsisInterface(string name, Input<string> id, IsisInterfaceState? state = null, CustomResourceOptions? options = null)
            : base("nxos:index/isisInterface:IsisInterface", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IsisInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IsisInterface Get(string name, Input<string> id, IsisInterfaceState? state = null, CustomResourceOptions? options = null)
        {
            return new IsisInterface(name, id, state, options);
        }
    }

    public sealed class IsisInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication Check for ISIS without specific level. - Default value: `true`
        /// </summary>
        [Input("authenticationCheck")]
        public Input<bool>? AuthenticationCheck { get; set; }

        /// <summary>
        /// Authentication Check for ISIS on Level-1. - Default value: `true`
        /// </summary>
        [Input("authenticationCheckL1")]
        public Input<bool>? AuthenticationCheckL1 { get; set; }

        /// <summary>
        /// Authentication Check for ISIS on Level-2. - Default value: `true`
        /// </summary>
        [Input("authenticationCheckL2")]
        public Input<bool>? AuthenticationCheckL2 { get; set; }

        /// <summary>
        /// Authentication Key for IS-IS without specific level.
        /// </summary>
        [Input("authenticationKey")]
        public Input<string>? AuthenticationKey { get; set; }

        /// <summary>
        /// Authentication Key for IS-IS on Level-1.
        /// </summary>
        [Input("authenticationKeyL1")]
        public Input<string>? AuthenticationKeyL1 { get; set; }

        /// <summary>
        /// Authentication Key for IS-IS on Level-2.
        /// </summary>
        [Input("authenticationKeyL2")]
        public Input<string>? AuthenticationKeyL2 { get; set; }

        /// <summary>
        /// IS-IS Authentication-Type without specific level. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
        /// </summary>
        [Input("authenticationType")]
        public Input<string>? AuthenticationType { get; set; }

        /// <summary>
        /// IS-IS Authentication-Type for Level-1. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
        /// </summary>
        [Input("authenticationTypeL1")]
        public Input<string>? AuthenticationTypeL1 { get; set; }

        /// <summary>
        /// IS-IS Authentication-Type for Level-2. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
        /// </summary>
        [Input("authenticationTypeL2")]
        public Input<string>? AuthenticationTypeL2 { get; set; }

        /// <summary>
        /// Circuit type. - Choices: `l1`, `l2`, `l12` - Default value: `l12`
        /// </summary>
        [Input("circuitType")]
        public Input<string>? CircuitType { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Hello interval. - Range: `1`-`65535` - Default value: `10`
        /// </summary>
        [Input("helloInterval")]
        public Input<int>? HelloInterval { get; set; }

        /// <summary>
        /// Hello interval Level-1. - Range: `1`-`65535` - Default value: `10`
        /// </summary>
        [Input("helloIntervalL1")]
        public Input<int>? HelloIntervalL1 { get; set; }

        /// <summary>
        /// Hello interval Level-2. - Range: `1`-`65535` - Default value: `10`
        /// </summary>
        [Input("helloIntervalL2")]
        public Input<int>? HelloIntervalL2 { get; set; }

        /// <summary>
        /// Hello multiplier. - Range: `3`-`1000` - Default value: `3`
        /// </summary>
        [Input("helloMultiplier")]
        public Input<int>? HelloMultiplier { get; set; }

        /// <summary>
        /// Hello multiplier Level-1. - Range: `3`-`1000` - Default value: `3`
        /// </summary>
        [Input("helloMultiplierL1")]
        public Input<int>? HelloMultiplierL1 { get; set; }

        /// <summary>
        /// Hello multiplier Level-2. - Range: `3`-`1000` - Default value: `3`
        /// </summary>
        [Input("helloMultiplierL2")]
        public Input<int>? HelloMultiplierL2 { get; set; }

        /// <summary>
        /// Hello padding. - Choices: `always`, `transient`, `never` - Default value: `always`
        /// </summary>
        [Input("helloPadding")]
        public Input<string>? HelloPadding { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        [Input("interfaceId", required: true)]
        public Input<string> InterfaceId { get; set; } = null!;

        /// <summary>
        /// Interface metric Level-1. - Range: `0`-`16777216` - Default value: `16777216`
        /// </summary>
        [Input("metricL1")]
        public Input<int>? MetricL1 { get; set; }

        /// <summary>
        /// Interface metric Level-2. - Range: `0`-`16777216` - Default value: `16777216`
        /// </summary>
        [Input("metricL2")]
        public Input<int>? MetricL2 { get; set; }

        /// <summary>
        /// MTU Check for IS-IS without specific level. - Default value: `false`
        /// </summary>
        [Input("mtuCheck")]
        public Input<bool>? MtuCheck { get; set; }

        /// <summary>
        /// MTU Check for IS-IS on Level-1. - Default value: `false`
        /// </summary>
        [Input("mtuCheckL1")]
        public Input<bool>? MtuCheckL1 { get; set; }

        /// <summary>
        /// MTU Check for IS-IS on Level-2. - Default value: `false`
        /// </summary>
        [Input("mtuCheckL2")]
        public Input<bool>? MtuCheckL2 { get; set; }

        /// <summary>
        /// Enabling Point-to-Point Network Type on IS-IS Interface. - Choices: `off`, `on`, `useAllISMac` - Default value: `off`
        /// </summary>
        [Input("networkTypeP2p")]
        public Input<string>? NetworkTypeP2p { get; set; }

        /// <summary>
        /// IS-IS Passive Interface Info. - Choices: `l1`, `l2`, `l12`, `noL1`, `noL2`, `noL12`, `inheritDef` - Default value:
        /// `inheritDef`
        /// </summary>
        [Input("passive")]
        public Input<string>? Passive { get; set; }

        /// <summary>
        /// Circuit priority. - Range: `0`-`127` - Default value: `64`
        /// </summary>
        [Input("priorityL1")]
        public Input<int>? PriorityL1 { get; set; }

        /// <summary>
        /// Circuit priority. - Range: `0`-`127` - Default value: `64`
        /// </summary>
        [Input("priorityL2")]
        public Input<int>? PriorityL2 { get; set; }

        /// <summary>
        /// VRF.
        /// </summary>
        [Input("vrf")]
        public Input<string>? Vrf { get; set; }

        public IsisInterfaceArgs()
        {
        }
        public static new IsisInterfaceArgs Empty => new IsisInterfaceArgs();
    }

    public sealed class IsisInterfaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Authentication Check for ISIS without specific level. - Default value: `true`
        /// </summary>
        [Input("authenticationCheck")]
        public Input<bool>? AuthenticationCheck { get; set; }

        /// <summary>
        /// Authentication Check for ISIS on Level-1. - Default value: `true`
        /// </summary>
        [Input("authenticationCheckL1")]
        public Input<bool>? AuthenticationCheckL1 { get; set; }

        /// <summary>
        /// Authentication Check for ISIS on Level-2. - Default value: `true`
        /// </summary>
        [Input("authenticationCheckL2")]
        public Input<bool>? AuthenticationCheckL2 { get; set; }

        /// <summary>
        /// Authentication Key for IS-IS without specific level.
        /// </summary>
        [Input("authenticationKey")]
        public Input<string>? AuthenticationKey { get; set; }

        /// <summary>
        /// Authentication Key for IS-IS on Level-1.
        /// </summary>
        [Input("authenticationKeyL1")]
        public Input<string>? AuthenticationKeyL1 { get; set; }

        /// <summary>
        /// Authentication Key for IS-IS on Level-2.
        /// </summary>
        [Input("authenticationKeyL2")]
        public Input<string>? AuthenticationKeyL2 { get; set; }

        /// <summary>
        /// IS-IS Authentication-Type without specific level. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
        /// </summary>
        [Input("authenticationType")]
        public Input<string>? AuthenticationType { get; set; }

        /// <summary>
        /// IS-IS Authentication-Type for Level-1. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
        /// </summary>
        [Input("authenticationTypeL1")]
        public Input<string>? AuthenticationTypeL1 { get; set; }

        /// <summary>
        /// IS-IS Authentication-Type for Level-2. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
        /// </summary>
        [Input("authenticationTypeL2")]
        public Input<string>? AuthenticationTypeL2 { get; set; }

        /// <summary>
        /// Circuit type. - Choices: `l1`, `l2`, `l12` - Default value: `l12`
        /// </summary>
        [Input("circuitType")]
        public Input<string>? CircuitType { get; set; }

        /// <summary>
        /// A device name from the provider configuration.
        /// </summary>
        [Input("device")]
        public Input<string>? Device { get; set; }

        /// <summary>
        /// Hello interval. - Range: `1`-`65535` - Default value: `10`
        /// </summary>
        [Input("helloInterval")]
        public Input<int>? HelloInterval { get; set; }

        /// <summary>
        /// Hello interval Level-1. - Range: `1`-`65535` - Default value: `10`
        /// </summary>
        [Input("helloIntervalL1")]
        public Input<int>? HelloIntervalL1 { get; set; }

        /// <summary>
        /// Hello interval Level-2. - Range: `1`-`65535` - Default value: `10`
        /// </summary>
        [Input("helloIntervalL2")]
        public Input<int>? HelloIntervalL2 { get; set; }

        /// <summary>
        /// Hello multiplier. - Range: `3`-`1000` - Default value: `3`
        /// </summary>
        [Input("helloMultiplier")]
        public Input<int>? HelloMultiplier { get; set; }

        /// <summary>
        /// Hello multiplier Level-1. - Range: `3`-`1000` - Default value: `3`
        /// </summary>
        [Input("helloMultiplierL1")]
        public Input<int>? HelloMultiplierL1 { get; set; }

        /// <summary>
        /// Hello multiplier Level-2. - Range: `3`-`1000` - Default value: `3`
        /// </summary>
        [Input("helloMultiplierL2")]
        public Input<int>? HelloMultiplierL2 { get; set; }

        /// <summary>
        /// Hello padding. - Choices: `always`, `transient`, `never` - Default value: `always`
        /// </summary>
        [Input("helloPadding")]
        public Input<string>? HelloPadding { get; set; }

        /// <summary>
        /// Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        /// </summary>
        [Input("interfaceId")]
        public Input<string>? InterfaceId { get; set; }

        /// <summary>
        /// Interface metric Level-1. - Range: `0`-`16777216` - Default value: `16777216`
        /// </summary>
        [Input("metricL1")]
        public Input<int>? MetricL1 { get; set; }

        /// <summary>
        /// Interface metric Level-2. - Range: `0`-`16777216` - Default value: `16777216`
        /// </summary>
        [Input("metricL2")]
        public Input<int>? MetricL2 { get; set; }

        /// <summary>
        /// MTU Check for IS-IS without specific level. - Default value: `false`
        /// </summary>
        [Input("mtuCheck")]
        public Input<bool>? MtuCheck { get; set; }

        /// <summary>
        /// MTU Check for IS-IS on Level-1. - Default value: `false`
        /// </summary>
        [Input("mtuCheckL1")]
        public Input<bool>? MtuCheckL1 { get; set; }

        /// <summary>
        /// MTU Check for IS-IS on Level-2. - Default value: `false`
        /// </summary>
        [Input("mtuCheckL2")]
        public Input<bool>? MtuCheckL2 { get; set; }

        /// <summary>
        /// Enabling Point-to-Point Network Type on IS-IS Interface. - Choices: `off`, `on`, `useAllISMac` - Default value: `off`
        /// </summary>
        [Input("networkTypeP2p")]
        public Input<string>? NetworkTypeP2p { get; set; }

        /// <summary>
        /// IS-IS Passive Interface Info. - Choices: `l1`, `l2`, `l12`, `noL1`, `noL2`, `noL12`, `inheritDef` - Default value:
        /// `inheritDef`
        /// </summary>
        [Input("passive")]
        public Input<string>? Passive { get; set; }

        /// <summary>
        /// Circuit priority. - Range: `0`-`127` - Default value: `64`
        /// </summary>
        [Input("priorityL1")]
        public Input<int>? PriorityL1 { get; set; }

        /// <summary>
        /// Circuit priority. - Range: `0`-`127` - Default value: `64`
        /// </summary>
        [Input("priorityL2")]
        public Input<int>? PriorityL2 { get; set; }

        /// <summary>
        /// VRF.
        /// </summary>
        [Input("vrf")]
        public Input<string>? Vrf { get; set; }

        public IsisInterfaceState()
        {
        }
        public static new IsisInterfaceState Empty => new IsisInterfaceState();
    }
}
