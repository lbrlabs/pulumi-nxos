// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBgpPeerTemplatePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBgpPeerTemplatePlainArgs Empty = new GetBgpPeerTemplatePlainArgs();

    /**
     * Autonomous system number.
     * 
     */
    @Import(name="asn", required=true)
    private String asn;

    /**
     * @return Autonomous system number.
     * 
     */
    public String asn() {
        return this.asn;
    }

    /**
     * A device name from the provider configuration.
     * 
     */
    @Import(name="device")
    private @Nullable String device;

    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }

    /**
     * Peer template name.
     * 
     */
    @Import(name="templateName", required=true)
    private String templateName;

    /**
     * @return Peer template name.
     * 
     */
    public String templateName() {
        return this.templateName;
    }

    private GetBgpPeerTemplatePlainArgs() {}

    private GetBgpPeerTemplatePlainArgs(GetBgpPeerTemplatePlainArgs $) {
        this.asn = $.asn;
        this.device = $.device;
        this.templateName = $.templateName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBgpPeerTemplatePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBgpPeerTemplatePlainArgs $;

        public Builder() {
            $ = new GetBgpPeerTemplatePlainArgs();
        }

        public Builder(GetBgpPeerTemplatePlainArgs defaults) {
            $ = new GetBgpPeerTemplatePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param asn Autonomous system number.
         * 
         * @return builder
         * 
         */
        public Builder asn(String asn) {
            $.asn = asn;
            return this;
        }

        /**
         * @param device A device name from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder device(@Nullable String device) {
            $.device = device;
            return this;
        }

        /**
         * @param templateName Peer template name.
         * 
         * @return builder
         * 
         */
        public Builder templateName(String templateName) {
            $.templateName = templateName;
            return this;
        }

        public GetBgpPeerTemplatePlainArgs build() {
            $.asn = Objects.requireNonNull($.asn, "expected parameter 'asn' to be non-null");
            $.templateName = Objects.requireNonNull($.templateName, "expected parameter 'templateName' to be non-null");
            return $;
        }
    }

}