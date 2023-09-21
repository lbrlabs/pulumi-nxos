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
    'GetBgpAdvertisedPrefixResult',
    'AwaitableGetBgpAdvertisedPrefixResult',
    'get_bgp_advertised_prefix',
    'get_bgp_advertised_prefix_output',
]

@pulumi.output_type
class GetBgpAdvertisedPrefixResult:
    """
    A collection of values returned by getBgpAdvertisedPrefix.
    """
    def __init__(__self__, address_family=None, asn=None, device=None, id=None, prefix=None, route_map=None, vrf=None):
        if address_family and not isinstance(address_family, str):
            raise TypeError("Expected argument 'address_family' to be a str")
        pulumi.set(__self__, "address_family", address_family)
        if asn and not isinstance(asn, str):
            raise TypeError("Expected argument 'asn' to be a str")
        pulumi.set(__self__, "asn", asn)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if prefix and not isinstance(prefix, str):
            raise TypeError("Expected argument 'prefix' to be a str")
        pulumi.set(__self__, "prefix", prefix)
        if route_map and not isinstance(route_map, str):
            raise TypeError("Expected argument 'route_map' to be a str")
        pulumi.set(__self__, "route_map", route_map)
        if vrf and not isinstance(vrf, str):
            raise TypeError("Expected argument 'vrf' to be a str")
        pulumi.set(__self__, "vrf", vrf)

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> str:
        return pulumi.get(self, "address_family")

    @property
    @pulumi.getter
    def asn(self) -> str:
        return pulumi.get(self, "asn")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def prefix(self) -> str:
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter(name="routeMap")
    def route_map(self) -> str:
        return pulumi.get(self, "route_map")

    @property
    @pulumi.getter
    def vrf(self) -> str:
        return pulumi.get(self, "vrf")


class AwaitableGetBgpAdvertisedPrefixResult(GetBgpAdvertisedPrefixResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBgpAdvertisedPrefixResult(
            address_family=self.address_family,
            asn=self.asn,
            device=self.device,
            id=self.id,
            prefix=self.prefix,
            route_map=self.route_map,
            vrf=self.vrf)


def get_bgp_advertised_prefix(address_family: Optional[str] = None,
                              asn: Optional[str] = None,
                              device: Optional[str] = None,
                              prefix: Optional[str] = None,
                              vrf: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBgpAdvertisedPrefixResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['addressFamily'] = address_family
    __args__['asn'] = asn
    __args__['device'] = device
    __args__['prefix'] = prefix
    __args__['vrf'] = vrf
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:nxos/getBgpAdvertisedPrefix:getBgpAdvertisedPrefix', __args__, opts=opts, typ=GetBgpAdvertisedPrefixResult).value

    return AwaitableGetBgpAdvertisedPrefixResult(
        address_family=pulumi.get(__ret__, 'address_family'),
        asn=pulumi.get(__ret__, 'asn'),
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        prefix=pulumi.get(__ret__, 'prefix'),
        route_map=pulumi.get(__ret__, 'route_map'),
        vrf=pulumi.get(__ret__, 'vrf'))


@_utilities.lift_output_func(get_bgp_advertised_prefix)
def get_bgp_advertised_prefix_output(address_family: Optional[pulumi.Input[str]] = None,
                                     asn: Optional[pulumi.Input[str]] = None,
                                     device: Optional[pulumi.Input[Optional[str]]] = None,
                                     prefix: Optional[pulumi.Input[str]] = None,
                                     vrf: Optional[pulumi.Input[str]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBgpAdvertisedPrefixResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...