# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BgpAdvertisedPrefixArgs', 'BgpAdvertisedPrefix']

@pulumi.input_type
class BgpAdvertisedPrefixArgs:
    def __init__(__self__, *,
                 address_family: pulumi.Input[str],
                 asn: pulumi.Input[str],
                 prefix: pulumi.Input[str],
                 vrf: pulumi.Input[str],
                 device: Optional[pulumi.Input[str]] = None,
                 route_map: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BgpAdvertisedPrefix resource.
        :param pulumi.Input[str] address_family: Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
               `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
               `ipv4-mdt` - Default value: `ipv4-ucast`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] prefix: IP address of the network or prefix to advertise.
        :param pulumi.Input[str] vrf: VRF name.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] route_map: Route map to modify attributes.
        """
        pulumi.set(__self__, "address_family", address_family)
        pulumi.set(__self__, "asn", asn)
        pulumi.set(__self__, "prefix", prefix)
        pulumi.set(__self__, "vrf", vrf)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if route_map is not None:
            pulumi.set(__self__, "route_map", route_map)

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
    def prefix(self) -> pulumi.Input[str]:
        """
        IP address of the network or prefix to advertise.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: pulumi.Input[str]):
        pulumi.set(self, "prefix", value)

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
    @pulumi.getter(name="routeMap")
    def route_map(self) -> Optional[pulumi.Input[str]]:
        """
        Route map to modify attributes.
        """
        return pulumi.get(self, "route_map")

    @route_map.setter
    def route_map(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_map", value)


@pulumi.input_type
class _BgpAdvertisedPrefixState:
    def __init__(__self__, *,
                 address_family: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 route_map: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BgpAdvertisedPrefix resources.
        :param pulumi.Input[str] address_family: Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
               `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
               `ipv4-mdt` - Default value: `ipv4-ucast`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] prefix: IP address of the network or prefix to advertise.
        :param pulumi.Input[str] route_map: Route map to modify attributes.
        :param pulumi.Input[str] vrf: VRF name.
        """
        if address_family is not None:
            pulumi.set(__self__, "address_family", address_family)
        if asn is not None:
            pulumi.set(__self__, "asn", asn)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if route_map is not None:
            pulumi.set(__self__, "route_map", route_map)
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
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        IP address of the network or prefix to advertise.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter(name="routeMap")
    def route_map(self) -> Optional[pulumi.Input[str]]:
        """
        Route map to modify attributes.
        """
        return pulumi.get(self, "route_map")

    @route_map.setter
    def route_map(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_map", value)

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


class BgpAdvertisedPrefix(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 route_map: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource can manage the BGP (VRF) advertised prefix configuration.

        - API Documentation: [bgpAdvPrefix](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:AdvPrefix/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.BgpAdvertisedPrefix("example",
            address_family="ipv4-ucast",
            asn="65001",
            prefix="192.168.1.0/24",
            route_map="rt-map",
            vrf="default")
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/bgpAdvertisedPrefix:BgpAdvertisedPrefix example "sys/bgp/inst/dom-[default]/af-[ipv4-ucast]/prefix-[192.168.1.0/24]"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
               `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
               `ipv4-mdt` - Default value: `ipv4-ucast`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] prefix: IP address of the network or prefix to advertise.
        :param pulumi.Input[str] route_map: Route map to modify attributes.
        :param pulumi.Input[str] vrf: VRF name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BgpAdvertisedPrefixArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage the BGP (VRF) advertised prefix configuration.

        - API Documentation: [bgpAdvPrefix](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:AdvPrefix/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.BgpAdvertisedPrefix("example",
            address_family="ipv4-ucast",
            asn="65001",
            prefix="192.168.1.0/24",
            route_map="rt-map",
            vrf="default")
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/bgpAdvertisedPrefix:BgpAdvertisedPrefix example "sys/bgp/inst/dom-[default]/af-[ipv4-ucast]/prefix-[192.168.1.0/24]"
        ```

        :param str resource_name: The name of the resource.
        :param BgpAdvertisedPrefixArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BgpAdvertisedPrefixArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 route_map: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BgpAdvertisedPrefixArgs.__new__(BgpAdvertisedPrefixArgs)

            if address_family is None and not opts.urn:
                raise TypeError("Missing required property 'address_family'")
            __props__.__dict__["address_family"] = address_family
            if asn is None and not opts.urn:
                raise TypeError("Missing required property 'asn'")
            __props__.__dict__["asn"] = asn
            __props__.__dict__["device"] = device
            if prefix is None and not opts.urn:
                raise TypeError("Missing required property 'prefix'")
            __props__.__dict__["prefix"] = prefix
            __props__.__dict__["route_map"] = route_map
            if vrf is None and not opts.urn:
                raise TypeError("Missing required property 'vrf'")
            __props__.__dict__["vrf"] = vrf
        super(BgpAdvertisedPrefix, __self__).__init__(
            'nxos:index/bgpAdvertisedPrefix:BgpAdvertisedPrefix',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address_family: Optional[pulumi.Input[str]] = None,
            asn: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None,
            prefix: Optional[pulumi.Input[str]] = None,
            route_map: Optional[pulumi.Input[str]] = None,
            vrf: Optional[pulumi.Input[str]] = None) -> 'BgpAdvertisedPrefix':
        """
        Get an existing BgpAdvertisedPrefix resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address_family: Address Family. - Choices: `ipv4-ucast`, `ipv4-mcast`, `vpnv4-ucast`, `ipv6-ucast`, `ipv6-mcast`, `vpnv6-ucast`,
               `vpnv6-mcast`, `l2vpn-evpn`, `ipv4-lucast`, `ipv6-lucast`, `lnkstate`, `ipv4-mvpn`, `ipv6-mvpn`, `l2vpn-vpls`,
               `ipv4-mdt` - Default value: `ipv4-ucast`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] prefix: IP address of the network or prefix to advertise.
        :param pulumi.Input[str] route_map: Route map to modify attributes.
        :param pulumi.Input[str] vrf: VRF name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BgpAdvertisedPrefixState.__new__(_BgpAdvertisedPrefixState)

        __props__.__dict__["address_family"] = address_family
        __props__.__dict__["asn"] = asn
        __props__.__dict__["device"] = device
        __props__.__dict__["prefix"] = prefix
        __props__.__dict__["route_map"] = route_map
        __props__.__dict__["vrf"] = vrf
        return BgpAdvertisedPrefix(resource_name, opts=opts, __props__=__props__)

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
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def prefix(self) -> pulumi.Output[str]:
        """
        IP address of the network or prefix to advertise.
        """
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter(name="routeMap")
    def route_map(self) -> pulumi.Output[Optional[str]]:
        """
        Route map to modify attributes.
        """
        return pulumi.get(self, "route_map")

    @property
    @pulumi.getter
    def vrf(self) -> pulumi.Output[str]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf")

