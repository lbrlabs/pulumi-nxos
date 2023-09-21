// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QueuingQosPolicyMapMatchClassMapPriorityState extends com.pulumi.resources.ResourceArgs {

    public static final QueuingQosPolicyMapMatchClassMapPriorityState Empty = new QueuingQosPolicyMapMatchClassMapPriorityState();

    /**
     * Class map name.
     * 
     */
    @Import(name="classMapName")
    private @Nullable Output<String> classMapName;

    /**
     * @return Class map name.
     * 
     */
    public Optional<Output<String>> classMapName() {
        return Optional.ofNullable(this.classMapName);
    }

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
     * Priority level. - Range: `1`-`8`
     * 
     */
    @Import(name="level")
    private @Nullable Output<Integer> level;

    /**
     * @return Priority level. - Range: `1`-`8`
     * 
     */
    public Optional<Output<Integer>> level() {
        return Optional.ofNullable(this.level);
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

    private QueuingQosPolicyMapMatchClassMapPriorityState() {}

    private QueuingQosPolicyMapMatchClassMapPriorityState(QueuingQosPolicyMapMatchClassMapPriorityState $) {
        this.classMapName = $.classMapName;
        this.device = $.device;
        this.level = $.level;
        this.policyMapName = $.policyMapName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QueuingQosPolicyMapMatchClassMapPriorityState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QueuingQosPolicyMapMatchClassMapPriorityState $;

        public Builder() {
            $ = new QueuingQosPolicyMapMatchClassMapPriorityState();
        }

        public Builder(QueuingQosPolicyMapMatchClassMapPriorityState defaults) {
            $ = new QueuingQosPolicyMapMatchClassMapPriorityState(Objects.requireNonNull(defaults));
        }

        /**
         * @param classMapName Class map name.
         * 
         * @return builder
         * 
         */
        public Builder classMapName(@Nullable Output<String> classMapName) {
            $.classMapName = classMapName;
            return this;
        }

        /**
         * @param classMapName Class map name.
         * 
         * @return builder
         * 
         */
        public Builder classMapName(String classMapName) {
            return classMapName(Output.of(classMapName));
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
         * @param level Priority level. - Range: `1`-`8`
         * 
         * @return builder
         * 
         */
        public Builder level(@Nullable Output<Integer> level) {
            $.level = level;
            return this;
        }

        /**
         * @param level Priority level. - Range: `1`-`8`
         * 
         * @return builder
         * 
         */
        public Builder level(Integer level) {
            return level(Output.of(level));
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

        public QueuingQosPolicyMapMatchClassMapPriorityState build() {
            return $;
        }
    }

}
