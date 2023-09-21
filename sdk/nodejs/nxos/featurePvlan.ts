// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class FeaturePvlan extends pulumi.CustomResource {
    /**
     * Get an existing FeaturePvlan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FeaturePvlanState, opts?: pulumi.CustomResourceOptions): FeaturePvlan {
        return new FeaturePvlan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:nxos/featurePvlan:FeaturePvlan';

    /**
     * Returns true if the given object is an instance of FeaturePvlan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FeaturePvlan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FeaturePvlan.__pulumiType;
    }

    /**
     * Administrative state. - Choices: `enabled`, `disabled`
     */
    public readonly adminState!: pulumi.Output<string>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;

    /**
     * Create a FeaturePvlan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FeaturePvlanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FeaturePvlanArgs | FeaturePvlanState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FeaturePvlanState | undefined;
            resourceInputs["adminState"] = state ? state.adminState : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
        } else {
            const args = argsOrState as FeaturePvlanArgs | undefined;
            if ((!args || args.adminState === undefined) && !opts.urn) {
                throw new Error("Missing required property 'adminState'");
            }
            resourceInputs["adminState"] = args ? args.adminState : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FeaturePvlan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FeaturePvlan resources.
 */
export interface FeaturePvlanState {
    /**
     * Administrative state. - Choices: `enabled`, `disabled`
     */
    adminState?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FeaturePvlan resource.
 */
export interface FeaturePvlanArgs {
    /**
     * Administrative state. - Choices: `enabled`, `disabled`
     */
    adminState: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}
