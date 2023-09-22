// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the system configuration.
 *
 * - API Documentation: [topSystem](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/System/top:System/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getSystem({});
 * ```
 */
export function getSystem(args?: GetSystemArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getSystem:getSystem", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getSystem.
 */
export interface GetSystemArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
}

/**
 * A collection of values returned by getSystem.
 */
export interface GetSystemResult {
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * The system name (hostname).
     */
    readonly name: string;
}
/**
 * This data source can read the system configuration.
 *
 * - API Documentation: [topSystem](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/System/top:System/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getSystem({});
 * ```
 */
export function getSystemOutput(args?: GetSystemOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSystemResult> {
    return pulumi.output(args).apply((a: any) => getSystem(a, opts))
}

/**
 * A collection of arguments for invoking getSystem.
 */
export interface GetSystemOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}
