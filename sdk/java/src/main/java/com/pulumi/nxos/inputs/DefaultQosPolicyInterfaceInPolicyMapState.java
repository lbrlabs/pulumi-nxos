// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DefaultQosPolicyInterfaceInPolicyMapState extends com.pulumi.resources.ResourceArgs {

    public static final DefaultQosPolicyInterfaceInPolicyMapState Empty = new DefaultQosPolicyInterfaceInPolicyMapState();

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
    @Import(name="interfaceId")
    private @Nullable Output<String> interfaceId;

    /**
     * @return Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     * 
     */
    public Optional<Output<String>> interfaceId() {
        return Optional.ofNullable(this.interfaceId);
    }

    /**
     * Policy map name.
     * 
     */
    @Import(name="policyMapName")
    private @Nullable Output<String> policyMapName;

    /**
     * @return Policy map name.
     * 
     */
    public Optional<Output<String>> policyMapName() {
        return Optional.ofNullable(this.policyMapName);
    }

    private DefaultQosPolicyInterfaceInPolicyMapState() {}

    private DefaultQosPolicyInterfaceInPolicyMapState(DefaultQosPolicyInterfaceInPolicyMapState $) {
        this.device = $.device;
        this.interfaceId = $.interfaceId;
        this.policyMapName = $.policyMapName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DefaultQosPolicyInterfaceInPolicyMapState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DefaultQosPolicyInterfaceInPolicyMapState $;

        public Builder() {
            $ = new DefaultQosPolicyInterfaceInPolicyMapState();
        }

        public Builder(DefaultQosPolicyInterfaceInPolicyMapState defaults) {
            $ = new DefaultQosPolicyInterfaceInPolicyMapState(Objects.requireNonNull(defaults));
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
        public Builder interfaceId(@Nullable Output<String> interfaceId) {
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
         * @param policyMapName Policy map name.
         * 
         * @return builder
         * 
         */
        public Builder policyMapName(@Nullable Output<String> policyMapName) {
            $.policyMapName = policyMapName;
            return this;
        }

        /**
         * @param policyMapName Policy map name.
         * 
         * @return builder
         * 
         */
        public Builder policyMapName(String policyMapName) {
            return policyMapName(Output.of(policyMapName));
        }

        public DefaultQosPolicyInterfaceInPolicyMapState build() {
            return $;
        }
    }

}