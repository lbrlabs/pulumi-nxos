// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class SpanningTreeInterface extends pulumi.CustomResource {
    /**
     * Get an existing SpanningTreeInterface resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SpanningTreeInterfaceState, opts?: pulumi.CustomResourceOptions): SpanningTreeInterface {
        return new SpanningTreeInterface(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:nxos/spanningTreeInterface:SpanningTreeInterface';

    /**
     * Returns true if the given object is an instance of SpanningTreeInterface.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SpanningTreeInterface {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SpanningTreeInterface.__pulumiType;
    }

    /**
     * The administrative state of the object or policy. - Choices: `enabled`, `disabled` - Default value: `enabled`
     */
    public readonly adminState!: pulumi.Output<string>;
    /**
     * BPDU filter mode. - Choices: `default`, `enable`, `disable` - Default value: `default`
     */
    public readonly bpduFilter!: pulumi.Output<string>;
    /**
     * BPDU guard mode. - Choices: `default`, `enable`, `disable` - Default value: `default`
     */
    public readonly bpduGuard!: pulumi.Output<string>;
    /**
     * Port path cost. - Range: `0`-`200000000` - Default value: `0`
     */
    public readonly cost!: pulumi.Output<number>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Guard mode. - Choices: `default`, `root`, `loop`, `none` - Default value: `default`
     */
    public readonly guard!: pulumi.Output<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    public readonly interfaceId!: pulumi.Output<string>;
    /**
     * Link type. - Choices: `auto`, `p2p`, `shared` - Default value: `auto`
     */
    public readonly linkType!: pulumi.Output<string>;
    /**
     * Port mode. - Choices: `default`, `edge`, `network`, `normal`, `trunk` - Default value: `default`
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * Port priority. - Range: `0`-`224` - Default value: `128`
     */
    public readonly priority!: pulumi.Output<number>;

    /**
     * Create a SpanningTreeInterface resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SpanningTreeInterfaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SpanningTreeInterfaceArgs | SpanningTreeInterfaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SpanningTreeInterfaceState | undefined;
            resourceInputs["adminState"] = state ? state.adminState : undefined;
            resourceInputs["bpduFilter"] = state ? state.bpduFilter : undefined;
            resourceInputs["bpduGuard"] = state ? state.bpduGuard : undefined;
            resourceInputs["cost"] = state ? state.cost : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["guard"] = state ? state.guard : undefined;
            resourceInputs["interfaceId"] = state ? state.interfaceId : undefined;
            resourceInputs["linkType"] = state ? state.linkType : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
        } else {
            const args = argsOrState as SpanningTreeInterfaceArgs | undefined;
            if ((!args || args.interfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'interfaceId'");
            }
            resourceInputs["adminState"] = args ? args.adminState : undefined;
            resourceInputs["bpduFilter"] = args ? args.bpduFilter : undefined;
            resourceInputs["bpduGuard"] = args ? args.bpduGuard : undefined;
            resourceInputs["cost"] = args ? args.cost : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["guard"] = args ? args.guard : undefined;
            resourceInputs["interfaceId"] = args ? args.interfaceId : undefined;
            resourceInputs["linkType"] = args ? args.linkType : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SpanningTreeInterface.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SpanningTreeInterface resources.
 */
export interface SpanningTreeInterfaceState {
    /**
     * The administrative state of the object or policy. - Choices: `enabled`, `disabled` - Default value: `enabled`
     */
    adminState?: pulumi.Input<string>;
    /**
     * BPDU filter mode. - Choices: `default`, `enable`, `disable` - Default value: `default`
     */
    bpduFilter?: pulumi.Input<string>;
    /**
     * BPDU guard mode. - Choices: `default`, `enable`, `disable` - Default value: `default`
     */
    bpduGuard?: pulumi.Input<string>;
    /**
     * Port path cost. - Range: `0`-`200000000` - Default value: `0`
     */
    cost?: pulumi.Input<number>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Guard mode. - Choices: `default`, `root`, `loop`, `none` - Default value: `default`
     */
    guard?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    interfaceId?: pulumi.Input<string>;
    /**
     * Link type. - Choices: `auto`, `p2p`, `shared` - Default value: `auto`
     */
    linkType?: pulumi.Input<string>;
    /**
     * Port mode. - Choices: `default`, `edge`, `network`, `normal`, `trunk` - Default value: `default`
     */
    mode?: pulumi.Input<string>;
    /**
     * Port priority. - Range: `0`-`224` - Default value: `128`
     */
    priority?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SpanningTreeInterface resource.
 */
export interface SpanningTreeInterfaceArgs {
    /**
     * The administrative state of the object or policy. - Choices: `enabled`, `disabled` - Default value: `enabled`
     */
    adminState?: pulumi.Input<string>;
    /**
     * BPDU filter mode. - Choices: `default`, `enable`, `disable` - Default value: `default`
     */
    bpduFilter?: pulumi.Input<string>;
    /**
     * BPDU guard mode. - Choices: `default`, `enable`, `disable` - Default value: `default`
     */
    bpduGuard?: pulumi.Input<string>;
    /**
     * Port path cost. - Range: `0`-`200000000` - Default value: `0`
     */
    cost?: pulumi.Input<number>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Guard mode. - Choices: `default`, `root`, `loop`, `none` - Default value: `default`
     */
    guard?: pulumi.Input<string>;
    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     */
    interfaceId: pulumi.Input<string>;
    /**
     * Link type. - Choices: `auto`, `p2p`, `shared` - Default value: `auto`
     */
    linkType?: pulumi.Input<string>;
    /**
     * Port mode. - Choices: `default`, `edge`, `network`, `normal`, `trunk` - Default value: `default`
     */
    mode?: pulumi.Input<string>;
    /**
     * Port priority. - Range: `0`-`224` - Default value: `128`
     */
    priority?: pulumi.Input<number>;
}
