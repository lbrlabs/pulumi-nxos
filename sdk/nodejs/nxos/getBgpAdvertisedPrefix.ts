// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getBgpAdvertisedPrefix(args: GetBgpAdvertisedPrefixArgs, opts?: pulumi.InvokeOptions): Promise<GetBgpAdvertisedPrefixResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getBgpAdvertisedPrefix:getBgpAdvertisedPrefix", {
        "addressFamily": args.addressFamily,
        "asn": args.asn,
        "device": args.device,
        "prefix": args.prefix,
        "vrf": args.vrf,
    }, opts);
}

/**
 * A collection of arguments for invoking getBgpAdvertisedPrefix.
 */
export interface GetBgpAdvertisedPrefixArgs {
    addressFamily: string;
    asn: string;
    device?: string;
    prefix: string;
    vrf: string;
}

/**
 * A collection of values returned by getBgpAdvertisedPrefix.
 */
export interface GetBgpAdvertisedPrefixResult {
    readonly addressFamily: string;
    readonly asn: string;
    readonly device?: string;
    readonly id: string;
    readonly prefix: string;
    readonly routeMap: string;
    readonly vrf: string;
}
export function getBgpAdvertisedPrefixOutput(args: GetBgpAdvertisedPrefixOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBgpAdvertisedPrefixResult> {
    return pulumi.output(args).apply((a: any) => getBgpAdvertisedPrefix(a, opts))
}

/**
 * A collection of arguments for invoking getBgpAdvertisedPrefix.
 */
export interface GetBgpAdvertisedPrefixOutputArgs {
    addressFamily: pulumi.Input<string>;
    asn: pulumi.Input<string>;
    device?: pulumi.Input<string>;
    prefix: pulumi.Input<string>;
    vrf: pulumi.Input<string>;
}