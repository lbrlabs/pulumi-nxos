// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetHmmPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetHmmPlainArgs Empty = new GetHmmPlainArgs();

    @Import(name="device")
    private @Nullable String device;

    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }

    private GetHmmPlainArgs() {}

    private GetHmmPlainArgs(GetHmmPlainArgs $) {
        this.device = $.device;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetHmmPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetHmmPlainArgs $;

        public Builder() {
            $ = new GetHmmPlainArgs();
        }

        public Builder(GetHmmPlainArgs defaults) {
            $ = new GetHmmPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder device(@Nullable String device) {
            $.device = device;
            return this;
        }

        public GetHmmPlainArgs build() {
            return $;
        }
    }

}
