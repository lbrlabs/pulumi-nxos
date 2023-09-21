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


public final class VpcInterfaceState extends com.pulumi.resources.ResourceArgs {

    public static final VpcInterfaceState Empty = new VpcInterfaceState();

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
     * Port-channel interface DN.
     * 
     */
    @Import(name="portChannelInterfaceDn")
    private @Nullable Output<String> portChannelInterfaceDn;

    /**
     * @return Port-channel interface DN.
     * 
     */
    public Optional<Output<String>> portChannelInterfaceDn() {
        return Optional.ofNullable(this.portChannelInterfaceDn);
    }

    /**
     * The vPC interface identifier. - Range: `1`-`16384`
     * 
     */
    @Import(name="vpcInterfaceId")
    private @Nullable Output<Integer> vpcInterfaceId;

    /**
     * @return The vPC interface identifier. - Range: `1`-`16384`
     * 
     */
    public Optional<Output<Integer>> vpcInterfaceId() {
        return Optional.ofNullable(this.vpcInterfaceId);
    }

    private VpcInterfaceState() {}

    private VpcInterfaceState(VpcInterfaceState $) {
        this.device = $.device;
        this.portChannelInterfaceDn = $.portChannelInterfaceDn;
        this.vpcInterfaceId = $.vpcInterfaceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VpcInterfaceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VpcInterfaceState $;

        public Builder() {
            $ = new VpcInterfaceState();
        }

        public Builder(VpcInterfaceState defaults) {
            $ = new VpcInterfaceState(Objects.requireNonNull(defaults));
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
         * @param portChannelInterfaceDn Port-channel interface DN.
         * 
         * @return builder
         * 
         */
        public Builder portChannelInterfaceDn(@Nullable Output<String> portChannelInterfaceDn) {
            $.portChannelInterfaceDn = portChannelInterfaceDn;
            return this;
        }

        /**
         * @param portChannelInterfaceDn Port-channel interface DN.
         * 
         * @return builder
         * 
         */
        public Builder portChannelInterfaceDn(String portChannelInterfaceDn) {
            return portChannelInterfaceDn(Output.of(portChannelInterfaceDn));
        }

        /**
         * @param vpcInterfaceId The vPC interface identifier. - Range: `1`-`16384`
         * 
         * @return builder
         * 
         */
        public Builder vpcInterfaceId(@Nullable Output<Integer> vpcInterfaceId) {
            $.vpcInterfaceId = vpcInterfaceId;
            return this;
        }

        /**
         * @param vpcInterfaceId The vPC interface identifier. - Range: `1`-`16384`
         * 
         * @return builder
         * 
         */
        public Builder vpcInterfaceId(Integer vpcInterfaceId) {
            return vpcInterfaceId(Output.of(vpcInterfaceId));
        }

        public VpcInterfaceState build() {
            return $;
        }
    }

}
