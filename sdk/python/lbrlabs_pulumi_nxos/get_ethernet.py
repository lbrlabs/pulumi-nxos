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
    'GetEthernetResult',
    'AwaitableGetEthernetResult',
    'get_ethernet',
    'get_ethernet_output',
]

@pulumi.output_type
class GetEthernetResult:
    """
    A collection of values returned by getEthernet.
    """
    def __init__(__self__, device=None, id=None, mtu=None):
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if mtu and not isinstance(mtu, int):
            raise TypeError("Expected argument 'mtu' to be a int")
        pulumi.set(__self__, "mtu", mtu)

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
    @pulumi.getter
    def mtu(self) -> int:
        """
        System jumbo MTU.
        """
        return pulumi.get(self, "mtu")


class AwaitableGetEthernetResult(GetEthernetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEthernetResult(
            device=self.device,
            id=self.id,
            mtu=self.mtu)


def get_ethernet(device: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEthernetResult:
    """
    This data source can read global ethernet settings.

    - API Documentation: [ethpmInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Interfaces/ethpm:Inst/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_ethernet()
    ```


    :param str device: A device name from the provider configuration.
    """
    __args__ = dict()
    __args__['device'] = device
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getEthernet:getEthernet', __args__, opts=opts, typ=GetEthernetResult).value

    return AwaitableGetEthernetResult(
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        mtu=pulumi.get(__ret__, 'mtu'))


@_utilities.lift_output_func(get_ethernet)
def get_ethernet_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEthernetResult]:
    """
    This data source can read global ethernet settings.

    - API Documentation: [ethpmInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Interfaces/ethpm:Inst/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_ethernet()
    ```


    :param str device: A device name from the provider configuration.
    """
    ...