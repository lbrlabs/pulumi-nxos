// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage the global PIM configuration.
 *
 * - API Documentation: [pimEntity](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:Entity/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.Pim("example", {adminState: "enabled"});
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/pim:Pim example "sys/pim"
 * ```
 */
export class Pim extends pulumi.CustomResource {
    /**
     * Get an existing Pim resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PimState, opts?: pulumi.CustomResourceOptions): Pim {
        return new Pim(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/pim:Pim';

    /**
     * Returns true if the given object is an instance of Pim.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Pim {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Pim.__pulumiType;
    }

    /**
     * Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     */
    public readonly adminState!: pulumi.Output<string>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;

    /**
     * Create a Pim resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: PimArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PimArgs | PimState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PimState | undefined;
            resourceInputs["adminState"] = state ? state.adminState : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
        } else {
            const args = argsOrState as PimArgs | undefined;
            resourceInputs["adminState"] = args ? args.adminState : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Pim.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Pim resources.
 */
export interface PimState {
    /**
     * Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     */
    adminState?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Pim resource.
 */
export interface PimArgs {
    /**
     * Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     */
    adminState?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}
