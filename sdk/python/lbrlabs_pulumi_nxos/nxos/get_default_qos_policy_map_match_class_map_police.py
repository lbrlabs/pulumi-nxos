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
    'GetDefaultQosPolicyMapMatchClassMapPoliceResult',
    'AwaitableGetDefaultQosPolicyMapMatchClassMapPoliceResult',
    'get_default_qos_policy_map_match_class_map_police',
    'get_default_qos_policy_map_match_class_map_police_output',
]

@pulumi.output_type
class GetDefaultQosPolicyMapMatchClassMapPoliceResult:
    """
    A collection of values returned by getDefaultQosPolicyMapMatchClassMapPolice.
    """
    def __init__(__self__, bc_rate=None, bc_unit=None, be_rate=None, be_unit=None, cir_rate=None, cir_unit=None, class_map_name=None, conform_action=None, conform_set_cos=None, conform_set_dscp=None, conform_set_precedence=None, conform_set_qos_group=None, device=None, exceed_action=None, exceed_set_cos=None, exceed_set_dscp=None, exceed_set_precedence=None, exceed_set_qos_group=None, id=None, pir_rate=None, pir_unit=None, policy_map_name=None, violate_action=None, violate_set_cos=None, violate_set_dscp=None, violate_set_precedence=None, violate_set_qos_group=None):
        if bc_rate and not isinstance(bc_rate, int):
            raise TypeError("Expected argument 'bc_rate' to be a int")
        pulumi.set(__self__, "bc_rate", bc_rate)
        if bc_unit and not isinstance(bc_unit, str):
            raise TypeError("Expected argument 'bc_unit' to be a str")
        pulumi.set(__self__, "bc_unit", bc_unit)
        if be_rate and not isinstance(be_rate, int):
            raise TypeError("Expected argument 'be_rate' to be a int")
        pulumi.set(__self__, "be_rate", be_rate)
        if be_unit and not isinstance(be_unit, str):
            raise TypeError("Expected argument 'be_unit' to be a str")
        pulumi.set(__self__, "be_unit", be_unit)
        if cir_rate and not isinstance(cir_rate, int):
            raise TypeError("Expected argument 'cir_rate' to be a int")
        pulumi.set(__self__, "cir_rate", cir_rate)
        if cir_unit and not isinstance(cir_unit, str):
            raise TypeError("Expected argument 'cir_unit' to be a str")
        pulumi.set(__self__, "cir_unit", cir_unit)
        if class_map_name and not isinstance(class_map_name, str):
            raise TypeError("Expected argument 'class_map_name' to be a str")
        pulumi.set(__self__, "class_map_name", class_map_name)
        if conform_action and not isinstance(conform_action, str):
            raise TypeError("Expected argument 'conform_action' to be a str")
        pulumi.set(__self__, "conform_action", conform_action)
        if conform_set_cos and not isinstance(conform_set_cos, int):
            raise TypeError("Expected argument 'conform_set_cos' to be a int")
        pulumi.set(__self__, "conform_set_cos", conform_set_cos)
        if conform_set_dscp and not isinstance(conform_set_dscp, int):
            raise TypeError("Expected argument 'conform_set_dscp' to be a int")
        pulumi.set(__self__, "conform_set_dscp", conform_set_dscp)
        if conform_set_precedence and not isinstance(conform_set_precedence, str):
            raise TypeError("Expected argument 'conform_set_precedence' to be a str")
        pulumi.set(__self__, "conform_set_precedence", conform_set_precedence)
        if conform_set_qos_group and not isinstance(conform_set_qos_group, int):
            raise TypeError("Expected argument 'conform_set_qos_group' to be a int")
        pulumi.set(__self__, "conform_set_qos_group", conform_set_qos_group)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if exceed_action and not isinstance(exceed_action, str):
            raise TypeError("Expected argument 'exceed_action' to be a str")
        pulumi.set(__self__, "exceed_action", exceed_action)
        if exceed_set_cos and not isinstance(exceed_set_cos, int):
            raise TypeError("Expected argument 'exceed_set_cos' to be a int")
        pulumi.set(__self__, "exceed_set_cos", exceed_set_cos)
        if exceed_set_dscp and not isinstance(exceed_set_dscp, int):
            raise TypeError("Expected argument 'exceed_set_dscp' to be a int")
        pulumi.set(__self__, "exceed_set_dscp", exceed_set_dscp)
        if exceed_set_precedence and not isinstance(exceed_set_precedence, str):
            raise TypeError("Expected argument 'exceed_set_precedence' to be a str")
        pulumi.set(__self__, "exceed_set_precedence", exceed_set_precedence)
        if exceed_set_qos_group and not isinstance(exceed_set_qos_group, int):
            raise TypeError("Expected argument 'exceed_set_qos_group' to be a int")
        pulumi.set(__self__, "exceed_set_qos_group", exceed_set_qos_group)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if pir_rate and not isinstance(pir_rate, int):
            raise TypeError("Expected argument 'pir_rate' to be a int")
        pulumi.set(__self__, "pir_rate", pir_rate)
        if pir_unit and not isinstance(pir_unit, str):
            raise TypeError("Expected argument 'pir_unit' to be a str")
        pulumi.set(__self__, "pir_unit", pir_unit)
        if policy_map_name and not isinstance(policy_map_name, str):
            raise TypeError("Expected argument 'policy_map_name' to be a str")
        pulumi.set(__self__, "policy_map_name", policy_map_name)
        if violate_action and not isinstance(violate_action, str):
            raise TypeError("Expected argument 'violate_action' to be a str")
        pulumi.set(__self__, "violate_action", violate_action)
        if violate_set_cos and not isinstance(violate_set_cos, int):
            raise TypeError("Expected argument 'violate_set_cos' to be a int")
        pulumi.set(__self__, "violate_set_cos", violate_set_cos)
        if violate_set_dscp and not isinstance(violate_set_dscp, int):
            raise TypeError("Expected argument 'violate_set_dscp' to be a int")
        pulumi.set(__self__, "violate_set_dscp", violate_set_dscp)
        if violate_set_precedence and not isinstance(violate_set_precedence, str):
            raise TypeError("Expected argument 'violate_set_precedence' to be a str")
        pulumi.set(__self__, "violate_set_precedence", violate_set_precedence)
        if violate_set_qos_group and not isinstance(violate_set_qos_group, int):
            raise TypeError("Expected argument 'violate_set_qos_group' to be a int")
        pulumi.set(__self__, "violate_set_qos_group", violate_set_qos_group)

    @property
    @pulumi.getter(name="bcRate")
    def bc_rate(self) -> int:
        return pulumi.get(self, "bc_rate")

    @property
    @pulumi.getter(name="bcUnit")
    def bc_unit(self) -> str:
        return pulumi.get(self, "bc_unit")

    @property
    @pulumi.getter(name="beRate")
    def be_rate(self) -> int:
        return pulumi.get(self, "be_rate")

    @property
    @pulumi.getter(name="beUnit")
    def be_unit(self) -> str:
        return pulumi.get(self, "be_unit")

    @property
    @pulumi.getter(name="cirRate")
    def cir_rate(self) -> int:
        return pulumi.get(self, "cir_rate")

    @property
    @pulumi.getter(name="cirUnit")
    def cir_unit(self) -> str:
        return pulumi.get(self, "cir_unit")

    @property
    @pulumi.getter(name="classMapName")
    def class_map_name(self) -> str:
        return pulumi.get(self, "class_map_name")

    @property
    @pulumi.getter(name="conformAction")
    def conform_action(self) -> str:
        return pulumi.get(self, "conform_action")

    @property
    @pulumi.getter(name="conformSetCos")
    def conform_set_cos(self) -> int:
        return pulumi.get(self, "conform_set_cos")

    @property
    @pulumi.getter(name="conformSetDscp")
    def conform_set_dscp(self) -> int:
        return pulumi.get(self, "conform_set_dscp")

    @property
    @pulumi.getter(name="conformSetPrecedence")
    def conform_set_precedence(self) -> str:
        return pulumi.get(self, "conform_set_precedence")

    @property
    @pulumi.getter(name="conformSetQosGroup")
    def conform_set_qos_group(self) -> int:
        return pulumi.get(self, "conform_set_qos_group")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="exceedAction")
    def exceed_action(self) -> str:
        return pulumi.get(self, "exceed_action")

    @property
    @pulumi.getter(name="exceedSetCos")
    def exceed_set_cos(self) -> int:
        return pulumi.get(self, "exceed_set_cos")

    @property
    @pulumi.getter(name="exceedSetDscp")
    def exceed_set_dscp(self) -> int:
        return pulumi.get(self, "exceed_set_dscp")

    @property
    @pulumi.getter(name="exceedSetPrecedence")
    def exceed_set_precedence(self) -> str:
        return pulumi.get(self, "exceed_set_precedence")

    @property
    @pulumi.getter(name="exceedSetQosGroup")
    def exceed_set_qos_group(self) -> int:
        return pulumi.get(self, "exceed_set_qos_group")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="pirRate")
    def pir_rate(self) -> int:
        return pulumi.get(self, "pir_rate")

    @property
    @pulumi.getter(name="pirUnit")
    def pir_unit(self) -> str:
        return pulumi.get(self, "pir_unit")

    @property
    @pulumi.getter(name="policyMapName")
    def policy_map_name(self) -> str:
        return pulumi.get(self, "policy_map_name")

    @property
    @pulumi.getter(name="violateAction")
    def violate_action(self) -> str:
        return pulumi.get(self, "violate_action")

    @property
    @pulumi.getter(name="violateSetCos")
    def violate_set_cos(self) -> int:
        return pulumi.get(self, "violate_set_cos")

    @property
    @pulumi.getter(name="violateSetDscp")
    def violate_set_dscp(self) -> int:
        return pulumi.get(self, "violate_set_dscp")

    @property
    @pulumi.getter(name="violateSetPrecedence")
    def violate_set_precedence(self) -> str:
        return pulumi.get(self, "violate_set_precedence")

    @property
    @pulumi.getter(name="violateSetQosGroup")
    def violate_set_qos_group(self) -> int:
        return pulumi.get(self, "violate_set_qos_group")


