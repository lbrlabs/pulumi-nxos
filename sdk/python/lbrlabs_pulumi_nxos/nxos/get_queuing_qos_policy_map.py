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
    'GetQueuingQosPolicyMapResult',
    'AwaitableGetQueuingQosPolicyMapResult',
    'get_queuing_qos_policy_map',
    'get_queuing_qos_policy_map_output',
]

@pulumi.output_type
class GetQueuingQosPolicyMapResult:
    """
    A collection of values returned by getQueuingQosPolicyMap.
    """
    def __init__(__self__, device=None, id=None, match_type=None, name=None):
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if match_type and not isinstance(match_type, str):
            raise TypeError("Expected argument 'match_type' to be a str")
        pulumi.set(__self__, "match_type", match_type)
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
    @pulumi.getter(name="matchType")
    def match_type(self) -> str:
        return pulumi.get(self, "match_type")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")


class AwaitableGetQueuingQosPolicyMapResult(GetQueuingQosPolicyMapResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetQueuingQosPolicyMapResult(
            device=self.device,
            id=self.id,
            match_type=self.match_type,
            name=self.name)


def get_queuing_qos_policy_map(device: Optional[str] = None,
                               name: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetQueuingQosPolicyMapResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:nxos/getQueuingQosPolicyMap:getQueuingQosPolicyMap', __args__, opts=opts, typ=GetQueuingQosPolicyMapResult).value

    return AwaitableGetQueuingQosPolicyMapResult(
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        match_type=pulumi.get(__ret__, 'match_type'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_queuing_qos_policy_map)
def get_queuing_qos_policy_map_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                                      name: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetQueuingQosPolicyMapResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
