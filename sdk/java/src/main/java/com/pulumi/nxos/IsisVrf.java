// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.IsisVrfArgs;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.inputs.IsisVrfState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * This resource can manage the IS-IS VRF configuration.
 * 
 * - API Documentation: [isisDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/isis:Dom/)
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.nxos.IsisVrf;
 * import com.pulumi.nxos.IsisVrfArgs;
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
 *         var example = new IsisVrf(&#34;example&#34;, IsisVrfArgs.builder()        
 *             .adminState(&#34;enabled&#34;)
 *             .authenticationCheckL1(false)
 *             .authenticationCheckL2(false)
 *             .authenticationKeyL1(&#34;&#34;)
 *             .authenticationKeyL2(&#34;&#34;)
 *             .authenticationTypeL1(&#34;unknown&#34;)
 *             .authenticationTypeL2(&#34;unknown&#34;)
 *             .bandwidthReference(400000)
 *             .banwidthReferenceUnit(&#34;mbps&#34;)
 *             .instanceName(&#34;ISIS1&#34;)
 *             .isType(&#34;l2&#34;)
 *             .metricType(&#34;wide&#34;)
 *             .mtu(2000)
 *             .net(&#34;49.0001.0000.0000.3333.00&#34;)
 *             .passiveDefault(&#34;l12&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ```sh
 *  $ pulumi import nxos:index/isisVrf:IsisVrf example &#34;sys/isis/inst-[ISIS1]/dom-[default]&#34;
 * ```
 * 
 */
@ResourceType(type="nxos:index/isisVrf:IsisVrf")
public class IsisVrf extends com.pulumi.resources.CustomResource {
    /**
     * Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     * 
     */
    @Export(name="adminState", refs={String.class}, tree="[0]")
    private Output<String> adminState;

    /**
     * @return Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     * 
     */
    public Output<String> adminState() {
        return this.adminState;
    }
    /**
     * Authentication Check for ISIS on Level-1. - Default value: `true`
     * 
     */
    @Export(name="authenticationCheckL1", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> authenticationCheckL1;

    /**
     * @return Authentication Check for ISIS on Level-1. - Default value: `true`
     * 
     */
    public Output<Boolean> authenticationCheckL1() {
        return this.authenticationCheckL1;
    }
    /**
     * Authentication Check for ISIS on Level-2. - Default value: `true`
     * 
     */
    @Export(name="authenticationCheckL2", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> authenticationCheckL2;

    /**
     * @return Authentication Check for ISIS on Level-2. - Default value: `true`
     * 
     */
    public Output<Boolean> authenticationCheckL2() {
        return this.authenticationCheckL2;
    }
    /**
     * Authentication Key for IS-IS on Level-1.
     * 
     */
    @Export(name="authenticationKeyL1", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authenticationKeyL1;

    /**
     * @return Authentication Key for IS-IS on Level-1.
     * 
     */
    public Output<Optional<String>> authenticationKeyL1() {
        return Codegen.optional(this.authenticationKeyL1);
    }
    /**
     * Authentication Key for IS-IS on Level-2.
     * 
     */
    @Export(name="authenticationKeyL2", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authenticationKeyL2;

    /**
     * @return Authentication Key for IS-IS on Level-2.
     * 
     */
    public Output<Optional<String>> authenticationKeyL2() {
        return Codegen.optional(this.authenticationKeyL2);
    }
    /**
     * IS-IS Authentication-Type for Level-1. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
     * 
     */
    @Export(name="authenticationTypeL1", refs={String.class}, tree="[0]")
    private Output<String> authenticationTypeL1;

    /**
     * @return IS-IS Authentication-Type for Level-1. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
     * 
     */
    public Output<String> authenticationTypeL1() {
        return this.authenticationTypeL1;
    }
    /**
     * IS-IS Authentication-Type for Level-2. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
     * 
     */
    @Export(name="authenticationTypeL2", refs={String.class}, tree="[0]")
    private Output<String> authenticationTypeL2;

    /**
     * @return IS-IS Authentication-Type for Level-2. - Choices: `clear`, `md5`, `unknown` - Default value: `unknown`
     * 
     */
    public Output<String> authenticationTypeL2() {
        return this.authenticationTypeL2;
    }
    /**
     * The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost
     * metric. - Range: `0`-`4294967295` - Default value: `40000`
     * 
     */
    @Export(name="bandwidthReference", refs={Integer.class}, tree="[0]")
    private Output<Integer> bandwidthReference;

    /**
     * @return The IS-IS domain bandwidth reference. This sets the default reference bandwidth used for calculating the IS-IS cost
     * metric. - Range: `0`-`4294967295` - Default value: `40000`
     * 
     */
    public Output<Integer> bandwidthReference() {
        return this.bandwidthReference;
    }
    /**
     * Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
     * 
     */
    @Export(name="banwidthReferenceUnit", refs={String.class}, tree="[0]")
    private Output<String> banwidthReferenceUnit;

    /**
     * @return Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
     * 
     */
    public Output<String> banwidthReferenceUnit() {
        return this.banwidthReferenceUnit;
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
     * IS-IS instance name.
     * 
     */
    @Export(name="instanceName", refs={String.class}, tree="[0]")
    private Output<String> instanceName;

    /**
     * @return IS-IS instance name.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
    }
    /**
     * IS-IS domain type. - Choices: `l1`, `l2`, `l12` - Default value: `l12`
     * 
     */
    @Export(name="isType", refs={String.class}, tree="[0]")
    private Output<String> isType;

    /**
     * @return IS-IS domain type. - Choices: `l1`, `l2`, `l12` - Default value: `l12`
     * 
     */
    public Output<String> isType() {
        return this.isType;
    }
    /**
     * IS-IS metric type. - Choices: `narrow`, `wide`, `transition` - Default value: `wide`
     * 
     */
    @Export(name="metricType", refs={String.class}, tree="[0]")
    private Output<String> metricType;

    /**
     * @return IS-IS metric type. - Choices: `narrow`, `wide`, `transition` - Default value: `wide`
     * 
     */
    public Output<String> metricType() {
        return this.metricType;
    }
    /**
     * The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352
     * bytes. - Range: `256`-`4352` - Default value: `1492`
     * 
     */
    @Export(name="mtu", refs={Integer.class}, tree="[0]")
    private Output<Integer> mtu;

    /**
     * @return The configuration of link-state packet (LSP) maximum transmission units (MTU) is supported. You can enable up to 4352
     * bytes. - Range: `256`-`4352` - Default value: `1492`
     * 
     */
    public Output<Integer> mtu() {
        return this.mtu;
    }
    /**
     * VRF name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return VRF name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Holds IS-IS domain NET (address) value.
     * 
     */
    @Export(name="net", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> net;

    /**
     * @return Holds IS-IS domain NET (address) value.
     * 
     */
    public Output<Optional<String>> net() {
        return Codegen.optional(this.net);
    }
    /**
     * IS-IS Domain passive-interface default level. - Choices: `l1`, `l2`, `l12`, `unknown` - Default value: `unknown`
     * 
     */
    @Export(name="passiveDefault", refs={String.class}, tree="[0]")
    private Output<String> passiveDefault;

    /**
     * @return IS-IS Domain passive-interface default level. - Choices: `l1`, `l2`, `l12`, `unknown` - Default value: `unknown`
     * 
     */
    public Output<String> passiveDefault() {
        return this.passiveDefault;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IsisVrf(String name) {
        this(name, IsisVrfArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IsisVrf(String name, IsisVrfArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IsisVrf(String name, IsisVrfArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/isisVrf:IsisVrf", name, args == null ? IsisVrfArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IsisVrf(String name, Output<String> id, @Nullable IsisVrfState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:index/isisVrf:IsisVrf", name, state, makeResourceOptions(options, id));
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
    public static IsisVrf get(String name, Output<String> id, @Nullable IsisVrfState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IsisVrf(name, id, state, options);
    }
}
