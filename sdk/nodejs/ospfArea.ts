// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage the OSPF Area configuration.
 *
 * - API Documentation: [ospfArea](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/ospf:Area/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.OspfArea("example", {
 *     areaId: "0.0.0.10",
 *     authenticationType: "none",
 *     cost: 10,
 *     instanceName: "OSPF1",
 *     type: "stub",
 *     vrfName: "VRF1",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/ospfArea:OspfArea example "sys/ospf/inst-[OSPF1]/dom-[VRF1]/area-[0.0.0.10]"
 * ```
 */
export class OspfArea extends pulumi.CustomResource {
    /**
     * Get an existing OspfArea resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OspfAreaState, opts?: pulumi.CustomResourceOptions): OspfArea {
        return new OspfArea(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/ospfArea:OspfArea';

    /**
     * Returns true if the given object is an instance of OspfArea.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OspfArea {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OspfArea.__pulumiType;
    }

    /**
     * Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
     */
    public readonly areaId!: pulumi.Output<string>;
    /**
     * Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
     */
    public readonly authenticationType!: pulumi.Output<string>;
    /**
     * Area cost, specifies cost for default summary LSAs. Used with nssa/stub area types. - Range: `0`-`16777215` - Default
     * value: `1`
     */
    public readonly cost!: pulumi.Output<number>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * OSPF instance name.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * Area type. - Choices: `regular`, `stub`, `nssa` - Default value: `regular`
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * VRF name.
     */
    public readonly vrfName!: pulumi.Output<string>;

    /**
     * Create a OspfArea resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OspfAreaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OspfAreaArgs | OspfAreaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OspfAreaState | undefined;
            resourceInputs["areaId"] = state ? state.areaId : undefined;
            resourceInputs["authenticationType"] = state ? state.authenticationType : undefined;
            resourceInputs["cost"] = state ? state.cost : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vrfName"] = state ? state.vrfName : undefined;
        } else {
            const args = argsOrState as OspfAreaArgs | undefined;
            if ((!args || args.areaId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'areaId'");
            }
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            if ((!args || args.vrfName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vrfName'");
            }
            resourceInputs["areaId"] = args ? args.areaId : undefined;
            resourceInputs["authenticationType"] = args ? args.authenticationType : undefined;
            resourceInputs["cost"] = args ? args.cost : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vrfName"] = args ? args.vrfName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OspfArea.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OspfArea resources.
 */
export interface OspfAreaState {
    /**
     * Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
     */
    areaId?: pulumi.Input<string>;
    /**
     * Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
     */
    authenticationType?: pulumi.Input<string>;
    /**
     * Area cost, specifies cost for default summary LSAs. Used with nssa/stub area types. - Range: `0`-`16777215` - Default
     * value: `1`
     */
    cost?: pulumi.Input<number>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * OSPF instance name.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Area type. - Choices: `regular`, `stub`, `nssa` - Default value: `regular`
     */
    type?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrfName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OspfArea resource.
 */
export interface OspfAreaArgs {
    /**
     * Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
     */
    areaId: pulumi.Input<string>;
    /**
     * Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
     */
    authenticationType?: pulumi.Input<string>;
    /**
     * Area cost, specifies cost for default summary LSAs. Used with nssa/stub area types. - Range: `0`-`16777215` - Default
     * value: `1`
     */
    cost?: pulumi.Input<number>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * OSPF instance name.
     */
    instanceName: pulumi.Input<string>;
    /**
     * Area type. - Choices: `regular`, `stub`, `nssa` - Default value: `regular`
     */
    type?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrfName: pulumi.Input<string>;
}