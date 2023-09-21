// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetIpv4AccessListPolicyIngressInterfaceResult {
    private String accessListName;
    private @Nullable String device;
    private String id;
    private String interfaceId;

    private GetIpv4AccessListPolicyIngressInterfaceResult() {}
    public String accessListName() {
        return this.accessListName;
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

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpv4AccessListPolicyIngressInterfaceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessListName;
        private @Nullable String device;
        private String id;
        private String interfaceId;
        public Builder() {}
        public Builder(GetIpv4AccessListPolicyIngressInterfaceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessListName = defaults.accessListName;
    	      this.device = defaults.device;
    	      this.id = defaults.id;
    	      this.interfaceId = defaults.interfaceId;
        }

        @CustomType.Setter
        public Builder accessListName(String accessListName) {
            this.accessListName = Objects.requireNonNull(accessListName);
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
        public GetIpv4AccessListPolicyIngressInterfaceResult build() {
            final var o = new GetIpv4AccessListPolicyIngressInterfaceResult();
            o.accessListName = accessListName;
            o.device = device;
            o.id = id;
            o.interfaceId = interfaceId;
            return o;
        }
    }
}
