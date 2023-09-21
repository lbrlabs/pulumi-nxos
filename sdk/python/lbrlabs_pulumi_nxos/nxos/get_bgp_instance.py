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
    'GetBgpInstanceResult',
    'AwaitableGetBgpInstanceResult',
    'get_bgp_instance',
    'get_bgp_instance_output',
]

@pulumi.output_type
class GetBgpInstanceResult:
    """
    A collection of values returned by getBgpInstance.
    """
    def __init__(__self__, admin_state=None, asn=None, device=None, enhanced_error_handling=None, id=None):
        if admin_state and not isinstance(admin_state, str):
            raise TypeError("Expected argument 'admin_state' to be a str")
        pulumi.set(__self__, "admin_state", admin_state)
        if asn and not isinstance(asn, str):
            raise TypeError("Expected argument 'asn' to be a str")
        pulumi.set(__self__, "asn", asn)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if enhanced_error_handling and not isinstance(enhanced_error_handling, bool):
            raise TypeError("Expected argument 'enhanced_error_handling' to be a bool")
        pulumi.set(__self__, "enhanced_error_handling", enhanced_error_handling)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> str:
        return pulumi.get(self, "admin_state")

    @property
    @pulumi.getter
    def asn(self) -> str:
        return pulumi.get(self, "asn")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="enhancedErrorHandling")
    def enhanced_error_handling(self) -> bool:
        return pulumi.get(self, "enhanced_error_handling")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")


class AwaitableGetBgpInstanceResult(GetBgpInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBgpInstanceResult(
            admin_state=self.admin_state,
            asn=self.asn,
            device=self.device,
            enhanced_error_handling=self.enhanced_error_handling,
            id=self.id)


def get_bgp_instance(device: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBgpInstanceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['device'] = device
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:nxos/getBgpInstance:getBgpInstance', __args__, opts=opts, typ=GetBgpInstanceResult).value

    return AwaitableGetBgpInstanceResult(
        admin_state=pulumi.get(__ret__, 'admin_state'),
        asn=pulumi.get(__ret__, 'asn'),
        device=pulumi.get(__ret__, 'device'),
        enhanced_error_handling=pulumi.get(__ret__, 'enhanced_error_handling'),
        id=pulumi.get(__ret__, 'id'))


@_utilities.lift_output_func(get_bgp_instance)
def get_bgp_instance_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBgpInstanceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
