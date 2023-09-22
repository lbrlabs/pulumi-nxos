// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.OspfArgs;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.inputs.OspfState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage the global OSPF configuration.
 * 
 * - API Documentation: [ospfEntity](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/ospf:Entity/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.nxos.Ospf;
 * import com.pulumi.nxos.OspfArgs;
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
 *         var example = new Ospf(&#34;example&#34;, OspfArgs.builder()        
 *             .adminState(&#34;enabled&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import nxos:index/ospf:Ospf example &#34;sys/ospf&#34;
 * ```
 * 
 */
@ResourceType(type="nxos:index/ospf:Ospf")
public class Ospf extends com.pulumi.resources.CustomResource {
    /**
     * Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     * 
     */
    @Export(name="adminState", refs={String.class}, tree="[0]")
    private Output<String> adminState;

    /**
     * @return Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     * 
     */
    public Output<String> adminState() {
        return this.adminState;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Ospf(String name) {
        this(name, OspfArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Ospf(String name, @Nullable OspfArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Ospf(String name, @Nullable OspfArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/ospf:Ospf", name, args == null ? OspfArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Ospf(String name, Output<String> id, @Nullable OspfState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/ospf:Ospf", name, state, makeResourceOptions(options, id));
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
    public static Ospf get(String name, Output<String> id, @Nullable OspfState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Ospf(name, id, state, options);
    }
}