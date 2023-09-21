// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class LoopbackInterface extends pulumi.CustomResource {
    /**
     * Get an existing LoopbackInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LoopbackInterfaceState, opts?: pulumi.CustomResourceOptions): LoopbackInterface {
        return new LoopbackInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:nxos/loopbackInterface:LoopbackInterface';

    /**
     * Returns true if the given object is an instance of LoopbackInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LoopbackInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LoopbackInterface.__pulumiType;
    }

    /**
     * Administrative state. - Choices: `up`, `down` - Default value: `up`
     */
    public readonly adminState!: pulumi.Output<string>;
    /**
     * Interface description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `lo123`.
     */
    public readonly interfaceId!: pulumi.Output<string>;

    /**
     * Create a LoopbackInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LoopbackInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LoopbackInterfaceArgs | LoopbackInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LoopbackInterfaceState | undefined;
            resourceInputs["adminState"] = state ? state.adminState : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["interfaceId"] = state ? state.interfaceId : undefined;
        } else {
            const args = argsOrState as LoopbackInterfaceArgs | undefined;
            if ((!args || args.interfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interfaceId'");
            }
            resourceInputs["adminState"] = args ? args.adminState : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["interfaceId"] = args ? args.interfaceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LoopbackInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LoopbackInterface resources.
 */
export interface LoopbackInterfaceState {
    /**
     * Administrative state. - Choices: `up`, `down` - Default value: `up`
     */
    adminState?: pulumi.Input<string>;
    /**
     * Interface description.
     */
    description?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `lo123`.
     */
    interfaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LoopbackInterface resource.
 */
export interface LoopbackInterfaceArgs {
    /**
     * Administrative state. - Choices: `up`, `down` - Default value: `up`
     */
    adminState?: pulumi.Input<string>;
    /**
     * Interface description.
     */
    description?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `lo123`.
     */
    interfaceId: pulumi.Input<string>;
}
