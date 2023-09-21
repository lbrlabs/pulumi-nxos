// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class BgpGracefulRestart extends pulumi.CustomResource {
    /**
     * Get an existing BgpGracefulRestart resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BgpGracefulRestartState, opts?: pulumi.CustomResourceOptions): BgpGracefulRestart {
        return new BgpGracefulRestart(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:nxos/bgpGracefulRestart:BgpGracefulRestart';

    /**
     * Returns true if the given object is an instance of BgpGracefulRestart.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BgpGracefulRestart {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BgpGracefulRestart.__pulumiType;
    }

    /**
     * Autonomous system number.
     */
    public readonly asn!: pulumi.Output<string>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
     */
    public readonly restartInterval!: pulumi.Output<number>;
    /**
     * The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
     */
    public readonly staleInterval!: pulumi.Output<number>;
    /**
     * VRF name.
     */
    public readonly vrf!: pulumi.Output<string>;

    /**
     * Create a BgpGracefulRestart resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BgpGracefulRestartArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BgpGracefulRestartArgs | BgpGracefulRestartState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BgpGracefulRestartState | undefined;
            resourceInputs["asn"] = state ? state.asn : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["restartInterval"] = state ? state.restartInterval : undefined;
            resourceInputs["staleInterval"] = state ? state.staleInterval : undefined;
            resourceInputs["vrf"] = state ? state.vrf : undefined;
        } else {
            const args = argsOrState as BgpGracefulRestartArgs | undefined;
            if ((!args || args.asn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'asn'");
            }
            if ((!args || args.vrf === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vrf'");
            }
            resourceInputs["asn"] = args ? args.asn : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["restartInterval"] = args ? args.restartInterval : undefined;
            resourceInputs["staleInterval"] = args ? args.staleInterval : undefined;
            resourceInputs["vrf"] = args ? args.vrf : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BgpGracefulRestart.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BgpGracefulRestart resources.
 */
export interface BgpGracefulRestartState {
    /**
     * Autonomous system number.
     */
    asn?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
     */
    restartInterval?: pulumi.Input<number>;
    /**
     * The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
     */
    staleInterval?: pulumi.Input<number>;
    /**
     * VRF name.
     */
    vrf?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BgpGracefulRestart resource.
 */
export interface BgpGracefulRestartArgs {
    /**
     * Autonomous system number.
     */
    asn: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
     */
    restartInterval?: pulumi.Input<number>;
    /**
     * The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
     */
    staleInterval?: pulumi.Input<number>;
    /**
     * VRF name.
     */
    vrf: pulumi.Input<string>;
}