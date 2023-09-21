// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.nxos.VrfRouteTargetArgs;
import com.pulumi.nxos.nxos.inputs.VrfRouteTargetState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="nxos:nxos/vrfRouteTarget:VrfRouteTarget")
public class VrfRouteTarget extends com.pulumi.resources.CustomResource {
    /**
     * Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
     * 
     */
    @Export(name="addressFamily", refs={String.class}, tree="[0]")
    private Output<String> addressFamily;

    /**
     * @return Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
     * 
     */
    public Output<String> addressFamily() {
        return this.addressFamily;
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
     * Route Target direction. - Choices: `import`, `export`
     * 
     */
    @Export(name="direction", refs={String.class}, tree="[0]")
    private Output<String> direction;

    /**
     * @return Route Target direction. - Choices: `import`, `export`
     * 
     */
    public Output<String> direction() {
        return this.direction;
    }
    /**
     * Route Target in NX-OS DME format.
     * 
     */
    @Export(name="routeTarget", refs={String.class}, tree="[0]")
    private Output<String> routeTarget;

    /**
     * @return Route Target in NX-OS DME format.
     * 
     */
    public Output<String> routeTarget() {
        return this.routeTarget;
    }
    /**
     * Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
     * 
     */
    @Export(name="routeTargetAddressFamily", refs={String.class}, tree="[0]")
    private Output<String> routeTargetAddressFamily;

    /**
     * @return Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
     * 
     */
    public Output<String> routeTargetAddressFamily() {
        return this.routeTargetAddressFamily;
    }
    /**
     * VRF name.
     * 
     */
    @Export(name="vrf", refs={String.class}, tree="[0]")
    private Output<String> vrf;

    /**
     * @return VRF name.
     * 
     */
    public Output<String> vrf() {
        return this.vrf;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VrfRouteTarget(String name) {
        this(name, VrfRouteTargetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VrfRouteTarget(String name, VrfRouteTargetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VrfRouteTarget(String name, VrfRouteTargetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/vrfRouteTarget:VrfRouteTarget", name, args == null ? VrfRouteTargetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VrfRouteTarget(String name, Output<String> id, @Nullable VrfRouteTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/vrfRouteTarget:VrfRouteTarget", name, state, makeResourceOptions(options, id));
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
    public static VrfRouteTarget get(String name, Output<String> id, @Nullable VrfRouteTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VrfRouteTarget(name, id, state, options);
    }
}