// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VrfAddressFamilyState extends com.pulumi.resources.ResourceArgs {

    public static final VrfAddressFamilyState Empty = new VrfAddressFamilyState();

    /**
     * Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
     * 
     */
    @Import(name="addressFamily")
    private @Nullable Output<String> addressFamily;

    /**
     * @return Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
     * 
     */
    public Optional<Output<String>> addressFamily() {
        return Optional.ofNullable(this.addressFamily);
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
     * VRF name.
     * 
     */
    @Import(name="vrf")
    private @Nullable Output<String> vrf;

    /**
     * @return VRF name.
     * 
     */
    public Optional<Output<String>> vrf() {
        return Optional.ofNullable(this.vrf);
    }

    private VrfAddressFamilyState() {}

    private VrfAddressFamilyState(VrfAddressFamilyState $) {
        this.addressFamily = $.addressFamily;
        this.device = $.device;
        this.vrf = $.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VrfAddressFamilyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VrfAddressFamilyState $;

        public Builder() {
            $ = new VrfAddressFamilyState();
        }

        public Builder(VrfAddressFamilyState defaults) {
            $ = new VrfAddressFamilyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressFamily Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(@Nullable Output<String> addressFamily) {
            $.addressFamily = addressFamily;
            return this;
        }

        /**
         * @param addressFamily Address family. - Choices: `ipv4-ucast`, `ipv6-ucast`
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(String addressFamily) {
            return addressFamily(Output.of(addressFamily));
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
         * @param vrf VRF name.
         * 
         * @return builder
         * 
         */
        public Builder vrf(@Nullable Output<String> vrf) {
            $.vrf = vrf;
            return this;
        }

        /**
         * @param vrf VRF name.
         * 
         * @return builder
         * 
         */
        public Builder vrf(String vrf) {
            return vrf(Output.of(vrf));
        }

        public VrfAddressFamilyState build() {
            return $;
        }
    }

}
