// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the VN Segment feature configuration.
 *
 * - API Documentation: [fmVnSegment](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:VnSegment/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureVnSegment({});
 * ```
 */
export function getFeatureVnSegment(args?: GetFeatureVnSegmentArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureVnSegmentResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getFeatureVnSegment:getFeatureVnSegment", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeatureVnSegment.
 */
export interface GetFeatureVnSegmentArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
}

/**
 * A collection of values returned by getFeatureVnSegment.
 */
export interface GetFeatureVnSegmentResult {
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
 * This data source can read the VN Segment feature configuration.
 *
 * - API Documentation: [fmVnSegment](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:VnSegment/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureVnSegment({});
 * ```
 */
export function getFeatureVnSegmentOutput(args?: GetFeatureVnSegmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureVnSegmentResult> {
    return pulumi.output(args).apply((a: any) => getFeatureVnSegment(a, opts))
}

/**
 * A collection of arguments for invoking getFeatureVnSegment.
 */
export interface GetFeatureVnSegmentOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}