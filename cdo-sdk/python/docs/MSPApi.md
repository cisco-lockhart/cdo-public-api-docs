# cdo_sdk_python.MSPApi

All URIs are relative to *https://edge.us.SCC.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_active_directory_groups_to_tenant_in_msp_portal**](MSPApi.md#add_active_directory_groups_to_tenant_in_msp_portal) | **POST** /v1/msp/tenants/{tenantUid}/users/groups | Add Active Directory Groups to SCC tenant in MSP Portal
[**add_msp_tenant**](MSPApi.md#add_msp_tenant) | **POST** /v1/msp/tenants | Add Tenant to MSP Portal (API token)
[**add_tenant_to_msp_portal**](MSPApi.md#add_tenant_to_msp_portal) | **POST** /v1/msp/tenants/{tenantUid} | Add tenant to MSP Portal
[**add_users_to_tenant_in_msp_portal**](MSPApi.md#add_users_to_tenant_in_msp_portal) | **POST** /v1/msp/tenants/{tenantUid}/users | Add users to SCC tenant in MSP Portal
[**create_tenant**](MSPApi.md#create_tenant) | **POST** /v1/msp/tenants/create | Create SCC Tenant
[**delete_user_groups_from_tenant_in_msp_portal**](MSPApi.md#delete_user_groups_from_tenant_in_msp_portal) | **POST** /v1/msp/tenants/{tenantUid}/users/groups/delete | Remove user groups from SCC tenant in MSP Portal
[**delete_users_from_tenant_in_msp_portal**](MSPApi.md#delete_users_from_tenant_in_msp_portal) | **POST** /v1/msp/tenants/{tenantUid}/users/delete | Remove users from SCC tenant in MSP Portal
[**enable_multicloud_defense_for_tenant_in_msp_portal**](MSPApi.md#enable_multicloud_defense_for_tenant_in_msp_portal) | **POST** /v1/msp/tenants/{tenantUid}/mcd | Enable Multicloud Defense for SCC tenant in MSP Portal
[**generate_api_token_for_user_in_tenant**](MSPApi.md#generate_api_token_for_user_in_tenant) | **POST** /v1/msp/tenants/{tenantUid}/users/{apiUserUid}/token | Generate token for API-only user on tenant managed by MSP portal
[**get_msp_managed_tenant**](MSPApi.md#get_msp_managed_tenant) | **GET** /v1/msp/tenants/{tenantUid} | Get SCC tenant managed by MSP Portal
[**get_msp_managed_tenants**](MSPApi.md#get_msp_managed_tenants) | **GET** /v1/msp/tenants | Get SCC tenants managed by MSP Portal
[**get_users_in_tenant_in_msp_portal**](MSPApi.md#get_users_in_tenant_in_msp_portal) | **GET** /v1/msp/tenants/{tenantUid}/users | Get Users associated with tenant in MSP portal
[**provision_cd_fmc_for_tenant_in_msp_portal**](MSPApi.md#provision_cd_fmc_for_tenant_in_msp_portal) | **POST** /v1/msp/tenants/{tenantUid}/cdfmc | Provision cdFMC for SCC tenant in MSP Portal
[**remove_tenant_from_msp_portal**](MSPApi.md#remove_tenant_from_msp_portal) | **DELETE** /v1/msp/tenants/{tenantUid} | Remove tenant from MSP Portal


# **add_active_directory_groups_to_tenant_in_msp_portal**
> CdoTransaction add_active_directory_groups_to_tenant_in_msp_portal(tenant_uid, active_directory_group_create_or_update_input)

Add Active Directory Groups to SCC tenant in MSP Portal

This is an asynchronous operation to add a list of Active Directory Groups to a tenant associated with the MSP Portal.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.active_directory_group_create_or_update_input import ActiveDirectoryGroupCreateOrUpdateInput
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant to which the user groups will be added. You can add a maximum of 50 user groups at a time.
    active_directory_group_create_or_update_input = [cdo_sdk_python.ActiveDirectoryGroupCreateOrUpdateInput()] # List[ActiveDirectoryGroupCreateOrUpdateInput] | 

    try:
        # Add Active Directory Groups to SCC tenant in MSP Portal
        api_response = api_instance.add_active_directory_groups_to_tenant_in_msp_portal(tenant_uid, active_directory_group_create_or_update_input)
        print("The response of MSPApi->add_active_directory_groups_to_tenant_in_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPApi->add_active_directory_groups_to_tenant_in_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant to which the user groups will be added. You can add a maximum of 50 user groups at a time. | 
 **active_directory_group_create_or_update_input** | [**List[ActiveDirectoryGroupCreateOrUpdateInput]**](ActiveDirectoryGroupCreateOrUpdateInput.md)|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | SCC Transaction object that can be used to track the status of the operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_msp_tenant**
> StatusInfo add_msp_tenant(msp_add_tenant_input)

Add Tenant to MSP Portal (API token)

Add a tenant to the MSP Portal by providing an API token for the tenant. Use this endpoint to add a tenant that the user performing the operation is not associated with, or a tenant in a different region to that of the MSP Portal. Note: This endpoint can only be executed by a super-admin in an MSP Portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.msp_add_tenant_input import MspAddTenantInput
from cdo_sdk_python.models.status_info import StatusInfo
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    msp_add_tenant_input = cdo_sdk_python.MspAddTenantInput() # MspAddTenantInput | 

    try:
        # Add Tenant to MSP Portal (API token)
        api_response = api_instance.add_msp_tenant(msp_add_tenant_input)
        print("The response of MSPApi->add_msp_tenant:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPApi->add_msp_tenant: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msp_add_tenant_input** | [**MspAddTenantInput**](MspAddTenantInput.md)|  | 

### Return type

[**StatusInfo**](StatusInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Status |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_tenant_to_msp_portal**
> CdoTransaction add_tenant_to_msp_portal(tenant_uid)

Add tenant to MSP Portal

This is an asynchronous operation to add an existing tenant that the user making the API call is associated with to the MSP Portal. This endpoint will add an API-only user to the tenant, generate an API token for the API-only user, and add the tenant to the MSP Portal. Note: this endpoint can only be executed by a super-admin in the MSP Portal who is also associated with the tenant being added.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant to add to the MSP portal.

    try:
        # Add tenant to MSP Portal
        api_response = api_instance.add_tenant_to_msp_portal(tenant_uid)
        print("The response of MSPApi->add_tenant_to_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPApi->add_tenant_to_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant to add to the MSP portal. | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | SCC Transaction object that can be used to track the status of the operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**409** | Conflict. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_users_to_tenant_in_msp_portal**
> CdoTransaction add_users_to_tenant_in_msp_portal(tenant_uid, msp_add_users_to_tenant_input)

Add users to SCC tenant in MSP Portal

This is an asynchronous operation to add a list of users to a tenant associated with the MSP Portal.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
from cdo_sdk_python.models.msp_add_users_to_tenant_input import MspAddUsersToTenantInput
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant to which the users will be added.
    msp_add_users_to_tenant_input = cdo_sdk_python.MspAddUsersToTenantInput() # MspAddUsersToTenantInput | 

    try:
        # Add users to SCC tenant in MSP Portal
        api_response = api_instance.add_users_to_tenant_in_msp_portal(tenant_uid, msp_add_users_to_tenant_input)
        print("The response of MSPApi->add_users_to_tenant_in_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPApi->add_users_to_tenant_in_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant to which the users will be added. | 
 **msp_add_users_to_tenant_input** | [**MspAddUsersToTenantInput**](MspAddUsersToTenantInput.md)|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | SCC Transaction object that can be used to track the status of the operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_tenant**
> CdoTransaction create_tenant(msp_create_tenant_input)

Create SCC Tenant

Create a new tenant in SCC from an MSP Portal. This endpoint creates a tenant, adds the tenant to the MSP Portal. If the user creating the tenant is not an API-only user, the user is also added to the tenant. Note: This endpoint can only be executed by a super-admin in an MSP Portal. You can create no more than 1 tenant every 30 seconds. Additionally, you are limited to creating a total of 100 tenants: please speak to support to raise this limit.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
from cdo_sdk_python.models.msp_create_tenant_input import MspCreateTenantInput
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    msp_create_tenant_input = cdo_sdk_python.MspCreateTenantInput() # MspCreateTenantInput | 

    try:
        # Create SCC Tenant
        api_response = api_instance.create_tenant(msp_create_tenant_input)
        print("The response of MSPApi->create_tenant:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPApi->create_tenant: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msp_create_tenant_input** | [**MspCreateTenantInput**](MspCreateTenantInput.md)|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | SCC Transaction object that can be used to track the status of the tenant creation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**409** | Conflict. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_user_groups_from_tenant_in_msp_portal**
> CdoTransaction delete_user_groups_from_tenant_in_msp_portal(tenant_uid, msp_delete_user_groups_from_tenant_input)

Remove user groups from SCC tenant in MSP Portal

This is an asynchronous operation to remove a list of user groups from a tenant associated with the MSP Portal.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
from cdo_sdk_python.models.msp_delete_user_groups_from_tenant_input import MspDeleteUserGroupsFromTenantInput
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant from which the user groups will be deleted
    msp_delete_user_groups_from_tenant_input = cdo_sdk_python.MspDeleteUserGroupsFromTenantInput() # MspDeleteUserGroupsFromTenantInput | 

    try:
        # Remove user groups from SCC tenant in MSP Portal
        api_response = api_instance.delete_user_groups_from_tenant_in_msp_portal(tenant_uid, msp_delete_user_groups_from_tenant_input)
        print("The response of MSPApi->delete_user_groups_from_tenant_in_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPApi->delete_user_groups_from_tenant_in_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant from which the user groups will be deleted | 
 **msp_delete_user_groups_from_tenant_input** | [**MspDeleteUserGroupsFromTenantInput**](MspDeleteUserGroupsFromTenantInput.md)|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | SCC Transaction object that can be used to track the status of the operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_users_from_tenant_in_msp_portal**
> CdoTransaction delete_users_from_tenant_in_msp_portal(tenant_uid, msp_delete_users_from_tenant_input)

Remove users from SCC tenant in MSP Portal

This is an asynchronous operation to remove a list of users from a tenant associated with the MSP Portal.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
from cdo_sdk_python.models.msp_delete_users_from_tenant_input import MspDeleteUsersFromTenantInput
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant from which the users will be deleted
    msp_delete_users_from_tenant_input = cdo_sdk_python.MspDeleteUsersFromTenantInput() # MspDeleteUsersFromTenantInput | 

    try:
        # Remove users from SCC tenant in MSP Portal
        api_response = api_instance.delete_users_from_tenant_in_msp_portal(tenant_uid, msp_delete_users_from_tenant_input)
        print("The response of MSPApi->delete_users_from_tenant_in_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPApi->delete_users_from_tenant_in_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant from which the users will be deleted | 
 **msp_delete_users_from_tenant_input** | [**MspDeleteUsersFromTenantInput**](MspDeleteUsersFromTenantInput.md)|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | SCC Transaction object that can be used to track the status of the operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **enable_multicloud_defense_for_tenant_in_msp_portal**
> CdoTransaction enable_multicloud_defense_for_tenant_in_msp_portal(tenant_uid)

Enable Multicloud Defense for SCC tenant in MSP Portal

This is an asynchronous operation to enable Multicloud Defense for a tenant associated with the MSP Portal.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant that Multicloud Defense will be enabled for.

    try:
        # Enable Multicloud Defense for SCC tenant in MSP Portal
        api_response = api_instance.enable_multicloud_defense_for_tenant_in_msp_portal(tenant_uid)
        print("The response of MSPApi->enable_multicloud_defense_for_tenant_in_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPApi->enable_multicloud_defense_for_tenant_in_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant that Multicloud Defense will be enabled for. | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | SCC Transaction object that can be used to track the status of the operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **generate_api_token_for_user_in_tenant**
> ApiTokenInfo generate_api_token_for_user_in_tenant(tenant_uid, api_user_uid)

Generate token for API-only user on tenant managed by MSP portal

Generate API Token for API-only user on a tenant managed by the MSP portal. API-only users are used for building automations with SCC. If the user ID provided is not that of an API-only user, this operation will fail.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.api_token_info import ApiTokenInfo
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | The unique identifier of the tenant in SCC.
    api_user_uid = 'api_user_uid_example' # str | The unique identifier of the API-only user in SCC.

    try:
        # Generate token for API-only user on tenant managed by MSP portal
        api_response = api_instance.generate_api_token_for_user_in_tenant(tenant_uid, api_user_uid)
        print("The response of MSPApi->generate_api_token_for_user_in_tenant:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPApi->generate_api_token_for_user_in_tenant: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| The unique identifier of the tenant in SCC. | 
 **api_user_uid** | **str**| The unique identifier of the API-only user in SCC. | 

### Return type

[**ApiTokenInfo**](ApiTokenInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | API Token |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_tenant**
> MspManagedTenant get_msp_managed_tenant(tenant_uid)

Get SCC tenant managed by MSP Portal

Get a SCC tenant managed by the MSP Portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.msp_managed_tenant import MspManagedTenant
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | The unique identifier of the tenant in SCC.

    try:
        # Get SCC tenant managed by MSP Portal
        api_response = api_instance.get_msp_managed_tenant(tenant_uid)
        print("The response of MSPApi->get_msp_managed_tenant:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPApi->get_msp_managed_tenant: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| The unique identifier of the tenant in SCC. | 

### Return type

[**MspManagedTenant**](MspManagedTenant.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | SCC tenant object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_tenants**
> MspManagedTenantPage get_msp_managed_tenants(limit=limit, offset=offset, q=q)

Get SCC tenants managed by MSP Portal

Get a list of SCC tenants managed by the MSP Portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.msp_managed_tenant_page import MspManagedTenantPage
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | The offset of the results retrieved. The SCC API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Get SCC tenants managed by MSP Portal
        api_response = api_instance.get_msp_managed_tenants(limit=limit, offset=offset, q=q)
        print("The response of MSPApi->get_msp_managed_tenants:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPApi->get_msp_managed_tenants: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| The offset of the results retrieved. The SCC API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 

### Return type

[**MspManagedTenantPage**](MspManagedTenantPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of SCC tenant objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_users_in_tenant_in_msp_portal**
> UserPage get_users_in_tenant_in_msp_portal(tenant_uid, limit=limit, offset=offset, q=q)

Get Users associated with tenant in MSP portal

Get a list of users associated with the SCC tenant managed by MSP portal.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.user_page import UserPage
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant to retrieve the users for.
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | The offset of the results retrieved. The SCC API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Get Users associated with tenant in MSP portal
        api_response = api_instance.get_users_in_tenant_in_msp_portal(tenant_uid, limit=limit, offset=offset, q=q)
        print("The response of MSPApi->get_users_in_tenant_in_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPApi->get_users_in_tenant_in_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant to retrieve the users for. | 
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| The offset of the results retrieved. The SCC API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 

### Return type

[**UserPage**](UserPage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of User objects |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provision_cd_fmc_for_tenant_in_msp_portal**
> CdoTransaction provision_cd_fmc_for_tenant_in_msp_portal(tenant_uid)

Provision cdFMC for SCC tenant in MSP Portal

This is an asynchronous operation to provision a cdFMC for a tenant associated with the MSP Portal. This operation does not wait for the cdFMC to be provisioned on the target tenant, and will be marked as successful once the provisioning has been successfully triggered. To monitor the transaction on the target tenant, use the UID in the `TRANSACTION_UID_IN_TARGET_TENANT` key in the transaction details.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant that cdFMC provisioning will be enabled for.

    try:
        # Provision cdFMC for SCC tenant in MSP Portal
        api_response = api_instance.provision_cd_fmc_for_tenant_in_msp_portal(tenant_uid)
        print("The response of MSPApi->provision_cd_fmc_for_tenant_in_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPApi->provision_cd_fmc_for_tenant_in_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant that cdFMC provisioning will be enabled for. | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | SCC Transaction object that can be used to track the status of the operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **remove_tenant_from_msp_portal**
> remove_tenant_from_msp_portal(tenant_uid)

Remove tenant from MSP Portal

Removes a tenant currently associated with the MSP Portal. Note: this endpoint can only be executed by a super-admin in the MSP Portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | The unique identifier, represented as a UUID, of the tenant in SCC.

    try:
        # Remove tenant from MSP Portal
        api_instance.remove_tenant_from_msp_portal(tenant_uid)
    except Exception as e:
        print("Exception when calling MSPApi->remove_tenant_from_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| The unique identifier, represented as a UUID, of the tenant in SCC. | 

### Return type

void (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | No Content |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

