# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['QueuingQosPolicyMapMatchClassMapRemainingBandwidthArgs', 'QueuingQosPolicyMapMatchClassMapRemainingBandwidth']

@pulumi.input_type
class QueuingQosPolicyMapMatchClassMapRemainingBandwidthArgs:
    def __init__(__self__, *,
                 class_map_name: pulumi.Input[str],
                 policy_map_name: pulumi.Input[str],
                 value: pulumi.Input[int],
                 device: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a QueuingQosPolicyMapMatchClassMapRemainingBandwidth resource.
        :param pulumi.Input[str] class_map_name: Class map name.
        :param pulumi.Input[str] policy_map_name: Policy map name.
        :param pulumi.Input[int] value: Remaining bandwidth percent. - Range: `0`-`100`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        """
        pulumi.set(__self__, "class_map_name", class_map_name)
        pulumi.set(__self__, "policy_map_name", policy_map_name)
        pulumi.set(__self__, "value", value)
        if device is not None:
            pulumi.set(__self__, "device", device)

    @property
    @pulumi.getter(name="classMapName")
    def class_map_name(self) -> pulumi.Input[str]:
        """
        Class map name.
        """
        return pulumi.get(self, "class_map_name")

    @class_map_name.setter
    def class_map_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "class_map_name", value)

    @property
    @pulumi.getter(name="policyMapName")
    def policy_map_name(self) -> pulumi.Input[str]:
        """
        Policy map name.
        """
        return pulumi.get(self, "policy_map_name")

    @policy_map_name.setter
    def policy_map_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_map_name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[int]:
        """
        Remaining bandwidth percent. - Range: `0`-`100`
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[int]):
        pulumi.set(self, "value", value)

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


@pulumi.input_type
class _QueuingQosPolicyMapMatchClassMapRemainingBandwidthState:
    def __init__(__self__, *,
                 class_map_name: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 policy_map_name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering QueuingQosPolicyMapMatchClassMapRemainingBandwidth resources.
        :param pulumi.Input[str] class_map_name: Class map name.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] policy_map_name: Policy map name.
        :param pulumi.Input[int] value: Remaining bandwidth percent. - Range: `0`-`100`
        """
        if class_map_name is not None:
            pulumi.set(__self__, "class_map_name", class_map_name)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if policy_map_name is not None:
            pulumi.set(__self__, "policy_map_name", policy_map_name)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="classMapName")
    def class_map_name(self) -> Optional[pulumi.Input[str]]:
        """
        Class map name.
        """
        return pulumi.get(self, "class_map_name")

    @class_map_name.setter
    def class_map_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "class_map_name", value)

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
    @pulumi.getter(name="policyMapName")
    def policy_map_name(self) -> Optional[pulumi.Input[str]]:
        """
        Policy map name.
        """
        return pulumi.get(self, "policy_map_name")

    @policy_map_name.setter
    def policy_map_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_map_name", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[int]]:
        """
        Remaining bandwidth percent. - Range: `0`-`100`
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "value", value)


class QueuingQosPolicyMapMatchClassMapRemainingBandwidth(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 class_map_name: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 policy_map_name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Create a QueuingQosPolicyMapMatchClassMapRemainingBandwidth resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] class_map_name: Class map name.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] policy_map_name: Policy map name.
        :param pulumi.Input[int] value: Remaining bandwidth percent. - Range: `0`-`100`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: QueuingQosPolicyMapMatchClassMapRemainingBandwidthArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a QueuingQosPolicyMapMatchClassMapRemainingBandwidth resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param QueuingQosPolicyMapMatchClassMapRemainingBandwidthArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QueuingQosPolicyMapMatchClassMapRemainingBandwidthArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 class_map_name: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 policy_map_name: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QueuingQosPolicyMapMatchClassMapRemainingBandwidthArgs.__new__(QueuingQosPolicyMapMatchClassMapRemainingBandwidthArgs)

            if class_map_name is None and not opts.urn:
                raise TypeError("Missing required property 'class_map_name'")
            __props__.__dict__["class_map_name"] = class_map_name
            __props__.__dict__["device"] = device
            if policy_map_name is None and not opts.urn:
                raise TypeError("Missing required property 'policy_map_name'")
            __props__.__dict__["policy_map_name"] = policy_map_name
            if value is None and not opts.urn:
                raise TypeError("Missing required property 'value'")
            __props__.__dict__["value"] = value
        super(QueuingQosPolicyMapMatchClassMapRemainingBandwidth, __self__).__init__(
            'nxos:nxos/queuingQosPolicyMapMatchClassMapRemainingBandwidth:QueuingQosPolicyMapMatchClassMapRemainingBandwidth',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            class_map_name: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None,
            policy_map_name: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[int]] = None) -> 'QueuingQosPolicyMapMatchClassMapRemainingBandwidth':
        """
        Get an existing QueuingQosPolicyMapMatchClassMapRemainingBandwidth resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] class_map_name: Class map name.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] policy_map_name: Policy map name.
        :param pulumi.Input[int] value: Remaining bandwidth percent. - Range: `0`-`100`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QueuingQosPolicyMapMatchClassMapRemainingBandwidthState.__new__(_QueuingQosPolicyMapMatchClassMapRemainingBandwidthState)

        __props__.__dict__["class_map_name"] = class_map_name
        __props__.__dict__["device"] = device
        __props__.__dict__["policy_map_name"] = policy_map_name
        __props__.__dict__["value"] = value
        return QueuingQosPolicyMapMatchClassMapRemainingBandwidth(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="classMapName")
    def class_map_name(self) -> pulumi.Output[str]:
        """
        Class map name.
        """
        return pulumi.get(self, "class_map_name")

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="policyMapName")
    def policy_map_name(self) -> pulumi.Output[str]:
        """
        Policy map name.
        """
        return pulumi.get(self, "policy_map_name")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[int]:
        """
        Remaining bandwidth percent. - Range: `0`-`100`
        """
        return pulumi.get(self, "value")

