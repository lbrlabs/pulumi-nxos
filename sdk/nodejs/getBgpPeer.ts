// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the BGP peer configuration.
 *
 * - API Documentation: [bgpPeer](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Peer/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getBgpPeer({
 *     address: "192.168.0.1",
 *     asn: "65001",
 *     vrf: "default",
 * });
 * ```
 */
export function getBgpPeer(args: GetBgpPeerArgs, opts?: pulumi.InvokeOptions): Promise<GetBgpPeerResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getBgpPeer:getBgpPeer", {
        "address": args.address,
        "asn": args.asn,
        "device": args.device,
        "vrf": args.vrf,
    }, opts);
}

/**
 * A collection of arguments for invoking getBgpPeer.
 */
export interface GetBgpPeerArgs {
    /**
     * Peer address.
     */
    address: string;
    /**
     * Autonomous system number.
     */
    asn: string;
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * VRF name.
     */
    vrf: string;
}

/**
 * A collection of values returned by getBgpPeer.
 */
export interface GetBgpPeerResult {
    /**
     * Peer address.
     */
    readonly address: string;
    /**
     * Autonomous system number.
     */
    readonly asn: string;
    /**
     * Peer description.
     */
    readonly description: string;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * BGP Hold Timer in seconds. The value must be greater than the keepalive timer
     */
    readonly holdTime: number;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * BGP Keepalive Timer in seconds
     */
    readonly keepalive: number;
    /**
     * Peer template name.
     */
    readonly peerTemplate: string;
    /**
     * Neighbor Fabric Type.
     */
    readonly peerType: string;
    /**
     * Peer autonomous system number.
     */
    readonly remoteAsn: string;
    /**
     * Source Interface. Must match first field in the output of `show intf brief`.
     */
    readonly sourceInterface: string;
    /**
     * VRF name.
     */
    readonly vrf: string;
}
/**
 * This data source can read the BGP peer configuration.
 *
 * - API Documentation: [bgpPeer](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Peer/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getBgpPeer({
 *     address: "192.168.0.1",
 *     asn: "65001",
 *     vrf: "default",
 * });
 * ```
 */
export function getBgpPeerOutput(args: GetBgpPeerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBgpPeerResult> {
    return pulumi.output(args).apply((a: any) => getBgpPeer(a, opts))
}

/**
 * A collection of arguments for invoking getBgpPeer.
 */
export interface GetBgpPeerOutputArgs {
    /**
     * Peer address.
     */
    address: pulumi.Input<string>;
    /**
     * Autonomous system number.
     */
    asn: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrf: pulumi.Input<string>;
}