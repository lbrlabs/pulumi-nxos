// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.nxos.VrfRouteTargetAddressFamilyArgs;
import com.pulumi.nxos.nxos.inputs.VrfRouteTargetAddressFamilyState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="nxos:nxos/vrfRouteTargetAddressFamily:VrfRouteTargetAddressFamily")
public class VrfRouteTargetAddressFamily extends com.pulumi.resources.CustomResource {
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
    public VrfRouteTargetAddressFamily(String name) {
        this(name, VrfRouteTargetAddressFamilyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VrfRouteTargetAddressFamily(String name, VrfRouteTargetAddressFamilyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VrfRouteTargetAddressFamily(String name, VrfRouteTargetAddressFamilyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/vrfRouteTargetAddressFamily:VrfRouteTargetAddressFamily", name, args == null ? VrfRouteTargetAddressFamilyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VrfRouteTargetAddressFamily(String name, Output<String> id, @Nullable VrfRouteTargetAddressFamilyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/vrfRouteTargetAddressFamily:VrfRouteTargetAddressFamily", name, state, makeResourceOptions(options, id));
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
    public static VrfRouteTargetAddressFamily get(String name, Output<String> id, @Nullable VrfRouteTargetAddressFamilyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VrfRouteTargetAddressFamily(name, id, state, options);
    }
}
