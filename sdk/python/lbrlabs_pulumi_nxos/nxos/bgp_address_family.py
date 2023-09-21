# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BgpAddressFamilyArgs', 'BgpAddressFamily']

@pulumi.input_type
class BgpAddressFamilyArgs:
    def __init__(__self__, *,
                 address_family: pulumi.Input[str],
                 asn: pulumi.Input[str],
                 vrf: pulumi.Input[str],
                 critical_nexthop_timeout: Optional[pulumi.Input[int]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 non_critical_nexthop_timeout: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a BgpAddressFamily resource.
        :param pulumi.Input[str] address_family: Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
               `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
               `ipv4-mdt` - Default value: `ipv4-ucast`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] vrf: VRF name.
        :param pulumi.Input[int] critical_nexthop_timeout: The next-hop address tracking delay timer for critical next-hop reachability routes. - Range: `1`-`4294967295` - Default
               value: `3000`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] non_critical_nexthop_timeout: The next-hop address tracking delay timer for non-critical next-hop reachability routes. - Range: `1`-`4294967295` -
               Default value: `10000`
        """
        pulumi.set(__self__, "address_family", address_family)
        pulumi.set(__self__, "asn", asn)
        pulumi.set(__self__, "vrf", vrf)
        if critical_nexthop_timeout is not None:
            pulumi.set(__self__, "critical_nexthop_timeout", critical_nexthop_timeout)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if non_critical_nexthop_timeout is not None:
            pulumi.set(__self__, "non_critical_nexthop_timeout", non_critical_nexthop_timeout)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> pulumi.Input[str]:
        """
        Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
        `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
        `ipv4-mdt` - Default value: `ipv4-ucast`
        """
        return pulumi.get(self, "address_family")

    @address_family.setter
    def address_family(self, value: pulumi.Input[str]):
        pulumi.set(self, "address_family", value)

    @property
    @pulumi.getter
    def asn(self) -> pulumi.Input[str]:
        """
        Autonomous system number.
        """
        return pulumi.get(self, "asn")

    @asn.setter
    def asn(self, value: pulumi.Input[str]):
        pulumi.set(self, "asn", value)

    @property
    @pulumi.getter
    def vrf(self) -> pulumi.Input[str]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf")

    @vrf.setter
    def vrf(self, value: pulumi.Input[str]):
        pulumi.set(self, "vrf", value)

    @property
    @pulumi.getter(name="criticalNexthopTimeout")
    def critical_nexthop_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The next-hop address tracking delay timer for critical next-hop reachability routes. - Range: `1`-`4294967295` - Default
        value: `3000`
        """
        return pulumi.get(self, "critical_nexthop_timeout")

    @critical_nexthop_timeout.setter
    def critical_nexthop_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "critical_nexthop_timeout", value)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter(name="nonCriticalNexthopTimeout")
    def non_critical_nexthop_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The next-hop address tracking delay timer for non-critical next-hop reachability routes. - Range: `1`-`4294967295` -
        Default value: `10000`
        """
        return pulumi.get(self, "non_critical_nexthop_timeout")

    @non_critical_nexthop_timeout.setter
    def non_critical_nexthop_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "non_critical_nexthop_timeout", value)


@pulumi.input_type
class _BgpAddressFamilyState:
    def __init__(__self__, *,
                 address_family: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 critical_nexthop_timeout: Optional[pulumi.Input[int]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 non_critical_nexthop_timeout: Optional[pulumi.Input[int]] = None,
                 vrf: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BgpAddressFamily resources.
        :param pulumi.Input[str] address_family: Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
               `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
               `ipv4-mdt` - Default value: `ipv4-ucast`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[int] critical_nexthop_timeout: The next-hop address tracking delay timer for critical next-hop reachability routes. - Range: `1`-`4294967295` - Default
               value: `3000`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] non_critical_nexthop_timeout: The next-hop address tracking delay timer for non-critical next-hop reachability routes. - Range: `1`-`4294967295` -
               Default value: `10000`
        :param pulumi.Input[str] vrf: VRF name.
        """
        if address_family is not None:
            pulumi.set(__self__, "address_family", address_family)
        if asn is not None:
            pulumi.set(__self__, "asn", asn)
        if critical_nexthop_timeout is not None:
            pulumi.set(__self__, "critical_nexthop_timeout", critical_nexthop_timeout)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if non_critical_nexthop_timeout is not None:
            pulumi.set(__self__, "non_critical_nexthop_timeout", non_critical_nexthop_timeout)
        if vrf is not None:
            pulumi.set(__self__, "vrf", vrf)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> Optional[pulumi.Input[str]]:
        """
        Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
        `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
        `ipv4-mdt` - Default value: `ipv4-ucast`
        """
        return pulumi.get(self, "address_family")

    @address_family.setter
    def address_family(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address_family", value)

    @property
    @pulumi.getter
    def asn(self) -> Optional[pulumi.Input[str]]:
        """
        Autonomous system number.
        """
        return pulumi.get(self, "asn")

    @asn.setter
    def asn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "asn", value)

    @property
    @pulumi.getter(name="criticalNexthopTimeout")
    def critical_nexthop_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The next-hop address tracking delay timer for critical next-hop reachability routes. - Range: `1`-`4294967295` - Default
        value: `3000`
        """
        return pulumi.get(self, "critical_nexthop_timeout")

    @critical_nexthop_timeout.setter
    def critical_nexthop_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "critical_nexthop_timeout", value)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter(name="nonCriticalNexthopTimeout")
    def non_critical_nexthop_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        The next-hop address tracking delay timer for non-critical next-hop reachability routes. - Range: `1`-`4294967295` -
        Default value: `10000`
        """
        return pulumi.get(self, "non_critical_nexthop_timeout")

    @non_critical_nexthop_timeout.setter
    def non_critical_nexthop_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "non_critical_nexthop_timeout", value)

    @property
    @pulumi.getter
    def vrf(self) -> Optional[pulumi.Input[str]]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf")

    @vrf.setter
    def vrf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vrf", value)


