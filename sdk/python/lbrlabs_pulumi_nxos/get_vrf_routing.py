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
    'GetVrfRoutingResult',
    'AwaitableGetVrfRoutingResult',
    'get_vrf_routing',
    'get_vrf_routing_output',
]

@pulumi.output_type
class GetVrfRoutingResult:
    """
    A collection of values returned by getVrfRouting.
    """
    def __init__(__self__, device=None, id=None, route_distinguisher=None, vrf=None):
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if route_distinguisher and not isinstance(route_distinguisher, str):
            raise TypeError("Expected argument 'route_distinguisher' to be a str")
        pulumi.set(__self__, "route_distinguisher", route_distinguisher)
        if vrf and not isinstance(vrf, str):
            raise TypeError("Expected argument 'vrf' to be a str")
        pulumi.set(__self__, "vrf", vrf)

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
    @pulumi.getter(name="routeDistinguisher")
    def route_distinguisher(self) -> str:
        """
        Route Distinguisher value in NX-OS DME format.
        """
        return pulumi.get(self, "route_distinguisher")

    @property
    @pulumi.getter
    def vrf(self) -> str:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf")


class AwaitableGetVrfRoutingResult(GetVrfRoutingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVrfRoutingResult(
            device=self.device,
            id=self.id,
            route_distinguisher=self.route_distinguisher,
            vrf=self.vrf)


def get_vrf_routing(device: Optional[str] = None,
                    vrf: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVrfRoutingResult:
    """
    This data source can read a VRF Route Distinguisher and VRF VNI.

    - API Documentation: [rtctrlDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:Dom/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_vrf_routing(vrf="VRF1")
    ```


    :param str device: A device name from the provider configuration.
    :param str vrf: VRF name.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['vrf'] = vrf
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getVrfRouting:getVrfRouting', __args__, opts=opts, typ=GetVrfRoutingResult).value

    return AwaitableGetVrfRoutingResult(
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        route_distinguisher=pulumi.get(__ret__, 'route_distinguisher'),
        vrf=pulumi.get(__ret__, 'vrf'))


@_utilities.lift_output_func(get_vrf_routing)
def get_vrf_routing_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                           vrf: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVrfRoutingResult]:
    """
    This data source can read a VRF Route Distinguisher and VRF VNI.

    - API Documentation: [rtctrlDom](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtctrl:Dom/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_vrf_routing(vrf="VRF1")
    ```


    :param str device: A device name from the provider configuration.
    :param str vrf: VRF name.
    """
    ...
