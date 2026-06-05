# scc_firewall_manager_sdk.UsersApi

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
[**create_active_directory_group**](UsersApi.md#create_active_directory_group) | **POST** /v1/users/groups | Create Active Directory Group in Security Cloud Control Tenant
[**create_multiple_users**](UsersApi.md#create_multiple_users) | **POST** /v1/users/bulk | Create multiple users in Security Cloud Control Tenant
[**create_user**](UsersApi.md#create_user) | **POST** /v1/users | Create User in Security Cloud Control Tenant
[**delete_active_directory_group**](UsersApi.md#delete_active_directory_group) | **DELETE** /v1/users/groups/{groupUid} | Delete Active Directory Group from Security Cloud Control Tenant
[**delete_multiple_users**](UsersApi.md#delete_multiple_users) | **POST** /v1/users/delete | Delete Multiple Users from Security Cloud Control Firewall Manager Tenant
[**delete_user**](UsersApi.md#delete_user) | **DELETE** /v1/users/{userUid} | Delete User from Security Cloud Control Firewall Manager Tenant
[**generate_api_token**](UsersApi.md#generate_api_token) | **POST** /v1/users/{apiUserUid}/apiToken/generate | Generate Token for API-only user
[**get_active_directory_group**](UsersApi.md#get_active_directory_group) | **GET** /v1/users/groups/{groupUid} | Get Active Directory Group
[**get_active_directory_groups**](UsersApi.md#get_active_directory_groups) | **GET** /v1/users/groups | Get Active Directory Groups
[**get_api_only_users**](UsersApi.md#get_api_only_users) | **GET** /v1/users/api-only | Get Tenant API-only Users
[**get_token**](UsersApi.md#get_token) | **GET** /v1/token | Get Token Info
[**get_user**](UsersApi.md#get_user) | **GET** /v1/users/{userUid} | Get Tenant User
[**get_users**](UsersApi.md#get_users) | **GET** /v1/users | Get Tenant Users
[**modify_active_directory_group**](UsersApi.md#modify_active_directory_group) | **PATCH** /v1/users/groups/{groupUid} | Modify Active Directory Group
[**revoke_api_token**](UsersApi.md#revoke_api_token) | **POST** /v1/users/{apiUserUid}/apiToken/revoke | Revoke API-only User&#39;s Token
[**revoke_token**](UsersApi.md#revoke_token) | **POST** /v1/token/revoke | Revoke Token


# **create_active_directory_group**
> ActiveDirectoryGroup create_active_directory_group(active_directory_group_create_or_update_input)

Create Active Directory Group in Security Cloud Control Tenant

Create an Active Directory Group in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.active_directory_group import ActiveDirectoryGroup
from scc_firewall_manager_sdk.models.active_directory_group_create_or_update_input import ActiveDirectoryGroupCreateOrUpdateInput
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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)
    active_directory_group_create_or_update_input = scc_firewall_manager_sdk.ActiveDirectoryGroupCreateOrUpdateInput() # ActiveDirectoryGroupCreateOrUpdateInput | 

    try:
        # Create Active Directory Group in Security Cloud Control Tenant
        api_response = api_instance.create_active_directory_group(active_directory_group_create_or_update_input)
        print("The response of UsersApi->create_active_directory_group:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->create_active_directory_group: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active_directory_group_create_or_update_input** | [**ActiveDirectoryGroupCreateOrUpdateInput**](ActiveDirectoryGroupCreateOrUpdateInput.md)|  | 

### Return type

[**ActiveDirectoryGroup**](ActiveDirectoryGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Active Directory Group object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_multiple_users**
> CdoTransaction create_multiple_users(user_create_or_update_input, x_calling_service=x_calling_service, x_request_metadata=x_request_metadata)

Create multiple users in Security Cloud Control Tenant

This is an asynchronous operation to create multiple users in the Security Cloud Control enterprise and assign the specified role to each user for access to the Firewall Manager application.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.user_create_or_update_input import UserCreateOrUpdateInput
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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)
    user_create_or_update_input = [scc_firewall_manager_sdk.UserCreateOrUpdateInput()] # List[UserCreateOrUpdateInput] | 
    x_calling_service = 'x_calling_service_example' # str |  (optional)
    x_request_metadata = 'x_request_metadata_example' # str | Custom metadata to identify the request (optional)

    try:
        # Create multiple users in Security Cloud Control Tenant
        api_response = api_instance.create_multiple_users(user_create_or_update_input, x_calling_service=x_calling_service, x_request_metadata=x_request_metadata)
        print("The response of UsersApi->create_multiple_users:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->create_multiple_users: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_create_or_update_input** | [**List[UserCreateOrUpdateInput]**](UserCreateOrUpdateInput.md)|  | 
 **x_calling_service** | **str**|  | [optional] 
 **x_request_metadata** | **str**| Custom metadata to identify the request | [optional] 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the creation operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_user**
> User create_user(user_create_or_update_input)

Create User in Security Cloud Control Tenant

Create a user in the Security Cloud Control enterprise and assigns the specified role for access to the Firewall Manager application.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.user import User
from scc_firewall_manager_sdk.models.user_create_or_update_input import UserCreateOrUpdateInput
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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)
    user_create_or_update_input = scc_firewall_manager_sdk.UserCreateOrUpdateInput() # UserCreateOrUpdateInput | 

    try:
        # Create User in Security Cloud Control Tenant
        api_response = api_instance.create_user(user_create_or_update_input)
        print("The response of UsersApi->create_user:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->create_user: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_create_or_update_input** | [**UserCreateOrUpdateInput**](UserCreateOrUpdateInput.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | User object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_active_directory_group**
> delete_active_directory_group(group_uid)

Delete Active Directory Group from Security Cloud Control Tenant

Delete a Active Directory Group by UID in the Security Cloud Control tenant.

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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)
    group_uid = 'group_uid_example' # str | The unique identifier, represented as a UUID, of the Active Directory Group in Security Cloud Control.

    try:
        # Delete Active Directory Group from Security Cloud Control Tenant
        api_instance.delete_active_directory_group(group_uid)
    except Exception as e:
        print("Exception when calling UsersApi->delete_active_directory_group: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_uid** | **str**| The unique identifier, represented as a UUID, of the Active Directory Group in Security Cloud Control. | 

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

# **delete_multiple_users**
> CdoTransaction delete_multiple_users(request_body, x_calling_service=x_calling_service, x_request_metadata=x_request_metadata)

Delete Multiple Users from Security Cloud Control Firewall Manager Tenant

This is an asynchronous operation that revokes a user's access to Firewall Manager within a specific Security Cloud Control enterprise. This action will not affect the user's access to other products in the enterprise or their access to Firewall Manager in other enterprises.

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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)
    request_body = ['request_body_example'] # List[str] | 
    x_calling_service = 'x_calling_service_example' # str |  (optional)
    x_request_metadata = 'x_request_metadata_example' # str | Custom metadata to identify the request (optional)

    try:
        # Delete Multiple Users from Security Cloud Control Firewall Manager Tenant
        api_response = api_instance.delete_multiple_users(request_body, x_calling_service=x_calling_service, x_request_metadata=x_request_metadata)
        print("The response of UsersApi->delete_multiple_users:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->delete_multiple_users: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request_body** | [**List[str]**](str.md)|  | 
 **x_calling_service** | **str**|  | [optional] 
 **x_request_metadata** | **str**| Custom metadata to identify the request | [optional] 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the deletion operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_user**
> delete_user(user_uid)

Delete User from Security Cloud Control Firewall Manager Tenant

Revokes a user's access to Firewall Manager within a specific Security Cloud Control enterprise. This action will not affect the user's access to other products in the enterprise or their access to Firewall Manager in other enterprises.

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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)
    user_uid = 'user_uid_example' # str | The unique identifier of the user in Security Cloud Control.

    try:
        # Delete User from Security Cloud Control Firewall Manager Tenant
        api_instance.delete_user(user_uid)
    except Exception as e:
        print("Exception when calling UsersApi->delete_user: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_uid** | **str**| The unique identifier of the user in Security Cloud Control. | 

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

# **generate_api_token**
> ApiTokenInfo generate_api_token(api_user_uid)

Generate Token for API-only user

Generate API Token for API-only user. API-only users are used for building automations with Security Cloud Control. If the user ID provided is not that of an API-only user, this operation will fail.

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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)
    api_user_uid = 'api_user_uid_example' # str | The unique identifier, represented as a UUID, of the API-only user in Security Cloud Control

    try:
        # Generate Token for API-only user
        api_response = api_instance.generate_api_token(api_user_uid)
        print("The response of UsersApi->generate_api_token:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->generate_api_token: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **api_user_uid** | **str**| The unique identifier, represented as a UUID, of the API-only user in Security Cloud Control | 

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

# **get_active_directory_group**
> ActiveDirectoryGroup get_active_directory_group(group_uid)

Get Active Directory Group

Fetch a active directory group by UID in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.active_directory_group import ActiveDirectoryGroup
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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)
    group_uid = 'group_uid_example' # str | The unique identifier, represented as a UUID, of the active directory group in Security Cloud Control.

    try:
        # Get Active Directory Group
        api_response = api_instance.get_active_directory_group(group_uid)
        print("The response of UsersApi->get_active_directory_group:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->get_active_directory_group: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_uid** | **str**| The unique identifier, represented as a UUID, of the active directory group in Security Cloud Control. | 

### Return type

[**ActiveDirectoryGroup**](ActiveDirectoryGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Active Directory Group object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_active_directory_groups**
> ActiveDirectoryGroupPage get_active_directory_groups(limit=limit, offset=offset, q=q)

Get Active Directory Groups

Get a list of active directory groups associated with the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.active_directory_group_page import ActiveDirectoryGroupPage
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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Get Active Directory Groups
        api_response = api_instance.get_active_directory_groups(limit=limit, offset=offset, q=q)
        print("The response of UsersApi->get_active_directory_groups:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->get_active_directory_groups: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 

### Return type

[**ActiveDirectoryGroupPage**](ActiveDirectoryGroupPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Active Directory Group objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_api_only_users**
> UserPage get_api_only_users(limit=limit, offset=offset, q=q)

Get Tenant API-only Users

Get a list of API-only users associated with the Security Cloud Control Firewall Manager tenant.

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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Get Tenant API-only Users
        api_response = api_instance.get_api_only_users(limit=limit, offset=offset, q=q)
        print("The response of UsersApi->get_api_only_users:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->get_api_only_users: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_token**
> CdoTokenInfo get_token()

Get Token Info

Fetch information on the current token. Each Security Cloud Control token is associated with a specific user and a specific tenant. A token can only be used to perform operations on the tenant it is associated with.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_token_info import CdoTokenInfo
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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)

    try:
        # Get Token Info
        api_response = api_instance.get_token()
        print("The response of UsersApi->get_token:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->get_token: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**CdoTokenInfo**](CdoTokenInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Token information |  -  |
**401** | Request not authorized. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_user**
> User get_user(user_uid)

Get Tenant User

Fetch a user by UID in the CDO tenant. This will only return 200 responses for users associated with the Security Cloud Control tenant and assigned to the SCC Firewall Manager.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.user import User
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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)
    user_uid = 'user_uid_example' # str | The unique identifier of the user in Security Cloud Control.

    try:
        # Get Tenant User
        api_response = api_instance.get_user(user_uid)
        print("The response of UsersApi->get_user:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->get_user: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_uid** | **str**| The unique identifier of the user in Security Cloud Control. | 

### Return type

[**User**](User.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | User object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_users**
> UserPage get_users(limit=limit, offset=offset, q=q)

Get Tenant Users

Get a list of (non-API-only) users associated with the Security Cloud Control tenant assigned to the SCC Firewall Manager. Note: This endpoint returns only human users, not API-only users.

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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)
    limit = 'limit_example' # str | The number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Get Tenant Users
        api_response = api_instance.get_users(limit=limit, offset=offset, q=q)
        print("The response of UsersApi->get_users:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->get_users: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| The number of results to retrieve. | [optional] 
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
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_active_directory_group**
> ActiveDirectoryGroup modify_active_directory_group(group_uid, active_directory_group_create_or_update_input)

Modify Active Directory Group

Modify an Active Directory Group by UID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.active_directory_group import ActiveDirectoryGroup
from scc_firewall_manager_sdk.models.active_directory_group_create_or_update_input import ActiveDirectoryGroupCreateOrUpdateInput
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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)
    group_uid = 'group_uid_example' # str | The unique identifier, represented as a UUID, of the active directory group in Security Cloud Control.
    active_directory_group_create_or_update_input = scc_firewall_manager_sdk.ActiveDirectoryGroupCreateOrUpdateInput() # ActiveDirectoryGroupCreateOrUpdateInput | 

    try:
        # Modify Active Directory Group
        api_response = api_instance.modify_active_directory_group(group_uid, active_directory_group_create_or_update_input)
        print("The response of UsersApi->modify_active_directory_group:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->modify_active_directory_group: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group_uid** | **str**| The unique identifier, represented as a UUID, of the active directory group in Security Cloud Control. | 
 **active_directory_group_create_or_update_input** | [**ActiveDirectoryGroupCreateOrUpdateInput**](ActiveDirectoryGroupCreateOrUpdateInput.md)|  | 

### Return type

[**ActiveDirectoryGroup**](ActiveDirectoryGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Active Directory Group object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **revoke_api_token**
> StatusInfo revoke_api_token(api_user_uid)

Revoke API-only User's Token

Revoke API Token of API-only user. If the user ID provided is not that of an API-only user, this operation will fail.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)
    api_user_uid = 'api_user_uid_example' # str | The unique identifier, represented as a UUID, of the API user in Security Cloud Control.

    try:
        # Revoke API-only User's Token
        api_response = api_instance.revoke_api_token(api_user_uid)
        print("The response of UsersApi->revoke_api_token:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->revoke_api_token: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **api_user_uid** | **str**| The unique identifier, represented as a UUID, of the API user in Security Cloud Control. | 

### Return type

[**StatusInfo**](StatusInfo.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
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

# **revoke_token**
> revoke_token()

Revoke Token

Revoke the current token. All subsequent requests with a revoked token will fail with 401 Unauthorized errors.

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
    api_instance = scc_firewall_manager_sdk.UsersApi(api_client)

    try:
        # Revoke Token
        api_instance.revoke_token()
    except Exception as e:
        print("Exception when calling UsersApi->revoke_token: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

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
**200** | Revoke success |  -  |
**401** | Request not authorized. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

