# scc_firewall_manager_sdk.FTDServicesApi

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
[**get_tasks**](FTDServicesApi.md#get_tasks) | **GET** /v1/ftd-services/tasks | Get tasks
[**get_tasks_attribute_values**](FTDServicesApi.md#get_tasks_attribute_values) | **GET** /v1/ftd-services/tasks/attribute-values | Get distinct attribute values for tasks


# **get_tasks**
> TaskDtoPage get_tasks(limit=limit, offset=offset, q=q, sort=sort)

Get tasks

Get FTD Services tasks executed on this Security Cloud Control tenant.

### Example


```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.task_dto_page import TaskDtoPage
from scc_firewall_manager_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.us.security.cisco.com/firewall
# See configuration.py for a list of all supported configuration parameters.
configuration = scc_firewall_manager_sdk.Configuration(
    host = "https://api.us.security.cisco.com/firewall"
)


# Enter a context with an instance of the API client
with scc_firewall_manager_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = scc_firewall_manager_sdk.FTDServicesApi(api_client)
    limit = '50' # str | Number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get tasks
        api_response = api_instance.get_tasks(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of FTDServicesApi->get_tasks:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FTDServicesApi->get_tasks: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**TaskDtoPage**](TaskDtoPage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of tasks |  -  |
**401** | Request not authorized. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_tasks_attribute_values**
> TaskDistinctAttributeValues get_tasks_attribute_values()

Get distinct attribute values for tasks

Get distinct values for multiple fields in the FTD services tasks executed on the Security Cloud Control tenant.

### Example


```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.task_distinct_attribute_values import TaskDistinctAttributeValues
from scc_firewall_manager_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.us.security.cisco.com/firewall
# See configuration.py for a list of all supported configuration parameters.
configuration = scc_firewall_manager_sdk.Configuration(
    host = "https://api.us.security.cisco.com/firewall"
)


# Enter a context with an instance of the API client
with scc_firewall_manager_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = scc_firewall_manager_sdk.FTDServicesApi(api_client)

    try:
        # Get distinct attribute values for tasks
        api_response = api_instance.get_tasks_attribute_values()
        print("The response of FTDServicesApi->get_tasks_attribute_values:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling FTDServicesApi->get_tasks_attribute_values: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**TaskDistinctAttributeValues**](TaskDistinctAttributeValues.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Distinct attribute values for MSP-managed devices |  -  |
**401** | Request not authorized. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

