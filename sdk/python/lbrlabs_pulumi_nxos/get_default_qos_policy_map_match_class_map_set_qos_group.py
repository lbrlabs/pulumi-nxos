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
    'GetDefaultQosPolicyMapMatchClassMapSetQosGroupResult',
    'AwaitableGetDefaultQosPolicyMapMatchClassMapSetQosGroupResult',
    'get_default_qos_policy_map_match_class_map_set_qos_group',
    'get_default_qos_policy_map_match_class_map_set_qos_group_output',
]

@pulumi.output_type
class GetDefaultQosPolicyMapMatchClassMapSetQosGroupResult:
    """
    A collection of values returned by getDefaultQosPolicyMapMatchClassMapSetQosGroup.
    """
    def __init__(__self__, class_map_name=None, device=None, id=None, policy_map_name=None, qos_group_id=None):
        if class_map_name and not isinstance(class_map_name, str):
            raise TypeError("Expected argument 'class_map_name' to be a str")
        pulumi.set(__self__, "class_map_name", class_map_name)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if policy_map_name and not isinstance(policy_map_name, str):
            raise TypeError("Expected argument 'policy_map_name' to be a str")
        pulumi.set(__self__, "policy_map_name", policy_map_name)
        if qos_group_id and not isinstance(qos_group_id, int):
            raise TypeError("Expected argument 'qos_group_id' to be a int")
        pulumi.set(__self__, "qos_group_id", qos_group_id)

    @property
    @pulumi.getter(name="classMapName")
    def class_map_name(self) -> str:
        """
        Class map name.
        """
        return pulumi.get(self, "class_map_name")

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
    @pulumi.getter(name="policyMapName")
    def policy_map_name(self) -> str:
        """
        Policy map name.
        """
        return pulumi.get(self, "policy_map_name")

    @property
    @pulumi.getter(name="qosGroupId")
    def qos_group_id(self) -> int:
        """
        QoS group ID.
        """
        return pulumi.get(self, "qos_group_id")


class AwaitableGetDefaultQosPolicyMapMatchClassMapSetQosGroupResult(GetDefaultQosPolicyMapMatchClassMapSetQosGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDefaultQosPolicyMapMatchClassMapSetQosGroupResult(
            class_map_name=self.class_map_name,
            device=self.device,
            id=self.id,
            policy_map_name=self.policy_map_name,
            qos_group_id=self.qos_group_id)


def get_default_qos_policy_map_match_class_map_set_qos_group(class_map_name: Optional[str] = None,
                                                             device: Optional[str] = None,
                                                             policy_map_name: Optional[str] = None,
                                                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDefaultQosPolicyMapMatchClassMapSetQosGroupResult:
    """
    This data source can read the default QoS policy map match class map set QoS group configuration.

    - API Documentation: [ipqosSetQoSGrp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:SetQoSGrp/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_default_qos_policy_map_match_class_map_set_qos_group(class_map_name="Voice",
        policy_map_name="PM1")
    ```


    :param str class_map_name: Class map name.
    :param str device: A device name from the provider configuration.
    :param str policy_map_name: Policy map name.
    """
    __args__ = dict()
    __args__['classMapName'] = class_map_name
    __args__['device'] = device
    __args__['policyMapName'] = policy_map_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getDefaultQosPolicyMapMatchClassMapSetQosGroup:getDefaultQosPolicyMapMatchClassMapSetQosGroup', __args__, opts=opts, typ=GetDefaultQosPolicyMapMatchClassMapSetQosGroupResult).value

    return AwaitableGetDefaultQosPolicyMapMatchClassMapSetQosGroupResult(
        class_map_name=pulumi.get(__ret__, 'class_map_name'),
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        policy_map_name=pulumi.get(__ret__, 'policy_map_name'),
        qos_group_id=pulumi.get(__ret__, 'qos_group_id'))


@_utilities.lift_output_func(get_default_qos_policy_map_match_class_map_set_qos_group)
def get_default_qos_policy_map_match_class_map_set_qos_group_output(class_map_name: Optional[pulumi.Input[str]] = None,
                                                                    device: Optional[pulumi.Input[Optional[str]]] = None,
                                                                    policy_map_name: Optional[pulumi.Input[str]] = None,
                                                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDefaultQosPolicyMapMatchClassMapSetQosGroupResult]:
    """
    This data source can read the default QoS policy map match class map set QoS group configuration.

    - API Documentation: [ipqosSetQoSGrp](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Qos/ipqos:SetQoSGrp/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_default_qos_policy_map_match_class_map_set_qos_group(class_map_name="Voice",
        policy_map_name="PM1")
    ```


    :param str class_map_name: Class map name.
    :param str device: A device name from the provider configuration.
    :param str policy_map_name: Policy map name.
    """
    ...
