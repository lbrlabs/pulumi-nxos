// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.nxos.NveVniIngressReplicationArgs;
import com.pulumi.nxos.nxos.inputs.NveVniIngressReplicationState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="nxos:nxos/nveVniIngressReplication:NveVniIngressReplication")
public class NveVniIngressReplication extends com.pulumi.resources.CustomResource {
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
     * Configure VxLAN Ingress Replication mode. - Choices: `bgp`, `unknown`, `static` - Default value: `unknown`
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output<String> protocol;

    /**
     * @return Configure VxLAN Ingress Replication mode. - Choices: `bgp`, `unknown`, `static` - Default value: `unknown`
     * 
     */
    public Output<String> protocol() {
        return this.protocol;
    }
    /**
     * Virtual Network ID. - Range: `1`-`16777214`
     * 
     */
    @Export(name="vni", refs={Integer.class}, tree="[0]")
    private Output<Integer> vni;

    /**
     * @return Virtual Network ID. - Range: `1`-`16777214`
     * 
     */
    public Output<Integer> vni() {
        return this.vni;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NveVniIngressReplication(String name) {
        this(name, NveVniIngressReplicationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NveVniIngressReplication(String name, NveVniIngressReplicationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NveVniIngressReplication(String name, NveVniIngressReplicationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/nveVniIngressReplication:NveVniIngressReplication", name, args == null ? NveVniIngressReplicationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NveVniIngressReplication(String name, Output<String> id, @Nullable NveVniIngressReplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/nveVniIngressReplication:NveVniIngressReplication", name, state, makeResourceOptions(options, id));
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
    public static NveVniIngressReplication get(String name, Output<String> id, @Nullable NveVniIngressReplicationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NveVniIngressReplication(name, id, state, options);
    }
}