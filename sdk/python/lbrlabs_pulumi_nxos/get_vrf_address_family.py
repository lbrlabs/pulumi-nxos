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
    'GetVrfAddressFamilyResult',
    'AwaitableGetVrfAddressFamilyResult',
    'get_vrf_address_family',
    'get_vrf_address_family_output',
]

@pulumi.output_type
class GetVrfAddressFamilyResult:
    """
    A collection of values returned by getVrfAddressFamily.
    """
    def __init__(__self__, address_family=None, device=None, id=None, vrf=None):
        if address_family and not isinstance(address_family, str):
            raise TypeError("Expected argument 'address_family' to be a str")
        pulumi.set(__self__, "address_family", address_family)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if vrf and not isinstance(vrf, str):
            raise TypeError("Expected argument 'vrf' to be a str")
        pulumi.set(__self__, "vrf", vrf)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> str:
        """
        Address family.
        """
        return pulumi.get(self, "address_family")

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
    def vrf(self) -> str:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf")


class AwaitableGetVrfAddressFamilyResult(GetVrfAddressFamilyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVrfAddressFamilyResult(
            address_family=self.address_family,
            device=self.device,
            id=self.id,
            vrf=self.vrf)


def get_vrf_address_family(address_family: Optional[str] = None,
                           device: Optional[str] = None,
                           vrf: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVrfAddressFamilyResult:
    """
    This data source can read a VRF Address Family.

    - API Documentation: [rtctrlDomAf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:DomAf/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_vrf_address_family(address_family="ipv4-ucast",
        vrf="VRF1")
    ```


    :param str address_family: Address family.
    :param str device: A device name from the provider configuration.
    :param str vrf: VRF name.
    """
    __args__ = dict()
    __args__['addressFamily'] = address_family
    __args__['device'] = device
    __args__['vrf'] = vrf
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getVrfAddressFamily:getVrfAddressFamily', __args__, opts=opts, typ=GetVrfAddressFamilyResult).value

    return AwaitableGetVrfAddressFamilyResult(
        address_family=pulumi.get(__ret__, 'address_family'),
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        vrf=pulumi.get(__ret__, 'vrf'))


@_utilities.lift_output_func(get_vrf_address_family)
def get_vrf_address_family_output(address_family: Optional[pulumi.Input[str]] = None,
                                  device: Optional[pulumi.Input[Optional[str]]] = None,
                                  vrf: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVrfAddressFamilyResult]:
    """
    This data source can read a VRF Address Family.

    - API Documentation: [rtctrlDomAf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:DomAf/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_vrf_address_family(address_family="ipv4-ucast",
        vrf="VRF1")
    ```


    :param str address_family: Address family.
    :param str device: A device name from the provider configuration.
    :param str vrf: VRF name.
    """
    ...
