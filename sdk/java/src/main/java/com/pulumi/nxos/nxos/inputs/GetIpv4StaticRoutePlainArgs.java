// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIpv4StaticRoutePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIpv4StaticRoutePlainArgs Empty = new GetIpv4StaticRoutePlainArgs();

    @Import(name="device")
    private @Nullable String device;

    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }

    @Import(name="prefix", required=true)
    private String prefix;

    public String prefix() {
        return this.prefix;
    }

    @Import(name="vrfName", required=true)
    private String vrfName;

    public String vrfName() {
        return this.vrfName;
    }

    private GetIpv4StaticRoutePlainArgs() {}

    private GetIpv4StaticRoutePlainArgs(GetIpv4StaticRoutePlainArgs $) {
        this.device = $.device;
        this.prefix = $.prefix;
        this.vrfName = $.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIpv4StaticRoutePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIpv4StaticRoutePlainArgs $;

        public Builder() {
            $ = new GetIpv4StaticRoutePlainArgs();
        }

        public Builder(GetIpv4StaticRoutePlainArgs defaults) {
            $ = new GetIpv4StaticRoutePlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder device(@Nullable String device) {
            $.device = device;
            return this;
        }

        public Builder prefix(String prefix) {
            $.prefix = prefix;
            return this;
        }

        public Builder vrfName(String vrfName) {
            $.vrfName = vrfName;
            return this;
        }

        public GetIpv4StaticRoutePlainArgs build() {
            $.prefix = Objects.requireNonNull($.prefix, "expected parameter 'prefix' to be non-null");
            $.vrfName = Objects.requireNonNull($.vrfName, "expected parameter 'vrfName' to be non-null");
            return $;
        }
    }

}
