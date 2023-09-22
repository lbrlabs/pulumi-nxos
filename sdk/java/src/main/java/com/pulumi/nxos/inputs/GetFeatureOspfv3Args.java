// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFeatureOspfv3Args extends com.pulumi.resources.InvokeArgs {

    public static final GetFeatureOspfv3Args Empty = new GetFeatureOspfv3Args();

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

    private GetFeatureOspfv3Args() {}

    private GetFeatureOspfv3Args(GetFeatureOspfv3Args $) {
        this.device = $.device;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFeatureOspfv3Args defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFeatureOspfv3Args $;

        public Builder() {
            $ = new GetFeatureOspfv3Args();
        }

        public Builder(GetFeatureOspfv3Args defaults) {
            $ = new GetFeatureOspfv3Args(Objects.requireNonNull(defaults));
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

        public GetFeatureOspfv3Args build() {
            return $;
        }
    }

}