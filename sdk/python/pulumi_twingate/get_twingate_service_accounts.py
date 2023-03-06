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

__all__ = [
    'GetTwingateServiceAccountsResult',
    'AwaitableGetTwingateServiceAccountsResult',
    'get_twingate_service_accounts',
    'get_twingate_service_accounts_output',
]

@pulumi.output_type
class GetTwingateServiceAccountsResult:
    """
    A collection of values returned by getTwingateServiceAccounts.
    """
    def __init__(__self__, id=None, name=None, service_accounts=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service_accounts and not isinstance(service_accounts, list):
            raise TypeError("Expected argument 'service_accounts' to be a list")
        pulumi.set(__self__, "service_accounts", service_accounts)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceAccounts")
    def service_accounts(self) -> Optional[Sequence['outputs.GetTwingateServiceAccountsServiceAccountResult']]:
        return pulumi.get(self, "service_accounts")


class AwaitableGetTwingateServiceAccountsResult(GetTwingateServiceAccountsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTwingateServiceAccountsResult(
            id=self.id,
            name=self.name,
            service_accounts=self.service_accounts)


def get_twingate_service_accounts(name: Optional[str] = None,
                                  service_accounts: Optional[Sequence[pulumi.InputType['GetTwingateServiceAccountsServiceAccountArgs']]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTwingateServiceAccountsResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['serviceAccounts'] = service_accounts
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('twingate:index/getTwingateServiceAccounts:getTwingateServiceAccounts', __args__, opts=opts, typ=GetTwingateServiceAccountsResult).value

    return AwaitableGetTwingateServiceAccountsResult(
        id=__ret__.id,
        name=__ret__.name,
        service_accounts=__ret__.service_accounts)


@_utilities.lift_output_func(get_twingate_service_accounts)
def get_twingate_service_accounts_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                                         service_accounts: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetTwingateServiceAccountsServiceAccountArgs']]]]] = None,
                                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTwingateServiceAccountsResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...