# cdo-python-sdk.ChangeRequestsApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_change_request**](ChangeRequestsApi.md#create_change_request) | **POST** /v1/changeRequests | Create Change Requests.
[**delete_change_request**](ChangeRequestsApi.md#delete_change_request) | **DELETE** /v1/changeRequests/{changeRequestUid} | Delete a Change Request by UID in the CDO tenant
[**get_change_request**](ChangeRequestsApi.md#get_change_request) | **GET** /v1/changeRequests/{changeRequestUid} | Fetch a Change Request by UID in the CDO tenant.
[**list_change_requests**](ChangeRequestsApi.md#list_change_requests) | **GET** /v1/changeRequests | Fetch a list of Change Requests.


# **create_change_request**
> ChangeRequest create_change_request(change_request_create_input)

Create Change Requests.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo-python-sdk
from cdo-python-sdk.models.change_request import ChangeRequest
from cdo-python-sdk.models.change_request_create_input import ChangeRequestCreateInput
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
    api_instance = cdo-python-sdk.ChangeRequestsApi(api_client)
    change_request_create_input = cdo-python-sdk.ChangeRequestCreateInput() # ChangeRequestCreateInput | 

    try:
        # Create Change Requests.
        api_response = api_instance.create_change_request(change_request_create_input)
        print("The response of ChangeRequestsApi->create_change_request:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChangeRequestsApi->create_change_request: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **change_request_create_input** | [**ChangeRequestCreateInput**](ChangeRequestCreateInput.md)|  | 

### Return type

[**ChangeRequest**](ChangeRequest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Created change Request object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_change_request**
> delete_change_request(change_request_uid)

Delete a Change Request by UID in the CDO tenant

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
    api_instance = cdo-python-sdk.ChangeRequestsApi(api_client)
    change_request_uid = 'change_request_uid_example' # str | The unique identifier of the Change Request in CDO.

    try:
        # Delete a Change Request by UID in the CDO tenant
        api_instance.delete_change_request(change_request_uid)
    except Exception as e:
        print("Exception when calling ChangeRequestsApi->delete_change_request: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **change_request_uid** | **str**| The unique identifier of the Change Request in CDO. | 

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
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_change_request**
> ChangeRequest get_change_request(change_request_uid)

Fetch a Change Request by UID in the CDO tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo-python-sdk
from cdo-python-sdk.models.change_request import ChangeRequest
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
    api_instance = cdo-python-sdk.ChangeRequestsApi(api_client)
    change_request_uid = 'change_request_uid_example' # str | The unique identifier of the Change Request in CDO.

    try:
        # Fetch a Change Request by UID in the CDO tenant.
        api_response = api_instance.get_change_request(change_request_uid)
        print("The response of ChangeRequestsApi->get_change_request:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChangeRequestsApi->get_change_request: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **change_request_uid** | **str**| The unique identifier of the Change Request in CDO. | 

### Return type

[**ChangeRequest**](ChangeRequest.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Change Request object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_change_requests**
> ChangeRequestPage list_change_requests(limit=limit, offset=offset, q=q)

Fetch a list of Change Requests.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo-python-sdk
from cdo-python-sdk.models.change_request_page import ChangeRequestPage
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
    api_instance = cdo-python-sdk.ChangeRequestsApi(api_client)
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | The offset of the results retrieved. The CDO Public API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'name:London-Office-ASA' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Fetch a list of Change Requests.
        api_response = api_instance.list_change_requests(limit=limit, offset=offset, q=q)
        print("The response of ChangeRequestsApi->list_change_requests:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChangeRequestsApi->list_change_requests: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| The offset of the results retrieved. The CDO Public API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 

### Return type

[**ChangeRequestPage**](ChangeRequestPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Change Request objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

