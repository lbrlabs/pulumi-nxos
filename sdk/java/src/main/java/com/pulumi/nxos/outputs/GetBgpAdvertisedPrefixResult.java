// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBgpAdvertisedPrefixResult {
    /**
     * @return Address Family.
     * 
     */
    private String addressFamily;
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
     * @return The distinguished name of the object.
     * 
     */
    private String id;
    /**
     * @return IP address of the network or prefix to advertise.
     * 
     */
    private String prefix;
    /**
     * @return Route map to modify attributes.
     * 
     */
    private String routeMap;
    /**
     * @return VRF name.
     * 
     */
    private String vrf;

    private GetBgpAdvertisedPrefixResult() {}
    /**
     * @return Address Family.
     * 
     */
    public String addressFamily() {
        return this.addressFamily;
    }
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
     * @return The distinguished name of the object.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return IP address of the network or prefix to advertise.
     * 
     */
    public String prefix() {
        return this.prefix;
    }
    /**
     * @return Route map to modify attributes.
     * 
     */
    public String routeMap() {
        return this.routeMap;
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

    public static Builder builder(GetBgpAdvertisedPrefixResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String addressFamily;
        private String asn;
        private @Nullable String device;
        private String id;
        private String prefix;
        private String routeMap;
        private String vrf;
        public Builder() {}
        public Builder(GetBgpAdvertisedPrefixResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.addressFamily = defaults.addressFamily;
    	      this.asn = defaults.asn;
    	      this.device = defaults.device;
    	      this.id = defaults.id;
    	      this.prefix = defaults.prefix;
    	      this.routeMap = defaults.routeMap;
    	      this.vrf = defaults.vrf;
        }

        @CustomType.Setter
        public Builder addressFamily(String addressFamily) {
            this.addressFamily = Objects.requireNonNull(addressFamily);
            return this;
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
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder prefix(String prefix) {
            this.prefix = Objects.requireNonNull(prefix);
            return this;
        }
        @CustomType.Setter
        public Builder routeMap(String routeMap) {
            this.routeMap = Objects.requireNonNull(routeMap);
            return this;
        }
        @CustomType.Setter
        public Builder vrf(String vrf) {
            this.vrf = Objects.requireNonNull(vrf);
            return this;
        }
        public GetBgpAdvertisedPrefixResult build() {
            final var o = new GetBgpAdvertisedPrefixResult();
            o.addressFamily = addressFamily;
            o.asn = asn;
            o.device = device;
            o.id = id;
            o.prefix = prefix;
            o.routeMap = routeMap;
            o.vrf = vrf;
            return o;
        }
    }
}
