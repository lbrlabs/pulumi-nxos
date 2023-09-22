// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the default QoS class map DSCP configuration.
 *
 * - API Documentation: [ipqosDscp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:Dscp/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getDefaultQosClassMapDscp({
 *     classMapName: "Voice",
 *     value: "ef",
 * });
 * ```
 */
export function getDefaultQosClassMapDscp(args: GetDefaultQosClassMapDscpArgs, opts?: pulumi.InvokeOptions): Promise<GetDefaultQosClassMapDscpResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getDefaultQosClassMapDscp:getDefaultQosClassMapDscp", {
        "classMapName": args.classMapName,
        "device": args.device,
        "value": args.value,
    }, opts);
}

/**
 * A collection of arguments for invoking getDefaultQosClassMapDscp.
 */
export interface GetDefaultQosClassMapDscpArgs {
    /**
     * Class map name.
     */
    classMapName: string;
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * DSCP value.
     */
    value: string;
}

/**
 * A collection of values returned by getDefaultQosClassMapDscp.
 */
export interface GetDefaultQosClassMapDscpResult {
    /**
     * Class map name.
     */
    readonly classMapName: string;
    /**
     * A device name from the provider configuration.
     */
    readonly device?: string;
    /**
     * The distinguished name of the object.
     */
    readonly id: string;
    /**
     * DSCP value.
     */
    readonly value: string;
}
/**
 * This data source can read the default QoS class map DSCP configuration.
 *
 * - API Documentation: [ipqosDscp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:Dscp/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getDefaultQosClassMapDscp({
 *     classMapName: "Voice",
 *     value: "ef",
 * });
 * ```
 */
export function getDefaultQosClassMapDscpOutput(args: GetDefaultQosClassMapDscpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDefaultQosClassMapDscpResult> {
    return pulumi.output(args).apply((a: any) => getDefaultQosClassMapDscp(a, opts))
}

/**
 * A collection of arguments for invoking getDefaultQosClassMapDscp.
 */
export interface GetDefaultQosClassMapDscpOutputArgs {
    /**
     * Class map name.
     */
    classMapName: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * DSCP value.
     */
    value: pulumi.Input<string>;
}
