// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read an IPv4 Access List Policy Egress Interface.
 *
 * - API Documentation: [aclIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Security%20and%20Policing/acl:If/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getIpv4AccessListPolicyEgressInterface({
 *     interfaceId: "eth1/10",
 * });
 * ```
 */
export function getIpv4AccessListPolicyEgressInterface(args: GetIpv4AccessListPolicyEgressInterfaceArgs, opts?: pulumi.InvokeOptions): Promise<GetIpv4AccessListPolicyEgressInterfaceResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getIpv4AccessListPolicyEgressInterface:getIpv4AccessListPolicyEgressInterface", {
        "device": args.device,
        "interfaceId": args.interfaceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpv4AccessListPolicyEgressInterface.
 */
export interface GetIpv4AccessListPolicyEgressInterfaceArgs {
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
 * A collection of values returned by getIpv4AccessListPolicyEgressInterface.
 */
export interface GetIpv4AccessListPolicyEgressInterfaceResult {
    /**
     * Access list name.
     */
    readonly accessListName: string;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    readonly interfaceId: string;
}
/**
 * This data source can read an IPv4 Access List Policy Egress Interface.
 *
 * - API Documentation: [aclIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Security%20and%20Policing/acl:If/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getIpv4AccessListPolicyEgressInterface({
 *     interfaceId: "eth1/10",
 * });
 * ```
 */
export function getIpv4AccessListPolicyEgressInterfaceOutput(args: GetIpv4AccessListPolicyEgressInterfaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIpv4AccessListPolicyEgressInterfaceResult> {
    return pulumi.output(args).apply((a: any) => getIpv4AccessListPolicyEgressInterface(a, opts))
}

/**
 * A collection of arguments for invoking getIpv4AccessListPolicyEgressInterface.
 */
export interface GetIpv4AccessListPolicyEgressInterfaceOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    interfaceId: pulumi.Input<string>;
}
