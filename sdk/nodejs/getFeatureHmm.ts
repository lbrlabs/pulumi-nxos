// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the HMM feature (aka `feature fabric forwarding`) configuration.
 *
 * - API Documentation: [fmHmm](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Hmm/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureHmm({});
 * ```
 */
export function getFeatureHmm(args?: GetFeatureHmmArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureHmmResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getFeatureHmm:getFeatureHmm", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeatureHmm.
 */
export interface GetFeatureHmmArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
}

/**
 * A collection of values returned by getFeatureHmm.
 */
export interface GetFeatureHmmResult {
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
 * This data source can read the HMM feature (aka `feature fabric forwarding`) configuration.
 *
 * - API Documentation: [fmHmm](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Hmm/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureHmm({});
 * ```
 */
export function getFeatureHmmOutput(args?: GetFeatureHmmOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureHmmResult> {
    return pulumi.output(args).apply((a: any) => getFeatureHmm(a, opts))
}

/**
 * A collection of arguments for invoking getFeatureHmm.
 */
export interface GetFeatureHmmOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}
