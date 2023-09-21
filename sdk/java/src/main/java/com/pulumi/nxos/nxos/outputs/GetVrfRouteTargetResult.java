// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVrfRouteTargetResult {
    private String addressFamily;
    private @Nullable String device;
    private String direction;
    private String id;
    private String routeTarget;
    private String routeTargetAddressFamily;
    private String vrf;

    private GetVrfRouteTargetResult() {}
    public String addressFamily() {
        return this.addressFamily;
    }
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public String direction() {
        return this.direction;
    }
    public String id() {
        return this.id;
    }
    public String routeTarget() {
        return this.routeTarget;
    }
    public String routeTargetAddressFamily() {
        return this.routeTargetAddressFamily;
    }
    public String vrf() {
        return this.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVrfRouteTargetResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String addressFamily;
        private @Nullable String device;
        private String direction;
        private String id;
        private String routeTarget;
        private String routeTargetAddressFamily;
        private String vrf;
        public Builder() {}
        public Builder(GetVrfRouteTargetResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.addressFamily = defaults.addressFamily;
    	      this.device = defaults.device;
    	      this.direction = defaults.direction;
    	      this.id = defaults.id;
    	      this.routeTarget = defaults.routeTarget;
    	      this.routeTargetAddressFamily = defaults.routeTargetAddressFamily;
    	      this.vrf = defaults.vrf;
        }

        @CustomType.Setter
        public Builder addressFamily(String addressFamily) {
            this.addressFamily = Objects.requireNonNull(addressFamily);
            return this;
        }
        @CustomType.Setter
        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        @CustomType.Setter
        public Builder direction(String direction) {
            this.direction = Objects.requireNonNull(direction);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder routeTarget(String routeTarget) {
            this.routeTarget = Objects.requireNonNull(routeTarget);
            return this;
        }
        @CustomType.Setter
        public Builder routeTargetAddressFamily(String routeTargetAddressFamily) {
            this.routeTargetAddressFamily = Objects.requireNonNull(routeTargetAddressFamily);
            return this;
        }
        @CustomType.Setter
        public Builder vrf(String vrf) {
            this.vrf = Objects.requireNonNull(vrf);
            return this;
        }
        public GetVrfRouteTargetResult build() {
            final var o = new GetVrfRouteTargetResult();
            o.addressFamily = addressFamily;
            o.device = device;
            o.direction = direction;
            o.id = id;
            o.routeTarget = routeTarget;
            o.routeTargetAddressFamily = routeTargetAddressFamily;
            o.vrf = vrf;
            return o;
        }
    }
}