// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the OSPFv3 feature configuration.
 *
 * - API Documentation: [fmOspfv3](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Ospfv3/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureOspfv3({});
 * ```
 */
export function getFeatureOspfv3(args?: GetFeatureOspfv3Args, opts?: pulumi.InvokeOptions): Promise<GetFeatureOspfv3Result> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getFeatureOspfv3:getFeatureOspfv3", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeatureOspfv3.
 */
export interface GetFeatureOspfv3Args {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
}

/**
 * A collection of values returned by getFeatureOspfv3.
 */
export interface GetFeatureOspfv3Result {
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
 * This data source can read the OSPFv3 feature configuration.
 *
 * - API Documentation: [fmOspfv3](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Ospfv3/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureOspfv3({});
 * ```
 */
export function getFeatureOspfv3Output(args?: GetFeatureOspfv3OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureOspfv3Result> {
    return pulumi.output(args).apply((a: any) => getFeatureOspfv3(a, opts))
}

/**
 * A collection of arguments for invoking getFeatureOspfv3.
 */
export interface GetFeatureOspfv3OutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}
