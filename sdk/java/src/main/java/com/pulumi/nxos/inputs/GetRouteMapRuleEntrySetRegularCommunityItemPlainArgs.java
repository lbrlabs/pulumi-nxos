// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetRouteMapRuleEntrySetRegularCommunityItemPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRouteMapRuleEntrySetRegularCommunityItemPlainArgs Empty = new GetRouteMapRuleEntrySetRegularCommunityItemPlainArgs();

    /**
     * Set Community.
     * 
     */
    @Import(name="community", required=true)
    private String community;

    /**
     * @return Set Community.
     * 
     */
    public String community() {
        return this.community;
    }

    /**
     * A device name from the provider configuration.
     * 
     */
    @Import(name="device")
    private @Nullable String device;

    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }

    /**
     * Route-Map Rule Entry order.
     * 
     */
    @Import(name="order", required=true)
    private Integer order;

    /**
     * @return Route-Map Rule Entry order.
     * 
     */
    public Integer order() {
        return this.order;
    }

    /**
     * Route Map rule name.
     * 
     */
    @Import(name="ruleName", required=true)
    private String ruleName;

    /**
     * @return Route Map rule name.
     * 
     */
    public String ruleName() {
        return this.ruleName;
    }

    private GetRouteMapRuleEntrySetRegularCommunityItemPlainArgs() {}

    private GetRouteMapRuleEntrySetRegularCommunityItemPlainArgs(GetRouteMapRuleEntrySetRegularCommunityItemPlainArgs $) {
        this.community = $.community;
        this.device = $.device;
        this.order = $.order;
        this.ruleName = $.ruleName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRouteMapRuleEntrySetRegularCommunityItemPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRouteMapRuleEntrySetRegularCommunityItemPlainArgs $;

        public Builder() {
            $ = new GetRouteMapRuleEntrySetRegularCommunityItemPlainArgs();
        }

        public Builder(GetRouteMapRuleEntrySetRegularCommunityItemPlainArgs defaults) {
            $ = new GetRouteMapRuleEntrySetRegularCommunityItemPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param community Set Community.
         * 
         * @return builder
         * 
         */
        public Builder community(String community) {
            $.community = community;
            return this;
        }

        /**
         * @param device A device name from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder device(@Nullable String device) {
            $.device = device;
            return this;
        }

        /**
         * @param order Route-Map Rule Entry order.
         * 
         * @return builder
         * 
         */
        public Builder order(Integer order) {
            $.order = order;
            return this;
        }

        /**
         * @param ruleName Route Map rule name.
         * 
         * @return builder
         * 
         */
        public Builder ruleName(String ruleName) {
            $.ruleName = ruleName;
            return this;
        }

        public GetRouteMapRuleEntrySetRegularCommunityItemPlainArgs build() {
            $.community = Objects.requireNonNull($.community, "expected parameter 'community' to be non-null");
            $.order = Objects.requireNonNull($.order, "expected parameter 'order' to be non-null");
            $.ruleName = Objects.requireNonNull($.ruleName, "expected parameter 'ruleName' to be non-null");
            return $;
        }
    }

}
