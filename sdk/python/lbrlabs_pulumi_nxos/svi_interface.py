# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SviInterfaceArgs', 'SviInterface']

@pulumi.input_type
class SviInterfaceArgs:
    def __init__(__self__, *,
                 interface_id: pulumi.Input[str],
                 admin_state: Optional[pulumi.Input[str]] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 delay: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 medium: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a SviInterface resource.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `vlan100`.
        :param pulumi.Input[str] admin_state: Administrative port state. - Choices: `up`, `down` - Default value: `up`
        :param pulumi.Input[int] bandwidth: Specifies the administrative port bandwidth. - Range: `1`-`400000000` - Default value: `1000000`
        :param pulumi.Input[int] delay: Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
        :param pulumi.Input[str] description: Interface description.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] medium: The administrative port medium type. - Choices: `bcast`, `p2p` - Default value: `bcast`
        :param pulumi.Input[int] mtu: Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
        """
        pulumi.set(__self__, "interface_id", interface_id)
        if admin_state is not None:
            pulumi.set(__self__, "admin_state", admin_state)
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if delay is not None:
            pulumi.set(__self__, "delay", delay)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if medium is not None:
            pulumi.set(__self__, "medium", medium)
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> pulumi.Input[str]:
        """
        Must match first field in the output of `show intf brief`. Example: `vlan100`.
        """
        return pulumi.get(self, "interface_id")

    @interface_id.setter
    def interface_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "interface_id", value)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> Optional[pulumi.Input[str]]:
        """
        Administrative port state. - Choices: `up`, `down` - Default value: `up`
        """
        return pulumi.get(self, "admin_state")

    @admin_state.setter
    def admin_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_state", value)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the administrative port bandwidth. - Range: `1`-`400000000` - Default value: `1000000`
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter
    def delay(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
        """
        return pulumi.get(self, "delay")

    @delay.setter
    def delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "delay", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Interface description.
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
    @pulumi.getter
    def medium(self) -> Optional[pulumi.Input[str]]:
        """
        The administrative port medium type. - Choices: `bcast`, `p2p` - Default value: `bcast`
        """
        return pulumi.get(self, "medium")

    @medium.setter
    def medium(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "medium", value)

    @property
    @pulumi.getter
    def mtu(self) -> Optional[pulumi.Input[int]]:
        """
        Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
        """
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mtu", value)


