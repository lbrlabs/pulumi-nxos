// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.FeatureIsisArgs;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.inputs.FeatureIsisState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage the ISIS feature configuration.
 * 
 * - API Documentation: [fmIsis](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Isis/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.nxos.FeatureIsis;
 * import com.pulumi.nxos.FeatureIsisArgs;
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
 *         var example = new FeatureIsis(&#34;example&#34;, FeatureIsisArgs.builder()        
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
 *  $ pulumi import nxos:index/featureIsis:FeatureIsis example &#34;sys/fm/isis&#34;
 * ```
 * 
 */
@ResourceType(type="nxos:index/featureIsis:FeatureIsis")
public class FeatureIsis extends com.pulumi.resources.CustomResource {
    /**
     * Administrative state. - Choices: `enabled`, `disabled`
     * 
     */
    @Export(name="adminState", refs={String.class}, tree="[0]")
    private Output<String> adminState;

    /**
     * @return Administrative state. - Choices: `enabled`, `disabled`
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
    public FeatureIsis(String name) {
        this(name, FeatureIsisArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FeatureIsis(String name, FeatureIsisArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FeatureIsis(String name, FeatureIsisArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/featureIsis:FeatureIsis", name, args == null ? FeatureIsisArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FeatureIsis(String name, Output<String> id, @Nullable FeatureIsisState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/featureIsis:FeatureIsis", name, state, makeResourceOptions(options, id));
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
    public static FeatureIsis get(String name, Output<String> id, @Nullable FeatureIsisState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FeatureIsis(name, id, state, options);
    }
}
