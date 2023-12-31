// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read a subinterface.
 *
 * - API Documentation: [l3EncRtdIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/l3:EncRtdIf/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getSubinterface({
 *     interfaceId: "eth1/10.124",
 * });
 * ```
 */
export function getSubinterface(args: GetSubinterfaceArgs, opts?: pulumi.InvokeOptions): Promise<GetSubinterfaceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getSubinterface:getSubinterface", {
        "device": args.device,
        "interfaceId": args.interfaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSubinterface.
 */
export interface GetSubinterfaceArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
     */
    interfaceId: string;
}

/**
 * A collection of values returned by getSubinterface.
 */
export interface GetSubinterfaceResult {
    /**
     * Administrative state.
     */
    readonly adminState: string;
    /**
     * Specifies the administrative port bandwidth.
     */
    readonly bandwidth: number;
    /**
     * Specifies the administrative port delay.
     */
    readonly delay: number;
    /**
     * Interface description.
     */
    readonly description: string;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * Subinterface encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     */
    readonly encap: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
     */
    readonly interfaceId: string;
    /**
     * Administrative link logging.
     */
    readonly linkLogging: string;
    /**
     * The administrative port medium type.
     */
    readonly medium: string;
    /**
     * Administrative port MTU.
     */
    readonly mtu: number;
}
/**
 * This data source can read a subinterface.
 *
 * - API Documentation: [l3EncRtdIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/l3:EncRtdIf/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getSubinterface({
 *     interfaceId: "eth1/10.124",
 * });
 * ```
 */
export function getSubinterfaceOutput(args: GetSubinterfaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSubinterfaceResult> {
    return pulumi.output(args).apply((a: any) => getSubinterface(a, opts))
}

/**
 * A collection of arguments for invoking getSubinterface.
 */
export interface GetSubinterfaceOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
     */
    interfaceId: pulumi.Input<string>;
}
