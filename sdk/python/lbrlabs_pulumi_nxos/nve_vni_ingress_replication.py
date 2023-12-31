# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NveVniIngressReplicationArgs', 'NveVniIngressReplication']

@pulumi.input_type
class NveVniIngressReplicationArgs:
    def __init__(__self__, *,
                 vni: pulumi.Input[int],
                 device: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NveVniIngressReplication resource.
        :param pulumi.Input[int] vni: Virtual Network ID. - Range: `1`-`16777214`
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] protocol: Configure VxLAN Ingress Replication mode. - Choices: `bgp`, `unknown`, `static` - Default value: `unknown`
        """
        pulumi.set(__self__, "vni", vni)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)

    @property
    @pulumi.getter
    def vni(self) -> pulumi.Input[int]:
        """
        Virtual Network ID. - Range: `1`-`16777214`
        """
        return pulumi.get(self, "vni")

    @vni.setter
    def vni(self, value: pulumi.Input[int]):
        pulumi.set(self, "vni", value)

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
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Configure VxLAN Ingress Replication mode. - Choices: `bgp`, `unknown`, `static` - Default value: `unknown`
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)


@pulumi.input_type
class _NveVniIngressReplicationState:
    def __init__(__self__, *,
                 device: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 vni: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering NveVniIngressReplication resources.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] protocol: Configure VxLAN Ingress Replication mode. - Choices: `bgp`, `unknown`, `static` - Default value: `unknown`
        :param pulumi.Input[int] vni: Virtual Network ID. - Range: `1`-`16777214`
        """
        if device is not None:
            pulumi.set(__self__, "device", device)
        if protocol is not None:
            pulumi.set(__self__, "protocol", protocol)
        if vni is not None:
            pulumi.set(__self__, "vni", vni)

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
    def protocol(self) -> Optional[pulumi.Input[str]]:
        """
        Configure VxLAN Ingress Replication mode. - Choices: `bgp`, `unknown`, `static` - Default value: `unknown`
        """
        return pulumi.get(self, "protocol")

    @protocol.setter
    def protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "protocol", value)

    @property
    @pulumi.getter
    def vni(self) -> Optional[pulumi.Input[int]]:
        """
        Virtual Network ID. - Range: `1`-`16777214`
        """
        return pulumi.get(self, "vni")

    @vni.setter
    def vni(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "vni", value)


class NveVniIngressReplication(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 vni: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        This resource can manage the configuration of Ingress Replication for Virtual Network ID (VNI).

        - API Documentation: [nvoIngRepl](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Network%20Virtualization/nvo:IngRepl/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.NveVniIngressReplication("example",
            protocol="bgp",
            vni=103100)
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/nveVniIngressReplication:NveVniIngressReplication example "sys/eps/epId-[1]/nws/vni-[103100]/IngRepl"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] protocol: Configure VxLAN Ingress Replication mode. - Choices: `bgp`, `unknown`, `static` - Default value: `unknown`
        :param pulumi.Input[int] vni: Virtual Network ID. - Range: `1`-`16777214`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NveVniIngressReplicationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage the configuration of Ingress Replication for Virtual Network ID (VNI).

        - API Documentation: [nvoIngRepl](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Network%20Virtualization/nvo:IngRepl/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.NveVniIngressReplication("example",
            protocol="bgp",
            vni=103100)
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/nveVniIngressReplication:NveVniIngressReplication example "sys/eps/epId-[1]/nws/vni-[103100]/IngRepl"
        ```

        :param str resource_name: The name of the resource.
        :param NveVniIngressReplicationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NveVniIngressReplicationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 protocol: Optional[pulumi.Input[str]] = None,
                 vni: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NveVniIngressReplicationArgs.__new__(NveVniIngressReplicationArgs)

            __props__.__dict__["device"] = device
            __props__.__dict__["protocol"] = protocol
            if vni is None and not opts.urn:
                raise TypeError("Missing required property 'vni'")
            __props__.__dict__["vni"] = vni
        super(NveVniIngressReplication, __self__).__init__(
            'nxos:index/nveVniIngressReplication:NveVniIngressReplication',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device: Optional[pulumi.Input[str]] = None,
            protocol: Optional[pulumi.Input[str]] = None,
            vni: Optional[pulumi.Input[int]] = None) -> 'NveVniIngressReplication':
        """
        Get an existing NveVniIngressReplication resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] protocol: Configure VxLAN Ingress Replication mode. - Choices: `bgp`, `unknown`, `static` - Default value: `unknown`
        :param pulumi.Input[int] vni: Virtual Network ID. - Range: `1`-`16777214`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NveVniIngressReplicationState.__new__(_NveVniIngressReplicationState)

        __props__.__dict__["device"] = device
        __props__.__dict__["protocol"] = protocol
        __props__.__dict__["vni"] = vni
        return NveVniIngressReplication(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter
    def protocol(self) -> pulumi.Output[str]:
        """
        Configure VxLAN Ingress Replication mode. - Choices: `bgp`, `unknown`, `static` - Default value: `unknown`
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter
    def vni(self) -> pulumi.Output[int]:
        """
        Virtual Network ID. - Range: `1`-`16777214`
        """
        return pulumi.get(self, "vni")

