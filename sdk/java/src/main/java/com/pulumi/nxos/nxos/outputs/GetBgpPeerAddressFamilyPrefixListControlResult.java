// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBgpPeerAddressFamilyPrefixListControlResult {
    private String address;
    private String addressFamily;
    private String asn;
    private @Nullable String device;
    private String direction;
    private String id;
    private String list;
    private String vrf;

    private GetBgpPeerAddressFamilyPrefixListControlResult() {}
    public String address() {
        return this.address;
    }
    public String addressFamily() {
        return this.addressFamily;
    }
    public String asn() {
        return this.asn;
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
    public String list() {
        return this.list;
    }
    public String vrf() {
        return this.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBgpPeerAddressFamilyPrefixListControlResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String address;
        private String addressFamily;
        private String asn;
        private @Nullable String device;
        private String direction;
        private String id;
        private String list;
        private String vrf;
        public Builder() {}
        public Builder(GetBgpPeerAddressFamilyPrefixListControlResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.address = defaults.address;
    	      this.addressFamily = defaults.addressFamily;
    	      this.asn = defaults.asn;
    	      this.device = defaults.device;
    	      this.direction = defaults.direction;
    	      this.id = defaults.id;
    	      this.list = defaults.list;
    	      this.vrf = defaults.vrf;
        }

        @CustomType.Setter
        public Builder address(String address) {
            this.address = Objects.requireNonNull(address);
            return this;
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
        public Builder list(String list) {
            this.list = Objects.requireNonNull(list);
            return this;
        }
        @CustomType.Setter
        public Builder vrf(String vrf) {
            this.vrf = Objects.requireNonNull(vrf);
            return this;
        }
        public GetBgpPeerAddressFamilyPrefixListControlResult build() {
            final var o = new GetBgpPeerAddressFamilyPrefixListControlResult();
            o.address = address;
            o.addressFamily = addressFamily;
            o.asn = asn;
            o.device = device;
            o.direction = direction;
            o.id = id;
            o.list = list;
            o.vrf = vrf;
            return o;
        }
    }
}
