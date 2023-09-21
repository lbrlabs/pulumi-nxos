// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getNveInterface(args?: GetNveInterfaceArgs, opts?: pulumi.InvokeOptions): Promise<GetNveInterfaceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getNveInterface:getNveInterface", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getNveInterface.
 */
export interface GetNveInterfaceArgs {
    device?: string;
}

/**
 * A collection of values returned by getNveInterface.
 */
export interface GetNveInterfaceResult {
    readonly adminState: string;
    readonly advertiseVirtualMac: boolean;
    readonly device?: string;
    readonly holdDownTime: number;
    readonly hostReachabilityProtocol: string;
    readonly id: string;
    readonly ingressReplicationProtocolBgp: boolean;
    readonly multicastGroupL2: string;
    readonly multicastGroupL3: string;
    readonly multisiteSourceInterface: string;
    readonly sourceInterface: string;
    readonly suppressArp: boolean;
    readonly suppressMacRoute: boolean;
}
export function getNveInterfaceOutput(args?: GetNveInterfaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNveInterfaceResult> {
    return pulumi.output(args).apply((a: any) => getNveInterface(a, opts))
}

/**
 * A collection of arguments for invoking getNveInterface.
 */
export interface GetNveInterfaceOutputArgs {
    device?: pulumi.Input<string>;
}