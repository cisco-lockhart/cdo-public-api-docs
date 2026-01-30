# scc_firewall_manager_sdk.MSPUserManagementApi

All URIs are relative to the base URL, which depends on the region your organization is deployed to.

Region | Base URL
------------- | -------------
US  | https://api.us.security.cisco.com/firewall
EU  | https://api.eu.security.cisco.com/firewall
APJ | https://api.apj.security.cisco.com/firewall
AU  | https://api.au.security.cisco.com/firewall
IN  | https://api.in.security.cisco.com/firewall
UAE | https://api.uae.security.cisco.com/firewall

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_users_to_tenant_in_msp_portal**](MSPUserManagementApi.md#add_users_to_tenant_in_msp_portal) | **POST** /v1/msp/tenants/{tenantUid}/users | Add users to Security Cloud Control tenant in MSP Portal
[**delete_users_from_tenant_in_msp_portal**](MSPUserManagementApi.md#delete_users_from_tenant_in_msp_portal) | **POST** /v1/msp/tenants/{tenantUid}/users/delete | Remove users from Security Cloud Control tenant in MSP Portal
[**get_api_only_users_in_msp_managed_tenant**](MSPUserManagementApi.md#get_api_only_users_in_msp_managed_tenant) | **GET** /v1/msp/tenants/{tenantUid}/users/api-only | Get API-only Users associated with tenant in MSP portal
[**get_human_users_in_msp_managed_tenant**](MSPUserManagementApi.md#get_human_users_in_msp_managed_tenant) | **GET** /v1/msp/tenants/{tenantUid}/users | Get non-API-only Users associated with tenant in MSP portal


# **add_users_to_tenant_in_msp_portal**
> CdoTransaction add_users_to_tenant_in_msp_portal(tenant_uid, msp_add_users_to_tenant_input)

Add users to Security Cloud Control tenant in MSP Portal

This is an asynchronous operation to add a list of users to a tenant associated with the MSP Portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.msp_add_users_to_tenant_input import MspAddUsersToTenantInput
from scc_firewall_manager_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.us.security.cisco.com/firewall
# See configuration.py for a list of all supported configuration parameters.
configuration = scc_firewall_manager_sdk.Configuration(
    host = "https://api.us.security.cisco.com/firewall"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = scc_firewall_manager_sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with scc_firewall_manager_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = scc_firewall_manager_sdk.MSPUserManagementApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant to which the users will be added.
    msp_add_users_to_tenant_input = scc_firewall_manager_sdk.MspAddUsersToTenantInput() # MspAddUsersToTenantInput | 

    try:
        # Add users to Security Cloud Control tenant in MSP Portal
        api_response = api_instance.add_users_to_tenant_in_msp_portal(tenant_uid, msp_add_users_to_tenant_input)
        print("The response of MSPUserManagementApi->add_users_to_tenant_in_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPUserManagementApi->add_users_to_tenant_in_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant to which the users will be added. | 
 **msp_add_users_to_tenant_input** | [**MspAddUsersToTenantInput**](MspAddUsersToTenantInput.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the status of the operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_users_from_tenant_in_msp_portal**
> CdoTransaction delete_users_from_tenant_in_msp_portal(tenant_uid, msp_delete_users_from_tenant_input)

Remove users from Security Cloud Control tenant in MSP Portal

This is an asynchronous operation to remove a list of users from a tenant associated with the MSP Portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.msp_delete_users_from_tenant_input import MspDeleteUsersFromTenantInput
from scc_firewall_manager_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.us.security.cisco.com/firewall
# See configuration.py for a list of all supported configuration parameters.
configuration = scc_firewall_manager_sdk.Configuration(
    host = "https://api.us.security.cisco.com/firewall"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = scc_firewall_manager_sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with scc_firewall_manager_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = scc_firewall_manager_sdk.MSPUserManagementApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant from which the users will be deleted
    msp_delete_users_from_tenant_input = scc_firewall_manager_sdk.MspDeleteUsersFromTenantInput() # MspDeleteUsersFromTenantInput | 

    try:
        # Remove users from Security Cloud Control tenant in MSP Portal
        api_response = api_instance.delete_users_from_tenant_in_msp_portal(tenant_uid, msp_delete_users_from_tenant_input)
        print("The response of MSPUserManagementApi->delete_users_from_tenant_in_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPUserManagementApi->delete_users_from_tenant_in_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant from which the users will be deleted | 
 **msp_delete_users_from_tenant_input** | [**MspDeleteUsersFromTenantInput**](MspDeleteUsersFromTenantInput.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the status of the operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_api_only_users_in_msp_managed_tenant**
> UserPage get_api_only_users_in_msp_managed_tenant(tenant_uid, limit=limit, offset=offset, q=q)

Get API-only Users associated with tenant in MSP portal

Get a list of API-only users associated with the Security Cloud Control tenant managed by MSP portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.user_page import UserPage
from scc_firewall_manager_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.us.security.cisco.com/firewall
# See configuration.py for a list of all supported configuration parameters.
configuration = scc_firewall_manager_sdk.Configuration(
    host = "https://api.us.security.cisco.com/firewall"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = scc_firewall_manager_sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with scc_firewall_manager_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = scc_firewall_manager_sdk.MSPUserManagementApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant to retrieve the users for.
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Get API-only Users associated with tenant in MSP portal
        api_response = api_instance.get_api_only_users_in_msp_managed_tenant(tenant_uid, limit=limit, offset=offset, q=q)
        print("The response of MSPUserManagementApi->get_api_only_users_in_msp_managed_tenant:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPUserManagementApi->get_api_only_users_in_msp_managed_tenant: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant to retrieve the users for. | 
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 

### Return type

[**UserPage**](UserPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of User objects |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_human_users_in_msp_managed_tenant**
> UserPage get_human_users_in_msp_managed_tenant(tenant_uid, limit=limit, offset=offset, q=q)

Get non-API-only Users associated with tenant in MSP portal

Get a list of (non-API-only) users associated with the Security Cloud Control tenant managed by MSP portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.user_page import UserPage
from scc_firewall_manager_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.us.security.cisco.com/firewall
# See configuration.py for a list of all supported configuration parameters.
configuration = scc_firewall_manager_sdk.Configuration(
    host = "https://api.us.security.cisco.com/firewall"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = scc_firewall_manager_sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with scc_firewall_manager_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = scc_firewall_manager_sdk.MSPUserManagementApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant to retrieve the users for.
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Get non-API-only Users associated with tenant in MSP portal
        api_response = api_instance.get_human_users_in_msp_managed_tenant(tenant_uid, limit=limit, offset=offset, q=q)
        print("The response of MSPUserManagementApi->get_human_users_in_msp_managed_tenant:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPUserManagementApi->get_human_users_in_msp_managed_tenant: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant to retrieve the users for. | 
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 

### Return type

[**UserPage**](UserPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of User objects |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

