// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetIpv4InterfaceResult {
    /**
     * @return A device name from the provider configuration.
     * 
     */
    private @Nullable String device;
    /**
     * @return ip drop-glean enabled/disabled.
     * 
     */
    private String dropGlean;
    /**
     * @return ip forward enabled/disabled.
     * 
     */
    private String forward;
    /**
     * @return The distinguished name of the object.
     * 
     */
    private String id;
    /**
     * @return Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     * 
     */
    private String interfaceId;
    /**
     * @return IP unnumbered. Reference to interface must match first field in the output of `show intf brief`. Example: `eth1/1`.
     * 
     */
    private String unnumbered;
    /**
     * @return URPF (unicast Reverse Path Forwarding).
     * 
     */
    private String urpf;
    /**
     * @return VRF name.
     * 
     */
    private String vrf;

    private GetIpv4InterfaceResult() {}
    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    /**
     * @return ip drop-glean enabled/disabled.
     * 
     */
    public String dropGlean() {
        return this.dropGlean;
    }
    /**
     * @return ip forward enabled/disabled.
     * 
     */
    public String forward() {
        return this.forward;
    }
    /**
     * @return The distinguished name of the object.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     * 
     */
    public String interfaceId() {
        return this.interfaceId;
    }
    /**
     * @return IP unnumbered. Reference to interface must match first field in the output of `show intf brief`. Example: `eth1/1`.
     * 
     */
    public String unnumbered() {
        return this.unnumbered;
    }
    /**
     * @return URPF (unicast Reverse Path Forwarding).
     * 
     */
    public String urpf() {
        return this.urpf;
    }
    /**
     * @return VRF name.
     * 
     */
    public String vrf() {
        return this.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpv4InterfaceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String device;
        private String dropGlean;
        private String forward;
        private String id;
        private String interfaceId;
        private String unnumbered;
        private String urpf;
        private String vrf;
        public Builder() {}
        public Builder(GetIpv4InterfaceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.device = defaults.device;
    	      this.dropGlean = defaults.dropGlean;
    	      this.forward = defaults.forward;
    	      this.id = defaults.id;
    	      this.interfaceId = defaults.interfaceId;
    	      this.unnumbered = defaults.unnumbered;
    	      this.urpf = defaults.urpf;
    	      this.vrf = defaults.vrf;
        }

        @CustomType.Setter
        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        @CustomType.Setter
        public Builder dropGlean(String dropGlean) {
            this.dropGlean = Objects.requireNonNull(dropGlean);
            return this;
        }
        @CustomType.Setter
        public Builder forward(String forward) {
            this.forward = Objects.requireNonNull(forward);
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
        public Builder unnumbered(String unnumbered) {
            this.unnumbered = Objects.requireNonNull(unnumbered);
            return this;
        }
        @CustomType.Setter
        public Builder urpf(String urpf) {
            this.urpf = Objects.requireNonNull(urpf);
            return this;
        }
        @CustomType.Setter
        public Builder vrf(String vrf) {
            this.vrf = Objects.requireNonNull(vrf);
            return this;
        }
        public GetIpv4InterfaceResult build() {
            final var o = new GetIpv4InterfaceResult();
            o.device = device;
            o.dropGlean = dropGlean;
            o.forward = forward;
            o.id = id;
            o.interfaceId = interfaceId;
            o.unnumbered = unnumbered;
            o.urpf = urpf;
            o.vrf = vrf;
            return o;
        }
    }
}
