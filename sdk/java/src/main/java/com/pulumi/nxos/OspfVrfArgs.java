// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OspfVrfArgs extends com.pulumi.resources.ResourceArgs {

    public static final OspfVrfArgs Empty = new OspfVrfArgs();

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
     * Bandwidth reference value. - Range: `0`-`4294967295` - Default value: `40000`
     * 
     */
    @Import(name="bandwidthReference")
    private @Nullable Output<Integer> bandwidthReference;

    /**
     * @return Bandwidth reference value. - Range: `0`-`4294967295` - Default value: `40000`
     * 
     */
    public Optional<Output<Integer>> bandwidthReference() {
        return Optional.ofNullable(this.bandwidthReference);
    }

    /**
     * Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
     * 
     */
    @Import(name="banwidthReferenceUnit")
    private @Nullable Output<String> banwidthReferenceUnit;

    /**
     * @return Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
     * 
     */
    public Optional<Output<String>> banwidthReferenceUnit() {
        return Optional.ofNullable(this.banwidthReferenceUnit);
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
     * Administrative distance preference. - Range: `1`-`255` - Default value: `110`
     * 
     */
    @Import(name="distance")
    private @Nullable Output<Integer> distance;

    /**
     * @return Administrative distance preference. - Range: `1`-`255` - Default value: `110`
     * 
     */
    public Optional<Output<Integer>> distance() {
        return Optional.ofNullable(this.distance);
    }

    /**
     * OSPF instance name.
     * 
     */
    @Import(name="instanceName", required=true)
    private Output<String> instanceName;

    /**
     * @return OSPF instance name.
     * 
     */
    public Output<String> instanceName() {
        return this.instanceName;
    }

    /**
     * VRF name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return VRF name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Router ID. - Default value: `0.0.0.0`
     * 
     */
    @Import(name="routerId")
    private @Nullable Output<String> routerId;

    /**
     * @return Router ID. - Default value: `0.0.0.0`
     * 
     */
    public Optional<Output<String>> routerId() {
        return Optional.ofNullable(this.routerId);
    }

    private OspfVrfArgs() {}

    private OspfVrfArgs(OspfVrfArgs $) {
        this.adminState = $.adminState;
        this.bandwidthReference = $.bandwidthReference;
        this.banwidthReferenceUnit = $.banwidthReferenceUnit;
        this.device = $.device;
        this.distance = $.distance;
        this.instanceName = $.instanceName;
        this.name = $.name;
        this.routerId = $.routerId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OspfVrfArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OspfVrfArgs $;

        public Builder() {
            $ = new OspfVrfArgs();
        }

        public Builder(OspfVrfArgs defaults) {
            $ = new OspfVrfArgs(Objects.requireNonNull(defaults));
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
         * @param bandwidthReference Bandwidth reference value. - Range: `0`-`4294967295` - Default value: `40000`
         * 
         * @return builder
         * 
         */
        public Builder bandwidthReference(@Nullable Output<Integer> bandwidthReference) {
            $.bandwidthReference = bandwidthReference;
            return this;
        }

        /**
         * @param bandwidthReference Bandwidth reference value. - Range: `0`-`4294967295` - Default value: `40000`
         * 
         * @return builder
         * 
         */
        public Builder bandwidthReference(Integer bandwidthReference) {
            return bandwidthReference(Output.of(bandwidthReference));
        }

        /**
         * @param banwidthReferenceUnit Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
         * 
         * @return builder
         * 
         */
        public Builder banwidthReferenceUnit(@Nullable Output<String> banwidthReferenceUnit) {
            $.banwidthReferenceUnit = banwidthReferenceUnit;
            return this;
        }

        /**
         * @param banwidthReferenceUnit Bandwidth reference unit. - Choices: `mbps`, `gbps` - Default value: `mbps`
         * 
         * @return builder
         * 
         */
        public Builder banwidthReferenceUnit(String banwidthReferenceUnit) {
            return banwidthReferenceUnit(Output.of(banwidthReferenceUnit));
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
         * @param distance Administrative distance preference. - Range: `1`-`255` - Default value: `110`
         * 
         * @return builder
         * 
         */
        public Builder distance(@Nullable Output<Integer> distance) {
            $.distance = distance;
            return this;
        }

        /**
         * @param distance Administrative distance preference. - Range: `1`-`255` - Default value: `110`
         * 
         * @return builder
         * 
         */
        public Builder distance(Integer distance) {
            return distance(Output.of(distance));
        }

        /**
         * @param instanceName OSPF instance name.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(Output<String> instanceName) {
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
         * @param name VRF name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name VRF name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param routerId Router ID. - Default value: `0.0.0.0`
         * 
         * @return builder
         * 
         */
        public Builder routerId(@Nullable Output<String> routerId) {
            $.routerId = routerId;
            return this;
        }

        /**
         * @param routerId Router ID. - Default value: `0.0.0.0`
         * 
         * @return builder
         * 
         */
        public Builder routerId(String routerId) {
            return routerId(Output.of(routerId));
        }

        public OspfVrfArgs build() {
            $.instanceName = Objects.requireNonNull($.instanceName, "expected parameter 'instanceName' to be non-null");
            return $;
        }
    }

}