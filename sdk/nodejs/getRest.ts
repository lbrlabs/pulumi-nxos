// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read one NX-OS DME object.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const l1PhysIf = nxos.getRest({
 *     dn: "sys/intf/phys-[eth1/1]",
 * });
 * ```
 */
export function getRest(args: GetRestArgs, opts?: pulumi.InvokeOptions): Promise<GetRestResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getRest:getRest", {
        "device": args.device,
        "dn": args.dn,
    }, opts);
}

/**
 * A collection of arguments for invoking getRest.
 */
export interface GetRestArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * Distinguished name of object to be retrieved, e.g. sys/intf/phys-[eth1/1].
     */
    dn: string;
}

/**
 * A collection of values returned by getRest.
 */
export interface GetRestResult {
    /**
     * Class name of object being retrieved.
     */
    readonly className: string;
    /**
     * Map of key-value pairs which represents the attributes of object being retrieved.
     */
    readonly content: {[key: string]: string};
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * Distinguished name of object to be retrieved, e.g. sys/intf/phys-[eth1/1].
     */
    readonly dn: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
}
/**
 * This data source can read one NX-OS DME object.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const l1PhysIf = nxos.getRest({
 *     dn: "sys/intf/phys-[eth1/1]",
 * });
 * ```
 */
export function getRestOutput(args: GetRestOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRestResult> {
    return pulumi.output(args).apply((a: any) => getRest(a, opts))
}

/**
 * A collection of arguments for invoking getRest.
 */
export interface GetRestOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Distinguished name of object to be retrieved, e.g. sys/intf/phys-[eth1/1].
     */
    dn: pulumi.Input<string>;
}
