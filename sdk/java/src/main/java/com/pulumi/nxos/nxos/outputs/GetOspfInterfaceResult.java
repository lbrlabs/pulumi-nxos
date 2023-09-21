// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetOspfInterfaceResult {
    private Boolean advertiseSecondaries;
    private String area;
    private String bfd;
    private Integer cost;
    private Integer deadInterval;
    private @Nullable String device;
    private Integer helloInterval;
    private String id;
    private String instanceName;
    private String interfaceId;
    private String networkType;
    private String passive;
    private Integer priority;
    private String vrfName;

    private GetOspfInterfaceResult() {}
    public Boolean advertiseSecondaries() {
        return this.advertiseSecondaries;
    }
    public String area() {
        return this.area;
    }
    public String bfd() {
        return this.bfd;
    }
    public Integer cost() {
        return this.cost;
    }
    public Integer deadInterval() {
        return this.deadInterval;
    }
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public Integer helloInterval() {
        return this.helloInterval;
    }
    public String id() {
        return this.id;
    }
    public String instanceName() {
        return this.instanceName;
    }
    public String interfaceId() {
        return this.interfaceId;
    }
    public String networkType() {
        return this.networkType;
    }
    public String passive() {
        return this.passive;
    }
    public Integer priority() {
        return this.priority;
    }
    public String vrfName() {
        return this.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOspfInterfaceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean advertiseSecondaries;
        private String area;
        private String bfd;
        private Integer cost;
        private Integer deadInterval;
        private @Nullable String device;
        private Integer helloInterval;
        private String id;
        private String instanceName;
        private String interfaceId;
        private String networkType;
        private String passive;
        private Integer priority;
        private String vrfName;
        public Builder() {}
        public Builder(GetOspfInterfaceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.advertiseSecondaries = defaults.advertiseSecondaries;
    	      this.area = defaults.area;
    	      this.bfd = defaults.bfd;
    	      this.cost = defaults.cost;
    	      this.deadInterval = defaults.deadInterval;
    	      this.device = defaults.device;
    	      this.helloInterval = defaults.helloInterval;
    	      this.id = defaults.id;
    	      this.instanceName = defaults.instanceName;
    	      this.interfaceId = defaults.interfaceId;
    	      this.networkType = defaults.networkType;
    	      this.passive = defaults.passive;
    	      this.priority = defaults.priority;
    	      this.vrfName = defaults.vrfName;
        }

        @CustomType.Setter
        public Builder advertiseSecondaries(Boolean advertiseSecondaries) {
            this.advertiseSecondaries = Objects.requireNonNull(advertiseSecondaries);
            return this;
        }
        @CustomType.Setter
        public Builder area(String area) {
            this.area = Objects.requireNonNull(area);
            return this;
        }
        @CustomType.Setter
        public Builder bfd(String bfd) {
            this.bfd = Objects.requireNonNull(bfd);
            return this;
        }
        @CustomType.Setter
        public Builder cost(Integer cost) {
            this.cost = Objects.requireNonNull(cost);
            return this;
        }
        @CustomType.Setter
        public Builder deadInterval(Integer deadInterval) {
            this.deadInterval = Objects.requireNonNull(deadInterval);
            return this;
        }
        @CustomType.Setter
        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        @CustomType.Setter
        public Builder helloInterval(Integer helloInterval) {
            this.helloInterval = Objects.requireNonNull(helloInterval);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder instanceName(String instanceName) {
            this.instanceName = Objects.requireNonNull(instanceName);
            return this;
        }
        @CustomType.Setter
        public Builder interfaceId(String interfaceId) {
            this.interfaceId = Objects.requireNonNull(interfaceId);
            return this;
        }
        @CustomType.Setter
        public Builder networkType(String networkType) {
            this.networkType = Objects.requireNonNull(networkType);
            return this;
        }
        @CustomType.Setter
        public Builder passive(String passive) {
            this.passive = Objects.requireNonNull(passive);
            return this;
        }
        @CustomType.Setter
        public Builder priority(Integer priority) {
            this.priority = Objects.requireNonNull(priority);
            return this;
        }
        @CustomType.Setter
        public Builder vrfName(String vrfName) {
            this.vrfName = Objects.requireNonNull(vrfName);
            return this;
        }
        public GetOspfInterfaceResult build() {
            final var o = new GetOspfInterfaceResult();
            o.advertiseSecondaries = advertiseSecondaries;
            o.area = area;
            o.bfd = bfd;
            o.cost = cost;
            o.deadInterval = deadInterval;
            o.device = device;
            o.helloInterval = helloInterval;
            o.id = id;
            o.instanceName = instanceName;
            o.interfaceId = interfaceId;
            o.networkType = networkType;
            o.passive = passive;
            o.priority = priority;
            o.vrfName = vrfName;
            return o;
        }
    }
}