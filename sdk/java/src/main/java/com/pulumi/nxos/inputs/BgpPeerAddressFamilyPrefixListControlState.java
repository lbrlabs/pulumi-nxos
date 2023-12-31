// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BgpPeerAddressFamilyPrefixListControlState extends com.pulumi.resources.ResourceArgs {

    public static final BgpPeerAddressFamilyPrefixListControlState Empty = new BgpPeerAddressFamilyPrefixListControlState();

    /**
     * Peer address.
     * 
     */
    @Import(name="address")
    private @Nullable Output<String> address;

    /**
     * @return Peer address.
     * 
     */
    public Optional<Output<String>> address() {
        return Optional.ofNullable(this.address);
    }

    /**
     * Address Family. - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate` - Default
     * value: `ipv4-ucast`
     * 
     */
    @Import(name="addressFamily")
    private @Nullable Output<String> addressFamily;

    /**
     * @return Address Family. - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate` - Default
     * value: `ipv4-ucast`
     * 
     */
    public Optional<Output<String>> addressFamily() {
        return Optional.ofNullable(this.addressFamily);
    }

    /**
     * Autonomous system number.
     * 
     */
    @Import(name="asn")
    private @Nullable Output<String> asn;

    /**
     * @return Autonomous system number.
     * 
     */
    public Optional<Output<String>> asn() {
        return Optional.ofNullable(this.asn);
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
     * Route Control direction. - Choices: `in`, `out` - Default value: `in`
     * 
     */
    @Import(name="direction")
    private @Nullable Output<String> direction;

    /**
     * @return Route Control direction. - Choices: `in`, `out` - Default value: `in`
     * 
     */
    public Optional<Output<String>> direction() {
        return Optional.ofNullable(this.direction);
    }

    /**
     * Route Control Prefix-List name.
     * 
     */
    @Import(name="list")
    private @Nullable Output<String> list;

    /**
     * @return Route Control Prefix-List name.
     * 
     */
    public Optional<Output<String>> list() {
        return Optional.ofNullable(this.list);
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

    private BgpPeerAddressFamilyPrefixListControlState() {}

    private BgpPeerAddressFamilyPrefixListControlState(BgpPeerAddressFamilyPrefixListControlState $) {
        this.address = $.address;
        this.addressFamily = $.addressFamily;
        this.asn = $.asn;
        this.device = $.device;
        this.direction = $.direction;
        this.list = $.list;
        this.vrf = $.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BgpPeerAddressFamilyPrefixListControlState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BgpPeerAddressFamilyPrefixListControlState $;

        public Builder() {
            $ = new BgpPeerAddressFamilyPrefixListControlState();
        }

        public Builder(BgpPeerAddressFamilyPrefixListControlState defaults) {
            $ = new BgpPeerAddressFamilyPrefixListControlState(Objects.requireNonNull(defaults));
        }

        /**
         * @param address Peer address.
         * 
         * @return builder
         * 
         */
        public Builder address(@Nullable Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address Peer address.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
        }

        /**
         * @param addressFamily Address Family. - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate` - Default
         * value: `ipv4-ucast`
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(@Nullable Output<String> addressFamily) {
            $.addressFamily = addressFamily;
            return this;
        }

        /**
         * @param addressFamily Address Family. - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate` - Default
         * value: `ipv4-ucast`
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(String addressFamily) {
            return addressFamily(Output.of(addressFamily));
        }

        /**
         * @param asn Autonomous system number.
         * 
         * @return builder
         * 
         */
        public Builder asn(@Nullable Output<String> asn) {
            $.asn = asn;
            return this;
        }

        /**
         * @param asn Autonomous system number.
         * 
         * @return builder
         * 
         */
        public Builder asn(String asn) {
            return asn(Output.of(asn));
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
         * @param direction Route Control direction. - Choices: `in`, `out` - Default value: `in`
         * 
         * @return builder
         * 
         */
        public Builder direction(@Nullable Output<String> direction) {
            $.direction = direction;
            return this;
        }

        /**
         * @param direction Route Control direction. - Choices: `in`, `out` - Default value: `in`
         * 
         * @return builder
         * 
         */
        public Builder direction(String direction) {
            return direction(Output.of(direction));
        }

        /**
         * @param list Route Control Prefix-List name.
         * 
         * @return builder
         * 
         */
        public Builder list(@Nullable Output<String> list) {
            $.list = list;
            return this;
        }

        /**
         * @param list Route Control Prefix-List name.
         * 
         * @return builder
         * 
         */
        public Builder list(String list) {
            return list(Output.of(list));
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

        public BgpPeerAddressFamilyPrefixListControlState build() {
            return $;
        }
    }

}
