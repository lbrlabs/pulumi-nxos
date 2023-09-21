// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEvpnVniArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEvpnVniArgs Empty = new GetEvpnVniArgs();

    @Import(name="device")
    private @Nullable Output<String> device;

    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    @Import(name="encap", required=true)
    private Output<String> encap;

    public Output<String> encap() {
        return this.encap;
    }

    private GetEvpnVniArgs() {}

    private GetEvpnVniArgs(GetEvpnVniArgs $) {
        this.device = $.device;
        this.encap = $.encap;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEvpnVniArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEvpnVniArgs $;

        public Builder() {
            $ = new GetEvpnVniArgs();
        }

        public Builder(GetEvpnVniArgs defaults) {
            $ = new GetEvpnVniArgs(Objects.requireNonNull(defaults));
        }

        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        public Builder device(String device) {
            return device(Output.of(device));
        }

        public Builder encap(Output<String> encap) {
            $.encap = encap;
            return this;
        }

        public Builder encap(String encap) {
            return encap(Output.of(encap));
        }

        public GetEvpnVniArgs build() {
            $.encap = Objects.requireNonNull($.encap, "expected parameter 'encap' to be non-null");
            return $;
        }
    }

}
