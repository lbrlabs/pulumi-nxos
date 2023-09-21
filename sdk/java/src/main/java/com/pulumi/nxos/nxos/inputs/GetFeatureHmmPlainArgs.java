// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFeatureHmmPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFeatureHmmPlainArgs Empty = new GetFeatureHmmPlainArgs();

    @Import(name="device")
    private @Nullable String device;

    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }

    private GetFeatureHmmPlainArgs() {}

    private GetFeatureHmmPlainArgs(GetFeatureHmmPlainArgs $) {
        this.device = $.device;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFeatureHmmPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFeatureHmmPlainArgs $;

        public Builder() {
            $ = new GetFeatureHmmPlainArgs();
        }

        public Builder(GetFeatureHmmPlainArgs defaults) {
            $ = new GetFeatureHmmPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder device(@Nullable String device) {
            $.device = device;
            return this;
        }

        public GetFeatureHmmPlainArgs build() {
            return $;
        }
    }

}
