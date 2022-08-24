# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 api_token: Optional[pulumi.Input[str]] = None,
                 http_max_retry: Optional[pulumi.Input[int]] = None,
                 http_timeout: Optional[pulumi.Input[int]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] api_token: The access key for API operations. You can retrieve this from the Twingate Admin Console
               ([documentation](https://docs.twingate.com/docs/api-overview)). Alternatively, this can be specified using the
               TWINGATE_API_TOKEN environment variable.
        :param pulumi.Input[int] http_max_retry: Specifies a retry limit for the http requests made. This default value is 10. Alternatively, this can be specified using
               the TWINGATE_HTTP_MAX_RETRY environment variable
        :param pulumi.Input[int] http_timeout: Specifies a time limit in seconds for the http requests made. The default value is 10 seconds. Alternatively, this can
               be specified using the TWINGATE_HTTP_TIMEOUT environment variable
        :param pulumi.Input[str] network: Your Twingate network ID for API operations. You can find it in the Admin Console URL, for example:
               `autoco.twingate.com`, where `autoco` is your network ID Alternatively, this can be specified using the TWINGATE_NETWORK
               environment variable.
        :param pulumi.Input[str] url: The default is 'twingate.com' This is optional and shouldn't be changed under normal circumstances.
        """
        if api_token is not None:
            pulumi.set(__self__, "api_token", api_token)
        if http_max_retry is None:
            http_max_retry = 5
        if http_max_retry is not None:
            pulumi.set(__self__, "http_max_retry", http_max_retry)
        if http_timeout is None:
            http_timeout = 10
        if http_timeout is not None:
            pulumi.set(__self__, "http_timeout", http_timeout)
        if network is not None:
            pulumi.set(__self__, "network", network)
        if url is not None:
            pulumi.set(__self__, "url", url)

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> Optional[pulumi.Input[str]]:
        """
        The access key for API operations. You can retrieve this from the Twingate Admin Console
        ([documentation](https://docs.twingate.com/docs/api-overview)). Alternatively, this can be specified using the
        TWINGATE_API_TOKEN environment variable.
        """
        return pulumi.get(self, "api_token")

    @api_token.setter
    def api_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_token", value)

    @property
    @pulumi.getter(name="httpMaxRetry")
    def http_max_retry(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies a retry limit for the http requests made. This default value is 10. Alternatively, this can be specified using
        the TWINGATE_HTTP_MAX_RETRY environment variable
        """
        return pulumi.get(self, "http_max_retry")

    @http_max_retry.setter
    def http_max_retry(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "http_max_retry", value)

    @property
    @pulumi.getter(name="httpTimeout")
    def http_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies a time limit in seconds for the http requests made. The default value is 10 seconds. Alternatively, this can
        be specified using the TWINGATE_HTTP_TIMEOUT environment variable
        """
        return pulumi.get(self, "http_timeout")

    @http_timeout.setter
    def http_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "http_timeout", value)

    @property
    @pulumi.getter
    def network(self) -> Optional[pulumi.Input[str]]:
        """
        Your Twingate network ID for API operations. You can find it in the Admin Console URL, for example:
        `autoco.twingate.com`, where `autoco` is your network ID Alternatively, this can be specified using the TWINGATE_NETWORK
        environment variable.
        """
        return pulumi.get(self, "network")

    @network.setter
    def network(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The default is 'twingate.com' This is optional and shouldn't be changed under normal circumstances.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_token: Optional[pulumi.Input[str]] = None,
                 http_max_retry: Optional[pulumi.Input[int]] = None,
                 http_timeout: Optional[pulumi.Input[int]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the twingate package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_token: The access key for API operations. You can retrieve this from the Twingate Admin Console
               ([documentation](https://docs.twingate.com/docs/api-overview)). Alternatively, this can be specified using the
               TWINGATE_API_TOKEN environment variable.
        :param pulumi.Input[int] http_max_retry: Specifies a retry limit for the http requests made. This default value is 10. Alternatively, this can be specified using
               the TWINGATE_HTTP_MAX_RETRY environment variable
        :param pulumi.Input[int] http_timeout: Specifies a time limit in seconds for the http requests made. The default value is 10 seconds. Alternatively, this can
               be specified using the TWINGATE_HTTP_TIMEOUT environment variable
        :param pulumi.Input[str] network: Your Twingate network ID for API operations. You can find it in the Admin Console URL, for example:
               `autoco.twingate.com`, where `autoco` is your network ID Alternatively, this can be specified using the TWINGATE_NETWORK
               environment variable.
        :param pulumi.Input[str] url: The default is 'twingate.com' This is optional and shouldn't be changed under normal circumstances.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the twingate package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_token: Optional[pulumi.Input[str]] = None,
                 http_max_retry: Optional[pulumi.Input[int]] = None,
                 http_timeout: Optional[pulumi.Input[int]] = None,
                 network: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["api_token"] = api_token
            if http_max_retry is None:
                http_max_retry = 5
            __props__.__dict__["http_max_retry"] = pulumi.Output.from_input(http_max_retry).apply(pulumi.runtime.to_json) if http_max_retry is not None else None
            if http_timeout is None:
                http_timeout = 10
            __props__.__dict__["http_timeout"] = pulumi.Output.from_input(http_timeout).apply(pulumi.runtime.to_json) if http_timeout is not None else None
            __props__.__dict__["network"] = network
            __props__.__dict__["url"] = url
        super(Provider, __self__).__init__(
            'twingate',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> pulumi.Output[Optional[str]]:
        """
        The access key for API operations. You can retrieve this from the Twingate Admin Console
        ([documentation](https://docs.twingate.com/docs/api-overview)). Alternatively, this can be specified using the
        TWINGATE_API_TOKEN environment variable.
        """
        return pulumi.get(self, "api_token")

    @property
    @pulumi.getter
    def network(self) -> pulumi.Output[Optional[str]]:
        """
        Your Twingate network ID for API operations. You can find it in the Admin Console URL, for example:
        `autoco.twingate.com`, where `autoco` is your network ID Alternatively, this can be specified using the TWINGATE_NETWORK
        environment variable.
        """
        return pulumi.get(self, "network")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[Optional[str]]:
        """
        The default is 'twingate.com' This is optional and shouldn't be changed under normal circumstances.
        """
        return pulumi.get(self, "url")

