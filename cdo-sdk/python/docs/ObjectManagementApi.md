# cdo_sdk_python.ObjectManagementApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_object**](ObjectManagementApi.md#create_object) | **POST** /v0/objects | Create Object
[**create_targets**](ObjectManagementApi.md#create_targets) | **POST** /v0/objects/{uid}/targets | Create Targets
[**delete_object**](ObjectManagementApi.md#delete_object) | **DELETE** /v0/objects/{uid} | Delete Object
[**delete_targets**](ObjectManagementApi.md#delete_targets) | **DELETE** /v0/objects/{uid}/targets | Delete Targets
[**get_object**](ObjectManagementApi.md#get_object) | **GET** /v0/objects/{uid} | Get Object
[**get_object_duplicates**](ObjectManagementApi.md#get_object_duplicates) | **GET** /v0/objects/{uid}/duplicates | 
[**get_object_usages**](ObjectManagementApi.md#get_object_usages) | **GET** /v0/objects/{uid}/usage | Get Object Usages
[**get_objects**](ObjectManagementApi.md#get_objects) | **GET** /v0/objects | Get Objects
[**modify_object**](ObjectManagementApi.md#modify_object) | **PATCH** /v0/objects/{uid} | Modify Object


# **create_object**
> ObjectResponse create_object(create_request)

Create Object

Create an object in the CDO tenant.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.create_request import CreateRequest
from cdo_sdk_python.models.object_response import ObjectResponse
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
    api_instance = cdo_sdk_python.ObjectManagementApi(api_client)
    create_request = cdo_sdk_python.CreateRequest() # CreateRequest | 

    try:
        # Create Object
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

Create Targets

Create targets for an object in the CDO tenant.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.targets_request import TargetsRequest
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
    api_instance = cdo_sdk_python.ObjectManagementApi(api_client)
    uid = 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' # str | The request UID of the object for which the targets are being added to.
    targets_request = cdo_sdk_python.TargetsRequest() # TargetsRequest | 

    try:
        # Create Targets
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

Delete Object

Delete an object in the CDO tenant

### Example


```python
import cdo_sdk_python
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
    api_instance = cdo_sdk_python.ObjectManagementApi(api_client)
    uid = 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' # str | The request UID of the object being deleted.

    try:
        # Delete Object
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

# **delete_targets**
> object delete_targets(uid, target_uuids)

Delete Targets

Delete targets from an object in the CDO tenant.

### Example


```python
import cdo_sdk_python
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
    api_instance = cdo_sdk_python.ObjectManagementApi(api_client)
    uid = 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' # str | The request UID of the object for which the targets are being removed from.
    target_uuids = ['[xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx, yyyyyyyy-yyyy-yyyy-yyyy-yyyyyyyyyyyy]'] # List[str] | The list of UIDs of the targets being removed.

    try:
        # Delete Targets
        api_response = api_instance.delete_targets(uid, target_uuids)
        print("The response of ObjectManagementApi->delete_targets:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ObjectManagementApi->delete_targets: %s\n" % e)
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

# **get_object**
> ObjectResponse get_object(uid, fields=fields)

Get Object

Gets an object in the CDO tenant

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.object_response import ObjectResponse
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
    api_instance = cdo_sdk_python.ObjectManagementApi(api_client)
    uid = 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' # str | The request UID of the object being retrieved.
    fields = '@detailed' # str | The scope of the fields to be retrieved. One of [\"@basic\", \"@detailed\"]. Defaults to \"@basic\". (optional)

    try:
        # Get Object
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
**200** | OK |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_object_duplicates**
> get_object_duplicates(uid, target_id)



### Example


```python
import cdo_sdk_python
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
    api_instance = cdo_sdk_python.ObjectManagementApi(api_client)
    uid = 'uid_example' # str | 
    target_id = 'target_id_example' # str | 

    try:
        api_instance.get_object_duplicates(uid, target_id)
    except Exception as e:
        print("Exception when calling ObjectManagementApi->get_object_duplicates: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **str**|  | 
 **target_id** | **str**|  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_object_usages**
> List[ReferenceInfo] get_object_usages(uid)

Get Object Usages

Get usages of an object in the CDO tenant.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.reference_info import ReferenceInfo
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
    api_instance = cdo_sdk_python.ObjectManagementApi(api_client)
    uid = 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' # str | The request UID of the object to retrieve usages for.

    try:
        # Get Object Usages
        api_response = api_instance.get_object_usages(uid)
        print("The response of ObjectManagementApi->get_object_usages:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ObjectManagementApi->get_object_usages: %s\n" % e)
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

# **get_objects**
> ListObjectResponse get_objects(q=q, offset=offset, limit=limit, sort_by=sort_by)

Get Objects

Get objects in the CDO tenant.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.list_object_response import ListObjectResponse
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
    api_instance = cdo_sdk_python.ObjectManagementApi(api_client)
    q = 'name:London-Office-ASA' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    offset = '0' # str | The offset of the results retrieved. The CDO Public API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    sort_by = 'name:DESC' # str | The fields to sort results by. (optional)

    try:
        # Get Objects
        api_response = api_instance.get_objects(q=q, offset=offset, limit=limit, sort_by=sort_by)
        print("The response of ObjectManagementApi->get_objects:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ObjectManagementApi->get_objects: %s\n" % e)
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

# **modify_object**
> ObjectResponse modify_object(uid, update_request)

Modify Object

Modify an object in the CDO tenant

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.object_response import ObjectResponse
from cdo_sdk_python.models.update_request import UpdateRequest
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
    api_instance = cdo_sdk_python.ObjectManagementApi(api_client)
    uid = 'xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx' # str | The request UID of the object being updated.
    update_request = cdo_sdk_python.UpdateRequest() # UpdateRequest | 

    try:
        # Modify Object
        api_response = api_instance.modify_object(uid, update_request)
        print("The response of ObjectManagementApi->modify_object:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ObjectManagementApi->modify_object: %s\n" % e)
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

