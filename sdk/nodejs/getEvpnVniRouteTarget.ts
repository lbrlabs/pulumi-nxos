// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read a EVPN VNI Route Target Entry.
 *
 * - API Documentation: [rtctrlRttEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:RttEntry/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getEvpnVniRouteTarget({
 *     direction: "import",
 *     encap: "vxlan-123456",
 *     routeTarget: "route-target:as2-nn2:2:2",
 * });
 * ```
 */
export function getEvpnVniRouteTarget(args: GetEvpnVniRouteTargetArgs, opts?: pulumi.InvokeOptions): Promise<GetEvpnVniRouteTargetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getEvpnVniRouteTarget:getEvpnVniRouteTarget", {
        "device": args.device,
        "direction": args.direction,
        "encap": args.encap,
        "routeTarget": args.routeTarget,
    }, opts);
}

/**
 * A collection of arguments for invoking getEvpnVniRouteTarget.
 */
export interface GetEvpnVniRouteTargetArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * Route Target direction.
     */
    direction: string;
    /**
     * Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     */
    encap: string;
    /**
     * Route Target in NX-OS DME format.
     */
    routeTarget: string;
}

/**
 * A collection of values returned by getEvpnVniRouteTarget.
 */
export interface GetEvpnVniRouteTargetResult {
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * Route Target direction.
     */
    readonly direction: string;
    /**
     * Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     */
    readonly encap: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * Route Target in NX-OS DME format.
     */
    readonly routeTarget: string;
}
/**
 * This data source can read a EVPN VNI Route Target Entry.
 *
 * - API Documentation: [rtctrlRttEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:RttEntry/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getEvpnVniRouteTarget({
 *     direction: "import",
 *     encap: "vxlan-123456",
 *     routeTarget: "route-target:as2-nn2:2:2",
 * });
 * ```
 */
export function getEvpnVniRouteTargetOutput(args: GetEvpnVniRouteTargetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEvpnVniRouteTargetResult> {
    return pulumi.output(args).apply((a: any) => getEvpnVniRouteTarget(a, opts))
}

/**
 * A collection of arguments for invoking getEvpnVniRouteTarget.
 */
export interface GetEvpnVniRouteTargetOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Route Target direction.
     */
    direction: pulumi.Input<string>;
    /**
     * Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     */
    encap: pulumi.Input<string>;
    /**
     * Route Target in NX-OS DME format.
     */
    routeTarget: pulumi.Input<string>;
}