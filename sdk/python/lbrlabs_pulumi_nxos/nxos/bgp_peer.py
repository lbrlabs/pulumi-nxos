# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BgpPeerArgs', 'BgpPeer']

@pulumi.input_type
class BgpPeerArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 asn: pulumi.Input[str],
                 vrf: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 hold_time: Optional[pulumi.Input[int]] = None,
                 keepalive: Optional[pulumi.Input[int]] = None,
                 peer_template: Optional[pulumi.Input[str]] = None,
                 peer_type: Optional[pulumi.Input[str]] = None,
                 remote_asn: Optional[pulumi.Input[str]] = None,
                 source_interface: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BgpPeer resource.
        :param pulumi.Input[str] address: Peer address.
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] vrf: VRF name.
        :param pulumi.Input[str] description: Peer description.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] hold_time: BGP Hold Timer in seconds. The value must be greater than the keepalive timer - Range: `3`-`3600` - Default value: `180`
        :param pulumi.Input[int] keepalive: BGP Keepalive Timer in seconds - Range: `0`-`3600` - Default value: `60`
        :param pulumi.Input[str] peer_template: Peer template name.
        :param pulumi.Input[str] peer_type: Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
               `fabric-internal`
        :param pulumi.Input[str] remote_asn: Peer autonomous system number.
        :param pulumi.Input[str] source_interface: Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
        """
        pulumi.set(__self__, "address", address)
        pulumi.set(__self__, "asn", asn)
        pulumi.set(__self__, "vrf", vrf)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if hold_time is not None:
            pulumi.set(__self__, "hold_time", hold_time)
        if keepalive is not None:
            pulumi.set(__self__, "keepalive", keepalive)
        if peer_template is not None:
            pulumi.set(__self__, "peer_template", peer_template)
        if peer_type is not None:
            pulumi.set(__self__, "peer_type", peer_type)
        if remote_asn is not None:
            pulumi.set(__self__, "remote_asn", remote_asn)
        if source_interface is not None:
            pulumi.set(__self__, "source_interface", source_interface)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        Peer address.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

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
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Peer description.
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
    @pulumi.getter(name="holdTime")
    def hold_time(self) -> Optional[pulumi.Input[int]]:
        """
        BGP Hold Timer in seconds. The value must be greater than the keepalive timer - Range: `3`-`3600` - Default value: `180`
        """
        return pulumi.get(self, "hold_time")

    @hold_time.setter
    def hold_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hold_time", value)

    @property
    @pulumi.getter
    def keepalive(self) -> Optional[pulumi.Input[int]]:
        """
        BGP Keepalive Timer in seconds - Range: `0`-`3600` - Default value: `60`
        """
        return pulumi.get(self, "keepalive")

    @keepalive.setter
    def keepalive(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "keepalive", value)

    @property
    @pulumi.getter(name="peerTemplate")
    def peer_template(self) -> Optional[pulumi.Input[str]]:
        """
        Peer template name.
        """
        return pulumi.get(self, "peer_template")

    @peer_template.setter
    def peer_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_template", value)

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
        Peer autonomous system number.
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
class _BgpPeerState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 hold_time: Optional[pulumi.Input[int]] = None,
                 keepalive: Optional[pulumi.Input[int]] = None,
                 peer_template: Optional[pulumi.Input[str]] = None,
                 peer_type: Optional[pulumi.Input[str]] = None,
                 remote_asn: Optional[pulumi.Input[str]] = None,
                 source_interface: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BgpPeer resources.
        :param pulumi.Input[str] address: Peer address.
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] description: Peer description.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] hold_time: BGP Hold Timer in seconds. The value must be greater than the keepalive timer - Range: `3`-`3600` - Default value: `180`
        :param pulumi.Input[int] keepalive: BGP Keepalive Timer in seconds - Range: `0`-`3600` - Default value: `60`
        :param pulumi.Input[str] peer_template: Peer template name.
        :param pulumi.Input[str] peer_type: Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
               `fabric-internal`
        :param pulumi.Input[str] remote_asn: Peer autonomous system number.
        :param pulumi.Input[str] source_interface: Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
        :param pulumi.Input[str] vrf: VRF name.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if asn is not None:
            pulumi.set(__self__, "asn", asn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if hold_time is not None:
            pulumi.set(__self__, "hold_time", hold_time)
        if keepalive is not None:
            pulumi.set(__self__, "keepalive", keepalive)
        if peer_template is not None:
            pulumi.set(__self__, "peer_template", peer_template)
        if peer_type is not None:
            pulumi.set(__self__, "peer_type", peer_type)
        if remote_asn is not None:
            pulumi.set(__self__, "remote_asn", remote_asn)
        if source_interface is not None:
            pulumi.set(__self__, "source_interface", source_interface)
        if vrf is not None:
            pulumi.set(__self__, "vrf", vrf)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Peer address.
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

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
        Peer description.
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
    @pulumi.getter(name="holdTime")
    def hold_time(self) -> Optional[pulumi.Input[int]]:
        """
        BGP Hold Timer in seconds. The value must be greater than the keepalive timer - Range: `3`-`3600` - Default value: `180`
        """
        return pulumi.get(self, "hold_time")

    @hold_time.setter
    def hold_time(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "hold_time", value)

    @property
    @pulumi.getter
    def keepalive(self) -> Optional[pulumi.Input[int]]:
        """
        BGP Keepalive Timer in seconds - Range: `0`-`3600` - Default value: `60`
        """
        return pulumi.get(self, "keepalive")

    @keepalive.setter
    def keepalive(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "keepalive", value)

    @property
    @pulumi.getter(name="peerTemplate")
    def peer_template(self) -> Optional[pulumi.Input[str]]:
        """
        Peer template name.
        """
        return pulumi.get(self, "peer_template")

    @peer_template.setter
    def peer_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_template", value)

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
        Peer autonomous system number.
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
    @pulumi.getter
    def vrf(self) -> Optional[pulumi.Input[str]]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf")

    @vrf.setter
    def vrf(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vrf", value)


class BgpPeer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 hold_time: Optional[pulumi.Input[int]] = None,
                 keepalive: Optional[pulumi.Input[int]] = None,
                 peer_template: Optional[pulumi.Input[str]] = None,
                 peer_type: Optional[pulumi.Input[str]] = None,
                 remote_asn: Optional[pulumi.Input[str]] = None,
                 source_interface: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a BgpPeer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Peer address.
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] description: Peer description.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] hold_time: BGP Hold Timer in seconds. The value must be greater than the keepalive timer - Range: `3`-`3600` - Default value: `180`
        :param pulumi.Input[int] keepalive: BGP Keepalive Timer in seconds - Range: `0`-`3600` - Default value: `60`
        :param pulumi.Input[str] peer_template: Peer template name.
        :param pulumi.Input[str] peer_type: Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
               `fabric-internal`
        :param pulumi.Input[str] remote_asn: Peer autonomous system number.
        :param pulumi.Input[str] source_interface: Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
        :param pulumi.Input[str] vrf: VRF name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BgpPeerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a BgpPeer resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param BgpPeerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BgpPeerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 asn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 hold_time: Optional[pulumi.Input[int]] = None,
                 keepalive: Optional[pulumi.Input[int]] = None,
                 peer_template: Optional[pulumi.Input[str]] = None,
                 peer_type: Optional[pulumi.Input[str]] = None,
                 remote_asn: Optional[pulumi.Input[str]] = None,
                 source_interface: Optional[pulumi.Input[str]] = None,
                 vrf: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BgpPeerArgs.__new__(BgpPeerArgs)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            if asn is None and not opts.urn:
                raise TypeError("Missing required property 'asn'")
            __props__.__dict__["asn"] = asn
            __props__.__dict__["description"] = description
            __props__.__dict__["device"] = device
            __props__.__dict__["hold_time"] = hold_time
            __props__.__dict__["keepalive"] = keepalive
            __props__.__dict__["peer_template"] = peer_template
            __props__.__dict__["peer_type"] = peer_type
            __props__.__dict__["remote_asn"] = remote_asn
            __props__.__dict__["source_interface"] = source_interface
            if vrf is None and not opts.urn:
                raise TypeError("Missing required property 'vrf'")
            __props__.__dict__["vrf"] = vrf
        super(BgpPeer, __self__).__init__(
            'nxos:nxos/bgpPeer:BgpPeer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            asn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None,
            hold_time: Optional[pulumi.Input[int]] = None,
            keepalive: Optional[pulumi.Input[int]] = None,
            peer_template: Optional[pulumi.Input[str]] = None,
            peer_type: Optional[pulumi.Input[str]] = None,
            remote_asn: Optional[pulumi.Input[str]] = None,
            source_interface: Optional[pulumi.Input[str]] = None,
            vrf: Optional[pulumi.Input[str]] = None) -> 'BgpPeer':
        """
        Get an existing BgpPeer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Peer address.
        :param pulumi.Input[str] asn: Autonomous system number.
        :param pulumi.Input[str] description: Peer description.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[int] hold_time: BGP Hold Timer in seconds. The value must be greater than the keepalive timer - Range: `3`-`3600` - Default value: `180`
        :param pulumi.Input[int] keepalive: BGP Keepalive Timer in seconds - Range: `0`-`3600` - Default value: `60`
        :param pulumi.Input[str] peer_template: Peer template name.
        :param pulumi.Input[str] peer_type: Neighbor Fabric Type. - Choices: `fabric-internal`, `fabric-external`, `fabric-border-leaf` - Default value:
               `fabric-internal`
        :param pulumi.Input[str] remote_asn: Peer autonomous system number.
        :param pulumi.Input[str] source_interface: Source Interface. Must match first field in the output of `show intf brief`. - Default value: `unspecified`
        :param pulumi.Input[str] vrf: VRF name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BgpPeerState.__new__(_BgpPeerState)

        __props__.__dict__["address"] = address
        __props__.__dict__["asn"] = asn
        __props__.__dict__["description"] = description
        __props__.__dict__["device"] = device
        __props__.__dict__["hold_time"] = hold_time
        __props__.__dict__["keepalive"] = keepalive
        __props__.__dict__["peer_template"] = peer_template
        __props__.__dict__["peer_type"] = peer_type
        __props__.__dict__["remote_asn"] = remote_asn
        __props__.__dict__["source_interface"] = source_interface
        __props__.__dict__["vrf"] = vrf
        return BgpPeer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        Peer address.
        """
        return pulumi.get(self, "address")

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
        Peer description.
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
    @pulumi.getter(name="holdTime")
    def hold_time(self) -> pulumi.Output[int]:
        """
        BGP Hold Timer in seconds. The value must be greater than the keepalive timer - Range: `3`-`3600` - Default value: `180`
        """
        return pulumi.get(self, "hold_time")

    @property
    @pulumi.getter
    def keepalive(self) -> pulumi.Output[int]:
        """
        BGP Keepalive Timer in seconds - Range: `0`-`3600` - Default value: `60`
        """
        return pulumi.get(self, "keepalive")

    @property
    @pulumi.getter(name="peerTemplate")
    def peer_template(self) -> pulumi.Output[Optional[str]]:
        """
        Peer template name.
        """
        return pulumi.get(self, "peer_template")

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
        Peer autonomous system number.
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
    @pulumi.getter
    def vrf(self) -> pulumi.Output[str]:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf")

