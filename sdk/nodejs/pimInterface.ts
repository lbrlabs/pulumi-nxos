// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage the PIM interface configuration.
 *
 * - API Documentation: [pimIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:If/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.PimInterface("example", {
 *     adminState: "enabled",
 *     bfd: "enabled",
 *     drPriority: 10,
 *     interfaceId: "eth1/10",
 *     passive: false,
 *     sparseMode: true,
 *     vrfName: "default",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/pimInterface:PimInterface example "sys/pim/inst/dom-[default]/if-[eth1/10]"
 * ```
 */
export class PimInterface extends pulumi.CustomResource {
    /**
     * Get an existing PimInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PimInterfaceState, opts?: pulumi.CustomResourceOptions): PimInterface {
        return new PimInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/pimInterface:PimInterface';

    /**
     * Returns true if the given object is an instance of PimInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PimInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PimInterface.__pulumiType;
    }

    /**
     * Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     */
    public readonly adminState!: pulumi.Output<string>;
    /**
     * BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
     */
    public readonly bfd!: pulumi.Output<string>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
     */
    public readonly drPriority!: pulumi.Output<number>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    public readonly interfaceId!: pulumi.Output<string>;
    /**
     * Passive interface. - Default value: `false`
     */
    public readonly passive!: pulumi.Output<boolean>;
    /**
     * Sparse mode. - Default value: `false`
     */
    public readonly sparseMode!: pulumi.Output<boolean>;
    /**
     * VRF name.
     */
    public readonly vrfName!: pulumi.Output<string>;

    /**
     * Create a PimInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PimInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PimInterfaceArgs | PimInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PimInterfaceState | undefined;
            resourceInputs["adminState"] = state ? state.adminState : undefined;
            resourceInputs["bfd"] = state ? state.bfd : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["drPriority"] = state ? state.drPriority : undefined;
            resourceInputs["interfaceId"] = state ? state.interfaceId : undefined;
            resourceInputs["passive"] = state ? state.passive : undefined;
            resourceInputs["sparseMode"] = state ? state.sparseMode : undefined;
            resourceInputs["vrfName"] = state ? state.vrfName : undefined;
        } else {
            const args = argsOrState as PimInterfaceArgs | undefined;
            if ((!args || args.interfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interfaceId'");
            }
            if ((!args || args.vrfName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vrfName'");
            }
            resourceInputs["adminState"] = args ? args.adminState : undefined;
            resourceInputs["bfd"] = args ? args.bfd : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["drPriority"] = args ? args.drPriority : undefined;
            resourceInputs["interfaceId"] = args ? args.interfaceId : undefined;
            resourceInputs["passive"] = args ? args.passive : undefined;
            resourceInputs["sparseMode"] = args ? args.sparseMode : undefined;
            resourceInputs["vrfName"] = args ? args.vrfName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PimInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PimInterface resources.
 */
export interface PimInterfaceState {
    /**
     * Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     */
    adminState?: pulumi.Input<string>;
    /**
     * BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
     */
    bfd?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
     */
    drPriority?: pulumi.Input<number>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    interfaceId?: pulumi.Input<string>;
    /**
     * Passive interface. - Default value: `false`
     */
    passive?: pulumi.Input<boolean>;
    /**
     * Sparse mode. - Default value: `false`
     */
    sparseMode?: pulumi.Input<boolean>;
    /**
     * VRF name.
     */
    vrfName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PimInterface resource.
 */
export interface PimInterfaceArgs {
    /**
     * Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     */
    adminState?: pulumi.Input<string>;
    /**
     * BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
     */
    bfd?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
     */
    drPriority?: pulumi.Input<number>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    interfaceId: pulumi.Input<string>;
    /**
     * Passive interface. - Default value: `false`
     */
    passive?: pulumi.Input<boolean>;
    /**
     * Sparse mode. - Default value: `false`
     */
    sparseMode?: pulumi.Input<boolean>;
    /**
     * VRF name.
     */
    vrfName: pulumi.Input<string>;
}