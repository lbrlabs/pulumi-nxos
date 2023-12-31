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
    'GetHmmInterfaceResult',
    'AwaitableGetHmmInterfaceResult',
    'get_hmm_interface',
    'get_hmm_interface_output',
]

@pulumi.output_type
class GetHmmInterfaceResult:
    """
    A collection of values returned by getHmmInterface.
    """
    def __init__(__self__, admin_state=None, device=None, id=None, interface_id=None, mode=None):
        if admin_state and not isinstance(admin_state, str):
            raise TypeError("Expected argument 'admin_state' to be a str")
        pulumi.set(__self__, "admin_state", admin_state)
        if device and not isinstance(device, str):
            raise TypeError("Expected argument 'device' to be a str")
        pulumi.set(__self__, "device", device)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if interface_id and not isinstance(interface_id, str):
            raise TypeError("Expected argument 'interface_id' to be a str")
        pulumi.set(__self__, "interface_id", interface_id)
        if mode and not isinstance(mode, str):
            raise TypeError("Expected argument 'mode' to be a str")
        pulumi.set(__self__, "mode", mode)

    @property
    @pulumi.getter(name="adminState")
    def admin_state(self) -> str:
        """
        Administrative state.
        """
        return pulumi.get(self, "admin_state")

    @property
    @pulumi.getter
    def device(self) -> Optional[str]:
        """
        A device name from the provider configuration.
        """
        return pulumi.get(self, "device")

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
        Must match first field in the output of `show intf brief`. Example: `vlan10`.
        """
        return pulumi.get(self, "interface_id")

    @property
    @pulumi.getter
    def mode(self) -> str:
        """
        HMM Fabric Forwarding mode information for the interface.
        """
        return pulumi.get(self, "mode")


class AwaitableGetHmmInterfaceResult(GetHmmInterfaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetHmmInterfaceResult(
            admin_state=self.admin_state,
            device=self.device,
            id=self.id,
            interface_id=self.interface_id,
            mode=self.mode)


def get_hmm_interface(device: Optional[str] = None,
                      interface_id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetHmmInterfaceResult:
    """
    This data source can read the HMM Fabric Forwarding mode information related to an Interface.

    - API Documentation: [hmmFwdIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Host%20Mobility/hmm:FwdIf/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_hmm_interface(interface_id="vlan10")
    ```


    :param str device: A device name from the provider configuration.
    :param str interface_id: Must match first field in the output of `show intf brief`. Example: `vlan10`.
    """
    __args__ = dict()
    __args__['device'] = device
    __args__['interfaceId'] = interface_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('nxos:index/getHmmInterface:getHmmInterface', __args__, opts=opts, typ=GetHmmInterfaceResult).value

    return AwaitableGetHmmInterfaceResult(
        admin_state=pulumi.get(__ret__, 'admin_state'),
        device=pulumi.get(__ret__, 'device'),
        id=pulumi.get(__ret__, 'id'),
        interface_id=pulumi.get(__ret__, 'interface_id'),
        mode=pulumi.get(__ret__, 'mode'))


@_utilities.lift_output_func(get_hmm_interface)
def get_hmm_interface_output(device: Optional[pulumi.Input[Optional[str]]] = None,
                             interface_id: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetHmmInterfaceResult]:
    """
    This data source can read the HMM Fabric Forwarding mode information related to an Interface.

    - API Documentation: [hmmFwdIf](https://pubhub.devnetcloud.com/media/dme-docs-10-2-2/docs/Host%20Mobility/hmm:FwdIf/)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_nxos as nxos

    example = nxos.get_hmm_interface(interface_id="vlan10")
    ```


    :param str device: A device name from the provider configuration.
    :param str interface_id: Must match first field in the output of `show intf brief`. Example: `vlan10`.
    """
    ...
