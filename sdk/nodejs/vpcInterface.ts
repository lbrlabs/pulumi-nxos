// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage the vPC interface configuration.
 *
 * - API Documentation: [vpcIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/System/vpc:If/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.VpcInterface("example", {
 *     portChannelInterfaceDn: "sys/intf/aggr-[po1]",
 *     vpcInterfaceId: 1,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/vpcInterface:VpcInterface example "sys/vpc/inst/dom/if-[1]"
 * ```
 */
export class VpcInterface extends pulumi.CustomResource {
    /**
     * Get an existing VpcInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VpcInterfaceState, opts?: pulumi.CustomResourceOptions): VpcInterface {
        return new VpcInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/vpcInterface:VpcInterface';

    /**
     * Returns true if the given object is an instance of VpcInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpcInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpcInterface.__pulumiType;
    }

    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Port-channel interface DN.
     */
    public readonly portChannelInterfaceDn!: pulumi.Output<string | undefined>;
    /**
     * The vPC interface identifier. - Range: `1`-`16384`
     */
    public readonly vpcInterfaceId!: pulumi.Output<number>;

    /**
     * Create a VpcInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpcInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpcInterfaceArgs | VpcInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VpcInterfaceState | undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["portChannelInterfaceDn"] = state ? state.portChannelInterfaceDn : undefined;
            resourceInputs["vpcInterfaceId"] = state ? state.vpcInterfaceId : undefined;
        } else {
            const args = argsOrState as VpcInterfaceArgs | undefined;
            if ((!args || args.vpcInterfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcInterfaceId'");
            }
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["portChannelInterfaceDn"] = args ? args.portChannelInterfaceDn : undefined;
            resourceInputs["vpcInterfaceId"] = args ? args.vpcInterfaceId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VpcInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VpcInterface resources.
 */
export interface VpcInterfaceState {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Port-channel interface DN.
     */
    portChannelInterfaceDn?: pulumi.Input<string>;
    /**
     * The vPC interface identifier. - Range: `1`-`16384`
     */
    vpcInterfaceId?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a VpcInterface resource.
 */
export interface VpcInterfaceArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Port-channel interface DN.
     */
    portChannelInterfaceDn?: pulumi.Input<string>;
    /**
     * The vPC interface identifier. - Range: `1`-`16384`
     */
    vpcInterfaceId: pulumi.Input<number>;
}
