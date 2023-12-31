// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SubinterfaceState extends com.pulumi.resources.ResourceArgs {

    public static final SubinterfaceState Empty = new SubinterfaceState();

    /**
     * Administrative state. - Choices: `up`, `down` - Default value: `up`
     * 
     */
    @Import(name="adminState")
    private @Nullable Output<String> adminState;

    /**
     * @return Administrative state. - Choices: `up`, `down` - Default value: `up`
     * 
     */
    public Optional<Output<String>> adminState() {
        return Optional.ofNullable(this.adminState);
    }

    /**
     * Specifies the administrative port bandwidth. - Range: `0`-`3200000000` - Default value: `0`
     * 
     */
    @Import(name="bandwidth")
    private @Nullable Output<Integer> bandwidth;

    /**
     * @return Specifies the administrative port bandwidth. - Range: `0`-`3200000000` - Default value: `0`
     * 
     */
    public Optional<Output<Integer>> bandwidth() {
        return Optional.ofNullable(this.bandwidth);
    }

    /**
     * Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
     * 
     */
    @Import(name="delay")
    private @Nullable Output<Integer> delay;

    /**
     * @return Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
     * 
     */
    public Optional<Output<Integer>> delay() {
        return Optional.ofNullable(this.delay);
    }

    /**
     * Interface description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Interface description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
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
     * Subinterface encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `unknown`
     * 
     */
    @Import(name="encap")
    private @Nullable Output<String> encap;

    /**
     * @return Subinterface encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `unknown`
     * 
     */
    public Optional<Output<String>> encap() {
        return Optional.ofNullable(this.encap);
    }

    /**
     * Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
     * 
     */
    @Import(name="interfaceId")
    private @Nullable Output<String> interfaceId;

    /**
     * @return Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
     * 
     */
    public Optional<Output<String>> interfaceId() {
        return Optional.ofNullable(this.interfaceId);
    }

    /**
     * Administrative link logging. - Choices: `default`, `enable`, `disable` - Default value: `default`
     * 
     */
    @Import(name="linkLogging")
    private @Nullable Output<String> linkLogging;

    /**
     * @return Administrative link logging. - Choices: `default`, `enable`, `disable` - Default value: `default`
     * 
     */
    public Optional<Output<String>> linkLogging() {
        return Optional.ofNullable(this.linkLogging);
    }

    /**
     * The administrative port medium type. - Choices: `broadcast`, `p2p` - Default value: `broadcast`
     * 
     */
    @Import(name="medium")
    private @Nullable Output<String> medium;

    /**
     * @return The administrative port medium type. - Choices: `broadcast`, `p2p` - Default value: `broadcast`
     * 
     */
    public Optional<Output<String>> medium() {
        return Optional.ofNullable(this.medium);
    }

    /**
     * Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
     * 
     */
    @Import(name="mtu")
    private @Nullable Output<Integer> mtu;

    /**
     * @return Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
     * 
     */
    public Optional<Output<Integer>> mtu() {
        return Optional.ofNullable(this.mtu);
    }

    private SubinterfaceState() {}

    private SubinterfaceState(SubinterfaceState $) {
        this.adminState = $.adminState;
        this.bandwidth = $.bandwidth;
        this.delay = $.delay;
        this.description = $.description;
        this.device = $.device;
        this.encap = $.encap;
        this.interfaceId = $.interfaceId;
        this.linkLogging = $.linkLogging;
        this.medium = $.medium;
        this.mtu = $.mtu;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SubinterfaceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SubinterfaceState $;

        public Builder() {
            $ = new SubinterfaceState();
        }

        public Builder(SubinterfaceState defaults) {
            $ = new SubinterfaceState(Objects.requireNonNull(defaults));
        }

        /**
         * @param adminState Administrative state. - Choices: `up`, `down` - Default value: `up`
         * 
         * @return builder
         * 
         */
        public Builder adminState(@Nullable Output<String> adminState) {
            $.adminState = adminState;
            return this;
        }

        /**
         * @param adminState Administrative state. - Choices: `up`, `down` - Default value: `up`
         * 
         * @return builder
         * 
         */
        public Builder adminState(String adminState) {
            return adminState(Output.of(adminState));
        }

        /**
         * @param bandwidth Specifies the administrative port bandwidth. - Range: `0`-`3200000000` - Default value: `0`
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(@Nullable Output<Integer> bandwidth) {
            $.bandwidth = bandwidth;
            return this;
        }

        /**
         * @param bandwidth Specifies the administrative port bandwidth. - Range: `0`-`3200000000` - Default value: `0`
         * 
         * @return builder
         * 
         */
        public Builder bandwidth(Integer bandwidth) {
            return bandwidth(Output.of(bandwidth));
        }

        /**
         * @param delay Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
         * 
         * @return builder
         * 
         */
        public Builder delay(@Nullable Output<Integer> delay) {
            $.delay = delay;
            return this;
        }

        /**
         * @param delay Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
         * 
         * @return builder
         * 
         */
        public Builder delay(Integer delay) {
            return delay(Output.of(delay));
        }

        /**
         * @param description Interface description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Interface description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
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
         * @param encap Subinterface encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `unknown`
         * 
         * @return builder
         * 
         */
        public Builder encap(@Nullable Output<String> encap) {
            $.encap = encap;
            return this;
        }

        /**
         * @param encap Subinterface encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`. - Default value: `unknown`
         * 
         * @return builder
         * 
         */
        public Builder encap(String encap) {
            return encap(Output.of(encap));
        }

        /**
         * @param interfaceId Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
         * 
         * @return builder
         * 
         */
        public Builder interfaceId(@Nullable Output<String> interfaceId) {
            $.interfaceId = interfaceId;
            return this;
        }

        /**
         * @param interfaceId Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
         * 
         * @return builder
         * 
         */
        public Builder interfaceId(String interfaceId) {
            return interfaceId(Output.of(interfaceId));
        }

        /**
         * @param linkLogging Administrative link logging. - Choices: `default`, `enable`, `disable` - Default value: `default`
         * 
         * @return builder
         * 
         */
        public Builder linkLogging(@Nullable Output<String> linkLogging) {
            $.linkLogging = linkLogging;
            return this;
        }

        /**
         * @param linkLogging Administrative link logging. - Choices: `default`, `enable`, `disable` - Default value: `default`
         * 
         * @return builder
         * 
         */
        public Builder linkLogging(String linkLogging) {
            return linkLogging(Output.of(linkLogging));
        }

        /**
         * @param medium The administrative port medium type. - Choices: `broadcast`, `p2p` - Default value: `broadcast`
         * 
         * @return builder
         * 
         */
        public Builder medium(@Nullable Output<String> medium) {
            $.medium = medium;
            return this;
        }

        /**
         * @param medium The administrative port medium type. - Choices: `broadcast`, `p2p` - Default value: `broadcast`
         * 
         * @return builder
         * 
         */
        public Builder medium(String medium) {
            return medium(Output.of(medium));
        }

        /**
         * @param mtu Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
         * 
         * @return builder
         * 
         */
        public Builder mtu(@Nullable Output<Integer> mtu) {
            $.mtu = mtu;
            return this;
        }

        /**
         * @param mtu Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
         * 
         * @return builder
         * 
         */
        public Builder mtu(Integer mtu) {
            return mtu(Output.of(mtu));
        }

        public SubinterfaceState build() {
            return $;
        }
    }

}
