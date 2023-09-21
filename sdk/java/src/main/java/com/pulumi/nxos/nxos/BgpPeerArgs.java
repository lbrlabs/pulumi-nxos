// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BgpPeerArgs extends com.pulumi.resources.ResourceArgs {

    public static final BgpPeerArgs Empty = new BgpPeerArgs();

    /**
     * Peer address.
     * 
     */
    @Import(name="address", required=true)
    private Output<String> address;

    /**
     * @return Peer address.
     * 
     */
    public Output<String> address() {
        return this.address;
    }

    /**
     * Autonomous system number.
     * 
     */
    @Import(name="asn", required=true)
    private Output<String> asn;

    /**
     * @return Autonomous system number.
     * 
     */
    public Output<String> asn() {
        return this.asn;
    }

    /**
     * Peer description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Peer description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
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
     * BGP Hold Timer in seconds. The value must be greater than the keepalive timer - Range: `3`-`3600` - Default value: `180`
     * 
     */
    @Import(name="holdTime")
    private @Nullable Output<Integer> holdTime;

    /**
     * @return BGP Hold Timer in seconds. The value must be greater than the keepalive timer - Range: `3`-`3600` - Default value: `180`
     * 
     */
    public Optional<Output<Integer>> holdTime() {
        return Optional.ofNullable(this.holdTime);
    }

    /**
     * BGP Keepalive Timer in seconds - Range: `0`-`3600` - Default value: `60`
     * 
     */
    @Import(name="keepalive")
    private @Nullable Output<Integer> keepalive;

    /**
     * @return BGP Keepalive Timer in seconds - Range: `0`-`3600` - Default value: `60`
     * 
     */
    public Optional<Output<Integer>> keepalive() {
        return Optional.ofNullable(this.keepalive);
    }

    /**
     * Peer template name.
     * 
     */
    @Import(name="peerTemplate")
    private @Nullable Output<String> peerTemplate;

    /**
     * @return Peer template name.
     * 
     */
    public Optional<Output<String>> peerTemplate() {
        return Optional.ofNullable(this.peerTemplate);
    }

    /**
     * Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
     * `fabric-internal`
     * 
     */
    @Import(name="peerType")
    private @Nullable Output<String> peerType;

    /**
     * @return Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
     * `fabric-internal`
     * 
     */
    public Optional<Output<String>> peerType() {
        return Optional.ofNullable(this.peerType);
    }

    /**
     * Peer autonomous system number.
     * 
     */
    @Import(name="remoteAsn")
    private @Nullable Output<String> remoteAsn;

    /**
     * @return Peer autonomous system number.
     * 
     */
    public Optional<Output<String>> remoteAsn() {
        return Optional.ofNullable(this.remoteAsn);
    }

    /**
     * Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
     * 
     */
    @Import(name="sourceInterface")
    private @Nullable Output<String> sourceInterface;

    /**
     * @return Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
     * 
     */
    public Optional<Output<String>> sourceInterface() {
        return Optional.ofNullable(this.sourceInterface);
    }

    /**
     * VRF name.
     * 
     */
    @Import(name="vrf", required=true)
    private Output<String> vrf;

    /**
     * @return VRF name.
     * 
     */
    public Output<String> vrf() {
        return this.vrf;
    }

    private BgpPeerArgs() {}

    private BgpPeerArgs(BgpPeerArgs $) {
        this.address = $.address;
        this.asn = $.asn;
        this.description = $.description;
        this.device = $.device;
        this.holdTime = $.holdTime;
        this.keepalive = $.keepalive;
        this.peerTemplate = $.peerTemplate;
        this.peerType = $.peerType;
        this.remoteAsn = $.remoteAsn;
        this.sourceInterface = $.sourceInterface;
        this.vrf = $.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BgpPeerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BgpPeerArgs $;

        public Builder() {
            $ = new BgpPeerArgs();
        }

        public Builder(BgpPeerArgs defaults) {
            $ = new BgpPeerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address Peer address.
         * 
         * @return builder
         * 
         */
        public Builder address(Output<String> address) {
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
         * @param asn Autonomous system number.
         * 
         * @return builder
         * 
         */
        public Builder asn(Output<String> asn) {
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
         * @param description Peer description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Peer description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
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
         * @param holdTime BGP Hold Timer in seconds. The value must be greater than the keepalive timer - Range: `3`-`3600` - Default value: `180`
         * 
         * @return builder
         * 
         */
        public Builder holdTime(@Nullable Output<Integer> holdTime) {
            $.holdTime = holdTime;
            return this;
        }

        /**
         * @param holdTime BGP Hold Timer in seconds. The value must be greater than the keepalive timer - Range: `3`-`3600` - Default value: `180`
         * 
         * @return builder
         * 
         */
        public Builder holdTime(Integer holdTime) {
            return holdTime(Output.of(holdTime));
        }

        /**
         * @param keepalive BGP Keepalive Timer in seconds - Range: `0`-`3600` - Default value: `60`
         * 
         * @return builder
         * 
         */
        public Builder keepalive(@Nullable Output<Integer> keepalive) {
            $.keepalive = keepalive;
            return this;
        }

        /**
         * @param keepalive BGP Keepalive Timer in seconds - Range: `0`-`3600` - Default value: `60`
         * 
         * @return builder
         * 
         */
        public Builder keepalive(Integer keepalive) {
            return keepalive(Output.of(keepalive));
        }

        /**
         * @param peerTemplate Peer template name.
         * 
         * @return builder
         * 
         */
        public Builder peerTemplate(@Nullable Output<String> peerTemplate) {
            $.peerTemplate = peerTemplate;
            return this;
        }

        /**
         * @param peerTemplate Peer template name.
         * 
         * @return builder
         * 
         */
        public Builder peerTemplate(String peerTemplate) {
            return peerTemplate(Output.of(peerTemplate));
        }

        /**
         * @param peerType Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
         * `fabric-internal`
         * 
         * @return builder
         * 
         */
        public Builder peerType(@Nullable Output<String> peerType) {
            $.peerType = peerType;
            return this;
        }

        /**
         * @param peerType Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
         * `fabric-internal`
         * 
         * @return builder
         * 
         */
        public Builder peerType(String peerType) {
            return peerType(Output.of(peerType));
        }

        /**
         * @param remoteAsn Peer autonomous system number.
         * 
         * @return builder
         * 
         */
        public Builder remoteAsn(@Nullable Output<String> remoteAsn) {
            $.remoteAsn = remoteAsn;
            return this;
        }

        /**
         * @param remoteAsn Peer autonomous system number.
         * 
         * @return builder
         * 
         */
        public Builder remoteAsn(String remoteAsn) {
            return remoteAsn(Output.of(remoteAsn));
        }

        /**
         * @param sourceInterface Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
         * 
         * @return builder
         * 
         */
        public Builder sourceInterface(@Nullable Output<String> sourceInterface) {
            $.sourceInterface = sourceInterface;
            return this;
        }

        /**
         * @param sourceInterface Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
         * 
         * @return builder
         * 
         */
        public Builder sourceInterface(String sourceInterface) {
            return sourceInterface(Output.of(sourceInterface));
        }

        /**
         * @param vrf VRF name.
         * 
         * @return builder
         * 
         */
        public Builder vrf(Output<String> vrf) {
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

        public BgpPeerArgs build() {
            $.address = Objects.requireNonNull($.address, "expected parameter 'address' to be non-null");
            $.asn = Objects.requireNonNull($.asn, "expected parameter 'asn' to be non-null");
            $.vrf = Objects.requireNonNull($.vrf, "expected parameter 'vrf' to be non-null");
            return $;
        }
    }

}
