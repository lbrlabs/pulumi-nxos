// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.nxos.inputs.RestChildrenArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RestArgs extends com.pulumi.resources.ResourceArgs {

    public static final RestArgs Empty = new RestArgs();

    /**
     * List of children.
     * 
     */
    @Import(name="childrens")
    private @Nullable Output<List<RestChildrenArgs>> childrens;

    /**
     * @return List of children.
     * 
     */
    public Optional<Output<List<RestChildrenArgs>>> childrens() {
        return Optional.ofNullable(this.childrens);
    }

    /**
     * Which class object is being created. (Make sure there is no colon in the classname)
     * 
     */
    @Import(name="className", required=true)
    private Output<String> className;

    /**
     * @return Which class object is being created. (Make sure there is no colon in the classname)
     * 
     */
    public Output<String> className() {
        return this.className;
    }

    /**
     * Map of key-value pairs that need to be passed to the Model object as parameters.
     * 
     */
    @Import(name="content")
    private @Nullable Output<Map<String,String>> content;

    /**
     * @return Map of key-value pairs that need to be passed to the Model object as parameters.
     * 
     */
    public Optional<Output<Map<String,String>>> content() {
        return Optional.ofNullable(this.content);
    }

    /**
     * Delete object during destroy operation. Default value is `true`.
     * 
     */
    @Import(name="delete")
    private @Nullable Output<Boolean> delete;

    /**
     * @return Delete object during destroy operation. Default value is `true`.
     * 
     */
    public Optional<Output<Boolean>> delete() {
        return Optional.ofNullable(this.delete);
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
     * Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].
     * 
     */
    @Import(name="dn", required=true)
    private Output<String> dn;

    /**
     * @return Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].
     * 
     */
    public Output<String> dn() {
        return this.dn;
    }

    private RestArgs() {}

    private RestArgs(RestArgs $) {
        this.childrens = $.childrens;
        this.className = $.className;
        this.content = $.content;
        this.delete = $.delete;
        this.device = $.device;
        this.dn = $.dn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RestArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RestArgs $;

        public Builder() {
            $ = new RestArgs();
        }

        public Builder(RestArgs defaults) {
            $ = new RestArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param childrens List of children.
         * 
         * @return builder
         * 
         */
        public Builder childrens(@Nullable Output<List<RestChildrenArgs>> childrens) {
            $.childrens = childrens;
            return this;
        }

        /**
         * @param childrens List of children.
         * 
         * @return builder
         * 
         */
        public Builder childrens(List<RestChildrenArgs> childrens) {
            return childrens(Output.of(childrens));
        }

        /**
         * @param childrens List of children.
         * 
         * @return builder
         * 
         */
        public Builder childrens(RestChildrenArgs... childrens) {
            return childrens(List.of(childrens));
        }

        /**
         * @param className Which class object is being created. (Make sure there is no colon in the classname)
         * 
         * @return builder
         * 
         */
        public Builder className(Output<String> className) {
            $.className = className;
            return this;
        }

        /**
         * @param className Which class object is being created. (Make sure there is no colon in the classname)
         * 
         * @return builder
         * 
         */
        public Builder className(String className) {
            return className(Output.of(className));
        }

        /**
         * @param content Map of key-value pairs that need to be passed to the Model object as parameters.
         * 
         * @return builder
         * 
         */
        public Builder content(@Nullable Output<Map<String,String>> content) {
            $.content = content;
            return this;
        }

        /**
         * @param content Map of key-value pairs that need to be passed to the Model object as parameters.
         * 
         * @return builder
         * 
         */
        public Builder content(Map<String,String> content) {
            return content(Output.of(content));
        }

        /**
         * @param delete Delete object during destroy operation. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder delete(@Nullable Output<Boolean> delete) {
            $.delete = delete;
            return this;
        }

        /**
         * @param delete Delete object during destroy operation. Default value is `true`.
         * 
         * @return builder
         * 
         */
        public Builder delete(Boolean delete) {
            return delete(Output.of(delete));
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
         * @param dn Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].
         * 
         * @return builder
         * 
         */
        public Builder dn(Output<String> dn) {
            $.dn = dn;
            return this;
        }

        /**
         * @param dn Distinguished name of object being managed including its relative name, e.g. sys/intf/phys-[eth1/1].
         * 
         * @return builder
         * 
         */
        public Builder dn(String dn) {
            return dn(Output.of(dn));
        }

        public RestArgs build() {
            $.className = Objects.requireNonNull($.className, "expected parameter 'className' to be non-null");
            $.dn = Objects.requireNonNull($.dn, "expected parameter 'dn' to be non-null");
            return $;
        }
    }

}
