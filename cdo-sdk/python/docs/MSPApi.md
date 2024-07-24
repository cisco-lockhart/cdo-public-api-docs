# cdo_sdk_python.MSPApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_active_directory_groups_to_tenant_in_msp_portal**](MSPApi.md#add_active_directory_groups_to_tenant_in_msp_portal) | **POST** /v1/msp/tenants/{tenantUid}/users/groups | Add Active Directory Groups to CDO tenant in MSP Portal
[**add_msp_tenant**](MSPApi.md#add_msp_tenant) | **POST** /v1/msp/tenants | Add Tenant to MSP Portal (API token)
[**add_tenant_to_msp_portal**](MSPApi.md#add_tenant_to_msp_portal) | **POST** /v1/msp/tenants/{tenantUid} | Add tenant to MSP Portal
[**add_users_to_tenant_in_msp_portal**](MSPApi.md#add_users_to_tenant_in_msp_portal) | **POST** /v1/msp/tenants/{tenantUid}/users | Add users to CDO tenant in MSP Portal
[**create_tenant**](MSPApi.md#create_tenant) | **POST** /v1/msp/tenants/create | Create CDO Tenant


# **add_active_directory_groups_to_tenant_in_msp_portal**
> CdoTransaction add_active_directory_groups_to_tenant_in_msp_portal(tenant_uid, active_directory_group_create_or_update_input)

Add Active Directory Groups to CDO tenant in MSP Portal

This is an asynchronous operation to add a list of Active Directory Groups to a tenant associated with the MSP Portal.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.active_directory_group_create_or_update_input import ActiveDirectoryGroupCreateOrUpdateInput
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant to which the user groups will be added.
    active_directory_group_create_or_update_input = [cdo_sdk_python.ActiveDirectoryGroupCreateOrUpdateInput()] # List[ActiveDirectoryGroupCreateOrUpdateInput] | 

    try:
        # Add Active Directory Groups to CDO tenant in MSP Portal
        api_response = api_instance.add_active_directory_groups_to_tenant_in_msp_portal(tenant_uid, active_directory_group_create_or_update_input)
        print("The response of MSPApi->add_active_directory_groups_to_tenant_in_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPApi->add_active_directory_groups_to_tenant_in_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant to which the user groups will be added. | 
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
**202** | CDO Transaction object that can be used to track the status of the operation. |  -  |
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

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
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

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
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
**202** | CDO Transaction object that can be used to track the status of the operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**409** | Conflict. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_users_to_tenant_in_msp_portal**
> CdoTransaction add_users_to_tenant_in_msp_portal(tenant_uid, msp_add_users_to_tenant_input)

Add users to CDO tenant in MSP Portal

This is an asynchronous operation to add a list of users to a tenant associated with the MSP Portal. Note: this endpoint cannot be used to add API-only users.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
from cdo_sdk_python.models.msp_add_users_to_tenant_input import MspAddUsersToTenantInput
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MSPApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant to which the users will be added.
    msp_add_users_to_tenant_input = cdo_sdk_python.MspAddUsersToTenantInput() # MspAddUsersToTenantInput | 

    try:
        # Add users to CDO tenant in MSP Portal
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
**202** | CDO Transaction object that can be used to track the status of the operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_tenant**
> CdoTransaction create_tenant(msp_create_tenant_input)

Create CDO Tenant

Create a new tenant in CDO from an MSP Portal. This endpoint creates a tenant, adds the tenant to the MSP Portal. If the user creating the tenant is not an API-only user, the user is also added to the tenant. Note: This endpoint can only be executed by a super-admin in an MSP Portal. You can create no more than 1 tenant every 30 seconds. Additionally, you are limited to creating a total of 100 tenants: please speak to support to raise this limit.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
from cdo_sdk_python.models.msp_create_tenant_input import MspCreateTenantInput
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
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
        # Create CDO Tenant
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
**202** | CDO Transaction object that can be used to track the status of the tenant creation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**409** | Conflict. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

