// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Manages NX-OS DME Objects via REST API calls. This resource can manage a single API object and its children. It is able to read the state and therefore reconcile configuration drift.
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/rest:Rest l1PhysIf" "l1PhysIf:sys/intf/phys-[eth1/1]"
 * ```
 */
export class Rest extends pulumi.CustomResource {
    /**
     * Get an existing Rest resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RestState, opts?: pulumi.CustomResourceOptions): Rest {
        return new Rest(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/rest:Rest';

    /**
     * Returns true if the given object is an instance of Rest.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Rest {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Rest.__pulumiType;
    }

    /**
     * List of children.
     */
    public readonly childrens!: pulumi.Output<outputs.RestChildren[] | undefined>;
    /**
     * Which class object is being created. (Make sure there is no colon in the classname)
     */
    public readonly className!: pulumi.Output<string>;
    /**
     * Map of key-value pairs that need to be passed to the Model object as parameters.
     */
    public readonly content!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Delete object during destroy operation. Default value is `true`.
     */
    public readonly delete!: pulumi.Output<boolean>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].
     */
    public readonly dn!: pulumi.Output<string>;

    /**
     * Create a Rest resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RestArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RestArgs | RestState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RestState | undefined;
            resourceInputs["childrens"] = state ? state.childrens : undefined;
            resourceInputs["className"] = state ? state.className : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["delete"] = state ? state.delete : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["dn"] = state ? state.dn : undefined;
        } else {
            const args = argsOrState as RestArgs | undefined;
            if ((!args || args.className === undefined) && !opts.urn) {
                throw new Error("Missing required property 'className'");
            }
            if ((!args || args.dn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dn'");
            }
            resourceInputs["childrens"] = args ? args.childrens : undefined;
            resourceInputs["className"] = args ? args.className : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["delete"] = args ? args.delete : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["dn"] = args ? args.dn : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Rest.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Rest resources.
 */
export interface RestState {
    /**
     * List of children.
     */
    childrens?: pulumi.Input<pulumi.Input<inputs.RestChildren>[]>;
    /**
     * Which class object is being created. (Make sure there is no colon in the classname)
     */
    className?: pulumi.Input<string>;
    /**
     * Map of key-value pairs that need to be passed to the Model object as parameters.
     */
    content?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Delete object during destroy operation. Default value is `true`.
     */
    delete?: pulumi.Input<boolean>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].
     */
    dn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Rest resource.
 */
export interface RestArgs {
    /**
     * List of children.
     */
    childrens?: pulumi.Input<pulumi.Input<inputs.RestChildren>[]>;
    /**
     * Which class object is being created. (Make sure there is no colon in the classname)
     */
    className: pulumi.Input<string>;
    /**
     * Map of key-value pairs that need to be passed to the Model object as parameters.
     */
    content?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Delete object during destroy operation. Default value is `true`.
     */
    delete?: pulumi.Input<boolean>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].
     */
    dn: pulumi.Input<string>;
}