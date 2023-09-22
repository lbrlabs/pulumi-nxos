// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVrfRouteTargetArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVrfRouteTargetArgs Empty = new GetVrfRouteTargetArgs();

    /**
     * Address family.
     * 
     */
    @Import(name="addressFamily", required=true)
    private Output<String> addressFamily;

    /**
     * @return Address family.
     * 
     */
    public Output<String> addressFamily() {
        return this.addressFamily;
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
     * Route Target direction.
     * 
     */
    @Import(name="direction", required=true)
    private Output<String> direction;

    /**
     * @return Route Target direction.
     * 
     */
    public Output<String> direction() {
        return this.direction;
    }

    /**
     * Route Target in NX-OS DME format.
     * 
     */
    @Import(name="routeTarget", required=true)
    private Output<String> routeTarget;

    /**
     * @return Route Target in NX-OS DME format.
     * 
     */
    public Output<String> routeTarget() {
        return this.routeTarget;
    }

    /**
     * Route Target Address Family.
     * 
     */
    @Import(name="routeTargetAddressFamily", required=true)
    private Output<String> routeTargetAddressFamily;

    /**
     * @return Route Target Address Family.
     * 
     */
    public Output<String> routeTargetAddressFamily() {
        return this.routeTargetAddressFamily;
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

    private GetVrfRouteTargetArgs() {}

    private GetVrfRouteTargetArgs(GetVrfRouteTargetArgs $) {
        this.addressFamily = $.addressFamily;
        this.device = $.device;
        this.direction = $.direction;
        this.routeTarget = $.routeTarget;
        this.routeTargetAddressFamily = $.routeTargetAddressFamily;
        this.vrf = $.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVrfRouteTargetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVrfRouteTargetArgs $;

        public Builder() {
            $ = new GetVrfRouteTargetArgs();
        }

        public Builder(GetVrfRouteTargetArgs defaults) {
            $ = new GetVrfRouteTargetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressFamily Address family.
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(Output<String> addressFamily) {
            $.addressFamily = addressFamily;
            return this;
        }

        /**
         * @param addressFamily Address family.
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(String addressFamily) {
            return addressFamily(Output.of(addressFamily));
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
         * @param direction Route Target direction.
         * 
         * @return builder
         * 
         */
        public Builder direction(Output<String> direction) {
            $.direction = direction;
            return this;
        }

        /**
         * @param direction Route Target direction.
         * 
         * @return builder
         * 
         */
        public Builder direction(String direction) {
            return direction(Output.of(direction));
        }

        /**
         * @param routeTarget Route Target in NX-OS DME format.
         * 
         * @return builder
         * 
         */
        public Builder routeTarget(Output<String> routeTarget) {
            $.routeTarget = routeTarget;
            return this;
        }

        /**
         * @param routeTarget Route Target in NX-OS DME format.
         * 
         * @return builder
         * 
         */
        public Builder routeTarget(String routeTarget) {
            return routeTarget(Output.of(routeTarget));
        }

        /**
         * @param routeTargetAddressFamily Route Target Address Family.
         * 
         * @return builder
         * 
         */
        public Builder routeTargetAddressFamily(Output<String> routeTargetAddressFamily) {
            $.routeTargetAddressFamily = routeTargetAddressFamily;
            return this;
        }

        /**
         * @param routeTargetAddressFamily Route Target Address Family.
         * 
         * @return builder
         * 
         */
        public Builder routeTargetAddressFamily(String routeTargetAddressFamily) {
            return routeTargetAddressFamily(Output.of(routeTargetAddressFamily));
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

        public GetVrfRouteTargetArgs build() {
            $.addressFamily = Objects.requireNonNull($.addressFamily, "expected parameter 'addressFamily' to be non-null");
            $.direction = Objects.requireNonNull($.direction, "expected parameter 'direction' to be non-null");
            $.routeTarget = Objects.requireNonNull($.routeTarget, "expected parameter 'routeTarget' to be non-null");
            $.routeTargetAddressFamily = Objects.requireNonNull($.routeTargetAddressFamily, "expected parameter 'routeTargetAddressFamily' to be non-null");
            $.vrf = Objects.requireNonNull($.vrf, "expected parameter 'vrf' to be non-null");
            return $;
        }
    }

}