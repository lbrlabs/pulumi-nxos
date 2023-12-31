// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage a DHCP relay address.
 *
 * - API Documentation: [dhcpRelayAddr](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/DHCP/dhcp:RelayAddr/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.DhcpRelayAddress("example", {
 *     address: "1.1.1.1",
 *     interfaceId: "eth1/10",
 *     vrf: "VRF1",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/dhcpRelayAddress:DhcpRelayAddress example "sys/dhcp/inst/relayif-[eth1/10]/addr-[VRF1]-[1.1.1.1]"
 * ```
 */
export class DhcpRelayAddress extends pulumi.CustomResource {
    /**
     * Get an existing DhcpRelayAddress resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DhcpRelayAddressState, opts?: pulumi.CustomResourceOptions): DhcpRelayAddress {
        return new DhcpRelayAddress(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/dhcpRelayAddress:DhcpRelayAddress';

    /**
     * Returns true if the given object is an instance of DhcpRelayAddress.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DhcpRelayAddress {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DhcpRelayAddress.__pulumiType;
    }

    /**
     * IPv4 or IPv6 address.
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    public readonly interfaceId!: pulumi.Output<string>;
    /**
     * VRF name.
     */
    public readonly vrf!: pulumi.Output<string>;

    /**
     * Create a DhcpRelayAddress resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DhcpRelayAddressArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DhcpRelayAddressArgs | DhcpRelayAddressState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DhcpRelayAddressState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["interfaceId"] = state ? state.interfaceId : undefined;
            resourceInputs["vrf"] = state ? state.vrf : undefined;
        } else {
            const args = argsOrState as DhcpRelayAddressArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.interfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interfaceId'");
            }
            if ((!args || args.vrf === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vrf'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["interfaceId"] = args ? args.interfaceId : undefined;
            resourceInputs["vrf"] = args ? args.vrf : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DhcpRelayAddress.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DhcpRelayAddress resources.
 */
export interface DhcpRelayAddressState {
    /**
     * IPv4 or IPv6 address.
     */
    address?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    interfaceId?: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrf?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DhcpRelayAddress resource.
 */
export interface DhcpRelayAddressArgs {
    /**
     * IPv4 or IPv6 address.
     */
    address: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    interfaceId: pulumi.Input<string>;
    /**
     * VRF name.
     */
    vrf: pulumi.Input<string>;
}
