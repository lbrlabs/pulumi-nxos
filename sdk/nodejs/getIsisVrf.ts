// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the IS-IS VRF configuration.
 *
 * - API Documentation: [isisDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:Dom/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getIsisVrf({
 *     instanceName: "ISIS1",
 *     name: "default",
 * });
 * ```
 */
export function getIsisVrf(args: GetIsisVrfArgs, opts?: pulumi.InvokeOptions): Promise<GetIsisVrfResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getIsisVrf:getIsisVrf", {
        "device": args.device,
        "instanceName": args.instanceName,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIsisVrf.
 */
export interface GetIsisVrfArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * IS-IS instance name.
     */
    instanceName: string;
    /**
     * VRF name.
     */
    name: string;
}

/**
 * A collection of values returned by getIsisVrf.
 */
export interface GetIsisVrfResult {
    /**
     * Administrative state.
     */
    readonly adminState: string;
    /**
     * Authentication Check for ISIS on Level-1.
     */
    readonly authenticationCheckL1: boolean;
    /**
     * Authentication Check for ISIS on Level-2.
     */
    readonly authenticationCheckL2: boolean;
    /**
     * Authentication Key for IS-IS on Level-1.
     */
    readonly authenticationKeyL1: string;
    /**
     * Authentication Key for IS-IS on Level-2.
     */
    readonly authenticationKeyL2: string;
    /**
     * IS-IS Authentication-Type for Level-1.
     */
    readonly authenticationTypeL1: string;
    /**
     * IS-IS Authentication-Type for Level-2.
     */
    readonly authenticationTypeL2: string;
    /**
     * The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost metric.
     */
    readonly bandwidthReference: number;
    /**
     * Bandwidth reference unit.
     */
    readonly banwidthReferenceUnit: string;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * IS-IS instance name.
     */
    readonly instanceName: string;
    /**
     * IS-IS domain type.
     */
    readonly isType: string;
    /**
     * IS-IS metric type.
     */
    readonly metricType: string;
    /**
     * The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352 bytes.
     */
    readonly mtu: number;
    /**
     * VRF name.
     */
    readonly name: string;
    /**
     * Holds IS-IS domain NET (address) value.
     */
    readonly net: string;
    /**
     * IS-IS Domain passive-interface default level.
     */
    readonly passiveDefault: string;
}
/**
 * This data source can read the IS-IS VRF configuration.
 *
 * - API Documentation: [isisDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:Dom/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getIsisVrf({
 *     instanceName: "ISIS1",
 *     name: "default",
 * });
 * ```
 */
export function getIsisVrfOutput(args: GetIsisVrfOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIsisVrfResult> {
    return pulumi.output(args).apply((a: any) => getIsisVrf(a, opts))
}

/**
 * A collection of arguments for invoking getIsisVrf.
 */
export interface GetIsisVrfOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * IS-IS instance name.
     */
    instanceName: pulumi.Input<string>;
    /**
     * VRF name.
     */
    name: pulumi.Input<string>;
}
