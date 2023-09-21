// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getFeatureNvOverlay(args?: GetFeatureNvOverlayArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureNvOverlayResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getFeatureNvOverlay:getFeatureNvOverlay", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeatureNvOverlay.
 */
export interface GetFeatureNvOverlayArgs {
    device?: string;
}

/**
 * A collection of values returned by getFeatureNvOverlay.
 */
export interface GetFeatureNvOverlayResult {
    readonly adminState: string;
    readonly device?: string;
    readonly id: string;
}
export function getFeatureNvOverlayOutput(args?: GetFeatureNvOverlayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureNvOverlayResult> {
    return pulumi.output(args).apply((a: any) => getFeatureNvOverlay(a, opts))
}

/**
 * A collection of arguments for invoking getFeatureNvOverlay.
 */
export interface GetFeatureNvOverlayOutputArgs {
    device?: pulumi.Input<string>;
}
