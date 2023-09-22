# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetFeaturePvlanResult',
    'AwaitableGetFeaturePvlanResult',
    'get_feature_pvlan',
    'get_feature_pvlan_output',
]

@pulumi.output_type
class GetFeaturePvlanResult:
    """
    A collection of values returned by getFeaturePvlan.
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
        """
        Administrative state.
        """
        return pulumi.get(self, "admin_state")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The distinguished name of the object.
        """
        return pulumi.get(self, "id")


class AwaitableGetFeaturePvlanResult(GetFeaturePvlanResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFeaturePvlanResult(
            admin_state=self.admin_state,
            device=self.device,
            id=self.id)


def get_feature_pvlan(device: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFeaturePvlanResult:
    """
    This data source can read the PVLAN feature configuration.

    - API Documentation: [fmPvlan](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Pvlan/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_feature_pvlan()
    ```


    :param str device: A device name from the provider configuration.
    """
    __args__ = dict()
    __args__['device'] = device
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getFeaturePvlan:getFeaturePvlan', __args__, opts=opts, typ=GetFeaturePvlanResult).value

    return AwaitableGetFeaturePvlanResult(
        admin_state=pulumi.get(__ret__, 'admin_state'),
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_feature_pvlan)
def get_feature_pvlan_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFeaturePvlanResult]:
    """
    This data source can read the PVLAN feature configuration.

    - API Documentation: [fmPvlan](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Pvlan/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_feature_pvlan()
    ```


    :param str device: A device name from the provider configuration.
    """
    ...