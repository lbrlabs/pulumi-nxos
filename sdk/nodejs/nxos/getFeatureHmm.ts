// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getFeatureHmm(args?: GetFeatureHmmArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureHmmResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getFeatureHmm:getFeatureHmm", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeatureHmm.
 */
export interface GetFeatureHmmArgs {
    device?: string;
}

/**
 * A collection of values returned by getFeatureHmm.
 */
export interface GetFeatureHmmResult {
    readonly adminState: string;
    readonly device?: string;
    readonly id: string;
}
export function getFeatureHmmOutput(args?: GetFeatureHmmOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureHmmResult> {
    return pulumi.output(args).apply((a: any) => getFeatureHmm(a, opts))
}

/**
 * A collection of arguments for invoking getFeatureHmm.
 */
export interface GetFeatureHmmOutputArgs {
    device?: pulumi.Input<string>;
}
