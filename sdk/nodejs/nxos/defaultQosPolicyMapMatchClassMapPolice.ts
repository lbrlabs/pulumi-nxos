// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class DefaultQosPolicyMapMatchClassMapPolice extends pulumi.CustomResource {
    /**
     * Get an existing DefaultQosPolicyMapMatchClassMapPolice resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefaultQosPolicyMapMatchClassMapPoliceState, opts?: pulumi.CustomResourceOptions): DefaultQosPolicyMapMatchClassMapPolice {
        return new DefaultQosPolicyMapMatchClassMapPolice(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:nxos/defaultQosPolicyMapMatchClassMapPolice:DefaultQosPolicyMapMatchClassMapPolice';

    /**
     * Returns true if the given object is an instance of DefaultQosPolicyMapMatchClassMapPolice.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefaultQosPolicyMapMatchClassMapPolice {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefaultQosPolicyMapMatchClassMapPolice.__pulumiType;
    }

    /**
     * CIR burst rate. - Range: `0`-`536870912` - Default value: `200`
     */
    public readonly bcRate!: pulumi.Output<number>;
    /**
     * CIR burst rate unit. - Choices: `unspecified`, `bytes`, `kbytes`, `mbytes`, `ms`, `us`, `packets` - Default value: `ms`
     */
    public readonly bcUnit!: pulumi.Output<string>;
    /**
     * PIR burst rate. - Range: `0`-`536870912` - Default value: `0`
     */
    public readonly beRate!: pulumi.Output<number>;
    /**
     * PIR burst rate unit. - Choices: `unspecified`, `bytes`, `kbytes`, `mbytes`, `ms`, `us`, `packets` - Default value:
     * `unspecified`
     */
    public readonly beUnit!: pulumi.Output<string>;
    /**
     * CIR rate. - Range: `0`-`100000000000` - Default value: `0`
     */
    public readonly cirRate!: pulumi.Output<number>;
    /**
     * CIR rate unit. - Choices: `unspecified`, `bps`, `kbps`, `mbps`, `gbps`, `pps`, `pct` - Default value: `bps`
     */
    public readonly cirUnit!: pulumi.Output<string>;
    /**
     * Class map name.
     */
    public readonly classMapName!: pulumi.Output<string>;
    /**
     * Conform action. - Choices: `unspecified`, `transmit`, `drop`, `set-cos-transmit`, `set-dscp-transmit`,
     * `set-prec-transmit`, `set-qos-transmit` - Default value: `transmit`
     */
    public readonly conformAction!: pulumi.Output<string>;
    /**
     * Set CoS for conforming traffic. - Range: `0`-`7` - Default value: `0`
     */
    public readonly conformSetCos!: pulumi.Output<number>;
    /**
     * Set DSCP for conforming traffic. - Range: `0`-`63` - Default value: `0`
     */
    public readonly conformSetDscp!: pulumi.Output<number>;
    /**
     * Set precedence for conforming traffic. - Choices: `routine`, `priority`, `immediate`, `flash`, `flash-override`,
     * `critical`, `internet`, `network` - Default value: `routine`
     */
    public readonly conformSetPrecedence!: pulumi.Output<string>;
    /**
     * Set qos-group for conforming traffic. - Range: `0`-`7` - Default value: `0`
     */
    public readonly conformSetQosGroup!: pulumi.Output<number>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Exceed action. - Choices: `unspecified`, `transmit`, `drop`, `set-cos-transmit`, `set-dscp-transmit`,
     * `set-prec-transmit`, `set-qos-transmit` - Default value: `unspecified`
     */
    public readonly exceedAction!: pulumi.Output<string>;
    /**
     * Set CoS for exceeding traffic. - Range: `0`-`7` - Default value: `0`
     */
    public readonly exceedSetCos!: pulumi.Output<number>;
    /**
     * Set DSCP for exceeding traffic. - Range: `0`-`63` - Default value: `0`
     */
    public readonly exceedSetDscp!: pulumi.Output<number>;
    /**
     * Set precedence for exceeding traffic. - Choices: `routine`, `priority`, `immediate`, `flash`, `flash-override`,
     * `critical`, `internet`, `network` - Default value: `routine`
     */
    public readonly exceedSetPrecedence!: pulumi.Output<string>;
    /**
     * Set qos-group for exceeding traffic. - Range: `0`-`7` - Default value: `0`
     */
    public readonly exceedSetQosGroup!: pulumi.Output<number>;
    /**
     * PIR rate. - Range: `0`-`100000000000` - Default value: `0`
     */
    public readonly pirRate!: pulumi.Output<number>;
    /**
     * PIR rate unit. - Choices: `unspecified`, `bps`, `kbps`, `mbps`, `gbps`, `pps`, `pct` - Default value: `unspecified`
     */
    public readonly pirUnit!: pulumi.Output<string>;
    /**
     * Policy map name.
     */
    public readonly policyMapName!: pulumi.Output<string>;
    /**
     * Violate action. - Choices: `unspecified`, `transmit`, `drop`, `set-cos-transmit`, `set-dscp-transmit`,
     * `set-prec-transmit`, `set-qos-transmit` - Default value: `drop`
     */
    public readonly violateAction!: pulumi.Output<string>;
    /**
     * Set CoS for violating traffic. - Range: `0`-`7` - Default value: `0`
     */
    public readonly violateSetCos!: pulumi.Output<number>;
    /**
     * Set DSCP for violating traffic. - Range: `0`-`63` - Default value: `0`
     */
    public readonly violateSetDscp!: pulumi.Output<number>;
    /**
     * Set precedence for violating traffic. - Choices: `routine`, `priority`, `immediate`, `flash`, `flash-override`,
     * `critical`, `internet`, `network` - Default value: `routine`
     */
    public readonly violateSetPrecedence!: pulumi.Output<string>;
    /**
     * Set qos-group for violating traffic. - Range: `0`-`7` - Default value: `0`
     */
    public readonly violateSetQosGroup!: pulumi.Output<number>;

    /**
     * Create a DefaultQosPolicyMapMatchClassMapPolice resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DefaultQosPolicyMapMatchClassMapPoliceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefaultQosPolicyMapMatchClassMapPoliceArgs | DefaultQosPolicyMapMatchClassMapPoliceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DefaultQosPolicyMapMatchClassMapPoliceState | undefined;
            resourceInputs["bcRate"] = state ? state.bcRate : undefined;
            resourceInputs["bcUnit"] = state ? state.bcUnit : undefined;
            resourceInputs["beRate"] = state ? state.beRate : undefined;
            resourceInputs["beUnit"] = state ? state.beUnit : undefined;
            resourceInputs["cirRate"] = state ? state.cirRate : undefined;
            resourceInputs["cirUnit"] = state ? state.cirUnit : undefined;
            resourceInputs["classMapName"] = state ? state.classMapName : undefined;
            resourceInputs["conformAction"] = state ? state.conformAction : undefined;
            resourceInputs["conformSetCos"] = state ? state.conformSetCos : undefined;
            resourceInputs["conformSetDscp"] = state ? state.conformSetDscp : undefined;
            resourceInputs["conformSetPrecedence"] = state ? state.conformSetPrecedence : undefined;
            resourceInputs["conformSetQosGroup"] = state ? state.conformSetQosGroup : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["exceedAction"] = state ? state.exceedAction : undefined;
            resourceInputs["exceedSetCos"] = state ? state.exceedSetCos : undefined;
            resourceInputs["exceedSetDscp"] = state ? state.exceedSetDscp : undefined;
            resourceInputs["exceedSetPrecedence"] = state ? state.exceedSetPrecedence : undefined;
            resourceInputs["exceedSetQosGroup"] = state ? state.exceedSetQosGroup : undefined;
            resourceInputs["pirRate"] = state ? state.pirRate : undefined;
            resourceInputs["pirUnit"] = state ? state.pirUnit : undefined;
            resourceInputs["policyMapName"] = state ? state.policyMapName : undefined;
            resourceInputs["violateAction"] = state ? state.violateAction : undefined;
            resourceInputs["violateSetCos"] = state ? state.violateSetCos : undefined;
            resourceInputs["violateSetDscp"] = state ? state.violateSetDscp : undefined;
            resourceInputs["violateSetPrecedence"] = state ? state.violateSetPrecedence : undefined;
            resourceInputs["violateSetQosGroup"] = state ? state.violateSetQosGroup : undefined;
        } else {
            const args = argsOrState as DefaultQosPolicyMapMatchClassMapPoliceArgs | undefined;
            if ((!args || args.cirRate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cirRate'");
            }
            if ((!args || args.classMapName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'classMapName'");
            }
            if ((!args || args.policyMapName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyMapName'");
            }
            resourceInputs["bcRate"] = args ? args.bcRate : undefined;
            resourceInputs["bcUnit"] = args ? args.bcUnit : undefined;
            resourceInputs["beRate"] = args ? args.beRate : undefined;
            resourceInputs["beUnit"] = args ? args.beUnit : undefined;
            resourceInputs["cirRate"] = args ? args.cirRate : undefined;
            resourceInputs["cirUnit"] = args ? args.cirUnit : undefined;
            resourceInputs["classMapName"] = args ? args.classMapName : undefined;
            resourceInputs["conformAction"] = args ? args.conformAction : undefined;
            resourceInputs["conformSetCos"] = args ? args.conformSetCos : undefined;
            resourceInputs["conformSetDscp"] = args ? args.conformSetDscp : undefined;
            resourceInputs["conformSetPrecedence"] = args ? args.conformSetPrecedence : undefined;
            resourceInputs["conformSetQosGroup"] = args ? args.conformSetQosGroup : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["exceedAction"] = args ? args.exceedAction : undefined;
            resourceInputs["exceedSetCos"] = args ? args.exceedSetCos : undefined;
            resourceInputs["exceedSetDscp"] = args ? args.exceedSetDscp : undefined;
            resourceInputs["exceedSetPrecedence"] = args ? args.exceedSetPrecedence : undefined;
            resourceInputs["exceedSetQosGroup"] = args ? args.exceedSetQosGroup : undefined;
            resourceInputs["pirRate"] = args ? args.pirRate : undefined;
            resourceInputs["pirUnit"] = args ? args.pirUnit : undefined;
            resourceInputs["policyMapName"] = args ? args.policyMapName : undefined;
            resourceInputs["violateAction"] = args ? args.violateAction : undefined;
            resourceInputs["violateSetCos"] = args ? args.violateSetCos : undefined;
            resourceInputs["violateSetDscp"] = args ? args.violateSetDscp : undefined;
            resourceInputs["violateSetPrecedence"] = args ? args.violateSetPrecedence : undefined;
            resourceInputs["violateSetQosGroup"] = args ? args.violateSetQosGroup : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DefaultQosPolicyMapMatchClassMapPolice.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefaultQosPolicyMapMatchClassMapPolice resources.
 */
