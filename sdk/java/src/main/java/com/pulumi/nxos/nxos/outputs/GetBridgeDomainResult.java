// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBridgeDomainResult {
    private String accessEncap;
    private @Nullable String device;
    private String fabricEncap;
    private String id;
    private String name;

    private GetBridgeDomainResult() {}
    public String accessEncap() {
        return this.accessEncap;
    }
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public String fabricEncap() {
        return this.fabricEncap;
    }
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBridgeDomainResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessEncap;
        private @Nullable String device;
        private String fabricEncap;
        private String id;
        private String name;
        public Builder() {}
        public Builder(GetBridgeDomainResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessEncap = defaults.accessEncap;
    	      this.device = defaults.device;
    	      this.fabricEncap = defaults.fabricEncap;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder accessEncap(String accessEncap) {
            this.accessEncap = Objects.requireNonNull(accessEncap);
            return this;
        }
        @CustomType.Setter
        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        @CustomType.Setter
        public Builder fabricEncap(String fabricEncap) {
            this.fabricEncap = Objects.requireNonNull(fabricEncap);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public GetBridgeDomainResult build() {
            final var o = new GetBridgeDomainResult();
            o.accessEncap = accessEncap;
            o.device = device;
            o.fabricEncap = fabricEncap;
            o.id = id;
            o.name = name;
            return o;
        }
    }
}
