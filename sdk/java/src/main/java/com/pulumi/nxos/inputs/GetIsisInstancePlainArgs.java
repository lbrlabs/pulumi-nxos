// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIsisInstancePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIsisInstancePlainArgs Empty = new GetIsisInstancePlainArgs();

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
     * IS-IS instance name.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return IS-IS instance name.
     * 
     */
    public String name() {
        return this.name;
    }

    private GetIsisInstancePlainArgs() {}

    private GetIsisInstancePlainArgs(GetIsisInstancePlainArgs $) {
        this.device = $.device;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIsisInstancePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIsisInstancePlainArgs $;

        public Builder() {
            $ = new GetIsisInstancePlainArgs();
        }

        public Builder(GetIsisInstancePlainArgs defaults) {
            $ = new GetIsisInstancePlainArgs(Objects.requireNonNull(defaults));
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
         * @param name IS-IS instance name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public GetIsisInstancePlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}