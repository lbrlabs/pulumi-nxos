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
    'GetIpv4VrfResult',
    'AwaitableGetIpv4VrfResult',
    'get_ipv4_vrf',
    'get_ipv4_vrf_output',
]

@pulumi.output_type
class GetIpv4VrfResult:
    """
    A collection of values returned by getIpv4Vrf.
    """
    def __init__(__self__, device=None, id=None, name=None):
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


class AwaitableGetIpv4VrfResult(GetIpv4VrfResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpv4VrfResult(
            device=self.device,
            id=self.id,
            name=self.name)


def get_ipv4_vrf(device: Optional[str] = None,
                 name: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpv4VrfResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:nxos/getIpv4Vrf:getIpv4Vrf', __args__, opts=opts, typ=GetIpv4VrfResult).value

    return AwaitableGetIpv4VrfResult(
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_ipv4_vrf)
def get_ipv4_vrf_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                        name: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIpv4VrfResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
