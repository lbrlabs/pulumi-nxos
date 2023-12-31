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
    'GetSubinterfaceResult',
    'AwaitableGetSubinterfaceResult',
    'get_subinterface',
    'get_subinterface_output',
]

@pulumi.output_type
class GetSubinterfaceResult:
    """
    A collection of values returned by getSubinterface.
    """
    def __init__(__self__, admin_state=None, bandwidth=None, delay=None, description=None, device=None, encap=None, id=None, interface_id=None, link_logging=None, medium=None, mtu=None):
        if admin_state and not isinstance(admin_state, str):
            raise TypeError("Expected argument 'admin_state' to be a str")
        pulumi.set(__self__, "admin_state", admin_state)
        if bandwidth and not isinstance(bandwidth, int):
            raise TypeError("Expected argument 'bandwidth' to be a int")
        pulumi.set(__self__, "bandwidth", bandwidth)
        if delay and not isinstance(delay, int):
            raise TypeError("Expected argument 'delay' to be a int")
        pulumi.set(__self__, "delay", delay)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if encap and not isinstance(encap, str):
            raise TypeError("Expected argument 'encap' to be a str")
        pulumi.set(__self__, "encap", encap)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface_id and not isinstance(interface_id, str):
            raise TypeError("Expected argument 'interface_id' to be a str")
        pulumi.set(__self__, "interface_id", interface_id)
        if link_logging and not isinstance(link_logging, str):
            raise TypeError("Expected argument 'link_logging' to be a str")
        pulumi.set(__self__, "link_logging", link_logging)
        if medium and not isinstance(medium, str):
            raise TypeError("Expected argument 'medium' to be a str")
        pulumi.set(__self__, "medium", medium)
        if mtu and not isinstance(mtu, int):
            raise TypeError("Expected argument 'mtu' to be a int")
        pulumi.set(__self__, "mtu", mtu)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> str:
        """
        Administrative state.
        """
        return pulumi.get(self, "admin_state")

    @property
    @pulumi.getter
    def bandwidth(self) -> int:
        """
        Specifies the administrative port bandwidth.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter
    def delay(self) -> int:
        """
        Specifies the administrative port delay.
        """
        return pulumi.get(self, "delay")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Interface description.
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
    def encap(self) -> str:
        """
        Subinterface encapsulation. Possible values are `unknown`, `vlan-XX` or `vxlan-XX`.
        """
        return pulumi.get(self, "encap")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The distinguished name of the object.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> str:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
        """
        return pulumi.get(self, "interface_id")

    @property
    @pulumi.getter(name="linkLogging")
    def link_logging(self) -> str:
        """
        Administrative link logging.
        """
        return pulumi.get(self, "link_logging")

    @property
    @pulumi.getter
    def medium(self) -> str:
        """
        The administrative port medium type.
        """
        return pulumi.get(self, "medium")

    @property
    @pulumi.getter
    def mtu(self) -> int:
        """
        Administrative port MTU.
        """
        return pulumi.get(self, "mtu")


class AwaitableGetSubinterfaceResult(GetSubinterfaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSubinterfaceResult(
            admin_state=self.admin_state,
            bandwidth=self.bandwidth,
            delay=self.delay,
            description=self.description,
            device=self.device,
            encap=self.encap,
            id=self.id,
            interface_id=self.interface_id,
            link_logging=self.link_logging,
            medium=self.medium,
            mtu=self.mtu)


def get_subinterface(device: Optional[str] = None,
                     interface_id: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSubinterfaceResult:
    """
    This data source can read a subinterface.

    - API Documentation: [l3EncRtdIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/l3:EncRtdIf/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_subinterface(interface_id="eth1/10.124")
    ```


    :param str device: A device name from the provider configuration.
    :param str interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['interfaceId'] = interface_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getSubinterface:getSubinterface', __args__, opts=opts, typ=GetSubinterfaceResult).value

    return AwaitableGetSubinterfaceResult(
        admin_state=pulumi.get(__ret__, 'admin_state'),
        bandwidth=pulumi.get(__ret__, 'bandwidth'),
        delay=pulumi.get(__ret__, 'delay'),
        description=pulumi.get(__ret__, 'description'),
        device=pulumi.get(__ret__, 'device'),
        encap=pulumi.get(__ret__, 'encap'),
        id=pulumi.get(__ret__, 'id'),
        interface_id=pulumi.get(__ret__, 'interface_id'),
        link_logging=pulumi.get(__ret__, 'link_logging'),
        medium=pulumi.get(__ret__, 'medium'),
        mtu=pulumi.get(__ret__, 'mtu'))


@_utilities.lift_output_func(get_subinterface)
def get_subinterface_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                            interface_id: Optional[pulumi.Input[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSubinterfaceResult]:
    """
    This data source can read a subinterface.

    - API Documentation: [l3EncRtdIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/l3:EncRtdIf/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_subinterface(interface_id="eth1/10.124")
    ```


    :param str device: A device name from the provider configuration.
    :param str interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1.10`.
    """
    ...
