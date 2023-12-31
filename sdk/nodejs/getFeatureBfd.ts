// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the BFD feature configuration.
 *
 * - API Documentation: [fmBfd](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Bfd/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureBfd({});
 * ```
 */
export function getFeatureBfd(args?: GetFeatureBfdArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureBfdResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getFeatureBfd:getFeatureBfd", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeatureBfd.
 */
export interface GetFeatureBfdArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
}

/**
 * A collection of values returned by getFeatureBfd.
 */
export interface GetFeatureBfdResult {
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
 * This data source can read the BFD feature configuration.
 *
 * - API Documentation: [fmBfd](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Bfd/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureBfd({});
 * ```
 */
export function getFeatureBfdOutput(args?: GetFeatureBfdOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureBfdResult> {
    return pulumi.output(args).apply((a: any) => getFeatureBfd(a, opts))
}

/**
 * A collection of arguments for invoking getFeatureBfd.
 */
export interface GetFeatureBfdOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}
