// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getRest(args: GetRestArgs, opts?: pulumi.InvokeOptions): Promise<GetRestResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getRest:getRest", {
        "device": args.device,
        "dn": args.dn,
    }, opts);
}

/**
 * A collection of arguments for invoking getRest.
 */
export interface GetRestArgs {
    device?: string;
    dn: string;
}

/**
 * A collection of values returned by getRest.
 */
export interface GetRestResult {
    readonly className: string;
    readonly content: {[key: string]: string};
    readonly device?: string;
    readonly dn: string;
    readonly id: string;
}
export function getRestOutput(args: GetRestOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRestResult> {
    return pulumi.output(args).apply((a: any) => getRest(a, opts))
}

/**
 * A collection of arguments for invoking getRest.
 */
export interface GetRestOutputArgs {
    device?: pulumi.Input<string>;
    dn: pulumi.Input<string>;
}
