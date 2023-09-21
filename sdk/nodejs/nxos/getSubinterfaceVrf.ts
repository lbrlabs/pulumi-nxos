// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getSubinterfaceVrf(args: GetSubinterfaceVrfArgs, opts?: pulumi.InvokeOptions): Promise<GetSubinterfaceVrfResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getSubinterfaceVrf:getSubinterfaceVrf", {
        "device": args.device,
        "interfaceId": args.interfaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubinterfaceVrf.
 */
export interface GetSubinterfaceVrfArgs {
    device?: string;
    interfaceId: string;
}

/**
 * A collection of values returned by getSubinterfaceVrf.
 */
export interface GetSubinterfaceVrfResult {
    readonly device?: string;
    readonly id: string;
    readonly interfaceId: string;
    readonly vrfDn: string;
}
export function getSubinterfaceVrfOutput(args: GetSubinterfaceVrfOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubinterfaceVrfResult> {
    return pulumi.output(args).apply((a: any) => getSubinterfaceVrf(a, opts))
}

/**
 * A collection of arguments for invoking getSubinterfaceVrf.
 */
export interface GetSubinterfaceVrfOutputArgs {
    device?: pulumi.Input<string>;
    interfaceId: pulumi.Input<string>;
}
