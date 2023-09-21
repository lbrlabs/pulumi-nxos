// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getSviInterfaceVrf(args: GetSviInterfaceVrfArgs, opts?: pulumi.InvokeOptions): Promise<GetSviInterfaceVrfResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getSviInterfaceVrf:getSviInterfaceVrf", {
        "device": args.device,
        "interfaceId": args.interfaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSviInterfaceVrf.
 */
export interface GetSviInterfaceVrfArgs {
    device?: string;
    interfaceId: string;
}

/**
 * A collection of values returned by getSviInterfaceVrf.
 */
export interface GetSviInterfaceVrfResult {
    readonly device?: string;
    readonly id: string;
    readonly interfaceId: string;
    readonly vrfDn: string;
}
export function getSviInterfaceVrfOutput(args: GetSviInterfaceVrfOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSviInterfaceVrfResult> {
    return pulumi.output(args).apply((a: any) => getSviInterfaceVrf(a, opts))
}

/**
 * A collection of arguments for invoking getSviInterfaceVrf.
 */
export interface GetSviInterfaceVrfOutputArgs {
    device?: pulumi.Input<string>;
    interfaceId: pulumi.Input<string>;
}
