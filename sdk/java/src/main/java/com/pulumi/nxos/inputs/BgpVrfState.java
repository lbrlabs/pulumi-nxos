// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BgpVrfState extends com.pulumi.resources.ResourceArgs {

    public static final BgpVrfState Empty = new BgpVrfState();

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
     * VRF name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return VRF name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Router ID.
     * 
     */
    @Import(name="routerId")
    private @Nullable Output<String> routerId;

    /**
     * @return Router ID.
     * 
     */
    public Optional<Output<String>> routerId() {
        return Optional.ofNullable(this.routerId);
    }

    private BgpVrfState() {}

    private BgpVrfState(BgpVrfState $) {
        this.asn = $.asn;
        this.device = $.device;
        this.name = $.name;
        this.routerId = $.routerId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BgpVrfState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BgpVrfState $;

        public Builder() {
            $ = new BgpVrfState();
        }

        public Builder(BgpVrfState defaults) {
            $ = new BgpVrfState(Objects.requireNonNull(defaults));
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
         * @param name VRF name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name VRF name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param routerId Router ID.
         * 
         * @return builder
         * 
         */
        public Builder routerId(@Nullable Output<String> routerId) {
            $.routerId = routerId;
            return this;
        }

        /**
         * @param routerId Router ID.
         * 
         * @return builder
         * 
         */
        public Builder routerId(String routerId) {
            return routerId(Output.of(routerId));
        }

        public BgpVrfState build() {
            return $;
        }
    }

}
