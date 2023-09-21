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
    'GetPimVrfResult',
    'AwaitableGetPimVrfResult',
    'get_pim_vrf',
    'get_pim_vrf_output',
]

@pulumi.output_type
class GetPimVrfResult:
    """
    A collection of values returned by getPimVrf.
    """
    def __init__(__self__, admin_state=None, bfd=None, device=None, id=None, name=None):
        if admin_state and not isinstance(admin_state, str):
            raise TypeError("Expected argument 'admin_state' to be a str")
        pulumi.set(__self__, "admin_state", admin_state)
        if bfd and not isinstance(bfd, bool):
            raise TypeError("Expected argument 'bfd' to be a bool")
        pulumi.set(__self__, "bfd", bfd)
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
    @pulumi.getter(name="adminState")
    def admin_state(self) -> str:
        return pulumi.get(self, "admin_state")

    @property
    @pulumi.getter
    def bfd(self) -> bool:
        return pulumi.get(self, "bfd")

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


class AwaitableGetPimVrfResult(GetPimVrfResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPimVrfResult(
            admin_state=self.admin_state,
            bfd=self.bfd,
            device=self.device,
            id=self.id,
            name=self.name)


def get_pim_vrf(device: Optional[str] = None,
                name: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPimVrfResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:nxos/getPimVrf:getPimVrf', __args__, opts=opts, typ=GetPimVrfResult).value

    return AwaitableGetPimVrfResult(
        admin_state=pulumi.get(__ret__, 'admin_state'),
        bfd=pulumi.get(__ret__, 'bfd'),
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_pim_vrf)
def get_pim_vrf_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                       name: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPimVrfResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...