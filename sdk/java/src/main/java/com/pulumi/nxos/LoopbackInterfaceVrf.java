// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.LoopbackInterfaceVrfArgs;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.inputs.LoopbackInterfaceVrfState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage a loopback interface VRF association.
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
 * import com.pulumi.nxos.LoopbackInterfaceVrf;
 * import com.pulumi.nxos.LoopbackInterfaceVrfArgs;
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
 *         var example = new LoopbackInterfaceVrf(&#34;example&#34;, LoopbackInterfaceVrfArgs.builder()        
 *             .interfaceId(&#34;lo123&#34;)
 *             .vrfDn(&#34;sys/inst-default&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import nxos:index/loopbackInterfaceVrf:LoopbackInterfaceVrf example &#34;sys/intf/lb-[lo123]/rtvrfMbr&#34;
 * ```
 * 
 */
@ResourceType(type="nxos:index/loopbackInterfaceVrf:LoopbackInterfaceVrf")
public class LoopbackInterfaceVrf extends com.pulumi.resources.CustomResource {
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
     * Must match first field in the output of `show intf brief`. Example: `lo123`.
     * 
     */
    @Export(name="interfaceId", refs={String.class}, tree="[0]")
    private Output<String> interfaceId;

    /**
     * @return Must match first field in the output of `show intf brief`. Example: `lo123`.
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
    public LoopbackInterfaceVrf(String name) {
        this(name, LoopbackInterfaceVrfArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LoopbackInterfaceVrf(String name, LoopbackInterfaceVrfArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LoopbackInterfaceVrf(String name, LoopbackInterfaceVrfArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/loopbackInterfaceVrf:LoopbackInterfaceVrf", name, args == null ? LoopbackInterfaceVrfArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private LoopbackInterfaceVrf(String name, Output<String> id, @Nullable LoopbackInterfaceVrfState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/loopbackInterfaceVrf:LoopbackInterfaceVrf", name, state, makeResourceOptions(options, id));
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
    public static LoopbackInterfaceVrf get(String name, Output<String> id, @Nullable LoopbackInterfaceVrfState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LoopbackInterfaceVrf(name, id, state, options);
    }
}
