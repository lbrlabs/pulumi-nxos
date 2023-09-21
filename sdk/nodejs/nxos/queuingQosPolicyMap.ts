// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class QueuingQosPolicyMap extends pulumi.CustomResource {
    /**
     * Get an existing QueuingQosPolicyMap resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QueuingQosPolicyMapState, opts?: pulumi.CustomResourceOptions): QueuingQosPolicyMap {
        return new QueuingQosPolicyMap(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:nxos/queuingQosPolicyMap:QueuingQosPolicyMap';

    /**
     * Returns true if the given object is an instance of QueuingQosPolicyMap.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QueuingQosPolicyMap {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QueuingQosPolicyMap.__pulumiType;
    }

    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
     */
    public readonly matchType!: pulumi.Output<string>;
    /**
     * Policy map name.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a QueuingQosPolicyMap resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: QueuingQosPolicyMapArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QueuingQosPolicyMapArgs | QueuingQosPolicyMapState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QueuingQosPolicyMapState | undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["matchType"] = state ? state.matchType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as QueuingQosPolicyMapArgs | undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["matchType"] = args ? args.matchType : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(QueuingQosPolicyMap.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering QueuingQosPolicyMap resources.
 */
export interface QueuingQosPolicyMapState {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
     */
    matchType?: pulumi.Input<string>;
    /**
     * Policy map name.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a QueuingQosPolicyMap resource.
 */
export interface QueuingQosPolicyMapArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Match type. - Choices: `match-any`, `match-all`, `match-first` - Default value: `match-all`
     */
    matchType?: pulumi.Input<string>;
    /**
     * Policy map name.
     */
    name?: pulumi.Input<string>;
}
