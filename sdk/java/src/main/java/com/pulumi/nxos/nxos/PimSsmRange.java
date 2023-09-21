// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.nxos.PimSsmRangeArgs;
import com.pulumi.nxos.nxos.inputs.PimSsmRangeState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="nxos:nxos/pimSsmRange:PimSsmRange")
public class PimSsmRange extends com.pulumi.resources.CustomResource {
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
     * Group list 1. - Default value: `0.0.0.0`
     * 
     */
    @Export(name="groupList1", refs={String.class}, tree="[0]")
    private Output<String> groupList1;

    /**
     * @return Group list 1. - Default value: `0.0.0.0`
     * 
     */
    public Output<String> groupList1() {
        return this.groupList1;
    }
    /**
     * Group list 2. - Default value: `0.0.0.0`
     * 
     */
    @Export(name="groupList2", refs={String.class}, tree="[0]")
    private Output<String> groupList2;

    /**
     * @return Group list 2. - Default value: `0.0.0.0`
     * 
     */
    public Output<String> groupList2() {
        return this.groupList2;
    }
    /**
     * Group list 3. - Default value: `0.0.0.0`
     * 
     */
    @Export(name="groupList3", refs={String.class}, tree="[0]")
    private Output<String> groupList3;

    /**
     * @return Group list 3. - Default value: `0.0.0.0`
     * 
     */
    public Output<String> groupList3() {
        return this.groupList3;
    }
    /**
     * Group list 4. - Default value: `0.0.0.0`
     * 
     */
    @Export(name="groupList4", refs={String.class}, tree="[0]")
    private Output<String> groupList4;

    /**
     * @return Group list 4. - Default value: `0.0.0.0`
     * 
     */
    public Output<String> groupList4() {
        return this.groupList4;
    }
    /**
     * Prefix list name.
     * 
     */
    @Export(name="prefixList", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> prefixList;

    /**
     * @return Prefix list name.
     * 
     */
    public Output<Optional<String>> prefixList() {
        return Codegen.optional(this.prefixList);
    }
    /**
     * Route map name.
     * 
     */
    @Export(name="routeMap", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> routeMap;

    /**
     * @return Route map name.
     * 
     */
    public Output<Optional<String>> routeMap() {
        return Codegen.optional(this.routeMap);
    }
    /**
     * Exclude standard SSM range (232.0.0.0/8). - Default value: `false`
     * 
     */
    @Export(name="ssmNone", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> ssmNone;

    /**
     * @return Exclude standard SSM range (232.0.0.0/8). - Default value: `false`
     * 
     */
    public Output<Boolean> ssmNone() {
        return this.ssmNone;
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
    public PimSsmRange(String name) {
        this(name, PimSsmRangeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PimSsmRange(String name, PimSsmRangeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PimSsmRange(String name, PimSsmRangeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/pimSsmRange:PimSsmRange", name, args == null ? PimSsmRangeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PimSsmRange(String name, Output<String> id, @Nullable PimSsmRangeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/pimSsmRange:PimSsmRange", name, state, makeResourceOptions(options, id));
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
    public static PimSsmRange get(String name, Output<String> id, @Nullable PimSsmRangeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PimSsmRange(name, id, state, options);
    }
}