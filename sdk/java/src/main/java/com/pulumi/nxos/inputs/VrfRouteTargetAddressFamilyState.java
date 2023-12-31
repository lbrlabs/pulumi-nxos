// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VrfRouteTargetAddressFamilyState extends com.pulumi.resources.ResourceArgs {

    public static final VrfRouteTargetAddressFamilyState Empty = new VrfRouteTargetAddressFamilyState();

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
     * Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
     * 
     */
    @Import(name="routeTargetAddressFamily")
    private @Nullable Output<String> routeTargetAddressFamily;

    /**
     * @return Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
     * 
     */
    public Optional<Output<String>> routeTargetAddressFamily() {
        return Optional.ofNullable(this.routeTargetAddressFamily);
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

    private VrfRouteTargetAddressFamilyState() {}

    private VrfRouteTargetAddressFamilyState(VrfRouteTargetAddressFamilyState $) {
        this.addressFamily = $.addressFamily;
        this.device = $.device;
        this.routeTargetAddressFamily = $.routeTargetAddressFamily;
        this.vrf = $.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VrfRouteTargetAddressFamilyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VrfRouteTargetAddressFamilyState $;

        public Builder() {
            $ = new VrfRouteTargetAddressFamilyState();
        }

        public Builder(VrfRouteTargetAddressFamilyState defaults) {
            $ = new VrfRouteTargetAddressFamilyState(Objects.requireNonNull(defaults));
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
         * @param routeTargetAddressFamily Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
         * 
         * @return builder
         * 
         */
        public Builder routeTargetAddressFamily(@Nullable Output<String> routeTargetAddressFamily) {
            $.routeTargetAddressFamily = routeTargetAddressFamily;
            return this;
        }

        /**
         * @param routeTargetAddressFamily Route Target Address Family. - Choices: `ipv4-ucast`, `ipv6-ucast`, `l2vpn-evpn`
         * 
         * @return builder
         * 
         */
        public Builder routeTargetAddressFamily(String routeTargetAddressFamily) {
            return routeTargetAddressFamily(Output.of(routeTargetAddressFamily));
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

        public VrfRouteTargetAddressFamilyState build() {
            return $;
        }
    }

}
