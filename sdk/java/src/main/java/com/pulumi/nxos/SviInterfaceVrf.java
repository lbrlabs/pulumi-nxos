// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.SviInterfaceVrfArgs;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.inputs.SviInterfaceVrfState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage an SVI interface VRF association.
 * 
 * - API Documentation: [nwRtVrfMbr](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/nw:RtVrfMbr/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.nxos.SviInterfaceVrf;
 * import com.pulumi.nxos.SviInterfaceVrfArgs;
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
 *         var example = new SviInterfaceVrf(&#34;example&#34;, SviInterfaceVrfArgs.builder()        
 *             .interfaceId(&#34;vlan293&#34;)
 *             .vrfDn(&#34;sys/inst-VRF123&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import nxos:index/sviInterfaceVrf:SviInterfaceVrf example &#34;sys/intf/svi-[vlan293]/rtvrfMbr&#34;
 * ```
 * 
 */
@ResourceType(type="nxos:index/sviInterfaceVrf:SviInterfaceVrf")
public class SviInterfaceVrf extends com.pulumi.resources.CustomResource {
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
     * DN of VRF. For example: `sys/inst-VRF1`.
     * 
     */
    @Export(name="vrfDn", refs={String.class}, tree="[0]")
    private Output<String> vrfDn;

    /**
     * @return DN of VRF. For example: `sys/inst-VRF1`.
     * 
     */
    public Output<String> vrfDn() {
        return this.vrfDn;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SviInterfaceVrf(String name) {
        this(name, SviInterfaceVrfArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SviInterfaceVrf(String name, SviInterfaceVrfArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SviInterfaceVrf(String name, SviInterfaceVrfArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/sviInterfaceVrf:SviInterfaceVrf", name, args == null ? SviInterfaceVrfArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SviInterfaceVrf(String name, Output<String> id, @Nullable SviInterfaceVrfState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/sviInterfaceVrf:SviInterfaceVrf", name, state, makeResourceOptions(options, id));
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
    public static SviInterfaceVrf get(String name, Output<String> id, @Nullable SviInterfaceVrfState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SviInterfaceVrf(name, id, state, options);
    }
}
