// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the BGP domain (VRF) graceful restart configuration.
 *
 * - API Documentation: [bgpGr](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Gr/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getBgpGracefulRestart({
 *     asn: "65001",
 *     vrf: "default",
 * });
 * ```
 */
export function getBgpGracefulRestart(args: GetBgpGracefulRestartArgs, opts?: pulumi.InvokeOptions): Promise<GetBgpGracefulRestartResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getBgpGracefulRestart:getBgpGracefulRestart", {
        "asn": args.asn,
        "device": args.device,
        "vrf": args.vrf,
    }, opts);
}

/**
 * A collection of arguments for invoking getBgpGracefulRestart.
 */
export interface GetBgpGracefulRestartArgs {
    /**
     * Autonomous system number.
     */
    asn: string;
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
 * A collection of values returned by getBgpGracefulRestart.
 */
export interface GetBgpGracefulRestartResult {
    /**
     * Autonomous system number.
     */
    readonly asn: string;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * The graceful restart interval.
     */
    readonly restartInterval: number;
    /**
     * The stale interval for routes advertised by the BGP peer.
     */
    readonly staleInterval: number;
    /**
     * VRF name.
     */
    readonly vrf: string;
}
/**
 * This data source can read the BGP domain (VRF) graceful restart configuration.
 *
 * - API Documentation: [bgpGr](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Gr/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getBgpGracefulRestart({
 *     asn: "65001",
 *     vrf: "default",
 * });
 * ```
 */
export function getBgpGracefulRestartOutput(args: GetBgpGracefulRestartOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBgpGracefulRestartResult> {
    return pulumi.output(args).apply((a: any) => getBgpGracefulRestart(a, opts))
}

/**
 * A collection of arguments for invoking getBgpGracefulRestart.
 */
export interface GetBgpGracefulRestartOutputArgs {
    /**
     * Autonomous system number.
     */
    asn: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrf: pulumi.Input<string>;
}
