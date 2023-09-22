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
    'GetDefaultQosClassMapResult',
    'AwaitableGetDefaultQosClassMapResult',
    'get_default_qos_class_map',
    'get_default_qos_class_map_output',
]

@pulumi.output_type
class GetDefaultQosClassMapResult:
    """
    A collection of values returned by getDefaultQosClassMap.
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

    @property
    @pulumi.getter(name="matchType")
    def match_type(self) -> str:
        """
        Match type.
        """
        return pulumi.get(self, "match_type")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Class map name.
        """
        return pulumi.get(self, "name")


class AwaitableGetDefaultQosClassMapResult(GetDefaultQosClassMapResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDefaultQosClassMapResult(
            device=self.device,
            id=self.id,
            match_type=self.match_type,
            name=self.name)


def get_default_qos_class_map(device: Optional[str] = None,
                              name: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDefaultQosClassMapResult:
    """
    This data source can read the default QoS class map configuration.

    - API Documentation: [ipqosCMapInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:CMapInst/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_default_qos_class_map(name="Voice")
    ```


    :param str device: A device name from the provider configuration.
    :param str name: Class map name.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getDefaultQosClassMap:getDefaultQosClassMap', __args__, opts=opts, typ=GetDefaultQosClassMapResult).value

    return AwaitableGetDefaultQosClassMapResult(
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        match_type=pulumi.get(__ret__, 'match_type'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_default_qos_class_map)
def get_default_qos_class_map_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                                     name: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDefaultQosClassMapResult]:
    """
    This data source can read the default QoS class map configuration.

    - API Documentation: [ipqosCMapInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:CMapInst/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_default_qos_class_map(name="Voice")
    ```


    :param str device: A device name from the provider configuration.
    :param str name: Class map name.
    """
    ...