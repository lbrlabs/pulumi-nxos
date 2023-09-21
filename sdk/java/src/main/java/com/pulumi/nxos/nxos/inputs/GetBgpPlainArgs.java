// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBgpPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBgpPlainArgs Empty = new GetBgpPlainArgs();

    @Import(name="device")
    private @Nullable String device;

    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }

    private GetBgpPlainArgs() {}

    private GetBgpPlainArgs(GetBgpPlainArgs $) {
        this.device = $.device;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBgpPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBgpPlainArgs $;

        public Builder() {
            $ = new GetBgpPlainArgs();
        }

        public Builder(GetBgpPlainArgs defaults) {
            $ = new GetBgpPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder device(@Nullable String device) {
            $.device = device;
            return this;
        }

        public GetBgpPlainArgs build() {
            return $;
        }
    }

}