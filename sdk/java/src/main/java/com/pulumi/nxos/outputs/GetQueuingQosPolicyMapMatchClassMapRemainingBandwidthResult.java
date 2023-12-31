// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetQueuingQosPolicyMapMatchClassMapRemainingBandwidthResult {
    /**
     * @return Class map name.
     * 
     */
    private String classMapName;
    /**
     * @return A device name from the provider configuration.
     * 
     */
    private @Nullable String device;
    /**
     * @return The distinguished name of the object.
     * 
     */
    private String id;
    /**
     * @return Policy map name.
     * 
     */
    private String policyMapName;
    /**
     * @return Remaining bandwidth percent.
     * 
     */
    private Integer value;

    private GetQueuingQosPolicyMapMatchClassMapRemainingBandwidthResult() {}
    /**
     * @return Class map name.
     * 
     */
    public String classMapName() {
        return this.classMapName;
    }
    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    /**
     * @return The distinguished name of the object.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Policy map name.
     * 
     */
    public String policyMapName() {
        return this.policyMapName;
    }
    /**
     * @return Remaining bandwidth percent.
     * 
     */
    public Integer value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetQueuingQosPolicyMapMatchClassMapRemainingBandwidthResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String classMapName;
        private @Nullable String device;
        private String id;
        private String policyMapName;
        private Integer value;
        public Builder() {}
        public Builder(GetQueuingQosPolicyMapMatchClassMapRemainingBandwidthResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.classMapName = defaults.classMapName;
    	      this.device = defaults.device;
    	      this.id = defaults.id;
    	      this.policyMapName = defaults.policyMapName;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder classMapName(String classMapName) {
            this.classMapName = Objects.requireNonNull(classMapName);
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
        public Builder policyMapName(String policyMapName) {
            this.policyMapName = Objects.requireNonNull(policyMapName);
            return this;
        }
        @CustomType.Setter
        public Builder value(Integer value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }
        public GetQueuingQosPolicyMapMatchClassMapRemainingBandwidthResult build() {
            final var o = new GetQueuingQosPolicyMapMatchClassMapRemainingBandwidthResult();
            o.classMapName = classMapName;
            o.device = device;
            o.id = id;
            o.policyMapName = policyMapName;
            o.value = value;
            return o;
        }
    }
}
