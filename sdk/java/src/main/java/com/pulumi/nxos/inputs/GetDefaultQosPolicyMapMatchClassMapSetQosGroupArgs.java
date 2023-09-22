// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDefaultQosPolicyMapMatchClassMapSetQosGroupArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDefaultQosPolicyMapMatchClassMapSetQosGroupArgs Empty = new GetDefaultQosPolicyMapMatchClassMapSetQosGroupArgs();

    /**
     * Class map name.
     * 
     */
    @Import(name="classMapName", required=true)
    private Output<String> classMapName;

    /**
     * @return Class map name.
     * 
     */
    public Output<String> classMapName() {
        return this.classMapName;
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

    private GetDefaultQosPolicyMapMatchClassMapSetQosGroupArgs() {}

    private GetDefaultQosPolicyMapMatchClassMapSetQosGroupArgs(GetDefaultQosPolicyMapMatchClassMapSetQosGroupArgs $) {
        this.classMapName = $.classMapName;
        this.device = $.device;
        this.policyMapName = $.policyMapName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDefaultQosPolicyMapMatchClassMapSetQosGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDefaultQosPolicyMapMatchClassMapSetQosGroupArgs $;

        public Builder() {
            $ = new GetDefaultQosPolicyMapMatchClassMapSetQosGroupArgs();
        }

        public Builder(GetDefaultQosPolicyMapMatchClassMapSetQosGroupArgs defaults) {
            $ = new GetDefaultQosPolicyMapMatchClassMapSetQosGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param classMapName Class map name.
         * 
         * @return builder
         * 
         */
        public Builder classMapName(Output<String> classMapName) {
            $.classMapName = classMapName;
            return this;
        }

        /**
         * @param classMapName Class map name.
         * 
         * @return builder
         * 
         */
        public Builder classMapName(String classMapName) {
            return classMapName(Output.of(classMapName));
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

        public GetDefaultQosPolicyMapMatchClassMapSetQosGroupArgs build() {
            $.classMapName = Objects.requireNonNull($.classMapName, "expected parameter 'classMapName' to be non-null");
            $.policyMapName = Objects.requireNonNull($.policyMapName, "expected parameter 'policyMapName' to be non-null");
            return $;
        }
    }

}