// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VrfState extends com.pulumi.resources.ResourceArgs {

    public static final VrfState Empty = new VrfState();

    /**
     * VRF description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return VRF description.
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
     * Encap for this Context, supported formats: `unknown`, `vlan-%d` or `vxlan-%d`. - Default value: `unknown`
     * 
     */
    @Import(name="encap")
    private @Nullable Output<String> encap;

    /**
     * @return Encap for this Context, supported formats: `unknown`, `vlan-%d` or `vxlan-%d`. - Default value: `unknown`
     * 
     */
    public Optional<Output<String>> encap() {
        return Optional.ofNullable(this.encap);
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

    private VrfState() {}

    private VrfState(VrfState $) {
        this.description = $.description;
        this.device = $.device;
        this.encap = $.encap;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VrfState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VrfState $;

        public Builder() {
            $ = new VrfState();
        }

        public Builder(VrfState defaults) {
            $ = new VrfState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description VRF description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description VRF description.
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
         * @param encap Encap for this Context, supported formats: `unknown`, `vlan-%d` or `vxlan-%d`. - Default value: `unknown`
         * 
         * @return builder
         * 
         */
        public Builder encap(@Nullable Output<String> encap) {
            $.encap = encap;
            return this;
        }

        /**
         * @param encap Encap for this Context, supported formats: `unknown`, `vlan-%d` or `vxlan-%d`. - Default value: `unknown`
         * 
         * @return builder
         * 
         */
        public Builder encap(String encap) {
            return encap(Output.of(encap));
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

        public VrfState build() {
            return $;
        }
    }

}
