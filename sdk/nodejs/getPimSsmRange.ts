// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the PIM SSM range configuration.
 *
 * - API Documentation: [pimSSMRangeP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:SSMRangeP/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getPimSsmRange({
 *     vrfName: "default",
 * });
 * ```
 */
export function getPimSsmRange(args: GetPimSsmRangeArgs, opts?: pulumi.InvokeOptions): Promise<GetPimSsmRangeResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getPimSsmRange:getPimSsmRange", {
        "device": args.device,
        "vrfName": args.vrfName,
    }, opts);
}

/**
 * A collection of arguments for invoking getPimSsmRange.
 */
export interface GetPimSsmRangeArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * VRF name.
     */
    vrfName: string;
}

/**
 * A collection of values returned by getPimSsmRange.
 */
export interface GetPimSsmRangeResult {
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * Group list 1.
     */
    readonly groupList1: string;
    /**
     * Group list 2.
     */
    readonly groupList2: string;
    /**
     * Group list 3.
     */
    readonly groupList3: string;
    /**
     * Group list 4.
     */
    readonly groupList4: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * Prefix list name.
     */
    readonly prefixList: string;
    /**
     * Route map name.
     */
    readonly routeMap: string;
    /**
     * Exclude standard SSM range (232.0.0.0/8).
     */
    readonly ssmNone: boolean;
    /**
     * VRF name.
     */
    readonly vrfName: string;
}
/**
 * This data source can read the PIM SSM range configuration.
 *
 * - API Documentation: [pimSSMRangeP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:SSMRangeP/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getPimSsmRange({
 *     vrfName: "default",
 * });
 * ```
 */
export function getPimSsmRangeOutput(args: GetPimSsmRangeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPimSsmRangeResult> {
    return pulumi.output(args).apply((a: any) => getPimSsmRange(a, opts))
}

/**
 * A collection of arguments for invoking getPimSsmRange.
 */
export interface GetPimSsmRangeOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrfName: pulumi.Input<string>;
}