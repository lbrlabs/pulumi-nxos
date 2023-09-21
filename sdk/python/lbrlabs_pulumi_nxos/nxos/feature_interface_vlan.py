# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FeatureInterfaceVlanArgs', 'FeatureInterfaceVlan']

@pulumi.input_type
class FeatureInterfaceVlanArgs:
    def __init__(__self__, *,
                 admin_state: pulumi.Input[str],
                 device: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a FeatureInterfaceVlan resource.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        """
        pulumi.set(__self__, "admin_state", admin_state)
        if device is not None:
            pulumi.set(__self__, "device", device)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> pulumi.Input[str]:
        """
        Administrative state. - Choices: `enabled`, `disabled`
        """
        return pulumi.get(self, "admin_state")

    @admin_state.setter
    def admin_state(self, value: pulumi.Input[str]):
        pulumi.set(self, "admin_state", value)

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
class _FeatureInterfaceVlanState:
    def __init__(__self__, *,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering FeatureInterfaceVlan resources.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        """
        if admin_state is not None:
            pulumi.set(__self__, "admin_state", admin_state)
        if device is not None:
            pulumi.set(__self__, "device", device)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> Optional[pulumi.Input[str]]:
        """
        Administrative state. - Choices: `enabled`, `disabled`
        """
        return pulumi.get(self, "admin_state")

    @admin_state.setter
    def admin_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_state", value)

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


class FeatureInterfaceVlan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a FeatureInterfaceVlan resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FeatureInterfaceVlanArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a FeatureInterfaceVlan resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param FeatureInterfaceVlanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FeatureInterfaceVlanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FeatureInterfaceVlanArgs.__new__(FeatureInterfaceVlanArgs)

            if admin_state is None and not opts.urn:
                raise TypeError("Missing required property 'admin_state'")
            __props__.__dict__["admin_state"] = admin_state
            __props__.__dict__["device"] = device
        super(FeatureInterfaceVlan, __self__).__init__(
            'nxos:nxos/featureInterfaceVlan:FeatureInterfaceVlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_state: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None) -> 'FeatureInterfaceVlan':
        """
        Get an existing FeatureInterfaceVlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FeatureInterfaceVlanState.__new__(_FeatureInterfaceVlanState)

        __props__.__dict__["admin_state"] = admin_state
        __props__.__dict__["device"] = device
        return FeatureInterfaceVlan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> pulumi.Output[str]:
        """
        Administrative state. - Choices: `enabled`, `disabled`
        """
        return pulumi.get(self, "admin_state")

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

