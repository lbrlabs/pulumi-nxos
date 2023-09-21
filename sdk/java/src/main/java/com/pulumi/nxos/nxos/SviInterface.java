// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.nxos.SviInterfaceArgs;
import com.pulumi.nxos.nxos.inputs.SviInterfaceState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="nxos:nxos/sviInterface:SviInterface")
public class SviInterface extends com.pulumi.resources.CustomResource {
    /**
     * Administrative port state. - Choices: `up`, `down` - Default value: `up`
     * 
     */
    @Export(name="adminState", refs={String.class}, tree="[0]")
    private Output<String> adminState;

    /**
     * @return Administrative port state. - Choices: `up`, `down` - Default value: `up`
     * 
     */
    public Output<String> adminState() {
        return this.adminState;
    }
    /**
     * Specifies the administrative port bandwidth. - Range: `1`-`400000000` - Default value: `1000000`
     * 
     */
    @Export(name="bandwidth", refs={Integer.class}, tree="[0]")
    private Output<Integer> bandwidth;

    /**
     * @return Specifies the administrative port bandwidth. - Range: `1`-`400000000` - Default value: `1000000`
     * 
     */
    public Output<Integer> bandwidth() {
        return this.bandwidth;
    }
    /**
     * Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
     * 
     */
    @Export(name="delay", refs={Integer.class}, tree="[0]")
    private Output<Integer> delay;

    /**
     * @return Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
     * 
     */
    public Output<Integer> delay() {
        return this.delay;
    }
    /**
     * Interface description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Interface description.
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
     * Must match first field in the output of `show intf brief`. Example: `vlan100`.
     * 
     */
    @Export(name="interfaceId", refs={String.class}, tree="[0]")
    private Output<String> interfaceId;

    /**
     * @return Must match first field in the output of `show intf brief`. Example: `vlan100`.
     * 
     */
    public Output<String> interfaceId() {
        return this.interfaceId;
    }
    /**
     * The administrative port medium type. - Choices: `bcast`, `p2p` - Default value: `bcast`
     * 
     */
    @Export(name="medium", refs={String.class}, tree="[0]")
    private Output<String> medium;

    /**
     * @return The administrative port medium type. - Choices: `bcast`, `p2p` - Default value: `bcast`
     * 
     */
    public Output<String> medium() {
        return this.medium;
    }
    /**
     * Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
     * 
     */
    @Export(name="mtu", refs={Integer.class}, tree="[0]")
    private Output<Integer> mtu;

    /**
     * @return Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
     * 
     */
    public Output<Integer> mtu() {
        return this.mtu;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SviInterface(String name) {
        this(name, SviInterfaceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SviInterface(String name, SviInterfaceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SviInterface(String name, SviInterfaceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/sviInterface:SviInterface", name, args == null ? SviInterfaceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SviInterface(String name, Output<String> id, @Nullable SviInterfaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/sviInterface:SviInterface", name, state, makeResourceOptions(options, id));
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
    public static SviInterface get(String name, Output<String> id, @Nullable SviInterfaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SviInterface(name, id, state, options);
    }
}
