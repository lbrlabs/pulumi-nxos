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
    'GetFeatureOspfResult',
    'AwaitableGetFeatureOspfResult',
    'get_feature_ospf',
    'get_feature_ospf_output',
]

@pulumi.output_type
class GetFeatureOspfResult:
    """
    A collection of values returned by getFeatureOspf.
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


class AwaitableGetFeatureOspfResult(GetFeatureOspfResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFeatureOspfResult(
            admin_state=self.admin_state,
            device=self.device,
            id=self.id)


def get_feature_ospf(device: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFeatureOspfResult:
    """
    This data source can read the OSPF feature configuration.

    - API Documentation: [fmOspf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Ospf/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_feature_ospf()
    ```


    :param str device: A device name from the provider configuration.
    """
    __args__ = dict()
    __args__['device'] = device
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getFeatureOspf:getFeatureOspf', __args__, opts=opts, typ=GetFeatureOspfResult).value

    return AwaitableGetFeatureOspfResult(
        admin_state=pulumi.get(__ret__, 'admin_state'),
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_feature_ospf)
def get_feature_ospf_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFeatureOspfResult]:
    """
    This data source can read the OSPF feature configuration.

    - API Documentation: [fmOspf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Feature%20Management/fm:Ospf/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_feature_ospf()
    ```


    :param str device: A device name from the provider configuration.
    """
    ...