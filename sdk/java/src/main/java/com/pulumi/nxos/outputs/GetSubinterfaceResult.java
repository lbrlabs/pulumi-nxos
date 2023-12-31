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
public final class GetSubinterfaceResult {
    /**
     * @return Administrative state.
     * 
     */
    private String adminState;
    /**
     * @return Specifies the administrative port bandwidth.
     * 
     */
    private Integer bandwidth;
    /**
     * @return Specifies the administrative port delay.
     * 
     */
    private Integer delay;
    /**
     * @return Interface description.
     * 
     */
    private String description;
    /**
     * @return A device name from the provider configuration.
     * 
     */
    private @Nullable String device;
    /**
     * @return Subinterface encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
     * 
     */
    private String encap;
    /**
     * @return The distinguished name of the object.
     * 
     */
    private String id;
    /**
     * @return Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
     * 
     */
    private String interfaceId;
    /**
     * @return Administrative link logging.
     * 
     */
    private String linkLogging;
    /**
     * @return The administrative port medium type.
     * 
     */
    private String medium;
    /**
     * @return Administrative port MTU.
     * 
     */
    private Integer mtu;

    private GetSubinterfaceResult() {}
    /**
     * @return Administrative state.
     * 
     */
    public String adminState() {
        return this.adminState;
    }
    /**
     * @return Specifies the administrative port bandwidth.
     * 
     */
    public Integer bandwidth() {
        return this.bandwidth;
    }
    /**
     * @return Specifies the administrative port delay.
     * 
     */
    public Integer delay() {
        return this.delay;
    }
    /**
     * @return Interface description.
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
     * @return Subinterface encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
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
     * @return Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
     * 
     */
    public String interfaceId() {
        return this.interfaceId;
    }
    /**
     * @return Administrative link logging.
     * 
     */
    public String linkLogging() {
        return this.linkLogging;
    }
    /**
     * @return The administrative port medium type.
     * 
     */
    public String medium() {
        return this.medium;
    }
    /**
     * @return Administrative port MTU.
     * 
     */
    public Integer mtu() {
        return this.mtu;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSubinterfaceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String adminState;
        private Integer bandwidth;
        private Integer delay;
        private String description;
        private @Nullable String device;
        private String encap;
        private String id;
        private String interfaceId;
        private String linkLogging;
        private String medium;
        private Integer mtu;
        public Builder() {}
        public Builder(GetSubinterfaceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adminState = defaults.adminState;
    	      this.bandwidth = defaults.bandwidth;
    	      this.delay = defaults.delay;
    	      this.description = defaults.description;
    	      this.device = defaults.device;
    	      this.encap = defaults.encap;
    	      this.id = defaults.id;
    	      this.interfaceId = defaults.interfaceId;
    	      this.linkLogging = defaults.linkLogging;
    	      this.medium = defaults.medium;
    	      this.mtu = defaults.mtu;
        }

        @CustomType.Setter
        public Builder adminState(String adminState) {
            this.adminState = Objects.requireNonNull(adminState);
            return this;
        }
        @CustomType.Setter
        public Builder bandwidth(Integer bandwidth) {
            this.bandwidth = Objects.requireNonNull(bandwidth);
            return this;
        }
        @CustomType.Setter
        public Builder delay(Integer delay) {
            this.delay = Objects.requireNonNull(delay);
            return this;
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
        public Builder interfaceId(String interfaceId) {
            this.interfaceId = Objects.requireNonNull(interfaceId);
            return this;
        }
        @CustomType.Setter
        public Builder linkLogging(String linkLogging) {
            this.linkLogging = Objects.requireNonNull(linkLogging);
            return this;
        }
        @CustomType.Setter
        public Builder medium(String medium) {
            this.medium = Objects.requireNonNull(medium);
            return this;
        }
        @CustomType.Setter
        public Builder mtu(Integer mtu) {
            this.mtu = Objects.requireNonNull(mtu);
            return this;
        }
        public GetSubinterfaceResult build() {
            final var o = new GetSubinterfaceResult();
            o.adminState = adminState;
            o.bandwidth = bandwidth;
            o.delay = delay;
            o.description = description;
            o.device = device;
            o.encap = encap;
            o.id = id;
            o.interfaceId = interfaceId;
            o.linkLogging = linkLogging;
            o.medium = medium;
            o.mtu = mtu;
            return o;
        }
    }
}
