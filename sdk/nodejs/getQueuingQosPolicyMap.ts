// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the queuing QoS policy map configuration.
 *
 * - API Documentation: [ipqosPMapInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:PMapInst/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getQueuingQosPolicyMap({
 *     name: "PM1",
 * });
 * ```
 */
export function getQueuingQosPolicyMap(args: GetQueuingQosPolicyMapArgs, opts?: pulumi.InvokeOptions): Promise<GetQueuingQosPolicyMapResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getQueuingQosPolicyMap:getQueuingQosPolicyMap", {
        "device": args.device,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getQueuingQosPolicyMap.
 */
export interface GetQueuingQosPolicyMapArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * Policy map name.
     */
    name: string;
}

/**
 * A collection of values returned by getQueuingQosPolicyMap.
 */
export interface GetQueuingQosPolicyMapResult {
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * Match type.
     */
    readonly matchType: string;
    /**
     * Policy map name.
     */
    readonly name: string;
}
/**
 * This data source can read the queuing QoS policy map configuration.
 *
 * - API Documentation: [ipqosPMapInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:PMapInst/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getQueuingQosPolicyMap({
 *     name: "PM1",
 * });
 * ```
 */
export function getQueuingQosPolicyMapOutput(args: GetQueuingQosPolicyMapOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetQueuingQosPolicyMapResult> {
    return pulumi.output(args).apply((a: any) => getQueuingQosPolicyMap(a, opts))
}

/**
 * A collection of arguments for invoking getQueuingQosPolicyMap.
 */
export interface GetQueuingQosPolicyMapOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Policy map name.
     */
    name: pulumi.Input<string>;
}