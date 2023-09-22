// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.VrfArgs;
import com.pulumi.nxos.inputs.VrfState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage a VRF.
 * 
 * - API Documentation: [l3Inst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/l3:Inst/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.nxos.Vrf;
 * import com.pulumi.nxos.VrfArgs;
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
 *         var example = new Vrf(&#34;example&#34;, VrfArgs.builder()        
 *             .description(&#34;My VRF1 Description&#34;)
 *             .encap(&#34;vxlan-103901&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import nxos:index/vrf:Vrf example &#34;sys/inst-[VRF1]&#34;
 * ```
 * 
 */
@ResourceType(type="nxos:index/vrf:Vrf")
public class Vrf extends com.pulumi.resources.CustomResource {
    /**
     * VRF description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return VRF description.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
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
     * Encap for this Context, supported formats: `unknown`, `vlan-%d` or `vxlan-%d`. - Default value: `unknown`
     * 
     */
    @Export(name="encap", refs={String.class}, tree="[0]")
    private Output<String> encap;

    /**
     * @return Encap for this Context, supported formats: `unknown`, `vlan-%d` or `vxlan-%d`. - Default value: `unknown`
     * 
     */
    public Output<String> encap() {
        return this.encap;
    }
    /**
     * VRF name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return VRF name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Vrf(String name) {
        this(name, VrfArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Vrf(String name, @Nullable VrfArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Vrf(String name, @Nullable VrfArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/vrf:Vrf", name, args == null ? VrfArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Vrf(String name, Output<String> id, @Nullable VrfState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/vrf:Vrf", name, state, makeResourceOptions(options, id));
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
    public static Vrf get(String name, Output<String> id, @Nullable VrfState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Vrf(name, id, state, options);
    }
}