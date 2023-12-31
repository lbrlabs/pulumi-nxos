// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPimStaticRpGroupListArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPimStaticRpGroupListArgs Empty = new GetPimStaticRpGroupListArgs();

    /**
     * Group list address information.
     * 
     */
    @Import(name="address", required=true)
    private Output<String> address;

    /**
     * @return Group list address information.
     * 
     */
    public Output<String> address() {
        return this.address;
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
     * RP address.
     * 
     */
    @Import(name="rpAddress", required=true)
    private Output<String> rpAddress;

    /**
     * @return RP address.
     * 
     */
    public Output<String> rpAddress() {
        return this.rpAddress;
    }

    /**
     * VRF name.
     * 
     */
    @Import(name="vrfName", required=true)
    private Output<String> vrfName;

    /**
     * @return VRF name.
     * 
     */
    public Output<String> vrfName() {
        return this.vrfName;
    }

    private GetPimStaticRpGroupListArgs() {}

    private GetPimStaticRpGroupListArgs(GetPimStaticRpGroupListArgs $) {
        this.address = $.address;
        this.device = $.device;
        this.rpAddress = $.rpAddress;
        this.vrfName = $.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPimStaticRpGroupListArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPimStaticRpGroupListArgs $;

        public Builder() {
            $ = new GetPimStaticRpGroupListArgs();
        }

        public Builder(GetPimStaticRpGroupListArgs defaults) {
            $ = new GetPimStaticRpGroupListArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address Group list address information.
         * 
         * @return builder
         * 
         */
        public Builder address(Output<String> address) {
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
         * @param rpAddress RP address.
         * 
         * @return builder
         * 
         */
        public Builder rpAddress(Output<String> rpAddress) {
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
        public Builder vrfName(Output<String> vrfName) {
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

        public GetPimStaticRpGroupListArgs build() {
            $.address = Objects.requireNonNull($.address, "expected parameter 'address' to be non-null");
            $.rpAddress = Objects.requireNonNull($.rpAddress, "expected parameter 'rpAddress' to be non-null");
            $.vrfName = Objects.requireNonNull($.vrfName, "expected parameter 'vrfName' to be non-null");
            return $;
        }
    }

}
