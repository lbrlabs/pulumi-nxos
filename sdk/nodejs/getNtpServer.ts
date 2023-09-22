// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read an ntp server or peer.
 *
 * - API Documentation: [datetimeNtpProvider](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/System/datetime:NtpProvider/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getNtpServer({
 *     name: "1.2.3.4",
 * });
 * ```
 */
export function getNtpServer(args: GetNtpServerArgs, opts?: pulumi.InvokeOptions): Promise<GetNtpServerResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getNtpServer:getNtpServer", {
        "device": args.device,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getNtpServer.
 */
export interface GetNtpServerArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * NTP server.
     */
    name: string;
}

/**
 * A collection of values returned by getNtpServer.
 */
export interface GetNtpServerResult {
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * NTP provider key ID. Possible range is from `1` to `65535`.
     */
    readonly keyId: number;
    /**
     * NTP maximum interval default in seconds. Possible range is from `4` to `16`.
     */
    readonly maxPoll: number;
    /**
     * NTP minimum interval default in seconds. Possible range is from `4` to `16`.
     */
    readonly minPoll: number;
    /**
     * NTP server.
     */
    readonly name: string;
    /**
     * NTP provider type. Possible values are `server` or `peer`.
     */
    readonly type: string;
    /**
     * Identifies the VRF for the NTP providers.
     */
    readonly vrf: string;
}
/**
 * This data source can read an ntp server or peer.
 *
 * - API Documentation: [datetimeNtpProvider](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/System/datetime:NtpProvider/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getNtpServer({
 *     name: "1.2.3.4",
 * });
 * ```
 */
export function getNtpServerOutput(args: GetNtpServerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNtpServerResult> {
    return pulumi.output(args).apply((a: any) => getNtpServer(a, opts))
}

/**
 * A collection of arguments for invoking getNtpServer.
 */
export interface GetNtpServerOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * NTP server.
     */
    name: pulumi.Input<string>;
}
