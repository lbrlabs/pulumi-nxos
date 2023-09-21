// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.nxos.Ipv4AccessListPolicyIngressInterfaceArgs;
import com.pulumi.nxos.nxos.inputs.Ipv4AccessListPolicyIngressInterfaceState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="nxos:nxos/ipv4AccessListPolicyIngressInterface:Ipv4AccessListPolicyIngressInterface")
public class Ipv4AccessListPolicyIngressInterface extends com.pulumi.resources.CustomResource {
    /**
     * Access list name.
     * 
     */
    @Export(name="accessListName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> accessListName;

    /**
     * @return Access list name.
     * 
     */
    public Output<Optional<String>> accessListName() {
        return Codegen.optional(this.accessListName);
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
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     * 
     */
    @Export(name="interfaceId", refs={String.class}, tree="[0]")
    private Output<String> interfaceId;

    /**
     * @return Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     * 
     */
    public Output<String> interfaceId() {
        return this.interfaceId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Ipv4AccessListPolicyIngressInterface(String name) {
        this(name, Ipv4AccessListPolicyIngressInterfaceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Ipv4AccessListPolicyIngressInterface(String name, Ipv4AccessListPolicyIngressInterfaceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Ipv4AccessListPolicyIngressInterface(String name, Ipv4AccessListPolicyIngressInterfaceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/ipv4AccessListPolicyIngressInterface:Ipv4AccessListPolicyIngressInterface", name, args == null ? Ipv4AccessListPolicyIngressInterfaceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Ipv4AccessListPolicyIngressInterface(String name, Output<String> id, @Nullable Ipv4AccessListPolicyIngressInterfaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/ipv4AccessListPolicyIngressInterface:Ipv4AccessListPolicyIngressInterface", name, state, makeResourceOptions(options, id));
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
    public static Ipv4AccessListPolicyIngressInterface get(String name, Output<String> id, @Nullable Ipv4AccessListPolicyIngressInterfaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Ipv4AccessListPolicyIngressInterface(name, id, state, options);
    }
}
