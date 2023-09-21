// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getPimVrf(args: GetPimVrfArgs, opts?: pulumi.InvokeOptions): Promise<GetPimVrfResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getPimVrf:getPimVrf", {
        "device": args.device,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getPimVrf.
 */
export interface GetPimVrfArgs {
    device?: string;
    name: string;
}

/**
 * A collection of values returned by getPimVrf.
 */
export interface GetPimVrfResult {
    readonly adminState: string;
    readonly bfd: boolean;
    readonly device?: string;
    readonly id: string;
    readonly name: string;
}
export function getPimVrfOutput(args: GetPimVrfOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPimVrfResult> {
    return pulumi.output(args).apply((a: any) => getPimVrf(a, opts))
}

/**
 * A collection of arguments for invoking getPimVrf.
 */
export interface GetPimVrfOutputArgs {
    device?: pulumi.Input<string>;
    name: pulumi.Input<string>;
}
