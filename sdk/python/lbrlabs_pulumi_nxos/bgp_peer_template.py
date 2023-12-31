# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['BgpPeerTemplateArgs', 'BgpPeerTemplate']

@pulumi.input_type
class BgpPeerTemplateArgs:
    def __init__(__self__, *,
                 asn: pulumi.Input[str],
                 template_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 peer_type: Optional[pulumi.Input[str]] = None,
                 remote_asn: Optional[pulumi.Input[str]] = None,
                 source_interface: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BgpPeerTemplate resource.
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] template_name: Peer template name.
        :param pulumi.Input[str] description: Peer template description.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] peer_type: Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
               `fabric-internal`
        :param pulumi.Input[str] remote_asn: Peer template autonomous system number.
        :param pulumi.Input[str] source_interface: Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
        """
        pulumi.set(__self__, "asn", asn)
        pulumi.set(__self__, "template_name", template_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if peer_type is not None:
            pulumi.set(__self__, "peer_type", peer_type)
        if remote_asn is not None:
            pulumi.set(__self__, "remote_asn", remote_asn)
        if source_interface is not None:
            pulumi.set(__self__, "source_interface", source_interface)

    @property
    @pulumi.getter
    def asn(self) -> pulumi.Input[str]:
        """
        Autonomous system number.
        """
        return pulumi.get(self, "asn")

    @asn.setter
    def asn(self, value: pulumi.Input[str]):
        pulumi.set(self, "asn", value)

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> pulumi.Input[str]:
        """
        Peer template name.
        """
        return pulumi.get(self, "template_name")

    @template_name.setter
    def template_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "template_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Peer template description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
    @pulumi.getter(name="peerType")
    def peer_type(self) -> Optional[pulumi.Input[str]]:
        """
        Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
        `fabric-internal`
        """
        return pulumi.get(self, "peer_type")

    @peer_type.setter
    def peer_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_type", value)

    @property
    @pulumi.getter(name="remoteAsn")
    def remote_asn(self) -> Optional[pulumi.Input[str]]:
        """
        Peer template autonomous system number.
        """
        return pulumi.get(self, "remote_asn")

    @remote_asn.setter
    def remote_asn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_asn", value)

    @property
    @pulumi.getter(name="sourceInterface")
    def source_interface(self) -> Optional[pulumi.Input[str]]:
        """
        Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
        """
        return pulumi.get(self, "source_interface")

    @source_interface.setter
    def source_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_interface", value)


@pulumi.input_type
class _BgpPeerTemplateState:
    def __init__(__self__, *,
                 asn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 peer_type: Optional[pulumi.Input[str]] = None,
                 remote_asn: Optional[pulumi.Input[str]] = None,
                 source_interface: Optional[pulumi.Input[str]] = None,
                 template_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BgpPeerTemplate resources.
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] description: Peer template description.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] peer_type: Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
               `fabric-internal`
        :param pulumi.Input[str] remote_asn: Peer template autonomous system number.
        :param pulumi.Input[str] source_interface: Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
        :param pulumi.Input[str] template_name: Peer template name.
        """
        if asn is not None:
            pulumi.set(__self__, "asn", asn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if peer_type is not None:
            pulumi.set(__self__, "peer_type", peer_type)
        if remote_asn is not None:
            pulumi.set(__self__, "remote_asn", remote_asn)
        if source_interface is not None:
            pulumi.set(__self__, "source_interface", source_interface)
        if template_name is not None:
            pulumi.set(__self__, "template_name", template_name)

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
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Peer template description.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

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
    @pulumi.getter(name="peerType")
    def peer_type(self) -> Optional[pulumi.Input[str]]:
        """
        Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
        `fabric-internal`
        """
        return pulumi.get(self, "peer_type")

    @peer_type.setter
    def peer_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_type", value)

    @property
    @pulumi.getter(name="remoteAsn")
    def remote_asn(self) -> Optional[pulumi.Input[str]]:
        """
        Peer template autonomous system number.
        """
        return pulumi.get(self, "remote_asn")

    @remote_asn.setter
    def remote_asn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "remote_asn", value)

    @property
    @pulumi.getter(name="sourceInterface")
    def source_interface(self) -> Optional[pulumi.Input[str]]:
        """
        Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
        """
        return pulumi.get(self, "source_interface")

    @source_interface.setter
    def source_interface(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "source_interface", value)

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> Optional[pulumi.Input[str]]:
        """
        Peer template name.
        """
        return pulumi.get(self, "template_name")

    @template_name.setter
    def template_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template_name", value)


