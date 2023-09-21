// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetFeatureMacsecResult {
    private String adminState;
    private @Nullable String device;
    private String id;

    private GetFeatureMacsecResult() {}
    public String adminState() {
        return this.adminState;
    }
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public String id() {
        return this.id;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetFeatureMacsecResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String adminState;
        private @Nullable String device;
        private String id;
        public Builder() {}
        public Builder(GetFeatureMacsecResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adminState = defaults.adminState;
    	      this.device = defaults.device;
    	      this.id = defaults.id;
        }

        @CustomType.Setter
        public Builder adminState(String adminState) {
            this.adminState = Objects.requireNonNull(adminState);
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
        public GetFeatureMacsecResult build() {
            final var o = new GetFeatureMacsecResult();
            o.adminState = adminState;
            o.device = device;
            o.id = id;
            return o;
        }
    }
}
