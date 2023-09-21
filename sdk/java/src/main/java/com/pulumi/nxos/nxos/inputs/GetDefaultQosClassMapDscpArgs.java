// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDefaultQosClassMapDscpArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDefaultQosClassMapDscpArgs Empty = new GetDefaultQosClassMapDscpArgs();

    @Import(name="classMapName", required=true)
    private Output<String> classMapName;

    public Output<String> classMapName() {
        return this.classMapName;
    }

    @Import(name="device")
    private @Nullable Output<String> device;

    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    @Import(name="value", required=true)
    private Output<String> value;

    public Output<String> value() {
        return this.value;
    }

    private GetDefaultQosClassMapDscpArgs() {}

    private GetDefaultQosClassMapDscpArgs(GetDefaultQosClassMapDscpArgs $) {
        this.classMapName = $.classMapName;
        this.device = $.device;
        this.value = $.value;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDefaultQosClassMapDscpArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDefaultQosClassMapDscpArgs $;

        public Builder() {
            $ = new GetDefaultQosClassMapDscpArgs();
        }

        public Builder(GetDefaultQosClassMapDscpArgs defaults) {
            $ = new GetDefaultQosClassMapDscpArgs(Objects.requireNonNull(defaults));
        }

        public Builder classMapName(Output<String> classMapName) {
            $.classMapName = classMapName;
            return this;
        }

        public Builder classMapName(String classMapName) {
            return classMapName(Output.of(classMapName));
        }

        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        public Builder device(String device) {
            return device(Output.of(device));
        }

        public Builder value(Output<String> value) {
            $.value = value;
            return this;
        }

        public Builder value(String value) {
            return value(Output.of(value));
        }

        public GetDefaultQosClassMapDscpArgs build() {
            $.classMapName = Objects.requireNonNull($.classMapName, "expected parameter 'classMapName' to be non-null");
            $.value = Objects.requireNonNull($.value, "expected parameter 'value' to be non-null");
            return $;
        }
    }

}