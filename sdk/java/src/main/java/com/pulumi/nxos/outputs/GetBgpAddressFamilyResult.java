// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBgpAddressFamilyResult {
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
     * @return The next-hop address tracking delay timer for critical next-hop reachability routes.
     * 
     */
    private Integer criticalNexthopTimeout;
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
     * @return The next-hop address tracking delay timer for non-critical next-hop reachability routes.
     * 
     */
    private Integer nonCriticalNexthopTimeout;
    /**
     * @return VRF name.
     * 
     */
    private String vrf;

    private GetBgpAddressFamilyResult() {}
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
     * @return The next-hop address tracking delay timer for critical next-hop reachability routes.
     * 
     */
    public Integer criticalNexthopTimeout() {
        return this.criticalNexthopTimeout;
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
     * @return The next-hop address tracking delay timer for non-critical next-hop reachability routes.
     * 
     */
    public Integer nonCriticalNexthopTimeout() {
        return this.nonCriticalNexthopTimeout;
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

    public static Builder builder(GetBgpAddressFamilyResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String addressFamily;
        private String asn;
        private Integer criticalNexthopTimeout;
        private @Nullable String device;
        private String id;
        private Integer nonCriticalNexthopTimeout;
        private String vrf;
        public Builder() {}
        public Builder(GetBgpAddressFamilyResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.addressFamily = defaults.addressFamily;
    	      this.asn = defaults.asn;
    	      this.criticalNexthopTimeout = defaults.criticalNexthopTimeout;
    	      this.device = defaults.device;
    	      this.id = defaults.id;
    	      this.nonCriticalNexthopTimeout = defaults.nonCriticalNexthopTimeout;
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
        public Builder criticalNexthopTimeout(Integer criticalNexthopTimeout) {
            this.criticalNexthopTimeout = Objects.requireNonNull(criticalNexthopTimeout);
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
        public Builder nonCriticalNexthopTimeout(Integer nonCriticalNexthopTimeout) {
            this.nonCriticalNexthopTimeout = Objects.requireNonNull(nonCriticalNexthopTimeout);
            return this;
        }
        @CustomType.Setter
        public Builder vrf(String vrf) {
            this.vrf = Objects.requireNonNull(vrf);
            return this;
        }
        public GetBgpAddressFamilyResult build() {
            final var o = new GetBgpAddressFamilyResult();
            o.addressFamily = addressFamily;
            o.asn = asn;
            o.criticalNexthopTimeout = criticalNexthopTimeout;
            o.device = device;
            o.id = id;
            o.nonCriticalNexthopTimeout = nonCriticalNexthopTimeout;
            o.vrf = vrf;
            return o;
        }
    }
}