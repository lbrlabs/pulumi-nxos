// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVrfResult {
    /**
     * @return VRF description.
     * 
     */
    private String description;
    /**
     * @return A device name from the provider configuration.
     * 
     */
    private @Nullable String device;
    /**
     * @return Encap for this Context, supported formats: `unknown`, `vlan-%d` or `vxlan-%d`.
     * 
     */
    private String encap;
    /**
     * @return The distinguished name of the object.
     * 
     */
    private String id;
    /**
     * @return VRF name.
     * 
     */
    private String name;

    private GetVrfResult() {}
    /**
     * @return VRF description.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    /**
     * @return Encap for this Context, supported formats: `unknown`, `vlan-%d` or `vxlan-%d`.
     * 
     */
    public String encap() {
        return this.encap;
    }
    /**
     * @return The distinguished name of the object.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return VRF name.
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVrfResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private @Nullable String device;
        private String encap;
        private String id;
        private String name;
        public Builder() {}
        public Builder(GetVrfResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.device = defaults.device;
    	      this.encap = defaults.encap;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        @CustomType.Setter
        public Builder encap(String encap) {
            this.encap = Objects.requireNonNull(encap);
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
        public GetVrfResult build() {
            final var o = new GetVrfResult();
            o.description = description;
            o.device = device;
            o.encap = encap;
            o.id = id;
            o.name = name;
            return o;
        }
    }
}