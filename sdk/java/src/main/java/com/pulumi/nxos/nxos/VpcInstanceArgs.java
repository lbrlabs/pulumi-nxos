// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VpcInstanceArgs extends com.pulumi.resources.ResourceArgs {

    public static final VpcInstanceArgs Empty = new VpcInstanceArgs();

    /**
     * Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     * 
     */
    @Import(name="adminState")
    private @Nullable Output<String> adminState;

    /**
     * @return Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
     * 
     */
    public Optional<Output<String>> adminState() {
        return Optional.ofNullable(this.adminState);
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

    private VpcInstanceArgs() {}

    private VpcInstanceArgs(VpcInstanceArgs $) {
        this.adminState = $.adminState;
        this.device = $.device;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcInstanceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcInstanceArgs $;

        public Builder() {
            $ = new VpcInstanceArgs();
        }

        public Builder(VpcInstanceArgs defaults) {
            $ = new VpcInstanceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminState Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
         * 
         * @return builder
         * 
         */
        public Builder adminState(@Nullable Output<String> adminState) {
            $.adminState = adminState;
            return this;
        }

        /**
         * @param adminState Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
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

        public VpcInstanceArgs build() {
            return $;
        }
    }

}
