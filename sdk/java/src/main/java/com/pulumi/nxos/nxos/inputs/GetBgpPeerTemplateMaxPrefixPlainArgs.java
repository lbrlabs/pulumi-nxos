// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBgpPeerTemplateMaxPrefixPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBgpPeerTemplateMaxPrefixPlainArgs Empty = new GetBgpPeerTemplateMaxPrefixPlainArgs();

    @Import(name="addressFamily", required=true)
    private String addressFamily;

    public String addressFamily() {
        return this.addressFamily;
    }

    @Import(name="asn", required=true)
    private String asn;

    public String asn() {
        return this.asn;
    }

    @Import(name="device")
    private @Nullable String device;

    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }

    @Import(name="templateName", required=true)
    private String templateName;

    public String templateName() {
        return this.templateName;
    }

    private GetBgpPeerTemplateMaxPrefixPlainArgs() {}

    private GetBgpPeerTemplateMaxPrefixPlainArgs(GetBgpPeerTemplateMaxPrefixPlainArgs $) {
        this.addressFamily = $.addressFamily;
        this.asn = $.asn;
        this.device = $.device;
        this.templateName = $.templateName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBgpPeerTemplateMaxPrefixPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBgpPeerTemplateMaxPrefixPlainArgs $;

        public Builder() {
            $ = new GetBgpPeerTemplateMaxPrefixPlainArgs();
        }

        public Builder(GetBgpPeerTemplateMaxPrefixPlainArgs defaults) {
            $ = new GetBgpPeerTemplateMaxPrefixPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder addressFamily(String addressFamily) {
            $.addressFamily = addressFamily;
            return this;
        }

        public Builder asn(String asn) {
            $.asn = asn;
            return this;
        }

        public Builder device(@Nullable String device) {
            $.device = device;
            return this;
        }

        public Builder templateName(String templateName) {
            $.templateName = templateName;
            return this;
        }

        public GetBgpPeerTemplateMaxPrefixPlainArgs build() {
            $.addressFamily = Objects.requireNonNull($.addressFamily, "expected parameter 'addressFamily' to be non-null");
            $.asn = Objects.requireNonNull($.asn, "expected parameter 'asn' to be non-null");
            $.templateName = Objects.requireNonNull($.templateName, "expected parameter 'templateName' to be non-null");
            return $;
        }
    }

}
