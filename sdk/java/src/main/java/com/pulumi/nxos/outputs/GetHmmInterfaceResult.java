// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetHmmInterfaceResult {
    /**
     * @return Administrative state.
     * 
     */
    private String adminState;
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
     * @return Must match first field in the output of `show intf brief`. Example: `vlan10`.
     * 
     */
    private String interfaceId;
    /**
     * @return HMM Fabric Forwarding mode information for the interface.
     * 
     */
    private String mode;

    private GetHmmInterfaceResult() {}
    /**
     * @return Administrative state.
     * 
     */
    public String adminState() {
        return this.adminState;
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
     * @return Must match first field in the output of `show intf brief`. Example: `vlan10`.
     * 
     */
    public String interfaceId() {
        return this.interfaceId;
    }
    /**
     * @return HMM Fabric Forwarding mode information for the interface.
     * 
     */
    public String mode() {
        return this.mode;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetHmmInterfaceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String adminState;
        private @Nullable String device;
        private String id;
        private String interfaceId;
        private String mode;
        public Builder() {}
        public Builder(GetHmmInterfaceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adminState = defaults.adminState;
    	      this.device = defaults.device;
    	      this.id = defaults.id;
    	      this.interfaceId = defaults.interfaceId;
    	      this.mode = defaults.mode;
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
        @CustomType.Setter
        public Builder interfaceId(String interfaceId) {
            this.interfaceId = Objects.requireNonNull(interfaceId);
            return this;
        }
        @CustomType.Setter
        public Builder mode(String mode) {
            this.mode = Objects.requireNonNull(mode);
            return this;
        }
        public GetHmmInterfaceResult build() {
            final var o = new GetHmmInterfaceResult();
            o.adminState = adminState;
            o.device = device;
            o.id = id;
            o.interfaceId = interfaceId;
            o.mode = mode;
            return o;
        }
    }
}
