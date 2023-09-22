// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetPortChannelInterfaceVrfPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPortChannelInterfaceVrfPlainArgs Empty = new GetPortChannelInterfaceVrfPlainArgs();

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
     * Must match first field in the output of `show intf brief`. Example: `po1`.
     * 
     */
    @Import(name="interfaceId", required=true)
    private String interfaceId;

    /**
     * @return Must match first field in the output of `show intf brief`. Example: `po1`.
     * 
     */
    public String interfaceId() {
        return this.interfaceId;
    }

    private GetPortChannelInterfaceVrfPlainArgs() {}

    private GetPortChannelInterfaceVrfPlainArgs(GetPortChannelInterfaceVrfPlainArgs $) {
        this.device = $.device;
        this.interfaceId = $.interfaceId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPortChannelInterfaceVrfPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPortChannelInterfaceVrfPlainArgs $;

        public Builder() {
            $ = new GetPortChannelInterfaceVrfPlainArgs();
        }

        public Builder(GetPortChannelInterfaceVrfPlainArgs defaults) {
            $ = new GetPortChannelInterfaceVrfPlainArgs(Objects.requireNonNull(defaults));
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
         * @param interfaceId Must match first field in the output of `show intf brief`. Example: `po1`.
         * 
         * @return builder
         * 
         */
        public Builder interfaceId(String interfaceId) {
            $.interfaceId = interfaceId;
            return this;
        }

        public GetPortChannelInterfaceVrfPlainArgs build() {
            $.interfaceId = Objects.requireNonNull($.interfaceId, "expected parameter 'interfaceId' to be non-null");
            return $;
        }
    }

}