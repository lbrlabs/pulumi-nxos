// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.EvpnVniRouteTargetArgs;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.inputs.EvpnVniRouteTargetState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage a EVPN VNI Route Target Entry.
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
 * import com.pulumi.nxos.EvpnVniRouteTarget;
 * import com.pulumi.nxos.EvpnVniRouteTargetArgs;
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
 *         var example = new EvpnVniRouteTarget(&#34;example&#34;, EvpnVniRouteTargetArgs.builder()        
 *             .direction(&#34;import&#34;)
 *             .encap(&#34;vxlan-123456&#34;)
 *             .routeTarget(&#34;route-target:as2-nn2:2:2&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import nxos:index/evpnVniRouteTarget:EvpnVniRouteTarget example &#34;sys/evpn/bdevi-[vxlan-123456]/rttp-[import]/ent-[route-target:as2-nn2:2:2]&#34;
 * ```
 * 
 */
@ResourceType(type="nxos:index/evpnVniRouteTarget:EvpnVniRouteTarget")
public class EvpnVniRouteTarget extends com.pulumi.resources.CustomResource {
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
     * Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     * 
     */
    @Export(name="encap", refs={String.class}, tree="[0]")
    private Output<String> encap;

    /**
     * @return Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     * 
     */
    public Output<String> encap() {
        return this.encap;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EvpnVniRouteTarget(String name) {
        this(name, EvpnVniRouteTargetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EvpnVniRouteTarget(String name, EvpnVniRouteTargetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EvpnVniRouteTarget(String name, EvpnVniRouteTargetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/evpnVniRouteTarget:EvpnVniRouteTarget", name, args == null ? EvpnVniRouteTargetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EvpnVniRouteTarget(String name, Output<String> id, @Nullable EvpnVniRouteTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/evpnVniRouteTarget:EvpnVniRouteTarget", name, state, makeResourceOptions(options, id));
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
    public static EvpnVniRouteTarget get(String name, Output<String> id, @Nullable EvpnVniRouteTargetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EvpnVniRouteTarget(name, id, state, options);
    }
}