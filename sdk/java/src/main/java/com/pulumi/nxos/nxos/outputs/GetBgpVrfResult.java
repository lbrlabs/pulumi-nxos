// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBgpVrfResult {
    private String asn;
    private @Nullable String device;
    private String id;
    private String name;
    private String routerId;

    private GetBgpVrfResult() {}
    public String asn() {
        return this.asn;
    }
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    public String routerId() {
        return this.routerId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBgpVrfResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String asn;
        private @Nullable String device;
        private String id;
        private String name;
        private String routerId;
        public Builder() {}
        public Builder(GetBgpVrfResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.asn = defaults.asn;
    	      this.device = defaults.device;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.routerId = defaults.routerId;
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
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder routerId(String routerId) {
            this.routerId = Objects.requireNonNull(routerId);
            return this;
        }
        public GetBgpVrfResult build() {
            final var o = new GetBgpVrfResult();
            o.asn = asn;
            o.device = device;
            o.id = id;
            o.name = name;
            o.routerId = routerId;
            return o;
        }
    }
}