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
    'GetVpcInterfaceResult',
    'AwaitableGetVpcInterfaceResult',
    'get_vpc_interface',
    'get_vpc_interface_output',
]

@pulumi.output_type
class GetVpcInterfaceResult:
    """
    A collection of values returned by getVpcInterface.
    """
    def __init__(__self__, device=None, id=None, port_channel_interface_dn=None, vpc_interface_id=None):
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if port_channel_interface_dn and not isinstance(port_channel_interface_dn, str):
            raise TypeError("Expected argument 'port_channel_interface_dn' to be a str")
        pulumi.set(__self__, "port_channel_interface_dn", port_channel_interface_dn)
        if vpc_interface_id and not isinstance(vpc_interface_id, int):
            raise TypeError("Expected argument 'vpc_interface_id' to be a int")
        pulumi.set(__self__, "vpc_interface_id", vpc_interface_id)

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="portChannelInterfaceDn")
    def port_channel_interface_dn(self) -> str:
        return pulumi.get(self, "port_channel_interface_dn")

    @property
    @pulumi.getter(name="vpcInterfaceId")
    def vpc_interface_id(self) -> int:
        return pulumi.get(self, "vpc_interface_id")


class AwaitableGetVpcInterfaceResult(GetVpcInterfaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVpcInterfaceResult(
            device=self.device,
            id=self.id,
            port_channel_interface_dn=self.port_channel_interface_dn,
            vpc_interface_id=self.vpc_interface_id)


def get_vpc_interface(device: Optional[str] = None,
                      vpc_interface_id: Optional[int] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVpcInterfaceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['vpcInterfaceId'] = vpc_interface_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:nxos/getVpcInterface:getVpcInterface', __args__, opts=opts, typ=GetVpcInterfaceResult).value

    return AwaitableGetVpcInterfaceResult(
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        port_channel_interface_dn=pulumi.get(__ret__, 'port_channel_interface_dn'),
        vpc_interface_id=pulumi.get(__ret__, 'vpc_interface_id'))


@_utilities.lift_output_func(get_vpc_interface)
def get_vpc_interface_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                             vpc_interface_id: Optional[pulumi.Input[int]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVpcInterfaceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
