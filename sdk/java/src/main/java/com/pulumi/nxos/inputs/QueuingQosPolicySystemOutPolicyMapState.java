// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QueuingQosPolicySystemOutPolicyMapState extends com.pulumi.resources.ResourceArgs {

    public static final QueuingQosPolicySystemOutPolicyMapState Empty = new QueuingQosPolicySystemOutPolicyMapState();

    /**
     * A device name from the provider configuration.
     * 
     */
    @Import(name="device")
    private @Nullable Output<String> device;

    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    /**
     * Policy map name.
     * 
     */
    @Import(name="policyMapName")
    private @Nullable Output<String> policyMapName;

    /**
     * @return Policy map name.
     * 
     */
    public Optional<Output<String>> policyMapName() {
        return Optional.ofNullable(this.policyMapName);
    }

    private QueuingQosPolicySystemOutPolicyMapState() {}

    private QueuingQosPolicySystemOutPolicyMapState(QueuingQosPolicySystemOutPolicyMapState $) {
        this.device = $.device;
        this.policyMapName = $.policyMapName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QueuingQosPolicySystemOutPolicyMapState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QueuingQosPolicySystemOutPolicyMapState $;

        public Builder() {
            $ = new QueuingQosPolicySystemOutPolicyMapState();
        }

        public Builder(QueuingQosPolicySystemOutPolicyMapState defaults) {
            $ = new QueuingQosPolicySystemOutPolicyMapState(Objects.requireNonNull(defaults));
        }

        /**
         * @param device A device name from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        /**
         * @param device A device name from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder device(String device) {
            return device(Output.of(device));
        }

        /**
         * @param policyMapName Policy map name.
         * 
         * @return builder
         * 
         */
        public Builder policyMapName(@Nullable Output<String> policyMapName) {
            $.policyMapName = policyMapName;
            return this;
        }

        /**
         * @param policyMapName Policy map name.
         * 
         * @return builder
         * 
         */
        public Builder policyMapName(String policyMapName) {
            return policyMapName(Output.of(policyMapName));
        }

        public QueuingQosPolicySystemOutPolicyMapState build() {
            return $;
        }
    }

}
