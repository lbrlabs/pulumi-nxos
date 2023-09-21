// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDhcpRelayAddressResult {
    private String address;
    private @Nullable String device;
    private String id;
    private String interfaceId;
    private String vrf;

    private GetDhcpRelayAddressResult() {}
    public String address() {
        return this.address;
    }
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public String id() {
        return this.id;
    }
    public String interfaceId() {
        return this.interfaceId;
    }
    public String vrf() {
        return this.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDhcpRelayAddressResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String address;
        private @Nullable String device;
        private String id;
        private String interfaceId;
        private String vrf;
        public Builder() {}
        public Builder(GetDhcpRelayAddressResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.address = defaults.address;
    	      this.device = defaults.device;
    	      this.id = defaults.id;
    	      this.interfaceId = defaults.interfaceId;
    	      this.vrf = defaults.vrf;
        }

        @CustomType.Setter
        public Builder address(String address) {
            this.address = Objects.requireNonNull(address);
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
        public Builder interfaceId(String interfaceId) {
            this.interfaceId = Objects.requireNonNull(interfaceId);
            return this;
        }
        @CustomType.Setter
        public Builder vrf(String vrf) {
            this.vrf = Objects.requireNonNull(vrf);
            return this;
        }
        public GetDhcpRelayAddressResult build() {
            final var o = new GetDhcpRelayAddressResult();
            o.address = address;
            o.device = device;
            o.id = id;
            o.interfaceId = interfaceId;
            o.vrf = vrf;
            return o;
        }
    }
}
