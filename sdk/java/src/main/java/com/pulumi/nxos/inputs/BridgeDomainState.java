// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class BridgeDomainState extends com.pulumi.resources.ResourceArgs {

    public static final BridgeDomainState Empty = new BridgeDomainState();

    /**
     * Access encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `unknown`
     * 
     */
    @Import(name="accessEncap")
    private @Nullable Output<String> accessEncap;

    /**
     * @return Access encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `unknown`
     * 
     */
    public Optional<Output<String>> accessEncap() {
        return Optional.ofNullable(this.accessEncap);
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
     * Fabric encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `unknown`
     * 
     */
    @Import(name="fabricEncap")
    private @Nullable Output<String> fabricEncap;

    /**
     * @return Fabric encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `unknown`
     * 
     */
    public Optional<Output<String>> fabricEncap() {
        return Optional.ofNullable(this.fabricEncap);
    }

    /**
     * Bridge domain name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Bridge domain name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private BridgeDomainState() {}

    private BridgeDomainState(BridgeDomainState $) {
        this.accessEncap = $.accessEncap;
        this.device = $.device;
        this.fabricEncap = $.fabricEncap;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(BridgeDomainState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private BridgeDomainState $;

        public Builder() {
            $ = new BridgeDomainState();
        }

        public Builder(BridgeDomainState defaults) {
            $ = new BridgeDomainState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessEncap Access encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `unknown`
         * 
         * @return builder
         * 
         */
        public Builder accessEncap(@Nullable Output<String> accessEncap) {
            $.accessEncap = accessEncap;
            return this;
        }

        /**
         * @param accessEncap Access encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `unknown`
         * 
         * @return builder
         * 
         */
        public Builder accessEncap(String accessEncap) {
            return accessEncap(Output.of(accessEncap));
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
         * @param fabricEncap Fabric encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `unknown`
         * 
         * @return builder
         * 
         */
        public Builder fabricEncap(@Nullable Output<String> fabricEncap) {
            $.fabricEncap = fabricEncap;
            return this;
        }

        /**
         * @param fabricEncap Fabric encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `unknown`
         * 
         * @return builder
         * 
         */
        public Builder fabricEncap(String fabricEncap) {
            return fabricEncap(Output.of(fabricEncap));
        }

        /**
         * @param name Bridge domain name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Bridge domain name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public BridgeDomainState build() {
            return $;
        }
    }

}
