# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['PimAnycastRpArgs', 'PimAnycastRp']

@pulumi.input_type
class PimAnycastRpArgs:
    def __init__(__self__, *,
                 vrf_name: pulumi.Input[str],
                 device: Optional[pulumi.Input[str]] = None,
                 local_interface: Optional[pulumi.Input[str]] = None,
                 source_interface: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PimAnycastRp resource.
        :param pulumi.Input[str] vrf_name: VRF name.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] local_interface: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[str] source_interface: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        pulumi.set(__self__, "vrf_name", vrf_name)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if local_interface is not None:
            pulumi.set(__self__, "local_interface", local_interface)
        if source_interface is not None:
            pulumi.set(__self__, "source_interface", source_interface)

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
    @pulumi.getter(name="localInterface")
    def local_interface(self) -> Optional[pulumi.Input[str]]:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "local_interface")

    @local_interface.setter
    def local_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_interface", value)

    @property
    @pulumi.getter(name="sourceInterface")
    def source_interface(self) -> Optional[pulumi.Input[str]]:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "source_interface")

    @source_interface.setter
    def source_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_interface", value)


@pulumi.input_type
class _PimAnycastRpState:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 local_interface: Optional[pulumi.Input[str]] = None,
                 source_interface: Optional[pulumi.Input[str]] = None,
                 vrf_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PimAnycastRp resources.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] local_interface: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[str] source_interface: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[str] vrf_name: VRF name.
        """
        if device is not None:
            pulumi.set(__self__, "device", device)
        if local_interface is not None:
            pulumi.set(__self__, "local_interface", local_interface)
        if source_interface is not None:
            pulumi.set(__self__, "source_interface", source_interface)
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
    @pulumi.getter(name="localInterface")
    def local_interface(self) -> Optional[pulumi.Input[str]]:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "local_interface")

    @local_interface.setter
    def local_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_interface", value)

    @property
    @pulumi.getter(name="sourceInterface")
    def source_interface(self) -> Optional[pulumi.Input[str]]:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "source_interface")

    @source_interface.setter
    def source_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_interface", value)

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


class PimAnycastRp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 local_interface: Optional[pulumi.Input[str]] = None,
                 source_interface: Optional[pulumi.Input[str]] = None,
                 vrf_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource can manage the PIM Anycast RP configuration.

        - API Documentation: [pimAcastRPFuncP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:AcastRPFuncP/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.PimAnycastRp("example",
            local_interface="eth1/10",
            source_interface="eth1/10",
            vrf_name="default")
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/pimAnycastRp:PimAnycastRp example "sys/pim/inst/dom-[default]/acastrpfunc"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] local_interface: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[str] source_interface: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[str] vrf_name: VRF name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PimAnycastRpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage the PIM Anycast RP configuration.

        - API Documentation: [pimAcastRPFuncP](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:AcastRPFuncP/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.PimAnycastRp("example",
            local_interface="eth1/10",
            source_interface="eth1/10",
            vrf_name="default")
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/pimAnycastRp:PimAnycastRp example "sys/pim/inst/dom-[default]/acastrpfunc"
        ```

        :param str resource_name: The name of the resource.
        :param PimAnycastRpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PimAnycastRpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 local_interface: Optional[pulumi.Input[str]] = None,
                 source_interface: Optional[pulumi.Input[str]] = None,
                 vrf_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PimAnycastRpArgs.__new__(PimAnycastRpArgs)

            __props__.__dict__["device"] = device
            __props__.__dict__["local_interface"] = local_interface
            __props__.__dict__["source_interface"] = source_interface
            if vrf_name is None and not opts.urn:
                raise TypeError("Missing required property 'vrf_name'")
            __props__.__dict__["vrf_name"] = vrf_name
        super(PimAnycastRp, __self__).__init__(
            'nxos:index/pimAnycastRp:PimAnycastRp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device: Optional[pulumi.Input[str]] = None,
            local_interface: Optional[pulumi.Input[str]] = None,
            source_interface: Optional[pulumi.Input[str]] = None,
            vrf_name: Optional[pulumi.Input[str]] = None) -> 'PimAnycastRp':
        """
        Get an existing PimAnycastRp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] local_interface: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[str] source_interface: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[str] vrf_name: VRF name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PimAnycastRpState.__new__(_PimAnycastRpState)

        __props__.__dict__["device"] = device
        __props__.__dict__["local_interface"] = local_interface
        __props__.__dict__["source_interface"] = source_interface
        __props__.__dict__["vrf_name"] = vrf_name
        return PimAnycastRp(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="localInterface")
    def local_interface(self) -> pulumi.Output[Optional[str]]:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "local_interface")

    @property
    @pulumi.getter(name="sourceInterface")
    def source_interface(self) -> pulumi.Output[Optional[str]]:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "source_interface")

    @property
    @pulumi.getter(name="vrfName")
    def vrf_name(self) -> pulumi.Output[str]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf_name")

