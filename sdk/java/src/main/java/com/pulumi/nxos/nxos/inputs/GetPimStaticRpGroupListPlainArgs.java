// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPimStaticRpGroupListPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPimStaticRpGroupListPlainArgs Empty = new GetPimStaticRpGroupListPlainArgs();

    @Import(name="address", required=true)
    private String address;

    public String address() {
        return this.address;
    }

    @Import(name="device")
    private @Nullable String device;

    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }

    @Import(name="rpAddress", required=true)
    private String rpAddress;

    public String rpAddress() {
        return this.rpAddress;
    }

    @Import(name="vrfName", required=true)
    private String vrfName;

    public String vrfName() {
        return this.vrfName;
    }

    private GetPimStaticRpGroupListPlainArgs() {}

    private GetPimStaticRpGroupListPlainArgs(GetPimStaticRpGroupListPlainArgs $) {
        this.address = $.address;
        this.device = $.device;
        this.rpAddress = $.rpAddress;
        this.vrfName = $.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPimStaticRpGroupListPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPimStaticRpGroupListPlainArgs $;

        public Builder() {
            $ = new GetPimStaticRpGroupListPlainArgs();
        }

        public Builder(GetPimStaticRpGroupListPlainArgs defaults) {
            $ = new GetPimStaticRpGroupListPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder address(String address) {
            $.address = address;
            return this;
        }

        public Builder device(@Nullable String device) {
            $.device = device;
            return this;
        }

        public Builder rpAddress(String rpAddress) {
            $.rpAddress = rpAddress;
            return this;
        }

        public Builder vrfName(String vrfName) {
            $.vrfName = vrfName;
            return this;
        }

        public GetPimStaticRpGroupListPlainArgs build() {
            $.address = Objects.requireNonNull($.address, "expected parameter 'address' to be non-null");
            $.rpAddress = Objects.requireNonNull($.rpAddress, "expected parameter 'rpAddress' to be non-null");
            $.vrfName = Objects.requireNonNull($.vrfName, "expected parameter 'vrfName' to be non-null");
            return $;
        }
    }

}