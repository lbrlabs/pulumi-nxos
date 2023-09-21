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


public final class BgpGracefulRestartState extends com.pulumi.resources.ResourceArgs {

    public static final BgpGracefulRestartState Empty = new BgpGracefulRestartState();

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
     * The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
     * 
     */
    @Import(name="restartInterval")
    private @Nullable Output<Integer> restartInterval;

    /**
     * @return The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
     * 
     */
    public Optional<Output<Integer>> restartInterval() {
        return Optional.ofNullable(this.restartInterval);
    }

    /**
     * The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
     * 
     */
    @Import(name="staleInterval")
    private @Nullable Output<Integer> staleInterval;

    /**
     * @return The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
     * 
     */
    public Optional<Output<Integer>> staleInterval() {
        return Optional.ofNullable(this.staleInterval);
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

    private BgpGracefulRestartState() {}

    private BgpGracefulRestartState(BgpGracefulRestartState $) {
        this.asn = $.asn;
        this.device = $.device;
        this.restartInterval = $.restartInterval;
        this.staleInterval = $.staleInterval;
        this.vrf = $.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BgpGracefulRestartState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BgpGracefulRestartState $;

        public Builder() {
            $ = new BgpGracefulRestartState();
        }

        public Builder(BgpGracefulRestartState defaults) {
            $ = new BgpGracefulRestartState(Objects.requireNonNull(defaults));
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
         * @param restartInterval The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
         * 
         * @return builder
         * 
         */
        public Builder restartInterval(@Nullable Output<Integer> restartInterval) {
            $.restartInterval = restartInterval;
            return this;
        }

        /**
         * @param restartInterval The graceful restart interval. - Range: `1`-`3600` - Default value: `120`
         * 
         * @return builder
         * 
         */
        public Builder restartInterval(Integer restartInterval) {
            return restartInterval(Output.of(restartInterval));
        }

        /**
         * @param staleInterval The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
         * 
         * @return builder
         * 
         */
        public Builder staleInterval(@Nullable Output<Integer> staleInterval) {
            $.staleInterval = staleInterval;
            return this;
        }

        /**
         * @param staleInterval The stale interval for routes advertised by the BGP peer. - Range: `1`-`3600` - Default value: `300`
         * 
         * @return builder
         * 
         */
        public Builder staleInterval(Integer staleInterval) {
            return staleInterval(Output.of(staleInterval));
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

        public BgpGracefulRestartState build() {
            return $;
        }
    }

}
