// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can read the default QoS policy map match class map set QoS group configuration.
 *
 * - API Documentation: [ipqosSetQoSGrp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:SetQoSGrp/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getDefaultQosPolicyMapMatchClassMapSetQosGroup({
 *     classMapName: "Voice",
 *     policyMapName: "PM1",
 * });
 * ```
 */
export function getDefaultQosPolicyMapMatchClassMapSetQosGroup(args: GetDefaultQosPolicyMapMatchClassMapSetQosGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetDefaultQosPolicyMapMatchClassMapSetQosGroupResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("nxos:index/getDefaultQosPolicyMapMatchClassMapSetQosGroup:getDefaultQosPolicyMapMatchClassMapSetQosGroup", {
        "classMapName": args.classMapName,
        "device": args.device,
        "policyMapName": args.policyMapName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDefaultQosPolicyMapMatchClassMapSetQosGroup.
 */
export interface GetDefaultQosPolicyMapMatchClassMapSetQosGroupArgs {
    /**
     * Class map name.
     */
    classMapName: string;
    /**
     * A device name from the provider configuration.
     */
    device?: string;
    /**
     * Policy map name.
     */
    policyMapName: string;
}

/**
 * A collection of values returned by getDefaultQosPolicyMapMatchClassMapSetQosGroup.
 */
export interface GetDefaultQosPolicyMapMatchClassMapSetQosGroupResult {
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
     * Policy map name.
     */
    readonly policyMapName: string;
    /**
     * QoS group ID.
     */
    readonly qosGroupId: number;
}
/**
 * This data source can read the default QoS policy map match class map set QoS group configuration.
 *
 * - API Documentation: [ipqosSetQoSGrp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:SetQoSGrp/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@pulumi/nxos";
 *
 * const example = nxos.getDefaultQosPolicyMapMatchClassMapSetQosGroup({
 *     classMapName: "Voice",
 *     policyMapName: "PM1",
 * });
 * ```
 */
export function getDefaultQosPolicyMapMatchClassMapSetQosGroupOutput(args: GetDefaultQosPolicyMapMatchClassMapSetQosGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDefaultQosPolicyMapMatchClassMapSetQosGroupResult> {
    return pulumi.output(args).apply((a: any) => getDefaultQosPolicyMapMatchClassMapSetQosGroup(a, opts))
}

/**
 * A collection of arguments for invoking getDefaultQosPolicyMapMatchClassMapSetQosGroup.
 */
export interface GetDefaultQosPolicyMapMatchClassMapSetQosGroupOutputArgs {
    /**
     * Class map name.
     */
    classMapName: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Policy map name.
     */
    policyMapName: pulumi.Input<string>;
}