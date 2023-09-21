# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetSviInterfaceResult',
    'AwaitableGetSviInterfaceResult',
    'get_svi_interface',
    'get_svi_interface_output',
]

@pulumi.output_type
class GetSviInterfaceResult:
    """
    A collection of values returned by getSviInterface.
    """
    def __init__(__self__, admin_state=None, bandwidth=None, delay=None, description=None, device=None, id=None, interface_id=None, medium=None, mtu=None):
        if admin_state and not isinstance(admin_state, str):
            raise TypeError("Expected argument 'admin_state' to be a str")
        pulumi.set(__self__, "admin_state", admin_state)
        if bandwidth and not isinstance(bandwidth, int):
            raise TypeError("Expected argument 'bandwidth' to be a int")
        pulumi.set(__self__, "bandwidth", bandwidth)
        if delay and not isinstance(delay, int):
            raise TypeError("Expected argument 'delay' to be a int")
        pulumi.set(__self__, "delay", delay)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface_id and not isinstance(interface_id, str):
            raise TypeError("Expected argument 'interface_id' to be a str")
        pulumi.set(__self__, "interface_id", interface_id)
        if medium and not isinstance(medium, str):
            raise TypeError("Expected argument 'medium' to be a str")
        pulumi.set(__self__, "medium", medium)
        if mtu and not isinstance(mtu, int):
            raise TypeError("Expected argument 'mtu' to be a int")
        pulumi.set(__self__, "mtu", mtu)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> str:
        return pulumi.get(self, "admin_state")

    @property
    @pulumi.getter
    def bandwidth(self) -> int:
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter
    def delay(self) -> int:
        return pulumi.get(self, "delay")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> str:
        return pulumi.get(self, "interface_id")

    @property
    @pulumi.getter
    def medium(self) -> str:
        return pulumi.get(self, "medium")

    @property
    @pulumi.getter
    def mtu(self) -> int:
        return pulumi.get(self, "mtu")


class AwaitableGetSviInterfaceResult(GetSviInterfaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSviInterfaceResult(
            admin_state=self.admin_state,
            bandwidth=self.bandwidth,
            delay=self.delay,
            description=self.description,
            device=self.device,
            id=self.id,
            interface_id=self.interface_id,
            medium=self.medium,
            mtu=self.mtu)


def get_svi_interface(device: Optional[str] = None,
                      interface_id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSviInterfaceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['interfaceId'] = interface_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:nxos/getSviInterface:getSviInterface', __args__, opts=opts, typ=GetSviInterfaceResult).value

    return AwaitableGetSviInterfaceResult(
        admin_state=pulumi.get(__ret__, 'admin_state'),
        bandwidth=pulumi.get(__ret__, 'bandwidth'),
        delay=pulumi.get(__ret__, 'delay'),
        description=pulumi.get(__ret__, 'description'),
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        interface_id=pulumi.get(__ret__, 'interface_id'),
        medium=pulumi.get(__ret__, 'medium'),
        mtu=pulumi.get(__ret__, 'mtu'))


@_utilities.lift_output_func(get_svi_interface)
def get_svi_interface_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                             interface_id: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSviInterfaceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...