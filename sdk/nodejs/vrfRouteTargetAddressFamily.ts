// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage a VRF Route Target Address Family.
 *
 * - API Documentation: [rtctrlAfCtrl](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:AfCtrl/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.VrfRouteTargetAddressFamily("example", {
 *     addressFamily: "ipv4-ucast",
 *     routeTargetAddressFamily: "ipv4-ucast",
 *     vrf: "VRF1",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/vrfRouteTargetAddressFamily:VrfRouteTargetAddressFamily example "sys/inst-[VRF1]/dom-[VRF1]/af-[ipv4-ucast]/ctrl-[ipv4-ucast]"
 * ```
 */
export class VrfRouteTargetAddressFamily extends pulumi.CustomResource {
    /**
     * Get an existing VrfRouteTargetAddressFamily resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VrfRouteTargetAddressFamilyState, opts?: pulumi.CustomResourceOptions): VrfRouteTargetAddressFamily {
        return new VrfRouteTargetAddressFamily(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/vrfRouteTargetAddressFamily:VrfRouteTargetAddressFamily';

    /**
     * Returns true if the given object is an instance of VrfRouteTargetAddressFamily.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VrfRouteTargetAddressFamily {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VrfRouteTargetAddressFamily.__pulumiType;
    }

    /**
     * Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
     */
    public readonly addressFamily!: pulumi.Output<string>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
     */
    public readonly routeTargetAddressFamily!: pulumi.Output<string>;
    /**
     * VRF name.
     */
    public readonly vrf!: pulumi.Output<string>;

    /**
     * Create a VrfRouteTargetAddressFamily resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VrfRouteTargetAddressFamilyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VrfRouteTargetAddressFamilyArgs | VrfRouteTargetAddressFamilyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VrfRouteTargetAddressFamilyState | undefined;
            resourceInputs["addressFamily"] = state ? state.addressFamily : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["routeTargetAddressFamily"] = state ? state.routeTargetAddressFamily : undefined;
            resourceInputs["vrf"] = state ? state.vrf : undefined;
        } else {
            const args = argsOrState as VrfRouteTargetAddressFamilyArgs | undefined;
            if ((!args || args.addressFamily === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if ((!args || args.routeTargetAddressFamily === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routeTargetAddressFamily'");
            }
            if ((!args || args.vrf === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vrf'");
            }
            resourceInputs["addressFamily"] = args ? args.addressFamily : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["routeTargetAddressFamily"] = args ? args.routeTargetAddressFamily : undefined;
            resourceInputs["vrf"] = args ? args.vrf : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VrfRouteTargetAddressFamily.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VrfRouteTargetAddressFamily resources.
 */
export interface VrfRouteTargetAddressFamilyState {
    /**
     * Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
     */
    addressFamily?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
     */
    routeTargetAddressFamily?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrf?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VrfRouteTargetAddressFamily resource.
 */
export interface VrfRouteTargetAddressFamilyArgs {
    /**
     * Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
     */
    addressFamily: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
     */
    routeTargetAddressFamily: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrf: pulumi.Input<string>;
}
