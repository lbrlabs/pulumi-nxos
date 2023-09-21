// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OspfAuthenticationState extends com.pulumi.resources.ResourceArgs {

    public static final OspfAuthenticationState Empty = new OspfAuthenticationState();

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
     * OSPF instance name.
     * 
     */
    @Import(name="instanceName")
    private @Nullable Output<String> instanceName;

    /**
     * @return OSPF instance name.
     * 
     */
    public Optional<Output<String>> instanceName() {
        return Optional.ofNullable(this.instanceName);
    }

    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     * 
     */
    @Import(name="interfaceId")
    private @Nullable Output<String> interfaceId;

    /**
     * @return Must match first field in the output of `show intf brief`. Example: `eth1/1`.
     * 
     */
    public Optional<Output<String>> interfaceId() {
        return Optional.ofNullable(this.interfaceId);
    }

    /**
     * Key used for authentication.
     * 
     */
    @Import(name="key")
    private @Nullable Output<String> key;

    /**
     * @return Key used for authentication.
     * 
     */
    public Optional<Output<String>> key() {
        return Optional.ofNullable(this.key);
    }

    /**
     * Key ID used for authentication. - Range: `0`-`255` - Default value: `0`
     * 
     */
    @Import(name="keyId")
    private @Nullable Output<Integer> keyId;

    /**
     * @return Key ID used for authentication. - Range: `0`-`255` - Default value: `0`
     * 
     */
    public Optional<Output<Integer>> keyId() {
        return Optional.ofNullable(this.keyId);
    }

    /**
     * Encrypted authentication key or plain text key. - Default value: `false`
     * 
     */
    @Import(name="keySecureMode")
    private @Nullable Output<Boolean> keySecureMode;

    /**
     * @return Encrypted authentication key or plain text key. - Default value: `false`
     * 
     */
    public Optional<Output<Boolean>> keySecureMode() {
        return Optional.ofNullable(this.keySecureMode);
    }

    /**
     * Authentication keychain.
     * 
     */
    @Import(name="keychain")
    private @Nullable Output<String> keychain;

    /**
     * @return Authentication keychain.
     * 
     */
    public Optional<Output<String>> keychain() {
        return Optional.ofNullable(this.keychain);
    }

    /**
     * Key used for md5 authentication.
     * 
     */
    @Import(name="md5Key")
    private @Nullable Output<String> md5Key;

    /**
     * @return Key used for md5 authentication.
     * 
     */
    public Optional<Output<String>> md5Key() {
        return Optional.ofNullable(this.md5Key);
    }

    /**
     * Encrypted authentication md5 key or plain text key. - Default value: `false`
     * 
     */
    @Import(name="md5KeySecureMode")
    private @Nullable Output<Boolean> md5KeySecureMode;

    /**
     * @return Encrypted authentication md5 key or plain text key. - Default value: `false`
     * 
     */
    public Optional<Output<Boolean>> md5KeySecureMode() {
        return Optional.ofNullable(this.md5KeySecureMode);
    }

    /**
     * Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * VRF name.
     * 
     */
    @Import(name="vrfName")
    private @Nullable Output<String> vrfName;

    /**
     * @return VRF name.
     * 
     */
    public Optional<Output<String>> vrfName() {
        return Optional.ofNullable(this.vrfName);
    }

    private OspfAuthenticationState() {}

    private OspfAuthenticationState(OspfAuthenticationState $) {
        this.device = $.device;
        this.instanceName = $.instanceName;
        this.interfaceId = $.interfaceId;
        this.key = $.key;
        this.keyId = $.keyId;
        this.keySecureMode = $.keySecureMode;
        this.keychain = $.keychain;
        this.md5Key = $.md5Key;
        this.md5KeySecureMode = $.md5KeySecureMode;
        this.type = $.type;
        this.vrfName = $.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OspfAuthenticationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OspfAuthenticationState $;

        public Builder() {
            $ = new OspfAuthenticationState();
        }

        public Builder(OspfAuthenticationState defaults) {
            $ = new OspfAuthenticationState(Objects.requireNonNull(defaults));
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
         * @param instanceName OSPF instance name.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(@Nullable Output<String> instanceName) {
            $.instanceName = instanceName;
            return this;
        }

        /**
         * @param instanceName OSPF instance name.
         * 
         * @return builder
         * 
         */
        public Builder instanceName(String instanceName) {
            return instanceName(Output.of(instanceName));
        }

        /**
         * @param interfaceId Must match first field in the output of `show intf brief`. Example: `eth1/1`.
         * 
         * @return builder
         * 
         */
        public Builder interfaceId(@Nullable Output<String> interfaceId) {
            $.interfaceId = interfaceId;
            return this;
        }

        /**
         * @param interfaceId Must match first field in the output of `show intf brief`. Example: `eth1/1`.
         * 
         * @return builder
         * 
         */
        public Builder interfaceId(String interfaceId) {
            return interfaceId(Output.of(interfaceId));
        }

        /**
         * @param key Key used for authentication.
         * 
         * @return builder
         * 
         */
        public Builder key(@Nullable Output<String> key) {
            $.key = key;
            return this;
        }

        /**
         * @param key Key used for authentication.
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            return key(Output.of(key));
        }

        /**
         * @param keyId Key ID used for authentication. - Range: `0`-`255` - Default value: `0`
         * 
         * @return builder
         * 
         */
        public Builder keyId(@Nullable Output<Integer> keyId) {
            $.keyId = keyId;
            return this;
        }

        /**
         * @param keyId Key ID used for authentication. - Range: `0`-`255` - Default value: `0`
         * 
         * @return builder
         * 
         */
        public Builder keyId(Integer keyId) {
            return keyId(Output.of(keyId));
        }

        /**
         * @param keySecureMode Encrypted authentication key or plain text key. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder keySecureMode(@Nullable Output<Boolean> keySecureMode) {
            $.keySecureMode = keySecureMode;
            return this;
        }

        /**
         * @param keySecureMode Encrypted authentication key or plain text key. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder keySecureMode(Boolean keySecureMode) {
            return keySecureMode(Output.of(keySecureMode));
        }

        /**
         * @param keychain Authentication keychain.
         * 
         * @return builder
         * 
         */
        public Builder keychain(@Nullable Output<String> keychain) {
            $.keychain = keychain;
            return this;
        }

        /**
         * @param keychain Authentication keychain.
         * 
         * @return builder
         * 
         */
        public Builder keychain(String keychain) {
            return keychain(Output.of(keychain));
        }

        /**
         * @param md5Key Key used for md5 authentication.
         * 
         * @return builder
         * 
         */
        public Builder md5Key(@Nullable Output<String> md5Key) {
            $.md5Key = md5Key;
            return this;
        }

        /**
         * @param md5Key Key used for md5 authentication.
         * 
         * @return builder
         * 
         */
        public Builder md5Key(String md5Key) {
            return md5Key(Output.of(md5Key));
        }

        /**
         * @param md5KeySecureMode Encrypted authentication md5 key or plain text key. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder md5KeySecureMode(@Nullable Output<Boolean> md5KeySecureMode) {
            $.md5KeySecureMode = md5KeySecureMode;
            return this;
        }

        /**
         * @param md5KeySecureMode Encrypted authentication md5 key or plain text key. - Default value: `false`
         * 
         * @return builder
         * 
         */
        public Builder md5KeySecureMode(Boolean md5KeySecureMode) {
            return md5KeySecureMode(Output.of(md5KeySecureMode));
        }

        /**
         * @param type Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Authentication type. - Choices: `none`, `simple`, `md5`, `unspecified` - Default value: `unspecified`
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param vrfName VRF name.
         * 
         * @return builder
         * 
         */
        public Builder vrfName(@Nullable Output<String> vrfName) {
            $.vrfName = vrfName;
            return this;
        }

        /**
         * @param vrfName VRF name.
         * 
         * @return builder
         * 
         */
        public Builder vrfName(String vrfName) {
            return vrfName(Output.of(vrfName));
        }

        public OspfAuthenticationState build() {
            return $;
        }
    }

}
