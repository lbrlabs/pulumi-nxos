// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getFeatureOspfv3(args?: GetFeatureOspfv3Args, opts?: pulumi.InvokeOptions): Promise<GetFeatureOspfv3Result> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getFeatureOspfv3:getFeatureOspfv3", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeatureOspfv3.
 */
export interface GetFeatureOspfv3Args {
    device?: string;
}

/**
 * A collection of values returned by getFeatureOspfv3.
 */
export interface GetFeatureOspfv3Result {
    readonly adminState: string;
    readonly device?: string;
    readonly id: string;
}
export function getFeatureOspfv3Output(args?: GetFeatureOspfv3OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureOspfv3Result> {
    return pulumi.output(args).apply((a: any) => getFeatureOspfv3(a, opts))
}

/**
 * A collection of arguments for invoking getFeatureOspfv3.
 */
export interface GetFeatureOspfv3OutputArgs {
    device?: pulumi.Input<string>;
}
