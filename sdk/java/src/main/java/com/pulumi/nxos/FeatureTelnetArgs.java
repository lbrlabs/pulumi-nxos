// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FeatureTelnetArgs extends com.pulumi.resources.ResourceArgs {

    public static final FeatureTelnetArgs Empty = new FeatureTelnetArgs();

    /**
     * Administrative state. - Choices: `enabled`, `disabled`
     * 
     */
    @Import(name="adminState", required=true)
    private Output<String> adminState;

    /**
     * @return Administrative state. - Choices: `enabled`, `disabled`
     * 
     */
    public Output<String> adminState() {
        return this.adminState;
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

    private FeatureTelnetArgs() {}

    private FeatureTelnetArgs(FeatureTelnetArgs $) {
        this.adminState = $.adminState;
        this.device = $.device;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FeatureTelnetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FeatureTelnetArgs $;

        public Builder() {
            $ = new FeatureTelnetArgs();
        }

        public Builder(FeatureTelnetArgs defaults) {
            $ = new FeatureTelnetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminState Administrative state. - Choices: `enabled`, `disabled`
         * 
         * @return builder
         * 
         */
        public Builder adminState(Output<String> adminState) {
            $.adminState = adminState;
            return this;
        }

        /**
         * @param adminState Administrative state. - Choices: `enabled`, `disabled`
         * 
         * @return builder
         * 
         */
        public Builder adminState(String adminState) {
            return adminState(Output.of(adminState));
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

        public FeatureTelnetArgs build() {
            $.adminState = Objects.requireNonNull($.adminState, "expected parameter 'adminState' to be non-null");
            return $;
        }
    }

}
