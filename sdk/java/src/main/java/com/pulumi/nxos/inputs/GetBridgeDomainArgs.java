// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBridgeDomainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBridgeDomainArgs Empty = new GetBridgeDomainArgs();

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
     * Fabric encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     * 
     */
    @Import(name="fabricEncap", required=true)
    private Output<String> fabricEncap;

    /**
     * @return Fabric encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     * 
     */
    public Output<String> fabricEncap() {
        return this.fabricEncap;
    }

    private GetBridgeDomainArgs() {}

    private GetBridgeDomainArgs(GetBridgeDomainArgs $) {
        this.device = $.device;
        this.fabricEncap = $.fabricEncap;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBridgeDomainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBridgeDomainArgs $;

        public Builder() {
            $ = new GetBridgeDomainArgs();
        }

        public Builder(GetBridgeDomainArgs defaults) {
            $ = new GetBridgeDomainArgs(Objects.requireNonNull(defaults));
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
         * @param fabricEncap Fabric encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
         * 
         * @return builder
         * 
         */
        public Builder fabricEncap(Output<String> fabricEncap) {
            $.fabricEncap = fabricEncap;
            return this;
        }

        /**
         * @param fabricEncap Fabric encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
         * 
         * @return builder
         * 
         */
        public Builder fabricEncap(String fabricEncap) {
            return fabricEncap(Output.of(fabricEncap));
        }

        public GetBridgeDomainArgs build() {
            $.fabricEncap = Objects.requireNonNull($.fabricEncap, "expected parameter 'fabricEncap' to be non-null");
            return $;
        }
    }

}
