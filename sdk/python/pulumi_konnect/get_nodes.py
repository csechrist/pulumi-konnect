# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetNodesResult',
    'AwaitableGetNodesResult',
    'get_nodes',
    'get_nodes_output',
]

@pulumi.output_type
class GetNodesResult:
    """
    A collection of values returned by getNodes.
    """
    def __init__(__self__, control_plane_id=None, id=None, nodes=None):
        if control_plane_id and not isinstance(control_plane_id, str):
            raise TypeError("Expected argument 'control_plane_id' to be a str")
        pulumi.set(__self__, "control_plane_id", control_plane_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if nodes and not isinstance(nodes, list):
            raise TypeError("Expected argument 'nodes' to be a list")
        pulumi.set(__self__, "nodes", nodes)

    @property
    @pulumi.getter(name="controlPlaneId")
    def control_plane_id(self) -> str:
        return pulumi.get(self, "control_plane_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def nodes(self) -> Sequence['outputs.GetNodesNodeResult']:
        """
        **(set{node})** Set of nodes belonging to control plane
        """
        return pulumi.get(self, "nodes")


class AwaitableGetNodesResult(GetNodesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNodesResult(
            control_plane_id=self.control_plane_id,
            id=self.id,
            nodes=self.nodes)


def get_nodes(control_plane_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNodesResult:
    """
    Represents all nodes of a control plane
    ## Example Usage

    ```python
    import pulumi
    import pulumi_konnect as konnect

    control_plane = konnect.get_control_plane(name="TestControlPlane")
    example = konnect.get_nodes(control_plane_id=control_plane.id)
    ```


    :param str control_plane_id: **(Required, String)** The id of the parent control plane.
    """
    __args__ = dict()
    __args__['controlPlaneId'] = control_plane_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('konnect:index/getNodes:getNodes', __args__, opts=opts, typ=GetNodesResult).value

    return AwaitableGetNodesResult(
        control_plane_id=pulumi.get(__ret__, 'control_plane_id'),
        id=pulumi.get(__ret__, 'id'),
        nodes=pulumi.get(__ret__, 'nodes'))


@_utilities.lift_output_func(get_nodes)
def get_nodes_output(control_plane_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNodesResult]:
    """
    Represents all nodes of a control plane
    ## Example Usage

    ```python
    import pulumi
    import pulumi_konnect as konnect

    control_plane = konnect.get_control_plane(name="TestControlPlane")
    example = konnect.get_nodes(control_plane_id=control_plane.id)
    ```


    :param str control_plane_id: **(Required, String)** The id of the parent control plane.
    """
    ...
