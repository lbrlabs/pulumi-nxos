// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class HmmInstanceState extends com.pulumi.resources.ResourceArgs {

    public static final HmmInstanceState Empty = new HmmInstanceState();

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
     * Anycast Gateway MAC address. - Default value: `enabled`
     * 
     */
    @Import(name="anycastMac")
    private @Nullable Output<String> anycastMac;

    /**
     * @return Anycast Gateway MAC address. - Default value: `enabled`
     * 
     */
    public Optional<Output<String>> anycastMac() {
        return Optional.ofNullable(this.anycastMac);
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

    private HmmInstanceState() {}

    private HmmInstanceState(HmmInstanceState $) {
        this.adminState = $.adminState;
        this.anycastMac = $.anycastMac;
        this.device = $.device;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(HmmInstanceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private HmmInstanceState $;

        public Builder() {
            $ = new HmmInstanceState();
        }

        public Builder(HmmInstanceState defaults) {
            $ = new HmmInstanceState(Objects.requireNonNull(defaults));
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
         * @param anycastMac Anycast Gateway MAC address. - Default value: `enabled`
         * 
         * @return builder
         * 
         */
        public Builder anycastMac(@Nullable Output<String> anycastMac) {
            $.anycastMac = anycastMac;
            return this;
        }

        /**
         * @param anycastMac Anycast Gateway MAC address. - Default value: `enabled`
         * 
         * @return builder
         * 
         */
        public Builder anycastMac(String anycastMac) {
            return anycastMac(Output.of(anycastMac));
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

        public HmmInstanceState build() {
            return $;
        }
    }

}
