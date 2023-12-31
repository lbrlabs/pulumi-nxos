// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFeatureTacacsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFeatureTacacsPlainArgs Empty = new GetFeatureTacacsPlainArgs();

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

    private GetFeatureTacacsPlainArgs() {}

    private GetFeatureTacacsPlainArgs(GetFeatureTacacsPlainArgs $) {
        this.device = $.device;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFeatureTacacsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFeatureTacacsPlainArgs $;

        public Builder() {
            $ = new GetFeatureTacacsPlainArgs();
        }

        public Builder(GetFeatureTacacsPlainArgs defaults) {
            $ = new GetFeatureTacacsPlainArgs(Objects.requireNonNull(defaults));
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

        public GetFeatureTacacsPlainArgs build() {
            return $;
        }
    }

}
