// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class PortChannelInterfaceVrf extends pulumi.CustomResource {
    /**
     * Get an existing PortChannelInterfaceVrf resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PortChannelInterfaceVrfState, opts?: pulumi.CustomResourceOptions): PortChannelInterfaceVrf {
        return new PortChannelInterfaceVrf(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:nxos/portChannelInterfaceVrf:PortChannelInterfaceVrf';

    /**
     * Returns true if the given object is an instance of PortChannelInterfaceVrf.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PortChannelInterfaceVrf {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PortChannelInterfaceVrf.__pulumiType;
    }

    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `po1`.
     */
    public readonly interfaceId!: pulumi.Output<string>;
    /**
     * DN of VRF. For example: `sys/inst-VRF1`.
     */
    public readonly vrfDn!: pulumi.Output<string>;

    /**
     * Create a PortChannelInterfaceVrf resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PortChannelInterfaceVrfArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PortChannelInterfaceVrfArgs | PortChannelInterfaceVrfState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PortChannelInterfaceVrfState | undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["interfaceId"] = state ? state.interfaceId : undefined;
            resourceInputs["vrfDn"] = state ? state.vrfDn : undefined;
        } else {
            const args = argsOrState as PortChannelInterfaceVrfArgs | undefined;
            if ((!args || args.interfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interfaceId'");
            }
            if ((!args || args.vrfDn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vrfDn'");
            }
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["interfaceId"] = args ? args.interfaceId : undefined;
            resourceInputs["vrfDn"] = args ? args.vrfDn : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PortChannelInterfaceVrf.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PortChannelInterfaceVrf resources.
 */
export interface PortChannelInterfaceVrfState {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `po1`.
     */
    interfaceId?: pulumi.Input<string>;
    /**
     * DN of VRF. For example: `sys/inst-VRF1`.
     */
    vrfDn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a PortChannelInterfaceVrf resource.
 */
export interface PortChannelInterfaceVrfArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `po1`.
     */
    interfaceId: pulumi.Input<string>;
    /**
     * DN of VRF. For example: `sys/inst-VRF1`.
     */
    vrfDn: pulumi.Input<string>;
}
