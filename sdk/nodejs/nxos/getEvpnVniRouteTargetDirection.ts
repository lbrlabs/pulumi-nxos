// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getEvpnVniRouteTargetDirection(args: GetEvpnVniRouteTargetDirectionArgs, opts?: pulumi.InvokeOptions): Promise<GetEvpnVniRouteTargetDirectionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getEvpnVniRouteTargetDirection:getEvpnVniRouteTargetDirection", {
        "device": args.device,
        "direction": args.direction,
        "encap": args.encap,
    }, opts);
}

/**
 * A collection of arguments for invoking getEvpnVniRouteTargetDirection.
 */
export interface GetEvpnVniRouteTargetDirectionArgs {
    device?: string;
    direction: string;
    encap: string;
}

/**
 * A collection of values returned by getEvpnVniRouteTargetDirection.
 */
export interface GetEvpnVniRouteTargetDirectionResult {
    readonly device?: string;
    readonly direction: string;
    readonly encap: string;
    readonly id: string;
}
export function getEvpnVniRouteTargetDirectionOutput(args: GetEvpnVniRouteTargetDirectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEvpnVniRouteTargetDirectionResult> {
    return pulumi.output(args).apply((a: any) => getEvpnVniRouteTargetDirection(a, opts))
}

/**
 * A collection of arguments for invoking getEvpnVniRouteTargetDirection.
 */
export interface GetEvpnVniRouteTargetDirectionOutputArgs {
    device?: pulumi.Input<string>;
    direction: pulumi.Input<string>;
    encap: pulumi.Input<string>;
}
