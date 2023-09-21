# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['LoopbackInterfaceVrfArgs', 'LoopbackInterfaceVrf']

@pulumi.input_type
class LoopbackInterfaceVrfArgs:
    def __init__(__self__, *,
                 interface_id: pulumi.Input[str],
                 vrf_dn: pulumi.Input[str],
                 device: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LoopbackInterfaceVrf resource.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `lo123`.
        :param pulumi.Input[str] vrf_dn: DN of VRF. For example: `sys/inst-VRF1`.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        """
        pulumi.set(__self__, "interface_id", interface_id)
        pulumi.set(__self__, "vrf_dn", vrf_dn)
        if device is not None:
            pulumi.set(__self__, "device", device)

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> pulumi.Input[str]:
        """
        Must match first field in the output of `show intf brief`. Example: `lo123`.
        """
        return pulumi.get(self, "interface_id")

    @interface_id.setter
    def interface_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface_id", value)

    @property
    @pulumi.getter(name="vrfDn")
    def vrf_dn(self) -> pulumi.Input[str]:
        """
        DN of VRF. For example: `sys/inst-VRF1`.
        """
        return pulumi.get(self, "vrf_dn")

    @vrf_dn.setter
    def vrf_dn(self, value: pulumi.Input[str]):
        pulumi.set(self, "vrf_dn", value)

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
class _LoopbackInterfaceVrfState:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 vrf_dn: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LoopbackInterfaceVrf resources.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `lo123`.
        :param pulumi.Input[str] vrf_dn: DN of VRF. For example: `sys/inst-VRF1`.
        """
        if device is not None:
            pulumi.set(__self__, "device", device)
        if interface_id is not None:
            pulumi.set(__self__, "interface_id", interface_id)
        if vrf_dn is not None:
            pulumi.set(__self__, "vrf_dn", vrf_dn)

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
        Must match first field in the output of `show intf brief`. Example: `lo123`.
        """
        return pulumi.get(self, "interface_id")

    @interface_id.setter
    def interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_id", value)

    @property
    @pulumi.getter(name="vrfDn")
    def vrf_dn(self) -> Optional[pulumi.Input[str]]:
        """
        DN of VRF. For example: `sys/inst-VRF1`.
        """
        return pulumi.get(self, "vrf_dn")

    @vrf_dn.setter
    def vrf_dn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vrf_dn", value)


class LoopbackInterfaceVrf(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 vrf_dn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a LoopbackInterfaceVrf resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `lo123`.
        :param pulumi.Input[str] vrf_dn: DN of VRF. For example: `sys/inst-VRF1`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LoopbackInterfaceVrfArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a LoopbackInterfaceVrf resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param LoopbackInterfaceVrfArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LoopbackInterfaceVrfArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 vrf_dn: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LoopbackInterfaceVrfArgs.__new__(LoopbackInterfaceVrfArgs)

            __props__.__dict__["device"] = device
            if interface_id is None and not opts.urn:
                raise TypeError("Missing required property 'interface_id'")
            __props__.__dict__["interface_id"] = interface_id
            if vrf_dn is None and not opts.urn:
                raise TypeError("Missing required property 'vrf_dn'")
            __props__.__dict__["vrf_dn"] = vrf_dn
        super(LoopbackInterfaceVrf, __self__).__init__(
            'nxos:nxos/loopbackInterfaceVrf:LoopbackInterfaceVrf',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device: Optional[pulumi.Input[str]] = None,
            interface_id: Optional[pulumi.Input[str]] = None,
            vrf_dn: Optional[pulumi.Input[str]] = None) -> 'LoopbackInterfaceVrf':
        """
        Get an existing LoopbackInterfaceVrf resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `lo123`.
        :param pulumi.Input[str] vrf_dn: DN of VRF. For example: `sys/inst-VRF1`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LoopbackInterfaceVrfState.__new__(_LoopbackInterfaceVrfState)

        __props__.__dict__["device"] = device
        __props__.__dict__["interface_id"] = interface_id
        __props__.__dict__["vrf_dn"] = vrf_dn
        return LoopbackInterfaceVrf(resource_name, opts=opts, __props__=__props__)

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
        Must match first field in the output of `show intf brief`. Example: `lo123`.
        """
        return pulumi.get(self, "interface_id")

    @property
    @pulumi.getter(name="vrfDn")
    def vrf_dn(self) -> pulumi.Output[str]:
        """
        DN of VRF. For example: `sys/inst-VRF1`.
        """
        return pulumi.get(self, "vrf_dn")

