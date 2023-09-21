// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getRouteMapRuleEntry(args: GetRouteMapRuleEntryArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteMapRuleEntryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:nxos/getRouteMapRuleEntry:getRouteMapRuleEntry", {
        "device": args.device,
        "order": args.order,
        "ruleName": args.ruleName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteMapRuleEntry.
 */
export interface GetRouteMapRuleEntryArgs {
    device?: string;
    order: number;
    ruleName: string;
}

/**
 * A collection of values returned by getRouteMapRuleEntry.
 */
export interface GetRouteMapRuleEntryResult {
    readonly action: string;
    readonly device?: string;
    readonly id: string;
    readonly order: number;
    readonly ruleName: string;
}
export function getRouteMapRuleEntryOutput(args: GetRouteMapRuleEntryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouteMapRuleEntryResult> {
    return pulumi.output(args).apply((a: any) => getRouteMapRuleEntry(a, opts))
}

/**
 * A collection of arguments for invoking getRouteMapRuleEntry.
 */
export interface GetRouteMapRuleEntryOutputArgs {
    device?: pulumi.Input<string>;
    order: pulumi.Input<number>;
    ruleName: pulumi.Input<string>;
}