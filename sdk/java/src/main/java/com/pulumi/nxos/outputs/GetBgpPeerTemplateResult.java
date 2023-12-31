// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBgpPeerTemplateResult {
    /**
     * @return Autonomous system number.
     * 
     */
    private String asn;
    /**
     * @return Peer template description.
     * 
     */
    private String description;
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
     * @return Neighbor Fabric Type.
     * 
     */
    private String peerType;
    /**
     * @return Peer template autonomous system number.
     * 
     */
    private String remoteAsn;
    /**
     * @return Source Interface. Must match first field in the output of `show intf brief`.
     * 
     */
    private String sourceInterface;
    /**
     * @return Peer template name.
     * 
     */
    private String templateName;

    private GetBgpPeerTemplateResult() {}
    /**
     * @return Autonomous system number.
     * 
     */
    public String asn() {
        return this.asn;
    }
    /**
     * @return Peer template description.
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
     * @return The distinguished name of the object.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Neighbor Fabric Type.
     * 
     */
    public String peerType() {
        return this.peerType;
    }
    /**
     * @return Peer template autonomous system number.
     * 
     */
    public String remoteAsn() {
        return this.remoteAsn;
    }
    /**
     * @return Source Interface. Must match first field in the output of `show intf brief`.
     * 
     */
    public String sourceInterface() {
        return this.sourceInterface;
    }
    /**
     * @return Peer template name.
     * 
     */
    public String templateName() {
        return this.templateName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBgpPeerTemplateResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String asn;
        private String description;
        private @Nullable String device;
        private String id;
        private String peerType;
        private String remoteAsn;
        private String sourceInterface;
        private String templateName;
        public Builder() {}
        public Builder(GetBgpPeerTemplateResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.asn = defaults.asn;
    	      this.description = defaults.description;
    	      this.device = defaults.device;
    	      this.id = defaults.id;
    	      this.peerType = defaults.peerType;
    	      this.remoteAsn = defaults.remoteAsn;
    	      this.sourceInterface = defaults.sourceInterface;
    	      this.templateName = defaults.templateName;
        }

        @CustomType.Setter
        public Builder asn(String asn) {
            this.asn = Objects.requireNonNull(asn);
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
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder peerType(String peerType) {
            this.peerType = Objects.requireNonNull(peerType);
            return this;
        }
        @CustomType.Setter
        public Builder remoteAsn(String remoteAsn) {
            this.remoteAsn = Objects.requireNonNull(remoteAsn);
            return this;
        }
        @CustomType.Setter
        public Builder sourceInterface(String sourceInterface) {
            this.sourceInterface = Objects.requireNonNull(sourceInterface);
            return this;
        }
        @CustomType.Setter
        public Builder templateName(String templateName) {
            this.templateName = Objects.requireNonNull(templateName);
            return this;
        }
        public GetBgpPeerTemplateResult build() {
            final var o = new GetBgpPeerTemplateResult();
            o.asn = asn;
            o.description = description;
            o.device = device;
            o.id = id;
            o.peerType = peerType;
            o.remoteAsn = remoteAsn;
            o.sourceInterface = sourceInterface;
            o.templateName = templateName;
            return o;
        }
    }
}
