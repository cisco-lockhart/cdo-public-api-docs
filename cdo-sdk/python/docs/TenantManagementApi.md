# cdo_sdk_python.TenantManagementApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**add_msp_tenant**](TenantManagementApi.md#add_msp_tenant) | **POST** /v1/msp/tenants | Add Tenant to MSP Portal (API token)
[**add_tenant_to_msp_portal**](TenantManagementApi.md#add_tenant_to_msp_portal) | **POST** /v1/msp/tenants/{tenantUid} | Add tenant to MSP Portal
[**add_users_to_tenant_in_msp_portal**](TenantManagementApi.md#add_users_to_tenant_in_msp_portal) | **POST** /v1/msp/tenants/{tenantUid}/users | Add users to CDO tenant in MSP Portal
[**create_tenant**](TenantManagementApi.md#create_tenant) | **POST** /v1/msp/tenants/create | Create CDO Tenant
[**get_feature_flags**](TenantManagementApi.md#get_feature_flags) | **GET** /v1/features | Get Feature Flags
[**get_tenant**](TenantManagementApi.md#get_tenant) | **GET** /v1/tenants/{tenantUid} | Get Tenant
[**get_tenant_settings**](TenantManagementApi.md#get_tenant_settings) | **GET** /v1/settings/tenant | Get Tenant Settings
[**get_tenants**](TenantManagementApi.md#get_tenants) | **GET** /v1/tenants | Get Tenants
[**modify_tenant_settings**](TenantManagementApi.md#modify_tenant_settings) | **PATCH** /v1/settings/tenant | Modify Tenant Settings


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
    api_instance = cdo_sdk_python.TenantManagementApi(api_client)
    msp_add_tenant_input = cdo_sdk_python.MspAddTenantInput() # MspAddTenantInput | 

    try:
        # Add Tenant to MSP Portal (API token)
        api_response = api_instance.add_msp_tenant(msp_add_tenant_input)
        print("The response of TenantManagementApi->add_msp_tenant:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TenantManagementApi->add_msp_tenant: %s\n" % e)
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
    api_instance = cdo_sdk_python.TenantManagementApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant to add to the MSP portal.

    try:
        # Add tenant to MSP Portal
        api_response = api_instance.add_tenant_to_msp_portal(tenant_uid)
        print("The response of TenantManagementApi->add_tenant_to_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TenantManagementApi->add_tenant_to_msp_portal: %s\n" % e)
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
**422** | Unprocessable entity. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **add_users_to_tenant_in_msp_portal**
> CdoTransaction add_users_to_tenant_in_msp_portal(tenant_uid, msp_add_users_to_tenant_input)

Add users to CDO tenant in MSP Portal

This is an asynchronous operation to add a list of users to a tenant associated with the MSP Portal. Note: this endpoint cannot be used to add API-only users.

### Example

* Bearer (JWT) Authentication (bearerAuth):

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
    api_instance = cdo_sdk_python.TenantManagementApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant to which the users will be added.
    msp_add_users_to_tenant_input = cdo_sdk_python.MspAddUsersToTenantInput() # MspAddUsersToTenantInput | 

    try:
        # Add users to CDO tenant in MSP Portal
        api_response = api_instance.add_users_to_tenant_in_msp_portal(tenant_uid, msp_add_users_to_tenant_input)
        print("The response of TenantManagementApi->add_users_to_tenant_in_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TenantManagementApi->add_users_to_tenant_in_msp_portal: %s\n" % e)
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
**202** | CDO Transaction object that can be used to track the status of the operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
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
    api_instance = cdo_sdk_python.TenantManagementApi(api_client)
    msp_create_tenant_input = cdo_sdk_python.MspCreateTenantInput() # MspCreateTenantInput | 

    try:
        # Create CDO Tenant
        api_response = api_instance.create_tenant(msp_create_tenant_input)
        print("The response of TenantManagementApi->create_tenant:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TenantManagementApi->create_tenant: %s\n" % e)
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
**422** | Unprocessable entity. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_feature_flags**
> List[str] get_feature_flags()

Get Feature Flags

Get the feature flags enabled in the CDO tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
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
    api_instance = cdo_sdk_python.TenantManagementApi(api_client)

    try:
        # Get Feature Flags
        api_response = api_instance.get_feature_flags()
        print("The response of TenantManagementApi->get_feature_flags:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TenantManagementApi->get_feature_flags: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

**List[str]**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Set of feature flags that are enabled for the user&#39;s tenant. |  -  |
**401** | Request not authorized. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_tenant**
> Tenant get_tenant(tenant_uid)

Get Tenant

Fetch a tenant by UID in CDO. This will return a 200 response only if the user is associated with the tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.tenant import Tenant
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
    api_instance = cdo_sdk_python.TenantManagementApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | The unique identifier of the tenant in CDO.

    try:
        # Get Tenant
        api_response = api_instance.get_tenant(tenant_uid)
        print("The response of TenantManagementApi->get_tenant:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TenantManagementApi->get_tenant: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| The unique identifier of the tenant in CDO. | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Tenant Object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_tenant_settings**
> TenantSettings get_tenant_settings()

Get Tenant Settings

Get the settings for the CDO tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.tenant_settings import TenantSettings
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
    api_instance = cdo_sdk_python.TenantManagementApi(api_client)

    try:
        # Get Tenant Settings
        api_response = api_instance.get_tenant_settings()
        print("The response of TenantManagementApi->get_tenant_settings:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TenantManagementApi->get_tenant_settings: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**TenantSettings**](TenantSettings.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Tenant Settings Object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_tenants**
> TenantPage get_tenants(limit=limit, offset=offset)

Get Tenants

Get a list of tenants with which the CDO user is associated.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.tenant_page import TenantPage
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
    api_instance = cdo_sdk_python.TenantManagementApi(api_client)
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')

    try:
        # Get Tenants
        api_response = api_instance.get_tenants(limit=limit, offset=offset)
        print("The response of TenantManagementApi->get_tenants:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TenantManagementApi->get_tenants: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]

### Return type

[**TenantPage**](TenantPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Tenant Objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_tenant_settings**
> TenantSettings modify_tenant_settings(tenant_settings)

Modify Tenant Settings

Modify the settings for the CDO tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.tenant_settings import TenantSettings
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
    api_instance = cdo_sdk_python.TenantManagementApi(api_client)
    tenant_settings = cdo_sdk_python.TenantSettings() # TenantSettings | 

    try:
        # Modify Tenant Settings
        api_response = api_instance.modify_tenant_settings(tenant_settings)
        print("The response of TenantManagementApi->modify_tenant_settings:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling TenantManagementApi->modify_tenant_settings: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_settings** | [**TenantSettings**](TenantSettings.md)|  | 

### Return type

[**TenantSettings**](TenantSettings.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Tenant Settings Object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

