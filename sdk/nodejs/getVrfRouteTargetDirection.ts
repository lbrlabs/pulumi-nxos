// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read a VRF Route Target direction.
 *
 * - API Documentation: [rtctrlRttP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:RttP/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getVrfRouteTargetDirection({
 *     addressFamily: "ipv4-ucast",
 *     direction: "import",
 *     routeTargetAddressFamily: "ipv4-ucast",
 *     vrf: "VRF1",
 * });
 * ```
 */
export function getVrfRouteTargetDirection(args: GetVrfRouteTargetDirectionArgs, opts?: pulumi.InvokeOptions): Promise<GetVrfRouteTargetDirectionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getVrfRouteTargetDirection:getVrfRouteTargetDirection", {
        "addressFamily": args.addressFamily,
        "device": args.device,
        "direction": args.direction,
        "routeTargetAddressFamily": args.routeTargetAddressFamily,
        "vrf": args.vrf,
    }, opts);
}

/**
 * A collection of arguments for invoking getVrfRouteTargetDirection.
 */
export interface GetVrfRouteTargetDirectionArgs {
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
     * Route Target Address Family.
     */
    routeTargetAddressFamily: string;
    /**
     * VRF name.
     */
    vrf: string;
}

/**
 * A collection of values returned by getVrfRouteTargetDirection.
 */
export interface GetVrfRouteTargetDirectionResult {
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
     * Route Target Address Family.
     */
    readonly routeTargetAddressFamily: string;
    /**
     * VRF name.
     */
    readonly vrf: string;
}
/**
 * This data source can read a VRF Route Target direction.
 *
 * - API Documentation: [rtctrlRttP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:RttP/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getVrfRouteTargetDirection({
 *     addressFamily: "ipv4-ucast",
 *     direction: "import",
 *     routeTargetAddressFamily: "ipv4-ucast",
 *     vrf: "VRF1",
 * });
 * ```
 */
export function getVrfRouteTargetDirectionOutput(args: GetVrfRouteTargetDirectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVrfRouteTargetDirectionResult> {
    return pulumi.output(args).apply((a: any) => getVrfRouteTargetDirection(a, opts))
}

/**
 * A collection of arguments for invoking getVrfRouteTargetDirection.
 */
export interface GetVrfRouteTargetDirectionOutputArgs {
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
     * Route Target Address Family.
     */
    routeTargetAddressFamily: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrf: pulumi.Input<string>;
}