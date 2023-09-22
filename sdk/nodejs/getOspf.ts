// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the global OSPF configuration.
 *
 * - API Documentation: [ospfEntity](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/ospf:Entity/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getOspf({});
 * ```
 */
export function getOspf(args?: GetOspfArgs, opts?: pulumi.InvokeOptions): Promise<GetOspfResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getOspf:getOspf", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getOspf.
 */
export interface GetOspfArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
}

/**
 * A collection of values returned by getOspf.
 */
export interface GetOspfResult {
    /**
     * Administrative state.
     */
    readonly adminState: string;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
}
/**
 * This data source can read the global OSPF configuration.
 *
 * - API Documentation: [ospfEntity](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/ospf:Entity/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getOspf({});
 * ```
 */
export function getOspfOutput(args?: GetOspfOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOspfResult> {
    return pulumi.output(args).apply((a: any) => getOspf(a, opts))
}

/**
 * A collection of arguments for invoking getOspf.
 */
export interface GetOspfOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}