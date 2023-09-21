// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getBgp(args?: GetBgpArgs, opts?: pulumi.InvokeOptions): Promise<GetBgpResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getBgp:getBgp", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getBgp.
 */
export interface GetBgpArgs {
    device?: string;
}

/**
 * A collection of values returned by getBgp.
 */
export interface GetBgpResult {
    readonly adminState: string;
    readonly device?: string;
    readonly id: string;
}
export function getBgpOutput(args?: GetBgpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBgpResult> {
    return pulumi.output(args).apply((a: any) => getBgp(a, opts))
}

/**
 * A collection of arguments for invoking getBgp.
 */
export interface GetBgpOutputArgs {
    device?: pulumi.Input<string>;
}