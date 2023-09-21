// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBgpRouteControlPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBgpRouteControlPlainArgs Empty = new GetBgpRouteControlPlainArgs();

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

    @Import(name="vrf", required=true)
    private String vrf;

    public String vrf() {
        return this.vrf;
    }

    private GetBgpRouteControlPlainArgs() {}

    private GetBgpRouteControlPlainArgs(GetBgpRouteControlPlainArgs $) {
        this.asn = $.asn;
        this.device = $.device;
        this.vrf = $.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBgpRouteControlPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBgpRouteControlPlainArgs $;

        public Builder() {
            $ = new GetBgpRouteControlPlainArgs();
        }

        public Builder(GetBgpRouteControlPlainArgs defaults) {
            $ = new GetBgpRouteControlPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder asn(String asn) {
            $.asn = asn;
            return this;
        }

        public Builder device(@Nullable String device) {
            $.device = device;
            return this;
        }

        public Builder vrf(String vrf) {
            $.vrf = vrf;
            return this;
        }

        public GetBgpRouteControlPlainArgs build() {
            $.asn = Objects.requireNonNull($.asn, "expected parameter 'asn' to be non-null");
            $.vrf = Objects.requireNonNull($.vrf, "expected parameter 'vrf' to be non-null");
            return $;
        }
    }

}
