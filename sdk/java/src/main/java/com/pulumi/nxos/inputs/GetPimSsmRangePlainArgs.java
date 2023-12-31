// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPimSsmRangePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPimSsmRangePlainArgs Empty = new GetPimSsmRangePlainArgs();

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
     * VRF name.
     * 
     */
    @Import(name="vrfName", required=true)
    private String vrfName;

    /**
     * @return VRF name.
     * 
     */
    public String vrfName() {
        return this.vrfName;
    }

    private GetPimSsmRangePlainArgs() {}

    private GetPimSsmRangePlainArgs(GetPimSsmRangePlainArgs $) {
        this.device = $.device;
        this.vrfName = $.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPimSsmRangePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPimSsmRangePlainArgs $;

        public Builder() {
            $ = new GetPimSsmRangePlainArgs();
        }

        public Builder(GetPimSsmRangePlainArgs defaults) {
            $ = new GetPimSsmRangePlainArgs(Objects.requireNonNull(defaults));
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
         * @param vrfName VRF name.
         * 
         * @return builder
         * 
         */
        public Builder vrfName(String vrfName) {
            $.vrfName = vrfName;
            return this;
        }

        public GetPimSsmRangePlainArgs build() {
            $.vrfName = Objects.requireNonNull($.vrfName, "expected parameter 'vrfName' to be non-null");
            return $;
        }
    }

}
