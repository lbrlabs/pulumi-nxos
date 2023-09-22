// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the vPC domain configuration.
 *
 * - API Documentation: [vpcDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/System/vpc:Dom/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getVpcDomain({});
 * ```
 */
export function getVpcDomain(args?: GetVpcDomainArgs, opts?: pulumi.InvokeOptions): Promise<GetVpcDomainResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getVpcDomain:getVpcDomain", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getVpcDomain.
 */
export interface GetVpcDomainArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
}

/**
 * A collection of values returned by getVpcDomain.
 */
export interface GetVpcDomainResult {
    /**
     * vPC suspend locally.
     */
    readonly adminState: string;
    /**
     * Auto Recovery.
     */
    readonly autoRecovery: string;
    /**
     * Auto Recovery interval.
     */
    readonly autoRecoveryInterval: number;
    /**
     * Delay restore for orphan ports.
     */
    readonly delayRestoreOrphanPort: number;
    /**
     * Delay restore for SVI.
     */
    readonly delayRestoreSvi: number;
    /**
     * Delay restore for vPC links.
     */
    readonly delayRestoreVpc: number;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * Domain id.
     */
    readonly domainId: number;
    /**
     * DSCP.
     */
    readonly dscp: number;
    /**
     * Fast Convergence.
     */
    readonly fastConvergence: string;
    /**
     * Graceful Type-1 Consistency Check.
     */
    readonly gracefulConsistencyCheck: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * L3 Peer Router.
     */
    readonly l3PeerRouter: string;
    /**
     * L3 Peer Router Syslog.
     */
    readonly l3PeerRouterSyslog: string;
    /**
     * L3 Peer Router Syslog Interval.
     */
    readonly l3PeerRouterSyslogInterval: number;
    /**
     * Peer Gateway.
     */
    readonly peerGateway: string;
    /**
     * vPC peer IP address.
     */
    readonly peerIp: string;
    /**
     * vPC pair switches.
     */
    readonly peerSwitch: string;
    /**
     * Role priority.
     */
    readonly rolePriority: number;
    /**
     * System MAC.
     */
    readonly sysMac: string;
    /**
     * System priority.
     */
    readonly systemPriority: number;
    /**
     * Tracking object to suspend vPC if object goes down.
     */
    readonly track: number;
    /**
     * vPC virtual IP address (vIP).
     */
    readonly virtualIp: string;
}
/**
 * This data source can read the vPC domain configuration.
 *
 * - API Documentation: [vpcDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/System/vpc:Dom/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getVpcDomain({});
 * ```
 */
export function getVpcDomainOutput(args?: GetVpcDomainOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVpcDomainResult> {
    return pulumi.output(args).apply((a: any) => getVpcDomain(a, opts))
}

/**
 * A collection of arguments for invoking getVpcDomain.
 */
export interface GetVpcDomainOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}