export interface DefaultQosPolicyMapMatchClassMapPoliceState {
    /**
     * CIR burst rate. - Range: `0`-`536870912` - Default value: `200`
     */
    bcRate?: pulumi.Input<number>;
    /**
     * CIR burst rate unit. - Choices: `unspecified`, `bytes`, `kbytes`, `mbytes`, `ms`, `us`, `packets` - Default value: `ms`
     */
    bcUnit?: pulumi.Input<string>;
    /**
     * PIR burst rate. - Range: `0`-`536870912` - Default value: `0`
     */
    beRate?: pulumi.Input<number>;
    /**
     * PIR burst rate unit. - Choices: `unspecified`, `bytes`, `kbytes`, `mbytes`, `ms`, `us`, `packets` - Default value:
     * `unspecified`
     */
    beUnit?: pulumi.Input<string>;
    /**
     * CIR rate. - Range: `0`-`100000000000` - Default value: `0`
     */
    cirRate?: pulumi.Input<number>;
    /**
     * CIR rate unit. - Choices: `unspecified`, `bps`, `kbps`, `mbps`, `gbps`, `pps`, `pct` - Default value: `bps`
     */
    cirUnit?: pulumi.Input<string>;
    /**
     * Class map name.
     */
    classMapName?: pulumi.Input<string>;
    /**
     * Conform action. - Choices: `unspecified`, `transmit`, `drop`, `set-cos-transmit`, `set-dscp-transmit`,
     * `set-prec-transmit`, `set-qos-transmit` - Default value: `transmit`
     */
    conformAction?: pulumi.Input<string>;
    /**
     * Set CoS for conforming traffic. - Range: `0`-`7` - Default value: `0`
     */
    conformSetCos?: pulumi.Input<number>;
    /**
     * Set DSCP for conforming traffic. - Range: `0`-`63` - Default value: `0`
     */
    conformSetDscp?: pulumi.Input<number>;
    /**
     * Set precedence for conforming traffic. - Choices: `routine`, `priority`, `immediate`, `flash`, `flash-override`,
     * `critical`, `internet`, `network` - Default value: `routine`
     */
    conformSetPrecedence?: pulumi.Input<string>;
    /**
     * Set qos-group for conforming traffic. - Range: `0`-`7` - Default value: `0`
     */
    conformSetQosGroup?: pulumi.Input<number>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Exceed action. - Choices: `unspecified`, `transmit`, `drop`, `set-cos-transmit`, `set-dscp-transmit`,
     * `set-prec-transmit`, `set-qos-transmit` - Default value: `unspecified`
     */
    exceedAction?: pulumi.Input<string>;
    /**
     * Set CoS for exceeding traffic. - Range: `0`-`7` - Default value: `0`
     */
    exceedSetCos?: pulumi.Input<number>;
    /**
     * Set DSCP for exceeding traffic. - Range: `0`-`63` - Default value: `0`
     */
    exceedSetDscp?: pulumi.Input<number>;
    /**
     * Set precedence for exceeding traffic. - Choices: `routine`, `priority`, `immediate`, `flash`, `flash-override`,
     * `critical`, `internet`, `network` - Default value: `routine`
     */
    exceedSetPrecedence?: pulumi.Input<string>;
    /**
     * Set qos-group for exceeding traffic. - Range: `0`-`7` - Default value: `0`
     */
    exceedSetQosGroup?: pulumi.Input<number>;
    /**
     * PIR rate. - Range: `0`-`100000000000` - Default value: `0`
     */
    pirRate?: pulumi.Input<number>;
    /**
     * PIR rate unit. - Choices: `unspecified`, `bps`, `kbps`, `mbps`, `gbps`, `pps`, `pct` - Default value: `unspecified`
     */
    pirUnit?: pulumi.Input<string>;
    /**
     * Policy map name.
     */
    policyMapName?: pulumi.Input<string>;
    /**
     * Violate action. - Choices: `unspecified`, `transmit`, `drop`, `set-cos-transmit`, `set-dscp-transmit`,
     * `set-prec-transmit`, `set-qos-transmit` - Default value: `drop`
     */
    violateAction?: pulumi.Input<string>;
    /**
     * Set CoS for violating traffic. - Range: `0`-`7` - Default value: `0`
     */
    violateSetCos?: pulumi.Input<number>;
    /**
     * Set DSCP for violating traffic. - Range: `0`-`63` - Default value: `0`
     */
    violateSetDscp?: pulumi.Input<number>;
    /**
     * Set precedence for violating traffic. - Choices: `routine`, `priority`, `immediate`, `flash`, `flash-override`,
     * `critical`, `internet`, `network` - Default value: `routine`
     */
    violateSetPrecedence?: pulumi.Input<string>;
    /**
     * Set qos-group for violating traffic. - Range: `0`-`7` - Default value: `0`
     */
    violateSetQosGroup?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a DefaultQosPolicyMapMatchClassMapPolice resource.
 */
export interface DefaultQosPolicyMapMatchClassMapPoliceArgs {
    /**
     * CIR burst rate. - Range: `0`-`536870912` - Default value: `200`
     */
    bcRate?: pulumi.Input<number>;
    /**
     * CIR burst rate unit. - Choices: `unspecified`, `bytes`, `kbytes`, `mbytes`, `ms`, `us`, `packets` - Default value: `ms`
     */
    bcUnit?: pulumi.Input<string>;
    /**
     * PIR burst rate. - Range: `0`-`536870912` - Default value: `0`
     */
    beRate?: pulumi.Input<number>;
    /**
     * PIR burst rate unit. - Choices: `unspecified`, `bytes`, `kbytes`, `mbytes`, `ms`, `us`, `packets` - Default value:
     * `unspecified`
     */
    beUnit?: pulumi.Input<string>;
    /**
     * CIR rate. - Range: `0`-`100000000000` - Default value: `0`
     */
    cirRate: pulumi.Input<number>;
    /**
     * CIR rate unit. - Choices: `unspecified`, `bps`, `kbps`, `mbps`, `gbps`, `pps`, `pct` - Default value: `bps`
     */
    cirUnit?: pulumi.Input<string>;
    /**
     * Class map name.
     */
    classMapName: pulumi.Input<string>;
    /**
     * Conform action. - Choices: `unspecified`, `transmit`, `drop`, `set-cos-transmit`, `set-dscp-transmit`,
     * `set-prec-transmit`, `set-qos-transmit` - Default value: `transmit`
     */
    conformAction?: pulumi.Input<string>;
    /**
     * Set CoS for conforming traffic. - Range: `0`-`7` - Default value: `0`
     */
    conformSetCos?: pulumi.Input<number>;
    /**
     * Set DSCP for conforming traffic. - Range: `0`-`63` - Default value: `0`
     */
    conformSetDscp?: pulumi.Input<number>;
    /**
     * Set precedence for conforming traffic. - Choices: `routine`, `priority`, `immediate`, `flash`, `flash-override`,
     * `critical`, `internet`, `network` - Default value: `routine`
     */
    conformSetPrecedence?: pulumi.Input<string>;
    /**
     * Set qos-group for conforming traffic. - Range: `0`-`7` - Default value: `0`
     */
    conformSetQosGroup?: pulumi.Input<number>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Exceed action. - Choices: `unspecified`, `transmit`, `drop`, `set-cos-transmit`, `set-dscp-transmit`,
     * `set-prec-transmit`, `set-qos-transmit` - Default value: `unspecified`
     */
    exceedAction?: pulumi.Input<string>;
    /**
     * Set CoS for exceeding traffic. - Range: `0`-`7` - Default value: `0`
     */
    exceedSetCos?: pulumi.Input<number>;
    /**
     * Set DSCP for exceeding traffic. - Range: `0`-`63` - Default value: `0`
     */
    exceedSetDscp?: pulumi.Input<number>;
    /**
     * Set precedence for exceeding traffic. - Choices: `routine`, `priority`, `immediate`, `flash`, `flash-override`,
     * `critical`, `internet`, `network` - Default value: `routine`
     */
    exceedSetPrecedence?: pulumi.Input<string>;
    /**
     * Set qos-group for exceeding traffic. - Range: `0`-`7` - Default value: `0`
     */
    exceedSetQosGroup?: pulumi.Input<number>;
    /**
     * PIR rate. - Range: `0`-`100000000000` - Default value: `0`
     */
    pirRate?: pulumi.Input<number>;
    /**
     * PIR rate unit. - Choices: `unspecified`, `bps`, `kbps`, `mbps`, `gbps`, `pps`, `pct` - Default value: `unspecified`
     */
    pirUnit?: pulumi.Input<string>;
    /**
     * Policy map name.
     */
    policyMapName: pulumi.Input<string>;
    /**
     * Violate action. - Choices: `unspecified`, `transmit`, `drop`, `set-cos-transmit`, `set-dscp-transmit`,
     * `set-prec-transmit`, `set-qos-transmit` - Default value: `drop`
     */
    violateAction?: pulumi.Input<string>;
    /**
     * Set CoS for violating traffic. - Range: `0`-`7` - Default value: `0`
     */
    violateSetCos?: pulumi.Input<number>;
    /**
     * Set DSCP for violating traffic. - Range: `0`-`63` - Default value: `0`
     */
    violateSetDscp?: pulumi.Input<number>;
    /**
     * Set precedence for violating traffic. - Choices: `routine`, `priority`, `immediate`, `flash`, `flash-override`,
     * `critical`, `internet`, `network` - Default value: `routine`
     */
    violateSetPrecedence?: pulumi.Input<string>;
    /**
     * Set qos-group for violating traffic. - Range: `0`-`7` - Default value: `0`
     */
    violateSetQosGroup?: pulumi.Input<number>;
}