# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VrfRoutingArgs', 'VrfRouting']

@pulumi.input_type
class VrfRoutingArgs:
    def __init__(__self__, *,
                 vrf: pulumi.Input[str],
                 device: Optional[pulumi.Input[str]] = None,
                 route_distinguisher: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VrfRouting resource.
        :param pulumi.Input[str] vrf: VRF name.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] route_distinguisher: Route Distinguisher value in NX-OS DME format. - Default value: `unknown:unknown:0:0`
        """
        pulumi.set(__self__, "vrf", vrf)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if route_distinguisher is not None:
            pulumi.set(__self__, "route_distinguisher", route_distinguisher)

    @property
    @pulumi.getter
    def vrf(self) -> pulumi.Input[str]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf")

    @vrf.setter
    def vrf(self, value: pulumi.Input[str]):
        pulumi.set(self, "vrf", value)

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
    @pulumi.getter(name="routeDistinguisher")
    def route_distinguisher(self) -> Optional[pulumi.Input[str]]:
        """
        Route Distinguisher value in NX-OS DME format. - Default value: `unknown:unknown:0:0`
        """
        return pulumi.get(self, "route_distinguisher")

    @route_distinguisher.setter
    def route_distinguisher(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_distinguisher", value)


@pulumi.input_type
class _VrfRoutingState:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 route_distinguisher: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VrfRouting resources.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] route_distinguisher: Route Distinguisher value in NX-OS DME format. - Default value: `unknown:unknown:0:0`
        :param pulumi.Input[str] vrf: VRF name.
        """
        if device is not None:
            pulumi.set(__self__, "device", device)
        if route_distinguisher is not None:
            pulumi.set(__self__, "route_distinguisher", route_distinguisher)
        if vrf is not None:
            pulumi.set(__self__, "vrf", vrf)

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
    @pulumi.getter(name="routeDistinguisher")
    def route_distinguisher(self) -> Optional[pulumi.Input[str]]:
        """
        Route Distinguisher value in NX-OS DME format. - Default value: `unknown:unknown:0:0`
        """
        return pulumi.get(self, "route_distinguisher")

    @route_distinguisher.setter
    def route_distinguisher(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_distinguisher", value)

    @property
    @pulumi.getter
    def vrf(self) -> Optional[pulumi.Input[str]]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf")

    @vrf.setter
    def vrf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vrf", value)


class VrfRouting(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 route_distinguisher: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a VrfRouting resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] route_distinguisher: Route Distinguisher value in NX-OS DME format. - Default value: `unknown:unknown:0:0`
        :param pulumi.Input[str] vrf: VRF name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VrfRoutingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VrfRouting resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VrfRoutingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VrfRoutingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 route_distinguisher: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VrfRoutingArgs.__new__(VrfRoutingArgs)

            __props__.__dict__["device"] = device
            __props__.__dict__["route_distinguisher"] = route_distinguisher
            if vrf is None and not opts.urn:
                raise TypeError("Missing required property 'vrf'")
            __props__.__dict__["vrf"] = vrf
        super(VrfRouting, __self__).__init__(
            'nxos:nxos/vrfRouting:VrfRouting',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device: Optional[pulumi.Input[str]] = None,
            route_distinguisher: Optional[pulumi.Input[str]] = None,
            vrf: Optional[pulumi.Input[str]] = None) -> 'VrfRouting':
        """
        Get an existing VrfRouting resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] route_distinguisher: Route Distinguisher value in NX-OS DME format. - Default value: `unknown:unknown:0:0`
        :param pulumi.Input[str] vrf: VRF name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VrfRoutingState.__new__(_VrfRoutingState)

        __props__.__dict__["device"] = device
        __props__.__dict__["route_distinguisher"] = route_distinguisher
        __props__.__dict__["vrf"] = vrf
        return VrfRouting(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="routeDistinguisher")
    def route_distinguisher(self) -> pulumi.Output[str]:
        """
        Route Distinguisher value in NX-OS DME format. - Default value: `unknown:unknown:0:0`
        """
        return pulumi.get(self, "route_distinguisher")

    @property
    @pulumi.getter
    def vrf(self) -> pulumi.Output[str]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf")

