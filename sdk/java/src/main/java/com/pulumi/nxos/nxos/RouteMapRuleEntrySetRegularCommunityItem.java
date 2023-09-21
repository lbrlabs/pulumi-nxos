// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.nxos.RouteMapRuleEntrySetRegularCommunityItemArgs;
import com.pulumi.nxos.nxos.inputs.RouteMapRuleEntrySetRegularCommunityItemState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="nxos:nxos/routeMapRuleEntrySetRegularCommunityItem:RouteMapRuleEntrySetRegularCommunityItem")
public class RouteMapRuleEntrySetRegularCommunityItem extends com.pulumi.resources.CustomResource {
    /**
     * Set Community.
     * 
     */
    @Export(name="community", refs={String.class}, tree="[0]")
    private Output<String> community;

    /**
     * @return Set Community.
     * 
     */
    public Output<String> community() {
        return this.community;
    }
    /**
     * A device name from the provider configuration.
     * 
     */
    @Export(name="device", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> device;

    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Output<Optional<String>> device() {
        return Codegen.optional(this.device);
    }
    /**
     * Route-Map Rule Entry order. - Range: `0`-`65535`
     * 
     */
    @Export(name="order", refs={Integer.class}, tree="[0]")
    private Output<Integer> order;

    /**
     * @return Route-Map Rule Entry order. - Range: `0`-`65535`
     * 
     */
    public Output<Integer> order() {
        return this.order;
    }
    /**
     * Route Map rule name.
     * 
     */
    @Export(name="ruleName", refs={String.class}, tree="[0]")
    private Output<String> ruleName;

    /**
     * @return Route Map rule name.
     * 
     */
    public Output<String> ruleName() {
        return this.ruleName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RouteMapRuleEntrySetRegularCommunityItem(String name) {
        this(name, RouteMapRuleEntrySetRegularCommunityItemArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RouteMapRuleEntrySetRegularCommunityItem(String name, RouteMapRuleEntrySetRegularCommunityItemArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RouteMapRuleEntrySetRegularCommunityItem(String name, RouteMapRuleEntrySetRegularCommunityItemArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/routeMapRuleEntrySetRegularCommunityItem:RouteMapRuleEntrySetRegularCommunityItem", name, args == null ? RouteMapRuleEntrySetRegularCommunityItemArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RouteMapRuleEntrySetRegularCommunityItem(String name, Output<String> id, @Nullable RouteMapRuleEntrySetRegularCommunityItemState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/routeMapRuleEntrySetRegularCommunityItem:RouteMapRuleEntrySetRegularCommunityItem", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static RouteMapRuleEntrySetRegularCommunityItem get(String name, Output<String> id, @Nullable RouteMapRuleEntrySetRegularCommunityItemState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RouteMapRuleEntrySetRegularCommunityItem(name, id, state, options);
    }
}