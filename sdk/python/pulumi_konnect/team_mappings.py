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
from ._inputs import *

__all__ = ['TeamMappingsArgs', 'TeamMappings']

@pulumi.input_type
class TeamMappingsArgs:
    def __init__(__self__, *,
                 mappings: Optional[pulumi.Input[Sequence[pulumi.Input['TeamMappingsMappingArgs']]]] = None):
        """
        The set of arguments for constructing a TeamMappings resource.
        :param pulumi.Input[Sequence[pulumi.Input['TeamMappingsMappingArgs']]] mappings: **(Optional, set{mapping})** Configuration block for a mapping.  Can be specified multiple times for each mapping.  Each block supports the fields documented below.
        """
        if mappings is not None:
            pulumi.set(__self__, "mappings", mappings)

    @property
    @pulumi.getter
    def mappings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TeamMappingsMappingArgs']]]]:
        """
        **(Optional, set{mapping})** Configuration block for a mapping.  Can be specified multiple times for each mapping.  Each block supports the fields documented below.
        """
        return pulumi.get(self, "mappings")

    @mappings.setter
    def mappings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TeamMappingsMappingArgs']]]]):
        pulumi.set(self, "mappings", value)


@pulumi.input_type
class _TeamMappingsState:
    def __init__(__self__, *,
                 mappings: Optional[pulumi.Input[Sequence[pulumi.Input['TeamMappingsMappingArgs']]]] = None):
        """
        Input properties used for looking up and filtering TeamMappings resources.
        :param pulumi.Input[Sequence[pulumi.Input['TeamMappingsMappingArgs']]] mappings: **(Optional, set{mapping})** Configuration block for a mapping.  Can be specified multiple times for each mapping.  Each block supports the fields documented below.
        """
        if mappings is not None:
            pulumi.set(__self__, "mappings", mappings)

    @property
    @pulumi.getter
    def mappings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TeamMappingsMappingArgs']]]]:
        """
        **(Optional, set{mapping})** Configuration block for a mapping.  Can be specified multiple times for each mapping.  Each block supports the fields documented below.
        """
        return pulumi.get(self, "mappings")

    @mappings.setter
    def mappings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TeamMappingsMappingArgs']]]]):
        pulumi.set(self, "mappings", value)


class TeamMappings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 mappings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamMappingsMappingArgs']]]]] = None,
                 __props__=None):
        """
        Represents the mappings between an external identity provider group and a Konnect team
        ## Example Usage

        ```python
        import pulumi
        import pulumi_konnect as konnect

        team = konnect.Team("team", description="testing")
        example = konnect.TeamMappings("example", mappings=[konnect.TeamMappingsMappingArgs(
            group="external IdP group",
            team_ids=[team.id],
        )])
        ```

        ## Import

        Team mappings can be imported using a proper value of `id` as described above

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamMappingsMappingArgs']]]] mappings: **(Optional, set{mapping})** Configuration block for a mapping.  Can be specified multiple times for each mapping.  Each block supports the fields documented below.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TeamMappingsArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents the mappings between an external identity provider group and a Konnect team
        ## Example Usage

        ```python
        import pulumi
        import pulumi_konnect as konnect

        team = konnect.Team("team", description="testing")
        example = konnect.TeamMappings("example", mappings=[konnect.TeamMappingsMappingArgs(
            group="external IdP group",
            team_ids=[team.id],
        )])
        ```

        ## Import

        Team mappings can be imported using a proper value of `id` as described above

        :param str resource_name: The name of the resource.
        :param TeamMappingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TeamMappingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 mappings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamMappingsMappingArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TeamMappingsArgs.__new__(TeamMappingsArgs)

            __props__.__dict__["mappings"] = mappings
        super(TeamMappings, __self__).__init__(
            'konnect:index/teamMappings:TeamMappings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            mappings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamMappingsMappingArgs']]]]] = None) -> 'TeamMappings':
        """
        Get an existing TeamMappings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TeamMappingsMappingArgs']]]] mappings: **(Optional, set{mapping})** Configuration block for a mapping.  Can be specified multiple times for each mapping.  Each block supports the fields documented below.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TeamMappingsState.__new__(_TeamMappingsState)

        __props__.__dict__["mappings"] = mappings
        return TeamMappings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def mappings(self) -> pulumi.Output[Optional[Sequence['outputs.TeamMappingsMapping']]]:
        """
        **(Optional, set{mapping})** Configuration block for a mapping.  Can be specified multiple times for each mapping.  Each block supports the fields documented below.
        """
        return pulumi.get(self, "mappings")

