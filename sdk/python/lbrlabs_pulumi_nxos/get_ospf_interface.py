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
    'GetOspfInterfaceResult',
    'AwaitableGetOspfInterfaceResult',
    'get_ospf_interface',
    'get_ospf_interface_output',
]

@pulumi.output_type
class GetOspfInterfaceResult:
    """
    A collection of values returned by getOspfInterface.
    """
    def __init__(__self__, advertise_secondaries=None, area=None, bfd=None, cost=None, dead_interval=None, device=None, hello_interval=None, id=None, instance_name=None, interface_id=None, network_type=None, passive=None, priority=None, vrf_name=None):
        if advertise_secondaries and not isinstance(advertise_secondaries, bool):
            raise TypeError("Expected argument 'advertise_secondaries' to be a bool")
        pulumi.set(__self__, "advertise_secondaries", advertise_secondaries)
        if area and not isinstance(area, str):
            raise TypeError("Expected argument 'area' to be a str")
        pulumi.set(__self__, "area", area)
        if bfd and not isinstance(bfd, str):
            raise TypeError("Expected argument 'bfd' to be a str")
        pulumi.set(__self__, "bfd", bfd)
        if cost and not isinstance(cost, int):
            raise TypeError("Expected argument 'cost' to be a int")
        pulumi.set(__self__, "cost", cost)
        if dead_interval and not isinstance(dead_interval, int):
            raise TypeError("Expected argument 'dead_interval' to be a int")
        pulumi.set(__self__, "dead_interval", dead_interval)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if hello_interval and not isinstance(hello_interval, int):
            raise TypeError("Expected argument 'hello_interval' to be a int")
        pulumi.set(__self__, "hello_interval", hello_interval)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instance_name and not isinstance(instance_name, str):
            raise TypeError("Expected argument 'instance_name' to be a str")
        pulumi.set(__self__, "instance_name", instance_name)
        if interface_id and not isinstance(interface_id, str):
            raise TypeError("Expected argument 'interface_id' to be a str")
        pulumi.set(__self__, "interface_id", interface_id)
        if network_type and not isinstance(network_type, str):
            raise TypeError("Expected argument 'network_type' to be a str")
        pulumi.set(__self__, "network_type", network_type)
        if passive and not isinstance(passive, str):
            raise TypeError("Expected argument 'passive' to be a str")
        pulumi.set(__self__, "passive", passive)
        if priority and not isinstance(priority, int):
            raise TypeError("Expected argument 'priority' to be a int")
        pulumi.set(__self__, "priority", priority)
        if vrf_name and not isinstance(vrf_name, str):
            raise TypeError("Expected argument 'vrf_name' to be a str")
        pulumi.set(__self__, "vrf_name", vrf_name)

    @property
    @pulumi.getter(name="advertiseSecondaries")
    def advertise_secondaries(self) -> bool:
        """
        Advertise secondary IP addresses.
        """
        return pulumi.get(self, "advertise_secondaries")

    @property
    @pulumi.getter
    def area(self) -> str:
        """
        Area identifier to which a network or interface belongs in IPv4 address format.
        """
        return pulumi.get(self, "area")

    @property
    @pulumi.getter
    def bfd(self) -> str:
        """
        Bidirectional Forwarding Detection (BFD).
        """
        return pulumi.get(self, "bfd")

    @property
    @pulumi.getter
    def cost(self) -> int:
        """
        Specifies the cost of interface.
        """
        return pulumi.get(self, "cost")

    @property
    @pulumi.getter(name="deadInterval")
    def dead_interval(self) -> int:
        """
        Dead interval, interval after which router declares that neighbor as down.
        """
        return pulumi.get(self, "dead_interval")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="helloInterval")
    def hello_interval(self) -> int:
        """
        Hello interval, interval between hello packets that OSPF sends on the interface.
        """
        return pulumi.get(self, "hello_interval")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The distinguished name of the object.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instanceName")
    def instance_name(self) -> str:
        """
        OSPF instance name.
        """
        return pulumi.get(self, "instance_name")

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> str:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "interface_id")

    @property
    @pulumi.getter(name="networkType")
    def network_type(self) -> str:
        """
        Network type.
        """
        return pulumi.get(self, "network_type")

    @property
    @pulumi.getter
    def passive(self) -> str:
        """
        Passive interface control. Interface can be configured as passive or non-passive.
        """
        return pulumi.get(self, "passive")

    @property
    @pulumi.getter
    def priority(self) -> int:
        """
        Priority, used in determining the designated router on this network.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="vrfName")
    def vrf_name(self) -> str:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf_name")


class AwaitableGetOspfInterfaceResult(GetOspfInterfaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOspfInterfaceResult(
            advertise_secondaries=self.advertise_secondaries,
            area=self.area,
            bfd=self.bfd,
            cost=self.cost,
            dead_interval=self.dead_interval,
            device=self.device,
            hello_interval=self.hello_interval,
            id=self.id,
            instance_name=self.instance_name,
            interface_id=self.interface_id,
            network_type=self.network_type,
            passive=self.passive,
            priority=self.priority,
            vrf_name=self.vrf_name)


def get_ospf_interface(device: Optional[str] = None,
                       instance_name: Optional[str] = None,
                       interface_id: Optional[str] = None,
                       vrf_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOspfInterfaceResult:
    """
    This data source can read the OSPF interface configuration.

    - API Documentation: [ospfIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/ospf:If/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_ospf_interface(instance_name="OSPF1",
        interface_id="eth1/10",
        vrf_name="VRF1")
    ```


    :param str device: A device name from the provider configuration.
    :param str instance_name: OSPF instance name.
    :param str interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
    :param str vrf_name: VRF name.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['instanceName'] = instance_name
    __args__['interfaceId'] = interface_id
    __args__['vrfName'] = vrf_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getOspfInterface:getOspfInterface', __args__, opts=opts, typ=GetOspfInterfaceResult).value

    return AwaitableGetOspfInterfaceResult(
        advertise_secondaries=pulumi.get(__ret__, 'advertise_secondaries'),
        area=pulumi.get(__ret__, 'area'),
        bfd=pulumi.get(__ret__, 'bfd'),
        cost=pulumi.get(__ret__, 'cost'),
        dead_interval=pulumi.get(__ret__, 'dead_interval'),
        device=pulumi.get(__ret__, 'device'),
        hello_interval=pulumi.get(__ret__, 'hello_interval'),
        id=pulumi.get(__ret__, 'id'),
        instance_name=pulumi.get(__ret__, 'instance_name'),
        interface_id=pulumi.get(__ret__, 'interface_id'),
        network_type=pulumi.get(__ret__, 'network_type'),
        passive=pulumi.get(__ret__, 'passive'),
        priority=pulumi.get(__ret__, 'priority'),
        vrf_name=pulumi.get(__ret__, 'vrf_name'))


@_utilities.lift_output_func(get_ospf_interface)
def get_ospf_interface_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                              instance_name: Optional[pulumi.Input[str]] = None,
                              interface_id: Optional[pulumi.Input[str]] = None,
                              vrf_name: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOspfInterfaceResult]:
    """
    This data source can read the OSPF interface configuration.

    - API Documentation: [ospfIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/ospf:If/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_ospf_interface(instance_name="OSPF1",
        interface_id="eth1/10",
        vrf_name="VRF1")
    ```


    :param str device: A device name from the provider configuration.
    :param str instance_name: OSPF instance name.
    :param str interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
    :param str vrf_name: VRF name.
    """
    ...
