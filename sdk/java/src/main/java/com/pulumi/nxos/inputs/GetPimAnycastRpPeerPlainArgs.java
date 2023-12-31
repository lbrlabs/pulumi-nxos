// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPimAnycastRpPeerPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPimAnycastRpPeerPlainArgs Empty = new GetPimAnycastRpPeerPlainArgs();

    /**
     * Anycast RP address.
     * 
     */
    @Import(name="address", required=true)
    private String address;

    /**
     * @return Anycast RP address.
     * 
     */
    public String address() {
        return this.address;
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
     * RP set address.
     * 
     */
    @Import(name="rpSetAddress", required=true)
    private String rpSetAddress;

    /**
     * @return RP set address.
     * 
     */
    public String rpSetAddress() {
        return this.rpSetAddress;
    }

    /**
     * VRF name.
     * 
     */
    @Import(name="vrfName", required=true)
    private String vrfName;

    /**
     * @return VRF name.
     * 
     */
    public String vrfName() {
        return this.vrfName;
    }

    private GetPimAnycastRpPeerPlainArgs() {}

    private GetPimAnycastRpPeerPlainArgs(GetPimAnycastRpPeerPlainArgs $) {
        this.address = $.address;
        this.device = $.device;
        this.rpSetAddress = $.rpSetAddress;
        this.vrfName = $.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPimAnycastRpPeerPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPimAnycastRpPeerPlainArgs $;

        public Builder() {
            $ = new GetPimAnycastRpPeerPlainArgs();
        }

        public Builder(GetPimAnycastRpPeerPlainArgs defaults) {
            $ = new GetPimAnycastRpPeerPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address Anycast RP address.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            $.address = address;
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
         * @param rpSetAddress RP set address.
         * 
         * @return builder
         * 
         */
        public Builder rpSetAddress(String rpSetAddress) {
            $.rpSetAddress = rpSetAddress;
            return this;
        }

        /**
         * @param vrfName VRF name.
         * 
         * @return builder
         * 
         */
        public Builder vrfName(String vrfName) {
            $.vrfName = vrfName;
            return this;
        }

        public GetPimAnycastRpPeerPlainArgs build() {
            $.address = Objects.requireNonNull($.address, "expected parameter 'address' to be non-null");
            $.rpSetAddress = Objects.requireNonNull($.rpSetAddress, "expected parameter 'rpSetAddress' to be non-null");
            $.vrfName = Objects.requireNonNull($.vrfName, "expected parameter 'vrfName' to be non-null");
            return $;
        }
    }

}
