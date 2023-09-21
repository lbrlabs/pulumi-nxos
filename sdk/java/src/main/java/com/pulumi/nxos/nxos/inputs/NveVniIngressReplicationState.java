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


public final class NveVniIngressReplicationState extends com.pulumi.resources.ResourceArgs {

    public static final NveVniIngressReplicationState Empty = new NveVniIngressReplicationState();

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
     * Configure VxLAN Ingress Replication mode. - Choices: `bgp`, `unknown`, `static` - Default value: `unknown`
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return Configure VxLAN Ingress Replication mode. - Choices: `bgp`, `unknown`, `static` - Default value: `unknown`
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * Virtual Network ID. - Range: `1`-`16777214`
     * 
     */
    @Import(name="vni")
    private @Nullable Output<Integer> vni;

    /**
     * @return Virtual Network ID. - Range: `1`-`16777214`
     * 
     */
    public Optional<Output<Integer>> vni() {
        return Optional.ofNullable(this.vni);
    }

    private NveVniIngressReplicationState() {}

    private NveVniIngressReplicationState(NveVniIngressReplicationState $) {
        this.device = $.device;
        this.protocol = $.protocol;
        this.vni = $.vni;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NveVniIngressReplicationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NveVniIngressReplicationState $;

        public Builder() {
            $ = new NveVniIngressReplicationState();
        }

        public Builder(NveVniIngressReplicationState defaults) {
            $ = new NveVniIngressReplicationState(Objects.requireNonNull(defaults));
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
         * @param protocol Configure VxLAN Ingress Replication mode. - Choices: `bgp`, `unknown`, `static` - Default value: `unknown`
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol Configure VxLAN Ingress Replication mode. - Choices: `bgp`, `unknown`, `static` - Default value: `unknown`
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param vni Virtual Network ID. - Range: `1`-`16777214`
         * 
         * @return builder
         * 
         */
        public Builder vni(@Nullable Output<Integer> vni) {
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

        public NveVniIngressReplicationState build() {
            return $;
        }
    }

}