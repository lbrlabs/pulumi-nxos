// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBridgeDomainPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBridgeDomainPlainArgs Empty = new GetBridgeDomainPlainArgs();

    @Import(name="device")
    private @Nullable String device;

    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }

    @Import(name="fabricEncap", required=true)
    private String fabricEncap;

    public String fabricEncap() {
        return this.fabricEncap;
    }

    private GetBridgeDomainPlainArgs() {}

    private GetBridgeDomainPlainArgs(GetBridgeDomainPlainArgs $) {
        this.device = $.device;
        this.fabricEncap = $.fabricEncap;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBridgeDomainPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBridgeDomainPlainArgs $;

        public Builder() {
            $ = new GetBridgeDomainPlainArgs();
        }

        public Builder(GetBridgeDomainPlainArgs defaults) {
            $ = new GetBridgeDomainPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder device(@Nullable String device) {
            $.device = device;
            return this;
        }

        public Builder fabricEncap(String fabricEncap) {
            $.fabricEncap = fabricEncap;
            return this;
        }

        public GetBridgeDomainPlainArgs build() {
            $.fabricEncap = Objects.requireNonNull($.fabricEncap, "expected parameter 'fabricEncap' to be non-null");
            return $;
        }
    }

}
