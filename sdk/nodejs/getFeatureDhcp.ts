// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the DHCP feature configuration.
 *
 * - API Documentation: [fmDhcp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Dhcp/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureDhcp({});
 * ```
 */
export function getFeatureDhcp(args?: GetFeatureDhcpArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureDhcpResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getFeatureDhcp:getFeatureDhcp", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeatureDhcp.
 */
export interface GetFeatureDhcpArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
}

/**
 * A collection of values returned by getFeatureDhcp.
 */
export interface GetFeatureDhcpResult {
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
 * This data source can read the DHCP feature configuration.
 *
 * - API Documentation: [fmDhcp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Dhcp/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureDhcp({});
 * ```
 */
export function getFeatureDhcpOutput(args?: GetFeatureDhcpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureDhcpResult> {
    return pulumi.output(args).apply((a: any) => getFeatureDhcp(a, opts))
}

/**
 * A collection of arguments for invoking getFeatureDhcp.
 */
export interface GetFeatureDhcpOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}
