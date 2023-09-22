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
    'GetBgpPeerTemplateResult',
    'AwaitableGetBgpPeerTemplateResult',
    'get_bgp_peer_template',
    'get_bgp_peer_template_output',
]

@pulumi.output_type
class GetBgpPeerTemplateResult:
    """
    A collection of values returned by getBgpPeerTemplate.
    """
    def __init__(__self__, asn=None, description=None, device=None, id=None, peer_type=None, remote_asn=None, source_interface=None, template_name=None):
        if asn and not isinstance(asn, str):
            raise TypeError("Expected argument 'asn' to be a str")
        pulumi.set(__self__, "asn", asn)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if peer_type and not isinstance(peer_type, str):
            raise TypeError("Expected argument 'peer_type' to be a str")
        pulumi.set(__self__, "peer_type", peer_type)
        if remote_asn and not isinstance(remote_asn, str):
            raise TypeError("Expected argument 'remote_asn' to be a str")
        pulumi.set(__self__, "remote_asn", remote_asn)
        if source_interface and not isinstance(source_interface, str):
            raise TypeError("Expected argument 'source_interface' to be a str")
        pulumi.set(__self__, "source_interface", source_interface)
        if template_name and not isinstance(template_name, str):
            raise TypeError("Expected argument 'template_name' to be a str")
        pulumi.set(__self__, "template_name", template_name)

    @property
    @pulumi.getter
    def asn(self) -> str:
        """
        Autonomous system number.
        """
        return pulumi.get(self, "asn")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Peer template description.
        """
        return pulumi.get(self, "description")

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
    @pulumi.getter(name="peerType")
    def peer_type(self) -> str:
        """
        Neighbor Fabric Type.
        """
        return pulumi.get(self, "peer_type")

    @property
    @pulumi.getter(name="remoteAsn")
    def remote_asn(self) -> str:
        """
        Peer template autonomous system number.
        """
        return pulumi.get(self, "remote_asn")

    @property
    @pulumi.getter(name="sourceInterface")
    def source_interface(self) -> str:
        """
        Source Interface. Must match first field in the output of `show intf brief`.
        """
        return pulumi.get(self, "source_interface")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> str:
        """
        Peer template name.
        """
        return pulumi.get(self, "template_name")


class AwaitableGetBgpPeerTemplateResult(GetBgpPeerTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBgpPeerTemplateResult(
            asn=self.asn,
            description=self.description,
            device=self.device,
            id=self.id,
            peer_type=self.peer_type,
            remote_asn=self.remote_asn,
            source_interface=self.source_interface,
            template_name=self.template_name)


def get_bgp_peer_template(asn: Optional[str] = None,
                          device: Optional[str] = None,
                          template_name: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBgpPeerTemplateResult:
    """
    This data source can read the BGP peer template configuration.

    - API Documentation: [bgpPeerCont](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:PeerCont/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_bgp_peer_template(asn="65001",
        template_name="SPINE-PEERS")
    ```


    :param str asn: Autonomous system number.
    :param str device: A device name from the provider configuration.
    :param str template_name: Peer template name.
    """
    __args__ = dict()
    __args__['asn'] = asn
    __args__['device'] = device
    __args__['templateName'] = template_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getBgpPeerTemplate:getBgpPeerTemplate', __args__, opts=opts, typ=GetBgpPeerTemplateResult).value

    return AwaitableGetBgpPeerTemplateResult(
        asn=pulumi.get(__ret__, 'asn'),
        description=pulumi.get(__ret__, 'description'),
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        peer_type=pulumi.get(__ret__, 'peer_type'),
        remote_asn=pulumi.get(__ret__, 'remote_asn'),
        source_interface=pulumi.get(__ret__, 'source_interface'),
        template_name=pulumi.get(__ret__, 'template_name'))


@_utilities.lift_output_func(get_bgp_peer_template)
def get_bgp_peer_template_output(asn: Optional[pulumi.Input[str]] = None,
                                 device: Optional[pulumi.Input[Optional[str]]] = None,
                                 template_name: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBgpPeerTemplateResult]:
    """
    This data source can read the BGP peer template configuration.

    - API Documentation: [bgpPeerCont](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:PeerCont/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_bgp_peer_template(asn="65001",
        template_name="SPINE-PEERS")
    ```


    :param str asn: Autonomous system number.
    :param str device: A device name from the provider configuration.
    :param str template_name: Peer template name.
    """
    ...