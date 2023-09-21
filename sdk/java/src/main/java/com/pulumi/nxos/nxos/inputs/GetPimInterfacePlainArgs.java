// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPimInterfacePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPimInterfacePlainArgs Empty = new GetPimInterfacePlainArgs();

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

    @Import(name="vrfName", required=true)
    private String vrfName;

    public String vrfName() {
        return this.vrfName;
    }

    private GetPimInterfacePlainArgs() {}

    private GetPimInterfacePlainArgs(GetPimInterfacePlainArgs $) {
        this.device = $.device;
        this.interfaceId = $.interfaceId;
        this.vrfName = $.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPimInterfacePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPimInterfacePlainArgs $;

        public Builder() {
            $ = new GetPimInterfacePlainArgs();
        }

        public Builder(GetPimInterfacePlainArgs defaults) {
            $ = new GetPimInterfacePlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder device(@Nullable String device) {
            $.device = device;
            return this;
        }

        public Builder interfaceId(String interfaceId) {
            $.interfaceId = interfaceId;
            return this;
        }

        public Builder vrfName(String vrfName) {
            $.vrfName = vrfName;
            return this;
        }

        public GetPimInterfacePlainArgs build() {
            $.interfaceId = Objects.requireNonNull($.interfaceId, "expected parameter 'interfaceId' to be non-null");
            $.vrfName = Objects.requireNonNull($.vrfName, "expected parameter 'vrfName' to be non-null");
            return $;
        }
    }

}