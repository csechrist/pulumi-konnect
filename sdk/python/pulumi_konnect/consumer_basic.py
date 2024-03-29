# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ConsumerBasicArgs', 'ConsumerBasic']

@pulumi.input_type
class ConsumerBasicArgs:
    def __init__(__self__, *,
                 consumer_id: pulumi.Input[str],
                 control_plane_id: pulumi.Input[str],
                 password: pulumi.Input[str],
                 username: pulumi.Input[str]):
        """
        The set of arguments for constructing a ConsumerBasic resource.
        :param pulumi.Input[str] consumer_id: **(Required, String)** The id of the consumer.
        :param pulumi.Input[str] control_plane_id: **(Required, String)** The id of the control plane.
        :param pulumi.Input[str] password: **(Required, String)** The password value.
        :param pulumi.Input[str] username: **(Required, String)** The username value.
        """
        pulumi.set(__self__, "consumer_id", consumer_id)
        pulumi.set(__self__, "control_plane_id", control_plane_id)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="consumerId")
    def consumer_id(self) -> pulumi.Input[str]:
        """
        **(Required, String)** The id of the consumer.
        """
        return pulumi.get(self, "consumer_id")

    @consumer_id.setter
    def consumer_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "consumer_id", value)

    @property
    @pulumi.getter(name="controlPlaneId")
    def control_plane_id(self) -> pulumi.Input[str]:
        """
        **(Required, String)** The id of the control plane.
        """
        return pulumi.get(self, "control_plane_id")

    @control_plane_id.setter
    def control_plane_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "control_plane_id", value)

    @property
    @pulumi.getter
    def password(self) -> pulumi.Input[str]:
        """
        **(Required, String)** The password value.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[str]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        **(Required, String)** The username value.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)


@pulumi.input_type
class _ConsumerBasicState:
    def __init__(__self__, *,
                 basic_id: Optional[pulumi.Input[str]] = None,
                 consumer_id: Optional[pulumi.Input[str]] = None,
                 control_plane_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 password_hash: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ConsumerBasic resources.
        :param pulumi.Input[str] basic_id: **(String)** Id of the consumer basic auth alone
        :param pulumi.Input[str] consumer_id: **(Required, String)** The id of the consumer.
        :param pulumi.Input[str] control_plane_id: **(Required, String)** The id of the control plane.
        :param pulumi.Input[str] password: **(Required, String)** The password value.
        :param pulumi.Input[str] password_hash: **(String)** Hash of the password
        :param pulumi.Input[str] username: **(Required, String)** The username value.
        """
        if basic_id is not None:
            pulumi.set(__self__, "basic_id", basic_id)
        if consumer_id is not None:
            pulumi.set(__self__, "consumer_id", consumer_id)
        if control_plane_id is not None:
            pulumi.set(__self__, "control_plane_id", control_plane_id)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if password_hash is not None:
            pulumi.set(__self__, "password_hash", password_hash)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="basicId")
    def basic_id(self) -> Optional[pulumi.Input[str]]:
        """
        **(String)** Id of the consumer basic auth alone
        """
        return pulumi.get(self, "basic_id")

    @basic_id.setter
    def basic_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "basic_id", value)

    @property
    @pulumi.getter(name="consumerId")
    def consumer_id(self) -> Optional[pulumi.Input[str]]:
        """
        **(Required, String)** The id of the consumer.
        """
        return pulumi.get(self, "consumer_id")

    @consumer_id.setter
    def consumer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consumer_id", value)

    @property
    @pulumi.getter(name="controlPlaneId")
    def control_plane_id(self) -> Optional[pulumi.Input[str]]:
        """
        **(Required, String)** The id of the control plane.
        """
        return pulumi.get(self, "control_plane_id")

    @control_plane_id.setter
    def control_plane_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "control_plane_id", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        """
        **(Required, String)** The password value.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter(name="passwordHash")
    def password_hash(self) -> Optional[pulumi.Input[str]]:
        """
        **(String)** Hash of the password
        """
        return pulumi.get(self, "password_hash")

    @password_hash.setter
    def password_hash(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password_hash", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        **(Required, String)** The username value.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class ConsumerBasic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_id: Optional[pulumi.Input[str]] = None,
                 control_plane_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Represents a basic auth credential for a consumer within a control plane
        ## Example Usage

        ```python
        import pulumi
        import pulumi_konnect as konnect

        control_plane = konnect.get_control_plane(name="TestControlPlane")
        consumer = konnect.get_consumer(control_plane_id=control_plane.id,
            search_username="Bob")
        example = konnect.ConsumerBasic("example",
            control_plane_id=control_plane.id,
            consumer_id=consumer.consumer_id,
            username="my-username",
            password="my-password")
        ```

        ## Import

        Consumer basics can be imported using a proper value of `id` as described above

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consumer_id: **(Required, String)** The id of the consumer.
        :param pulumi.Input[str] control_plane_id: **(Required, String)** The id of the control plane.
        :param pulumi.Input[str] password: **(Required, String)** The password value.
        :param pulumi.Input[str] username: **(Required, String)** The username value.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ConsumerBasicArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Represents a basic auth credential for a consumer within a control plane
        ## Example Usage

        ```python
        import pulumi
        import pulumi_konnect as konnect

        control_plane = konnect.get_control_plane(name="TestControlPlane")
        consumer = konnect.get_consumer(control_plane_id=control_plane.id,
            search_username="Bob")
        example = konnect.ConsumerBasic("example",
            control_plane_id=control_plane.id,
            consumer_id=consumer.consumer_id,
            username="my-username",
            password="my-password")
        ```

        ## Import

        Consumer basics can be imported using a proper value of `id` as described above

        :param str resource_name: The name of the resource.
        :param ConsumerBasicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ConsumerBasicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consumer_id: Optional[pulumi.Input[str]] = None,
                 control_plane_id: Optional[pulumi.Input[str]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ConsumerBasicArgs.__new__(ConsumerBasicArgs)

            if consumer_id is None and not opts.urn:
                raise TypeError("Missing required property 'consumer_id'")
            __props__.__dict__["consumer_id"] = consumer_id
            if control_plane_id is None and not opts.urn:
                raise TypeError("Missing required property 'control_plane_id'")
            __props__.__dict__["control_plane_id"] = control_plane_id
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = password
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
            __props__.__dict__["basic_id"] = None
            __props__.__dict__["password_hash"] = None
        super(ConsumerBasic, __self__).__init__(
            'konnect:index/consumerBasic:ConsumerBasic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            basic_id: Optional[pulumi.Input[str]] = None,
            consumer_id: Optional[pulumi.Input[str]] = None,
            control_plane_id: Optional[pulumi.Input[str]] = None,
            password: Optional[pulumi.Input[str]] = None,
            password_hash: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'ConsumerBasic':
        """
        Get an existing ConsumerBasic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] basic_id: **(String)** Id of the consumer basic auth alone
        :param pulumi.Input[str] consumer_id: **(Required, String)** The id of the consumer.
        :param pulumi.Input[str] control_plane_id: **(Required, String)** The id of the control plane.
        :param pulumi.Input[str] password: **(Required, String)** The password value.
        :param pulumi.Input[str] password_hash: **(String)** Hash of the password
        :param pulumi.Input[str] username: **(Required, String)** The username value.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ConsumerBasicState.__new__(_ConsumerBasicState)

        __props__.__dict__["basic_id"] = basic_id
        __props__.__dict__["consumer_id"] = consumer_id
        __props__.__dict__["control_plane_id"] = control_plane_id
        __props__.__dict__["password"] = password
        __props__.__dict__["password_hash"] = password_hash
        __props__.__dict__["username"] = username
        return ConsumerBasic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="basicId")
    def basic_id(self) -> pulumi.Output[str]:
        """
        **(String)** Id of the consumer basic auth alone
        """
        return pulumi.get(self, "basic_id")

    @property
    @pulumi.getter(name="consumerId")
    def consumer_id(self) -> pulumi.Output[str]:
        """
        **(Required, String)** The id of the consumer.
        """
        return pulumi.get(self, "consumer_id")

    @property
    @pulumi.getter(name="controlPlaneId")
    def control_plane_id(self) -> pulumi.Output[str]:
        """
        **(Required, String)** The id of the control plane.
        """
        return pulumi.get(self, "control_plane_id")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[str]:
        """
        **(Required, String)** The password value.
        """
        return pulumi.get(self, "password")

    @property
    @pulumi.getter(name="passwordHash")
    def password_hash(self) -> pulumi.Output[str]:
        """
        **(String)** Hash of the password
        """
        return pulumi.get(self, "password_hash")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        **(Required, String)** The username value.
        """
        return pulumi.get(self, "username")
