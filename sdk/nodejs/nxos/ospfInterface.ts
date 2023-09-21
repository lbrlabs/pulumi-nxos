// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class OspfInterface extends pulumi.CustomResource {
    /**
     * Get an existing OspfInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OspfInterfaceState, opts?: pulumi.CustomResourceOptions): OspfInterface {
        return new OspfInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:nxos/ospfInterface:OspfInterface';

    /**
     * Returns true if the given object is an instance of OspfInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OspfInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OspfInterface.__pulumiType;
    }

    /**
     * Advertise secondary IP addresses. - Default value: `true`
     */
    public readonly advertiseSecondaries!: pulumi.Output<boolean>;
    /**
     * Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
     */
    public readonly area!: pulumi.Output<string>;
    /**
     * Bidirectional Forwarding Detection (BFD). - Choices: `unspecified`, `enabled`, `disabled` - Default value: `unspecified`
     */
    public readonly bfd!: pulumi.Output<string>;
    /**
     * Specifies the cost of interface. - Range: `0`-`65535` - Default value: `0`
     */
    public readonly cost!: pulumi.Output<number>;
    /**
     * Dead interval, interval after which router declares that neighbor as down. - Range: `0`-`65535` - Default value: `0`
     */
    public readonly deadInterval!: pulumi.Output<number>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Hello interval, interval between hello packets that OSPF sends on the interface. - Range: `0`-`65535` - Default value:
     * `10`
     */
    public readonly helloInterval!: pulumi.Output<number>;
    /**
     * OSPF instance name.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    public readonly interfaceId!: pulumi.Output<string>;
    /**
     * Network type. - Choices: `unspecified`, `p2p`, `bcast` - Default value: `unspecified`
     */
    public readonly networkType!: pulumi.Output<string>;
    /**
     * Passive interface control. Interface can be configured as passive or non-passive. - Choices: `unspecified`, `enabled`,
     * `disabled` - Default value: `unspecified`
     */
    public readonly passive!: pulumi.Output<string>;
    /**
     * Priority, used in determining the designated router on this network. - Range: `0`-`255` - Default value: `1`
     */
    public readonly priority!: pulumi.Output<number>;
    /**
     * VRF name.
     */
    public readonly vrfName!: pulumi.Output<string>;

    /**
     * Create a OspfInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OspfInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OspfInterfaceArgs | OspfInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OspfInterfaceState | undefined;
            resourceInputs["advertiseSecondaries"] = state ? state.advertiseSecondaries : undefined;
            resourceInputs["area"] = state ? state.area : undefined;
            resourceInputs["bfd"] = state ? state.bfd : undefined;
            resourceInputs["cost"] = state ? state.cost : undefined;
            resourceInputs["deadInterval"] = state ? state.deadInterval : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["helloInterval"] = state ? state.helloInterval : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["interfaceId"] = state ? state.interfaceId : undefined;
            resourceInputs["networkType"] = state ? state.networkType : undefined;
            resourceInputs["passive"] = state ? state.passive : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["vrfName"] = state ? state.vrfName : undefined;
        } else {
            const args = argsOrState as OspfInterfaceArgs | undefined;
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            if ((!args || args.interfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interfaceId'");
            }
            if ((!args || args.vrfName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vrfName'");
            }
            resourceInputs["advertiseSecondaries"] = args ? args.advertiseSecondaries : undefined;
            resourceInputs["area"] = args ? args.area : undefined;
            resourceInputs["bfd"] = args ? args.bfd : undefined;
            resourceInputs["cost"] = args ? args.cost : undefined;
            resourceInputs["deadInterval"] = args ? args.deadInterval : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["helloInterval"] = args ? args.helloInterval : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["interfaceId"] = args ? args.interfaceId : undefined;
            resourceInputs["networkType"] = args ? args.networkType : undefined;
            resourceInputs["passive"] = args ? args.passive : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["vrfName"] = args ? args.vrfName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OspfInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OspfInterface resources.
 */
export interface OspfInterfaceState {
    /**
     * Advertise secondary IP addresses. - Default value: `true`
     */
    advertiseSecondaries?: pulumi.Input<boolean>;
    /**
     * Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
     */
    area?: pulumi.Input<string>;
    /**
     * Bidirectional Forwarding Detection (BFD). - Choices: `unspecified`, `enabled`, `disabled` - Default value: `unspecified`
     */
    bfd?: pulumi.Input<string>;
    /**
     * Specifies the cost of interface. - Range: `0`-`65535` - Default value: `0`
     */
    cost?: pulumi.Input<number>;
    /**
     * Dead interval, interval after which router declares that neighbor as down. - Range: `0`-`65535` - Default value: `0`
     */
    deadInterval?: pulumi.Input<number>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Hello interval, interval between hello packets that OSPF sends on the interface. - Range: `0`-`65535` - Default value:
     * `10`
     */
    helloInterval?: pulumi.Input<number>;
    /**
     * OSPF instance name.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    interfaceId?: pulumi.Input<string>;
    /**
     * Network type. - Choices: `unspecified`, `p2p`, `bcast` - Default value: `unspecified`
     */
    networkType?: pulumi.Input<string>;
    /**
     * Passive interface control. Interface can be configured as passive or non-passive. - Choices: `unspecified`, `enabled`,
     * `disabled` - Default value: `unspecified`
     */
    passive?: pulumi.Input<string>;
    /**
     * Priority, used in determining the designated router on this network. - Range: `0`-`255` - Default value: `1`
     */
    priority?: pulumi.Input<number>;
    /**
     * VRF name.
     */
    vrfName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OspfInterface resource.
 */
export interface OspfInterfaceArgs {
    /**
     * Advertise secondary IP addresses. - Default value: `true`
     */
    advertiseSecondaries?: pulumi.Input<boolean>;
    /**
     * Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
     */
    area?: pulumi.Input<string>;
    /**
     * Bidirectional Forwarding Detection (BFD). - Choices: `unspecified`, `enabled`, `disabled` - Default value: `unspecified`
     */
    bfd?: pulumi.Input<string>;
    /**
     * Specifies the cost of interface. - Range: `0`-`65535` - Default value: `0`
     */
    cost?: pulumi.Input<number>;
    /**
     * Dead interval, interval after which router declares that neighbor as down. - Range: `0`-`65535` - Default value: `0`
     */
    deadInterval?: pulumi.Input<number>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Hello interval, interval between hello packets that OSPF sends on the interface. - Range: `0`-`65535` - Default value:
     * `10`
     */
    helloInterval?: pulumi.Input<number>;
    /**
     * OSPF instance name.
     */
    instanceName: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    interfaceId: pulumi.Input<string>;
    /**
     * Network type. - Choices: `unspecified`, `p2p`, `bcast` - Default value: `unspecified`
     */
    networkType?: pulumi.Input<string>;
    /**
     * Passive interface control. Interface can be configured as passive or non-passive. - Choices: `unspecified`, `enabled`,
     * `disabled` - Default value: `unspecified`
     */
    passive?: pulumi.Input<string>;
    /**
     * Priority, used in determining the designated router on this network. - Range: `0`-`255` - Default value: `1`
     */
    priority?: pulumi.Input<number>;
    /**
     * VRF name.
     */
    vrfName: pulumi.Input<string>;
}
