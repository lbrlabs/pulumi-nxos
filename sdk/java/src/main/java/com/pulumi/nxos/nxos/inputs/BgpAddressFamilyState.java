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


public final class BgpAddressFamilyState extends com.pulumi.resources.ResourceArgs {

    public static final BgpAddressFamilyState Empty = new BgpAddressFamilyState();

    /**
     * Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
     * `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
     * `ipv4-mdt` - Default value: `ipv4-ucast`
     * 
     */
    @Import(name="addressFamily")
    private @Nullable Output<String> addressFamily;

    /**
     * @return Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
     * `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
     * `ipv4-mdt` - Default value: `ipv4-ucast`
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
     * The next-hop address tracking delay timer for critical next-hop reachability routes. - Range: `1`-`4294967295` - Default
     * value: `3000`
     * 
     */
    @Import(name="criticalNexthopTimeout")
    private @Nullable Output<Integer> criticalNexthopTimeout;

    /**
     * @return The next-hop address tracking delay timer for critical next-hop reachability routes. - Range: `1`-`4294967295` - Default
     * value: `3000`
     * 
     */
    public Optional<Output<Integer>> criticalNexthopTimeout() {
        return Optional.ofNullable(this.criticalNexthopTimeout);
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
     * The next-hop address tracking delay timer for non-critical next-hop reachability routes. - Range: `1`-`4294967295` -
     * Default value: `10000`
     * 
     */
    @Import(name="nonCriticalNexthopTimeout")
    private @Nullable Output<Integer> nonCriticalNexthopTimeout;

    /**
     * @return The next-hop address tracking delay timer for non-critical next-hop reachability routes. - Range: `1`-`4294967295` -
     * Default value: `10000`
     * 
     */
    public Optional<Output<Integer>> nonCriticalNexthopTimeout() {
        return Optional.ofNullable(this.nonCriticalNexthopTimeout);
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

    private BgpAddressFamilyState() {}

    private BgpAddressFamilyState(BgpAddressFamilyState $) {
        this.addressFamily = $.addressFamily;
        this.asn = $.asn;
        this.criticalNexthopTimeout = $.criticalNexthopTimeout;
        this.device = $.device;
        this.nonCriticalNexthopTimeout = $.nonCriticalNexthopTimeout;
        this.vrf = $.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BgpAddressFamilyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BgpAddressFamilyState $;

        public Builder() {
            $ = new BgpAddressFamilyState();
        }

        public Builder(BgpAddressFamilyState defaults) {
            $ = new BgpAddressFamilyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressFamily Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
         * `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
         * `ipv4-mdt` - Default value: `ipv4-ucast`
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(@Nullable Output<String> addressFamily) {
            $.addressFamily = addressFamily;
            return this;
        }

        /**
         * @param addressFamily Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
         * `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
         * `ipv4-mdt` - Default value: `ipv4-ucast`
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
         * @param criticalNexthopTimeout The next-hop address tracking delay timer for critical next-hop reachability routes. - Range: `1`-`4294967295` - Default
         * value: `3000`
         * 
         * @return builder
         * 
         */
        public Builder criticalNexthopTimeout(@Nullable Output<Integer> criticalNexthopTimeout) {
            $.criticalNexthopTimeout = criticalNexthopTimeout;
            return this;
        }

        /**
         * @param criticalNexthopTimeout The next-hop address tracking delay timer for critical next-hop reachability routes. - Range: `1`-`4294967295` - Default
         * value: `3000`
         * 
         * @return builder
         * 
         */
        public Builder criticalNexthopTimeout(Integer criticalNexthopTimeout) {
            return criticalNexthopTimeout(Output.of(criticalNexthopTimeout));
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
         * @param nonCriticalNexthopTimeout The next-hop address tracking delay timer for non-critical next-hop reachability routes. - Range: `1`-`4294967295` -
         * Default value: `10000`
         * 
         * @return builder
         * 
         */
        public Builder nonCriticalNexthopTimeout(@Nullable Output<Integer> nonCriticalNexthopTimeout) {
            $.nonCriticalNexthopTimeout = nonCriticalNexthopTimeout;
            return this;
        }

        /**
         * @param nonCriticalNexthopTimeout The next-hop address tracking delay timer for non-critical next-hop reachability routes. - Range: `1`-`4294967295` -
         * Default value: `10000`
         * 
         * @return builder
         * 
         */
        public Builder nonCriticalNexthopTimeout(Integer nonCriticalNexthopTimeout) {
            return nonCriticalNexthopTimeout(Output.of(nonCriticalNexthopTimeout));
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

        public BgpAddressFamilyState build() {
            return $;
        }
    }

}