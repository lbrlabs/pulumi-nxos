// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIpv4InterfaceAddressArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIpv4InterfaceAddressArgs Empty = new GetIpv4InterfaceAddressArgs();

    /**
     * IPv4 address.
     * 
     */
    @Import(name="address", required=true)
    private Output<String> address;

    /**
     * @return IPv4 address.
     * 
     */
    public Output<String> address() {
        return this.address;
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
     * VRF name.
     * 
     */
    @Import(name="vrf", required=true)
    private Output<String> vrf;

    /**
     * @return VRF name.
     * 
     */
    public Output<String> vrf() {
        return this.vrf;
    }

    private GetIpv4InterfaceAddressArgs() {}

    private GetIpv4InterfaceAddressArgs(GetIpv4InterfaceAddressArgs $) {
        this.address = $.address;
        this.device = $.device;
        this.interfaceId = $.interfaceId;
        this.vrf = $.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIpv4InterfaceAddressArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIpv4InterfaceAddressArgs $;

        public Builder() {
            $ = new GetIpv4InterfaceAddressArgs();
        }

        public Builder(GetIpv4InterfaceAddressArgs defaults) {
            $ = new GetIpv4InterfaceAddressArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param address IPv4 address.
         * 
         * @return builder
         * 
         */
        public Builder address(Output<String> address) {
            $.address = address;
            return this;
        }

        /**
         * @param address IPv4 address.
         * 
         * @return builder
         * 
         */
        public Builder address(String address) {
            return address(Output.of(address));
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
         * @param vrf VRF name.
         * 
         * @return builder
         * 
         */
        public Builder vrf(Output<String> vrf) {
            $.vrf = vrf;
            return this;
        }

        /**
         * @param vrf VRF name.
         * 
         * @return builder
         * 
         */
        public Builder vrf(String vrf) {
            return vrf(Output.of(vrf));
        }

        public GetIpv4InterfaceAddressArgs build() {
            $.address = Objects.requireNonNull($.address, "expected parameter 'address' to be non-null");
            $.interfaceId = Objects.requireNonNull($.interfaceId, "expected parameter 'interfaceId' to be non-null");
            $.vrf = Objects.requireNonNull($.vrf, "expected parameter 'vrf' to be non-null");
            return $;
        }
    }

}
