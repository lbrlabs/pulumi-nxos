// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getPortChannelInterfaceMember(args: GetPortChannelInterfaceMemberArgs, opts?: pulumi.InvokeOptions): Promise<GetPortChannelInterfaceMemberResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getPortChannelInterfaceMember:getPortChannelInterfaceMember", {
        "device": args.device,
        "interfaceDn": args.interfaceDn,
        "interfaceId": args.interfaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getPortChannelInterfaceMember.
 */
export interface GetPortChannelInterfaceMemberArgs {
    device?: string;
    interfaceDn: string;
    interfaceId: string;
}

/**
 * A collection of values returned by getPortChannelInterfaceMember.
 */
export interface GetPortChannelInterfaceMemberResult {
    readonly device?: string;
    readonly id: string;
    readonly interfaceDn: string;
    readonly interfaceId: string;
}
export function getPortChannelInterfaceMemberOutput(args: GetPortChannelInterfaceMemberOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPortChannelInterfaceMemberResult> {
    return pulumi.output(args).apply((a: any) => getPortChannelInterfaceMember(a, opts))
}

/**
 * A collection of arguments for invoking getPortChannelInterfaceMember.
 */
export interface GetPortChannelInterfaceMemberOutputArgs {
    device?: pulumi.Input<string>;
    interfaceDn: pulumi.Input<string>;
    interfaceId: pulumi.Input<string>;
}