# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DhcpRelayAddressArgs', 'DhcpRelayAddress']

@pulumi.input_type
class DhcpRelayAddressArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 interface_id: pulumi.Input[str],
                 vrf: pulumi.Input[str],
                 device: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DhcpRelayAddress resource.
        :param pulumi.Input[str] address: IPv4 or IPv6 address.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[str] vrf: VRF name.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "interface_id", interface_id)
        pulumi.set(__self__, "vrf", vrf)
        if device is not None:
            pulumi.set(__self__, "device", device)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        IPv4 or IPv6 address.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> pulumi.Input[str]:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "interface_id")

    @interface_id.setter
    def interface_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface_id", value)

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


@pulumi.input_type
class _DhcpRelayAddressState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DhcpRelayAddress resources.
        :param pulumi.Input[str] address: IPv4 or IPv6 address.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[str] vrf: VRF name.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if interface_id is not None:
            pulumi.set(__self__, "interface_id", interface_id)
        if vrf is not None:
            pulumi.set(__self__, "vrf", vrf)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 or IPv6 address.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

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
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "interface_id")

    @interface_id.setter
    def interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_id", value)

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


class DhcpRelayAddress(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a DhcpRelayAddress resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: IPv4 or IPv6 address.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[str] vrf: VRF name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DhcpRelayAddressArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a DhcpRelayAddress resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param DhcpRelayAddressArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DhcpRelayAddressArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DhcpRelayAddressArgs.__new__(DhcpRelayAddressArgs)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            __props__.__dict__["device"] = device
            if interface_id is None and not opts.urn:
                raise TypeError("Missing required property 'interface_id'")
            __props__.__dict__["interface_id"] = interface_id
            if vrf is None and not opts.urn:
                raise TypeError("Missing required property 'vrf'")
            __props__.__dict__["vrf"] = vrf
        super(DhcpRelayAddress, __self__).__init__(
            'nxos:nxos/dhcpRelayAddress:DhcpRelayAddress',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None,
            interface_id: Optional[pulumi.Input[str]] = None,
            vrf: Optional[pulumi.Input[str]] = None) -> 'DhcpRelayAddress':
        """
        Get an existing DhcpRelayAddress resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: IPv4 or IPv6 address.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[str] vrf: VRF name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DhcpRelayAddressState.__new__(_DhcpRelayAddressState)

        __props__.__dict__["address"] = address
        __props__.__dict__["device"] = device
        __props__.__dict__["interface_id"] = interface_id
        __props__.__dict__["vrf"] = vrf
        return DhcpRelayAddress(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        IPv4 or IPv6 address.
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> pulumi.Output[str]:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "interface_id")

    @property
    @pulumi.getter
    def vrf(self) -> pulumi.Output[str]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf")

