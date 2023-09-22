// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the PIM Static RP configuration.
 *
 * - API Documentation: [pimStaticRP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:StaticRP/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getPimStaticRp({
 *     address: "1.2.3.4",
 *     vrfName: "default",
 * });
 * ```
 */
export function getPimStaticRp(args: GetPimStaticRpArgs, opts?: pulumi.InvokeOptions): Promise<GetPimStaticRpResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getPimStaticRp:getPimStaticRp", {
        "address": args.address,
        "device": args.device,
        "vrfName": args.vrfName,
    }, opts);
}

/**
 * A collection of arguments for invoking getPimStaticRp.
 */
export interface GetPimStaticRpArgs {
    /**
     * Address.
     */
    address: string;
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * VRF name.
     */
    vrfName: string;
}

/**
 * A collection of values returned by getPimStaticRp.
 */
export interface GetPimStaticRpResult {
    /**
     * Address.
     */
    readonly address: string;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * VRF name.
     */
    readonly vrfName: string;
}
/**
 * This data source can read the PIM Static RP configuration.
 *
 * - API Documentation: [pimStaticRP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:StaticRP/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getPimStaticRp({
 *     address: "1.2.3.4",
 *     vrfName: "default",
 * });
 * ```
 */
export function getPimStaticRpOutput(args: GetPimStaticRpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPimStaticRpResult> {
    return pulumi.output(args).apply((a: any) => getPimStaticRp(a, opts))
}

/**
 * A collection of arguments for invoking getPimStaticRp.
 */
export interface GetPimStaticRpOutputArgs {
    /**
     * Address.
     */
    address: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrfName: pulumi.Input<string>;
}
