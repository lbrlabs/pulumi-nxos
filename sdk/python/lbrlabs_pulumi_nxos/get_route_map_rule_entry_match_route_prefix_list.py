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
    'GetRouteMapRuleEntryMatchRoutePrefixListResult',
    'AwaitableGetRouteMapRuleEntryMatchRoutePrefixListResult',
    'get_route_map_rule_entry_match_route_prefix_list',
    'get_route_map_rule_entry_match_route_prefix_list_output',
]

@pulumi.output_type
class GetRouteMapRuleEntryMatchRoutePrefixListResult:
    """
    A collection of values returned by getRouteMapRuleEntryMatchRoutePrefixList.
    """
    def __init__(__self__, device=None, id=None, order=None, prefix_list_dn=None, rule_name=None):
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if order and not isinstance(order, int):
            raise TypeError("Expected argument 'order' to be a int")
        pulumi.set(__self__, "order", order)
        if prefix_list_dn and not isinstance(prefix_list_dn, str):
            raise TypeError("Expected argument 'prefix_list_dn' to be a str")
        pulumi.set(__self__, "prefix_list_dn", prefix_list_dn)
        if rule_name and not isinstance(rule_name, str):
            raise TypeError("Expected argument 'rule_name' to be a str")
        pulumi.set(__self__, "rule_name", rule_name)

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
    @pulumi.getter
    def order(self) -> int:
        """
        Route-Map Rule Entry order.
        """
        return pulumi.get(self, "order")

    @property
    @pulumi.getter(name="prefixListDn")
    def prefix_list_dn(self) -> str:
        """
        DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
        """
        return pulumi.get(self, "prefix_list_dn")

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> str:
        """
        Route Map rule name.
        """
        return pulumi.get(self, "rule_name")


class AwaitableGetRouteMapRuleEntryMatchRoutePrefixListResult(GetRouteMapRuleEntryMatchRoutePrefixListResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRouteMapRuleEntryMatchRoutePrefixListResult(
            device=self.device,
            id=self.id,
            order=self.order,
            prefix_list_dn=self.prefix_list_dn,
            rule_name=self.rule_name)


def get_route_map_rule_entry_match_route_prefix_list(device: Optional[str] = None,
                                                     order: Optional[int] = None,
                                                     prefix_list_dn: Optional[str] = None,
                                                     rule_name: Optional[str] = None,
                                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRouteMapRuleEntryMatchRoutePrefixListResult:
    """
    This data source can read a Match Route Prefix List in Route-Map Rule Entry configuration.

    - API Documentation: [rtmapRsRtDstAtt](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:RsRtDstAtt/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_route_map_rule_entry_match_route_prefix_list(order=10,
        prefix_list_dn="sys/rpm/pfxlistv4-[LIST1]",
        rule_name="RULE1")
    ```


    :param str device: A device name from the provider configuration.
    :param int order: Route-Map Rule Entry order.
    :param str prefix_list_dn: DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
    :param str rule_name: Route Map rule name.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['order'] = order
    __args__['prefixListDn'] = prefix_list_dn
    __args__['ruleName'] = rule_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getRouteMapRuleEntryMatchRoutePrefixList:getRouteMapRuleEntryMatchRoutePrefixList', __args__, opts=opts, typ=GetRouteMapRuleEntryMatchRoutePrefixListResult).value

    return AwaitableGetRouteMapRuleEntryMatchRoutePrefixListResult(
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        order=pulumi.get(__ret__, 'order'),
        prefix_list_dn=pulumi.get(__ret__, 'prefix_list_dn'),
        rule_name=pulumi.get(__ret__, 'rule_name'))


@_utilities.lift_output_func(get_route_map_rule_entry_match_route_prefix_list)
def get_route_map_rule_entry_match_route_prefix_list_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                                                            order: Optional[pulumi.Input[int]] = None,
                                                            prefix_list_dn: Optional[pulumi.Input[str]] = None,
                                                            rule_name: Optional[pulumi.Input[str]] = None,
                                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRouteMapRuleEntryMatchRoutePrefixListResult]:
    """
    This data source can read a Match Route Prefix List in Route-Map Rule Entry configuration.

    - API Documentation: [rtmapRsRtDstAtt](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/rtmap:RsRtDstAtt/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_route_map_rule_entry_match_route_prefix_list(order=10,
        prefix_list_dn="sys/rpm/pfxlistv4-[LIST1]",
        rule_name="RULE1")
    ```


    :param str device: A device name from the provider configuration.
    :param int order: Route-Map Rule Entry order.
    :param str prefix_list_dn: DN of Prefix List. For example: `sys/rpm/pfxlistv4-[LIST1]`
    :param str rule_name: Route Map rule name.
    """
    ...