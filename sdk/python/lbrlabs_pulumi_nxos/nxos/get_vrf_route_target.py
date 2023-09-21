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
    'GetVrfRouteTargetResult',
    'AwaitableGetVrfRouteTargetResult',
    'get_vrf_route_target',
    'get_vrf_route_target_output',
]

@pulumi.output_type
class GetVrfRouteTargetResult:
    """
    A collection of values returned by getVrfRouteTarget.
    """
    def __init__(__self__, address_family=None, device=None, direction=None, id=None, route_target=None, route_target_address_family=None, vrf=None):
        if address_family and not isinstance(address_family, str):
            raise TypeError("Expected argument 'address_family' to be a str")
        pulumi.set(__self__, "address_family", address_family)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if direction and not isinstance(direction, str):
            raise TypeError("Expected argument 'direction' to be a str")
        pulumi.set(__self__, "direction", direction)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if route_target and not isinstance(route_target, str):
            raise TypeError("Expected argument 'route_target' to be a str")
        pulumi.set(__self__, "route_target", route_target)
        if route_target_address_family and not isinstance(route_target_address_family, str):
            raise TypeError("Expected argument 'route_target_address_family' to be a str")
        pulumi.set(__self__, "route_target_address_family", route_target_address_family)
        if vrf and not isinstance(vrf, str):
            raise TypeError("Expected argument 'vrf' to be a str")
        pulumi.set(__self__, "vrf", vrf)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> str:
        return pulumi.get(self, "address_family")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def direction(self) -> str:
        return pulumi.get(self, "direction")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="routeTarget")
    def route_target(self) -> str:
        return pulumi.get(self, "route_target")

    @property
    @pulumi.getter(name="routeTargetAddressFamily")
    def route_target_address_family(self) -> str:
        return pulumi.get(self, "route_target_address_family")

    @property
    @pulumi.getter
    def vrf(self) -> str:
        return pulumi.get(self, "vrf")


class AwaitableGetVrfRouteTargetResult(GetVrfRouteTargetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVrfRouteTargetResult(
            address_family=self.address_family,
            device=self.device,
            direction=self.direction,
            id=self.id,
            route_target=self.route_target,
            route_target_address_family=self.route_target_address_family,
            vrf=self.vrf)


def get_vrf_route_target(address_family: Optional[str] = None,
                         device: Optional[str] = None,
                         direction: Optional[str] = None,
                         route_target: Optional[str] = None,
                         route_target_address_family: Optional[str] = None,
                         vrf: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVrfRouteTargetResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['addressFamily'] = address_family
    __args__['device'] = device
    __args__['direction'] = direction
    __args__['routeTarget'] = route_target
    __args__['routeTargetAddressFamily'] = route_target_address_family
    __args__['vrf'] = vrf
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:nxos/getVrfRouteTarget:getVrfRouteTarget', __args__, opts=opts, typ=GetVrfRouteTargetResult).value

    return AwaitableGetVrfRouteTargetResult(
        address_family=pulumi.get(__ret__, 'address_family'),
        device=pulumi.get(__ret__, 'device'),
        direction=pulumi.get(__ret__, 'direction'),
        id=pulumi.get(__ret__, 'id'),
        route_target=pulumi.get(__ret__, 'route_target'),
        route_target_address_family=pulumi.get(__ret__, 'route_target_address_family'),
        vrf=pulumi.get(__ret__, 'vrf'))


@_utilities.lift_output_func(get_vrf_route_target)
def get_vrf_route_target_output(address_family: Optional[pulumi.Input[str]] = None,
                                device: Optional[pulumi.Input[Optional[str]]] = None,
                                direction: Optional[pulumi.Input[str]] = None,
                                route_target: Optional[pulumi.Input[str]] = None,
                                route_target_address_family: Optional[pulumi.Input[str]] = None,
                                vrf: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVrfRouteTargetResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
