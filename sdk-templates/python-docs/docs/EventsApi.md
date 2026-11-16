# scc_firewall_manager_sdk.EventsApi

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
[**get_sal_event_usage**](EventsApi.md#get_sal_event_usage) | **GET** /v1/events/usage | Get SAL Event Usage
[**get_search_reports**](EventsApi.md#get_search_reports) | **GET** /v1/events/reports | Get event search reports


# **get_sal_event_usage**
> EventUsageResponse get_sal_event_usage(time_range)

Get SAL Event Usage

Returns the amount of firewall and SD-WAN event data ingested for the tenant. SAL reports usage in monthly buckets, so each usage entry represents one UTC calendar month with an inclusive start date and exclusive end date. The response includes totals and monthly averages for the requested time range. Security Analytics and Logging (SAL) must be enabled for the tenant. Usage for the current month includes data ingested from the start of the current calendar month up to the time of the request.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.event_usage_response import EventUsageResponse
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
    api_instance = scc_firewall_manager_sdk.EventsApi(api_client)
    time_range = '3M' # str | The calendar-month range to return. The current month is included. Use 1M for the current month, 3M for the current and previous two months, 6M for the current and previous five months, or 1Y for the current and previous eleven months.

    try:
        # Get SAL Event Usage
        api_response = api_instance.get_sal_event_usage(time_range)
        print("The response of EventsApi->get_sal_event_usage:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EventsApi->get_sal_event_usage: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time_range** | **str**| The calendar-month range to return. The current month is included. Use 1M for the current month, 3M for the current and previous two months, 6M for the current and previous five months, or 1Y for the current and previous eleven months. | 

### Return type

[**EventUsageResponse**](EventUsageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | EventUsageResponse object. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_search_reports**
> Page get_search_reports(limit=limit, offset=offset, q=q, sort=sort)

Get event search reports

Get a list of reports containing the results of event log searches executed using the Report feature in the Event Logging page.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.page import Page
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
    api_instance = scc_firewall_manager_sdk.EventsApi(api_client)
    limit = '50' # str | Number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get event search reports
        api_response = api_instance.get_search_reports(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of EventsApi->get_search_reports:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling EventsApi->get_search_reports: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**Page**](Page.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of reports, including a presigned URL to download the reports |  -  |
**401** | Request not authorized. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

