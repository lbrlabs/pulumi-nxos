// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetHmmInterfacePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetHmmInterfacePlainArgs Empty = new GetHmmInterfacePlainArgs();

    @Import(name="device")
    private @Nullable String device;

    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }

    @Import(name="interfaceId", required=true)
    private String interfaceId;

    public String interfaceId() {
        return this.interfaceId;
    }

    private GetHmmInterfacePlainArgs() {}

    private GetHmmInterfacePlainArgs(GetHmmInterfacePlainArgs $) {
        this.device = $.device;
        this.interfaceId = $.interfaceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetHmmInterfacePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetHmmInterfacePlainArgs $;

        public Builder() {
            $ = new GetHmmInterfacePlainArgs();
        }

        public Builder(GetHmmInterfacePlainArgs defaults) {
            $ = new GetHmmInterfacePlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder device(@Nullable String device) {
            $.device = device;
            return this;
        }

        public Builder interfaceId(String interfaceId) {
            $.interfaceId = interfaceId;
            return this;
        }

        public GetHmmInterfacePlainArgs build() {
            $.interfaceId = Objects.requireNonNull($.interfaceId, "expected parameter 'interfaceId' to be non-null");
            return $;
        }
    }

}
