// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage the PIM Anycast RP configuration.
 *
 * - API Documentation: [pimAcastRPFuncP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:AcastRPFuncP/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.PimAnycastRp("example", {
 *     localInterface: "eth1/10",
 *     sourceInterface: "eth1/10",
 *     vrfName: "default",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/pimAnycastRp:PimAnycastRp example "sys/pim/inst/dom-[default]/acastrpfunc"
 * ```
 */
export class PimAnycastRp extends pulumi.CustomResource {
    /**
     * Get an existing PimAnycastRp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PimAnycastRpState, opts?: pulumi.CustomResourceOptions): PimAnycastRp {
        return new PimAnycastRp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/pimAnycastRp:PimAnycastRp';

    /**
     * Returns true if the given object is an instance of PimAnycastRp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PimAnycastRp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PimAnycastRp.__pulumiType;
    }

    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    public readonly localInterface!: pulumi.Output<string | undefined>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    public readonly sourceInterface!: pulumi.Output<string | undefined>;
    /**
     * VRF name.
     */
    public readonly vrfName!: pulumi.Output<string>;

    /**
     * Create a PimAnycastRp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PimAnycastRpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PimAnycastRpArgs | PimAnycastRpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PimAnycastRpState | undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["localInterface"] = state ? state.localInterface : undefined;
            resourceInputs["sourceInterface"] = state ? state.sourceInterface : undefined;
            resourceInputs["vrfName"] = state ? state.vrfName : undefined;
        } else {
            const args = argsOrState as PimAnycastRpArgs | undefined;
            if ((!args || args.vrfName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vrfName'");
            }
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["localInterface"] = args ? args.localInterface : undefined;
            resourceInputs["sourceInterface"] = args ? args.sourceInterface : undefined;
            resourceInputs["vrfName"] = args ? args.vrfName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PimAnycastRp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PimAnycastRp resources.
 */
export interface PimAnycastRpState {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    localInterface?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    sourceInterface?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrfName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PimAnycastRp resource.
 */
export interface PimAnycastRpArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    localInterface?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    sourceInterface?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrfName: pulumi.Input<string>;
}
