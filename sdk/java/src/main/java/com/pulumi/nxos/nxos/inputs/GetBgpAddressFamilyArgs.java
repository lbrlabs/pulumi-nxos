// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetBgpAddressFamilyArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetBgpAddressFamilyArgs Empty = new GetBgpAddressFamilyArgs();

    @Import(name="addressFamily", required=true)
    private Output<String> addressFamily;

    public Output<String> addressFamily() {
        return this.addressFamily;
    }

    @Import(name="asn", required=true)
    private Output<String> asn;

    public Output<String> asn() {
        return this.asn;
    }

    @Import(name="device")
    private @Nullable Output<String> device;

    public Optional<Output<String>> device() {
        return Optional.ofNullable(this.device);
    }

    @Import(name="vrf", required=true)
    private Output<String> vrf;

    public Output<String> vrf() {
        return this.vrf;
    }

    private GetBgpAddressFamilyArgs() {}

    private GetBgpAddressFamilyArgs(GetBgpAddressFamilyArgs $) {
        this.addressFamily = $.addressFamily;
        this.asn = $.asn;
        this.device = $.device;
        this.vrf = $.vrf;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetBgpAddressFamilyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetBgpAddressFamilyArgs $;

        public Builder() {
            $ = new GetBgpAddressFamilyArgs();
        }

        public Builder(GetBgpAddressFamilyArgs defaults) {
            $ = new GetBgpAddressFamilyArgs(Objects.requireNonNull(defaults));
        }

        public Builder addressFamily(Output<String> addressFamily) {
            $.addressFamily = addressFamily;
            return this;
        }

        public Builder addressFamily(String addressFamily) {
            return addressFamily(Output.of(addressFamily));
        }

        public Builder asn(Output<String> asn) {
            $.asn = asn;
            return this;
        }

        public Builder asn(String asn) {
            return asn(Output.of(asn));
        }

        public Builder device(@Nullable Output<String> device) {
            $.device = device;
            return this;
        }

        public Builder device(String device) {
            return device(Output.of(device));
        }

        public Builder vrf(Output<String> vrf) {
            $.vrf = vrf;
            return this;
        }

        public Builder vrf(String vrf) {
            return vrf(Output.of(vrf));
        }

        public GetBgpAddressFamilyArgs build() {
            $.addressFamily = Objects.requireNonNull($.addressFamily, "expected parameter 'addressFamily' to be non-null");
            $.asn = Objects.requireNonNull($.asn, "expected parameter 'asn' to be non-null");
            $.vrf = Objects.requireNonNull($.vrf, "expected parameter 'vrf' to be non-null");
            return $;
        }
    }

}
