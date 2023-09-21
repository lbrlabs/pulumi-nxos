// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getVrfAddressFamily(args: GetVrfAddressFamilyArgs, opts?: pulumi.InvokeOptions): Promise<GetVrfAddressFamilyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getVrfAddressFamily:getVrfAddressFamily", {
        "addressFamily": args.addressFamily,
        "device": args.device,
        "vrf": args.vrf,
    }, opts);
}

/**
 * A collection of arguments for invoking getVrfAddressFamily.
 */
export interface GetVrfAddressFamilyArgs {
    addressFamily: string;
    device?: string;
    vrf: string;
}

/**
 * A collection of values returned by getVrfAddressFamily.
 */
export interface GetVrfAddressFamilyResult {
    readonly addressFamily: string;
    readonly device?: string;
    readonly id: string;
    readonly vrf: string;
}
export function getVrfAddressFamilyOutput(args: GetVrfAddressFamilyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVrfAddressFamilyResult> {
    return pulumi.output(args).apply((a: any) => getVrfAddressFamily(a, opts))
}

/**
 * A collection of arguments for invoking getVrfAddressFamily.
 */
export interface GetVrfAddressFamilyOutputArgs {
    addressFamily: pulumi.Input<string>;
    device?: pulumi.Input<string>;
    vrf: pulumi.Input<string>;
}
