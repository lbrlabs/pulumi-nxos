// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PimInterfaceArgs extends com.pulumi.resources.ResourceArgs {

    public static final PimInterfaceArgs Empty = new PimInterfaceArgs();

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
     * BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
     * 
     */
    @Import(name="bfd")
    private @Nullable Output<String> bfd;

    /**
     * @return BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
     * 
     */
    public Optional<Output<String>> bfd() {
        return Optional.ofNullable(this.bfd);
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
     * Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
     * 
     */
    @Import(name="drPriority")
    private @Nullable Output<Integer> drPriority;

    /**
     * @return Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
     * 
     */
    public Optional<Output<Integer>> drPriority() {
        return Optional.ofNullable(this.drPriority);
    }

    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     * 
     */
    @Import(name="interfaceId", required=true)
    private Output<String> interfaceId;

    /**
     * @return Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     * 
     */
    public Output<String> interfaceId() {
        return this.interfaceId;
    }

    /**
     * Passive interface. - Default value: `false`
     * 
     */
    @Import(name="passive")
    private @Nullable Output<Boolean> passive;

    /**
     * @return Passive interface. - Default value: `false`
     * 
     */
    public Optional<Output<Boolean>> passive() {
        return Optional.ofNullable(this.passive);
    }

    /**
     * Sparse mode. - Default value: `false`
     * 
     */
    @Import(name="sparseMode")
    private @Nullable Output<Boolean> sparseMode;

    /**
     * @return Sparse mode. - Default value: `false`
     * 
     */
    public Optional<Output<Boolean>> sparseMode() {
        return Optional.ofNullable(this.sparseMode);
    }

    /**
     * VRF name.
     * 
     */
    @Import(name="vrfName", required=true)
    private Output<String> vrfName;

    /**
     * @return VRF name.
     * 
     */
    public Output<String> vrfName() {
        return this.vrfName;
    }

    private PimInterfaceArgs() {}

    private PimInterfaceArgs(PimInterfaceArgs $) {
        this.adminState = $.adminState;
        this.bfd = $.bfd;
        this.device = $.device;
        this.drPriority = $.drPriority;
        this.interfaceId = $.interfaceId;
        this.passive = $.passive;
        this.sparseMode = $.sparseMode;
        this.vrfName = $.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PimInterfaceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PimInterfaceArgs $;

        public Builder() {
            $ = new PimInterfaceArgs();
        }

        public Builder(PimInterfaceArgs defaults) {
            $ = new PimInterfaceArgs(Objects.requireNonNull(defaults));
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
         * @param bfd BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
         * 
         * @return builder
         * 
         */
        public Builder bfd(@Nullable Output<String> bfd) {
            $.bfd = bfd;
            return this;
        }

        /**
         * @param bfd BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
         * 
         * @return builder
         * 
         */
        public Builder bfd(String bfd) {
            return bfd(Output.of(bfd));
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
         * @param drPriority Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
         * 
         * @return builder
         * 
         */
        public Builder drPriority(@Nullable Output<Integer> drPriority) {
            $.drPriority = drPriority;
            return this;
        }

        /**
         * @param drPriority Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
         * 
         * @return builder
         * 
         */
        public Builder drPriority(Integer drPriority) {
            return drPriority(Output.of(drPriority));
        }

        /**
         * @param interfaceId Must match first field in the output of `show intf brief`. Example: `eth1/1`.
         * 
         * @return builder
         * 
         */
        public Builder interfaceId(Output<String> interfaceId) {
            $.interfaceId = interfaceId;
            return this;
        }

        /**
         * @param interfaceId Must match first field in the output of `show intf brief`. Example: `eth1/1`.
         * 
         * @return builder
         * 
         */
        public Builder interfaceId(String interfaceId) {
            return interfaceId(Output.of(interfaceId));
        }

        /**
         * @param passive Passive interface. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder passive(@Nullable Output<Boolean> passive) {
            $.passive = passive;
            return this;
        }

        /**
         * @param passive Passive interface. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder passive(Boolean passive) {
            return passive(Output.of(passive));
        }

        /**
         * @param sparseMode Sparse mode. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder sparseMode(@Nullable Output<Boolean> sparseMode) {
            $.sparseMode = sparseMode;
            return this;
        }

        /**
         * @param sparseMode Sparse mode. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder sparseMode(Boolean sparseMode) {
            return sparseMode(Output.of(sparseMode));
        }

        /**
         * @param vrfName VRF name.
         * 
         * @return builder
         * 
         */
        public Builder vrfName(Output<String> vrfName) {
            $.vrfName = vrfName;
            return this;
        }

        /**
         * @param vrfName VRF name.
         * 
         * @return builder
         * 
         */
        public Builder vrfName(String vrfName) {
            return vrfName(Output.of(vrfName));
        }

        public PimInterfaceArgs build() {
            $.interfaceId = Objects.requireNonNull($.interfaceId, "expected parameter 'interfaceId' to be non-null");
            $.vrfName = Objects.requireNonNull($.vrfName, "expected parameter 'vrfName' to be non-null");
            return $;
        }
    }

}
