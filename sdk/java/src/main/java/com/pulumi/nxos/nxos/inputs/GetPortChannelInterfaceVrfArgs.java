// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPortChannelInterfaceVrfArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPortChannelInterfaceVrfArgs Empty = new GetPortChannelInterfaceVrfArgs();

    @Import(name="device")
    private @Nullable Output<String> device;

    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    @Import(name="interfaceId", required=true)
    private Output<String> interfaceId;

    public Output<String> interfaceId() {
        return this.interfaceId;
    }

    private GetPortChannelInterfaceVrfArgs() {}

    private GetPortChannelInterfaceVrfArgs(GetPortChannelInterfaceVrfArgs $) {
        this.device = $.device;
        this.interfaceId = $.interfaceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPortChannelInterfaceVrfArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPortChannelInterfaceVrfArgs $;

        public Builder() {
            $ = new GetPortChannelInterfaceVrfArgs();
        }

        public Builder(GetPortChannelInterfaceVrfArgs defaults) {
            $ = new GetPortChannelInterfaceVrfArgs(Objects.requireNonNull(defaults));
        }

        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        public Builder device(String device) {
            return device(Output.of(device));
        }

        public Builder interfaceId(Output<String> interfaceId) {
            $.interfaceId = interfaceId;
            return this;
        }

        public Builder interfaceId(String interfaceId) {
            return interfaceId(Output.of(interfaceId));
        }

        public GetPortChannelInterfaceVrfArgs build() {
            $.interfaceId = Objects.requireNonNull($.interfaceId, "expected parameter 'interfaceId' to be non-null");
            return $;
        }
    }

}