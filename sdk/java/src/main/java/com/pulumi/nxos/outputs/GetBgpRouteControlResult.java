// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBgpRouteControlResult {
    /**
     * @return Autonomous system number.
     * 
     */
    private String asn;
    /**
     * @return A device name from the provider configuration.
     * 
     */
    private @Nullable String device;
    /**
     * @return Enforce First AS For Ebgp. Can be configured only for VRF default.
     * 
     */
    private String enforceFirstAs;
    /**
     * @return Accelerate the hardware updates for IP/IPv6 adjacencies for neighbor. Can be configured only for VRF default.
     * 
     */
    private String fibAccelerate;
    /**
     * @return The distinguished name of the object.
     * 
     */
    private String id;
    /**
     * @return Log Neighbor Changes.
     * 
     */
    private String logNeighborChanges;
    /**
     * @return Suppress Routes: Advertise only routes that are programmed in hardware to peers. Can be configured only for VRF default.
     * 
     */
    private String suppressRoutes;
    /**
     * @return VRF name.
     * 
     */
    private String vrf;

    private GetBgpRouteControlResult() {}
    /**
     * @return Autonomous system number.
     * 
     */
    public String asn() {
        return this.asn;
    }
    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    /**
     * @return Enforce First AS For Ebgp. Can be configured only for VRF default.
     * 
     */
    public String enforceFirstAs() {
        return this.enforceFirstAs;
    }
    /**
     * @return Accelerate the hardware updates for IP/IPv6 adjacencies for neighbor. Can be configured only for VRF default.
     * 
     */
    public String fibAccelerate() {
        return this.fibAccelerate;
    }
    /**
     * @return The distinguished name of the object.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Log Neighbor Changes.
     * 
     */
    public String logNeighborChanges() {
        return this.logNeighborChanges;
    }
    /**
     * @return Suppress Routes: Advertise only routes that are programmed in hardware to peers. Can be configured only for VRF default.
     * 
     */
    public String suppressRoutes() {
        return this.suppressRoutes;
    }
    /**
     * @return VRF name.
     * 
     */
    public String vrf() {
        return this.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBgpRouteControlResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String asn;
        private @Nullable String device;
        private String enforceFirstAs;
        private String fibAccelerate;
        private String id;
        private String logNeighborChanges;
        private String suppressRoutes;
        private String vrf;
        public Builder() {}
        public Builder(GetBgpRouteControlResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.asn = defaults.asn;
    	      this.device = defaults.device;
    	      this.enforceFirstAs = defaults.enforceFirstAs;
    	      this.fibAccelerate = defaults.fibAccelerate;
    	      this.id = defaults.id;
    	      this.logNeighborChanges = defaults.logNeighborChanges;
    	      this.suppressRoutes = defaults.suppressRoutes;
    	      this.vrf = defaults.vrf;
        }

        @CustomType.Setter
        public Builder asn(String asn) {
            this.asn = Objects.requireNonNull(asn);
            return this;
        }
        @CustomType.Setter
        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        @CustomType.Setter
        public Builder enforceFirstAs(String enforceFirstAs) {
            this.enforceFirstAs = Objects.requireNonNull(enforceFirstAs);
            return this;
        }
        @CustomType.Setter
        public Builder fibAccelerate(String fibAccelerate) {
            this.fibAccelerate = Objects.requireNonNull(fibAccelerate);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder logNeighborChanges(String logNeighborChanges) {
            this.logNeighborChanges = Objects.requireNonNull(logNeighborChanges);
            return this;
        }
        @CustomType.Setter
        public Builder suppressRoutes(String suppressRoutes) {
            this.suppressRoutes = Objects.requireNonNull(suppressRoutes);
            return this;
        }
        @CustomType.Setter
        public Builder vrf(String vrf) {
            this.vrf = Objects.requireNonNull(vrf);
            return this;
        }
        public GetBgpRouteControlResult build() {
            final var o = new GetBgpRouteControlResult();
            o.asn = asn;
            o.device = device;
            o.enforceFirstAs = enforceFirstAs;
            o.fibAccelerate = fibAccelerate;
            o.id = id;
            o.logNeighborChanges = logNeighborChanges;
            o.suppressRoutes = suppressRoutes;
            o.vrf = vrf;
            return o;
        }
    }
}