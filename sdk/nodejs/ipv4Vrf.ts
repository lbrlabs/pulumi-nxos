// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage the IPv4 VRF information.
 *
 * - API Documentation: [ipv4Dom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/ipv4:Dom/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.Ipv4Vrf("example", {});
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/ipv4Vrf:Ipv4Vrf example "sys/ipv4/inst/dom-[VRF1]"
 * ```
 */
export class Ipv4Vrf extends pulumi.CustomResource {
    /**
     * Get an existing Ipv4Vrf resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ipv4VrfState, opts?: pulumi.CustomResourceOptions): Ipv4Vrf {
        return new Ipv4Vrf(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/ipv4Vrf:Ipv4Vrf';

    /**
     * Returns true if the given object is an instance of Ipv4Vrf.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ipv4Vrf {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ipv4Vrf.__pulumiType;
    }

    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * VRF name.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a Ipv4Vrf resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: Ipv4VrfArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ipv4VrfArgs | Ipv4VrfState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Ipv4VrfState | undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as Ipv4VrfArgs | undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ipv4Vrf.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ipv4Vrf resources.
 */
export interface Ipv4VrfState {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ipv4Vrf resource.
 */
export interface Ipv4VrfArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    name?: pulumi.Input<string>;
}
