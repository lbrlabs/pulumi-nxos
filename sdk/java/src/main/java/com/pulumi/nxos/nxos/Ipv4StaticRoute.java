// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.nxos.Ipv4StaticRouteArgs;
import com.pulumi.nxos.nxos.inputs.Ipv4StaticRouteState;
import com.pulumi.nxos.nxos.outputs.Ipv4StaticRouteNextHop;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="nxos:nxos/ipv4StaticRoute:Ipv4StaticRoute")
public class Ipv4StaticRoute extends com.pulumi.resources.CustomResource {
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
     * List of next hops.
     * 
     */
    @Export(name="nextHops", refs={List.class,Ipv4StaticRouteNextHop.class}, tree="[0,1]")
    private Output<List<Ipv4StaticRouteNextHop>> nextHops;

    /**
     * @return List of next hops.
     * 
     */
    public Output<List<Ipv4StaticRouteNextHop>> nextHops() {
        return this.nextHops;
    }
    /**
     * Prefix.
     * 
     */
    @Export(name="prefix", refs={String.class}, tree="[0]")
    private Output<String> prefix;

    /**
     * @return Prefix.
     * 
     */
    public Output<String> prefix() {
        return this.prefix;
    }
    /**
     * VRF name.
     * 
     */
    @Export(name="vrfName", refs={String.class}, tree="[0]")
    private Output<String> vrfName;

    /**
     * @return VRF name.
     * 
     */
    public Output<String> vrfName() {
        return this.vrfName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Ipv4StaticRoute(String name) {
        this(name, Ipv4StaticRouteArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Ipv4StaticRoute(String name, Ipv4StaticRouteArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Ipv4StaticRoute(String name, Ipv4StaticRouteArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/ipv4StaticRoute:Ipv4StaticRoute", name, args == null ? Ipv4StaticRouteArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Ipv4StaticRoute(String name, Output<String> id, @Nullable Ipv4StaticRouteState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/ipv4StaticRoute:Ipv4StaticRoute", name, state, makeResourceOptions(options, id));
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
    public static Ipv4StaticRoute get(String name, Output<String> id, @Nullable Ipv4StaticRouteState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Ipv4StaticRoute(name, id, state, options);
    }
}
