# cdo_sdk_python.ASAAccessGroupsApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_access_group**](ASAAccessGroupsApi.md#create_access_group) | **POST** /v1/policies/asa/accessgroups | Create Access Group
[**delete_access_group**](ASAAccessGroupsApi.md#delete_access_group) | **DELETE** /v1/policies/asa/accessgroups/{accessGroupUid} | Delete Access Group
[**fetch_access_group**](ASAAccessGroupsApi.md#fetch_access_group) | **GET** /v1/policies/asa/accessgroups/{accessGroupUid} | Get Access Group
[**list_access_groups**](ASAAccessGroupsApi.md#list_access_groups) | **GET** /v1/policies/asa/accessgroups | Get Access Groups


# **create_access_group**
> AccessGroup create_access_group(access_group_create_or_update_input)

Create Access Group

Create an Access Group

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.access_group import AccessGroup
from cdo_sdk_python.models.access_group_create_or_update_input import AccessGroupCreateOrUpdateInput
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
    api_instance = cdo_sdk_python.ASAAccessGroupsApi(api_client)
    access_group_create_or_update_input = cdo_sdk_python.AccessGroupCreateOrUpdateInput() # AccessGroupCreateOrUpdateInput | 

    try:
        # Create Access Group
        api_response = api_instance.create_access_group(access_group_create_or_update_input)
        print("The response of ASAAccessGroupsApi->create_access_group:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAAccessGroupsApi->create_access_group: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **access_group_create_or_update_input** | [**AccessGroupCreateOrUpdateInput**](AccessGroupCreateOrUpdateInput.md)|  | 

### Return type

[**AccessGroup**](AccessGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | CDO Access Group |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_access_group**
> delete_access_group(access_group_uid)

Delete Access Group

Delete Access Group by UID in the CDO tenant.

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
    api_instance = cdo_sdk_python.ASAAccessGroupsApi(api_client)
    access_group_uid = 'access_group_uid_example' # str | The unique identifier of the Access Group in CDO.

    try:
        # Delete Access Group
        api_instance.delete_access_group(access_group_uid)
    except Exception as e:
        print("Exception when calling ASAAccessGroupsApi->delete_access_group: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **access_group_uid** | **str**| The unique identifier of the Access Group in CDO. | 

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
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **fetch_access_group**
> AccessGroup fetch_access_group(access_group_uid)

Get Access Group

Get a single ASA Access Group by UUID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.access_group import AccessGroup
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
    api_instance = cdo_sdk_python.ASAAccessGroupsApi(api_client)
    access_group_uid = 'access_group_uid_example' # str | The unique identifier of the Access Group in CDO.

    try:
        # Get Access Group
        api_response = api_instance.fetch_access_group(access_group_uid)
        print("The response of ASAAccessGroupsApi->fetch_access_group:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAAccessGroupsApi->fetch_access_group: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **access_group_uid** | **str**| The unique identifier of the Access Group in CDO. | 

### Return type

[**AccessGroup**](AccessGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Access Group object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_access_groups**
> AccessGroupPage list_access_groups(limit=limit, offset=offset, q=q, sort=sort)

Get Access Groups

Get a list of ASA Access Groups.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.access_group_page import AccessGroupPage
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
    api_instance = cdo_sdk_python.ASAAccessGroupsApi(api_client)
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Access Groups
        api_response = api_instance.list_access_groups(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of ASAAccessGroupsApi->list_access_groups:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAAccessGroupsApi->list_access_groups: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**AccessGroupPage**](AccessGroupPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Access Groups |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

