// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SubinterfaceVrfArgs extends com.pulumi.resources.ResourceArgs {

    public static final SubinterfaceVrfArgs Empty = new SubinterfaceVrfArgs();

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
     * Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
     * 
     */
    @Import(name="interfaceId", required=true)
    private Output<String> interfaceId;

    /**
     * @return Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
     * 
     */
    public Output<String> interfaceId() {
        return this.interfaceId;
    }

    /**
     * DN of VRF. For example: `sys/inst-VRF1`.
     * 
     */
    @Import(name="vrfDn", required=true)
    private Output<String> vrfDn;

    /**
     * @return DN of VRF. For example: `sys/inst-VRF1`.
     * 
     */
    public Output<String> vrfDn() {
        return this.vrfDn;
    }

    private SubinterfaceVrfArgs() {}

    private SubinterfaceVrfArgs(SubinterfaceVrfArgs $) {
        this.device = $.device;
        this.interfaceId = $.interfaceId;
        this.vrfDn = $.vrfDn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SubinterfaceVrfArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SubinterfaceVrfArgs $;

        public Builder() {
            $ = new SubinterfaceVrfArgs();
        }

        public Builder(SubinterfaceVrfArgs defaults) {
            $ = new SubinterfaceVrfArgs(Objects.requireNonNull(defaults));
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
         * @param interfaceId Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
         * 
         * @return builder
         * 
         */
        public Builder interfaceId(Output<String> interfaceId) {
            $.interfaceId = interfaceId;
            return this;
        }

        /**
         * @param interfaceId Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
         * 
         * @return builder
         * 
         */
        public Builder interfaceId(String interfaceId) {
            return interfaceId(Output.of(interfaceId));
        }

        /**
         * @param vrfDn DN of VRF. For example: `sys/inst-VRF1`.
         * 
         * @return builder
         * 
         */
        public Builder vrfDn(Output<String> vrfDn) {
            $.vrfDn = vrfDn;
            return this;
        }

        /**
         * @param vrfDn DN of VRF. For example: `sys/inst-VRF1`.
         * 
         * @return builder
         * 
         */
        public Builder vrfDn(String vrfDn) {
            return vrfDn(Output.of(vrfDn));
        }

        public SubinterfaceVrfArgs build() {
            $.interfaceId = Objects.requireNonNull($.interfaceId, "expected parameter 'interfaceId' to be non-null");
            $.vrfDn = Objects.requireNonNull($.vrfDn, "expected parameter 'vrfDn' to be non-null");
            return $;
        }
    }

}
