// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getBgpPeerAddressFamilyRouteControl(args: GetBgpPeerAddressFamilyRouteControlArgs, opts?: pulumi.InvokeOptions): Promise<GetBgpPeerAddressFamilyRouteControlResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getBgpPeerAddressFamilyRouteControl:getBgpPeerAddressFamilyRouteControl", {
        "address": args.address,
        "addressFamily": args.addressFamily,
        "asn": args.asn,
        "device": args.device,
        "direction": args.direction,
        "vrf": args.vrf,
    }, opts);
}

/**
 * A collection of arguments for invoking getBgpPeerAddressFamilyRouteControl.
 */
export interface GetBgpPeerAddressFamilyRouteControlArgs {
    address: string;
    addressFamily: string;
    asn: string;
    device?: string;
    direction: string;
    vrf: string;
}

/**
 * A collection of values returned by getBgpPeerAddressFamilyRouteControl.
 */
export interface GetBgpPeerAddressFamilyRouteControlResult {
    readonly address: string;
    readonly addressFamily: string;
    readonly asn: string;
    readonly device?: string;
    readonly direction: string;
    readonly id: string;
    readonly routeMapName: string;
    readonly vrf: string;
}
export function getBgpPeerAddressFamilyRouteControlOutput(args: GetBgpPeerAddressFamilyRouteControlOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBgpPeerAddressFamilyRouteControlResult> {
    return pulumi.output(args).apply((a: any) => getBgpPeerAddressFamilyRouteControl(a, opts))
}

/**
 * A collection of arguments for invoking getBgpPeerAddressFamilyRouteControl.
 */
export interface GetBgpPeerAddressFamilyRouteControlOutputArgs {
    address: pulumi.Input<string>;
    addressFamily: pulumi.Input<string>;
    asn: pulumi.Input<string>;
    device?: pulumi.Input<string>;
    direction: pulumi.Input<string>;
    vrf: pulumi.Input<string>;
}
