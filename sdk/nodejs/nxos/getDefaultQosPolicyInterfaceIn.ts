// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getDefaultQosPolicyInterfaceIn(args: GetDefaultQosPolicyInterfaceInArgs, opts?: pulumi.InvokeOptions): Promise<GetDefaultQosPolicyInterfaceInResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getDefaultQosPolicyInterfaceIn:getDefaultQosPolicyInterfaceIn", {
        "device": args.device,
        "interfaceId": args.interfaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getDefaultQosPolicyInterfaceIn.
 */
export interface GetDefaultQosPolicyInterfaceInArgs {
    device?: string;
    interfaceId: string;
}

/**
 * A collection of values returned by getDefaultQosPolicyInterfaceIn.
 */
export interface GetDefaultQosPolicyInterfaceInResult {
    readonly device?: string;
    readonly id: string;
    readonly interfaceId: string;
}
export function getDefaultQosPolicyInterfaceInOutput(args: GetDefaultQosPolicyInterfaceInOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDefaultQosPolicyInterfaceInResult> {
    return pulumi.output(args).apply((a: any) => getDefaultQosPolicyInterfaceIn(a, opts))
}

/**
 * A collection of arguments for invoking getDefaultQosPolicyInterfaceIn.
 */
export interface GetDefaultQosPolicyInterfaceInOutputArgs {
    device?: pulumi.Input<string>;
    interfaceId: pulumi.Input<string>;
}
