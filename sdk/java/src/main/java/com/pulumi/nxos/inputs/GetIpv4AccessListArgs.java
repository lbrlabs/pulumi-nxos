// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIpv4AccessListArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIpv4AccessListArgs Empty = new GetIpv4AccessListArgs();

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
     * Access list name.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Access list name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private GetIpv4AccessListArgs() {}

    private GetIpv4AccessListArgs(GetIpv4AccessListArgs $) {
        this.device = $.device;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIpv4AccessListArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIpv4AccessListArgs $;

        public Builder() {
            $ = new GetIpv4AccessListArgs();
        }

        public Builder(GetIpv4AccessListArgs defaults) {
            $ = new GetIpv4AccessListArgs(Objects.requireNonNull(defaults));
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
         * @param name Access list name.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Access list name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetIpv4AccessListArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
