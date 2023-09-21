// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getPimSsmRange(args: GetPimSsmRangeArgs, opts?: pulumi.InvokeOptions): Promise<GetPimSsmRangeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getPimSsmRange:getPimSsmRange", {
        "device": args.device,
        "vrfName": args.vrfName,
    }, opts);
}

/**
 * A collection of arguments for invoking getPimSsmRange.
 */
export interface GetPimSsmRangeArgs {
    device?: string;
    vrfName: string;
}

/**
 * A collection of values returned by getPimSsmRange.
 */
export interface GetPimSsmRangeResult {
    readonly device?: string;
    readonly groupList1: string;
    readonly groupList2: string;
    readonly groupList3: string;
    readonly groupList4: string;
    readonly id: string;
    readonly prefixList: string;
    readonly routeMap: string;
    readonly ssmNone: boolean;
    readonly vrfName: string;
}
export function getPimSsmRangeOutput(args: GetPimSsmRangeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPimSsmRangeResult> {
    return pulumi.output(args).apply((a: any) => getPimSsmRange(a, opts))
}

/**
 * A collection of arguments for invoking getPimSsmRange.
 */
export interface GetPimSsmRangeOutputArgs {
    device?: pulumi.Input<string>;
    vrfName: pulumi.Input<string>;
}
