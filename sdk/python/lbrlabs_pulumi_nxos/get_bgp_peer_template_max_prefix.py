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
    'GetBgpPeerTemplateMaxPrefixResult',
    'AwaitableGetBgpPeerTemplateMaxPrefixResult',
    'get_bgp_peer_template_max_prefix',
    'get_bgp_peer_template_max_prefix_output',
]

@pulumi.output_type
class GetBgpPeerTemplateMaxPrefixResult:
    """
    A collection of values returned by getBgpPeerTemplateMaxPrefix.
    """
    def __init__(__self__, action=None, address_family=None, asn=None, device=None, id=None, maximum_prefix=None, restart_time=None, template_name=None, threshold=None):
        if action and not isinstance(action, str):
            raise TypeError("Expected argument 'action' to be a str")
        pulumi.set(__self__, "action", action)
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
        if maximum_prefix and not isinstance(maximum_prefix, int):
            raise TypeError("Expected argument 'maximum_prefix' to be a int")
        pulumi.set(__self__, "maximum_prefix", maximum_prefix)
        if restart_time and not isinstance(restart_time, int):
            raise TypeError("Expected argument 'restart_time' to be a int")
        pulumi.set(__self__, "restart_time", restart_time)
        if template_name and not isinstance(template_name, str):
            raise TypeError("Expected argument 'template_name' to be a str")
        pulumi.set(__self__, "template_name", template_name)
        if threshold and not isinstance(threshold, int):
            raise TypeError("Expected argument 'threshold' to be a int")
        pulumi.set(__self__, "threshold", threshold)

    @property
    @pulumi.getter
    def action(self) -> str:
        """
        Action to do when limit is exceeded.
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter(name="addressFamily")
    def address_family(self) -> str:
        """
        Address Family.
        """
        return pulumi.get(self, "address_family")

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
    @pulumi.getter
    def id(self) -> str:
        """
        The distinguished name of the object.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="maximumPrefix")
    def maximum_prefix(self) -> int:
        """
        Maximum number of prefixes allowed from the peer.
        """
        return pulumi.get(self, "maximum_prefix")

    @property
    @pulumi.getter(name="restartTime")
    def restart_time(self) -> int:
        """
        The period of time in minutes before restarting the peer when the prefix limit is reached.
        """
        return pulumi.get(self, "restart_time")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> str:
        """
        Peer template name.
        """
        return pulumi.get(self, "template_name")

    @property
    @pulumi.getter
    def threshold(self) -> int:
        """
        The period of time in minutes before restarting the peer when the prefix limit is reached.
        """
        return pulumi.get(self, "threshold")


class AwaitableGetBgpPeerTemplateMaxPrefixResult(GetBgpPeerTemplateMaxPrefixResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBgpPeerTemplateMaxPrefixResult(
            action=self.action,
            address_family=self.address_family,
            asn=self.asn,
            device=self.device,
            id=self.id,
            maximum_prefix=self.maximum_prefix,
            restart_time=self.restart_time,
            template_name=self.template_name,
            threshold=self.threshold)


def get_bgp_peer_template_max_prefix(address_family: Optional[str] = None,
                                     asn: Optional[str] = None,
                                     device: Optional[str] = None,
                                     template_name: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBgpPeerTemplateMaxPrefixResult:
    """
    This data source can read the BGP peer template Maximum Prefix Policy configuration.

    - API Documentation: [bgpMaxPfxP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:MaxPfxP/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_bgp_peer_template_max_prefix(address_family="ipv4-ucast",
        asn="65001",
        template_name="SPINE-PEERS")
    ```


    :param str address_family: Address Family.
    :param str asn: Autonomous system number.
    :param str device: A device name from the provider configuration.
    :param str template_name: Peer template name.
    """
    __args__ = dict()
    __args__['addressFamily'] = address_family
    __args__['asn'] = asn
    __args__['device'] = device
    __args__['templateName'] = template_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getBgpPeerTemplateMaxPrefix:getBgpPeerTemplateMaxPrefix', __args__, opts=opts, typ=GetBgpPeerTemplateMaxPrefixResult).value

    return AwaitableGetBgpPeerTemplateMaxPrefixResult(
        action=pulumi.get(__ret__, 'action'),
        address_family=pulumi.get(__ret__, 'address_family'),
        asn=pulumi.get(__ret__, 'asn'),
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        maximum_prefix=pulumi.get(__ret__, 'maximum_prefix'),
        restart_time=pulumi.get(__ret__, 'restart_time'),
        template_name=pulumi.get(__ret__, 'template_name'),
        threshold=pulumi.get(__ret__, 'threshold'))


@_utilities.lift_output_func(get_bgp_peer_template_max_prefix)
def get_bgp_peer_template_max_prefix_output(address_family: Optional[pulumi.Input[str]] = None,
                                            asn: Optional[pulumi.Input[str]] = None,
                                            device: Optional[pulumi.Input[Optional[str]]] = None,
                                            template_name: Optional[pulumi.Input[str]] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBgpPeerTemplateMaxPrefixResult]:
    """
    This data source can read the BGP peer template Maximum Prefix Policy configuration.

    - API Documentation: [bgpMaxPfxP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:MaxPfxP/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_bgp_peer_template_max_prefix(address_family="ipv4-ucast",
        asn="65001",
        template_name="SPINE-PEERS")
    ```


    :param str address_family: Address Family.
    :param str asn: Autonomous system number.
    :param str device: A device name from the provider configuration.
    :param str template_name: Peer template name.
    """
    ...
