# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BgpInstanceArgs', 'BgpInstance']

@pulumi.input_type
class BgpInstanceArgs:
    def __init__(__self__, *,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 enhanced_error_handling: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a BgpInstance resource.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[bool] enhanced_error_handling: Enable BGP Enhanced Error Handling. - Default value: `true`
        """
        if admin_state is not None:
            pulumi.set(__self__, "admin_state", admin_state)
        if asn is not None:
            pulumi.set(__self__, "asn", asn)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if enhanced_error_handling is not None:
            pulumi.set(__self__, "enhanced_error_handling", enhanced_error_handling)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> Optional[pulumi.Input[str]]:
        """
        Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        """
        return pulumi.get(self, "admin_state")

    @admin_state.setter
    def admin_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_state", value)

    @property
    @pulumi.getter
    def asn(self) -> Optional[pulumi.Input[str]]:
        """
        Autonomous system number.
        """
        return pulumi.get(self, "asn")

    @asn.setter
    def asn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "asn", value)

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
    @pulumi.getter(name="enhancedErrorHandling")
    def enhanced_error_handling(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable BGP Enhanced Error Handling. - Default value: `true`
        """
        return pulumi.get(self, "enhanced_error_handling")

    @enhanced_error_handling.setter
    def enhanced_error_handling(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enhanced_error_handling", value)


@pulumi.input_type
class _BgpInstanceState:
    def __init__(__self__, *,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 enhanced_error_handling: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering BgpInstance resources.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[bool] enhanced_error_handling: Enable BGP Enhanced Error Handling. - Default value: `true`
        """
        if admin_state is not None:
            pulumi.set(__self__, "admin_state", admin_state)
        if asn is not None:
            pulumi.set(__self__, "asn", asn)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if enhanced_error_handling is not None:
            pulumi.set(__self__, "enhanced_error_handling", enhanced_error_handling)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> Optional[pulumi.Input[str]]:
        """
        Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        """
        return pulumi.get(self, "admin_state")

    @admin_state.setter
    def admin_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_state", value)

    @property
    @pulumi.getter
    def asn(self) -> Optional[pulumi.Input[str]]:
        """
        Autonomous system number.
        """
        return pulumi.get(self, "asn")

    @asn.setter
    def asn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "asn", value)

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
    @pulumi.getter(name="enhancedErrorHandling")
    def enhanced_error_handling(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable BGP Enhanced Error Handling. - Default value: `true`
        """
        return pulumi.get(self, "enhanced_error_handling")

    @enhanced_error_handling.setter
    def enhanced_error_handling(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enhanced_error_handling", value)


class BgpInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 enhanced_error_handling: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        This resource can manage the BGP instance configuration.

        - API Documentation: [bgpInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Inst/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.BgpInstance("example",
            admin_state="enabled",
            asn="65001",
            enhanced_error_handling=False)
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/bgpInstance:BgpInstance example "sys/bgp/inst"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[bool] enhanced_error_handling: Enable BGP Enhanced Error Handling. - Default value: `true`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[BgpInstanceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage the BGP instance configuration.

        - API Documentation: [bgpInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:Inst/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.BgpInstance("example",
            admin_state="enabled",
            asn="65001",
            enhanced_error_handling=False)
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/bgpInstance:BgpInstance example "sys/bgp/inst"
        ```

        :param str resource_name: The name of the resource.
        :param BgpInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BgpInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 enhanced_error_handling: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BgpInstanceArgs.__new__(BgpInstanceArgs)

            __props__.__dict__["admin_state"] = admin_state
            __props__.__dict__["asn"] = asn
            __props__.__dict__["device"] = device
            __props__.__dict__["enhanced_error_handling"] = enhanced_error_handling
        super(BgpInstance, __self__).__init__(
            'nxos:index/bgpInstance:BgpInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_state: Optional[pulumi.Input[str]] = None,
            asn: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None,
            enhanced_error_handling: Optional[pulumi.Input[bool]] = None) -> 'BgpInstance':
        """
        Get an existing BgpInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[bool] enhanced_error_handling: Enable BGP Enhanced Error Handling. - Default value: `true`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BgpInstanceState.__new__(_BgpInstanceState)

        __props__.__dict__["admin_state"] = admin_state
        __props__.__dict__["asn"] = asn
        __props__.__dict__["device"] = device
        __props__.__dict__["enhanced_error_handling"] = enhanced_error_handling
        return BgpInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> pulumi.Output[str]:
        """
        Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        """
        return pulumi.get(self, "admin_state")

    @property
    @pulumi.getter
    def asn(self) -> pulumi.Output[Optional[str]]:
        """
        Autonomous system number.
        """
        return pulumi.get(self, "asn")

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="enhancedErrorHandling")
    def enhanced_error_handling(self) -> pulumi.Output[bool]:
        """
        Enable BGP Enhanced Error Handling. - Default value: `true`
        """
        return pulumi.get(self, "enhanced_error_handling")

