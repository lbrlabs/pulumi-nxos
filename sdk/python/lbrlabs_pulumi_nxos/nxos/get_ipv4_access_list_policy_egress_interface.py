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
    'GetIpv4AccessListPolicyEgressInterfaceResult',
    'AwaitableGetIpv4AccessListPolicyEgressInterfaceResult',
    'get_ipv4_access_list_policy_egress_interface',
    'get_ipv4_access_list_policy_egress_interface_output',
]

@pulumi.output_type
class GetIpv4AccessListPolicyEgressInterfaceResult:
    """
    A collection of values returned by getIpv4AccessListPolicyEgressInterface.
    """
    def __init__(__self__, access_list_name=None, device=None, id=None, interface_id=None):
        if access_list_name and not isinstance(access_list_name, str):
            raise TypeError("Expected argument 'access_list_name' to be a str")
        pulumi.set(__self__, "access_list_name", access_list_name)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface_id and not isinstance(interface_id, str):
            raise TypeError("Expected argument 'interface_id' to be a str")
        pulumi.set(__self__, "interface_id", interface_id)

    @property
    @pulumi.getter(name="accessListName")
    def access_list_name(self) -> str:
        return pulumi.get(self, "access_list_name")

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


class AwaitableGetIpv4AccessListPolicyEgressInterfaceResult(GetIpv4AccessListPolicyEgressInterfaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpv4AccessListPolicyEgressInterfaceResult(
            access_list_name=self.access_list_name,
            device=self.device,
            id=self.id,
            interface_id=self.interface_id)


def get_ipv4_access_list_policy_egress_interface(device: Optional[str] = None,
                                                 interface_id: Optional[str] = None,
                                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpv4AccessListPolicyEgressInterfaceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['interfaceId'] = interface_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:nxos/getIpv4AccessListPolicyEgressInterface:getIpv4AccessListPolicyEgressInterface', __args__, opts=opts, typ=GetIpv4AccessListPolicyEgressInterfaceResult).value

    return AwaitableGetIpv4AccessListPolicyEgressInterfaceResult(
        access_list_name=pulumi.get(__ret__, 'access_list_name'),
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        interface_id=pulumi.get(__ret__, 'interface_id'))


@_utilities.lift_output_func(get_ipv4_access_list_policy_egress_interface)
def get_ipv4_access_list_policy_egress_interface_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                                                        interface_id: Optional[pulumi.Input[str]] = None,
                                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIpv4AccessListPolicyEgressInterfaceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