@pulumi.input_type
class _SviInterfaceState:
    def __init__(__self__, *,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 delay: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 medium: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering SviInterface resources.
        :param pulumi.Input[str] admin_state: Administrative port state. - Choices: `up`, `down` - Default value: `up`
        :param pulumi.Input[int] bandwidth: Specifies the administrative port bandwidth. - Range: `1`-`400000000` - Default value: `1000000`
        :param pulumi.Input[int] delay: Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
        :param pulumi.Input[str] description: Interface description.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `vlan100`.
        :param pulumi.Input[str] medium: The administrative port medium type. - Choices: `bcast`, `p2p` - Default value: `bcast`
        :param pulumi.Input[int] mtu: Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
        """
        if admin_state is not None:
            pulumi.set(__self__, "admin_state", admin_state)
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if delay is not None:
            pulumi.set(__self__, "delay", delay)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if device is not None:
            pulumi.set(__self__, "device", device)
        if interface_id is not None:
            pulumi.set(__self__, "interface_id", interface_id)
        if medium is not None:
            pulumi.set(__self__, "medium", medium)
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> Optional[pulumi.Input[str]]:
        """
        Administrative port state. - Choices: `up`, `down` - Default value: `up`
        """
        return pulumi.get(self, "admin_state")

    @admin_state.setter
    def admin_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "admin_state", value)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the administrative port bandwidth. - Range: `1`-`400000000` - Default value: `1000000`
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter
    def delay(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
        """
        return pulumi.get(self, "delay")

    @delay.setter
    def delay(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "delay", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Interface description.
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
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        Must match first field in the output of `show intf brief`. Example: `vlan100`.
        """
        return pulumi.get(self, "interface_id")

    @interface_id.setter
    def interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_id", value)

    @property
    @pulumi.getter
    def medium(self) -> Optional[pulumi.Input[str]]:
        """
        The administrative port medium type. - Choices: `bcast`, `p2p` - Default value: `bcast`
        """
        return pulumi.get(self, "medium")

    @medium.setter
    def medium(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "medium", value)

    @property
    @pulumi.getter
    def mtu(self) -> Optional[pulumi.Input[int]]:
        """
        Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
        """
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mtu", value)


class SviInterface(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 delay: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 medium: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        This resource can manage an SVI interface.

        - API Documentation: [sviIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Interfaces/svi:If/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.SviInterface("example",
            admin_state="down",
            bandwidth=1000,
            delay=10,
            description="My Description",
            interface_id="vlan293",
            medium="bcast",
            mtu=9216)
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/sviInterface:SviInterface example "sys/intf/svi-[vlan293]"
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_state: Administrative port state. - Choices: `up`, `down` - Default value: `up`
        :param pulumi.Input[int] bandwidth: Specifies the administrative port bandwidth. - Range: `1`-`400000000` - Default value: `1000000`
        :param pulumi.Input[int] delay: Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
        :param pulumi.Input[str] description: Interface description.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `vlan100`.
        :param pulumi.Input[str] medium: The administrative port medium type. - Choices: `bcast`, `p2p` - Default value: `bcast`
        :param pulumi.Input[int] mtu: Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SviInterfaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource can manage an SVI interface.

        - API Documentation: [sviIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Interfaces/svi:If/)

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_nxos as nxos

        example = nxos.SviInterface("example",
            admin_state="down",
            bandwidth=1000,
            delay=10,
            description="My Description",
            interface_id="vlan293",
            medium="bcast",
            mtu=9216)
        ```

        ## Import

        ```sh
         $ pulumi import nxos:index/sviInterface:SviInterface example "sys/intf/svi-[vlan293]"
        ```

        :param str resource_name: The name of the resource.
        :param SviInterfaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SviInterfaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_state: Optional[pulumi.Input[str]] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 delay: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 device: Optional[pulumi.Input[str]] = None,
                 interface_id: Optional[pulumi.Input[str]] = None,
                 medium: Optional[pulumi.Input[str]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SviInterfaceArgs.__new__(SviInterfaceArgs)

            __props__.__dict__["admin_state"] = admin_state
            __props__.__dict__["bandwidth"] = bandwidth
            __props__.__dict__["delay"] = delay
            __props__.__dict__["description"] = description
            __props__.__dict__["device"] = device
            if interface_id is None and not opts.urn:
                raise TypeError("Missing required property 'interface_id'")
            __props__.__dict__["interface_id"] = interface_id
            __props__.__dict__["medium"] = medium
            __props__.__dict__["mtu"] = mtu
        super(SviInterface, __self__).__init__(
            'nxos:index/sviInterface:SviInterface',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_state: Optional[pulumi.Input[str]] = None,
            bandwidth: Optional[pulumi.Input[int]] = None,
            delay: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            device: Optional[pulumi.Input[str]] = None,
            interface_id: Optional[pulumi.Input[str]] = None,
            medium: Optional[pulumi.Input[str]] = None,
            mtu: Optional[pulumi.Input[int]] = None) -> 'SviInterface':
        """
        Get an existing SviInterface resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] admin_state: Administrative port state. - Choices: `up`, `down` - Default value: `up`
        :param pulumi.Input[int] bandwidth: Specifies the administrative port bandwidth. - Range: `1`-`400000000` - Default value: `1000000`
        :param pulumi.Input[int] delay: Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
        :param pulumi.Input[str] description: Interface description.
        :param pulumi.Input[str] device: A device name from the provider configuration.
        :param pulumi.Input[str] interface_id: Must match first field in the output of `show intf brief`. Example: `vlan100`.
        :param pulumi.Input[str] medium: The administrative port medium type. - Choices: `bcast`, `p2p` - Default value: `bcast`
        :param pulumi.Input[int] mtu: Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SviInterfaceState.__new__(_SviInterfaceState)

        __props__.__dict__["admin_state"] = admin_state
        __props__.__dict__["bandwidth"] = bandwidth
        __props__.__dict__["delay"] = delay
        __props__.__dict__["description"] = description
        __props__.__dict__["device"] = device
        __props__.__dict__["interface_id"] = interface_id
        __props__.__dict__["medium"] = medium
        __props__.__dict__["mtu"] = mtu
        return SviInterface(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> pulumi.Output[str]:
        """
        Administrative port state. - Choices: `up`, `down` - Default value: `up`
        """
        return pulumi.get(self, "admin_state")

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output[int]:
        """
        Specifies the administrative port bandwidth. - Range: `1`-`400000000` - Default value: `1000000`
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter
    def delay(self) -> pulumi.Output[int]:
        """
        Specifies the administrative port delay. - Range: `1`-`16777215` - Default value: `1`
        """
        return pulumi.get(self, "delay")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Interface description.
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
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> pulumi.Output[str]:
        """
        Must match first field in the output of `show intf brief`. Example: `vlan100`.
        """
        return pulumi.get(self, "interface_id")

    @property
    @pulumi.getter
    def medium(self) -> pulumi.Output[str]:
        """
        The administrative port medium type. - Choices: `bcast`, `p2p` - Default value: `bcast`
        """
        return pulumi.get(self, "medium")

    @property
    @pulumi.getter
    def mtu(self) -> pulumi.Output[int]:
        """
        Administrative port MTU. - Range: `576`-`9216` - Default value: `1500`
        """
        return pulumi.get(self, "mtu")
