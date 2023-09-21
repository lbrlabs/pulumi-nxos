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
    'GetFeatureOspfv3Result',
    'AwaitableGetFeatureOspfv3Result',
    'get_feature_ospfv3',
    'get_feature_ospfv3_output',
]

@pulumi.output_type
class GetFeatureOspfv3Result:
    """
    A collection of values returned by getFeatureOspfv3.
    """
    def __init__(__self__, admin_state=None, device=None, id=None):
        if admin_state and not isinstance(admin_state, str):
            raise TypeError("Expected argument 'admin_state' to be a str")
        pulumi.set(__self__, "admin_state", admin_state)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> str:
        return pulumi.get(self, "admin_state")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")


class AwaitableGetFeatureOspfv3Result(GetFeatureOspfv3Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFeatureOspfv3Result(
            admin_state=self.admin_state,
            device=self.device,
            id=self.id)


def get_feature_ospfv3(device: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFeatureOspfv3Result:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['device'] = device
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:nxos/getFeatureOspfv3:getFeatureOspfv3', __args__, opts=opts, typ=GetFeatureOspfv3Result).value

    return AwaitableGetFeatureOspfv3Result(
        admin_state=pulumi.get(__ret__, 'admin_state'),
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_feature_ospfv3)
def get_feature_ospfv3_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFeatureOspfv3Result]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
