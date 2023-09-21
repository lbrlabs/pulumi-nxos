// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NveVniArgs extends com.pulumi.resources.ResourceArgs {

    public static final NveVniArgs Empty = new NveVniArgs();

    /**
     * Configures VNI as L3 VNI. - Default value: `false`
     * 
     */
    @Import(name="associateVrf")
    private @Nullable Output<Boolean> associateVrf;

    /**
     * @return Configures VNI as L3 VNI. - Default value: `false`
     * 
     */
    public Optional<Output<Boolean>> associateVrf() {
        return Optional.ofNullable(this.associateVrf);
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
     * Configures multicast group address for VNI. - Default value: `0.0.0.0`
     * 
     */
    @Import(name="multicastGroup")
    private @Nullable Output<String> multicastGroup;

    /**
     * @return Configures multicast group address for VNI. - Default value: `0.0.0.0`
     * 
     */
    public Optional<Output<String>> multicastGroup() {
        return Optional.ofNullable(this.multicastGroup);
    }

    /**
     * Enable or disable Multisite Ingress Replication for VNI(s). - Choices: `enable`, `disable`, `enableOptimized` - Default
     * value: `disable`
     * 
     */
    @Import(name="multisiteIngressReplication")
    private @Nullable Output<String> multisiteIngressReplication;

    /**
     * @return Enable or disable Multisite Ingress Replication for VNI(s). - Choices: `enable`, `disable`, `enableOptimized` - Default
     * value: `disable`
     * 
     */
    public Optional<Output<String>> multisiteIngressReplication() {
        return Optional.ofNullable(this.multisiteIngressReplication);
    }

    /**
     * Enable or disable ARP suppression for VNI(s). - Choices: `enabled`, `disabled`, `off` - Default value: `off`
     * 
     */
    @Import(name="suppressArp")
    private @Nullable Output<String> suppressArp;

    /**
     * @return Enable or disable ARP suppression for VNI(s). - Choices: `enabled`, `disabled`, `off` - Default value: `off`
     * 
     */
    public Optional<Output<String>> suppressArp() {
        return Optional.ofNullable(this.suppressArp);
    }

    /**
     * Virtual Network ID. - Range: `1`-`16777214`
     * 
     */
    @Import(name="vni", required=true)
    private Output<Integer> vni;

    /**
     * @return Virtual Network ID. - Range: `1`-`16777214`
     * 
     */
    public Output<Integer> vni() {
        return this.vni;
    }

    private NveVniArgs() {}

    private NveVniArgs(NveVniArgs $) {
        this.associateVrf = $.associateVrf;
        this.device = $.device;
        this.multicastGroup = $.multicastGroup;
        this.multisiteIngressReplication = $.multisiteIngressReplication;
        this.suppressArp = $.suppressArp;
        this.vni = $.vni;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NveVniArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NveVniArgs $;

        public Builder() {
            $ = new NveVniArgs();
        }

        public Builder(NveVniArgs defaults) {
            $ = new NveVniArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param associateVrf Configures VNI as L3 VNI. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder associateVrf(@Nullable Output<Boolean> associateVrf) {
            $.associateVrf = associateVrf;
            return this;
        }

        /**
         * @param associateVrf Configures VNI as L3 VNI. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder associateVrf(Boolean associateVrf) {
            return associateVrf(Output.of(associateVrf));
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
         * @param multicastGroup Configures multicast group address for VNI. - Default value: `0.0.0.0`
         * 
         * @return builder
         * 
         */
        public Builder multicastGroup(@Nullable Output<String> multicastGroup) {
            $.multicastGroup = multicastGroup;
            return this;
        }

        /**
         * @param multicastGroup Configures multicast group address for VNI. - Default value: `0.0.0.0`
         * 
         * @return builder
         * 
         */
        public Builder multicastGroup(String multicastGroup) {
            return multicastGroup(Output.of(multicastGroup));
        }

        /**
         * @param multisiteIngressReplication Enable or disable Multisite Ingress Replication for VNI(s). - Choices: `enable`, `disable`, `enableOptimized` - Default
         * value: `disable`
         * 
         * @return builder
         * 
         */
        public Builder multisiteIngressReplication(@Nullable Output<String> multisiteIngressReplication) {
            $.multisiteIngressReplication = multisiteIngressReplication;
            return this;
        }

        /**
         * @param multisiteIngressReplication Enable or disable Multisite Ingress Replication for VNI(s). - Choices: `enable`, `disable`, `enableOptimized` - Default
         * value: `disable`
         * 
         * @return builder
         * 
         */
        public Builder multisiteIngressReplication(String multisiteIngressReplication) {
            return multisiteIngressReplication(Output.of(multisiteIngressReplication));
        }

        /**
         * @param suppressArp Enable or disable ARP suppression for VNI(s). - Choices: `enabled`, `disabled`, `off` - Default value: `off`
         * 
         * @return builder
         * 
         */
        public Builder suppressArp(@Nullable Output<String> suppressArp) {
            $.suppressArp = suppressArp;
            return this;
        }

        /**
         * @param suppressArp Enable or disable ARP suppression for VNI(s). - Choices: `enabled`, `disabled`, `off` - Default value: `off`
         * 
         * @return builder
         * 
         */
        public Builder suppressArp(String suppressArp) {
            return suppressArp(Output.of(suppressArp));
        }

        /**
         * @param vni Virtual Network ID. - Range: `1`-`16777214`
         * 
         * @return builder
         * 
         */
        public Builder vni(Output<Integer> vni) {
            $.vni = vni;
            return this;
        }

        /**
         * @param vni Virtual Network ID. - Range: `1`-`16777214`
         * 
         * @return builder
         * 
         */
        public Builder vni(Integer vni) {
            return vni(Output.of(vni));
        }

        public NveVniArgs build() {
            $.vni = Objects.requireNonNull($.vni, "expected parameter 'vni' to be non-null");
            return $;
        }
    }

}
