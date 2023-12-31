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
    'GetBgpRouteControlResult',
    'AwaitableGetBgpRouteControlResult',
    'get_bgp_route_control',
    'get_bgp_route_control_output',
]

@pulumi.output_type
class GetBgpRouteControlResult:
    """
    A collection of values returned by getBgpRouteControl.
    """
    def __init__(__self__, asn=None, device=None, enforce_first_as=None, fib_accelerate=None, id=None, log_neighbor_changes=None, suppress_routes=None, vrf=None):
        if asn and not isinstance(asn, str):
            raise TypeError("Expected argument 'asn' to be a str")
        pulumi.set(__self__, "asn", asn)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if enforce_first_as and not isinstance(enforce_first_as, str):
            raise TypeError("Expected argument 'enforce_first_as' to be a str")
        pulumi.set(__self__, "enforce_first_as", enforce_first_as)
        if fib_accelerate and not isinstance(fib_accelerate, str):
            raise TypeError("Expected argument 'fib_accelerate' to be a str")
        pulumi.set(__self__, "fib_accelerate", fib_accelerate)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if log_neighbor_changes and not isinstance(log_neighbor_changes, str):
            raise TypeError("Expected argument 'log_neighbor_changes' to be a str")
        pulumi.set(__self__, "log_neighbor_changes", log_neighbor_changes)
        if suppress_routes and not isinstance(suppress_routes, str):
            raise TypeError("Expected argument 'suppress_routes' to be a str")
        pulumi.set(__self__, "suppress_routes", suppress_routes)
        if vrf and not isinstance(vrf, str):
            raise TypeError("Expected argument 'vrf' to be a str")
        pulumi.set(__self__, "vrf", vrf)

    @property
    @pulumi.getter
    def asn(self) -> str:
        """
        Autonomous system number.
        """
        return pulumi.get(self, "asn")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="enforceFirstAs")
    def enforce_first_as(self) -> str:
        """
        Enforce First AS For Ebgp. Can be configured only for VRF default.
        """
        return pulumi.get(self, "enforce_first_as")

    @property
    @pulumi.getter(name="fibAccelerate")
    def fib_accelerate(self) -> str:
        """
        Accelerate the hardware updates for IP/IPv6 adjacencies for neighbor. Can be configured only for VRF default.
        """
        return pulumi.get(self, "fib_accelerate")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The distinguished name of the object.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="logNeighborChanges")
    def log_neighbor_changes(self) -> str:
        """
        Log Neighbor Changes.
        """
        return pulumi.get(self, "log_neighbor_changes")

    @property
    @pulumi.getter(name="suppressRoutes")
    def suppress_routes(self) -> str:
        """
        Suppress Routes: Advertise only routes that are programmed in hardware to peers. Can be configured only for VRF default.
        """
        return pulumi.get(self, "suppress_routes")

    @property
    @pulumi.getter
    def vrf(self) -> str:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf")


class AwaitableGetBgpRouteControlResult(GetBgpRouteControlResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBgpRouteControlResult(
            asn=self.asn,
            device=self.device,
            enforce_first_as=self.enforce_first_as,
            fib_accelerate=self.fib_accelerate,
            id=self.id,
            log_neighbor_changes=self.log_neighbor_changes,
            suppress_routes=self.suppress_routes,
            vrf=self.vrf)


def get_bgp_route_control(asn: Optional[str] = None,
                          device: Optional[str] = None,
                          vrf: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBgpRouteControlResult:
    """
    This data source can read the BGP Route Control configuration.

    - API Documentation: [bgpRtCtrl](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:RtCtrl/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_bgp_route_control(asn="65001",
        vrf="default")
    ```


    :param str asn: Autonomous system number.
    :param str device: A device name from the provider configuration.
    :param str vrf: VRF name.
    """
    __args__ = dict()
    __args__['asn'] = asn
    __args__['device'] = device
    __args__['vrf'] = vrf
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getBgpRouteControl:getBgpRouteControl', __args__, opts=opts, typ=GetBgpRouteControlResult).value

    return AwaitableGetBgpRouteControlResult(
        asn=pulumi.get(__ret__, 'asn'),
        device=pulumi.get(__ret__, 'device'),
        enforce_first_as=pulumi.get(__ret__, 'enforce_first_as'),
        fib_accelerate=pulumi.get(__ret__, 'fib_accelerate'),
        id=pulumi.get(__ret__, 'id'),
        log_neighbor_changes=pulumi.get(__ret__, 'log_neighbor_changes'),
        suppress_routes=pulumi.get(__ret__, 'suppress_routes'),
        vrf=pulumi.get(__ret__, 'vrf'))


@_utilities.lift_output_func(get_bgp_route_control)
def get_bgp_route_control_output(asn: Optional[pulumi.Input[str]] = None,
                                 device: Optional[pulumi.Input[Optional[str]]] = None,
                                 vrf: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBgpRouteControlResult]:
    """
    This data source can read the BGP Route Control configuration.

    - API Documentation: [bgpRtCtrl](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:RtCtrl/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_bgp_route_control(asn="65001",
        vrf="default")
    ```


    :param str asn: Autonomous system number.
    :param str device: A device name from the provider configuration.
    :param str vrf: VRF name.
    """
    ...
