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
    'GetBgpPeerTemplateAddressFamilyResult',
    'AwaitableGetBgpPeerTemplateAddressFamilyResult',
    'get_bgp_peer_template_address_family',
    'get_bgp_peer_template_address_family_output',
]

@pulumi.output_type
class GetBgpPeerTemplateAddressFamilyResult:
    """
    A collection of values returned by getBgpPeerTemplateAddressFamily.
    """
    def __init__(__self__, address_family=None, asn=None, control=None, device=None, id=None, send_community_extended=None, send_community_standard=None, template_name=None):
        if address_family and not isinstance(address_family, str):
            raise TypeError("Expected argument 'address_family' to be a str")
        pulumi.set(__self__, "address_family", address_family)
        if asn and not isinstance(asn, str):
            raise TypeError("Expected argument 'asn' to be a str")
        pulumi.set(__self__, "asn", asn)
        if control and not isinstance(control, str):
            raise TypeError("Expected argument 'control' to be a str")
        pulumi.set(__self__, "control", control)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if send_community_extended and not isinstance(send_community_extended, str):
            raise TypeError("Expected argument 'send_community_extended' to be a str")
        pulumi.set(__self__, "send_community_extended", send_community_extended)
        if send_community_standard and not isinstance(send_community_standard, str):
            raise TypeError("Expected argument 'send_community_standard' to be a str")
        pulumi.set(__self__, "send_community_standard", send_community_standard)
        if template_name and not isinstance(template_name, str):
            raise TypeError("Expected argument 'template_name' to be a str")
        pulumi.set(__self__, "template_name", template_name)

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
    def control(self) -> str:
        return pulumi.get(self, "control")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="sendCommunityExtended")
    def send_community_extended(self) -> str:
        return pulumi.get(self, "send_community_extended")

    @property
    @pulumi.getter(name="sendCommunityStandard")
    def send_community_standard(self) -> str:
        return pulumi.get(self, "send_community_standard")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> str:
        return pulumi.get(self, "template_name")


class AwaitableGetBgpPeerTemplateAddressFamilyResult(GetBgpPeerTemplateAddressFamilyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBgpPeerTemplateAddressFamilyResult(
            address_family=self.address_family,
            asn=self.asn,
            control=self.control,
            device=self.device,
            id=self.id,
            send_community_extended=self.send_community_extended,
            send_community_standard=self.send_community_standard,
            template_name=self.template_name)


def get_bgp_peer_template_address_family(address_family: Optional[str] = None,
                                         asn: Optional[str] = None,
                                         device: Optional[str] = None,
                                         template_name: Optional[str] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBgpPeerTemplateAddressFamilyResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['addressFamily'] = address_family
    __args__['asn'] = asn
    __args__['device'] = device
    __args__['templateName'] = template_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:nxos/getBgpPeerTemplateAddressFamily:getBgpPeerTemplateAddressFamily', __args__, opts=opts, typ=GetBgpPeerTemplateAddressFamilyResult).value

    return AwaitableGetBgpPeerTemplateAddressFamilyResult(
        address_family=pulumi.get(__ret__, 'address_family'),
        asn=pulumi.get(__ret__, 'asn'),
        control=pulumi.get(__ret__, 'control'),
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        send_community_extended=pulumi.get(__ret__, 'send_community_extended'),
        send_community_standard=pulumi.get(__ret__, 'send_community_standard'),
        template_name=pulumi.get(__ret__, 'template_name'))


@_utilities.lift_output_func(get_bgp_peer_template_address_family)
def get_bgp_peer_template_address_family_output(address_family: Optional[pulumi.Input[str]] = None,
                                                asn: Optional[pulumi.Input[str]] = None,
                                                device: Optional[pulumi.Input[Optional[str]]] = None,
                                                template_name: Optional[pulumi.Input[str]] = None,
                                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBgpPeerTemplateAddressFamilyResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...