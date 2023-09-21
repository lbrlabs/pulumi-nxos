# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PimStaticRpPolicyArgs', 'PimStaticRpPolicy']

@pulumi.input_type
class PimStaticRpPolicyArgs:
    def __init__(__self__, *,
                 vrf_name: pulumi.Input[str],
                 device: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PimStaticRpPolicy resource.
        :param pulumi.Input[str] vrf_name: VRF name.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] name: Policy name.
        """
        pulumi.set(__self__, "vrf_name", vrf_name)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="vrfName")
    def vrf_name(self) -> pulumi.Input[str]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf_name")

    @vrf_name.setter
    def vrf_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "vrf_name", value)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _PimStaticRpPolicyState:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vrf_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PimStaticRpPolicy resources.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] name: Policy name.
        :param pulumi.Input[str] vrf_name: VRF name.
        """
        if device is not None:
            pulumi.set(__self__, "device", device)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if vrf_name is not None:
            pulumi.set(__self__, "vrf_name", vrf_name)

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
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="vrfName")
    def vrf_name(self) -> Optional[pulumi.Input[str]]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf_name")

    @vrf_name.setter
    def vrf_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vrf_name", value)


class PimStaticRpPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vrf_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a PimStaticRpPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] name: Policy name.
        :param pulumi.Input[str] vrf_name: VRF name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PimStaticRpPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a PimStaticRpPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param PimStaticRpPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PimStaticRpPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 vrf_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PimStaticRpPolicyArgs.__new__(PimStaticRpPolicyArgs)

            __props__.__dict__["device"] = device
            __props__.__dict__["name"] = name
            if vrf_name is None and not opts.urn:
                raise TypeError("Missing required property 'vrf_name'")
            __props__.__dict__["vrf_name"] = vrf_name
        super(PimStaticRpPolicy, __self__).__init__(
            'nxos:nxos/pimStaticRpPolicy:PimStaticRpPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            vrf_name: Optional[pulumi.Input[str]] = None) -> 'PimStaticRpPolicy':
        """
        Get an existing PimStaticRpPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] name: Policy name.
        :param pulumi.Input[str] vrf_name: VRF name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PimStaticRpPolicyState.__new__(_PimStaticRpPolicyState)

        __props__.__dict__["device"] = device
        __props__.__dict__["name"] = name
        __props__.__dict__["vrf_name"] = vrf_name
        return PimStaticRpPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Policy name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="vrfName")
    def vrf_name(self) -> pulumi.Output[str]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf_name")

