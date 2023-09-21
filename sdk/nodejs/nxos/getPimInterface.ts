// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getPimInterface(args: GetPimInterfaceArgs, opts?: pulumi.InvokeOptions): Promise<GetPimInterfaceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getPimInterface:getPimInterface", {
        "device": args.device,
        "interfaceId": args.interfaceId,
        "vrfName": args.vrfName,
    }, opts);
}

/**
 * A collection of arguments for invoking getPimInterface.
 */
export interface GetPimInterfaceArgs {
    device?: string;
    interfaceId: string;
    vrfName: string;
}

/**
 * A collection of values returned by getPimInterface.
 */
export interface GetPimInterfaceResult {
    readonly adminState: string;
    readonly bfd: string;
    readonly device?: string;
    readonly drPriority: number;
    readonly id: string;
    readonly interfaceId: string;
    readonly passive: boolean;
    readonly sparseMode: boolean;
    readonly vrfName: string;
}
export function getPimInterfaceOutput(args: GetPimInterfaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPimInterfaceResult> {
    return pulumi.output(args).apply((a: any) => getPimInterface(a, opts))
}

/**
 * A collection of arguments for invoking getPimInterface.
 */
export interface GetPimInterfaceOutputArgs {
    device?: pulumi.Input<string>;
    interfaceId: pulumi.Input<string>;
    vrfName: pulumi.Input<string>;
}
