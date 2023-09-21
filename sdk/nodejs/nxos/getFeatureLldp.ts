// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getFeatureLldp(args?: GetFeatureLldpArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureLldpResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getFeatureLldp:getFeatureLldp", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeatureLldp.
 */
export interface GetFeatureLldpArgs {
    device?: string;
}

/**
 * A collection of values returned by getFeatureLldp.
 */
export interface GetFeatureLldpResult {
    readonly adminState: string;
    readonly device?: string;
    readonly id: string;
}
export function getFeatureLldpOutput(args?: GetFeatureLldpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureLldpResult> {
    return pulumi.output(args).apply((a: any) => getFeatureLldp(a, opts))
}

/**
 * A collection of arguments for invoking getFeatureLldp.
 */
export interface GetFeatureLldpOutputArgs {
    device?: pulumi.Input<string>;
}