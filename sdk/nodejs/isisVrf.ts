// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage the IS-IS VRF configuration.
 *
 * - API Documentation: [isisDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:Dom/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.IsisVrf("example", {
 *     adminState: "enabled",
 *     authenticationCheckL1: false,
 *     authenticationCheckL2: false,
 *     authenticationKeyL1: "",
 *     authenticationKeyL2: "",
 *     authenticationTypeL1: "unknown",
 *     authenticationTypeL2: "unknown",
 *     bandwidthReference: 400000,
 *     banwidthReferenceUnit: "mbps",
 *     instanceName: "ISIS1",
 *     isType: "l2",
 *     metricType: "wide",
 *     mtu: 2000,
 *     net: "49.0001.0000.0000.3333.00",
 *     passiveDefault: "l12",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/isisVrf:IsisVrf example "sys/isis/inst-[ISIS1]/dom-[default]"
 * ```
 */
export class IsisVrf extends pulumi.CustomResource {
    /**
     * Get an existing IsisVrf resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IsisVrfState, opts?: pulumi.CustomResourceOptions): IsisVrf {
        return new IsisVrf(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/isisVrf:IsisVrf';

    /**
     * Returns true if the given object is an instance of IsisVrf.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IsisVrf {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IsisVrf.__pulumiType;
    }

    /**
     * Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     */
    public readonly adminState!: pulumi.Output<string>;
    /**
     * Authentication Check for ISIS on Level-1. - Default value: `true`
     */
    public readonly authenticationCheckL1!: pulumi.Output<boolean>;
    /**
     * Authentication Check for ISIS on Level-2. - Default value: `true`
     */
    public readonly authenticationCheckL2!: pulumi.Output<boolean>;
    /**
     * Authentication Key for IS-IS on Level-1.
     */
    public readonly authenticationKeyL1!: pulumi.Output<string | undefined>;
    /**
     * Authentication Key for IS-IS on Level-2.
     */
    public readonly authenticationKeyL2!: pulumi.Output<string | undefined>;
    /**
     * IS-IS Authentication-Type for Level-1. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
     */
    public readonly authenticationTypeL1!: pulumi.Output<string>;
    /**
     * IS-IS Authentication-Type for Level-2. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
     */
    public readonly authenticationTypeL2!: pulumi.Output<string>;
    /**
     * The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost
     * metric. - Range: `0`-`4294967295` - Default value: `40000`
     */
    public readonly bandwidthReference!: pulumi.Output<number>;
    /**
     * Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
     */
    public readonly banwidthReferenceUnit!: pulumi.Output<string>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * IS-IS instance name.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * IS-IS domain type. - Choices: `l1`, `l2`, `l12` - Default value: `l12`
     */
    public readonly isType!: pulumi.Output<string>;
    /**
     * IS-IS metric type. - Choices: `narrow`, `wide`, `transition` - Default value: `wide`
     */
    public readonly metricType!: pulumi.Output<string>;
    /**
     * The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352
     * bytes. - Range: `256`-`4352` - Default value: `1492`
     */
    public readonly mtu!: pulumi.Output<number>;
    /**
     * VRF name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Holds IS-IS domain NET (address) value.
     */
    public readonly net!: pulumi.Output<string | undefined>;
    /**
     * IS-IS Domain passive-interface default level. - Choices: `l1`, `l2`, `l12`, `unknown` - Default value: `unknown`
     */
    public readonly passiveDefault!: pulumi.Output<string>;

    /**
     * Create a IsisVrf resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IsisVrfArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IsisVrfArgs | IsisVrfState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IsisVrfState | undefined;
            resourceInputs["adminState"] = state ? state.adminState : undefined;
            resourceInputs["authenticationCheckL1"] = state ? state.authenticationCheckL1 : undefined;
            resourceInputs["authenticationCheckL2"] = state ? state.authenticationCheckL2 : undefined;
            resourceInputs["authenticationKeyL1"] = state ? state.authenticationKeyL1 : undefined;
            resourceInputs["authenticationKeyL2"] = state ? state.authenticationKeyL2 : undefined;
            resourceInputs["authenticationTypeL1"] = state ? state.authenticationTypeL1 : undefined;
            resourceInputs["authenticationTypeL2"] = state ? state.authenticationTypeL2 : undefined;
            resourceInputs["bandwidthReference"] = state ? state.bandwidthReference : undefined;
            resourceInputs["banwidthReferenceUnit"] = state ? state.banwidthReferenceUnit : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["isType"] = state ? state.isType : undefined;
            resourceInputs["metricType"] = state ? state.metricType : undefined;
            resourceInputs["mtu"] = state ? state.mtu : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["net"] = state ? state.net : undefined;
            resourceInputs["passiveDefault"] = state ? state.passiveDefault : undefined;
        } else {
            const args = argsOrState as IsisVrfArgs | undefined;
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            resourceInputs["adminState"] = args ? args.adminState : undefined;
            resourceInputs["authenticationCheckL1"] = args ? args.authenticationCheckL1 : undefined;
            resourceInputs["authenticationCheckL2"] = args ? args.authenticationCheckL2 : undefined;
            resourceInputs["authenticationKeyL1"] = args ? args.authenticationKeyL1 : undefined;
            resourceInputs["authenticationKeyL2"] = args ? args.authenticationKeyL2 : undefined;
            resourceInputs["authenticationTypeL1"] = args ? args.authenticationTypeL1 : undefined;
            resourceInputs["authenticationTypeL2"] = args ? args.authenticationTypeL2 : undefined;
            resourceInputs["bandwidthReference"] = args ? args.bandwidthReference : undefined;
            resourceInputs["banwidthReferenceUnit"] = args ? args.banwidthReferenceUnit : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["isType"] = args ? args.isType : undefined;
            resourceInputs["metricType"] = args ? args.metricType : undefined;
            resourceInputs["mtu"] = args ? args.mtu : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["net"] = args ? args.net : undefined;
            resourceInputs["passiveDefault"] = args ? args.passiveDefault : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IsisVrf.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IsisVrf resources.
 */
export interface IsisVrfState {
    /**
     * Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     */
    adminState?: pulumi.Input<string>;
    /**
     * Authentication Check for ISIS on Level-1. - Default value: `true`
     */
    authenticationCheckL1?: pulumi.Input<boolean>;
    /**
     * Authentication Check for ISIS on Level-2. - Default value: `true`
     */
    authenticationCheckL2?: pulumi.Input<boolean>;
    /**
     * Authentication Key for IS-IS on Level-1.
     */
    authenticationKeyL1?: pulumi.Input<string>;
    /**
     * Authentication Key for IS-IS on Level-2.
     */
    authenticationKeyL2?: pulumi.Input<string>;
    /**
     * IS-IS Authentication-Type for Level-1. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
     */
    authenticationTypeL1?: pulumi.Input<string>;
    /**
     * IS-IS Authentication-Type for Level-2. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
     */
    authenticationTypeL2?: pulumi.Input<string>;
    /**
     * The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost
     * metric. - Range: `0`-`4294967295` - Default value: `40000`
     */
    bandwidthReference?: pulumi.Input<number>;
    /**
     * Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
     */
    banwidthReferenceUnit?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * IS-IS instance name.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * IS-IS domain type. - Choices: `l1`, `l2`, `l12` - Default value: `l12`
     */
    isType?: pulumi.Input<string>;
    /**
     * IS-IS metric type. - Choices: `narrow`, `wide`, `transition` - Default value: `wide`
     */
    metricType?: pulumi.Input<string>;
    /**
     * The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352
     * bytes. - Range: `256`-`4352` - Default value: `1492`
     */
    mtu?: pulumi.Input<number>;
    /**
     * VRF name.
     */
    name?: pulumi.Input<string>;
    /**
     * Holds IS-IS domain NET (address) value.
     */
    net?: pulumi.Input<string>;
    /**
     * IS-IS Domain passive-interface default level. - Choices: `l1`, `l2`, `l12`, `unknown` - Default value: `unknown`
     */
    passiveDefault?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IsisVrf resource.
 */
export interface IsisVrfArgs {
    /**
     * Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     */
    adminState?: pulumi.Input<string>;
    /**
     * Authentication Check for ISIS on Level-1. - Default value: `true`
     */
    authenticationCheckL1?: pulumi.Input<boolean>;
    /**
     * Authentication Check for ISIS on Level-2. - Default value: `true`
     */
    authenticationCheckL2?: pulumi.Input<boolean>;
    /**
     * Authentication Key for IS-IS on Level-1.
     */
    authenticationKeyL1?: pulumi.Input<string>;
    /**
     * Authentication Key for IS-IS on Level-2.
     */
    authenticationKeyL2?: pulumi.Input<string>;
    /**
     * IS-IS Authentication-Type for Level-1. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
     */
    authenticationTypeL1?: pulumi.Input<string>;
    /**
     * IS-IS Authentication-Type for Level-2. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
     */
    authenticationTypeL2?: pulumi.Input<string>;
    /**
     * The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost
     * metric. - Range: `0`-`4294967295` - Default value: `40000`
     */
    bandwidthReference?: pulumi.Input<number>;
    /**
     * Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
     */
    banwidthReferenceUnit?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * IS-IS instance name.
     */
    instanceName: pulumi.Input<string>;
    /**
     * IS-IS domain type. - Choices: `l1`, `l2`, `l12` - Default value: `l12`
     */
    isType?: pulumi.Input<string>;
    /**
     * IS-IS metric type. - Choices: `narrow`, `wide`, `transition` - Default value: `wide`
     */
    metricType?: pulumi.Input<string>;
    /**
     * The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352
     * bytes. - Range: `256`-`4352` - Default value: `1492`
     */
    mtu?: pulumi.Input<number>;
    /**
     * VRF name.
     */
    name?: pulumi.Input<string>;
    /**
     * Holds IS-IS domain NET (address) value.
     */
    net?: pulumi.Input<string>;
    /**
     * IS-IS Domain passive-interface default level. - Choices: `l1`, `l2`, `l12`, `unknown` - Default value: `unknown`
     */
    passiveDefault?: pulumi.Input<string>;
}
