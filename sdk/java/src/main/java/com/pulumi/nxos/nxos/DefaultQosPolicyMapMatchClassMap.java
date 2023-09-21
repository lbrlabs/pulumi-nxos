// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.nxos.DefaultQosPolicyMapMatchClassMapArgs;
import com.pulumi.nxos.nxos.inputs.DefaultQosPolicyMapMatchClassMapState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="nxos:nxos/defaultQosPolicyMapMatchClassMap:DefaultQosPolicyMapMatchClassMap")
public class DefaultQosPolicyMapMatchClassMap extends com.pulumi.resources.CustomResource {
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
     * Class map name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Class map name.
     * 
     */
    public Output<String> name() {
        return this.name;
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
    public DefaultQosPolicyMapMatchClassMap(String name) {
        this(name, DefaultQosPolicyMapMatchClassMapArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DefaultQosPolicyMapMatchClassMap(String name, DefaultQosPolicyMapMatchClassMapArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DefaultQosPolicyMapMatchClassMap(String name, DefaultQosPolicyMapMatchClassMapArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/defaultQosPolicyMapMatchClassMap:DefaultQosPolicyMapMatchClassMap", name, args == null ? DefaultQosPolicyMapMatchClassMapArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DefaultQosPolicyMapMatchClassMap(String name, Output<String> id, @Nullable DefaultQosPolicyMapMatchClassMapState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/defaultQosPolicyMapMatchClassMap:DefaultQosPolicyMapMatchClassMap", name, state, makeResourceOptions(options, id));
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
    public static DefaultQosPolicyMapMatchClassMap get(String name, Output<String> id, @Nullable DefaultQosPolicyMapMatchClassMapState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DefaultQosPolicyMapMatchClassMap(name, id, state, options);
    }
}
