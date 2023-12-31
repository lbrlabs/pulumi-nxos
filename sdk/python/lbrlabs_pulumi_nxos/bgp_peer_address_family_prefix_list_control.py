# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BgpPeerAddressFamilyPrefixListControlArgs', 'BgpPeerAddressFamilyPrefixListControl']

@pulumi.input_type
class BgpPeerAddressFamilyPrefixListControlArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 address_family: pulumi.Input[str],
                 asn: pulumi.Input[str],
                 direction: pulumi.Input[str],
                 vrf: pulumi.Input[str],
                 device: Optional[pulumi.Input[str]] = None,
                 list: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BgpPeerAddressFamilyPrefixListControl resource.
        :param pulumi.Input[str] address: Peer address.
        :param pulumi.Input[str] address_family: Address Family. - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate` - Default
               value: `ipv4-ucast`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] direction: Route Control direction. - Choices: `in`, `out` - Default value: `in`
        :param pulumi.Input[str] vrf: VRF name.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] list: Route Control Prefix-List name.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "address_family", address_family)
        pulumi.set(__self__, "asn", asn)
        pulumi.set(__self__, "direction", direction)
        pulumi.set(__self__, "vrf", vrf)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if list is not None:
            pulumi.set(__self__, "list", list)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        Peer address.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> pulumi.Input[str]:
        """
        Address Family. - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate` - Default
        value: `ipv4-ucast`
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
    def direction(self) -> pulumi.Input[str]:
        """
        Route Control direction. - Choices: `in`, `out` - Default value: `in`
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: pulumi.Input[str]):
        pulumi.set(self, "direction", value)

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
    @pulumi.getter
    def list(self) -> Optional[pulumi.Input[str]]:
        """
        Route Control Prefix-List name.
        """
        return pulumi.get(self, "list")

    @list.setter
    def list(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "list", value)


@pulumi.input_type
class _BgpPeerAddressFamilyPrefixListControlState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 list: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BgpPeerAddressFamilyPrefixListControl resources.
        :param pulumi.Input[str] address: Peer address.
        :param pulumi.Input[str] address_family: Address Family. - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate` - Default
               value: `ipv4-ucast`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] direction: Route Control direction. - Choices: `in`, `out` - Default value: `in`
        :param pulumi.Input[str] list: Route Control Prefix-List name.
        :param pulumi.Input[str] vrf: VRF name.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if address_family is not None:
            pulumi.set(__self__, "address_family", address_family)
        if asn is not None:
            pulumi.set(__self__, "asn", asn)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if direction is not None:
            pulumi.set(__self__, "direction", direction)
        if list is not None:
            pulumi.set(__self__, "list", list)
        if vrf is not None:
            pulumi.set(__self__, "vrf", vrf)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Peer address.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> Optional[pulumi.Input[str]]:
        """
        Address Family. - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate` - Default
        value: `ipv4-ucast`
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
    def direction(self) -> Optional[pulumi.Input[str]]:
        """
        Route Control direction. - Choices: `in`, `out` - Default value: `in`
        """
        return pulumi.get(self, "direction")

    @direction.setter
    def direction(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direction", value)

    @property
    @pulumi.getter
    def list(self) -> Optional[pulumi.Input[str]]:
        """
        Route Control Prefix-List name.
        """
        return pulumi.get(self, "list")

    @list.setter
    def list(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "list", value)

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


class BgpPeerAddressFamilyPrefixListControl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 list: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource can manage the BGP peer address family prefix list control configuration.

        - API Documentation: [bgpPfxCtrlP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:PfxCtrlP/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.BgpPeerAddressFamilyPrefixListControl("example",
            address="192.168.0.1",
            address_family="ipv4-ucast",
            asn="65001",
            direction="in",
            list="PREFIX_LIST1",
            vrf="default")
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/bgpPeerAddressFamilyPrefixListControl:BgpPeerAddressFamilyPrefixListControl example "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]/af-[ipv4-ucast]/pfxctrl-[in]"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Peer address.
        :param pulumi.Input[str] address_family: Address Family. - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate` - Default
               value: `ipv4-ucast`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] direction: Route Control direction. - Choices: `in`, `out` - Default value: `in`
        :param pulumi.Input[str] list: Route Control Prefix-List name.
        :param pulumi.Input[str] vrf: VRF name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BgpPeerAddressFamilyPrefixListControlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage the BGP peer address family prefix list control configuration.

        - API Documentation: [bgpPfxCtrlP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:PfxCtrlP/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.BgpPeerAddressFamilyPrefixListControl("example",
            address="192.168.0.1",
            address_family="ipv4-ucast",
            asn="65001",
            direction="in",
            list="PREFIX_LIST1",
            vrf="default")
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/bgpPeerAddressFamilyPrefixListControl:BgpPeerAddressFamilyPrefixListControl example "sys/bgp/inst/dom-[default]/peer-[192.168.0.1]/af-[ipv4-ucast]/pfxctrl-[in]"
        ```

        :param str resource_name: The name of the resource.
        :param BgpPeerAddressFamilyPrefixListControlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BgpPeerAddressFamilyPrefixListControlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 address_family: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 direction: Optional[pulumi.Input[str]] = None,
                 list: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BgpPeerAddressFamilyPrefixListControlArgs.__new__(BgpPeerAddressFamilyPrefixListControlArgs)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            if address_family is None and not opts.urn:
                raise TypeError("Missing required property 'address_family'")
            __props__.__dict__["address_family"] = address_family
            if asn is None and not opts.urn:
                raise TypeError("Missing required property 'asn'")
            __props__.__dict__["asn"] = asn
            __props__.__dict__["device"] = device
            if direction is None and not opts.urn:
                raise TypeError("Missing required property 'direction'")
            __props__.__dict__["direction"] = direction
            __props__.__dict__["list"] = list
            if vrf is None and not opts.urn:
                raise TypeError("Missing required property 'vrf'")
            __props__.__dict__["vrf"] = vrf
        super(BgpPeerAddressFamilyPrefixListControl, __self__).__init__(
            'nxos:index/bgpPeerAddressFamilyPrefixListControl:BgpPeerAddressFamilyPrefixListControl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            address_family: Optional[pulumi.Input[str]] = None,
            asn: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None,
            direction: Optional[pulumi.Input[str]] = None,
            list: Optional[pulumi.Input[str]] = None,
            vrf: Optional[pulumi.Input[str]] = None) -> 'BgpPeerAddressFamilyPrefixListControl':
        """
        Get an existing BgpPeerAddressFamilyPrefixListControl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Peer address.
        :param pulumi.Input[str] address_family: Address Family. - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate` - Default
               value: `ipv4-ucast`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] direction: Route Control direction. - Choices: `in`, `out` - Default value: `in`
        :param pulumi.Input[str] list: Route Control Prefix-List name.
        :param pulumi.Input[str] vrf: VRF name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BgpPeerAddressFamilyPrefixListControlState.__new__(_BgpPeerAddressFamilyPrefixListControlState)

        __props__.__dict__["address"] = address
        __props__.__dict__["address_family"] = address_family
        __props__.__dict__["asn"] = asn
        __props__.__dict__["device"] = device
        __props__.__dict__["direction"] = direction
        __props__.__dict__["list"] = list
        __props__.__dict__["vrf"] = vrf
        return BgpPeerAddressFamilyPrefixListControl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        Peer address.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> pulumi.Output[str]:
        """
        Address Family. - Choices: `ipv4-ucast`, `vpnv4-ucast`, `ipv6-ucast`, `vpnv6-ucast`, `l2vpn-evpn`, `lnkstate` - Default
        value: `ipv4-ucast`
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
    def direction(self) -> pulumi.Output[str]:
        """
        Route Control direction. - Choices: `in`, `out` - Default value: `in`
        """
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def list(self) -> pulumi.Output[Optional[str]]:
        """
        Route Control Prefix-List name.
        """
        return pulumi.get(self, "list")

    @property
    @pulumi.getter
    def vrf(self) -> pulumi.Output[str]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf")

