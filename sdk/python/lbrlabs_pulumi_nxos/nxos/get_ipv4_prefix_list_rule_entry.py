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
    'GetIpv4PrefixListRuleEntryResult',
    'AwaitableGetIpv4PrefixListRuleEntryResult',
    'get_ipv4_prefix_list_rule_entry',
    'get_ipv4_prefix_list_rule_entry_output',
]

@pulumi.output_type
class GetIpv4PrefixListRuleEntryResult:
    """
    A collection of values returned by getIpv4PrefixListRuleEntry.
    """
    def __init__(__self__, action=None, criteria=None, device=None, from_range=None, id=None, order=None, prefix=None, rule_name=None, to_range=None):
        if action and not isinstance(action, str):
            raise TypeError("Expected argument 'action' to be a str")
        pulumi.set(__self__, "action", action)
        if criteria and not isinstance(criteria, str):
            raise TypeError("Expected argument 'criteria' to be a str")
        pulumi.set(__self__, "criteria", criteria)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if from_range and not isinstance(from_range, int):
            raise TypeError("Expected argument 'from_range' to be a int")
        pulumi.set(__self__, "from_range", from_range)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if order and not isinstance(order, int):
            raise TypeError("Expected argument 'order' to be a int")
        pulumi.set(__self__, "order", order)
        if prefix and not isinstance(prefix, str):
            raise TypeError("Expected argument 'prefix' to be a str")
        pulumi.set(__self__, "prefix", prefix)
        if rule_name and not isinstance(rule_name, str):
            raise TypeError("Expected argument 'rule_name' to be a str")
        pulumi.set(__self__, "rule_name", rule_name)
        if to_range and not isinstance(to_range, int):
            raise TypeError("Expected argument 'to_range' to be a int")
        pulumi.set(__self__, "to_range", to_range)

    @property
    @pulumi.getter
    def action(self) -> str:
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def criteria(self) -> str:
        return pulumi.get(self, "criteria")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="fromRange")
    def from_range(self) -> int:
        return pulumi.get(self, "from_range")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def order(self) -> int:
        return pulumi.get(self, "order")

    @property
    @pulumi.getter
    def prefix(self) -> str:
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> str:
        return pulumi.get(self, "rule_name")

    @property
    @pulumi.getter(name="toRange")
    def to_range(self) -> int:
        return pulumi.get(self, "to_range")


class AwaitableGetIpv4PrefixListRuleEntryResult(GetIpv4PrefixListRuleEntryResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpv4PrefixListRuleEntryResult(
            action=self.action,
            criteria=self.criteria,
            device=self.device,
            from_range=self.from_range,
            id=self.id,
            order=self.order,
            prefix=self.prefix,
            rule_name=self.rule_name,
            to_range=self.to_range)


def get_ipv4_prefix_list_rule_entry(device: Optional[str] = None,
                                    order: Optional[int] = None,
                                    rule_name: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpv4PrefixListRuleEntryResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['order'] = order
    __args__['ruleName'] = rule_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:nxos/getIpv4PrefixListRuleEntry:getIpv4PrefixListRuleEntry', __args__, opts=opts, typ=GetIpv4PrefixListRuleEntryResult).value

    return AwaitableGetIpv4PrefixListRuleEntryResult(
        action=pulumi.get(__ret__, 'action'),
        criteria=pulumi.get(__ret__, 'criteria'),
        device=pulumi.get(__ret__, 'device'),
        from_range=pulumi.get(__ret__, 'from_range'),
        id=pulumi.get(__ret__, 'id'),
        order=pulumi.get(__ret__, 'order'),
        prefix=pulumi.get(__ret__, 'prefix'),
        rule_name=pulumi.get(__ret__, 'rule_name'),
        to_range=pulumi.get(__ret__, 'to_range'))


@_utilities.lift_output_func(get_ipv4_prefix_list_rule_entry)
def get_ipv4_prefix_list_rule_entry_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                                           order: Optional[pulumi.Input[int]] = None,
                                           rule_name: Optional[pulumi.Input[str]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIpv4PrefixListRuleEntryResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
