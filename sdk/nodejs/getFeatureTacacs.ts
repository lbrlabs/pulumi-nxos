// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the TACACS+ feature configuration.
 *
 * - API Documentation: [fmTacacsplus](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Tacacsplus/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureTacacs({});
 * ```
 */
export function getFeatureTacacs(args?: GetFeatureTacacsArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureTacacsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getFeatureTacacs:getFeatureTacacs", {
        "device": args.device,
    }, opts);
}

/**
 * A collection of arguments for invoking getFeatureTacacs.
 */
export interface GetFeatureTacacsArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: string;
}

/**
 * A collection of values returned by getFeatureTacacs.
 */
export interface GetFeatureTacacsResult {
    /**
     * Administrative state.
     */
    readonly adminState: string;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
}
/**
 * This data source can read the TACACS+ feature configuration.
 *
 * - API Documentation: [fmTacacsplus](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Tacacsplus/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getFeatureTacacs({});
 * ```
 */
export function getFeatureTacacsOutput(args?: GetFeatureTacacsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureTacacsResult> {
    return pulumi.output(args).apply((a: any) => getFeatureTacacs(a, opts))
}

/**
 * A collection of arguments for invoking getFeatureTacacs.
 */
export interface GetFeatureTacacsOutputArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
}
