// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read a VRF Route Distinguisher and VRF VNI.
 *
 * - API Documentation: [rtctrlDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:Dom/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getVrfRouting({
 *     vrf: "VRF1",
 * });
 * ```
 */
export function getVrfRouting(args: GetVrfRoutingArgs, opts?: pulumi.InvokeOptions): Promise<GetVrfRoutingResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getVrfRouting:getVrfRouting", {
        "device": args.device,
        "vrf": args.vrf,
    }, opts);
}

/**
 * A collection of arguments for invoking getVrfRouting.
 */
export interface GetVrfRoutingArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * VRF name.
     */
    vrf: string;
}

/**
 * A collection of values returned by getVrfRouting.
 */
export interface GetVrfRoutingResult {
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * Route Distinguisher value in NX-OS DME format.
     */
    readonly routeDistinguisher: string;
    /**
     * VRF name.
     */
    readonly vrf: string;
}
/**
 * This data source can read a VRF Route Distinguisher and VRF VNI.
 *
 * - API Documentation: [rtctrlDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:Dom/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getVrfRouting({
 *     vrf: "VRF1",
 * });
 * ```
 */
export function getVrfRoutingOutput(args: GetVrfRoutingOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVrfRoutingResult> {
    return pulumi.output(args).apply((a: any) => getVrfRouting(a, opts))
}

/**
 * A collection of arguments for invoking getVrfRouting.
 */
export interface GetVrfRoutingOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrf: pulumi.Input<string>;
}
