// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getFeatureVpc(args?: GetFeatureVpcArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureVpcResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getFeatureVpc:getFeatureVpc", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeatureVpc.
 */
export interface GetFeatureVpcArgs {
    device?: string;
}

/**
 * A collection of values returned by getFeatureVpc.
 */
export interface GetFeatureVpcResult {
    readonly adminState: string;
    readonly device?: string;
    readonly id: string;
}
export function getFeatureVpcOutput(args?: GetFeatureVpcOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureVpcResult> {
    return pulumi.output(args).apply((a: any) => getFeatureVpc(a, opts))
}

/**
 * A collection of arguments for invoking getFeatureVpc.
 */
export interface GetFeatureVpcOutputArgs {
    device?: pulumi.Input<string>;
}
