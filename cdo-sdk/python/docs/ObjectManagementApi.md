# openapi_client.ObjectManagementApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_object**](ObjectManagementApi.md#create_object) | **POST** /v0/objects | Creates an object in the CDO tenant
[**create_targets**](ObjectManagementApi.md#create_targets) | **POST** /v0/objects/{uid}/targets | Create targets for an object in the CDO tenant
[**delete_object**](ObjectManagementApi.md#delete_object) | **DELETE** /v0/objects/{uid} | Deletes an object in the CDO tenant
[**get_object**](ObjectManagementApi.md#get_object) | **GET** /v0/objects/{uid} | Retrieves an object in the CDO tenant
[**get_object_usage**](ObjectManagementApi.md#get_object_usage) | **GET** /v0/objects/{uid}/usage | Retrieves usages of an object in the CDO tenant
[**list_objects**](ObjectManagementApi.md#list_objects) | **GET** /v0/objects | Retrieves objects in the CDO tenant
[**remove_targets**](ObjectManagementApi.md#remove_targets) | **DELETE** /v0/objects/{uid}/targets | Removes targets from an object in the CDO tenant
[**update_object**](ObjectManagementApi.md#update_object) | **PATCH** /v0/objects/{uid} | Updates an object in the CDO tenant


# **create_object**
> ObjectResponse create_object(create_request)

Creates an object in the CDO tenant

### Example


```python
import openapi_client
from openapi_client.models.create_request import CreateRequest
from openapi_client.models.object_response import ObjectResponse
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ObjectManagementApi(api_client)
    create_request = openapi_client.CreateRequest() # CreateRequest | 

    try:
        # Creates an object in the CDO tenant
        api_response = api_instance.create_object(create_request)
        print("The response of ObjectManagementApi->create_object:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ObjectManagementApi->create_object: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **create_request** | [**CreateRequest**](CreateRequest.md)|  | 

### Return type

[**ObjectResponse**](ObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | A descriptive representation of the created CDO object. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_targets**
> object create_targets(uid, targets_request)

Create targets for an object in the CDO tenant

### Example


```python
import openapi_client
from openapi_client.models.targets_request import TargetsRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ObjectManagementApi(api_client)
    uid = 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' # str | The request UID of the object for which the targets are being added to.
    targets_request = openapi_client.TargetsRequest() # TargetsRequest | 

    try:
        # Create targets for an object in the CDO tenant
        api_response = api_instance.create_targets(uid, targets_request)
        print("The response of ObjectManagementApi->create_targets:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ObjectManagementApi->create_targets: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **str**| The request UID of the object for which the targets are being added to. | 
 **targets_request** | [**TargetsRequest**](TargetsRequest.md)|  | 

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*, application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | No content |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Not found |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_object**
> object delete_object(uid)

Deletes an object in the CDO tenant

### Example


```python
import openapi_client
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ObjectManagementApi(api_client)
    uid = 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' # str | The request UID of the object being deleted.

    try:
        # Deletes an object in the CDO tenant
        api_response = api_instance.delete_object(uid)
        print("The response of ObjectManagementApi->delete_object:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ObjectManagementApi->delete_object: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **str**| The request UID of the object being deleted. | 

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | No content |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_object**
> ObjectResponse get_object(uid, fields=fields)

Retrieves an object in the CDO tenant

### Example


```python
import openapi_client
from openapi_client.models.object_response import ObjectResponse
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ObjectManagementApi(api_client)
    uid = 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' # str | The request UID of the object being retrieved.
    fields = '@detailed' # str | The scope of the fields to be retrieved. One of [\"@basic\", \"@detailed\"]. Defaults to \"@basic\". (optional)

    try:
        # Retrieves an object in the CDO tenant
        api_response = api_instance.get_object(uid, fields=fields)
        print("The response of ObjectManagementApi->get_object:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ObjectManagementApi->get_object: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **str**| The request UID of the object being retrieved. | 
 **fields** | **str**| The scope of the fields to be retrieved. One of [\&quot;@basic\&quot;, \&quot;@detailed\&quot;]. Defaults to \&quot;@basic\&quot;. | [optional] 

### Return type

[**ObjectResponse**](ObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Fetch an object by UID in the CDO tenant |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_object_usage**
> List[ReferenceInfo] get_object_usage(uid)

Retrieves usages of an object in the CDO tenant

### Example


```python
import openapi_client
from openapi_client.models.reference_info import ReferenceInfo
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ObjectManagementApi(api_client)
    uid = 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' # str | The request UID of the object to retrieve usages for.

    try:
        # Retrieves usages of an object in the CDO tenant
        api_response = api_instance.get_object_usage(uid)
        print("The response of ObjectManagementApi->get_object_usage:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ObjectManagementApi->get_object_usage: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **str**| The request UID of the object to retrieve usages for. | 

### Return type

[**List[ReferenceInfo]**](ReferenceInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A descriptive representation of the created CDO object. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_objects**
> ListObjectResponse list_objects(q=q, offset=offset, limit=limit, sort_by=sort_by)

Retrieves objects in the CDO tenant

### Example


```python
import openapi_client
from openapi_client.models.list_object_response import ListObjectResponse
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ObjectManagementApi(api_client)
    q = 'name:London-Office-ASA' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    offset = '0' # str | The offset of the results retrieved. The CDO Public API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    sort_by = 'name:DESC' # str | The fields to sort results by. (optional)

    try:
        # Retrieves objects in the CDO tenant
        api_response = api_instance.list_objects(q=q, offset=offset, limit=limit, sort_by=sort_by)
        print("The response of ObjectManagementApi->list_objects:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ObjectManagementApi->list_objects: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **offset** | **str**| The offset of the results retrieved. The CDO Public API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **sort_by** | **str**| The fields to sort results by. | [optional] 

### Return type

[**ListObjectResponse**](ListObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A paginated view of the CDO objects. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **remove_targets**
> object remove_targets(uid, target_uuids)

Removes targets from an object in the CDO tenant

### Example


```python
import openapi_client
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ObjectManagementApi(api_client)
    uid = 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' # str | The request UID of the object for which the targets are being removed from.
    target_uuids = ['[xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx, yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy]'] # List[str] | The list of UIDs of the targets being removed.

    try:
        # Removes targets from an object in the CDO tenant
        api_response = api_instance.remove_targets(uid, target_uuids)
        print("The response of ObjectManagementApi->remove_targets:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ObjectManagementApi->remove_targets: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **str**| The request UID of the object for which the targets are being removed from. | 
 **target_uuids** | [**List[str]**](str.md)| The list of UIDs of the targets being removed. | 

### Return type

**object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | No content |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Not found |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_object**
> ObjectResponse update_object(uid, update_request)

Updates an object in the CDO tenant

### Example


```python
import openapi_client
from openapi_client.models.object_response import ObjectResponse
from openapi_client.models.update_request import UpdateRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.ObjectManagementApi(api_client)
    uid = 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' # str | The request UID of the object being updated.
    update_request = openapi_client.UpdateRequest() # UpdateRequest | 

    try:
        # Updates an object in the CDO tenant
        api_response = api_instance.update_object(uid, update_request)
        print("The response of ObjectManagementApi->update_object:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ObjectManagementApi->update_object: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **str**| The request UID of the object being updated. | 
 **update_request** | [**UpdateRequest**](UpdateRequest.md)|  | 

### Return type

[**ObjectResponse**](ObjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A descriptive representation of the updated CDO object. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

