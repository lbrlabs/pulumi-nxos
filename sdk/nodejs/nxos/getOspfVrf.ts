// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getOspfVrf(args: GetOspfVrfArgs, opts?: pulumi.InvokeOptions): Promise<GetOspfVrfResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getOspfVrf:getOspfVrf", {
        "device": args.device,
        "instanceName": args.instanceName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getOspfVrf.
 */
export interface GetOspfVrfArgs {
    device?: string;
    instanceName: string;
    name: string;
}

/**
 * A collection of values returned by getOspfVrf.
 */
export interface GetOspfVrfResult {
    readonly adminState: string;
    readonly bandwidthReference: number;
    readonly banwidthReferenceUnit: string;
    readonly device?: string;
    readonly distance: number;
    readonly id: string;
    readonly instanceName: string;
    readonly name: string;
    readonly routerId: string;
}
export function getOspfVrfOutput(args: GetOspfVrfOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOspfVrfResult> {
    return pulumi.output(args).apply((a: any) => getOspfVrf(a, opts))
}

/**
 * A collection of arguments for invoking getOspfVrf.
 */
export interface GetOspfVrfOutputArgs {
    device?: pulumi.Input<string>;
    instanceName: pulumi.Input<string>;
    name: pulumi.Input<string>;
}