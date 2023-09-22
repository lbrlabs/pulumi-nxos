// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBgpPeerTemplateMaxPrefixArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBgpPeerTemplateMaxPrefixArgs Empty = new GetBgpPeerTemplateMaxPrefixArgs();

    /**
     * Address Family.
     * 
     */
    @Import(name="addressFamily", required=true)
    private Output<String> addressFamily;

    /**
     * @return Address Family.
     * 
     */
    public Output<String> addressFamily() {
        return this.addressFamily;
    }

    /**
     * Autonomous system number.
     * 
     */
    @Import(name="asn", required=true)
    private Output<String> asn;

    /**
     * @return Autonomous system number.
     * 
     */
    public Output<String> asn() {
        return this.asn;
    }

    /**
     * A device name from the provider configuration.
     * 
     */
    @Import(name="device")
    private @Nullable Output<String> device;

    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    /**
     * Peer template name.
     * 
     */
    @Import(name="templateName", required=true)
    private Output<String> templateName;

    /**
     * @return Peer template name.
     * 
     */
    public Output<String> templateName() {
        return this.templateName;
    }

    private GetBgpPeerTemplateMaxPrefixArgs() {}

    private GetBgpPeerTemplateMaxPrefixArgs(GetBgpPeerTemplateMaxPrefixArgs $) {
        this.addressFamily = $.addressFamily;
        this.asn = $.asn;
        this.device = $.device;
        this.templateName = $.templateName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBgpPeerTemplateMaxPrefixArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBgpPeerTemplateMaxPrefixArgs $;

        public Builder() {
            $ = new GetBgpPeerTemplateMaxPrefixArgs();
        }

        public Builder(GetBgpPeerTemplateMaxPrefixArgs defaults) {
            $ = new GetBgpPeerTemplateMaxPrefixArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addressFamily Address Family.
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(Output<String> addressFamily) {
            $.addressFamily = addressFamily;
            return this;
        }

        /**
         * @param addressFamily Address Family.
         * 
         * @return builder
         * 
         */
        public Builder addressFamily(String addressFamily) {
            return addressFamily(Output.of(addressFamily));
        }

        /**
         * @param asn Autonomous system number.
         * 
         * @return builder
         * 
         */
        public Builder asn(Output<String> asn) {
            $.asn = asn;
            return this;
        }

        /**
         * @param asn Autonomous system number.
         * 
         * @return builder
         * 
         */
        public Builder asn(String asn) {
            return asn(Output.of(asn));
        }

        /**
         * @param device A device name from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        /**
         * @param device A device name from the provider configuration.
         * 
         * @return builder
         * 
         */
        public Builder device(String device) {
            return device(Output.of(device));
        }

        /**
         * @param templateName Peer template name.
         * 
         * @return builder
         * 
         */
        public Builder templateName(Output<String> templateName) {
            $.templateName = templateName;
            return this;
        }

        /**
         * @param templateName Peer template name.
         * 
         * @return builder
         * 
         */
        public Builder templateName(String templateName) {
            return templateName(Output.of(templateName));
        }

        public GetBgpPeerTemplateMaxPrefixArgs build() {
            $.addressFamily = Objects.requireNonNull($.addressFamily, "expected parameter 'addressFamily' to be non-null");
            $.asn = Objects.requireNonNull($.asn, "expected parameter 'asn' to be non-null");
            $.templateName = Objects.requireNonNull($.templateName, "expected parameter 'templateName' to be non-null");
            return $;
        }
    }

}