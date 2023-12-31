// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.VrfRouteTargetArgs;
import com.pulumi.nxos.inputs.VrfRouteTargetState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage a VRF Route Target Entry.
 * 
 * - API Documentation: [rtctrlRttEntry](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:RttEntry/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.nxos.VrfRouteTarget;
 * import com.pulumi.nxos.VrfRouteTargetArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new VrfRouteTarget(&#34;example&#34;, VrfRouteTargetArgs.builder()        
 *             .addressFamily(&#34;ipv4-ucast&#34;)
 *             .direction(&#34;import&#34;)
 *             .routeTarget(&#34;route-target:as2-nn2:2:2&#34;)
 *             .routeTargetAddressFamily(&#34;ipv4-ucast&#34;)
 *             .vrf(&#34;VRF1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import nxos:index/vrfRouteTarget:VrfRouteTarget example &#34;sys/inst-[VRF1]/dom-[VRF1]/af-[ipv4-ucast]/ctrl-[ipv4-ucast]/rttp-[import]/ent-[route-target:as2-nn2:2:2]&#34;
 * ```
 * 
 */
@ResourceType(type="nxos:index/vrfRouteTarget:VrfRouteTarget")
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
        super("nxos:index/vrfRouteTarget:VrfRouteTarget", name, args == null ? VrfRouteTargetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VrfRouteTarget(String name, Output<String> id, @Nullable VrfRouteTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/vrfRouteTarget:VrfRouteTarget", name, state, makeResourceOptions(options, id));
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
