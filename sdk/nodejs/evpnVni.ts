// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage a EVPN VNI Route Distinguisher.
 *
 * - API Documentation: [rtctrlBDEvi](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:BDEvi/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.EvpnVni("example", {
 *     encap: "vxlan-123456",
 *     routeDistinguisher: "rd:unknown:0:0",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/evpnVni:EvpnVni example "sys/evpn/bdevi-[vxlan-123456]"
 * ```
 */
export class EvpnVni extends pulumi.CustomResource {
    /**
     * Get an existing EvpnVni resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EvpnVniState, opts?: pulumi.CustomResourceOptions): EvpnVni {
        return new EvpnVni(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/evpnVni:EvpnVni';

    /**
     * Returns true if the given object is an instance of EvpnVni.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EvpnVni {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EvpnVni.__pulumiType;
    }

    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     */
    public readonly encap!: pulumi.Output<string>;
    /**
     * Route Distinguisher value in NX-OS DME format. - Default value: `unknown:unknown:0:0`
     */
    public readonly routeDistinguisher!: pulumi.Output<string>;

    /**
     * Create a EvpnVni resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EvpnVniArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EvpnVniArgs | EvpnVniState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EvpnVniState | undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["encap"] = state ? state.encap : undefined;
            resourceInputs["routeDistinguisher"] = state ? state.routeDistinguisher : undefined;
        } else {
            const args = argsOrState as EvpnVniArgs | undefined;
            if ((!args || args.encap === undefined) && !opts.urn) {
                throw new Error("Missing required property 'encap'");
            }
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["encap"] = args ? args.encap : undefined;
            resourceInputs["routeDistinguisher"] = args ? args.routeDistinguisher : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EvpnVni.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EvpnVni resources.
 */
export interface EvpnVniState {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     */
    encap?: pulumi.Input<string>;
    /**
     * Route Distinguisher value in NX-OS DME format. - Default value: `unknown:unknown:0:0`
     */
    routeDistinguisher?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EvpnVni resource.
 */
export interface EvpnVniArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     */
    encap: pulumi.Input<string>;
    /**
     * Route Distinguisher value in NX-OS DME format. - Default value: `unknown:unknown:0:0`
     */
    routeDistinguisher?: pulumi.Input<string>;
}
