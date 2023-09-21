// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIpv4StaticRouteNextHop {
    private String address;
    private String description;
    private String interfaceId;
    private Integer object;
    private Integer preference;
    private Integer tag;
    private String vrfName;

    private GetIpv4StaticRouteNextHop() {}
    public String address() {
        return this.address;
    }
    public String description() {
        return this.description;
    }
    public String interfaceId() {
        return this.interfaceId;
    }
    public Integer object() {
        return this.object;
    }
    public Integer preference() {
        return this.preference;
    }
    public Integer tag() {
        return this.tag;
    }
    public String vrfName() {
        return this.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpv4StaticRouteNextHop defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String address;
        private String description;
        private String interfaceId;
        private Integer object;
        private Integer preference;
        private Integer tag;
        private String vrfName;
        public Builder() {}
        public Builder(GetIpv4StaticRouteNextHop defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.address = defaults.address;
    	      this.description = defaults.description;
    	      this.interfaceId = defaults.interfaceId;
    	      this.object = defaults.object;
    	      this.preference = defaults.preference;
    	      this.tag = defaults.tag;
    	      this.vrfName = defaults.vrfName;
        }

        @CustomType.Setter
        public Builder address(String address) {
            this.address = Objects.requireNonNull(address);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder interfaceId(String interfaceId) {
            this.interfaceId = Objects.requireNonNull(interfaceId);
            return this;
        }
        @CustomType.Setter
        public Builder object(Integer object) {
            this.object = Objects.requireNonNull(object);
            return this;
        }
        @CustomType.Setter
        public Builder preference(Integer preference) {
            this.preference = Objects.requireNonNull(preference);
            return this;
        }
        @CustomType.Setter
        public Builder tag(Integer tag) {
            this.tag = Objects.requireNonNull(tag);
            return this;
        }
        @CustomType.Setter
        public Builder vrfName(String vrfName) {
            this.vrfName = Objects.requireNonNull(vrfName);
            return this;
        }
        public GetIpv4StaticRouteNextHop build() {
            final var o = new GetIpv4StaticRouteNextHop();
            o.address = address;
            o.description = description;
            o.interfaceId = interfaceId;
            o.object = object;
            o.preference = preference;
            o.tag = tag;
            o.vrfName = vrfName;
            return o;
        }
    }
}
