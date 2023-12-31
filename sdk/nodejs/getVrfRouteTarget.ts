// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read a VRF Route Target Entry.
 *
 * - API Documentation: [rtctrlRttEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:RttEntry/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getVrfRouteTarget({
 *     addressFamily: "ipv4-ucast",
 *     direction: "import",
 *     routeTarget: "route-target:as2-nn2:2:2",
 *     routeTargetAddressFamily: "ipv4-ucast",
 *     vrf: "VRF1",
 * });
 * ```
 */
export function getVrfRouteTarget(args: GetVrfRouteTargetArgs, opts?: pulumi.InvokeOptions): Promise<GetVrfRouteTargetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getVrfRouteTarget:getVrfRouteTarget", {
        "addressFamily": args.addressFamily,
        "device": args.device,
        "direction": args.direction,
        "routeTarget": args.routeTarget,
        "routeTargetAddressFamily": args.routeTargetAddressFamily,
        "vrf": args.vrf,
    }, opts);
}

/**
 * A collection of arguments for invoking getVrfRouteTarget.
 */
export interface GetVrfRouteTargetArgs {
    /**
     * Address family.
     */
    addressFamily: string;
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * Route Target direction.
     */
    direction: string;
    /**
     * Route Target in NX-OS DME format.
     */
    routeTarget: string;
    /**
     * Route Target Address Family.
     */
    routeTargetAddressFamily: string;
    /**
     * VRF name.
     */
    vrf: string;
}

/**
 * A collection of values returned by getVrfRouteTarget.
 */
export interface GetVrfRouteTargetResult {
    /**
     * Address family.
     */
    readonly addressFamily: string;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * Route Target direction.
     */
    readonly direction: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * Route Target in NX-OS DME format.
     */
    readonly routeTarget: string;
    /**
     * Route Target Address Family.
     */
    readonly routeTargetAddressFamily: string;
    /**
     * VRF name.
     */
    readonly vrf: string;
}
/**
 * This data source can read a VRF Route Target Entry.
 *
 * - API Documentation: [rtctrlRttEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:RttEntry/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getVrfRouteTarget({
 *     addressFamily: "ipv4-ucast",
 *     direction: "import",
 *     routeTarget: "route-target:as2-nn2:2:2",
 *     routeTargetAddressFamily: "ipv4-ucast",
 *     vrf: "VRF1",
 * });
 * ```
 */
export function getVrfRouteTargetOutput(args: GetVrfRouteTargetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVrfRouteTargetResult> {
    return pulumi.output(args).apply((a: any) => getVrfRouteTarget(a, opts))
}

/**
 * A collection of arguments for invoking getVrfRouteTarget.
 */
export interface GetVrfRouteTargetOutputArgs {
    /**
     * Address family.
     */
    addressFamily: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Route Target direction.
     */
    direction: pulumi.Input<string>;
    /**
     * Route Target in NX-OS DME format.
     */
    routeTarget: pulumi.Input<string>;
    /**
     * Route Target Address Family.
     */
    routeTargetAddressFamily: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrf: pulumi.Input<string>;
}
