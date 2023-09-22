// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This resource can manage a Route-Map Rule Entry configuration.
 *
 * - API Documentation: [rtmapEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:Entry/)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as nxos from "@lbrlabs/pulumi-nxos";
 *
 * const example = new nxos.RouteMapRuleEntry("example", {
 *     action: "permit",
 *     order: 10,
 *     ruleName: "RULE1",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 *  $ pulumi import nxos:index/routeMapRuleEntry:RouteMapRuleEntry example "sys/rpm/rtmap-[RULE1]/ent-[10]"
 * ```
 */
export class RouteMapRuleEntry extends pulumi.CustomResource {
    /**
     * Get an existing RouteMapRuleEntry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RouteMapRuleEntryState, opts?: pulumi.CustomResourceOptions): RouteMapRuleEntry {
        return new RouteMapRuleEntry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'nxos:index/routeMapRuleEntry:RouteMapRuleEntry';

    /**
     * Returns true if the given object is an instance of RouteMapRuleEntry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RouteMapRuleEntry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RouteMapRuleEntry.__pulumiType;
    }

    /**
     * Route-Map Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * A device name from the provider configuration.
     */
    public readonly device!: pulumi.Output<string | undefined>;
    /**
     * Route-Map Rule Entry order. - Range: `0`-`65535`
     */
    public readonly order!: pulumi.Output<number>;
    /**
     * Route Map rule name.
     */
    public readonly ruleName!: pulumi.Output<string>;

    /**
     * Create a RouteMapRuleEntry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteMapRuleEntryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RouteMapRuleEntryArgs | RouteMapRuleEntryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RouteMapRuleEntryState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["device"] = state ? state.device : undefined;
            resourceInputs["order"] = state ? state.order : undefined;
            resourceInputs["ruleName"] = state ? state.ruleName : undefined;
        } else {
            const args = argsOrState as RouteMapRuleEntryArgs | undefined;
            if ((!args || args.order === undefined) && !opts.urn) {
                throw new Error("Missing required property 'order'");
            }
            if ((!args || args.ruleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ruleName'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["device"] = args ? args.device : undefined;
            resourceInputs["order"] = args ? args.order : undefined;
            resourceInputs["ruleName"] = args ? args.ruleName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RouteMapRuleEntry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RouteMapRuleEntry resources.
 */
export interface RouteMapRuleEntryState {
    /**
     * Route-Map Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
     */
    action?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Route-Map Rule Entry order. - Range: `0`-`65535`
     */
    order?: pulumi.Input<number>;
    /**
     * Route Map rule name.
     */
    ruleName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RouteMapRuleEntry resource.
 */
export interface RouteMapRuleEntryArgs {
    /**
     * Route-Map Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
     */
    action?: pulumi.Input<string>;
    /**
     * A device name from the provider configuration.
     */
    device?: pulumi.Input<string>;
    /**
     * Route-Map Rule Entry order. - Range: `0`-`65535`
     */
    order: pulumi.Input<number>;
    /**
     * Route Map rule name.
     */
    ruleName: pulumi.Input<string>;
}