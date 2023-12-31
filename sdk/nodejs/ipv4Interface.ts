// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage an IPv4 interface.
 *
 * - API Documentation: [ipv4If](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/ipv4:If/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.Ipv4Interface("example", {
 *     dropGlean: "disabled",
 *     forward: "disabled",
 *     interfaceId: "eth1/10",
 *     unnumbered: "unspecified",
 *     urpf: "disabled",
 *     vrf: "default",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/ipv4Interface:Ipv4Interface example "sys/ipv4/inst/dom-[default]/if-[eth1/10]"
 * ```
 */
export class Ipv4Interface extends pulumi.CustomResource {
    /**
     * Get an existing Ipv4Interface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: Ipv4InterfaceState, opts?: pulumi.CustomResourceOptions): Ipv4Interface {
        return new Ipv4Interface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/ipv4Interface:Ipv4Interface';

    /**
     * Returns true if the given object is an instance of Ipv4Interface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Ipv4Interface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Ipv4Interface.__pulumiType;
    }

    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * ip drop-glean enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
     */
    public readonly dropGlean!: pulumi.Output<string>;
    /**
     * ip forward enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
     */
    public readonly forward!: pulumi.Output<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    public readonly interfaceId!: pulumi.Output<string>;
    /**
     * IP unnumbered. Reference to interface must match first field in the output of `show intf brief`. Example: `eth1/1`. -
     * Default value: `unspecified`
     */
    public readonly unnumbered!: pulumi.Output<string>;
    /**
     * URPF (unicast Reverse Path Forwarding). - Choices: `disabled`, `strict`, `loose`, `loose-allow-default`,
     * `strict-allow-vni-hosts` - Default value: `disabled`
     */
    public readonly urpf!: pulumi.Output<string>;
    /**
     * VRF name.
     */
    public readonly vrf!: pulumi.Output<string>;

    /**
     * Create a Ipv4Interface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: Ipv4InterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: Ipv4InterfaceArgs | Ipv4InterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as Ipv4InterfaceState | undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["dropGlean"] = state ? state.dropGlean : undefined;
            resourceInputs["forward"] = state ? state.forward : undefined;
            resourceInputs["interfaceId"] = state ? state.interfaceId : undefined;
            resourceInputs["unnumbered"] = state ? state.unnumbered : undefined;
            resourceInputs["urpf"] = state ? state.urpf : undefined;
            resourceInputs["vrf"] = state ? state.vrf : undefined;
        } else {
            const args = argsOrState as Ipv4InterfaceArgs | undefined;
            if ((!args || args.interfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interfaceId'");
            }
            if ((!args || args.vrf === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vrf'");
            }
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["dropGlean"] = args ? args.dropGlean : undefined;
            resourceInputs["forward"] = args ? args.forward : undefined;
            resourceInputs["interfaceId"] = args ? args.interfaceId : undefined;
            resourceInputs["unnumbered"] = args ? args.unnumbered : undefined;
            resourceInputs["urpf"] = args ? args.urpf : undefined;
            resourceInputs["vrf"] = args ? args.vrf : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Ipv4Interface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Ipv4Interface resources.
 */
export interface Ipv4InterfaceState {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * ip drop-glean enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
     */
    dropGlean?: pulumi.Input<string>;
    /**
     * ip forward enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
     */
    forward?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    interfaceId?: pulumi.Input<string>;
    /**
     * IP unnumbered. Reference to interface must match first field in the output of `show intf brief`. Example: `eth1/1`. -
     * Default value: `unspecified`
     */
    unnumbered?: pulumi.Input<string>;
    /**
     * URPF (unicast Reverse Path Forwarding). - Choices: `disabled`, `strict`, `loose`, `loose-allow-default`,
     * `strict-allow-vni-hosts` - Default value: `disabled`
     */
    urpf?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrf?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Ipv4Interface resource.
 */
export interface Ipv4InterfaceArgs {
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * ip drop-glean enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
     */
    dropGlean?: pulumi.Input<string>;
    /**
     * ip forward enabled/disabled. - Choices: `enabled`, `disabled` - Default value: `disabled`
     */
    forward?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    interfaceId: pulumi.Input<string>;
    /**
     * IP unnumbered. Reference to interface must match first field in the output of `show intf brief`. Example: `eth1/1`. -
     * Default value: `unspecified`
     */
    unnumbered?: pulumi.Input<string>;
    /**
     * URPF (unicast Reverse Path Forwarding). - Choices: `disabled`, `strict`, `loose`, `loose-allow-default`,
     * `strict-allow-vni-hosts` - Default value: `disabled`
     */
    urpf?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrf: pulumi.Input<string>;
}
