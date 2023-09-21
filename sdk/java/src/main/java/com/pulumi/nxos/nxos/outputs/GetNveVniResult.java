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
public final class GetNveVniResult {
    private Boolean associateVrf;
    private @Nullable String device;
    private String id;
    private String multicastGroup;
    private String multisiteIngressReplication;
    private String suppressArp;
    private Integer vni;

    private GetNveVniResult() {}
    public Boolean associateVrf() {
        return this.associateVrf;
    }
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public String id() {
        return this.id;
    }
    public String multicastGroup() {
        return this.multicastGroup;
    }
    public String multisiteIngressReplication() {
        return this.multisiteIngressReplication;
    }
    public String suppressArp() {
        return this.suppressArp;
    }
    public Integer vni() {
        return this.vni;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetNveVniResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean associateVrf;
        private @Nullable String device;
        private String id;
        private String multicastGroup;
        private String multisiteIngressReplication;
        private String suppressArp;
        private Integer vni;
        public Builder() {}
        public Builder(GetNveVniResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.associateVrf = defaults.associateVrf;
    	      this.device = defaults.device;
    	      this.id = defaults.id;
    	      this.multicastGroup = defaults.multicastGroup;
    	      this.multisiteIngressReplication = defaults.multisiteIngressReplication;
    	      this.suppressArp = defaults.suppressArp;
    	      this.vni = defaults.vni;
        }

        @CustomType.Setter
        public Builder associateVrf(Boolean associateVrf) {
            this.associateVrf = Objects.requireNonNull(associateVrf);
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
        public Builder multicastGroup(String multicastGroup) {
            this.multicastGroup = Objects.requireNonNull(multicastGroup);
            return this;
        }
        @CustomType.Setter
        public Builder multisiteIngressReplication(String multisiteIngressReplication) {
            this.multisiteIngressReplication = Objects.requireNonNull(multisiteIngressReplication);
            return this;
        }
        @CustomType.Setter
        public Builder suppressArp(String suppressArp) {
            this.suppressArp = Objects.requireNonNull(suppressArp);
            return this;
        }
        @CustomType.Setter
        public Builder vni(Integer vni) {
            this.vni = Objects.requireNonNull(vni);
            return this;
        }
        public GetNveVniResult build() {
            final var o = new GetNveVniResult();
            o.associateVrf = associateVrf;
            o.device = device;
            o.id = id;
            o.multicastGroup = multicastGroup;
            o.multisiteIngressReplication = multisiteIngressReplication;
            o.suppressArp = suppressArp;
            o.vni = vni;
            return o;
        }
    }
}