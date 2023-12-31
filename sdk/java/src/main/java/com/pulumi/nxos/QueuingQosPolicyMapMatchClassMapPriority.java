// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.QueuingQosPolicyMapMatchClassMapPriorityArgs;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.inputs.QueuingQosPolicyMapMatchClassMapPriorityState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage the queuing QoS policy map match class map priority configuration.
 * 
 * - API Documentation: [ipqosPriority](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:Priority/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.nxos.QueuingQosPolicyMapMatchClassMapPriority;
 * import com.pulumi.nxos.QueuingQosPolicyMapMatchClassMapPriorityArgs;
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
 *         var example = new QueuingQosPolicyMapMatchClassMapPriority(&#34;example&#34;, QueuingQosPolicyMapMatchClassMapPriorityArgs.builder()        
 *             .classMapName(&#34;c-out-q1&#34;)
 *             .level(1)
 *             .policyMapName(&#34;PM1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import nxos:index/queuingQosPolicyMapMatchClassMapPriority:QueuingQosPolicyMapMatchClassMapPriority example &#34;sys/ipqos/queuing/p/name-[PM1]/cmap-[c-out-q1]/prio&#34;
 * ```
 * 
 */
@ResourceType(type="nxos:index/queuingQosPolicyMapMatchClassMapPriority:QueuingQosPolicyMapMatchClassMapPriority")
public class QueuingQosPolicyMapMatchClassMapPriority extends com.pulumi.resources.CustomResource {
    /**
     * Class map name.
     * 
     */
    @Export(name="classMapName", refs={String.class}, tree="[0]")
    private Output<String> classMapName;

    /**
     * @return Class map name.
     * 
     */
    public Output<String> classMapName() {
        return this.classMapName;
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
     * Priority level. - Range: `1`-`8`
     * 
     */
    @Export(name="level", refs={Integer.class}, tree="[0]")
    private Output<Integer> level;

    /**
     * @return Priority level. - Range: `1`-`8`
     * 
     */
    public Output<Integer> level() {
        return this.level;
    }
    /**
     * Policy map name.
     * 
     */
    @Export(name="policyMapName", refs={String.class}, tree="[0]")
    private Output<String> policyMapName;

    /**
     * @return Policy map name.
     * 
     */
    public Output<String> policyMapName() {
        return this.policyMapName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public QueuingQosPolicyMapMatchClassMapPriority(String name) {
        this(name, QueuingQosPolicyMapMatchClassMapPriorityArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public QueuingQosPolicyMapMatchClassMapPriority(String name, QueuingQosPolicyMapMatchClassMapPriorityArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public QueuingQosPolicyMapMatchClassMapPriority(String name, QueuingQosPolicyMapMatchClassMapPriorityArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/queuingQosPolicyMapMatchClassMapPriority:QueuingQosPolicyMapMatchClassMapPriority", name, args == null ? QueuingQosPolicyMapMatchClassMapPriorityArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private QueuingQosPolicyMapMatchClassMapPriority(String name, Output<String> id, @Nullable QueuingQosPolicyMapMatchClassMapPriorityState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/queuingQosPolicyMapMatchClassMapPriority:QueuingQosPolicyMapMatchClassMapPriority", name, state, makeResourceOptions(options, id));
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
    public static QueuingQosPolicyMapMatchClassMapPriority get(String name, Output<String> id, @Nullable QueuingQosPolicyMapMatchClassMapPriorityState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new QueuingQosPolicyMapMatchClassMapPriority(name, id, state, options);
    }
}
