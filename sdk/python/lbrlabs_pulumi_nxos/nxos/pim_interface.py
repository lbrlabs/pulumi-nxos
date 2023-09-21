# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PimInterfaceArgs', 'PimInterface']

@pulumi.input_type
class PimInterfaceArgs:
    def __init__(__self__, *,
                 interface_id: pulumi.Input[str],
                 vrf_name: pulumi.Input[str],
                 admin_state: Optional[pulumi.Input[str]] = None,
                 bfd: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 dr_priority: Optional[pulumi.Input[int]] = None,
                 passive: Optional[pulumi.Input[bool]] = None,
                 sparse_mode: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a PimInterface resource.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[str] vrf_name: VRF name.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        :param pulumi.Input[str] bfd: BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] dr_priority: Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
        :param pulumi.Input[bool] passive: Passive interface. - Default value: `false`
        :param pulumi.Input[bool] sparse_mode: Sparse mode. - Default value: `false`
        """
        pulumi.set(__self__, "interface_id", interface_id)
        pulumi.set(__self__, "vrf_name", vrf_name)
        if admin_state is not None:
            pulumi.set(__self__, "admin_state", admin_state)
        if bfd is not None:
            pulumi.set(__self__, "bfd", bfd)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if dr_priority is not None:
            pulumi.set(__self__, "dr_priority", dr_priority)
        if passive is not None:
            pulumi.set(__self__, "passive", passive)
        if sparse_mode is not None:
            pulumi.set(__self__, "sparse_mode", sparse_mode)

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
    @pulumi.getter(name="vrfName")
    def vrf_name(self) -> pulumi.Input[str]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf_name")

    @vrf_name.setter
    def vrf_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "vrf_name", value)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> Optional[pulumi.Input[str]]:
        """
        Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        """
        return pulumi.get(self, "admin_state")

    @admin_state.setter
    def admin_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_state", value)

    @property
    @pulumi.getter
    def bfd(self) -> Optional[pulumi.Input[str]]:
        """
        BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
        """
        return pulumi.get(self, "bfd")

    @bfd.setter
    def bfd(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bfd", value)

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
    @pulumi.getter(name="drPriority")
    def dr_priority(self) -> Optional[pulumi.Input[int]]:
        """
        Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
        """
        return pulumi.get(self, "dr_priority")

    @dr_priority.setter
    def dr_priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dr_priority", value)

    @property
    @pulumi.getter
    def passive(self) -> Optional[pulumi.Input[bool]]:
        """
        Passive interface. - Default value: `false`
        """
        return pulumi.get(self, "passive")

    @passive.setter
    def passive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "passive", value)

    @property
    @pulumi.getter(name="sparseMode")
    def sparse_mode(self) -> Optional[pulumi.Input[bool]]:
        """
        Sparse mode. - Default value: `false`
        """
        return pulumi.get(self, "sparse_mode")

    @sparse_mode.setter
    def sparse_mode(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sparse_mode", value)


@pulumi.input_type
class _PimInterfaceState:
    def __init__(__self__, *,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 bfd: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 dr_priority: Optional[pulumi.Input[int]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 passive: Optional[pulumi.Input[bool]] = None,
                 sparse_mode: Optional[pulumi.Input[bool]] = None,
                 vrf_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PimInterface resources.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        :param pulumi.Input[str] bfd: BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] dr_priority: Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[bool] passive: Passive interface. - Default value: `false`
        :param pulumi.Input[bool] sparse_mode: Sparse mode. - Default value: `false`
        :param pulumi.Input[str] vrf_name: VRF name.
        """
        if admin_state is not None:
            pulumi.set(__self__, "admin_state", admin_state)
        if bfd is not None:
            pulumi.set(__self__, "bfd", bfd)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if dr_priority is not None:
            pulumi.set(__self__, "dr_priority", dr_priority)
        if interface_id is not None:
            pulumi.set(__self__, "interface_id", interface_id)
        if passive is not None:
            pulumi.set(__self__, "passive", passive)
        if sparse_mode is not None:
            pulumi.set(__self__, "sparse_mode", sparse_mode)
        if vrf_name is not None:
            pulumi.set(__self__, "vrf_name", vrf_name)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> Optional[pulumi.Input[str]]:
        """
        Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        """
        return pulumi.get(self, "admin_state")

    @admin_state.setter
    def admin_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_state", value)

    @property
    @pulumi.getter
    def bfd(self) -> Optional[pulumi.Input[str]]:
        """
        BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
        """
        return pulumi.get(self, "bfd")

    @bfd.setter
    def bfd(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bfd", value)

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
    @pulumi.getter(name="drPriority")
    def dr_priority(self) -> Optional[pulumi.Input[int]]:
        """
        Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
        """
        return pulumi.get(self, "dr_priority")

    @dr_priority.setter
    def dr_priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "dr_priority", value)

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
    def passive(self) -> Optional[pulumi.Input[bool]]:
        """
        Passive interface. - Default value: `false`
        """
        return pulumi.get(self, "passive")

    @passive.setter
    def passive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "passive", value)

    @property
    @pulumi.getter(name="sparseMode")
    def sparse_mode(self) -> Optional[pulumi.Input[bool]]:
        """
        Sparse mode. - Default value: `false`
        """
        return pulumi.get(self, "sparse_mode")

    @sparse_mode.setter
    def sparse_mode(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sparse_mode", value)

    @property
    @pulumi.getter(name="vrfName")
    def vrf_name(self) -> Optional[pulumi.Input[str]]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf_name")

    @vrf_name.setter
    def vrf_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vrf_name", value)


class PimInterface(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 bfd: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 dr_priority: Optional[pulumi.Input[int]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 passive: Optional[pulumi.Input[bool]] = None,
                 sparse_mode: Optional[pulumi.Input[bool]] = None,
                 vrf_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a PimInterface resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        :param pulumi.Input[str] bfd: BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] dr_priority: Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[bool] passive: Passive interface. - Default value: `false`
        :param pulumi.Input[bool] sparse_mode: Sparse mode. - Default value: `false`
        :param pulumi.Input[str] vrf_name: VRF name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PimInterfaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PimInterface resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PimInterfaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PimInterfaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 bfd: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 dr_priority: Optional[pulumi.Input[int]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 passive: Optional[pulumi.Input[bool]] = None,
                 sparse_mode: Optional[pulumi.Input[bool]] = None,
                 vrf_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PimInterfaceArgs.__new__(PimInterfaceArgs)

            __props__.__dict__["admin_state"] = admin_state
            __props__.__dict__["bfd"] = bfd
            __props__.__dict__["device"] = device
            __props__.__dict__["dr_priority"] = dr_priority
            if interface_id is None and not opts.urn:
                raise TypeError("Missing required property 'interface_id'")
            __props__.__dict__["interface_id"] = interface_id
            __props__.__dict__["passive"] = passive
            __props__.__dict__["sparse_mode"] = sparse_mode
            if vrf_name is None and not opts.urn:
                raise TypeError("Missing required property 'vrf_name'")
            __props__.__dict__["vrf_name"] = vrf_name
        super(PimInterface, __self__).__init__(
            'nxos:nxos/pimInterface:PimInterface',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_state: Optional[pulumi.Input[str]] = None,
            bfd: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None,
            dr_priority: Optional[pulumi.Input[int]] = None,
            interface_id: Optional[pulumi.Input[str]] = None,
            passive: Optional[pulumi.Input[bool]] = None,
            sparse_mode: Optional[pulumi.Input[bool]] = None,
            vrf_name: Optional[pulumi.Input[str]] = None) -> 'PimInterface':
        """
        Get an existing PimInterface resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        :param pulumi.Input[str] bfd: BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] dr_priority: Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[bool] passive: Passive interface. - Default value: `false`
        :param pulumi.Input[bool] sparse_mode: Sparse mode. - Default value: `false`
        :param pulumi.Input[str] vrf_name: VRF name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PimInterfaceState.__new__(_PimInterfaceState)

        __props__.__dict__["admin_state"] = admin_state
        __props__.__dict__["bfd"] = bfd
        __props__.__dict__["device"] = device
        __props__.__dict__["dr_priority"] = dr_priority
        __props__.__dict__["interface_id"] = interface_id
        __props__.__dict__["passive"] = passive
        __props__.__dict__["sparse_mode"] = sparse_mode
        __props__.__dict__["vrf_name"] = vrf_name
        return PimInterface(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> pulumi.Output[str]:
        """
        Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        """
        return pulumi.get(self, "admin_state")

    @property
    @pulumi.getter
    def bfd(self) -> pulumi.Output[str]:
        """
        BFD. - Choices: `none`, `enabled`, `disabled` - Default value: `none`
        """
        return pulumi.get(self, "bfd")

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="drPriority")
    def dr_priority(self) -> pulumi.Output[int]:
        """
        Designated Router priority level. - Range: `1`-`4294967295` - Default value: `1`
        """
        return pulumi.get(self, "dr_priority")

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> pulumi.Output[str]:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "interface_id")

    @property
    @pulumi.getter
    def passive(self) -> pulumi.Output[bool]:
        """
        Passive interface. - Default value: `false`
        """
        return pulumi.get(self, "passive")

    @property
    @pulumi.getter(name="sparseMode")
    def sparse_mode(self) -> pulumi.Output[bool]:
        """
        Sparse mode. - Default value: `false`
        """
        return pulumi.get(self, "sparse_mode")

    @property
    @pulumi.getter(name="vrfName")
    def vrf_name(self) -> pulumi.Output[str]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf_name")
