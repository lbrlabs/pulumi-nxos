// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QueuingQosPolicyMapMatchClassMapArgs extends com.pulumi.resources.ResourceArgs {

    public static final QueuingQosPolicyMapMatchClassMapArgs Empty = new QueuingQosPolicyMapMatchClassMapArgs();

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
     * Class map name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Class map name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Policy map name.
     * 
     */
    @Import(name="policyMapName", required=true)
    private Output<String> policyMapName;

    /**
     * @return Policy map name.
     * 
     */
    public Output<String> policyMapName() {
        return this.policyMapName;
    }

    private QueuingQosPolicyMapMatchClassMapArgs() {}

    private QueuingQosPolicyMapMatchClassMapArgs(QueuingQosPolicyMapMatchClassMapArgs $) {
        this.device = $.device;
        this.name = $.name;
        this.policyMapName = $.policyMapName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QueuingQosPolicyMapMatchClassMapArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QueuingQosPolicyMapMatchClassMapArgs $;

        public Builder() {
            $ = new QueuingQosPolicyMapMatchClassMapArgs();
        }

        public Builder(QueuingQosPolicyMapMatchClassMapArgs defaults) {
            $ = new QueuingQosPolicyMapMatchClassMapArgs(Objects.requireNonNull(defaults));
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
         * @param name Class map name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Class map name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param policyMapName Policy map name.
         * 
         * @return builder
         * 
         */
        public Builder policyMapName(Output<String> policyMapName) {
            $.policyMapName = policyMapName;
            return this;
        }

        /**
         * @param policyMapName Policy map name.
         * 
         * @return builder
         * 
         */
        public Builder policyMapName(String policyMapName) {
            return policyMapName(Output.of(policyMapName));
        }

        public QueuingQosPolicyMapMatchClassMapArgs build() {
            $.policyMapName = Objects.requireNonNull($.policyMapName, "expected parameter 'policyMapName' to be non-null");
            return $;
        }
    }

}