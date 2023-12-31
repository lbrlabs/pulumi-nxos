// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the HSRP feature configuration.
 *
 * - API Documentation: [fmHsrp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Hsrp/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureHsrp({});
 * ```
 */
export function getFeatureHsrp(args?: GetFeatureHsrpArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureHsrpResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getFeatureHsrp:getFeatureHsrp", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeatureHsrp.
 */
export interface GetFeatureHsrpArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
}

/**
 * A collection of values returned by getFeatureHsrp.
 */
export interface GetFeatureHsrpResult {
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
 * This data source can read the HSRP feature configuration.
 *
 * - API Documentation: [fmHsrp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Hsrp/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureHsrp({});
 * ```
 */
export function getFeatureHsrpOutput(args?: GetFeatureHsrpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureHsrpResult> {
    return pulumi.output(args).apply((a: any) => getFeatureHsrp(a, opts))
}

/**
 * A collection of arguments for invoking getFeatureHsrp.
 */
export interface GetFeatureHsrpOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}
