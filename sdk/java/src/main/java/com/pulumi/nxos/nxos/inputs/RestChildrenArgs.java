// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RestChildrenArgs extends com.pulumi.resources.ResourceArgs {

    public static final RestChildrenArgs Empty = new RestChildrenArgs();

    @Import(name="className", required=true)
    private Output<String> className;

    public Output<String> className() {
        return this.className;
    }

    @Import(name="content")
    private @Nullable Output<Map<String,String>> content;

    public Optional<Output<Map<String,String>>> content() {
        return Optional.ofNullable(this.content);
    }

    @Import(name="rn", required=true)
    private Output<String> rn;

    public Output<String> rn() {
        return this.rn;
    }

    private RestChildrenArgs() {}

    private RestChildrenArgs(RestChildrenArgs $) {
        this.className = $.className;
        this.content = $.content;
        this.rn = $.rn;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RestChildrenArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RestChildrenArgs $;

        public Builder() {
            $ = new RestChildrenArgs();
        }

        public Builder(RestChildrenArgs defaults) {
            $ = new RestChildrenArgs(Objects.requireNonNull(defaults));
        }

        public Builder className(Output<String> className) {
            $.className = className;
            return this;
        }

        public Builder className(String className) {
            return className(Output.of(className));
        }

        public Builder content(@Nullable Output<Map<String,String>> content) {
            $.content = content;
            return this;
        }

        public Builder content(Map<String,String> content) {
            return content(Output.of(content));
        }

        public Builder rn(Output<String> rn) {
            $.rn = rn;
            return this;
        }

        public Builder rn(String rn) {
            return rn(Output.of(rn));
        }

        public RestChildrenArgs build() {
            $.className = Objects.requireNonNull($.className, "expected parameter 'className' to be non-null");
            $.rn = Objects.requireNonNull($.rn, "expected parameter 'rn' to be non-null");
            return $;
        }
    }

}
