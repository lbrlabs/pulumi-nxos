// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRouteMapRuleEntryArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRouteMapRuleEntryArgs Empty = new GetRouteMapRuleEntryArgs();

    @Import(name="device")
    private @Nullable Output<String> device;

    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    @Import(name="order", required=true)
    private Output<Integer> order;

    public Output<Integer> order() {
        return this.order;
    }

    @Import(name="ruleName", required=true)
    private Output<String> ruleName;

    public Output<String> ruleName() {
        return this.ruleName;
    }

    private GetRouteMapRuleEntryArgs() {}

    private GetRouteMapRuleEntryArgs(GetRouteMapRuleEntryArgs $) {
        this.device = $.device;
        this.order = $.order;
        this.ruleName = $.ruleName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRouteMapRuleEntryArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRouteMapRuleEntryArgs $;

        public Builder() {
            $ = new GetRouteMapRuleEntryArgs();
        }

        public Builder(GetRouteMapRuleEntryArgs defaults) {
            $ = new GetRouteMapRuleEntryArgs(Objects.requireNonNull(defaults));
        }

        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        public Builder device(String device) {
            return device(Output.of(device));
        }

        public Builder order(Output<Integer> order) {
            $.order = order;
            return this;
        }

        public Builder order(Integer order) {
            return order(Output.of(order));
        }

        public Builder ruleName(Output<String> ruleName) {
            $.ruleName = ruleName;
            return this;
        }

        public Builder ruleName(String ruleName) {
            return ruleName(Output.of(ruleName));
        }

        public GetRouteMapRuleEntryArgs build() {
            $.order = Objects.requireNonNull($.order, "expected parameter 'order' to be non-null");
            $.ruleName = Objects.requireNonNull($.ruleName, "expected parameter 'ruleName' to be non-null");
            return $;
        }
    }

}