class AwaitableGetDefaultQosPolicyMapMatchClassMapPoliceResult(GetDefaultQosPolicyMapMatchClassMapPoliceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDefaultQosPolicyMapMatchClassMapPoliceResult(
            bc_rate=self.bc_rate,
            bc_unit=self.bc_unit,
            be_rate=self.be_rate,
            be_unit=self.be_unit,
            cir_rate=self.cir_rate,
            cir_unit=self.cir_unit,
            class_map_name=self.class_map_name,
            conform_action=self.conform_action,
            conform_set_cos=self.conform_set_cos,
            conform_set_dscp=self.conform_set_dscp,
            conform_set_precedence=self.conform_set_precedence,
            conform_set_qos_group=self.conform_set_qos_group,
            device=self.device,
            exceed_action=self.exceed_action,
            exceed_set_cos=self.exceed_set_cos,
            exceed_set_dscp=self.exceed_set_dscp,
            exceed_set_precedence=self.exceed_set_precedence,
            exceed_set_qos_group=self.exceed_set_qos_group,
            id=self.id,
            pir_rate=self.pir_rate,
            pir_unit=self.pir_unit,
            policy_map_name=self.policy_map_name,
            violate_action=self.violate_action,
            violate_set_cos=self.violate_set_cos,
            violate_set_dscp=self.violate_set_dscp,
            violate_set_precedence=self.violate_set_precedence,
            violate_set_qos_group=self.violate_set_qos_group)


def get_default_qos_policy_map_match_class_map_police(class_map_name: Optional[str] = None,
                                                      device: Optional[str] = None,
                                                      policy_map_name: Optional[str] = None,
                                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDefaultQosPolicyMapMatchClassMapPoliceResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['classMapName'] = class_map_name
    __args__['device'] = device
    __args__['policyMapName'] = policy_map_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:nxos/getDefaultQosPolicyMapMatchClassMapPolice:getDefaultQosPolicyMapMatchClassMapPolice', __args__, opts=opts, typ=GetDefaultQosPolicyMapMatchClassMapPoliceResult).value

    return AwaitableGetDefaultQosPolicyMapMatchClassMapPoliceResult(
        bc_rate=pulumi.get(__ret__, 'bc_rate'),
        bc_unit=pulumi.get(__ret__, 'bc_unit'),
        be_rate=pulumi.get(__ret__, 'be_rate'),
        be_unit=pulumi.get(__ret__, 'be_unit'),
        cir_rate=pulumi.get(__ret__, 'cir_rate'),
        cir_unit=pulumi.get(__ret__, 'cir_unit'),
        class_map_name=pulumi.get(__ret__, 'class_map_name'),
        conform_action=pulumi.get(__ret__, 'conform_action'),
        conform_set_cos=pulumi.get(__ret__, 'conform_set_cos'),
        conform_set_dscp=pulumi.get(__ret__, 'conform_set_dscp'),
        conform_set_precedence=pulumi.get(__ret__, 'conform_set_precedence'),
        conform_set_qos_group=pulumi.get(__ret__, 'conform_set_qos_group'),
        device=pulumi.get(__ret__, 'device'),
        exceed_action=pulumi.get(__ret__, 'exceed_action'),
        exceed_set_cos=pulumi.get(__ret__, 'exceed_set_cos'),
        exceed_set_dscp=pulumi.get(__ret__, 'exceed_set_dscp'),
        exceed_set_precedence=pulumi.get(__ret__, 'exceed_set_precedence'),
        exceed_set_qos_group=pulumi.get(__ret__, 'exceed_set_qos_group'),
        id=pulumi.get(__ret__, 'id'),
        pir_rate=pulumi.get(__ret__, 'pir_rate'),
        pir_unit=pulumi.get(__ret__, 'pir_unit'),
        policy_map_name=pulumi.get(__ret__, 'policy_map_name'),
        violate_action=pulumi.get(__ret__, 'violate_action'),
        violate_set_cos=pulumi.get(__ret__, 'violate_set_cos'),
        violate_set_dscp=pulumi.get(__ret__, 'violate_set_dscp'),
        violate_set_precedence=pulumi.get(__ret__, 'violate_set_precedence'),
        violate_set_qos_group=pulumi.get(__ret__, 'violate_set_qos_group'))


@_utilities.lift_output_func(get_default_qos_policy_map_match_class_map_police)
def get_default_qos_policy_map_match_class_map_police_output(class_map_name: Optional[pulumi.Input[str]] = None,
                                                             device: Optional[pulumi.Input[Optional[str]]] = None,
                                                             policy_map_name: Optional[pulumi.Input[str]] = None,
                                                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDefaultQosPolicyMapMatchClassMapPoliceResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
