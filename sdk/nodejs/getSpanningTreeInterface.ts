// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the Spanning Tree interface configuration.
 *
 * - API Documentation: [stpIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Discovery%20Protocols/stp:If/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getSpanningTreeInterface({
 *     interfaceId: "eth1/9",
 * });
 * ```
 */
export function getSpanningTreeInterface(args: GetSpanningTreeInterfaceArgs, opts?: pulumi.InvokeOptions): Promise<GetSpanningTreeInterfaceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getSpanningTreeInterface:getSpanningTreeInterface", {
        "device": args.device,
        "interfaceId": args.interfaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSpanningTreeInterface.
 */
export interface GetSpanningTreeInterfaceArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    interfaceId: string;
}

/**
 * A collection of values returned by getSpanningTreeInterface.
 */
export interface GetSpanningTreeInterfaceResult {
    /**
     * The administrative state of the object or policy.
     */
    readonly adminState: string;
    /**
     * BPDU filter mode.
     */
    readonly bpduFilter: string;
    /**
     * BPDU guard mode.
     */
    readonly bpduGuard: string;
    /**
     * Port path cost.
     */
    readonly cost: number;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * Guard mode.
     */
    readonly guard: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    readonly interfaceId: string;
    /**
     * Link type.
     */
    readonly linkType: string;
    /**
     * Port mode.
     */
    readonly mode: string;
    /**
     * Port priority.
     */
    readonly priority: number;
}
/**
 * This data source can read the Spanning Tree interface configuration.
 *
 * - API Documentation: [stpIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Discovery%20Protocols/stp:If/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getSpanningTreeInterface({
 *     interfaceId: "eth1/9",
 * });
 * ```
 */
export function getSpanningTreeInterfaceOutput(args: GetSpanningTreeInterfaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSpanningTreeInterfaceResult> {
    return pulumi.output(args).apply((a: any) => getSpanningTreeInterface(a, opts))
}

/**
 * A collection of arguments for invoking getSpanningTreeInterface.
 */
export interface GetSpanningTreeInterfaceOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    interfaceId: pulumi.Input<string>;
}
