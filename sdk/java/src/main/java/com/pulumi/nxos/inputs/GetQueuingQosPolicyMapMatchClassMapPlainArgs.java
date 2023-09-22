// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetQueuingQosPolicyMapMatchClassMapPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetQueuingQosPolicyMapMatchClassMapPlainArgs Empty = new GetQueuingQosPolicyMapMatchClassMapPlainArgs();

    /**
     * A device name from the provider configuration.
     * 
     */
    @Import(name="device")
    private @Nullable String device;

    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }

    /**
     * Class map name.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Class map name.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * Policy map name.
     * 
     */
    @Import(name="policyMapName", required=true)
    private String policyMapName;

    /**
     * @return Policy map name.
     * 
     */
    public String policyMapName() {
        return this.policyMapName;
    }

    private GetQueuingQosPolicyMapMatchClassMapPlainArgs() {}

    private GetQueuingQosPolicyMapMatchClassMapPlainArgs(GetQueuingQosPolicyMapMatchClassMapPlainArgs $) {
        this.device = $.device;
        this.name = $.name;
        this.policyMapName = $.policyMapName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetQueuingQosPolicyMapMatchClassMapPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetQueuingQosPolicyMapMatchClassMapPlainArgs $;

        public Builder() {
            $ = new GetQueuingQosPolicyMapMatchClassMapPlainArgs();
        }

        public Builder(GetQueuingQosPolicyMapMatchClassMapPlainArgs defaults) {
            $ = new GetQueuingQosPolicyMapMatchClassMapPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param device A device name from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder device(@Nullable String device) {
            $.device = device;
            return this;
        }

        /**
         * @param name Class map name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param policyMapName Policy map name.
         * 
         * @return builder
         * 
         */
        public Builder policyMapName(String policyMapName) {
            $.policyMapName = policyMapName;
            return this;
        }

        public GetQueuingQosPolicyMapMatchClassMapPlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.policyMapName = Objects.requireNonNull($.policyMapName, "expected parameter 'policyMapName' to be non-null");
            return $;
        }
    }

}