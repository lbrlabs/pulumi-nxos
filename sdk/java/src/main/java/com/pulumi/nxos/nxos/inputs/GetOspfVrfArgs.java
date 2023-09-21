// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetOspfVrfArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOspfVrfArgs Empty = new GetOspfVrfArgs();

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

    @Import(name="name", required=true)
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }

    private GetOspfVrfArgs() {}

    private GetOspfVrfArgs(GetOspfVrfArgs $) {
        this.device = $.device;
        this.instanceName = $.instanceName;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOspfVrfArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOspfVrfArgs $;

        public Builder() {
            $ = new GetOspfVrfArgs();
        }

        public Builder(GetOspfVrfArgs defaults) {
            $ = new GetOspfVrfArgs(Objects.requireNonNull(defaults));
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

        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public GetOspfVrfArgs build() {
            $.instanceName = Objects.requireNonNull($.instanceName, "expected parameter 'instanceName' to be non-null");
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
