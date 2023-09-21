// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getSubinterface(args: GetSubinterfaceArgs, opts?: pulumi.InvokeOptions): Promise<GetSubinterfaceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getSubinterface:getSubinterface", {
        "device": args.device,
        "interfaceId": args.interfaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubinterface.
 */
export interface GetSubinterfaceArgs {
    device?: string;
    interfaceId: string;
}

/**
 * A collection of values returned by getSubinterface.
 */
export interface GetSubinterfaceResult {
    readonly adminState: string;
    readonly bandwidth: number;
    readonly delay: number;
    readonly description: string;
    readonly device?: string;
    readonly encap: string;
    readonly id: string;
    readonly interfaceId: string;
    readonly linkLogging: string;
    readonly medium: string;
    readonly mtu: number;
}
export function getSubinterfaceOutput(args: GetSubinterfaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubinterfaceResult> {
    return pulumi.output(args).apply((a: any) => getSubinterface(a, opts))
}

/**
 * A collection of arguments for invoking getSubinterface.
 */
export interface GetSubinterfaceOutputArgs {
    device?: pulumi.Input<string>;
    interfaceId: pulumi.Input<string>;
}
