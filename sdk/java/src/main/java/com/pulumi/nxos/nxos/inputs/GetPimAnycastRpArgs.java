// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPimAnycastRpArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPimAnycastRpArgs Empty = new GetPimAnycastRpArgs();

    @Import(name="device")
    private @Nullable Output<String> device;

    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    @Import(name="vrfName", required=true)
    private Output<String> vrfName;

    public Output<String> vrfName() {
        return this.vrfName;
    }

    private GetPimAnycastRpArgs() {}

    private GetPimAnycastRpArgs(GetPimAnycastRpArgs $) {
        this.device = $.device;
        this.vrfName = $.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPimAnycastRpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPimAnycastRpArgs $;

        public Builder() {
            $ = new GetPimAnycastRpArgs();
        }

        public Builder(GetPimAnycastRpArgs defaults) {
            $ = new GetPimAnycastRpArgs(Objects.requireNonNull(defaults));
        }

        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        public Builder device(String device) {
            return device(Output.of(device));
        }

        public Builder vrfName(Output<String> vrfName) {
            $.vrfName = vrfName;
            return this;
        }

        public Builder vrfName(String vrfName) {
            return vrfName(Output.of(vrfName));
        }

        public GetPimAnycastRpArgs build() {
            $.vrfName = Objects.requireNonNull($.vrfName, "expected parameter 'vrfName' to be non-null");
            return $;
        }
    }

}