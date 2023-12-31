// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read a Match Route Prefix List in Route-Map Rule Entry configuration.
 *
 * - API Documentation: [rtmapRsRtDstAtt](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:RsRtDstAtt/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getRouteMapRuleEntryMatchRoutePrefixList({
 *     order: 10,
 *     prefixListDn: "sys/rpm/pfxlistv4-[LIST1]",
 *     ruleName: "RULE1",
 * });
 * ```
 */
export function getRouteMapRuleEntryMatchRoutePrefixList(args: GetRouteMapRuleEntryMatchRoutePrefixListArgs, opts?: pulumi.InvokeOptions): Promise<GetRouteMapRuleEntryMatchRoutePrefixListResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getRouteMapRuleEntryMatchRoutePrefixList:getRouteMapRuleEntryMatchRoutePrefixList", {
        "device": args.device,
        "order": args.order,
        "prefixListDn": args.prefixListDn,
        "ruleName": args.ruleName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouteMapRuleEntryMatchRoutePrefixList.
 */
export interface GetRouteMapRuleEntryMatchRoutePrefixListArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * Route-Map Rule Entry order.
     */
    order: number;
    /**
     * DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
     */
    prefixListDn: string;
    /**
     * Route Map rule name.
     */
    ruleName: string;
}

/**
 * A collection of values returned by getRouteMapRuleEntryMatchRoutePrefixList.
 */
export interface GetRouteMapRuleEntryMatchRoutePrefixListResult {
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
     * DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
     */
    readonly prefixListDn: string;
    /**
     * Route Map rule name.
     */
    readonly ruleName: string;
}
/**
 * This data source can read a Match Route Prefix List in Route-Map Rule Entry configuration.
 *
 * - API Documentation: [rtmapRsRtDstAtt](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:RsRtDstAtt/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getRouteMapRuleEntryMatchRoutePrefixList({
 *     order: 10,
 *     prefixListDn: "sys/rpm/pfxlistv4-[LIST1]",
 *     ruleName: "RULE1",
 * });
 * ```
 */
export function getRouteMapRuleEntryMatchRoutePrefixListOutput(args: GetRouteMapRuleEntryMatchRoutePrefixListOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouteMapRuleEntryMatchRoutePrefixListResult> {
    return pulumi.output(args).apply((a: any) => getRouteMapRuleEntryMatchRoutePrefixList(a, opts))
}

/**
 * A collection of arguments for invoking getRouteMapRuleEntryMatchRoutePrefixList.
 */
export interface GetRouteMapRuleEntryMatchRoutePrefixListOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Route-Map Rule Entry order.
     */
    order: pulumi.Input<number>;
    /**
     * DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
     */
    prefixListDn: pulumi.Input<string>;
    /**
     * Route Map rule name.
     */
    ruleName: pulumi.Input<string>;
}
