// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetOspfInterfaceArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOspfInterfaceArgs Empty = new GetOspfInterfaceArgs();

    @Import(name="device")
    private @Nullable Output<String> device;

    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    @Import(name="instanceName", required=true)
    private Output<String> instanceName;

    public Output<String> instanceName() {
        return this.instanceName;
    }

    @Import(name="interfaceId", required=true)
    private Output<String> interfaceId;

    public Output<String> interfaceId() {
        return this.interfaceId;
    }

    @Import(name="vrfName", required=true)
    private Output<String> vrfName;

    public Output<String> vrfName() {
        return this.vrfName;
    }

    private GetOspfInterfaceArgs() {}

    private GetOspfInterfaceArgs(GetOspfInterfaceArgs $) {
        this.device = $.device;
        this.instanceName = $.instanceName;
        this.interfaceId = $.interfaceId;
        this.vrfName = $.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOspfInterfaceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOspfInterfaceArgs $;

        public Builder() {
            $ = new GetOspfInterfaceArgs();
        }

        public Builder(GetOspfInterfaceArgs defaults) {
            $ = new GetOspfInterfaceArgs(Objects.requireNonNull(defaults));
        }

        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        public Builder device(String device) {
            return device(Output.of(device));
        }

        public Builder instanceName(Output<String> instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        public Builder instanceName(String instanceName) {
            return instanceName(Output.of(instanceName));
        }

        public Builder interfaceId(Output<String> interfaceId) {
            $.interfaceId = interfaceId;
            return this;
        }

        public Builder interfaceId(String interfaceId) {
            return interfaceId(Output.of(interfaceId));
        }

        public Builder vrfName(Output<String> vrfName) {
            $.vrfName = vrfName;
            return this;
        }

        public Builder vrfName(String vrfName) {
            return vrfName(Output.of(vrfName));
        }

        public GetOspfInterfaceArgs build() {
            $.instanceName = Objects.requireNonNull($.instanceName, "expected parameter 'instanceName' to be non-null");
            $.interfaceId = Objects.requireNonNull($.interfaceId, "expected parameter 'interfaceId' to be non-null");
            $.vrfName = Objects.requireNonNull($.vrfName, "expected parameter 'vrfName' to be non-null");
            return $;
        }
    }

}