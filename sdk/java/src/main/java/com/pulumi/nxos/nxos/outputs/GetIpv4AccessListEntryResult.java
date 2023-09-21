// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetIpv4AccessListEntryResult {
    private Boolean ack;
    private String action;
    private String destinationAddressGroup;
    private String destinationPort1;
    private String destinationPort2;
    private String destinationPortGroup;
    private String destinationPortMask;
    private String destinationPortOperator;
    private String destinationPrefix;
    private String destinationPrefixLength;
    private String destinationPrefixMask;
    private @Nullable String device;
    private Integer dscp;
    private Boolean est;
    private Boolean fin;
    private Boolean fragment;
    private String httpOptionType;
    private Integer icmpCode;
    private Integer icmpType;
    private String id;
    private Boolean logging;
    private String name;
    private String packetLength1;
    private String packetLength2;
    private String packetLengthOperator;
    private String precedence;
    private String protocol;
    private String protocolMask;
    private Boolean psh;
    private String redirect;
    private String remark;
    private Boolean rev;
    private Boolean rst;
    private Integer sequenceNumber;
    private String sourceAddressGroup;
    private String sourcePort1;
    private String sourcePort2;
    private String sourcePortGroup;
    private String sourcePortMask;
    private String sourcePortOperator;
    private String sourcePrefix;
    private String sourcePrefixLength;
    private String sourcePrefixMask;
    private Boolean syn;
    private String timeRange;
    private Integer ttl;
    private Boolean urg;
    private Integer vlan;
    private String vni;

    private GetIpv4AccessListEntryResult() {}
    public Boolean ack() {
        return this.ack;
    }
    public String action() {
        return this.action;
    }
    public String destinationAddressGroup() {
        return this.destinationAddressGroup;
    }
    public String destinationPort1() {
        return this.destinationPort1;
    }
    public String destinationPort2() {
        return this.destinationPort2;
    }
    public String destinationPortGroup() {
        return this.destinationPortGroup;
    }
    public String destinationPortMask() {
        return this.destinationPortMask;
    }
    public String destinationPortOperator() {
        return this.destinationPortOperator;
    }
    public String destinationPrefix() {
        return this.destinationPrefix;
    }
    public String destinationPrefixLength() {
        return this.destinationPrefixLength;
    }
    public String destinationPrefixMask() {
        return this.destinationPrefixMask;
    }
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public Integer dscp() {
        return this.dscp;
    }
    public Boolean est() {
        return this.est;
    }
    public Boolean fin() {
        return this.fin;
    }
    public Boolean fragment() {
        return this.fragment;
    }
    public String httpOptionType() {
        return this.httpOptionType;
    }
    public Integer icmpCode() {
        return this.icmpCode;
    }
    public Integer icmpType() {
        return this.icmpType;
    }
    public String id() {
        return this.id;
    }
    public Boolean logging() {
        return this.logging;
    }
    public String name() {
        return this.name;
    }
    public String packetLength1() {
        return this.packetLength1;
    }
    public String packetLength2() {
        return this.packetLength2;
    }
    public String packetLengthOperator() {
        return this.packetLengthOperator;
    }
    public String precedence() {
        return this.precedence;
    }
    public String protocol() {
        return this.protocol;
    }
    public String protocolMask() {
        return this.protocolMask;
    }
    public Boolean psh() {
        return this.psh;
    }
    public String redirect() {
        return this.redirect;
    }
    public String remark() {
        return this.remark;
    }
    public Boolean rev() {
        return this.rev;
    }
    public Boolean rst() {
        return this.rst;
    }
    public Integer sequenceNumber() {
        return this.sequenceNumber;
    }
    public String sourceAddressGroup() {
        return this.sourceAddressGroup;
    }
    public String sourcePort1() {
        return this.sourcePort1;
    }
    public String sourcePort2() {
        return this.sourcePort2;
    }
    public String sourcePortGroup() {
        return this.sourcePortGroup;
    }
    public String sourcePortMask() {
        return this.sourcePortMask;
    }
    public String sourcePortOperator() {
        return this.sourcePortOperator;
    }
    public String sourcePrefix() {
        return this.sourcePrefix;
    }
    public String sourcePrefixLength() {
        return this.sourcePrefixLength;
    }
    public String sourcePrefixMask() {
        return this.sourcePrefixMask;
    }
    public Boolean syn() {
        return this.syn;
    }
    public String timeRange() {
        return this.timeRange;
    }
    public Integer ttl() {
        return this.ttl;
    }
    public Boolean urg() {
        return this.urg;
    }
    public Integer vlan() {
        return this.vlan;
    }
    public String vni() {
        return this.vni;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIpv4AccessListEntryResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean ack;
        private String action;
        private String destinationAddressGroup;
        private String destinationPort1;
        private String destinationPort2;
        private String destinationPortGroup;
        private String destinationPortMask;
        private String destinationPortOperator;
        private String destinationPrefix;
        private String destinationPrefixLength;
        private String destinationPrefixMask;
        private @Nullable String device;
        private Integer dscp;
        private Boolean est;
        private Boolean fin;
        private Boolean fragment;
        private String httpOptionType;
        private Integer icmpCode;
        private Integer icmpType;
        private String id;
        private Boolean logging;
        private String name;
        private String packetLength1;
        private String packetLength2;
        private String packetLengthOperator;
        private String precedence;
        private String protocol;
        private String protocolMask;
        private Boolean psh;
        private String redirect;
        private String remark;
        private Boolean rev;
        private Boolean rst;
        private Integer sequenceNumber;
        private String sourceAddressGroup;
        private String sourcePort1;
        private String sourcePort2;
        private String sourcePortGroup;
        private String sourcePortMask;
        private String sourcePortOperator;
        private String sourcePrefix;
        private String sourcePrefixLength;
        private String sourcePrefixMask;
        private Boolean syn;
        private String timeRange;
        private Integer ttl;
        private Boolean urg;
        private Integer vlan;
        private String vni;
        public Builder() {}
        public Builder(GetIpv4AccessListEntryResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ack = defaults.ack;
    	      this.action = defaults.action;
    	      this.destinationAddressGroup = defaults.destinationAddressGroup;
    	      this.destinationPort1 = defaults.destinationPort1;
    	      this.destinationPort2 = defaults.destinationPort2;
    	      this.destinationPortGroup = defaults.destinationPortGroup;
    	      this.destinationPortMask = defaults.destinationPortMask;
    	      this.destinationPortOperator = defaults.destinationPortOperator;
    	      this.destinationPrefix = defaults.destinationPrefix;
    	      this.destinationPrefixLength = defaults.destinationPrefixLength;
    	      this.destinationPrefixMask = defaults.destinationPrefixMask;
    	      this.device = defaults.device;
    	      this.dscp = defaults.dscp;
    	      this.est = defaults.est;
    	      this.fin = defaults.fin;
    	      this.fragment = defaults.fragment;
    	      this.httpOptionType = defaults.httpOptionType;
    	      this.icmpCode = defaults.icmpCode;
    	      this.icmpType = defaults.icmpType;
    	      this.id = defaults.id;
    	      this.logging = defaults.logging;
    	      this.name = defaults.name;
    	      this.packetLength1 = defaults.packetLength1;
    	      this.packetLength2 = defaults.packetLength2;
    	      this.packetLengthOperator = defaults.packetLengthOperator;
    	      this.precedence = defaults.precedence;
    	      this.protocol = defaults.protocol;
    	      this.protocolMask = defaults.protocolMask;
    	      this.psh = defaults.psh;
    	      this.redirect = defaults.redirect;
    	      this.remark = defaults.remark;
    	      this.rev = defaults.rev;
    	      this.rst = defaults.rst;
    	      this.sequenceNumber = defaults.sequenceNumber;
    	      this.sourceAddressGroup = defaults.sourceAddressGroup;
    	      this.sourcePort1 = defaults.sourcePort1;
    	      this.sourcePort2 = defaults.sourcePort2;
    	      this.sourcePortGroup = defaults.sourcePortGroup;
    	      this.sourcePortMask = defaults.sourcePortMask;
    	      this.sourcePortOperator = defaults.sourcePortOperator;
    	      this.sourcePrefix = defaults.sourcePrefix;
    	      this.sourcePrefixLength = defaults.sourcePrefixLength;
    	      this.sourcePrefixMask = defaults.sourcePrefixMask;
    	      this.syn = defaults.syn;
    	      this.timeRange = defaults.timeRange;
    	      this.ttl = defaults.ttl;
    	      this.urg = defaults.urg;
    	      this.vlan = defaults.vlan;
    	      this.vni = defaults.vni;
        }

        @CustomType.Setter
        public Builder ack(Boolean ack) {
            this.ack = Objects.requireNonNull(ack);
            return this;
        }
        @CustomType.Setter
        public Builder action(String action) {
            this.action = Objects.requireNonNull(action);
            return this;
        }
        @CustomType.Setter
        public Builder destinationAddressGroup(String destinationAddressGroup) {
            this.destinationAddressGroup = Objects.requireNonNull(destinationAddressGroup);
            return this;
        }
        @CustomType.Setter
        public Builder destinationPort1(String destinationPort1) {
            this.destinationPort1 = Objects.requireNonNull(destinationPort1);
            return this;
        }
        @CustomType.Setter
        public Builder destinationPort2(String destinationPort2) {
            this.destinationPort2 = Objects.requireNonNull(destinationPort2);
            return this;
        }
        @CustomType.Setter
        public Builder destinationPortGroup(String destinationPortGroup) {
            this.destinationPortGroup = Objects.requireNonNull(destinationPortGroup);
            return this;
        }
        @CustomType.Setter
        public Builder destinationPortMask(String destinationPortMask) {
            this.destinationPortMask = Objects.requireNonNull(destinationPortMask);
            return this;
        }
        @CustomType.Setter
        public Builder destinationPortOperator(String destinationPortOperator) {
            this.destinationPortOperator = Objects.requireNonNull(destinationPortOperator);
            return this;
        }
        @CustomType.Setter
        public Builder destinationPrefix(String destinationPrefix) {
            this.destinationPrefix = Objects.requireNonNull(destinationPrefix);
            return this;
        }
        @CustomType.Setter
        public Builder destinationPrefixLength(String destinationPrefixLength) {
            this.destinationPrefixLength = Objects.requireNonNull(destinationPrefixLength);
            return this;
        }
        @CustomType.Setter
        public Builder destinationPrefixMask(String destinationPrefixMask) {
            this.destinationPrefixMask = Objects.requireNonNull(destinationPrefixMask);
            return this;
        }
        @CustomType.Setter
        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        @CustomType.Setter
        public Builder dscp(Integer dscp) {
            this.dscp = Objects.requireNonNull(dscp);
            return this;
        }
        @CustomType.Setter
        public Builder est(Boolean est) {
            this.est = Objects.requireNonNull(est);
            return this;
        }
        @CustomType.Setter
        public Builder fin(Boolean fin) {
            this.fin = Objects.requireNonNull(fin);
            return this;
        }
        @CustomType.Setter
        public Builder fragment(Boolean fragment) {
            this.fragment = Objects.requireNonNull(fragment);
            return this;
        }
        @CustomType.Setter
        public Builder httpOptionType(String httpOptionType) {
            this.httpOptionType = Objects.requireNonNull(httpOptionType);
            return this;
        }
        @CustomType.Setter
        public Builder icmpCode(Integer icmpCode) {
            this.icmpCode = Objects.requireNonNull(icmpCode);
            return this;
        }
        @CustomType.Setter
        public Builder icmpType(Integer icmpType) {
            this.icmpType = Objects.requireNonNull(icmpType);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder logging(Boolean logging) {
            this.logging = Objects.requireNonNull(logging);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder packetLength1(String packetLength1) {
            this.packetLength1 = Objects.requireNonNull(packetLength1);
            return this;
        }
        @CustomType.Setter
        public Builder packetLength2(String packetLength2) {
            this.packetLength2 = Objects.requireNonNull(packetLength2);
            return this;
        }
        @CustomType.Setter
        public Builder packetLengthOperator(String packetLengthOperator) {
            this.packetLengthOperator = Objects.requireNonNull(packetLengthOperator);
            return this;
        }
        @CustomType.Setter
        public Builder precedence(String precedence) {
            this.precedence = Objects.requireNonNull(precedence);
            return this;
        }
        @CustomType.Setter
        public Builder protocol(String protocol) {
            this.protocol = Objects.requireNonNull(protocol);
            return this;
        }
        @CustomType.Setter
        public Builder protocolMask(String protocolMask) {
            this.protocolMask = Objects.requireNonNull(protocolMask);
            return this;
        }
        @CustomType.Setter
        public Builder psh(Boolean psh) {
            this.psh = Objects.requireNonNull(psh);
            return this;
        }
        @CustomType.Setter
        public Builder redirect(String redirect) {
            this.redirect = Objects.requireNonNull(redirect);
            return this;
        }
        @CustomType.Setter
        public Builder remark(String remark) {
            this.remark = Objects.requireNonNull(remark);
            return this;
        }
        @CustomType.Setter
        public Builder rev(Boolean rev) {
            this.rev = Objects.requireNonNull(rev);
            return this;
        }
        @CustomType.Setter
        public Builder rst(Boolean rst) {
            this.rst = Objects.requireNonNull(rst);
            return this;
        }
        @CustomType.Setter
        public Builder sequenceNumber(Integer sequenceNumber) {
            this.sequenceNumber = Objects.requireNonNull(sequenceNumber);
            return this;
        }
        @CustomType.Setter
        public Builder sourceAddressGroup(String sourceAddressGroup) {
            this.sourceAddressGroup = Objects.requireNonNull(sourceAddressGroup);
            return this;
        }
        @CustomType.Setter
        public Builder sourcePort1(String sourcePort1) {
            this.sourcePort1 = Objects.requireNonNull(sourcePort1);
            return this;
        }
        @CustomType.Setter
        public Builder sourcePort2(String sourcePort2) {
            this.sourcePort2 = Objects.requireNonNull(sourcePort2);
            return this;
        }
        @CustomType.Setter
        public Builder sourcePortGroup(String sourcePortGroup) {
            this.sourcePortGroup = Objects.requireNonNull(sourcePortGroup);
            return this;
        }
        @CustomType.Setter
        public Builder sourcePortMask(String sourcePortMask) {
            this.sourcePortMask = Objects.requireNonNull(sourcePortMask);
            return this;
        }
        @CustomType.Setter
        public Builder sourcePortOperator(String sourcePortOperator) {
            this.sourcePortOperator = Objects.requireNonNull(sourcePortOperator);
            return this;
        }
        @CustomType.Setter
        public Builder sourcePrefix(String sourcePrefix) {
            this.sourcePrefix = Objects.requireNonNull(sourcePrefix);
            return this;
        }
        @CustomType.Setter
        public Builder sourcePrefixLength(String sourcePrefixLength) {
            this.sourcePrefixLength = Objects.requireNonNull(sourcePrefixLength);
            return this;
        }
        @CustomType.Setter
        public Builder sourcePrefixMask(String sourcePrefixMask) {
            this.sourcePrefixMask = Objects.requireNonNull(sourcePrefixMask);
            return this;
        }
        @CustomType.Setter
        public Builder syn(Boolean syn) {
            this.syn = Objects.requireNonNull(syn);
            return this;
        }
        @CustomType.Setter
        public Builder timeRange(String timeRange) {
            this.timeRange = Objects.requireNonNull(timeRange);
            return this;
        }
        @CustomType.Setter
        public Builder ttl(Integer ttl) {
            this.ttl = Objects.requireNonNull(ttl);
            return this;
        }
        @CustomType.Setter
        public Builder urg(Boolean urg) {
            this.urg = Objects.requireNonNull(urg);
            return this;
        }
        @CustomType.Setter
        public Builder vlan(Integer vlan) {
            this.vlan = Objects.requireNonNull(vlan);
            return this;
        }
        @CustomType.Setter
        public Builder vni(String vni) {
            this.vni = Objects.requireNonNull(vni);
            return this;
        }
        public GetIpv4AccessListEntryResult build() {
            final var o = new GetIpv4AccessListEntryResult();
            o.ack = ack;
            o.action = action;
            o.destinationAddressGroup = destinationAddressGroup;
            o.destinationPort1 = destinationPort1;
            o.destinationPort2 = destinationPort2;
            o.destinationPortGroup = destinationPortGroup;
            o.destinationPortMask = destinationPortMask;
            o.destinationPortOperator = destinationPortOperator;
            o.destinationPrefix = destinationPrefix;
            o.destinationPrefixLength = destinationPrefixLength;
            o.destinationPrefixMask = destinationPrefixMask;
            o.device = device;
            o.dscp = dscp;
            o.est = est;
            o.fin = fin;
            o.fragment = fragment;
            o.httpOptionType = httpOptionType;
            o.icmpCode = icmpCode;
            o.icmpType = icmpType;
            o.id = id;
            o.logging = logging;
            o.name = name;
            o.packetLength1 = packetLength1;
            o.packetLength2 = packetLength2;
            o.packetLengthOperator = packetLengthOperator;
            o.precedence = precedence;
            o.protocol = protocol;
            o.protocolMask = protocolMask;
            o.psh = psh;
            o.redirect = redirect;
            o.remark = remark;
            o.rev = rev;
            o.rst = rst;
            o.sequenceNumber = sequenceNumber;
            o.sourceAddressGroup = sourceAddressGroup;
            o.sourcePort1 = sourcePort1;
            o.sourcePort2 = sourcePort2;
            o.sourcePortGroup = sourcePortGroup;
            o.sourcePortMask = sourcePortMask;
            o.sourcePortOperator = sourcePortOperator;
            o.sourcePrefix = sourcePrefix;
            o.sourcePrefixLength = sourcePrefixLength;
            o.sourcePrefixMask = sourcePrefixMask;
            o.syn = syn;
            o.timeRange = timeRange;
            o.ttl = ttl;
            o.urg = urg;
            o.vlan = vlan;
            o.vni = vni;
            return o;
        }
    }
}