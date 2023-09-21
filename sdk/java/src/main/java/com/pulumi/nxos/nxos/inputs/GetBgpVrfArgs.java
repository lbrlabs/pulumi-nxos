// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBgpVrfArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBgpVrfArgs Empty = new GetBgpVrfArgs();

    @Import(name="asn", required=true)
    private Output<String> asn;

    public Output<String> asn() {
        return this.asn;
    }

    @Import(name="device")
    private @Nullable Output<String> device;

    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    @Import(name="name", required=true)
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }

    private GetBgpVrfArgs() {}

    private GetBgpVrfArgs(GetBgpVrfArgs $) {
        this.asn = $.asn;
        this.device = $.device;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBgpVrfArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBgpVrfArgs $;

        public Builder() {
            $ = new GetBgpVrfArgs();
        }

        public Builder(GetBgpVrfArgs defaults) {
            $ = new GetBgpVrfArgs(Objects.requireNonNull(defaults));
        }

        public Builder asn(Output<String> asn) {
            $.asn = asn;
            return this;
        }

        public Builder asn(String asn) {
            return asn(Output.of(asn));
        }

        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        public Builder device(String device) {
            return device(Output.of(device));
        }

        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetBgpVrfArgs build() {
            $.asn = Objects.requireNonNull($.asn, "expected parameter 'asn' to be non-null");
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
