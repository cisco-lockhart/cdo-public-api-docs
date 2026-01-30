# scc_firewall_manager_sdk.MSPTenantManagementApi

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
[**add_msp_tenant**](MSPTenantManagementApi.md#add_msp_tenant) | **POST** /v1/msp/tenants | Add Tenant to MSP Portal (API token)
[**add_tenant_to_msp_portal**](MSPTenantManagementApi.md#add_tenant_to_msp_portal) | **POST** /v1/msp/tenants/{tenantUid} | Add tenant to MSP Portal
[**create_tenant**](MSPTenantManagementApi.md#create_tenant) | **POST** /v1/msp/tenants/create | Create Security Cloud Control Tenant
[**enable_multicloud_defense_for_tenant_in_msp_portal**](MSPTenantManagementApi.md#enable_multicloud_defense_for_tenant_in_msp_portal) | **POST** /v1/msp/tenants/{tenantUid}/mcd | Enable Multicloud Defense for Security Cloud Control tenant in MSP Portal
[**generate_api_token_for_user_in_tenant**](MSPTenantManagementApi.md#generate_api_token_for_user_in_tenant) | **POST** /v1/msp/tenants/{tenantUid}/users/{apiUserUid}/token | Generate token for API-only user on tenant managed by MSP portal
[**get_msp_managed_tenant**](MSPTenantManagementApi.md#get_msp_managed_tenant) | **GET** /v1/msp/tenants/{tenantUid} | Get Security Cloud Control tenant managed by MSP Portal
[**get_msp_managed_tenants**](MSPTenantManagementApi.md#get_msp_managed_tenants) | **GET** /v1/msp/tenants | Get Security Cloud Control tenants managed by MSP Portal
[**provision_cd_fmc_for_tenant_in_msp_portal**](MSPTenantManagementApi.md#provision_cd_fmc_for_tenant_in_msp_portal) | **POST** /v1/msp/tenants/{tenantUid}/cdfmc | Provision cdFMC for Security Cloud Control tenant in MSP Portal
[**remove_tenant_from_msp_portal**](MSPTenantManagementApi.md#remove_tenant_from_msp_portal) | **DELETE** /v1/msp/tenants/{tenantUid} | Remove tenant from MSP Portal


# **add_msp_tenant**
> StatusInfo add_msp_tenant(msp_add_tenant_input)

Add Tenant to MSP Portal (API token)

Add a tenant to the MSP Portal by providing an API token for the tenant. Use this endpoint to add a tenant that the user performing the operation is not associated with, or a tenant in a different region to that of the MSP Portal. Note: This endpoint can only be executed by a super-admin in an MSP Portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_add_tenant_input import MspAddTenantInput
from scc_firewall_manager_sdk.models.status_info import StatusInfo
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
    api_instance = scc_firewall_manager_sdk.MSPTenantManagementApi(api_client)
    msp_add_tenant_input = scc_firewall_manager_sdk.MspAddTenantInput() # MspAddTenantInput | 

    try:
        # Add Tenant to MSP Portal (API token)
        api_response = api_instance.add_msp_tenant(msp_add_tenant_input)
        print("The response of MSPTenantManagementApi->add_msp_tenant:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPTenantManagementApi->add_msp_tenant: %s\n" % e)
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
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
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
    api_instance = scc_firewall_manager_sdk.MSPTenantManagementApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant to add to the MSP portal.

    try:
        # Add tenant to MSP Portal
        api_response = api_instance.add_tenant_to_msp_portal(tenant_uid)
        print("The response of MSPTenantManagementApi->add_tenant_to_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPTenantManagementApi->add_tenant_to_msp_portal: %s\n" % e)
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
**202** | Security Cloud Control Transaction object that can be used to track the status of the operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**409** | Conflict. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_tenant**
> CdoTransaction create_tenant(msp_create_tenant_input)

Create Security Cloud Control Tenant

Create a new tenant in Security Cloud Control from an MSP Portal. This endpoint creates a tenant, adds the tenant to the MSP Portal. If the user creating the tenant is not an API-only user, the user is also added to the tenant. Note: This endpoint can only be executed by a super-admin in an MSP Portal. You can create no more than 1 tenant every 30 seconds. Additionally, you are limited to creating a total of 100 tenants: please speak to support to raise this limit.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.msp_create_tenant_input import MspCreateTenantInput
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
    api_instance = scc_firewall_manager_sdk.MSPTenantManagementApi(api_client)
    msp_create_tenant_input = scc_firewall_manager_sdk.MspCreateTenantInput() # MspCreateTenantInput | 

    try:
        # Create Security Cloud Control Tenant
        api_response = api_instance.create_tenant(msp_create_tenant_input)
        print("The response of MSPTenantManagementApi->create_tenant:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPTenantManagementApi->create_tenant: %s\n" % e)
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
**202** | Security Cloud Control Transaction object that can be used to track the status of the tenant creation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**409** | Conflict. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **enable_multicloud_defense_for_tenant_in_msp_portal**
> CdoTransaction enable_multicloud_defense_for_tenant_in_msp_portal(tenant_uid)

Enable Multicloud Defense for Security Cloud Control tenant in MSP Portal

This is an asynchronous operation to enable Multicloud Defense for a tenant associated with the MSP Portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
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
    api_instance = scc_firewall_manager_sdk.MSPTenantManagementApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant that Multicloud Defense will be enabled for.

    try:
        # Enable Multicloud Defense for Security Cloud Control tenant in MSP Portal
        api_response = api_instance.enable_multicloud_defense_for_tenant_in_msp_portal(tenant_uid)
        print("The response of MSPTenantManagementApi->enable_multicloud_defense_for_tenant_in_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPTenantManagementApi->enable_multicloud_defense_for_tenant_in_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant that Multicloud Defense will be enabled for. | 

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
**202** | Security Cloud Control Transaction object that can be used to track the status of the operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **generate_api_token_for_user_in_tenant**
> ApiTokenInfo generate_api_token_for_user_in_tenant(tenant_uid, api_user_uid)

Generate token for API-only user on tenant managed by MSP portal

Generate API Token for API-only user on a tenant managed by the MSP portal. API-only users are used for building automations with Security Cloud Control. If the user ID provided is not that of an API-only user, this operation will fail.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.api_token_info import ApiTokenInfo
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
    api_instance = scc_firewall_manager_sdk.MSPTenantManagementApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | The unique identifier of the tenant in Security Cloud Control.
    api_user_uid = 'api_user_uid_example' # str | The unique identifier of the API-only user in Security Cloud Control.

    try:
        # Generate token for API-only user on tenant managed by MSP portal
        api_response = api_instance.generate_api_token_for_user_in_tenant(tenant_uid, api_user_uid)
        print("The response of MSPTenantManagementApi->generate_api_token_for_user_in_tenant:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPTenantManagementApi->generate_api_token_for_user_in_tenant: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| The unique identifier of the tenant in Security Cloud Control. | 
 **api_user_uid** | **str**| The unique identifier of the API-only user in Security Cloud Control. | 

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

Get Security Cloud Control tenant managed by MSP Portal

Get a Security Cloud Control tenant managed by the MSP Portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_managed_tenant import MspManagedTenant
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
    api_instance = scc_firewall_manager_sdk.MSPTenantManagementApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | The unique identifier of the tenant in Security Cloud Control.

    try:
        # Get Security Cloud Control tenant managed by MSP Portal
        api_response = api_instance.get_msp_managed_tenant(tenant_uid)
        print("The response of MSPTenantManagementApi->get_msp_managed_tenant:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPTenantManagementApi->get_msp_managed_tenant: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| The unique identifier of the tenant in Security Cloud Control. | 

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
**200** | Security Cloud Control tenant object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_tenants**
> MspManagedTenantPage get_msp_managed_tenants(limit=limit, offset=offset, q=q)

Get Security Cloud Control tenants managed by MSP Portal

Get a list of Security Cloud Control tenants managed by the MSP Portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_managed_tenant_page import MspManagedTenantPage
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
    api_instance = scc_firewall_manager_sdk.MSPTenantManagementApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Get Security Cloud Control tenants managed by MSP Portal
        api_response = api_instance.get_msp_managed_tenants(limit=limit, offset=offset, q=q)
        print("The response of MSPTenantManagementApi->get_msp_managed_tenants:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPTenantManagementApi->get_msp_managed_tenants: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
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
**200** | List of Security Cloud Control tenant objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provision_cd_fmc_for_tenant_in_msp_portal**
> CdoTransaction provision_cd_fmc_for_tenant_in_msp_portal(tenant_uid, enable_cd_fmc_for_tenant_request=enable_cd_fmc_for_tenant_request)

Provision cdFMC for Security Cloud Control tenant in MSP Portal

This is an asynchronous operation to provision a cdFMC for a tenant associated with the MSP Portal. This operation does not wait for the cdFMC to be provisioned on the target tenant, and will be marked as successful once the provisioning has been successfully triggered. To monitor the transaction on the target tenant, use the UID in the `TRANSACTION_UID_IN_TARGET_TENANT` key in the transaction details.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.enable_cd_fmc_for_tenant_request import EnableCdFmcForTenantRequest
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
    api_instance = scc_firewall_manager_sdk.MSPTenantManagementApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | Unique identifier of the tenant that cdFMC provisioning will be enabled for.
    enable_cd_fmc_for_tenant_request = scc_firewall_manager_sdk.EnableCdFmcForTenantRequest() # EnableCdFmcForTenantRequest |  (optional)

    try:
        # Provision cdFMC for Security Cloud Control tenant in MSP Portal
        api_response = api_instance.provision_cd_fmc_for_tenant_in_msp_portal(tenant_uid, enable_cd_fmc_for_tenant_request=enable_cd_fmc_for_tenant_request)
        print("The response of MSPTenantManagementApi->provision_cd_fmc_for_tenant_in_msp_portal:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPTenantManagementApi->provision_cd_fmc_for_tenant_in_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| Unique identifier of the tenant that cdFMC provisioning will be enabled for. | 
 **enable_cd_fmc_for_tenant_request** | [**EnableCdFmcForTenantRequest**](EnableCdFmcForTenantRequest.md)|  | [optional] 

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

# **remove_tenant_from_msp_portal**
> remove_tenant_from_msp_portal(tenant_uid)

Remove tenant from MSP Portal

Removes a tenant currently associated with the MSP Portal. Note: this endpoint can only be executed by a super-admin in the MSP Portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
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
    api_instance = scc_firewall_manager_sdk.MSPTenantManagementApi(api_client)
    tenant_uid = 'tenant_uid_example' # str | The unique identifier, represented as a UUID, of the tenant in Security Cloud Control.

    try:
        # Remove tenant from MSP Portal
        api_instance.remove_tenant_from_msp_portal(tenant_uid)
    except Exception as e:
        print("Exception when calling MSPTenantManagementApi->remove_tenant_from_msp_portal: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenant_uid** | **str**| The unique identifier, represented as a UUID, of the tenant in Security Cloud Control. | 

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

