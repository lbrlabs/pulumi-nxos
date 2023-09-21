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
public final class GetVpcInterfaceResult {
    private @Nullable String device;
    private String id;
    private String portChannelInterfaceDn;
    private Integer vpcInterfaceId;

    private GetVpcInterfaceResult() {}
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public String id() {
        return this.id;
    }
    public String portChannelInterfaceDn() {
        return this.portChannelInterfaceDn;
    }
    public Integer vpcInterfaceId() {
        return this.vpcInterfaceId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVpcInterfaceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String device;
        private String id;
        private String portChannelInterfaceDn;
        private Integer vpcInterfaceId;
        public Builder() {}
        public Builder(GetVpcInterfaceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.device = defaults.device;
    	      this.id = defaults.id;
    	      this.portChannelInterfaceDn = defaults.portChannelInterfaceDn;
    	      this.vpcInterfaceId = defaults.vpcInterfaceId;
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
        public Builder portChannelInterfaceDn(String portChannelInterfaceDn) {
            this.portChannelInterfaceDn = Objects.requireNonNull(portChannelInterfaceDn);
            return this;
        }
        @CustomType.Setter
        public Builder vpcInterfaceId(Integer vpcInterfaceId) {
            this.vpcInterfaceId = Objects.requireNonNull(vpcInterfaceId);
            return this;
        }
        public GetVpcInterfaceResult build() {
            final var o = new GetVpcInterfaceResult();
            o.device = device;
            o.id = id;
            o.portChannelInterfaceDn = portChannelInterfaceDn;
            o.vpcInterfaceId = vpcInterfaceId;
            return o;
        }
    }
}