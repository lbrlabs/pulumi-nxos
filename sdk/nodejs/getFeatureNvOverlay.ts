// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the NV Overlay feature configuration.
 *
 * - API Documentation: [fmNvo](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Nvo/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureNvOverlay({});
 * ```
 */
export function getFeatureNvOverlay(args?: GetFeatureNvOverlayArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureNvOverlayResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getFeatureNvOverlay:getFeatureNvOverlay", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeatureNvOverlay.
 */
export interface GetFeatureNvOverlayArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
}

/**
 * A collection of values returned by getFeatureNvOverlay.
 */
export interface GetFeatureNvOverlayResult {
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
 * This data source can read the NV Overlay feature configuration.
 *
 * - API Documentation: [fmNvo](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Nvo/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureNvOverlay({});
 * ```
 */
export function getFeatureNvOverlayOutput(args?: GetFeatureNvOverlayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureNvOverlayResult> {
    return pulumi.output(args).apply((a: any) => getFeatureNvOverlay(a, opts))
}

/**
 * A collection of arguments for invoking getFeatureNvOverlay.
 */
export interface GetFeatureNvOverlayOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}