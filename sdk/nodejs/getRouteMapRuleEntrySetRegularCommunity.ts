// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read a Set Community configuration in a Route-Map Rule Entry.
 *
 * - API Documentation: [rtmapSetRegComm](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:SetRegComm/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getRouteMapRuleEntrySetRegularCommunity({
 *     order: 10,
 *     ruleName: "RULE1",
 * });
 * ```
 */
export function getRouteMapRuleEntrySetRegularCommunity(args: GetRouteMapRuleEntrySetRegularCommunityArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteMapRuleEntrySetRegularCommunityResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getRouteMapRuleEntrySetRegularCommunity:getRouteMapRuleEntrySetRegularCommunity", {
        "device": args.device,
        "order": args.order,
        "ruleName": args.ruleName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteMapRuleEntrySetRegularCommunity.
 */
export interface GetRouteMapRuleEntrySetRegularCommunityArgs {
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
 * A collection of values returned by getRouteMapRuleEntrySetRegularCommunity.
 */
export interface GetRouteMapRuleEntrySetRegularCommunityResult {
    /**
     * Option to add to an existing community.
     */
    readonly additive: string;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * Option to have no community attribute.
     */
    readonly noCommunity: string;
    /**
     * Route-Map Rule Entry order.
     */
    readonly order: number;
    /**
     * Route Map rule name.
     */
    readonly ruleName: string;
    /**
     * Operation on the current community.
     */
    readonly setCriteria: string;
}
/**
 * This data source can read a Set Community configuration in a Route-Map Rule Entry.
 *
 * - API Documentation: [rtmapSetRegComm](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:SetRegComm/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getRouteMapRuleEntrySetRegularCommunity({
 *     order: 10,
 *     ruleName: "RULE1",
 * });
 * ```
 */
export function getRouteMapRuleEntrySetRegularCommunityOutput(args: GetRouteMapRuleEntrySetRegularCommunityOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouteMapRuleEntrySetRegularCommunityResult> {
    return pulumi.output(args).apply((a: any) => getRouteMapRuleEntrySetRegularCommunity(a, opts))
}

/**
 * A collection of arguments for invoking getRouteMapRuleEntrySetRegularCommunity.
 */
export interface GetRouteMapRuleEntrySetRegularCommunityOutputArgs {
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
