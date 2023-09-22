// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read a bridge domain.
 *
 * - API Documentation: [l2BD](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%202/l2:BD/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getBridgeDomain({
 *     fabricEncap: "vlan-10",
 * });
 * ```
 */
export function getBridgeDomain(args: GetBridgeDomainArgs, opts?: pulumi.InvokeOptions): Promise<GetBridgeDomainResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getBridgeDomain:getBridgeDomain", {
        "device": args.device,
        "fabricEncap": args.fabricEncap,
    }, opts);
}

/**
 * A collection of arguments for invoking getBridgeDomain.
 */
export interface GetBridgeDomainArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * Fabric encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     */
    fabricEncap: string;
}

/**
 * A collection of values returned by getBridgeDomain.
 */
export interface GetBridgeDomainResult {
    /**
     * Access encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     */
    readonly accessEncap: string;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * Fabric encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     */
    readonly fabricEncap: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * Bridge domain name.
     */
    readonly name: string;
}
/**
 * This data source can read a bridge domain.
 *
 * - API Documentation: [l2BD](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%202/l2:BD/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getBridgeDomain({
 *     fabricEncap: "vlan-10",
 * });
 * ```
 */
export function getBridgeDomainOutput(args: GetBridgeDomainOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBridgeDomainResult> {
    return pulumi.output(args).apply((a: any) => getBridgeDomain(a, opts))
}

/**
 * A collection of arguments for invoking getBridgeDomain.
 */
export interface GetBridgeDomainOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Fabric encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     */
    fabricEncap: pulumi.Input<string>;
}