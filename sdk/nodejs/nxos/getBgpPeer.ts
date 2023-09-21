// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getBgpPeer(args: GetBgpPeerArgs, opts?: pulumi.InvokeOptions): Promise<GetBgpPeerResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getBgpPeer:getBgpPeer", {
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
    address: string;
    asn: string;
    device?: string;
    vrf: string;
}

/**
 * A collection of values returned by getBgpPeer.
 */
export interface GetBgpPeerResult {
    readonly address: string;
    readonly asn: string;
    readonly description: string;
    readonly device?: string;
    readonly holdTime: number;
    readonly id: string;
    readonly keepalive: number;
    readonly peerTemplate: string;
    readonly peerType: string;
    readonly remoteAsn: string;
    readonly sourceInterface: string;
    readonly vrf: string;
}
export function getBgpPeerOutput(args: GetBgpPeerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBgpPeerResult> {
    return pulumi.output(args).apply((a: any) => getBgpPeer(a, opts))
}

/**
 * A collection of arguments for invoking getBgpPeer.
 */
export interface GetBgpPeerOutputArgs {
    address: pulumi.Input<string>;
    asn: pulumi.Input<string>;
    device?: pulumi.Input<string>;
    vrf: pulumi.Input<string>;
}
