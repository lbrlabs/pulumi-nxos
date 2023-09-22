// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read a Route-Map Rule Entry configuration.
 *
 * - API Documentation: [rtmapEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:Entry/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getRouteMapRuleEntry({
 *     order: 10,
 *     ruleName: "RULE1",
 * });
 * ```
 */
export function getRouteMapRuleEntry(args: GetRouteMapRuleEntryArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteMapRuleEntryResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getRouteMapRuleEntry:getRouteMapRuleEntry", {
        "device": args.device,
        "order": args.order,
        "ruleName": args.ruleName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteMapRuleEntry.
 */
export interface GetRouteMapRuleEntryArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * Route-Map Rule Entry order.
     */
    order: number;
    /**
     * Route Map rule name.
     */
    ruleName: string;
}

/**
 * A collection of values returned by getRouteMapRuleEntry.
 */
export interface GetRouteMapRuleEntryResult {
    /**
     * Route-Map Rule Entry action.
     */
    readonly action: string;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * Route-Map Rule Entry order.
     */
    readonly order: number;
    /**
     * Route Map rule name.
     */
    readonly ruleName: string;
}
/**
 * This data source can read a Route-Map Rule Entry configuration.
 *
 * - API Documentation: [rtmapEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:Entry/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getRouteMapRuleEntry({
 *     order: 10,
 *     ruleName: "RULE1",
 * });
 * ```
 */
export function getRouteMapRuleEntryOutput(args: GetRouteMapRuleEntryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouteMapRuleEntryResult> {
    return pulumi.output(args).apply((a: any) => getRouteMapRuleEntry(a, opts))
}

/**
 * A collection of arguments for invoking getRouteMapRuleEntry.
 */
export interface GetRouteMapRuleEntryOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Route-Map Rule Entry order.
     */
    order: pulumi.Input<number>;
    /**
     * Route Map rule name.
     */
    ruleName: pulumi.Input<string>;
}
