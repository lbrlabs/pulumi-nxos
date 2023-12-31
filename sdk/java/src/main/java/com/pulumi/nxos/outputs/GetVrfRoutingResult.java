// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVrfRoutingResult {
    /**
     * @return A device name from the provider configuration.
     * 
     */
    private @Nullable String device;
    /**
     * @return The distinguished name of the object.
     * 
     */
    private String id;
    /**
     * @return Route Distinguisher value in NX-OS DME format.
     * 
     */
    private String routeDistinguisher;
    /**
     * @return VRF name.
     * 
     */
    private String vrf;

    private GetVrfRoutingResult() {}
    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    /**
     * @return The distinguished name of the object.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Route Distinguisher value in NX-OS DME format.
     * 
     */
    public String routeDistinguisher() {
        return this.routeDistinguisher;
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

    public static Builder builder(GetVrfRoutingResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String device;
        private String id;
        private String routeDistinguisher;
        private String vrf;
        public Builder() {}
        public Builder(GetVrfRoutingResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.device = defaults.device;
    	      this.id = defaults.id;
    	      this.routeDistinguisher = defaults.routeDistinguisher;
    	      this.vrf = defaults.vrf;
        }

        @CustomType.Setter
        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder routeDistinguisher(String routeDistinguisher) {
            this.routeDistinguisher = Objects.requireNonNull(routeDistinguisher);
            return this;
        }
        @CustomType.Setter
        public Builder vrf(String vrf) {
            this.vrf = Objects.requireNonNull(vrf);
            return this;
        }
        public GetVrfRoutingResult build() {
            final var o = new GetVrfRoutingResult();
            o.device = device;
            o.id = id;
            o.routeDistinguisher = routeDistinguisher;
            o.vrf = vrf;
            return o;
        }
    }
}
