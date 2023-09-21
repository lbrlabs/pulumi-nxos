// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetPimAnycastRpPeerResult {
    private String address;
    private @Nullable String device;
    private String id;
    private String rpSetAddress;
    private String vrfName;

    private GetPimAnycastRpPeerResult() {}
    public String address() {
        return this.address;
    }
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public String id() {
        return this.id;
    }
    public String rpSetAddress() {
        return this.rpSetAddress;
    }
    public String vrfName() {
        return this.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPimAnycastRpPeerResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String address;
        private @Nullable String device;
        private String id;
        private String rpSetAddress;
        private String vrfName;
        public Builder() {}
        public Builder(GetPimAnycastRpPeerResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.address = defaults.address;
    	      this.device = defaults.device;
    	      this.id = defaults.id;
    	      this.rpSetAddress = defaults.rpSetAddress;
    	      this.vrfName = defaults.vrfName;
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
        public Builder rpSetAddress(String rpSetAddress) {
            this.rpSetAddress = Objects.requireNonNull(rpSetAddress);
            return this;
        }
        @CustomType.Setter
        public Builder vrfName(String vrfName) {
            this.vrfName = Objects.requireNonNull(vrfName);
            return this;
        }
        public GetPimAnycastRpPeerResult build() {
            final var o = new GetPimAnycastRpPeerResult();
            o.address = address;
            o.device = device;
            o.id = id;
            o.rpSetAddress = rpSetAddress;
            o.vrfName = vrfName;
            return o;
        }
    }
}
