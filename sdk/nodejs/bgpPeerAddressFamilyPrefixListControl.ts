// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage the BGP peer address family prefix list control configuration.
 *
 * - API Documentation: [bgpPfxCtrlP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:PfxCtrlP/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.BgpPeerAddressFamilyPrefixListControl("example", {
 *     address: "192.168.0.1",
 *     addressFamily: "ipv4-ucast",
 *     asn: "65001",
 *     direction: "in",
 *     list: "PREFIX_LIST1",
 *     vrf: "default",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/bgpPeerAddressFamilyPrefixListControl:BgpPeerAddressFamilyPrefixListControl example "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]/af-[ipv4-ucast]/pfxctrl-[in]"
 * ```
 */
export class BgpPeerAddressFamilyPrefixListControl extends pulumi.CustomResource {
    /**
     * Get an existing BgpPeerAddressFamilyPrefixListControl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BgpPeerAddressFamilyPrefixListControlState, opts?: pulumi.CustomResourceOptions): BgpPeerAddressFamilyPrefixListControl {
        return new BgpPeerAddressFamilyPrefixListControl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/bgpPeerAddressFamilyPrefixListControl:BgpPeerAddressFamilyPrefixListControl';

    /**
     * Returns true if the given object is an instance of BgpPeerAddressFamilyPrefixListControl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BgpPeerAddressFamilyPrefixListControl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BgpPeerAddressFamilyPrefixListControl.__pulumiType;
    }

    /**
     * Peer address.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * Address Family. - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate` - Default
     * value: `ipv4-ucast`
     */
    public readonly addressFamily!: pulumi.Output<string>;
    /**
     * Autonomous system number.
     */
    public readonly asn!: pulumi.Output<string>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Route Control direction. - Choices: `in`, `out` - Default value: `in`
     */
    public readonly direction!: pulumi.Output<string>;
    /**
     * Route Control Prefix-List name.
     */
    public readonly list!: pulumi.Output<string | undefined>;
    /**
     * VRF name.
     */
    public readonly vrf!: pulumi.Output<string>;

    /**
     * Create a BgpPeerAddressFamilyPrefixListControl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BgpPeerAddressFamilyPrefixListControlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BgpPeerAddressFamilyPrefixListControlArgs | BgpPeerAddressFamilyPrefixListControlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BgpPeerAddressFamilyPrefixListControlState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["addressFamily"] = state ? state.addressFamily : undefined;
            resourceInputs["asn"] = state ? state.asn : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["direction"] = state ? state.direction : undefined;
            resourceInputs["list"] = state ? state.list : undefined;
            resourceInputs["vrf"] = state ? state.vrf : undefined;
        } else {
            const args = argsOrState as BgpPeerAddressFamilyPrefixListControlArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.addressFamily === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if ((!args || args.asn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'asn'");
            }
            if ((!args || args.direction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'direction'");
            }
            if ((!args || args.vrf === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vrf'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["addressFamily"] = args ? args.addressFamily : undefined;
            resourceInputs["asn"] = args ? args.asn : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["direction"] = args ? args.direction : undefined;
            resourceInputs["list"] = args ? args.list : undefined;
            resourceInputs["vrf"] = args ? args.vrf : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BgpPeerAddressFamilyPrefixListControl.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BgpPeerAddressFamilyPrefixListControl resources.
 */
export interface BgpPeerAddressFamilyPrefixListControlState {
    /**
     * Peer address.
     */
    address?: pulumi.Input<string>;
    /**
     * Address Family. - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate` - Default
     * value: `ipv4-ucast`
     */
    addressFamily?: pulumi.Input<string>;
    /**
     * Autonomous system number.
     */
    asn?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Route Control direction. - Choices: `in`, `out` - Default value: `in`
     */
    direction?: pulumi.Input<string>;
    /**
     * Route Control Prefix-List name.
     */
    list?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrf?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BgpPeerAddressFamilyPrefixListControl resource.
 */
export interface BgpPeerAddressFamilyPrefixListControlArgs {
    /**
     * Peer address.
     */
    address: pulumi.Input<string>;
    /**
     * Address Family. - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate` - Default
     * value: `ipv4-ucast`
     */
    addressFamily: pulumi.Input<string>;
    /**
     * Autonomous system number.
     */
    asn: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Route Control direction. - Choices: `in`, `out` - Default value: `in`
     */
    direction: pulumi.Input<string>;
    /**
     * Route Control Prefix-List name.
     */
    list?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrf: pulumi.Input<string>;
}
