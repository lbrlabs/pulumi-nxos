// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.OspfAuthenticationArgs;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.inputs.OspfAuthenticationState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage the OSPF authentication configuration.
 * 
 * - API Documentation: [ospfAuthNewP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/ospf:AuthNewP/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.nxos.OspfAuthentication;
 * import com.pulumi.nxos.OspfAuthenticationArgs;
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
 *         var example = new OspfAuthentication(&#34;example&#34;, OspfAuthenticationArgs.builder()        
 *             .instanceName(&#34;OSPF1&#34;)
 *             .interfaceId(&#34;eth1/10&#34;)
 *             .key(&#34;0 mykey&#34;)
 *             .keyId(1)
 *             .keySecureMode(false)
 *             .keychain(&#34;mykeychain&#34;)
 *             .md5Key(&#34;0 mymd5key&#34;)
 *             .md5KeySecureMode(false)
 *             .type(&#34;none&#34;)
 *             .vrfName(&#34;VRF1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import nxos:index/ospfAuthentication:OspfAuthentication example &#34;sys/ospf/inst-[OSPF1]/dom-[VRF1]/if-[eth1/10]/authnew&#34;
 * ```
 * 
 */
@ResourceType(type="nxos:index/ospfAuthentication:OspfAuthentication")
public class OspfAuthentication extends com.pulumi.resources.CustomResource {
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
     * OSPF instance name.
     * 
     */
    @Export(name="instanceName", refs={String.class}, tree="[0]")
    private Output<String> instanceName;

    /**
     * @return OSPF instance name.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
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
     * Key used for authentication.
     * 
     */
    @Export(name="key", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> key;

    /**
     * @return Key used for authentication.
     * 
     */
    public Output<Optional<String>> key() {
        return Codegen.optional(this.key);
    }
    /**
     * Key ID used for authentication. - Range: `0`-`255` - Default value: `0`
     * 
     */
    @Export(name="keyId", refs={Integer.class}, tree="[0]")
    private Output<Integer> keyId;

    /**
     * @return Key ID used for authentication. - Range: `0`-`255` - Default value: `0`
     * 
     */
    public Output<Integer> keyId() {
        return this.keyId;
    }
    /**
     * Encrypted authentication key or plain text key. - Default value: `false`
     * 
     */
    @Export(name="keySecureMode", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> keySecureMode;

    /**
     * @return Encrypted authentication key or plain text key. - Default value: `false`
     * 
     */
    public Output<Boolean> keySecureMode() {
        return this.keySecureMode;
    }
    /**
     * Authentication keychain.
     * 
     */
    @Export(name="keychain", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> keychain;

    /**
     * @return Authentication keychain.
     * 
     */
    public Output<Optional<String>> keychain() {
        return Codegen.optional(this.keychain);
    }
    /**
     * Key used for md5 authentication.
     * 
     */
    @Export(name="md5Key", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> md5Key;

    /**
     * @return Key used for md5 authentication.
     * 
     */
    public Output<Optional<String>> md5Key() {
        return Codegen.optional(this.md5Key);
    }
    /**
     * Encrypted authentication md5 key or plain text key. - Default value: `false`
     * 
     */
    @Export(name="md5KeySecureMode", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> md5KeySecureMode;

    /**
     * @return Encrypted authentication md5 key or plain text key. - Default value: `false`
     * 
     */
    public Output<Boolean> md5KeySecureMode() {
        return this.md5KeySecureMode;
    }
    /**
     * Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
     * 
     */
    public Output<String> type() {
        return this.type;
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
    public OspfAuthentication(String name) {
        this(name, OspfAuthenticationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OspfAuthentication(String name, OspfAuthenticationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OspfAuthentication(String name, OspfAuthenticationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/ospfAuthentication:OspfAuthentication", name, args == null ? OspfAuthenticationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private OspfAuthentication(String name, Output<String> id, @Nullable OspfAuthenticationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/ospfAuthentication:OspfAuthentication", name, state, makeResourceOptions(options, id));
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
    public static OspfAuthentication get(String name, Output<String> id, @Nullable OspfAuthenticationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OspfAuthentication(name, id, state, options);
    }
}