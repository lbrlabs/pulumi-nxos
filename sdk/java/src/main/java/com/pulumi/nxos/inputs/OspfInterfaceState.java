// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OspfInterfaceState extends com.pulumi.resources.ResourceArgs {

    public static final OspfInterfaceState Empty = new OspfInterfaceState();

    /**
     * Advertise secondary IP addresses. - Default value: `true`
     * 
     */
    @Import(name="advertiseSecondaries")
    private @Nullable Output<Boolean> advertiseSecondaries;

    /**
     * @return Advertise secondary IP addresses. - Default value: `true`
     * 
     */
    public Optional<Output<Boolean>> advertiseSecondaries() {
        return Optional.ofNullable(this.advertiseSecondaries);
    }

    /**
     * Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
     * 
     */
    @Import(name="area")
    private @Nullable Output<String> area;

    /**
     * @return Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
     * 
     */
    public Optional<Output<String>> area() {
        return Optional.ofNullable(this.area);
    }

    /**
     * Bidirectional Forwarding Detection (BFD). - Choices: `unspecified`, `enabled`, `disabled` - Default value: `unspecified`
     * 
     */
    @Import(name="bfd")
    private @Nullable Output<String> bfd;

    /**
     * @return Bidirectional Forwarding Detection (BFD). - Choices: `unspecified`, `enabled`, `disabled` - Default value: `unspecified`
     * 
     */
    public Optional<Output<String>> bfd() {
        return Optional.ofNullable(this.bfd);
    }

    /**
     * Specifies the cost of interface. - Range: `0`-`65535` - Default value: `0`
     * 
     */
    @Import(name="cost")
    private @Nullable Output<Integer> cost;

    /**
     * @return Specifies the cost of interface. - Range: `0`-`65535` - Default value: `0`
     * 
     */
    public Optional<Output<Integer>> cost() {
        return Optional.ofNullable(this.cost);
    }

    /**
     * Dead interval, interval after which router declares that neighbor as down. - Range: `0`-`65535` - Default value: `0`
     * 
     */
    @Import(name="deadInterval")
    private @Nullable Output<Integer> deadInterval;

    /**
     * @return Dead interval, interval after which router declares that neighbor as down. - Range: `0`-`65535` - Default value: `0`
     * 
     */
    public Optional<Output<Integer>> deadInterval() {
        return Optional.ofNullable(this.deadInterval);
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
     * Hello interval, interval between hello packets that OSPF sends on the interface. - Range: `0`-`65535` - Default value:
     * `10`
     * 
     */
    @Import(name="helloInterval")
    private @Nullable Output<Integer> helloInterval;

    /**
     * @return Hello interval, interval between hello packets that OSPF sends on the interface. - Range: `0`-`65535` - Default value:
     * `10`
     * 
     */
    public Optional<Output<Integer>> helloInterval() {
        return Optional.ofNullable(this.helloInterval);
    }

    /**
     * OSPF instance name.
     * 
     */
    @Import(name="instanceName")
    private @Nullable Output<String> instanceName;

    /**
     * @return OSPF instance name.
     * 
     */
    public Optional<Output<String>> instanceName() {
        return Optional.ofNullable(this.instanceName);
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
     * Network type. - Choices: `unspecified`, `p2p`, `bcast` - Default value: `unspecified`
     * 
     */
    @Import(name="networkType")
    private @Nullable Output<String> networkType;

    /**
     * @return Network type. - Choices: `unspecified`, `p2p`, `bcast` - Default value: `unspecified`
     * 
     */
    public Optional<Output<String>> networkType() {
        return Optional.ofNullable(this.networkType);
    }

    /**
     * Passive interface control. Interface can be configured as passive or non-passive. - Choices: `unspecified`, `enabled`,
     * `disabled` - Default value: `unspecified`
     * 
     */
    @Import(name="passive")
    private @Nullable Output<String> passive;

    /**
     * @return Passive interface control. Interface can be configured as passive or non-passive. - Choices: `unspecified`, `enabled`,
     * `disabled` - Default value: `unspecified`
     * 
     */
    public Optional<Output<String>> passive() {
        return Optional.ofNullable(this.passive);
    }

    /**
     * Priority, used in determining the designated router on this network. - Range: `0`-`255` - Default value: `1`
     * 
     */
    @Import(name="priority")
    private @Nullable Output<Integer> priority;

    /**
     * @return Priority, used in determining the designated router on this network. - Range: `0`-`255` - Default value: `1`
     * 
     */
    public Optional<Output<Integer>> priority() {
        return Optional.ofNullable(this.priority);
    }

    /**
     * VRF name.
     * 
     */
    @Import(name="vrfName")
    private @Nullable Output<String> vrfName;

    /**
     * @return VRF name.
     * 
     */
    public Optional<Output<String>> vrfName() {
        return Optional.ofNullable(this.vrfName);
    }

    private OspfInterfaceState() {}

    private OspfInterfaceState(OspfInterfaceState $) {
        this.advertiseSecondaries = $.advertiseSecondaries;
        this.area = $.area;
        this.bfd = $.bfd;
        this.cost = $.cost;
        this.deadInterval = $.deadInterval;
        this.device = $.device;
        this.helloInterval = $.helloInterval;
        this.instanceName = $.instanceName;
        this.interfaceId = $.interfaceId;
        this.networkType = $.networkType;
        this.passive = $.passive;
        this.priority = $.priority;
        this.vrfName = $.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OspfInterfaceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OspfInterfaceState $;

        public Builder() {
            $ = new OspfInterfaceState();
        }

        public Builder(OspfInterfaceState defaults) {
            $ = new OspfInterfaceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param advertiseSecondaries Advertise secondary IP addresses. - Default value: `true`
         * 
         * @return builder
         * 
         */
        public Builder advertiseSecondaries(@Nullable Output<Boolean> advertiseSecondaries) {
            $.advertiseSecondaries = advertiseSecondaries;
            return this;
        }

        /**
         * @param advertiseSecondaries Advertise secondary IP addresses. - Default value: `true`
         * 
         * @return builder
         * 
         */
        public Builder advertiseSecondaries(Boolean advertiseSecondaries) {
            return advertiseSecondaries(Output.of(advertiseSecondaries));
        }

        /**
         * @param area Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
         * 
         * @return builder
         * 
         */
        public Builder area(@Nullable Output<String> area) {
            $.area = area;
            return this;
        }

        /**
         * @param area Area identifier to which a network or interface belongs in IPv4 address format. - Default value: `0.0.0.0`
         * 
         * @return builder
         * 
         */
        public Builder area(String area) {
            return area(Output.of(area));
        }

        /**
         * @param bfd Bidirectional Forwarding Detection (BFD). - Choices: `unspecified`, `enabled`, `disabled` - Default value: `unspecified`
         * 
         * @return builder
         * 
         */
        public Builder bfd(@Nullable Output<String> bfd) {
            $.bfd = bfd;
            return this;
        }

        /**
         * @param bfd Bidirectional Forwarding Detection (BFD). - Choices: `unspecified`, `enabled`, `disabled` - Default value: `unspecified`
         * 
         * @return builder
         * 
         */
        public Builder bfd(String bfd) {
            return bfd(Output.of(bfd));
        }

        /**
         * @param cost Specifies the cost of interface. - Range: `0`-`65535` - Default value: `0`
         * 
         * @return builder
         * 
         */
        public Builder cost(@Nullable Output<Integer> cost) {
            $.cost = cost;
            return this;
        }

        /**
         * @param cost Specifies the cost of interface. - Range: `0`-`65535` - Default value: `0`
         * 
         * @return builder
         * 
         */
        public Builder cost(Integer cost) {
            return cost(Output.of(cost));
        }

        /**
         * @param deadInterval Dead interval, interval after which router declares that neighbor as down. - Range: `0`-`65535` - Default value: `0`
         * 
         * @return builder
         * 
         */
        public Builder deadInterval(@Nullable Output<Integer> deadInterval) {
            $.deadInterval = deadInterval;
            return this;
        }

        /**
         * @param deadInterval Dead interval, interval after which router declares that neighbor as down. - Range: `0`-`65535` - Default value: `0`
         * 
         * @return builder
         * 
         */
        public Builder deadInterval(Integer deadInterval) {
            return deadInterval(Output.of(deadInterval));
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
         * @param helloInterval Hello interval, interval between hello packets that OSPF sends on the interface. - Range: `0`-`65535` - Default value:
         * `10`
         * 
         * @return builder
         * 
         */
        public Builder helloInterval(@Nullable Output<Integer> helloInterval) {
            $.helloInterval = helloInterval;
            return this;
        }

        /**
         * @param helloInterval Hello interval, interval between hello packets that OSPF sends on the interface. - Range: `0`-`65535` - Default value:
         * `10`
         * 
         * @return builder
         * 
         */
        public Builder helloInterval(Integer helloInterval) {
            return helloInterval(Output.of(helloInterval));
        }

        /**
         * @param instanceName OSPF instance name.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(@Nullable Output<String> instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        /**
         * @param instanceName OSPF instance name.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(String instanceName) {
            return instanceName(Output.of(instanceName));
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
         * @param networkType Network type. - Choices: `unspecified`, `p2p`, `bcast` - Default value: `unspecified`
         * 
         * @return builder
         * 
         */
        public Builder networkType(@Nullable Output<String> networkType) {
            $.networkType = networkType;
            return this;
        }

        /**
         * @param networkType Network type. - Choices: `unspecified`, `p2p`, `bcast` - Default value: `unspecified`
         * 
         * @return builder
         * 
         */
        public Builder networkType(String networkType) {
            return networkType(Output.of(networkType));
        }

        /**
         * @param passive Passive interface control. Interface can be configured as passive or non-passive. - Choices: `unspecified`, `enabled`,
         * `disabled` - Default value: `unspecified`
         * 
         * @return builder
         * 
         */
        public Builder passive(@Nullable Output<String> passive) {
            $.passive = passive;
            return this;
        }

        /**
         * @param passive Passive interface control. Interface can be configured as passive or non-passive. - Choices: `unspecified`, `enabled`,
         * `disabled` - Default value: `unspecified`
         * 
         * @return builder
         * 
         */
        public Builder passive(String passive) {
            return passive(Output.of(passive));
        }

        /**
         * @param priority Priority, used in determining the designated router on this network. - Range: `0`-`255` - Default value: `1`
         * 
         * @return builder
         * 
         */
        public Builder priority(@Nullable Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority Priority, used in determining the designated router on this network. - Range: `0`-`255` - Default value: `1`
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
        }

        /**
         * @param vrfName VRF name.
         * 
         * @return builder
         * 
         */
        public Builder vrfName(@Nullable Output<String> vrfName) {
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

        public OspfInterfaceState build() {
            return $;
        }
    }

}
