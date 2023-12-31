// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the configuration of Ingress Replication for Virtual Network ID (VNI).
 *
 * - API Documentation: [nvoIngRepl](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Network%20Virtualization/nvo:IngRepl/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getNveVniIngressReplication({
 *     vni: 103100,
 * });
 * ```
 */
export function getNveVniIngressReplication(args: GetNveVniIngressReplicationArgs, opts?: pulumi.InvokeOptions): Promise<GetNveVniIngressReplicationResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getNveVniIngressReplication:getNveVniIngressReplication", {
        "device": args.device,
        "vni": args.vni,
    }, opts);
}

/**
 * A collection of arguments for invoking getNveVniIngressReplication.
 */
export interface GetNveVniIngressReplicationArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * Virtual Network ID.
     */
    vni: number;
}

/**
 * A collection of values returned by getNveVniIngressReplication.
 */
export interface GetNveVniIngressReplicationResult {
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * Configure VxLAN Ingress Replication mode.
     */
    readonly protocol: string;
    /**
     * Virtual Network ID.
     */
    readonly vni: number;
}
/**
 * This data source can read the configuration of Ingress Replication for Virtual Network ID (VNI).
 *
 * - API Documentation: [nvoIngRepl](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Network%20Virtualization/nvo:IngRepl/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getNveVniIngressReplication({
 *     vni: 103100,
 * });
 * ```
 */
export function getNveVniIngressReplicationOutput(args: GetNveVniIngressReplicationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNveVniIngressReplicationResult> {
    return pulumi.output(args).apply((a: any) => getNveVniIngressReplication(a, opts))
}

/**
 * A collection of arguments for invoking getNveVniIngressReplication.
 */
export interface GetNveVniIngressReplicationOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Virtual Network ID.
     */
    vni: pulumi.Input<number>;
}
