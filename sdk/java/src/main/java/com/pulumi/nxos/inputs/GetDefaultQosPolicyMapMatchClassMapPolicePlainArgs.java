// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDefaultQosPolicyMapMatchClassMapPolicePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDefaultQosPolicyMapMatchClassMapPolicePlainArgs Empty = new GetDefaultQosPolicyMapMatchClassMapPolicePlainArgs();

    /**
     * Class map name.
     * 
     */
    @Import(name="classMapName", required=true)
    private String classMapName;

    /**
     * @return Class map name.
     * 
     */
    public String classMapName() {
        return this.classMapName;
    }

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

    private GetDefaultQosPolicyMapMatchClassMapPolicePlainArgs() {}

    private GetDefaultQosPolicyMapMatchClassMapPolicePlainArgs(GetDefaultQosPolicyMapMatchClassMapPolicePlainArgs $) {
        this.classMapName = $.classMapName;
        this.device = $.device;
        this.policyMapName = $.policyMapName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDefaultQosPolicyMapMatchClassMapPolicePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDefaultQosPolicyMapMatchClassMapPolicePlainArgs $;

        public Builder() {
            $ = new GetDefaultQosPolicyMapMatchClassMapPolicePlainArgs();
        }

        public Builder(GetDefaultQosPolicyMapMatchClassMapPolicePlainArgs defaults) {
            $ = new GetDefaultQosPolicyMapMatchClassMapPolicePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param classMapName Class map name.
         * 
         * @return builder
         * 
         */
        public Builder classMapName(String classMapName) {
            $.classMapName = classMapName;
            return this;
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
         * @param policyMapName Policy map name.
         * 
         * @return builder
         * 
         */
        public Builder policyMapName(String policyMapName) {
            $.policyMapName = policyMapName;
            return this;
        }

        public GetDefaultQosPolicyMapMatchClassMapPolicePlainArgs build() {
            $.classMapName = Objects.requireNonNull($.classMapName, "expected parameter 'classMapName' to be non-null");
            $.policyMapName = Objects.requireNonNull($.policyMapName, "expected parameter 'policyMapName' to be non-null");
            return $;
        }
    }

}