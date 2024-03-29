# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str],
                 full_name: Optional[pulumi.Input[str]] = None,
                 preferred_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[str] email: **(Required, String)** The email of the user.
        :param pulumi.Input[str] full_name: **(Optional, String)** The full name of the user.
        :param pulumi.Input[str] preferred_name: **(Optional, String)** The preferred name of the user.
        """
        pulumi.set(__self__, "email", email)
        if full_name is not None:
            pulumi.set(__self__, "full_name", full_name)
        if preferred_name is not None:
            pulumi.set(__self__, "preferred_name", preferred_name)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        **(Required, String)** The email of the user.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> Optional[pulumi.Input[str]]:
        """
        **(Optional, String)** The full name of the user.
        """
        return pulumi.get(self, "full_name")

    @full_name.setter
    def full_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "full_name", value)

    @property
    @pulumi.getter(name="preferredName")
    def preferred_name(self) -> Optional[pulumi.Input[str]]:
        """
        **(Optional, String)** The preferred name of the user.
        """
        return pulumi.get(self, "preferred_name")

    @preferred_name.setter
    def preferred_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "preferred_name", value)


@pulumi.input_type
class _UserState:
    def __init__(__self__, *,
                 active: Optional[pulumi.Input[bool]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 preferred_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering User resources.
        :param pulumi.Input[bool] active: **(Boolean)** Whether the user is active.
        :param pulumi.Input[str] email: **(Required, String)** The email of the user.
        :param pulumi.Input[str] full_name: **(Optional, String)** The full name of the user.
        :param pulumi.Input[str] preferred_name: **(Optional, String)** The preferred name of the user.
        """
        if active is not None:
            pulumi.set(__self__, "active", active)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if full_name is not None:
            pulumi.set(__self__, "full_name", full_name)
        if preferred_name is not None:
            pulumi.set(__self__, "preferred_name", preferred_name)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[bool]]:
        """
        **(Boolean)** Whether the user is active.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        **(Required, String)** The email of the user.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> Optional[pulumi.Input[str]]:
        """
        **(Optional, String)** The full name of the user.
        """
        return pulumi.get(self, "full_name")

    @full_name.setter
    def full_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "full_name", value)

    @property
    @pulumi.getter(name="preferredName")
    def preferred_name(self) -> Optional[pulumi.Input[str]]:
        """
        **(Optional, String)** The preferred name of the user.
        """
        return pulumi.get(self, "preferred_name")

    @preferred_name.setter
    def preferred_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "preferred_name", value)


class User(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 preferred_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Represents a user
        ## Example Usage

        ```python
        import pulumi
        import pulumi_konnect as konnect

        example = konnect.User("example",
            email="Joe.Burrow@example.com",
            full_name="Joe Burrow",
            preferred_name="Joe")
        ```

        ## Import

        Users can be imported using a proper value of `id` as described above

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] email: **(Required, String)** The email of the user.
        :param pulumi.Input[str] full_name: **(Optional, String)** The full name of the user.
        :param pulumi.Input[str] preferred_name: **(Optional, String)** The preferred name of the user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a user
        ## Example Usage

        ```python
        import pulumi
        import pulumi_konnect as konnect

        example = konnect.User("example",
            email="Joe.Burrow@example.com",
            full_name="Joe Burrow",
            preferred_name="Joe")
        ```

        ## Import

        Users can be imported using a proper value of `id` as described above

        :param str resource_name: The name of the resource.
        :param UserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 full_name: Optional[pulumi.Input[str]] = None,
                 preferred_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserArgs.__new__(UserArgs)

            if email is None and not opts.urn:
                raise TypeError("Missing required property 'email'")
            __props__.__dict__["email"] = email
            __props__.__dict__["full_name"] = full_name
            __props__.__dict__["preferred_name"] = preferred_name
            __props__.__dict__["active"] = None
        super(User, __self__).__init__(
            'konnect:index/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active: Optional[pulumi.Input[bool]] = None,
            email: Optional[pulumi.Input[str]] = None,
            full_name: Optional[pulumi.Input[str]] = None,
            preferred_name: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: **(Boolean)** Whether the user is active.
        :param pulumi.Input[str] email: **(Required, String)** The email of the user.
        :param pulumi.Input[str] full_name: **(Optional, String)** The full name of the user.
        :param pulumi.Input[str] preferred_name: **(Optional, String)** The preferred name of the user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserState.__new__(_UserState)

        __props__.__dict__["active"] = active
        __props__.__dict__["email"] = email
        __props__.__dict__["full_name"] = full_name
        __props__.__dict__["preferred_name"] = preferred_name
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def active(self) -> pulumi.Output[bool]:
        """
        **(Boolean)** Whether the user is active.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[str]:
        """
        **(Required, String)** The email of the user.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> pulumi.Output[Optional[str]]:
        """
        **(Optional, String)** The full name of the user.
        """
        return pulumi.get(self, "full_name")

    @property
    @pulumi.getter(name="preferredName")
    def preferred_name(self) -> pulumi.Output[Optional[str]]:
        """
        **(Optional, String)** The preferred name of the user.
        """
        return pulumi.get(self, "preferred_name")

