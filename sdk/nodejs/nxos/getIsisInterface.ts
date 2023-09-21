// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getIsisInterface(args: GetIsisInterfaceArgs, opts?: pulumi.InvokeOptions): Promise<GetIsisInterfaceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getIsisInterface:getIsisInterface", {
        "device": args.device,
        "interfaceId": args.interfaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIsisInterface.
 */
export interface GetIsisInterfaceArgs {
    device?: string;
    interfaceId: string;
}

/**
 * A collection of values returned by getIsisInterface.
 */
export interface GetIsisInterfaceResult {
    readonly authenticationCheck: boolean;
    readonly authenticationCheckL1: boolean;
    readonly authenticationCheckL2: boolean;
    readonly authenticationKey: string;
    readonly authenticationKeyL1: string;
    readonly authenticationKeyL2: string;
    readonly authenticationType: string;
    readonly authenticationTypeL1: string;
    readonly authenticationTypeL2: string;
    readonly circuitType: string;
    readonly device?: string;
    readonly helloInterval: number;
    readonly helloIntervalL1: number;
    readonly helloIntervalL2: number;
    readonly helloMultiplier: number;
    readonly helloMultiplierL1: number;
    readonly helloMultiplierL2: number;
    readonly helloPadding: string;
    readonly id: string;
    readonly interfaceId: string;
    readonly metricL1: number;
    readonly metricL2: number;
    readonly mtuCheck: boolean;
    readonly mtuCheckL1: boolean;
    readonly mtuCheckL2: boolean;
    readonly networkTypeP2p: string;
    readonly passive: string;
    readonly priorityL1: number;
    readonly priorityL2: number;
    readonly vrf: string;
}
export function getIsisInterfaceOutput(args: GetIsisInterfaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIsisInterfaceResult> {
    return pulumi.output(args).apply((a: any) => getIsisInterface(a, opts))
}

/**
 * A collection of arguments for invoking getIsisInterface.
 */
export interface GetIsisInterfaceOutputArgs {
    device?: pulumi.Input<string>;
    interfaceId: pulumi.Input<string>;
}