// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BgpPeerAddressFamilyState extends com.pulumi.resources.ResourceArgs {

    public static final BgpPeerAddressFamilyState Empty = new BgpPeerAddressFamilyState();

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
     * Peer address-family control. Choices: `rr-client`, `nh-self`, `dis-peer-as-check`, `allow-self-as`, `default-originate`,
     * `advertisement-interval`, `suppress-inactive`, `nh-self-all`. Can be an empty string. Allowed formats: - Single value.
     * Example: `nh-self` - Multiple values (comma-separated). Example:
     * `dis-peer-as-check,nh-self,rr-client,suppress-inactive`. In this case values must be in alphabetical order.
     * 
     */
    @Import(name="control")
    private @Nullable Output<String> control;

    /**
     * @return Peer address-family control. Choices: `rr-client`, `nh-self`, `dis-peer-as-check`, `allow-self-as`, `default-originate`,
     * `advertisement-interval`, `suppress-inactive`, `nh-self-all`. Can be an empty string. Allowed formats: - Single value.
     * Example: `nh-self` - Multiple values (comma-separated). Example:
     * `dis-peer-as-check,nh-self,rr-client,suppress-inactive`. In this case values must be in alphabetical order.
     * 
     */
    public Optional<Output<String>> control() {
        return Optional.ofNullable(this.control);
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
     * Send-community extended. - Choices: `enabled`, `disabled` - Default value: `disabled`
     * 
     */
    @Import(name="sendCommunityExtended")
    private @Nullable Output<String> sendCommunityExtended;

    /**
     * @return Send-community extended. - Choices: `enabled`, `disabled` - Default value: `disabled`
     * 
     */
    public Optional<Output<String>> sendCommunityExtended() {
        return Optional.ofNullable(this.sendCommunityExtended);
    }

    /**
     * Send-community standard. - Choices: `enabled`, `disabled` - Default value: `disabled`
     * 
     */
    @Import(name="sendCommunityStandard")
    private @Nullable Output<String> sendCommunityStandard;

    /**
     * @return Send-community standard. - Choices: `enabled`, `disabled` - Default value: `disabled`
     * 
     */
    public Optional<Output<String>> sendCommunityStandard() {
        return Optional.ofNullable(this.sendCommunityStandard);
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

    private BgpPeerAddressFamilyState() {}

    private BgpPeerAddressFamilyState(BgpPeerAddressFamilyState $) {
        this.address = $.address;
        this.addressFamily = $.addressFamily;
        this.asn = $.asn;
        this.control = $.control;
        this.device = $.device;
        this.sendCommunityExtended = $.sendCommunityExtended;
        this.sendCommunityStandard = $.sendCommunityStandard;
        this.vrf = $.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BgpPeerAddressFamilyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BgpPeerAddressFamilyState $;

        public Builder() {
            $ = new BgpPeerAddressFamilyState();
        }

        public Builder(BgpPeerAddressFamilyState defaults) {
            $ = new BgpPeerAddressFamilyState(Objects.requireNonNull(defaults));
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
         * @param control Peer address-family control. Choices: `rr-client`, `nh-self`, `dis-peer-as-check`, `allow-self-as`, `default-originate`,
         * `advertisement-interval`, `suppress-inactive`, `nh-self-all`. Can be an empty string. Allowed formats: - Single value.
         * Example: `nh-self` - Multiple values (comma-separated). Example:
         * `dis-peer-as-check,nh-self,rr-client,suppress-inactive`. In this case values must be in alphabetical order.
         * 
         * @return builder
         * 
         */
        public Builder control(@Nullable Output<String> control) {
            $.control = control;
            return this;
        }

        /**
         * @param control Peer address-family control. Choices: `rr-client`, `nh-self`, `dis-peer-as-check`, `allow-self-as`, `default-originate`,
         * `advertisement-interval`, `suppress-inactive`, `nh-self-all`. Can be an empty string. Allowed formats: - Single value.
         * Example: `nh-self` - Multiple values (comma-separated). Example:
         * `dis-peer-as-check,nh-self,rr-client,suppress-inactive`. In this case values must be in alphabetical order.
         * 
         * @return builder
         * 
         */
        public Builder control(String control) {
            return control(Output.of(control));
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
         * @param sendCommunityExtended Send-community extended. - Choices: `enabled`, `disabled` - Default value: `disabled`
         * 
         * @return builder
         * 
         */
        public Builder sendCommunityExtended(@Nullable Output<String> sendCommunityExtended) {
            $.sendCommunityExtended = sendCommunityExtended;
            return this;
        }

        /**
         * @param sendCommunityExtended Send-community extended. - Choices: `enabled`, `disabled` - Default value: `disabled`
         * 
         * @return builder
         * 
         */
        public Builder sendCommunityExtended(String sendCommunityExtended) {
            return sendCommunityExtended(Output.of(sendCommunityExtended));
        }

        /**
         * @param sendCommunityStandard Send-community standard. - Choices: `enabled`, `disabled` - Default value: `disabled`
         * 
         * @return builder
         * 
         */
        public Builder sendCommunityStandard(@Nullable Output<String> sendCommunityStandard) {
            $.sendCommunityStandard = sendCommunityStandard;
            return this;
        }

        /**
         * @param sendCommunityStandard Send-community standard. - Choices: `enabled`, `disabled` - Default value: `disabled`
         * 
         * @return builder
         * 
         */
        public Builder sendCommunityStandard(String sendCommunityStandard) {
            return sendCommunityStandard(Output.of(sendCommunityStandard));
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

        public BgpPeerAddressFamilyState build() {
            return $;
        }
    }

}
