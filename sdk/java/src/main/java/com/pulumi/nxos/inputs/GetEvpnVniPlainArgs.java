// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEvpnVniPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEvpnVniPlainArgs Empty = new GetEvpnVniPlainArgs();

    /**
     * A device name from the provider configuration.
     * 
     */
    @Import(name="device")
    private @Nullable String device;

    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }

    /**
     * Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     * 
     */
    @Import(name="encap", required=true)
    private String encap;

    /**
     * @return Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     * 
     */
    public String encap() {
        return this.encap;
    }

    private GetEvpnVniPlainArgs() {}

    private GetEvpnVniPlainArgs(GetEvpnVniPlainArgs $) {
        this.device = $.device;
        this.encap = $.encap;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEvpnVniPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEvpnVniPlainArgs $;

        public Builder() {
            $ = new GetEvpnVniPlainArgs();
        }

        public Builder(GetEvpnVniPlainArgs defaults) {
            $ = new GetEvpnVniPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param device A device name from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder device(@Nullable String device) {
            $.device = device;
            return this;
        }

        /**
         * @param encap Encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
         * 
         * @return builder
         * 
         */
        public Builder encap(String encap) {
            $.encap = encap;
            return this;
        }

        public GetEvpnVniPlainArgs build() {
            $.encap = Objects.requireNonNull($.encap, "expected parameter 'encap' to be non-null");
            return $;
        }
    }

}