// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetDefaultQosClassMapResult {
    private @Nullable String device;
    private String id;
    private String matchType;
    private String name;

    private GetDefaultQosClassMapResult() {}
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public String id() {
        return this.id;
    }
    public String matchType() {
        return this.matchType;
    }
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDefaultQosClassMapResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String device;
        private String id;
        private String matchType;
        private String name;
        public Builder() {}
        public Builder(GetDefaultQosClassMapResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.device = defaults.device;
    	      this.id = defaults.id;
    	      this.matchType = defaults.matchType;
    	      this.name = defaults.name;
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
        public Builder matchType(String matchType) {
            this.matchType = Objects.requireNonNull(matchType);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public GetDefaultQosClassMapResult build() {
            final var o = new GetDefaultQosClassMapResult();
            o.device = device;
            o.id = id;
            o.matchType = matchType;
            o.name = name;
            return o;
        }
    }
}
