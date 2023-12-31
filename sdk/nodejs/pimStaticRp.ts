// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage the PIM Static RP configuration.
 *
 * - API Documentation: [pimStaticRP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:StaticRP/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.PimStaticRp("example", {
 *     address: "1.2.3.4",
 *     vrfName: "default",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/pimStaticRp:PimStaticRp example "sys/pim/inst/dom-[default]/staticrp/rp-[1.2.3.4]"
 * ```
 */
export class PimStaticRp extends pulumi.CustomResource {
    /**
     * Get an existing PimStaticRp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PimStaticRpState, opts?: pulumi.CustomResourceOptions): PimStaticRp {
        return new PimStaticRp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/pimStaticRp:PimStaticRp';

    /**
     * Returns true if the given object is an instance of PimStaticRp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PimStaticRp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PimStaticRp.__pulumiType;
    }

    /**
     * Address.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * VRF name.
     */
    public readonly vrfName!: pulumi.Output<string>;

    /**
     * Create a PimStaticRp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PimStaticRpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PimStaticRpArgs | PimStaticRpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PimStaticRpState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["vrfName"] = state ? state.vrfName : undefined;
        } else {
            const args = argsOrState as PimStaticRpArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.vrfName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vrfName'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["vrfName"] = args ? args.vrfName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PimStaticRp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PimStaticRp resources.
 */
export interface PimStaticRpState {
    /**
     * Address.
     */
    address?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrfName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PimStaticRp resource.
 */
export interface PimStaticRpArgs {
    /**
     * Address.
     */
    address: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrfName: pulumi.Input<string>;
}
