// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage the BGP peer template Maximum Prefix Policy configuration.
 *
 * - API Documentation: [bgpMaxPfxP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:MaxPfxP/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.BgpPeerTemplateMaxPrefix("example", {
 *     action: "log",
 *     addressFamily: "ipv4-ucast",
 *     asn: "65001",
 *     maximumPrefix: 10000,
 *     restartTime: 0,
 *     templateName: "SPINE-PEERS",
 *     threshold: 30,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/bgpPeerTemplateMaxPrefix:BgpPeerTemplateMaxPrefix example "sys/bgp/inst/dom-[default]/peercont-[SPINE-PEERS]/af-[ipv4-ucast]/maxpfxp"
 * ```
 */
export class BgpPeerTemplateMaxPrefix extends pulumi.CustomResource {
    /**
     * Get an existing BgpPeerTemplateMaxPrefix resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BgpPeerTemplateMaxPrefixState, opts?: pulumi.CustomResourceOptions): BgpPeerTemplateMaxPrefix {
        return new BgpPeerTemplateMaxPrefix(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/bgpPeerTemplateMaxPrefix:BgpPeerTemplateMaxPrefix';

    /**
     * Returns true if the given object is an instance of BgpPeerTemplateMaxPrefix.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BgpPeerTemplateMaxPrefix {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BgpPeerTemplateMaxPrefix.__pulumiType;
    }

    /**
     * Action to do when limit is exceeded. - Choices: `log`, `shut`, `restart` - Default value: `shut`
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
     * `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
     * `ipv4-mdt` - Default value: `ipv4-ucast`
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
     * Maximum number of prefixes allowed from the peer. - Range: `0`-`4294967295`
     */
    public readonly maximumPrefix!: pulumi.Output<number | undefined>;
    /**
     * The period of time in minutes before restarting the peer when the prefix limit is reached. - Default value: `0`
     */
    public readonly restartTime!: pulumi.Output<number>;
    /**
     * Peer template name.
     */
    public readonly templateName!: pulumi.Output<string>;
    /**
     * The period of time in minutes before restarting the peer when the prefix limit is reached. - Range: `0`-`100` - Default
     * value: `0`
     */
    public readonly threshold!: pulumi.Output<number>;

    /**
     * Create a BgpPeerTemplateMaxPrefix resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BgpPeerTemplateMaxPrefixArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BgpPeerTemplateMaxPrefixArgs | BgpPeerTemplateMaxPrefixState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BgpPeerTemplateMaxPrefixState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["addressFamily"] = state ? state.addressFamily : undefined;
            resourceInputs["asn"] = state ? state.asn : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["maximumPrefix"] = state ? state.maximumPrefix : undefined;
            resourceInputs["restartTime"] = state ? state.restartTime : undefined;
            resourceInputs["templateName"] = state ? state.templateName : undefined;
            resourceInputs["threshold"] = state ? state.threshold : undefined;
        } else {
            const args = argsOrState as BgpPeerTemplateMaxPrefixArgs | undefined;
            if ((!args || args.addressFamily === undefined) && !opts.urn) {
                throw new Error("Missing required property 'addressFamily'");
            }
            if ((!args || args.asn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'asn'");
            }
            if ((!args || args.templateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateName'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["addressFamily"] = args ? args.addressFamily : undefined;
            resourceInputs["asn"] = args ? args.asn : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["maximumPrefix"] = args ? args.maximumPrefix : undefined;
            resourceInputs["restartTime"] = args ? args.restartTime : undefined;
            resourceInputs["templateName"] = args ? args.templateName : undefined;
            resourceInputs["threshold"] = args ? args.threshold : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BgpPeerTemplateMaxPrefix.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BgpPeerTemplateMaxPrefix resources.
 */
export interface BgpPeerTemplateMaxPrefixState {
    /**
     * Action to do when limit is exceeded. - Choices: `log`, `shut`, `restart` - Default value: `shut`
     */
    action?: pulumi.Input<string>;
    /**
     * Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
     * `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
     * `ipv4-mdt` - Default value: `ipv4-ucast`
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
     * Maximum number of prefixes allowed from the peer. - Range: `0`-`4294967295`
     */
    maximumPrefix?: pulumi.Input<number>;
    /**
     * The period of time in minutes before restarting the peer when the prefix limit is reached. - Default value: `0`
     */
    restartTime?: pulumi.Input<number>;
    /**
     * Peer template name.
     */
    templateName?: pulumi.Input<string>;
    /**
     * The period of time in minutes before restarting the peer when the prefix limit is reached. - Range: `0`-`100` - Default
     * value: `0`
     */
    threshold?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a BgpPeerTemplateMaxPrefix resource.
 */
export interface BgpPeerTemplateMaxPrefixArgs {
    /**
     * Action to do when limit is exceeded. - Choices: `log`, `shut`, `restart` - Default value: `shut`
     */
    action?: pulumi.Input<string>;
    /**
     * Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
     * `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
     * `ipv4-mdt` - Default value: `ipv4-ucast`
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
     * Maximum number of prefixes allowed from the peer. - Range: `0`-`4294967295`
     */
    maximumPrefix?: pulumi.Input<number>;
    /**
     * The period of time in minutes before restarting the peer when the prefix limit is reached. - Default value: `0`
     */
    restartTime?: pulumi.Input<number>;
    /**
     * Peer template name.
     */
    templateName: pulumi.Input<string>;
    /**
     * The period of time in minutes before restarting the peer when the prefix limit is reached. - Range: `0`-`100` - Default
     * value: `0`
     */
    threshold?: pulumi.Input<number>;
}
