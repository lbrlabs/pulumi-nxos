// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PimStaticRpGroupListState extends com.pulumi.resources.ResourceArgs {

    public static final PimStaticRpGroupListState Empty = new PimStaticRpGroupListState();

    /**
     * Group list address information.
     * 
     */
    @Import(name="address")
    private @Nullable Output<String> address;

    /**
     * @return Group list address information.
     * 
     */
    public Optional<Output<String>> address() {
        return Optional.ofNullable(this.address);
    }

    /**
     * Flag to treat Group Ranges as BiDir. - Default value: `false`
     * 
     */
    @Import(name="bidir")
    private @Nullable Output<Boolean> bidir;

    /**
     * @return Flag to treat Group Ranges as BiDir. - Default value: `false`
     * 
     */
    public Optional<Output<Boolean>> bidir() {
        return Optional.ofNullable(this.bidir);
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
     * Flag to override RP preference to use Static over Dynamic RP. - Default value: `false`
     * 
     */
    @Import(name="override")
    private @Nullable Output<Boolean> override;

    /**
     * @return Flag to override RP preference to use Static over Dynamic RP. - Default value: `false`
     * 
     */
    public Optional<Output<Boolean>> override() {
        return Optional.ofNullable(this.override);
    }

    /**
     * RP address.
     * 
     */
    @Import(name="rpAddress")
    private @Nullable Output<String> rpAddress;

    /**
     * @return RP address.
     * 
     */
    public Optional<Output<String>> rpAddress() {
        return Optional.ofNullable(this.rpAddress);
    }

    /**
     * VRF name.
     * 
     */
    @Import(name="vrfName")
    private @Nullable Output<String> vrfName;

    /**
     * @return VRF name.
     * 
     */
    public Optional<Output<String>> vrfName() {
        return Optional.ofNullable(this.vrfName);
    }

    private PimStaticRpGroupListState() {}

    private PimStaticRpGroupListState(PimStaticRpGroupListState $) {
        this.address = $.address;
        this.bidir = $.bidir;
        this.device = $.device;
        this.override = $.override;
        this.rpAddress = $.rpAddress;
        this.vrfName = $.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PimStaticRpGroupListState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PimStaticRpGroupListState $;

        public Builder() {
            $ = new PimStaticRpGroupListState();
        }

        public Builder(PimStaticRpGroupListState defaults) {
            $ = new PimStaticRpGroupListState(Objects.requireNonNull(defaults));
        }

        /**
         * @param address Group list address information.
         * 
         * @return builder
         * 
         */
        public Builder address(@Nullable Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address Group list address information.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param bidir Flag to treat Group Ranges as BiDir. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder bidir(@Nullable Output<Boolean> bidir) {
            $.bidir = bidir;
            return this;
        }

        /**
         * @param bidir Flag to treat Group Ranges as BiDir. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder bidir(Boolean bidir) {
            return bidir(Output.of(bidir));
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
         * @param override Flag to override RP preference to use Static over Dynamic RP. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder override(@Nullable Output<Boolean> override) {
            $.override = override;
            return this;
        }

        /**
         * @param override Flag to override RP preference to use Static over Dynamic RP. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder override(Boolean override) {
            return override(Output.of(override));
        }

        /**
         * @param rpAddress RP address.
         * 
         * @return builder
         * 
         */
        public Builder rpAddress(@Nullable Output<String> rpAddress) {
            $.rpAddress = rpAddress;
            return this;
        }

        /**
         * @param rpAddress RP address.
         * 
         * @return builder
         * 
         */
        public Builder rpAddress(String rpAddress) {
            return rpAddress(Output.of(rpAddress));
        }

        /**
         * @param vrfName VRF name.
         * 
         * @return builder
         * 
         */
        public Builder vrfName(@Nullable Output<String> vrfName) {
            $.vrfName = vrfName;
            return this;
        }

        /**
         * @param vrfName VRF name.
         * 
         * @return builder
         * 
         */
        public Builder vrfName(String vrfName) {
            return vrfName(Output.of(vrfName));
        }

        public PimStaticRpGroupListState build() {
            return $;
        }
    }

}