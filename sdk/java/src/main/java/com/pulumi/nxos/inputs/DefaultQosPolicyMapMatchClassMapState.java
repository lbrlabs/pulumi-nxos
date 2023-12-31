// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DefaultQosPolicyMapMatchClassMapState extends com.pulumi.resources.ResourceArgs {

    public static final DefaultQosPolicyMapMatchClassMapState Empty = new DefaultQosPolicyMapMatchClassMapState();

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
     * Class map name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Class map name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
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

    private DefaultQosPolicyMapMatchClassMapState() {}

    private DefaultQosPolicyMapMatchClassMapState(DefaultQosPolicyMapMatchClassMapState $) {
        this.device = $.device;
        this.name = $.name;
        this.policyMapName = $.policyMapName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DefaultQosPolicyMapMatchClassMapState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DefaultQosPolicyMapMatchClassMapState $;

        public Builder() {
            $ = new DefaultQosPolicyMapMatchClassMapState();
        }

        public Builder(DefaultQosPolicyMapMatchClassMapState defaults) {
            $ = new DefaultQosPolicyMapMatchClassMapState(Objects.requireNonNull(defaults));
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
         * @param name Class map name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Class map name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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

        public DefaultQosPolicyMapMatchClassMapState build() {
            return $;
        }
    }

}
