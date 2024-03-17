# cdo-python-sdk.UsersApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_user**](UsersApi.md#create_user) | **POST** /v1/users | Create a user in the CDO tenant
[**delete_user**](UsersApi.md#delete_user) | **DELETE** /v1/users/{userUid} | Delete a User by UID in the CDO tenant
[**generate_api_token**](UsersApi.md#generate_api_token) | **POST** /v1/users/{apiUserId}/apiToken/generate | Generate API Token for API-only user
[**get_token**](UsersApi.md#get_token) | **GET** /v1/token | Fetch information on the current token
[**get_user**](UsersApi.md#get_user) | **GET** /v1/users/{userUid} | Fetch a user by UID in the CDO tenant.
[**list_users**](UsersApi.md#list_users) | **GET** /v1/users | Fetch a list of users associated with the CDO tenant.
[**revoke_api_token**](UsersApi.md#revoke_api_token) | **POST** /v1/users/{apiUserId}/apiToken/revoke | Revoke API Token of API-only user
[**revoke_token**](UsersApi.md#revoke_token) | **POST** /v1/token/revoke | Revoke the current token


# **create_user**
> User create_user(user_create_or_update_input)

Create a user in the CDO tenant

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo-python-sdk
from cdo-python-sdk.models.user import User
from cdo-python-sdk.models.user_create_or_update_input import UserCreateOrUpdateInput
from cdo-python-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo-python-sdk.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo-python-sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo-python-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo-python-sdk.UsersApi(api_client)
    user_create_or_update_input = cdo-python-sdk.UserCreateOrUpdateInput() # UserCreateOrUpdateInput | 

    try:
        # Create a user in the CDO tenant
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

# **delete_user**
> delete_user(user_uid)

Delete a User by UID in the CDO tenant

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo-python-sdk
from cdo-python-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo-python-sdk.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo-python-sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo-python-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo-python-sdk.UsersApi(api_client)
    user_uid = 'user_uid_example' # str | The unique identifier of the user in CDO.

    try:
        # Delete a User by UID in the CDO tenant
        api_instance.delete_user(user_uid)
    except Exception as e:
        print("Exception when calling UsersApi->delete_user: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_uid** | **str**| The unique identifier of the user in CDO. | 

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
> ApiTokenInfo generate_api_token(api_user_id)

Generate API Token for API-only user

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo-python-sdk
from cdo-python-sdk.models.api_token_info import ApiTokenInfo
from cdo-python-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo-python-sdk.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo-python-sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo-python-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo-python-sdk.UsersApi(api_client)
    api_user_id = 'api_user_id_example' # str | The unique identifier of the API user in CDO.

    try:
        # Generate API Token for API-only user
        api_response = api_instance.generate_api_token(api_user_id)
        print("The response of UsersApi->generate_api_token:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->generate_api_token: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **api_user_id** | **str**| The unique identifier of the API user in CDO. | 

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

# **get_token**
> CdoTokenInfo get_token()

Fetch information on the current token

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo-python-sdk
from cdo-python-sdk.models.cdo_token_info import CdoTokenInfo
from cdo-python-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo-python-sdk.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo-python-sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo-python-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo-python-sdk.UsersApi(api_client)

    try:
        # Fetch information on the current token
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

Fetch a user by UID in the CDO tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo-python-sdk
from cdo-python-sdk.models.user import User
from cdo-python-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo-python-sdk.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo-python-sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo-python-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo-python-sdk.UsersApi(api_client)
    user_uid = 'user_uid_example' # str | The unique identifier of the user in CDO.

    try:
        # Fetch a user by UID in the CDO tenant.
        api_response = api_instance.get_user(user_uid)
        print("The response of UsersApi->get_user:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->get_user: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user_uid** | **str**| The unique identifier of the user in CDO. | 

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

# **list_users**
> UserPage list_users(limit=limit, offset=offset, q=q)

Fetch a list of users associated with the CDO tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo-python-sdk
from cdo-python-sdk.models.user_page import UserPage
from cdo-python-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo-python-sdk.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo-python-sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo-python-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo-python-sdk.UsersApi(api_client)
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | The offset of the results retrieved. The CDO Public API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'name:London-Office-ASA' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Fetch a list of users associated with the CDO tenant.
        api_response = api_instance.list_users(limit=limit, offset=offset, q=q)
        print("The response of UsersApi->list_users:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->list_users: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| The offset of the results retrieved. The CDO Public API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
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

# **revoke_api_token**
> StatusInfo revoke_api_token(api_user_id)

Revoke API Token of API-only user

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo-python-sdk
from cdo-python-sdk.models.status_info import StatusInfo
from cdo-python-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo-python-sdk.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo-python-sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo-python-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo-python-sdk.UsersApi(api_client)
    api_user_id = 'api_user_id_example' # str | The unique identifier of the API user in CDO.

    try:
        # Revoke API Token of API-only user
        api_response = api_instance.revoke_api_token(api_user_id)
        print("The response of UsersApi->revoke_api_token:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling UsersApi->revoke_api_token: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **api_user_id** | **str**| The unique identifier of the API user in CDO. | 

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

Revoke the current token

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo-python-sdk
from cdo-python-sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo-python-sdk.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo-python-sdk.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo-python-sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo-python-sdk.UsersApi(api_client)

    try:
        # Revoke the current token
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

