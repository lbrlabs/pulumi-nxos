// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetSpanningTreeInterfaceResult {
    private String adminState;
    private String bpduFilter;
    private String bpduGuard;
    private Integer cost;
    private @Nullable String device;
    private String guard;
    private String id;
    private String interfaceId;
    private String linkType;
    private String mode;
    private Integer priority;

    private GetSpanningTreeInterfaceResult() {}
    public String adminState() {
        return this.adminState;
    }
    public String bpduFilter() {
        return this.bpduFilter;
    }
    public String bpduGuard() {
        return this.bpduGuard;
    }
    public Integer cost() {
        return this.cost;
    }
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public String guard() {
        return this.guard;
    }
    public String id() {
        return this.id;
    }
    public String interfaceId() {
        return this.interfaceId;
    }
    public String linkType() {
        return this.linkType;
    }
    public String mode() {
        return this.mode;
    }
    public Integer priority() {
        return this.priority;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSpanningTreeInterfaceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String adminState;
        private String bpduFilter;
        private String bpduGuard;
        private Integer cost;
        private @Nullable String device;
        private String guard;
        private String id;
        private String interfaceId;
        private String linkType;
        private String mode;
        private Integer priority;
        public Builder() {}
        public Builder(GetSpanningTreeInterfaceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adminState = defaults.adminState;
    	      this.bpduFilter = defaults.bpduFilter;
    	      this.bpduGuard = defaults.bpduGuard;
    	      this.cost = defaults.cost;
    	      this.device = defaults.device;
    	      this.guard = defaults.guard;
    	      this.id = defaults.id;
    	      this.interfaceId = defaults.interfaceId;
    	      this.linkType = defaults.linkType;
    	      this.mode = defaults.mode;
    	      this.priority = defaults.priority;
        }

        @CustomType.Setter
        public Builder adminState(String adminState) {
            this.adminState = Objects.requireNonNull(adminState);
            return this;
        }
        @CustomType.Setter
        public Builder bpduFilter(String bpduFilter) {
            this.bpduFilter = Objects.requireNonNull(bpduFilter);
            return this;
        }
        @CustomType.Setter
        public Builder bpduGuard(String bpduGuard) {
            this.bpduGuard = Objects.requireNonNull(bpduGuard);
            return this;
        }
        @CustomType.Setter
        public Builder cost(Integer cost) {
            this.cost = Objects.requireNonNull(cost);
            return this;
        }
        @CustomType.Setter
        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        @CustomType.Setter
        public Builder guard(String guard) {
            this.guard = Objects.requireNonNull(guard);
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
        public Builder linkType(String linkType) {
            this.linkType = Objects.requireNonNull(linkType);
            return this;
        }
        @CustomType.Setter
        public Builder mode(String mode) {
            this.mode = Objects.requireNonNull(mode);
            return this;
        }
        @CustomType.Setter
        public Builder priority(Integer priority) {
            this.priority = Objects.requireNonNull(priority);
            return this;
        }
        public GetSpanningTreeInterfaceResult build() {
            final var o = new GetSpanningTreeInterfaceResult();
            o.adminState = adminState;
            o.bpduFilter = bpduFilter;
            o.bpduGuard = bpduGuard;
            o.cost = cost;
            o.device = device;
            o.guard = guard;
            o.id = id;
            o.interfaceId = interfaceId;
            o.linkType = linkType;
            o.mode = mode;
            o.priority = priority;
            return o;
        }
    }
}
