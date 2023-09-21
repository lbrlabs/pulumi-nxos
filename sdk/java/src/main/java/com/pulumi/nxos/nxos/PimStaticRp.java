// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.nxos.PimStaticRpArgs;
import com.pulumi.nxos.nxos.inputs.PimStaticRpState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="nxos:nxos/pimStaticRp:PimStaticRp")
public class PimStaticRp extends com.pulumi.resources.CustomResource {
    /**
     * Address.
     * 
     */
    @Export(name="address", refs={String.class}, tree="[0]")
    private Output<String> address;

    /**
     * @return Address.
     * 
     */
    public Output<String> address() {
        return this.address;
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
    public PimStaticRp(String name) {
        this(name, PimStaticRpArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PimStaticRp(String name, PimStaticRpArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PimStaticRp(String name, PimStaticRpArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/pimStaticRp:PimStaticRp", name, args == null ? PimStaticRpArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PimStaticRp(String name, Output<String> id, @Nullable PimStaticRpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/pimStaticRp:PimStaticRp", name, state, makeResourceOptions(options, id));
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
    public static PimStaticRp get(String name, Output<String> id, @Nullable PimStaticRpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PimStaticRp(name, id, state, options);
    }
}
