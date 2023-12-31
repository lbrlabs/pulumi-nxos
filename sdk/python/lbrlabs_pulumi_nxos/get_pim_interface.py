# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetPimInterfaceResult',
    'AwaitableGetPimInterfaceResult',
    'get_pim_interface',
    'get_pim_interface_output',
]

@pulumi.output_type
class GetPimInterfaceResult:
    """
    A collection of values returned by getPimInterface.
    """
    def __init__(__self__, admin_state=None, bfd=None, device=None, dr_priority=None, id=None, interface_id=None, passive=None, sparse_mode=None, vrf_name=None):
        if admin_state and not isinstance(admin_state, str):
            raise TypeError("Expected argument 'admin_state' to be a str")
        pulumi.set(__self__, "admin_state", admin_state)
        if bfd and not isinstance(bfd, str):
            raise TypeError("Expected argument 'bfd' to be a str")
        pulumi.set(__self__, "bfd", bfd)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if dr_priority and not isinstance(dr_priority, int):
            raise TypeError("Expected argument 'dr_priority' to be a int")
        pulumi.set(__self__, "dr_priority", dr_priority)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface_id and not isinstance(interface_id, str):
            raise TypeError("Expected argument 'interface_id' to be a str")
        pulumi.set(__self__, "interface_id", interface_id)
        if passive and not isinstance(passive, bool):
            raise TypeError("Expected argument 'passive' to be a bool")
        pulumi.set(__self__, "passive", passive)
        if sparse_mode and not isinstance(sparse_mode, bool):
            raise TypeError("Expected argument 'sparse_mode' to be a bool")
        pulumi.set(__self__, "sparse_mode", sparse_mode)
        if vrf_name and not isinstance(vrf_name, str):
            raise TypeError("Expected argument 'vrf_name' to be a str")
        pulumi.set(__self__, "vrf_name", vrf_name)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> str:
        """
        Administrative state.
        """
        return pulumi.get(self, "admin_state")

    @property
    @pulumi.getter
    def bfd(self) -> str:
        """
        BFD.
        """
        return pulumi.get(self, "bfd")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

    @property
    @pulumi.getter(name="drPriority")
    def dr_priority(self) -> int:
        """
        Designated Router priority level.
        """
        return pulumi.get(self, "dr_priority")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The distinguished name of the object.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="interfaceId")
    def interface_id(self) -> str:
        """
        Must match first field in the output of `show intf brief`. Example: `eth1/1`.
        """
        return pulumi.get(self, "interface_id")

    @property
    @pulumi.getter
    def passive(self) -> bool:
        """
        Passive interface.
        """
        return pulumi.get(self, "passive")

    @property
    @pulumi.getter(name="sparseMode")
    def sparse_mode(self) -> bool:
        """
        Sparse mode.
        """
        return pulumi.get(self, "sparse_mode")

    @property
    @pulumi.getter(name="vrfName")
    def vrf_name(self) -> str:
        """
        VRF name.
        """
        return pulumi.get(self, "vrf_name")


class AwaitableGetPimInterfaceResult(GetPimInterfaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPimInterfaceResult(
            admin_state=self.admin_state,
            bfd=self.bfd,
            device=self.device,
            dr_priority=self.dr_priority,
            id=self.id,
            interface_id=self.interface_id,
            passive=self.passive,
            sparse_mode=self.sparse_mode,
            vrf_name=self.vrf_name)


def get_pim_interface(device: Optional[str] = None,
                      interface_id: Optional[str] = None,
                      vrf_name: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPimInterfaceResult:
    """
    This data source can read the PIM interface configuration.

    - API Documentation: [pimIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:If/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_pim_interface(interface_id="eth1/10",
        vrf_name="default")
    ```


    :param str device: A device name from the provider configuration.
    :param str interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
    :param str vrf_name: VRF name.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['interfaceId'] = interface_id
    __args__['vrfName'] = vrf_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getPimInterface:getPimInterface', __args__, opts=opts, typ=GetPimInterfaceResult).value

    return AwaitableGetPimInterfaceResult(
        admin_state=pulumi.get(__ret__, 'admin_state'),
        bfd=pulumi.get(__ret__, 'bfd'),
        device=pulumi.get(__ret__, 'device'),
        dr_priority=pulumi.get(__ret__, 'dr_priority'),
        id=pulumi.get(__ret__, 'id'),
        interface_id=pulumi.get(__ret__, 'interface_id'),
        passive=pulumi.get(__ret__, 'passive'),
        sparse_mode=pulumi.get(__ret__, 'sparse_mode'),
        vrf_name=pulumi.get(__ret__, 'vrf_name'))


@_utilities.lift_output_func(get_pim_interface)
def get_pim_interface_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                             interface_id: Optional[pulumi.Input[str]] = None,
                             vrf_name: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPimInterfaceResult]:
    """
    This data source can read the PIM interface configuration.

    - API Documentation: [pimIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Layer%203/pim:If/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_pim_interface(interface_id="eth1/10",
        vrf_name="default")
    ```


    :param str device: A device name from the provider configuration.
    :param str interface_id: Must match first field in the output of `show intf brief`. Example: `eth1/1`.
    :param str vrf_name: VRF name.
    """
    ...
