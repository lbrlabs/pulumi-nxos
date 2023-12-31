# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['OspfInstanceArgs', 'OspfInstance']

@pulumi.input_type
class OspfInstanceArgs:
    def __init__(__self__, *,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OspfInstance resource.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] name: OSPF instance name.
        """
        if admin_state is not None:
            pulumi.set(__self__, "admin_state", admin_state)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if name is not None:
            pulumi.set(__self__, "name", name)

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
        OSPF instance name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _OspfInstanceState:
    def __init__(__self__, *,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OspfInstance resources.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] name: OSPF instance name.
        """
        if admin_state is not None:
            pulumi.set(__self__, "admin_state", admin_state)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if name is not None:
            pulumi.set(__self__, "name", name)

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
        OSPF instance name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class OspfInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource can manage the OSPF instance configuration.

        - API Documentation: [ospfInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/ospf:Inst/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.OspfInstance("example", admin_state="enabled")
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/ospfInstance:OspfInstance example "sys/ospf/inst-[OSPF1]"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] name: OSPF instance name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OspfInstanceArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage the OSPF instance configuration.

        - API Documentation: [ospfInst](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/ospf:Inst/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.OspfInstance("example", admin_state="enabled")
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/ospfInstance:OspfInstance example "sys/ospf/inst-[OSPF1]"
        ```

        :param str resource_name: The name of the resource.
        :param OspfInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OspfInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OspfInstanceArgs.__new__(OspfInstanceArgs)

            __props__.__dict__["admin_state"] = admin_state
            __props__.__dict__["device"] = device
            __props__.__dict__["name"] = name
        super(OspfInstance, __self__).__init__(
            'nxos:index/ospfInstance:OspfInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_state: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'OspfInstance':
        """
        Get an existing OspfInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_state: Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] name: OSPF instance name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OspfInstanceState.__new__(_OspfInstanceState)

        __props__.__dict__["admin_state"] = admin_state
        __props__.__dict__["device"] = device
        __props__.__dict__["name"] = name
        return OspfInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> pulumi.Output[str]:
        """
        Administrative state. - Choices: `enabled`, `disabled` - Default value: `enabled`
        """
        return pulumi.get(self, "admin_state")

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
        OSPF instance name.
        """
        return pulumi.get(self, "name")

