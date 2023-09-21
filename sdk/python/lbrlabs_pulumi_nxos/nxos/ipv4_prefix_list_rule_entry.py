# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['Ipv4PrefixListRuleEntryArgs', 'Ipv4PrefixListRuleEntry']

@pulumi.input_type
class Ipv4PrefixListRuleEntryArgs:
    def __init__(__self__, *,
                 order: pulumi.Input[int],
                 rule_name: pulumi.Input[str],
                 action: Optional[pulumi.Input[str]] = None,
                 criteria: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 from_range: Optional[pulumi.Input[int]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 to_range: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Ipv4PrefixListRuleEntry resource.
        :param pulumi.Input[int] order: IPv4 Prefix List Rule Entry order. - Range: `0`-`4294967294`
        :param pulumi.Input[str] rule_name: IPv4 Prefix List Rule name.
        :param pulumi.Input[str] action: IPv4 Prefix List Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
        :param pulumi.Input[str] criteria: IPv4 Prefix List Rule Entry criteria. - Choices: `exact`, `inexact` - Default value: `exact`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] from_range: IPv4 Prefix List Rule Entry start range. - Range: `0`-`128` - Default value: `0`
        :param pulumi.Input[str] prefix: IPv4 Prefix List Rule Entry prefix.
        :param pulumi.Input[int] to_range: IPv4 Prefix List Rule Entry end range. - Range: `0`-`128` - Default value: `0`
        """
        pulumi.set(__self__, "order", order)
        pulumi.set(__self__, "rule_name", rule_name)
        if action is not None:
            pulumi.set(__self__, "action", action)
        if criteria is not None:
            pulumi.set(__self__, "criteria", criteria)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if from_range is not None:
            pulumi.set(__self__, "from_range", from_range)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if to_range is not None:
            pulumi.set(__self__, "to_range", to_range)

    @property
    @pulumi.getter
    def order(self) -> pulumi.Input[int]:
        """
        IPv4 Prefix List Rule Entry order. - Range: `0`-`4294967294`
        """
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: pulumi.Input[int]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> pulumi.Input[str]:
        """
        IPv4 Prefix List Rule name.
        """
        return pulumi.get(self, "rule_name")

    @rule_name.setter
    def rule_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "rule_name", value)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 Prefix List Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def criteria(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 Prefix List Rule Entry criteria. - Choices: `exact`, `inexact` - Default value: `exact`
        """
        return pulumi.get(self, "criteria")

    @criteria.setter
    def criteria(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "criteria", value)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter(name="fromRange")
    def from_range(self) -> Optional[pulumi.Input[int]]:
        """
        IPv4 Prefix List Rule Entry start range. - Range: `0`-`128` - Default value: `0`
        """
        return pulumi.get(self, "from_range")

    @from_range.setter
    def from_range(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "from_range", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 Prefix List Rule Entry prefix.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter(name="toRange")
    def to_range(self) -> Optional[pulumi.Input[int]]:
        """
        IPv4 Prefix List Rule Entry end range. - Range: `0`-`128` - Default value: `0`
        """
        return pulumi.get(self, "to_range")

    @to_range.setter
    def to_range(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "to_range", value)


@pulumi.input_type
class _Ipv4PrefixListRuleEntryState:
    def __init__(__self__, *,
                 action: Optional[pulumi.Input[str]] = None,
                 criteria: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 from_range: Optional[pulumi.Input[int]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 to_range: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Ipv4PrefixListRuleEntry resources.
        :param pulumi.Input[str] action: IPv4 Prefix List Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
        :param pulumi.Input[str] criteria: IPv4 Prefix List Rule Entry criteria. - Choices: `exact`, `inexact` - Default value: `exact`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] from_range: IPv4 Prefix List Rule Entry start range. - Range: `0`-`128` - Default value: `0`
        :param pulumi.Input[int] order: IPv4 Prefix List Rule Entry order. - Range: `0`-`4294967294`
        :param pulumi.Input[str] prefix: IPv4 Prefix List Rule Entry prefix.
        :param pulumi.Input[str] rule_name: IPv4 Prefix List Rule name.
        :param pulumi.Input[int] to_range: IPv4 Prefix List Rule Entry end range. - Range: `0`-`128` - Default value: `0`
        """
        if action is not None:
            pulumi.set(__self__, "action", action)
        if criteria is not None:
            pulumi.set(__self__, "criteria", criteria)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if from_range is not None:
            pulumi.set(__self__, "from_range", from_range)
        if order is not None:
            pulumi.set(__self__, "order", order)
        if prefix is not None:
            pulumi.set(__self__, "prefix", prefix)
        if rule_name is not None:
            pulumi.set(__self__, "rule_name", rule_name)
        if to_range is not None:
            pulumi.set(__self__, "to_range", to_range)

    @property
    @pulumi.getter
    def action(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 Prefix List Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
        """
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def criteria(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 Prefix List Rule Entry criteria. - Choices: `exact`, `inexact` - Default value: `exact`
        """
        return pulumi.get(self, "criteria")

    @criteria.setter
    def criteria(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "criteria", value)

    @property
    @pulumi.getter
    def device(self) -> Optional[pulumi.Input[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @device.setter
    def device(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device", value)

    @property
    @pulumi.getter(name="fromRange")
    def from_range(self) -> Optional[pulumi.Input[int]]:
        """
        IPv4 Prefix List Rule Entry start range. - Range: `0`-`128` - Default value: `0`
        """
        return pulumi.get(self, "from_range")

    @from_range.setter
    def from_range(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "from_range", value)

    @property
    @pulumi.getter
    def order(self) -> Optional[pulumi.Input[int]]:
        """
        IPv4 Prefix List Rule Entry order. - Range: `0`-`4294967294`
        """
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter
    def prefix(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 Prefix List Rule Entry prefix.
        """
        return pulumi.get(self, "prefix")

    @prefix.setter
    def prefix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix", value)

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> Optional[pulumi.Input[str]]:
        """
        IPv4 Prefix List Rule name.
        """
        return pulumi.get(self, "rule_name")

    @rule_name.setter
    def rule_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rule_name", value)

    @property
    @pulumi.getter(name="toRange")
    def to_range(self) -> Optional[pulumi.Input[int]]:
        """
        IPv4 Prefix List Rule Entry end range. - Range: `0`-`128` - Default value: `0`
        """
        return pulumi.get(self, "to_range")

    @to_range.setter
    def to_range(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "to_range", value)


class Ipv4PrefixListRuleEntry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 criteria: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 from_range: Optional[pulumi.Input[int]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 to_range: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a Ipv4PrefixListRuleEntry resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: IPv4 Prefix List Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
        :param pulumi.Input[str] criteria: IPv4 Prefix List Rule Entry criteria. - Choices: `exact`, `inexact` - Default value: `exact`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] from_range: IPv4 Prefix List Rule Entry start range. - Range: `0`-`128` - Default value: `0`
        :param pulumi.Input[int] order: IPv4 Prefix List Rule Entry order. - Range: `0`-`4294967294`
        :param pulumi.Input[str] prefix: IPv4 Prefix List Rule Entry prefix.
        :param pulumi.Input[str] rule_name: IPv4 Prefix List Rule name.
        :param pulumi.Input[int] to_range: IPv4 Prefix List Rule Entry end range. - Range: `0`-`128` - Default value: `0`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Ipv4PrefixListRuleEntryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Ipv4PrefixListRuleEntry resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param Ipv4PrefixListRuleEntryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Ipv4PrefixListRuleEntryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 criteria: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 from_range: Optional[pulumi.Input[int]] = None,
                 order: Optional[pulumi.Input[int]] = None,
                 prefix: Optional[pulumi.Input[str]] = None,
                 rule_name: Optional[pulumi.Input[str]] = None,
                 to_range: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Ipv4PrefixListRuleEntryArgs.__new__(Ipv4PrefixListRuleEntryArgs)

            __props__.__dict__["action"] = action
            __props__.__dict__["criteria"] = criteria
            __props__.__dict__["device"] = device
            __props__.__dict__["from_range"] = from_range
            if order is None and not opts.urn:
                raise TypeError("Missing required property 'order'")
            __props__.__dict__["order"] = order
            __props__.__dict__["prefix"] = prefix
            if rule_name is None and not opts.urn:
                raise TypeError("Missing required property 'rule_name'")
            __props__.__dict__["rule_name"] = rule_name
            __props__.__dict__["to_range"] = to_range
        super(Ipv4PrefixListRuleEntry, __self__).__init__(
            'nxos:nxos/ipv4PrefixListRuleEntry:Ipv4PrefixListRuleEntry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            action: Optional[pulumi.Input[str]] = None,
            criteria: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None,
            from_range: Optional[pulumi.Input[int]] = None,
            order: Optional[pulumi.Input[int]] = None,
            prefix: Optional[pulumi.Input[str]] = None,
            rule_name: Optional[pulumi.Input[str]] = None,
            to_range: Optional[pulumi.Input[int]] = None) -> 'Ipv4PrefixListRuleEntry':
        """
        Get an existing Ipv4PrefixListRuleEntry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] action: IPv4 Prefix List Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
        :param pulumi.Input[str] criteria: IPv4 Prefix List Rule Entry criteria. - Choices: `exact`, `inexact` - Default value: `exact`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] from_range: IPv4 Prefix List Rule Entry start range. - Range: `0`-`128` - Default value: `0`
        :param pulumi.Input[int] order: IPv4 Prefix List Rule Entry order. - Range: `0`-`4294967294`
        :param pulumi.Input[str] prefix: IPv4 Prefix List Rule Entry prefix.
        :param pulumi.Input[str] rule_name: IPv4 Prefix List Rule name.
        :param pulumi.Input[int] to_range: IPv4 Prefix List Rule Entry end range. - Range: `0`-`128` - Default value: `0`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Ipv4PrefixListRuleEntryState.__new__(_Ipv4PrefixListRuleEntryState)

        __props__.__dict__["action"] = action
        __props__.__dict__["criteria"] = criteria
        __props__.__dict__["device"] = device
        __props__.__dict__["from_range"] = from_range
        __props__.__dict__["order"] = order
        __props__.__dict__["prefix"] = prefix
        __props__.__dict__["rule_name"] = rule_name
        __props__.__dict__["to_range"] = to_range
        return Ipv4PrefixListRuleEntry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Output[str]:
        """
        IPv4 Prefix List Rule Entry action. - Choices: `deny`, `permit` - Default value: `permit`
        """
        return pulumi.get(self, "action")

    @property
    @pulumi.getter
    def criteria(self) -> pulumi.Output[str]:
        """
        IPv4 Prefix List Rule Entry criteria. - Choices: `exact`, `inexact` - Default value: `exact`
        """
        return pulumi.get(self, "criteria")

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="fromRange")
    def from_range(self) -> pulumi.Output[int]:
        """
        IPv4 Prefix List Rule Entry start range. - Range: `0`-`128` - Default value: `0`
        """
        return pulumi.get(self, "from_range")

    @property
    @pulumi.getter
    def order(self) -> pulumi.Output[int]:
        """
        IPv4 Prefix List Rule Entry order. - Range: `0`-`4294967294`
        """
        return pulumi.get(self, "order")

    @property
    @pulumi.getter
    def prefix(self) -> pulumi.Output[Optional[str]]:
        """
        IPv4 Prefix List Rule Entry prefix.
        """
        return pulumi.get(self, "prefix")

    @property
    @pulumi.getter(name="ruleName")
    def rule_name(self) -> pulumi.Output[str]:
        """
        IPv4 Prefix List Rule name.
        """
        return pulumi.get(self, "rule_name")

    @property
    @pulumi.getter(name="toRange")
    def to_range(self) -> pulumi.Output[int]:
        """
        IPv4 Prefix List Rule Entry end range. - Range: `0`-`128` - Default value: `0`
        """
        return pulumi.get(self, "to_range")
