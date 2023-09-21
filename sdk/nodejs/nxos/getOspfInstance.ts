// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getOspfInstance(args: GetOspfInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetOspfInstanceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getOspfInstance:getOspfInstance", {
        "device": args.device,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getOspfInstance.
 */
export interface GetOspfInstanceArgs {
    device?: string;
    name: string;
}

/**
 * A collection of values returned by getOspfInstance.
 */
export interface GetOspfInstanceResult {
    readonly adminState: string;
    readonly device?: string;
    readonly id: string;
    readonly name: string;
}
export function getOspfInstanceOutput(args: GetOspfInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOspfInstanceResult> {
    return pulumi.output(args).apply((a: any) => getOspfInstance(a, opts))
}

/**
 * A collection of arguments for invoking getOspfInstance.
 */
export interface GetOspfInstanceOutputArgs {
    device?: pulumi.Input<string>;
    name: pulumi.Input<string>;
}