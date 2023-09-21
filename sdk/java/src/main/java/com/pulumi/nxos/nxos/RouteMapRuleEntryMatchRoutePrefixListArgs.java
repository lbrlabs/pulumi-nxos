// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RouteMapRuleEntryMatchRoutePrefixListArgs extends com.pulumi.resources.ResourceArgs {

    public static final RouteMapRuleEntryMatchRoutePrefixListArgs Empty = new RouteMapRuleEntryMatchRoutePrefixListArgs();

    /**
     * A device name from the provider configuration.
     * 
     */
    @Import(name="device")
    private @Nullable Output<String> device;

    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    /**
     * Route-Map Rule Entry order. - Range: `0`-`65535`
     * 
     */
    @Import(name="order", required=true)
    private Output<Integer> order;

    /**
     * @return Route-Map Rule Entry order. - Range: `0`-`65535`
     * 
     */
    public Output<Integer> order() {
        return this.order;
    }

    /**
     * DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
     * 
     */
    @Import(name="prefixListDn", required=true)
    private Output<String> prefixListDn;

    /**
     * @return DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
     * 
     */
    public Output<String> prefixListDn() {
        return this.prefixListDn;
    }

    /**
     * Route Map rule name.
     * 
     */
    @Import(name="ruleName", required=true)
    private Output<String> ruleName;

    /**
     * @return Route Map rule name.
     * 
     */
    public Output<String> ruleName() {
        return this.ruleName;
    }

    private RouteMapRuleEntryMatchRoutePrefixListArgs() {}

    private RouteMapRuleEntryMatchRoutePrefixListArgs(RouteMapRuleEntryMatchRoutePrefixListArgs $) {
        this.device = $.device;
        this.order = $.order;
        this.prefixListDn = $.prefixListDn;
        this.ruleName = $.ruleName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RouteMapRuleEntryMatchRoutePrefixListArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RouteMapRuleEntryMatchRoutePrefixListArgs $;

        public Builder() {
            $ = new RouteMapRuleEntryMatchRoutePrefixListArgs();
        }

        public Builder(RouteMapRuleEntryMatchRoutePrefixListArgs defaults) {
            $ = new RouteMapRuleEntryMatchRoutePrefixListArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param device A device name from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        /**
         * @param device A device name from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder device(String device) {
            return device(Output.of(device));
        }

        /**
         * @param order Route-Map Rule Entry order. - Range: `0`-`65535`
         * 
         * @return builder
         * 
         */
        public Builder order(Output<Integer> order) {
            $.order = order;
            return this;
        }

        /**
         * @param order Route-Map Rule Entry order. - Range: `0`-`65535`
         * 
         * @return builder
         * 
         */
        public Builder order(Integer order) {
            return order(Output.of(order));
        }

        /**
         * @param prefixListDn DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
         * 
         * @return builder
         * 
         */
        public Builder prefixListDn(Output<String> prefixListDn) {
            $.prefixListDn = prefixListDn;
            return this;
        }

        /**
         * @param prefixListDn DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
         * 
         * @return builder
         * 
         */
        public Builder prefixListDn(String prefixListDn) {
            return prefixListDn(Output.of(prefixListDn));
        }

        /**
         * @param ruleName Route Map rule name.
         * 
         * @return builder
         * 
         */
        public Builder ruleName(Output<String> ruleName) {
            $.ruleName = ruleName;
            return this;
        }

        /**
         * @param ruleName Route Map rule name.
         * 
         * @return builder
         * 
         */
        public Builder ruleName(String ruleName) {
            return ruleName(Output.of(ruleName));
        }

        public RouteMapRuleEntryMatchRoutePrefixListArgs build() {
            $.order = Objects.requireNonNull($.order, "expected parameter 'order' to be non-null");
            $.prefixListDn = Objects.requireNonNull($.prefixListDn, "expected parameter 'prefixListDn' to be non-null");
            $.ruleName = Objects.requireNonNull($.ruleName, "expected parameter 'ruleName' to be non-null");
            return $;
        }
    }

}
