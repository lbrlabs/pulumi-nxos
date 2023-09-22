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
    'GetDefaultQosPolicyInterfaceInPolicyMapResult',
    'AwaitableGetDefaultQosPolicyInterfaceInPolicyMapResult',
    'get_default_qos_policy_interface_in_policy_map',
    'get_default_qos_policy_interface_in_policy_map_output',
]

@pulumi.output_type
class GetDefaultQosPolicyInterfaceInPolicyMapResult:
    """
    A collection of values returned by getDefaultQosPolicyInterfaceInPolicyMap.
    """
    def __init__(__self__, device=None, id=None, interface_id=None, policy_map_name=None):
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface_id and not isinstance(interface_id, str):
            raise TypeError("Expected argument 'interface_id' to be a str")
        pulumi.set(__self__, "interface_id", interface_id)
        if policy_map_name and not isinstance(policy_map_name, str):
            raise TypeError("Expected argument 'policy_map_name' to be a str")
        pulumi.set(__self__, "policy_map_name", policy_map_name)

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
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> str:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "interface_id")

    @property
    @pulumi.getter(name="policyMapName")
    def policy_map_name(self) -> str:
        """
        Policy map name.
        """
        return pulumi.get(self, "policy_map_name")


class AwaitableGetDefaultQosPolicyInterfaceInPolicyMapResult(GetDefaultQosPolicyInterfaceInPolicyMapResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDefaultQosPolicyInterfaceInPolicyMapResult(
            device=self.device,
            id=self.id,
            interface_id=self.interface_id,
            policy_map_name=self.policy_map_name)


def get_default_qos_policy_interface_in_policy_map(device: Optional[str] = None,
                                                   interface_id: Optional[str] = None,
                                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDefaultQosPolicyInterfaceInPolicyMapResult:
    """
    This data source can read the default QoS policy interface in policy map configuration.

    - API Documentation: [ipqosInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:Inst/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_default_qos_policy_interface_in_policy_map(interface_id="eth1/10")
    ```


    :param str device: A device name from the provider configuration.
    :param str interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['interfaceId'] = interface_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getDefaultQosPolicyInterfaceInPolicyMap:getDefaultQosPolicyInterfaceInPolicyMap', __args__, opts=opts, typ=GetDefaultQosPolicyInterfaceInPolicyMapResult).value

    return AwaitableGetDefaultQosPolicyInterfaceInPolicyMapResult(
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        interface_id=pulumi.get(__ret__, 'interface_id'),
        policy_map_name=pulumi.get(__ret__, 'policy_map_name'))


@_utilities.lift_output_func(get_default_qos_policy_interface_in_policy_map)
def get_default_qos_policy_interface_in_policy_map_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                                                          interface_id: Optional[pulumi.Input[str]] = None,
                                                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDefaultQosPolicyInterfaceInPolicyMapResult]:
    """
    This data source can read the default QoS policy interface in policy map configuration.

    - API Documentation: [ipqosInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:Inst/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_default_qos_policy_interface_in_policy_map(interface_id="eth1/10")
    ```


    :param str device: A device name from the provider configuration.
    :param str interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
    """
    ...