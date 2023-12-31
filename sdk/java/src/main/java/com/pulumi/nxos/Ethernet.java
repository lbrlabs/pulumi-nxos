// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.EthernetArgs;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.inputs.EthernetState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage global ethernet settings.
 * 
 * - API Documentation: [ethpmInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Interfaces/ethpm:Inst/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.nxos.Ethernet;
 * import com.pulumi.nxos.EthernetArgs;
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
 *         var example = new Ethernet(&#34;example&#34;, EthernetArgs.builder()        
 *             .mtu(9216)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import nxos:index/ethernet:Ethernet example &#34;sys/ethpm/inst&#34;
 * ```
 * 
 */
@ResourceType(type="nxos:index/ethernet:Ethernet")
public class Ethernet extends com.pulumi.resources.CustomResource {
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
     * System jumbo MTU. - Range: `576`-`9216` - Default value: `9216`
     * 
     */
    @Export(name="mtu", refs={Integer.class}, tree="[0]")
    private Output<Integer> mtu;

    /**
     * @return System jumbo MTU. - Range: `576`-`9216` - Default value: `9216`
     * 
     */
    public Output<Integer> mtu() {
        return this.mtu;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Ethernet(String name) {
        this(name, EthernetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Ethernet(String name, @Nullable EthernetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Ethernet(String name, @Nullable EthernetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/ethernet:Ethernet", name, args == null ? EthernetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Ethernet(String name, Output<String> id, @Nullable EthernetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/ethernet:Ethernet", name, state, makeResourceOptions(options, id));
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
    public static Ethernet get(String name, Output<String> id, @Nullable EthernetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Ethernet(name, id, state, options);
    }
}