class BgpAddressFamily(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 critical_nexthop_timeout: Optional[pulumi.Input[int]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 non_critical_nexthop_timeout: Optional[pulumi.Input[int]] = None,
                 vrf: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a BgpAddressFamily resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
               `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
               `ipv4-mdt` - Default value: `ipv4-ucast`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[int] critical_nexthop_timeout: The next-hop address tracking delay timer for critical next-hop reachability routes. - Range: `1`-`4294967295` - Default
               value: `3000`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] non_critical_nexthop_timeout: The next-hop address tracking delay timer for non-critical next-hop reachability routes. - Range: `1`-`4294967295` -
               Default value: `10000`
        :param pulumi.Input[str] vrf: VRF name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BgpAddressFamilyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a BgpAddressFamily resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param BgpAddressFamilyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BgpAddressFamilyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 critical_nexthop_timeout: Optional[pulumi.Input[int]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 non_critical_nexthop_timeout: Optional[pulumi.Input[int]] = None,
                 vrf: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BgpAddressFamilyArgs.__new__(BgpAddressFamilyArgs)

            if address_family is None and not opts.urn:
                raise TypeError("Missing required property 'address_family'")
            __props__.__dict__["address_family"] = address_family
            if asn is None and not opts.urn:
                raise TypeError("Missing required property 'asn'")
            __props__.__dict__["asn"] = asn
            __props__.__dict__["critical_nexthop_timeout"] = critical_nexthop_timeout
            __props__.__dict__["device"] = device
            __props__.__dict__["non_critical_nexthop_timeout"] = non_critical_nexthop_timeout
            if vrf is None and not opts.urn:
                raise TypeError("Missing required property 'vrf'")
            __props__.__dict__["vrf"] = vrf
        super(BgpAddressFamily, __self__).__init__(
            'nxos:nxos/bgpAddressFamily:BgpAddressFamily',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_family: Optional[pulumi.Input[str]] = None,
            asn: Optional[pulumi.Input[str]] = None,
            critical_nexthop_timeout: Optional[pulumi.Input[int]] = None,
            device: Optional[pulumi.Input[str]] = None,
            non_critical_nexthop_timeout: Optional[pulumi.Input[int]] = None,
            vrf: Optional[pulumi.Input[str]] = None) -> 'BgpAddressFamily':
        """
        Get an existing BgpAddressFamily resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
               `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
               `ipv4-mdt` - Default value: `ipv4-ucast`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[int] critical_nexthop_timeout: The next-hop address tracking delay timer for critical next-hop reachability routes. - Range: `1`-`4294967295` - Default
               value: `3000`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] non_critical_nexthop_timeout: The next-hop address tracking delay timer for non-critical next-hop reachability routes. - Range: `1`-`4294967295` -
               Default value: `10000`
        :param pulumi.Input[str] vrf: VRF name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BgpAddressFamilyState.__new__(_BgpAddressFamilyState)

        __props__.__dict__["address_family"] = address_family
        __props__.__dict__["asn"] = asn
        __props__.__dict__["critical_nexthop_timeout"] = critical_nexthop_timeout
        __props__.__dict__["device"] = device
        __props__.__dict__["non_critical_nexthop_timeout"] = non_critical_nexthop_timeout
        __props__.__dict__["vrf"] = vrf
        return BgpAddressFamily(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> pulumi.Output[str]:
        """
        Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
        `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
        `ipv4-mdt` - Default value: `ipv4-ucast`
        """
        return pulumi.get(self, "address_family")

    @property
    @pulumi.getter
    def asn(self) -> pulumi.Output[str]:
        """
        Autonomous system number.
        """
        return pulumi.get(self, "asn")

    @property
    @pulumi.getter(name="criticalNexthopTimeout")
    def critical_nexthop_timeout(self) -> pulumi.Output[int]:
        """
        The next-hop address tracking delay timer for critical next-hop reachability routes. - Range: `1`-`4294967295` - Default
        value: `3000`
        """
        return pulumi.get(self, "critical_nexthop_timeout")

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="nonCriticalNexthopTimeout")
    def non_critical_nexthop_timeout(self) -> pulumi.Output[int]:
        """
        The next-hop address tracking delay timer for non-critical next-hop reachability routes. - Range: `1`-`4294967295` -
        Default value: `10000`
        """
        return pulumi.get(self, "non_critical_nexthop_timeout")

    @property
    @pulumi.getter
    def vrf(self) -> pulumi.Output[str]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf")

