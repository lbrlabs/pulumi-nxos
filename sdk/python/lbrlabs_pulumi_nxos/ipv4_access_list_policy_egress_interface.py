# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['Ipv4AccessListPolicyEgressInterfaceArgs', 'Ipv4AccessListPolicyEgressInterface']

@pulumi.input_type
class Ipv4AccessListPolicyEgressInterfaceArgs:
    def __init__(__self__, *,
                 interface_id: pulumi.Input[str],
                 access_list_name: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Ipv4AccessListPolicyEgressInterface resource.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        :param pulumi.Input[str] access_list_name: Access list name.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        """
        pulumi.set(__self__, "interface_id", interface_id)
        if access_list_name is not None:
            pulumi.set(__self__, "access_list_name", access_list_name)
        if device is not None:
            pulumi.set(__self__, "device", device)

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> pulumi.Input[str]:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "interface_id")

    @interface_id.setter
    def interface_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface_id", value)

    @property
    @pulumi.getter(name="accessListName")
    def access_list_name(self) -> Optional[pulumi.Input[str]]:
        """
        Access list name.
        """
        return pulumi.get(self, "access_list_name")

    @access_list_name.setter
    def access_list_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_list_name", value)

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
class _Ipv4AccessListPolicyEgressInterfaceState:
    def __init__(__self__, *,
                 access_list_name: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Ipv4AccessListPolicyEgressInterface resources.
        :param pulumi.Input[str] access_list_name: Access list name.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        if access_list_name is not None:
            pulumi.set(__self__, "access_list_name", access_list_name)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if interface_id is not None:
            pulumi.set(__self__, "interface_id", interface_id)

    @property
    @pulumi.getter(name="accessListName")
    def access_list_name(self) -> Optional[pulumi.Input[str]]:
        """
        Access list name.
        """
        return pulumi.get(self, "access_list_name")

    @access_list_name.setter
    def access_list_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_list_name", value)

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
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "interface_id")

    @interface_id.setter
    def interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_id", value)


class Ipv4AccessListPolicyEgressInterface(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_list_name: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource can manage an IPv4 Access List Policy Egress Interface.

        - API Documentation: [aclIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Security%20and%20Policing/acl:If/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.Ipv4AccessListPolicyEgressInterface("example",
            access_list_name="ACL1",
            interface_id="eth1/10")
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/ipv4AccessListPolicyEgressInterface:Ipv4AccessListPolicyEgressInterface example "sys/acl/ipv4/policy/egress/intf-[eth1/10]"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_list_name: Access list name.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Ipv4AccessListPolicyEgressInterfaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage an IPv4 Access List Policy Egress Interface.

        - API Documentation: [aclIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Security%20and%20Policing/acl:If/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.Ipv4AccessListPolicyEgressInterface("example",
            access_list_name="ACL1",
            interface_id="eth1/10")
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/ipv4AccessListPolicyEgressInterface:Ipv4AccessListPolicyEgressInterface example "sys/acl/ipv4/policy/egress/intf-[eth1/10]"
        ```

        :param str resource_name: The name of the resource.
        :param Ipv4AccessListPolicyEgressInterfaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(Ipv4AccessListPolicyEgressInterfaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_list_name: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = Ipv4AccessListPolicyEgressInterfaceArgs.__new__(Ipv4AccessListPolicyEgressInterfaceArgs)

            __props__.__dict__["access_list_name"] = access_list_name
            __props__.__dict__["device"] = device
            if interface_id is None and not opts.urn:
                raise TypeError("Missing required property 'interface_id'")
            __props__.__dict__["interface_id"] = interface_id
        super(Ipv4AccessListPolicyEgressInterface, __self__).__init__(
            'nxos:index/ipv4AccessListPolicyEgressInterface:Ipv4AccessListPolicyEgressInterface',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_list_name: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None,
            interface_id: Optional[pulumi.Input[str]] = None) -> 'Ipv4AccessListPolicyEgressInterface':
        """
        Get an existing Ipv4AccessListPolicyEgressInterface resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_list_name: Access list name.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _Ipv4AccessListPolicyEgressInterfaceState.__new__(_Ipv4AccessListPolicyEgressInterfaceState)

        __props__.__dict__["access_list_name"] = access_list_name
        __props__.__dict__["device"] = device
        __props__.__dict__["interface_id"] = interface_id
        return Ipv4AccessListPolicyEgressInterface(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessListName")
    def access_list_name(self) -> pulumi.Output[Optional[str]]:
        """
        Access list name.
        """
        return pulumi.get(self, "access_list_name")

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> pulumi.Output[str]:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "interface_id")

