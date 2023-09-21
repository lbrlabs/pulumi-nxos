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
public final class GetPhysicalInterfaceResult {
    private String accessVlan;
    private String adminState;
    private String autoNegotiation;
    private Integer bandwidth;
    private Integer delay;
    private String description;
    private @Nullable String device;
    private String duplex;
    private String fecMode;
    private String id;
    private String interfaceId;
    private String layer;
    private Integer linkDebounceDown;
    private Integer linkDebounceUp;
    private String linkLogging;
    private String medium;
    private String mode;
    private Integer mtu;
    private String nativeVlan;
    private String speed;
    private String speedGroup;
    private String trunkVlans;
    private String uniDirectionalEthernet;
    private String userConfiguredFlags;

    private GetPhysicalInterfaceResult() {}
    public String accessVlan() {
        return this.accessVlan;
    }
    public String adminState() {
        return this.adminState;
    }
    public String autoNegotiation() {
        return this.autoNegotiation;
    }
    public Integer bandwidth() {
        return this.bandwidth;
    }
    public Integer delay() {
        return this.delay;
    }
    public String description() {
        return this.description;
    }
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public String duplex() {
        return this.duplex;
    }
    public String fecMode() {
        return this.fecMode;
    }
    public String id() {
        return this.id;
    }
    public String interfaceId() {
        return this.interfaceId;
    }
    public String layer() {
        return this.layer;
    }
    public Integer linkDebounceDown() {
        return this.linkDebounceDown;
    }
    public Integer linkDebounceUp() {
        return this.linkDebounceUp;
    }
    public String linkLogging() {
        return this.linkLogging;
    }
    public String medium() {
        return this.medium;
    }
    public String mode() {
        return this.mode;
    }
    public Integer mtu() {
        return this.mtu;
    }
    public String nativeVlan() {
        return this.nativeVlan;
    }
    public String speed() {
        return this.speed;
    }
    public String speedGroup() {
        return this.speedGroup;
    }
    public String trunkVlans() {
        return this.trunkVlans;
    }
    public String uniDirectionalEthernet() {
        return this.uniDirectionalEthernet;
    }
    public String userConfiguredFlags() {
        return this.userConfiguredFlags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPhysicalInterfaceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accessVlan;
        private String adminState;
        private String autoNegotiation;
        private Integer bandwidth;
        private Integer delay;
        private String description;
        private @Nullable String device;
        private String duplex;
        private String fecMode;
        private String id;
        private String interfaceId;
        private String layer;
        private Integer linkDebounceDown;
        private Integer linkDebounceUp;
        private String linkLogging;
        private String medium;
        private String mode;
        private Integer mtu;
        private String nativeVlan;
        private String speed;
        private String speedGroup;
        private String trunkVlans;
        private String uniDirectionalEthernet;
        private String userConfiguredFlags;
        public Builder() {}
        public Builder(GetPhysicalInterfaceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessVlan = defaults.accessVlan;
    	      this.adminState = defaults.adminState;
    	      this.autoNegotiation = defaults.autoNegotiation;
    	      this.bandwidth = defaults.bandwidth;
    	      this.delay = defaults.delay;
    	      this.description = defaults.description;
    	      this.device = defaults.device;
    	      this.duplex = defaults.duplex;
    	      this.fecMode = defaults.fecMode;
    	      this.id = defaults.id;
    	      this.interfaceId = defaults.interfaceId;
    	      this.layer = defaults.layer;
    	      this.linkDebounceDown = defaults.linkDebounceDown;
    	      this.linkDebounceUp = defaults.linkDebounceUp;
    	      this.linkLogging = defaults.linkLogging;
    	      this.medium = defaults.medium;
    	      this.mode = defaults.mode;
    	      this.mtu = defaults.mtu;
    	      this.nativeVlan = defaults.nativeVlan;
    	      this.speed = defaults.speed;
    	      this.speedGroup = defaults.speedGroup;
    	      this.trunkVlans = defaults.trunkVlans;
    	      this.uniDirectionalEthernet = defaults.uniDirectionalEthernet;
    	      this.userConfiguredFlags = defaults.userConfiguredFlags;
        }

        @CustomType.Setter
        public Builder accessVlan(String accessVlan) {
            this.accessVlan = Objects.requireNonNull(accessVlan);
            return this;
        }
        @CustomType.Setter
        public Builder adminState(String adminState) {
            this.adminState = Objects.requireNonNull(adminState);
            return this;
        }
        @CustomType.Setter
        public Builder autoNegotiation(String autoNegotiation) {
            this.autoNegotiation = Objects.requireNonNull(autoNegotiation);
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
        public Builder duplex(String duplex) {
            this.duplex = Objects.requireNonNull(duplex);
            return this;
        }
        @CustomType.Setter
        public Builder fecMode(String fecMode) {
            this.fecMode = Objects.requireNonNull(fecMode);
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
        public Builder layer(String layer) {
            this.layer = Objects.requireNonNull(layer);
            return this;
        }
        @CustomType.Setter
        public Builder linkDebounceDown(Integer linkDebounceDown) {
            this.linkDebounceDown = Objects.requireNonNull(linkDebounceDown);
            return this;
        }
        @CustomType.Setter
        public Builder linkDebounceUp(Integer linkDebounceUp) {
            this.linkDebounceUp = Objects.requireNonNull(linkDebounceUp);
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
        public Builder mode(String mode) {
            this.mode = Objects.requireNonNull(mode);
            return this;
        }
        @CustomType.Setter
        public Builder mtu(Integer mtu) {
            this.mtu = Objects.requireNonNull(mtu);
            return this;
        }
        @CustomType.Setter
        public Builder nativeVlan(String nativeVlan) {
            this.nativeVlan = Objects.requireNonNull(nativeVlan);
            return this;
        }
        @CustomType.Setter
        public Builder speed(String speed) {
            this.speed = Objects.requireNonNull(speed);
            return this;
        }
        @CustomType.Setter
        public Builder speedGroup(String speedGroup) {
            this.speedGroup = Objects.requireNonNull(speedGroup);
            return this;
        }
        @CustomType.Setter
        public Builder trunkVlans(String trunkVlans) {
            this.trunkVlans = Objects.requireNonNull(trunkVlans);
            return this;
        }
        @CustomType.Setter
        public Builder uniDirectionalEthernet(String uniDirectionalEthernet) {
            this.uniDirectionalEthernet = Objects.requireNonNull(uniDirectionalEthernet);
            return this;
        }
        @CustomType.Setter
        public Builder userConfiguredFlags(String userConfiguredFlags) {
            this.userConfiguredFlags = Objects.requireNonNull(userConfiguredFlags);
            return this;
        }
        public GetPhysicalInterfaceResult build() {
            final var o = new GetPhysicalInterfaceResult();
            o.accessVlan = accessVlan;
            o.adminState = adminState;
            o.autoNegotiation = autoNegotiation;
            o.bandwidth = bandwidth;
            o.delay = delay;
            o.description = description;
            o.device = device;
            o.duplex = duplex;
            o.fecMode = fecMode;
            o.id = id;
            o.interfaceId = interfaceId;
            o.layer = layer;
            o.linkDebounceDown = linkDebounceDown;
            o.linkDebounceUp = linkDebounceUp;
            o.linkLogging = linkLogging;
            o.medium = medium;
            o.mode = mode;
            o.mtu = mtu;
            o.nativeVlan = nativeVlan;
            o.speed = speed;
            o.speedGroup = speedGroup;
            o.trunkVlans = trunkVlans;
            o.uniDirectionalEthernet = uniDirectionalEthernet;
            o.userConfiguredFlags = userConfiguredFlags;
            return o;
        }
    }
}
