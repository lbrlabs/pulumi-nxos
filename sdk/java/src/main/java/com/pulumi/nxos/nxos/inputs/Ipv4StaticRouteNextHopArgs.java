// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class Ipv4StaticRouteNextHopArgs extends com.pulumi.resources.ResourceArgs {

    public static final Ipv4StaticRouteNextHopArgs Empty = new Ipv4StaticRouteNextHopArgs();

    @Import(name="address", required=true)
    private Output<String> address;

    public Output<String> address() {
        return this.address;
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="interfaceId", required=true)
    private Output<String> interfaceId;

    public Output<String> interfaceId() {
        return this.interfaceId;
    }

    @Import(name="object")
    private @Nullable Output<Integer> object;

    public Optional<Output<Integer>> object() {
        return Optional.ofNullable(this.object);
    }

    @Import(name="preference")
    private @Nullable Output<Integer> preference;

    public Optional<Output<Integer>> preference() {
        return Optional.ofNullable(this.preference);
    }

    @Import(name="tag")
    private @Nullable Output<Integer> tag;

    public Optional<Output<Integer>> tag() {
        return Optional.ofNullable(this.tag);
    }

    @Import(name="vrfName", required=true)
    private Output<String> vrfName;

    public Output<String> vrfName() {
        return this.vrfName;
    }

    private Ipv4StaticRouteNextHopArgs() {}

    private Ipv4StaticRouteNextHopArgs(Ipv4StaticRouteNextHopArgs $) {
        this.address = $.address;
        this.description = $.description;
        this.interfaceId = $.interfaceId;
        this.object = $.object;
        this.preference = $.preference;
        this.tag = $.tag;
        this.vrfName = $.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(Ipv4StaticRouteNextHopArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Ipv4StaticRouteNextHopArgs $;

        public Builder() {
            $ = new Ipv4StaticRouteNextHopArgs();
        }

        public Builder(Ipv4StaticRouteNextHopArgs defaults) {
            $ = new Ipv4StaticRouteNextHopArgs(Objects.requireNonNull(defaults));
        }

        public Builder address(Output<String> address) {
            $.address = address;
            return this;
        }

        public Builder address(String address) {
            return address(Output.of(address));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder interfaceId(Output<String> interfaceId) {
            $.interfaceId = interfaceId;
            return this;
        }

        public Builder interfaceId(String interfaceId) {
            return interfaceId(Output.of(interfaceId));
        }

        public Builder object(@Nullable Output<Integer> object) {
            $.object = object;
            return this;
        }

        public Builder object(Integer object) {
            return object(Output.of(object));
        }

        public Builder preference(@Nullable Output<Integer> preference) {
            $.preference = preference;
            return this;
        }

        public Builder preference(Integer preference) {
            return preference(Output.of(preference));
        }

        public Builder tag(@Nullable Output<Integer> tag) {
            $.tag = tag;
            return this;
        }

        public Builder tag(Integer tag) {
            return tag(Output.of(tag));
        }

        public Builder vrfName(Output<String> vrfName) {
            $.vrfName = vrfName;
            return this;
        }

        public Builder vrfName(String vrfName) {
            return vrfName(Output.of(vrfName));
        }

        public Ipv4StaticRouteNextHopArgs build() {
            $.address = Objects.requireNonNull($.address, "expected parameter 'address' to be non-null");
            $.interfaceId = Objects.requireNonNull($.interfaceId, "expected parameter 'interfaceId' to be non-null");
            $.vrfName = Objects.requireNonNull($.vrfName, "expected parameter 'vrfName' to be non-null");
            return $;
        }
    }

}