class BgpPeerTemplate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 peer_type: Optional[pulumi.Input[str]] = None,
                 remote_asn: Optional[pulumi.Input[str]] = None,
                 source_interface: Optional[pulumi.Input[str]] = None,
                 template_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource can manage the BGP peer template configuration.

        - API Documentation: [bgpPeerCont](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:PeerCont/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.BgpPeerTemplate("example",
            asn="65001",
            description="My Description",
            peer_type="fabric-internal",
            remote_asn="65002",
            source_interface="lo0",
            template_name="SPINE-PEERS")
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/bgpPeerTemplate:BgpPeerTemplate example "sys/bgp/inst/dom-[default]/peercont-[SPINE-PEERS]"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] description: Peer template description.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] peer_type: Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
               `fabric-internal`
        :param pulumi.Input[str] remote_asn: Peer template autonomous system number.
        :param pulumi.Input[str] source_interface: Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
        :param pulumi.Input[str] template_name: Peer template name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BgpPeerTemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage the BGP peer template configuration.

        - API Documentation: [bgpPeerCont](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Routing%20and%20Forwarding/bgp:PeerCont/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.BgpPeerTemplate("example",
            asn="65001",
            description="My Description",
            peer_type="fabric-internal",
            remote_asn="65002",
            source_interface="lo0",
            template_name="SPINE-PEERS")
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/bgpPeerTemplate:BgpPeerTemplate example "sys/bgp/inst/dom-[default]/peercont-[SPINE-PEERS]"
        ```

        :param str resource_name: The name of the resource.
        :param BgpPeerTemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BgpPeerTemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 peer_type: Optional[pulumi.Input[str]] = None,
                 remote_asn: Optional[pulumi.Input[str]] = None,
                 source_interface: Optional[pulumi.Input[str]] = None,
                 template_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BgpPeerTemplateArgs.__new__(BgpPeerTemplateArgs)

            if asn is None and not opts.urn:
                raise TypeError("Missing required property 'asn'")
            __props__.__dict__["asn"] = asn
            __props__.__dict__["description"] = description
            __props__.__dict__["device"] = device
            __props__.__dict__["peer_type"] = peer_type
            __props__.__dict__["remote_asn"] = remote_asn
            __props__.__dict__["source_interface"] = source_interface
            if template_name is None and not opts.urn:
                raise TypeError("Missing required property 'template_name'")
            __props__.__dict__["template_name"] = template_name
        super(BgpPeerTemplate, __self__).__init__(
            'nxos:index/bgpPeerTemplate:BgpPeerTemplate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            asn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None,
            peer_type: Optional[pulumi.Input[str]] = None,
            remote_asn: Optional[pulumi.Input[str]] = None,
            source_interface: Optional[pulumi.Input[str]] = None,
            template_name: Optional[pulumi.Input[str]] = None) -> 'BgpPeerTemplate':
        """
        Get an existing BgpPeerTemplate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] description: Peer template description.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] peer_type: Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
               `fabric-internal`
        :param pulumi.Input[str] remote_asn: Peer template autonomous system number.
        :param pulumi.Input[str] source_interface: Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
        :param pulumi.Input[str] template_name: Peer template name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BgpPeerTemplateState.__new__(_BgpPeerTemplateState)

        __props__.__dict__["asn"] = asn
        __props__.__dict__["description"] = description
        __props__.__dict__["device"] = device
        __props__.__dict__["peer_type"] = peer_type
        __props__.__dict__["remote_asn"] = remote_asn
        __props__.__dict__["source_interface"] = source_interface
        __props__.__dict__["template_name"] = template_name
        return BgpPeerTemplate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def asn(self) -> pulumi.Output[str]:
        """
        Autonomous system number.
        """
        return pulumi.get(self, "asn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Peer template description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def device(self) -> pulumi.Output[Optional[str]]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="peerType")
    def peer_type(self) -> pulumi.Output[str]:
        """
        Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
        `fabric-internal`
        """
        return pulumi.get(self, "peer_type")

    @property
    @pulumi.getter(name="remoteAsn")
    def remote_asn(self) -> pulumi.Output[Optional[str]]:
        """
        Peer template autonomous system number.
        """
        return pulumi.get(self, "remote_asn")

    @property
    @pulumi.getter(name="sourceInterface")
    def source_interface(self) -> pulumi.Output[str]:
        """
        Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
        """
        return pulumi.get(self, "source_interface")

    @property
    @pulumi.getter(name="templateName")
    def template_name(self) -> pulumi.Output[str]:
        """
        Peer template name.
        """
        return pulumi.get(self, "template_name")

