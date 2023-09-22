// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class Ipv4PrefixListRuleEntryState extends com.pulumi.resources.ResourceArgs {

    public static final Ipv4PrefixListRuleEntryState Empty = new Ipv4PrefixListRuleEntryState();

    /**
     * IPv4 Prefix List Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
     * 
     */
    @Import(name="action")
    private @Nullable Output<String> action;

    /**
     * @return IPv4 Prefix List Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
     * 
     */
    public Optional<Output<String>> action() {
        return Optional.ofNullable(this.action);
    }

    /**
     * IPv4 Prefix List Rule Entry criteria. - Choices: `exact`, `inexact` - Default value: `exact`
     * 
     */
    @Import(name="criteria")
    private @Nullable Output<String> criteria;

    /**
     * @return IPv4 Prefix List Rule Entry criteria. - Choices: `exact`, `inexact` - Default value: `exact`
     * 
     */
    public Optional<Output<String>> criteria() {
        return Optional.ofNullable(this.criteria);
    }

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
     * IPv4 Prefix List Rule Entry start range. - Range: `0`-`128` - Default value: `0`
     * 
     */
    @Import(name="fromRange")
    private @Nullable Output<Integer> fromRange;

    /**
     * @return IPv4 Prefix List Rule Entry start range. - Range: `0`-`128` - Default value: `0`
     * 
     */
    public Optional<Output<Integer>> fromRange() {
        return Optional.ofNullable(this.fromRange);
    }

    /**
     * IPv4 Prefix List Rule Entry order. - Range: `0`-`4294967294`
     * 
     */
    @Import(name="order")
    private @Nullable Output<Integer> order;

    /**
     * @return IPv4 Prefix List Rule Entry order. - Range: `0`-`4294967294`
     * 
     */
    public Optional<Output<Integer>> order() {
        return Optional.ofNullable(this.order);
    }

    /**
     * IPv4 Prefix List Rule Entry prefix.
     * 
     */
    @Import(name="prefix")
    private @Nullable Output<String> prefix;

    /**
     * @return IPv4 Prefix List Rule Entry prefix.
     * 
     */
    public Optional<Output<String>> prefix() {
        return Optional.ofNullable(this.prefix);
    }

    /**
     * IPv4 Prefix List Rule name.
     * 
     */
    @Import(name="ruleName")
    private @Nullable Output<String> ruleName;

    /**
     * @return IPv4 Prefix List Rule name.
     * 
     */
    public Optional<Output<String>> ruleName() {
        return Optional.ofNullable(this.ruleName);
    }

    /**
     * IPv4 Prefix List Rule Entry end range. - Range: `0`-`128` - Default value: `0`
     * 
     */
    @Import(name="toRange")
    private @Nullable Output<Integer> toRange;

    /**
     * @return IPv4 Prefix List Rule Entry end range. - Range: `0`-`128` - Default value: `0`
     * 
     */
    public Optional<Output<Integer>> toRange() {
        return Optional.ofNullable(this.toRange);
    }

    private Ipv4PrefixListRuleEntryState() {}

    private Ipv4PrefixListRuleEntryState(Ipv4PrefixListRuleEntryState $) {
        this.action = $.action;
        this.criteria = $.criteria;
        this.device = $.device;
        this.fromRange = $.fromRange;
        this.order = $.order;
        this.prefix = $.prefix;
        this.ruleName = $.ruleName;
        this.toRange = $.toRange;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Ipv4PrefixListRuleEntryState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Ipv4PrefixListRuleEntryState $;

        public Builder() {
            $ = new Ipv4PrefixListRuleEntryState();
        }

        public Builder(Ipv4PrefixListRuleEntryState defaults) {
            $ = new Ipv4PrefixListRuleEntryState(Objects.requireNonNull(defaults));
        }

        /**
         * @param action IPv4 Prefix List Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
         * 
         * @return builder
         * 
         */
        public Builder action(@Nullable Output<String> action) {
            $.action = action;
            return this;
        }

        /**
         * @param action IPv4 Prefix List Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
         * 
         * @return builder
         * 
         */
        public Builder action(String action) {
            return action(Output.of(action));
        }

        /**
         * @param criteria IPv4 Prefix List Rule Entry criteria. - Choices: `exact`, `inexact` - Default value: `exact`
         * 
         * @return builder
         * 
         */
        public Builder criteria(@Nullable Output<String> criteria) {
            $.criteria = criteria;
            return this;
        }

        /**
         * @param criteria IPv4 Prefix List Rule Entry criteria. - Choices: `exact`, `inexact` - Default value: `exact`
         * 
         * @return builder
         * 
         */
        public Builder criteria(String criteria) {
            return criteria(Output.of(criteria));
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
         * @param fromRange IPv4 Prefix List Rule Entry start range. - Range: `0`-`128` - Default value: `0`
         * 
         * @return builder
         * 
         */
        public Builder fromRange(@Nullable Output<Integer> fromRange) {
            $.fromRange = fromRange;
            return this;
        }

        /**
         * @param fromRange IPv4 Prefix List Rule Entry start range. - Range: `0`-`128` - Default value: `0`
         * 
         * @return builder
         * 
         */
        public Builder fromRange(Integer fromRange) {
            return fromRange(Output.of(fromRange));
        }

        /**
         * @param order IPv4 Prefix List Rule Entry order. - Range: `0`-`4294967294`
         * 
         * @return builder
         * 
         */
        public Builder order(@Nullable Output<Integer> order) {
            $.order = order;
            return this;
        }

        /**
         * @param order IPv4 Prefix List Rule Entry order. - Range: `0`-`4294967294`
         * 
         * @return builder
         * 
         */
        public Builder order(Integer order) {
            return order(Output.of(order));
        }

        /**
         * @param prefix IPv4 Prefix List Rule Entry prefix.
         * 
         * @return builder
         * 
         */
        public Builder prefix(@Nullable Output<String> prefix) {
            $.prefix = prefix;
            return this;
        }

        /**
         * @param prefix IPv4 Prefix List Rule Entry prefix.
         * 
         * @return builder
         * 
         */
        public Builder prefix(String prefix) {
            return prefix(Output.of(prefix));
        }

        /**
         * @param ruleName IPv4 Prefix List Rule name.
         * 
         * @return builder
         * 
         */
        public Builder ruleName(@Nullable Output<String> ruleName) {
            $.ruleName = ruleName;
            return this;
        }

        /**
         * @param ruleName IPv4 Prefix List Rule name.
         * 
         * @return builder
         * 
         */
        public Builder ruleName(String ruleName) {
            return ruleName(Output.of(ruleName));
        }

        /**
         * @param toRange IPv4 Prefix List Rule Entry end range. - Range: `0`-`128` - Default value: `0`
         * 
         * @return builder
         * 
         */
        public Builder toRange(@Nullable Output<Integer> toRange) {
            $.toRange = toRange;
            return this;
        }

        /**
         * @param toRange IPv4 Prefix List Rule Entry end range. - Range: `0`-`128` - Default value: `0`
         * 
         * @return builder
         * 
         */
        public Builder toRange(Integer toRange) {
            return toRange(Output.of(toRange));
        }

        public Ipv4PrefixListRuleEntryState build() {
            return $;
        }
    }

}