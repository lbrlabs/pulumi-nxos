// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.nxos.Utilities;
import com.pulumi.nxos.nxos.Ipv4AccessListEntryArgs;
import com.pulumi.nxos.nxos.inputs.Ipv4AccessListEntryState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="nxos:nxos/ipv4AccessListEntry:Ipv4AccessListEntry")
public class Ipv4AccessListEntry extends com.pulumi.resources.CustomResource {
    /**
     * Match TCP ACK flag.
     * 
     */
    @Export(name="ack", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ack;

    /**
     * @return Match TCP ACK flag.
     * 
     */
    public Output<Optional<Boolean>> ack() {
        return Codegen.optional(this.ack);
    }
    /**
     * Action. - Choices: `invalid`, `permit`, `deny` - Default value: `invalid`
     * 
     */
    @Export(name="action", refs={String.class}, tree="[0]")
    private Output<String> action;

    /**
     * @return Action. - Choices: `invalid`, `permit`, `deny` - Default value: `invalid`
     * 
     */
    public Output<String> action() {
        return this.action;
    }
    /**
     * Destination address group.
     * 
     */
    @Export(name="destinationAddressGroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destinationAddressGroup;

    /**
     * @return Destination address group.
     * 
     */
    public Output<Optional<String>> destinationAddressGroup() {
        return Codegen.optional(this.destinationAddressGroup);
    }
    /**
     * First destination port number or name. - Choices: `echo`, `discard`, `daytime`, `chargen`, `ftp-data`, `ftp`, `telnet`,
     * `smtp`, `time`, `nameserver`, `whois`, `tacacs`, `domain`, `bootps`, `bootpc`, `tftp`, `gopher`, `finger`, `www`,
     * `hostname`, `pop2`, `pop3`, `sunrpc`, `ident`, `nntp`, `ntp`, `netbios-ns`, `netbios-dgm`, `netbios-ss`, `snmp`,
     * `snmptrap`, `xdmcp`, `bgp`, `irc`, `dnsix`, `mobile-ip`, `pim-auto-rp`, `isakmp`, `biff`, `exec`, `who`, `login`,
     * `syslog`, `cmd`, `lpd`, `talk`, `rip`, `uucp`, `klogin`, `kshell`, `drip`, `non500-isakmp`
     * 
     */
    @Export(name="destinationPort1", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destinationPort1;

    /**
     * @return First destination port number or name. - Choices: `echo`, `discard`, `daytime`, `chargen`, `ftp-data`, `ftp`, `telnet`,
     * `smtp`, `time`, `nameserver`, `whois`, `tacacs`, `domain`, `bootps`, `bootpc`, `tftp`, `gopher`, `finger`, `www`,
     * `hostname`, `pop2`, `pop3`, `sunrpc`, `ident`, `nntp`, `ntp`, `netbios-ns`, `netbios-dgm`, `netbios-ss`, `snmp`,
     * `snmptrap`, `xdmcp`, `bgp`, `irc`, `dnsix`, `mobile-ip`, `pim-auto-rp`, `isakmp`, `biff`, `exec`, `who`, `login`,
     * `syslog`, `cmd`, `lpd`, `talk`, `rip`, `uucp`, `klogin`, `kshell`, `drip`, `non500-isakmp`
     * 
     */
    public Output<Optional<String>> destinationPort1() {
        return Codegen.optional(this.destinationPort1);
    }
    /**
     * Second destination port number or name. - Choices: `echo`, `discard`, `daytime`, `chargen`, `ftp-data`, `ftp`, `telnet`,
     * `smtp`, `time`, `nameserver`, `whois`, `tacacs`, `domain`, `bootps`, `bootpc`, `tftp`, `gopher`, `finger`, `www`,
     * `hostname`, `pop2`, `pop3`, `sunrpc`, `ident`, `nntp`, `ntp`, `netbios-ns`, `netbios-dgm`, `netbios-ss`, `snmp`,
     * `snmptrap`, `xdmcp`, `bgp`, `irc`, `dnsix`, `mobile-ip`, `pim-auto-rp`, `isakmp`, `biff`, `exec`, `who`, `login`,
     * `syslog`, `cmd`, `lpd`, `talk`, `rip`, `uucp`, `klogin`, `kshell`, `drip`, `non500-isakmp`
     * 
     */
    @Export(name="destinationPort2", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destinationPort2;

    /**
     * @return Second destination port number or name. - Choices: `echo`, `discard`, `daytime`, `chargen`, `ftp-data`, `ftp`, `telnet`,
     * `smtp`, `time`, `nameserver`, `whois`, `tacacs`, `domain`, `bootps`, `bootpc`, `tftp`, `gopher`, `finger`, `www`,
     * `hostname`, `pop2`, `pop3`, `sunrpc`, `ident`, `nntp`, `ntp`, `netbios-ns`, `netbios-dgm`, `netbios-ss`, `snmp`,
     * `snmptrap`, `xdmcp`, `bgp`, `irc`, `dnsix`, `mobile-ip`, `pim-auto-rp`, `isakmp`, `biff`, `exec`, `who`, `login`,
     * `syslog`, `cmd`, `lpd`, `talk`, `rip`, `uucp`, `klogin`, `kshell`, `drip`, `non500-isakmp`
     * 
     */
    public Output<Optional<String>> destinationPort2() {
        return Codegen.optional(this.destinationPort2);
    }
    /**
     * Destination port group.
     * 
     */
    @Export(name="destinationPortGroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destinationPortGroup;

    /**
     * @return Destination port group.
     * 
     */
    public Output<Optional<String>> destinationPortGroup() {
        return Codegen.optional(this.destinationPortGroup);
    }
    /**
     * Destination port mask number or name. - Choices: `echo`, `discard`, `daytime`, `chargen`, `ftp-data`, `ftp`, `telnet`,
     * `smtp`, `time`, `nameserver`, `whois`, `tacacs`, `domain`, `bootps`, `bootpc`, `tftp`, `gopher`, `finger`, `www`,
     * `hostname`, `pop2`, `pop3`, `sunrpc`, `ident`, `nntp`, `ntp`, `netbios-ns`, `netbios-dgm`, `netbios-ss`, `snmp`,
     * `snmptrap`, `xdmcp`, `bgp`, `irc`, `dnsix`, `mobile-ip`, `pim-auto-rp`, `isakmp`, `biff`, `exec`, `who`, `login`,
     * `syslog`, `cmd`, `lpd`, `talk`, `rip`, `uucp`, `klogin`, `kshell`, `drip`, `non500-isakmp`
     * 
     */
    @Export(name="destinationPortMask", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destinationPortMask;

    /**
     * @return Destination port mask number or name. - Choices: `echo`, `discard`, `daytime`, `chargen`, `ftp-data`, `ftp`, `telnet`,
     * `smtp`, `time`, `nameserver`, `whois`, `tacacs`, `domain`, `bootps`, `bootpc`, `tftp`, `gopher`, `finger`, `www`,
     * `hostname`, `pop2`, `pop3`, `sunrpc`, `ident`, `nntp`, `ntp`, `netbios-ns`, `netbios-dgm`, `netbios-ss`, `snmp`,
     * `snmptrap`, `xdmcp`, `bgp`, `irc`, `dnsix`, `mobile-ip`, `pim-auto-rp`, `isakmp`, `biff`, `exec`, `who`, `login`,
     * `syslog`, `cmd`, `lpd`, `talk`, `rip`, `uucp`, `klogin`, `kshell`, `drip`, `non500-isakmp`
     * 
     */
    public Output<Optional<String>> destinationPortMask() {
        return Codegen.optional(this.destinationPortMask);
    }
    /**
     * Destination port operator. - Choices: `none`, `lt`, `gt`, `eq`, `neq`, `range` - Default value: `none`
     * 
     */
    @Export(name="destinationPortOperator", refs={String.class}, tree="[0]")
    private Output<String> destinationPortOperator;

    /**
     * @return Destination port operator. - Choices: `none`, `lt`, `gt`, `eq`, `neq`, `range` - Default value: `none`
     * 
     */
    public Output<String> destinationPortOperator() {
        return this.destinationPortOperator;
    }
    /**
     * Destination prefix.
     * 
     */
    @Export(name="destinationPrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destinationPrefix;

    /**
     * @return Destination prefix.
     * 
     */
    public Output<Optional<String>> destinationPrefix() {
        return Codegen.optional(this.destinationPrefix);
    }
    /**
     * Destination prefix length.
     * 
     */
    @Export(name="destinationPrefixLength", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destinationPrefixLength;

    /**
     * @return Destination prefix length.
     * 
     */
    public Output<Optional<String>> destinationPrefixLength() {
        return Codegen.optional(this.destinationPrefixLength);
    }
    /**
     * Destination prefix mask.
     * 
     */
    @Export(name="destinationPrefixMask", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> destinationPrefixMask;

    /**
     * @return Destination prefix mask.
     * 
     */
    public Output<Optional<String>> destinationPrefixMask() {
        return Codegen.optional(this.destinationPrefixMask);
    }
    /**
     * A device name from the provider configuration.
     * 
     */
    @Export(name="device", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> device;

    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Output<Optional<String>> device() {
        return Codegen.optional(this.device);
    }
    /**
     * Match DSCP. - Range: `0`-`63`
     * 
     */
    @Export(name="dscp", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> dscp;

    /**
     * @return Match DSCP. - Range: `0`-`63`
     * 
     */
    public Output<Optional<Integer>> dscp() {
        return Codegen.optional(this.dscp);
    }
    /**
     * Match TCP EST flag.
     * 
     */
    @Export(name="est", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> est;

    /**
     * @return Match TCP EST flag.
     * 
     */
    public Output<Optional<Boolean>> est() {
        return Codegen.optional(this.est);
    }
    /**
     * Match TCP FIN flag.
     * 
     */
    @Export(name="fin", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> fin;

    /**
     * @return Match TCP FIN flag.
     * 
     */
    public Output<Optional<Boolean>> fin() {
        return Codegen.optional(this.fin);
    }
    /**
     * Match non-initial fragment.
     * 
     */
    @Export(name="fragment", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> fragment;

    /**
     * @return Match non-initial fragment.
     * 
     */
    public Output<Optional<Boolean>> fragment() {
        return Codegen.optional(this.fragment);
    }
    /**
     * HTTP option method. - Choices: `invalid`, `get`, `put`, `head`, `post`, `delete`, `trace`, `connect` - Default value:
     * `invalid`
     * 
     */
    @Export(name="httpOptionType", refs={String.class}, tree="[0]")
    private Output<String> httpOptionType;

    /**
     * @return HTTP option method. - Choices: `invalid`, `get`, `put`, `head`, `post`, `delete`, `trace`, `connect` - Default value:
     * `invalid`
     * 
     */
    public Output<String> httpOptionType() {
        return this.httpOptionType;
    }
    /**
     * ICMP code. - Range: `0`-`256` - Default value: `256`
     * 
     */
    @Export(name="icmpCode", refs={Integer.class}, tree="[0]")
    private Output<Integer> icmpCode;

    /**
     * @return ICMP code. - Range: `0`-`256` - Default value: `256`
     * 
     */
    public Output<Integer> icmpCode() {
        return this.icmpCode;
    }
    /**
     * ICMP type. - Range: `0`-`256` - Default value: `256`
     * 
     */
    @Export(name="icmpType", refs={Integer.class}, tree="[0]")
    private Output<Integer> icmpType;

    /**
     * @return ICMP type. - Range: `0`-`256` - Default value: `256`
     * 
     */
    public Output<Integer> icmpType() {
        return this.icmpType;
    }
    /**
     * Log matches against ACL entry. - Default value: `false`
     * 
     */
    @Export(name="logging", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> logging;

    /**
     * @return Log matches against ACL entry. - Default value: `false`
     * 
     */
    public Output<Boolean> logging() {
        return this.logging;
    }
    /**
     * Access list name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Access list name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * First packet length. Either `invalid` or a number between 19 and 9210. - Default value: `invalid`
     * 
     */
    @Export(name="packetLength1", refs={String.class}, tree="[0]")
    private Output<String> packetLength1;

    /**
     * @return First packet length. Either `invalid` or a number between 19 and 9210. - Default value: `invalid`
     * 
     */
    public Output<String> packetLength1() {
        return this.packetLength1;
    }
    /**
     * Second packet length. Either `invalid` or a number between 19 and 9210. - Default value: `invalid`
     * 
     */
    @Export(name="packetLength2", refs={String.class}, tree="[0]")
    private Output<String> packetLength2;

    /**
     * @return Second packet length. Either `invalid` or a number between 19 and 9210. - Default value: `invalid`
     * 
     */
    public Output<String> packetLength2() {
        return this.packetLength2;
    }
    /**
     * Packet length operator. - Choices: `none`, `lt`, `gt`, `eq`, `neq`, `range` - Default value: `none`
     * 
     */
    @Export(name="packetLengthOperator", refs={String.class}, tree="[0]")
    private Output<String> packetLengthOperator;

    /**
     * @return Packet length operator. - Choices: `none`, `lt`, `gt`, `eq`, `neq`, `range` - Default value: `none`
     * 
     */
    public Output<String> packetLengthOperator() {
        return this.packetLengthOperator;
    }
    /**
     * Precedence. Either `unspecified` or a number between 0 and 7. - Default value: `unspecified`
     * 
     */
    @Export(name="precedence", refs={String.class}, tree="[0]")
    private Output<String> precedence;

    /**
     * @return Precedence. Either `unspecified` or a number between 0 and 7. - Default value: `unspecified`
     * 
     */
    public Output<String> precedence() {
        return this.precedence;
    }
    /**
     * Protocol name or number. - Choices: `ip`, `icmp`, `igmp`, `tcp`, `udp`, `gre`, `esp`, `ahp`, `eigrp`, `ospf`, `nos`,
     * `pim`, `pcp`, `udf`
     * 
     */
    @Export(name="protocol", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> protocol;

    /**
     * @return Protocol name or number. - Choices: `ip`, `icmp`, `igmp`, `tcp`, `udp`, `gre`, `esp`, `ahp`, `eigrp`, `ospf`, `nos`,
     * `pim`, `pcp`, `udf`
     * 
     */
    public Output<Optional<String>> protocol() {
        return Codegen.optional(this.protocol);
    }
    /**
     * Protocol mask name or number. - Choices: `ip`, `icmp`, `igmp`, `tcp`, `udp`, `gre`, `esp`, `ahp`, `eigrp`, `ospf`,
     * `nos`, `pim`, `pcp`, `udf`
     * 
     */
    @Export(name="protocolMask", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> protocolMask;

    /**
     * @return Protocol mask name or number. - Choices: `ip`, `icmp`, `igmp`, `tcp`, `udp`, `gre`, `esp`, `ahp`, `eigrp`, `ospf`,
     * `nos`, `pim`, `pcp`, `udf`
     * 
     */
    public Output<Optional<String>> protocolMask() {
        return Codegen.optional(this.protocolMask);
    }
    /**
     * Match TCP PSH flag.
     * 
     */
    @Export(name="psh", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> psh;

    /**
     * @return Match TCP PSH flag.
     * 
     */
    public Output<Optional<Boolean>> psh() {
        return Codegen.optional(this.psh);
    }
    /**
     * Redirect action.
     * 
     */
    @Export(name="redirect", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> redirect;

    /**
     * @return Redirect action.
     * 
     */
    public Output<Optional<String>> redirect() {
        return Codegen.optional(this.redirect);
    }
    /**
     * ACL comment.
     * 
     */
    @Export(name="remark", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> remark;

    /**
     * @return ACL comment.
     * 
     */
    public Output<Optional<String>> remark() {
        return Codegen.optional(this.remark);
    }
    /**
     * Match TCP REV flag.
     * 
     */
    @Export(name="rev", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> rev;

    /**
     * @return Match TCP REV flag.
     * 
     */
    public Output<Optional<Boolean>> rev() {
        return Codegen.optional(this.rev);
    }
    /**
     * Match TCP RST flag.
     * 
     */
    @Export(name="rst", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> rst;

    /**
     * @return Match TCP RST flag.
     * 
     */
    public Output<Optional<Boolean>> rst() {
        return Codegen.optional(this.rst);
    }
    /**
     * Sequence number.
     * 
     */
    @Export(name="sequenceNumber", refs={Integer.class}, tree="[0]")
    private Output<Integer> sequenceNumber;

    /**
     * @return Sequence number.
     * 
     */
    public Output<Integer> sequenceNumber() {
        return this.sequenceNumber;
    }
    /**
     * Source address group.
     * 
     */
    @Export(name="sourceAddressGroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourceAddressGroup;

    /**
     * @return Source address group.
     * 
     */
    public Output<Optional<String>> sourceAddressGroup() {
        return Codegen.optional(this.sourceAddressGroup);
    }
    /**
     * First source port name or number. - Choices: `echo`, `discard`, `daytime`, `chargen`, `ftp-data`, `ftp`, `telnet`,
     * `smtp`, `time`, `nameserver`, `whois`, `tacacs`, `domain`, `bootps`, `bootpc`, `tftp`, `gopher`, `finger`, `www`,
     * `hostname`, `pop2`, `pop3`, `sunrpc`, `ident`, `nntp`, `ntp`, `netbios-ns`, `netbios-dgm`, `netbios-ss`, `snmp`,
     * `snmptrap`, `xdmcp`, `bgp`, `irc`, `dnsix`, `mobile-ip`, `pim-auto-rp`, `isakmp`, `biff`, `exec`, `who`, `login`,
     * `syslog`, `cmd`, `lpd`, `talk`, `rip`, `uucp`, `klogin`, `kshell`, `drip`, `non500-isakmp`
     * 
     */
    @Export(name="sourcePort1", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourcePort1;

    /**
     * @return First source port name or number. - Choices: `echo`, `discard`, `daytime`, `chargen`, `ftp-data`, `ftp`, `telnet`,
     * `smtp`, `time`, `nameserver`, `whois`, `tacacs`, `domain`, `bootps`, `bootpc`, `tftp`, `gopher`, `finger`, `www`,
     * `hostname`, `pop2`, `pop3`, `sunrpc`, `ident`, `nntp`, `ntp`, `netbios-ns`, `netbios-dgm`, `netbios-ss`, `snmp`,
     * `snmptrap`, `xdmcp`, `bgp`, `irc`, `dnsix`, `mobile-ip`, `pim-auto-rp`, `isakmp`, `biff`, `exec`, `who`, `login`,
     * `syslog`, `cmd`, `lpd`, `talk`, `rip`, `uucp`, `klogin`, `kshell`, `drip`, `non500-isakmp`
     * 
     */
    public Output<Optional<String>> sourcePort1() {
        return Codegen.optional(this.sourcePort1);
    }
    /**
     * Second source port name or number. - Choices: `echo`, `discard`, `daytime`, `chargen`, `ftp-data`, `ftp`, `telnet`,
     * `smtp`, `time`, `nameserver`, `whois`, `tacacs`, `domain`, `bootps`, `bootpc`, `tftp`, `gopher`, `finger`, `www`,
     * `hostname`, `pop2`, `pop3`, `sunrpc`, `ident`, `nntp`, `ntp`, `netbios-ns`, `netbios-dgm`, `netbios-ss`, `snmp`,
     * `snmptrap`, `xdmcp`, `bgp`, `irc`, `dnsix`, `mobile-ip`, `pim-auto-rp`, `isakmp`, `biff`, `exec`, `who`, `login`,
     * `syslog`, `cmd`, `lpd`, `talk`, `rip`, `uucp`, `klogin`, `kshell`, `drip`, `non500-isakmp`
     * 
     */
    @Export(name="sourcePort2", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourcePort2;

    /**
     * @return Second source port name or number. - Choices: `echo`, `discard`, `daytime`, `chargen`, `ftp-data`, `ftp`, `telnet`,
     * `smtp`, `time`, `nameserver`, `whois`, `tacacs`, `domain`, `bootps`, `bootpc`, `tftp`, `gopher`, `finger`, `www`,
     * `hostname`, `pop2`, `pop3`, `sunrpc`, `ident`, `nntp`, `ntp`, `netbios-ns`, `netbios-dgm`, `netbios-ss`, `snmp`,
     * `snmptrap`, `xdmcp`, `bgp`, `irc`, `dnsix`, `mobile-ip`, `pim-auto-rp`, `isakmp`, `biff`, `exec`, `who`, `login`,
     * `syslog`, `cmd`, `lpd`, `talk`, `rip`, `uucp`, `klogin`, `kshell`, `drip`, `non500-isakmp`
     * 
     */
    public Output<Optional<String>> sourcePort2() {
        return Codegen.optional(this.sourcePort2);
    }
    /**
     * Source port group.
     * 
     */
    @Export(name="sourcePortGroup", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourcePortGroup;

    /**
     * @return Source port group.
     * 
     */
    public Output<Optional<String>> sourcePortGroup() {
        return Codegen.optional(this.sourcePortGroup);
    }
    /**
     * Source port mask name or number. - Choices: `echo`, `discard`, `daytime`, `chargen`, `ftp-data`, `ftp`, `telnet`,
     * `smtp`, `time`, `nameserver`, `whois`, `tacacs`, `domain`, `bootps`, `bootpc`, `tftp`, `gopher`, `finger`, `www`,
     * `hostname`, `pop2`, `pop3`, `sunrpc`, `ident`, `nntp`, `ntp`, `netbios-ns`, `netbios-dgm`, `netbios-ss`, `snmp`,
     * `snmptrap`, `xdmcp`, `bgp`, `irc`, `dnsix`, `mobile-ip`, `pim-auto-rp`, `isakmp`, `biff`, `exec`, `who`, `login`,
     * `syslog`, `cmd`, `lpd`, `talk`, `rip`, `uucp`, `klogin`, `kshell`, `drip`, `non500-isakmp`
     * 
     */
    @Export(name="sourcePortMask", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourcePortMask;

    /**
     * @return Source port mask name or number. - Choices: `echo`, `discard`, `daytime`, `chargen`, `ftp-data`, `ftp`, `telnet`,
     * `smtp`, `time`, `nameserver`, `whois`, `tacacs`, `domain`, `bootps`, `bootpc`, `tftp`, `gopher`, `finger`, `www`,
     * `hostname`, `pop2`, `pop3`, `sunrpc`, `ident`, `nntp`, `ntp`, `netbios-ns`, `netbios-dgm`, `netbios-ss`, `snmp`,
     * `snmptrap`, `xdmcp`, `bgp`, `irc`, `dnsix`, `mobile-ip`, `pim-auto-rp`, `isakmp`, `biff`, `exec`, `who`, `login`,
     * `syslog`, `cmd`, `lpd`, `talk`, `rip`, `uucp`, `klogin`, `kshell`, `drip`, `non500-isakmp`
     * 
     */
    public Output<Optional<String>> sourcePortMask() {
        return Codegen.optional(this.sourcePortMask);
    }
    /**
     * Source port operator. - Choices: `none`, `lt`, `gt`, `eq`, `neq`, `range` - Default value: `none`
     * 
     */
    @Export(name="sourcePortOperator", refs={String.class}, tree="[0]")
    private Output<String> sourcePortOperator;

    /**
     * @return Source port operator. - Choices: `none`, `lt`, `gt`, `eq`, `neq`, `range` - Default value: `none`
     * 
     */
    public Output<String> sourcePortOperator() {
        return this.sourcePortOperator;
    }
    /**
     * Source prefix.
     * 
     */
    @Export(name="sourcePrefix", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourcePrefix;

    /**
     * @return Source prefix.
     * 
     */
    public Output<Optional<String>> sourcePrefix() {
        return Codegen.optional(this.sourcePrefix);
    }
    /**
     * Source prefix length.
     * 
     */
    @Export(name="sourcePrefixLength", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourcePrefixLength;

    /**
     * @return Source prefix length.
     * 
     */
    public Output<Optional<String>> sourcePrefixLength() {
        return Codegen.optional(this.sourcePrefixLength);
    }
    /**
     * Source prefix mask.
     * 
     */
    @Export(name="sourcePrefixMask", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> sourcePrefixMask;

    /**
     * @return Source prefix mask.
     * 
     */
    public Output<Optional<String>> sourcePrefixMask() {
        return Codegen.optional(this.sourcePrefixMask);
    }
    /**
     * Match TCP SYN flag.
     * 
     */
    @Export(name="syn", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> syn;

    /**
     * @return Match TCP SYN flag.
     * 
     */
    public Output<Optional<Boolean>> syn() {
        return Codegen.optional(this.syn);
    }
    /**
     * Time range name.
     * 
     */
    @Export(name="timeRange", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> timeRange;

    /**
     * @return Time range name.
     * 
     */
    public Output<Optional<String>> timeRange() {
        return Codegen.optional(this.timeRange);
    }
    /**
     * TTL. - Range: `0`-`255` - Default value: `0`
     * 
     */
    @Export(name="ttl", refs={Integer.class}, tree="[0]")
    private Output<Integer> ttl;

    /**
     * @return TTL. - Range: `0`-`255` - Default value: `0`
     * 
     */
    public Output<Integer> ttl() {
        return this.ttl;
    }
    /**
     * Match TCP URG flag.
     * 
     */
    @Export(name="urg", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> urg;

    /**
     * @return Match TCP URG flag.
     * 
     */
    public Output<Optional<Boolean>> urg() {
        return Codegen.optional(this.urg);
    }
    /**
     * VLAN ID. - Range: `0`-`4095` - Default value: `4095`
     * 
     */
    @Export(name="vlan", refs={Integer.class}, tree="[0]")
    private Output<Integer> vlan;

    /**
     * @return VLAN ID. - Range: `0`-`4095` - Default value: `4095`
     * 
     */
    public Output<Integer> vlan() {
        return this.vlan;
    }
    /**
     * NVE VNI ID. Either `invalid` or a number between 0 and 16777216. - Default value: `invalid`
     * 
     */
    @Export(name="vni", refs={String.class}, tree="[0]")
    private Output<String> vni;

    /**
     * @return NVE VNI ID. Either `invalid` or a number between 0 and 16777216. - Default value: `invalid`
     * 
     */
    public Output<String> vni() {
        return this.vni;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Ipv4AccessListEntry(String name) {
        this(name, Ipv4AccessListEntryArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Ipv4AccessListEntry(String name, Ipv4AccessListEntryArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Ipv4AccessListEntry(String name, Ipv4AccessListEntryArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/ipv4AccessListEntry:Ipv4AccessListEntry", name, args == null ? Ipv4AccessListEntryArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Ipv4AccessListEntry(String name, Output<String> id, @Nullable Ipv4AccessListEntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("nxos:nxos/ipv4AccessListEntry:Ipv4AccessListEntry", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Ipv4AccessListEntry get(String name, Output<String> id, @Nullable Ipv4AccessListEntryState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Ipv4AccessListEntry(name, id, state, options);
    }
}
