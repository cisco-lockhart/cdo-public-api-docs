# scc_firewall_manager_sdk.MSPEventsApi

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
[**get_msp_sal_event_rate_daily**](MSPEventsApi.md#get_msp_sal_event_rate_daily) | **GET** /v1/msp/events/rates/daily | Get MSP SAL Daily Event Rate
[**get_msp_sal_event_rate_monthly**](MSPEventsApi.md#get_msp_sal_event_rate_monthly) | **GET** /v1/msp/events/rates/monthly | Get MSP SAL Monthly Event Rate
[**get_msp_sal_event_usage**](MSPEventsApi.md#get_msp_sal_event_usage) | **GET** /v1/msp/events/usage | Get MSP SAL Event Usage
[**get_msp_sal_event_usage_daily**](MSPEventsApi.md#get_msp_sal_event_usage_daily) | **GET** /v1/msp/events/usage/daily | Get MSP SAL Daily Event Usage


# **get_msp_sal_event_rate_daily**
> MspSalRateDailyPage get_msp_sal_event_rate_daily(time_range, limit=limit, offset=offset, q=q, sort=sort)

Get MSP SAL Daily Event Rate

Returns the average number of firewall and SD-WAN events per second sent by each managed tenant in the MSP portal that has Security Analytics and Logging (SAL) enabled. Each managed tenant includes the average and peak daily rate for the requested time range, and a rate entry for each stored, completed UTC day in that range. Days without data and the current day are not listed. Managed tenants without collected rates are returned with no entries. Results can be filtered, sorted, and paginated.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_sal_rate_daily_page import MspSalRateDailyPage
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
    api_instance = scc_firewall_manager_sdk.MSPEventsApi(api_client)
    time_range = 'time_range_example' # str | UTC calendar-month range to include.
    limit = '50' # str | Number of managed tenants to return. (optional) (default to '50')
    offset = '0' # str | Index of the first managed tenant to return. Must be 0 or a multiple of limit. (optional) (default to '0')
    q = 'managedTenantName:*soln23* AND managedTenantDisplayName:*auto* AND payType:INTERNAL' # str | Filter by managedTenantName, managedTenantDisplayName, or payType using Lucene syntax. Name fields use case-insensitive matching; use '*' for wildcard matching. payType uses an exact match. Only '*' is supported as a wildcard. (optional)
    sort = ['peakEventsPerSecond:desc'] # List[str] | Sort by averageEventsPerSecond or peakEventsPerSecond over the selected timeRange. Append :asc or :desc (uppercase directions are also accepted); a bare field sorts ascending. Repeat sort parameters to apply clauses in the supplied order. If sort is omitted, tenants are returned by UID in ascending order. (optional)

    try:
        # Get MSP SAL Daily Event Rate
        api_response = api_instance.get_msp_sal_event_rate_daily(time_range, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPEventsApi->get_msp_sal_event_rate_daily:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPEventsApi->get_msp_sal_event_rate_daily: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time_range** | **str**| UTC calendar-month range to include. | 
 **limit** | **str**| Number of managed tenants to return. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| Index of the first managed tenant to return. Must be 0 or a multiple of limit. | [optional] [default to &#39;0&#39;]
 **q** | **str**| Filter by managedTenantName, managedTenantDisplayName, or payType using Lucene syntax. Name fields use case-insensitive matching; use &#39;*&#39; for wildcard matching. payType uses an exact match. Only &#39;*&#39; is supported as a wildcard. | [optional] 
 **sort** | [**List[str]**](str.md)| Sort by averageEventsPerSecond or peakEventsPerSecond over the selected timeRange. Append :asc or :desc (uppercase directions are also accepted); a bare field sorts ascending. Repeat sort parameters to apply clauses in the supplied order. If sort is omitted, tenants are returned by UID in ascending order. | [optional] 

### Return type

[**MspSalRateDailyPage**](MspSalRateDailyPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | SAL daily event rates for enabled managed tenants |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_sal_event_rate_monthly**
> MspSalRateMonthlyPage get_msp_sal_event_rate_monthly(time_range, limit=limit, offset=offset, q=q, sort=sort)

Get MSP SAL Monthly Event Rate

Returns the average number of firewall and SD-WAN events per second sent by each managed tenant in the MSP portal that has Security Analytics and Logging (SAL) enabled. Each managed tenant includes the average and peak monthly rate for the requested time range, and a rate entry for each stored UTC calendar month in that range. The current month's rate is null until its first UTC day is over. Managed tenants without collected rates are returned with no entries. Results can be filtered, sorted, and paginated.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_sal_rate_monthly_page import MspSalRateMonthlyPage
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
    api_instance = scc_firewall_manager_sdk.MSPEventsApi(api_client)
    time_range = 'time_range_example' # str | UTC calendar-month range to include.
    limit = '50' # str | Number of managed tenants to return. (optional) (default to '50')
    offset = '0' # str | Index of the first managed tenant to return. Must be 0 or a multiple of limit. (optional) (default to '0')
    q = 'managedTenantName:*soln23* AND managedTenantDisplayName:*auto* AND payType:INTERNAL' # str | Filter by managedTenantName, managedTenantDisplayName, or payType using Lucene syntax. Name fields use case-insensitive matching; use '*' for wildcard matching. payType uses an exact match. Only '*' is supported as a wildcard. (optional)
    sort = ['peakEventsPerSecond:desc'] # List[str] | Sort by averageEventsPerSecond or peakEventsPerSecond over the selected timeRange. Append :asc or :desc (uppercase directions are also accepted); a bare field sorts ascending. Repeat sort parameters to apply clauses in the supplied order. If sort is omitted, tenants are returned by UID in ascending order. (optional)

    try:
        # Get MSP SAL Monthly Event Rate
        api_response = api_instance.get_msp_sal_event_rate_monthly(time_range, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPEventsApi->get_msp_sal_event_rate_monthly:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPEventsApi->get_msp_sal_event_rate_monthly: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time_range** | **str**| UTC calendar-month range to include. | 
 **limit** | **str**| Number of managed tenants to return. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| Index of the first managed tenant to return. Must be 0 or a multiple of limit. | [optional] [default to &#39;0&#39;]
 **q** | **str**| Filter by managedTenantName, managedTenantDisplayName, or payType using Lucene syntax. Name fields use case-insensitive matching; use &#39;*&#39; for wildcard matching. payType uses an exact match. Only &#39;*&#39; is supported as a wildcard. | [optional] 
 **sort** | [**List[str]**](str.md)| Sort by averageEventsPerSecond or peakEventsPerSecond over the selected timeRange. Append :asc or :desc (uppercase directions are also accepted); a bare field sorts ascending. Repeat sort parameters to apply clauses in the supplied order. If sort is omitted, tenants are returned by UID in ascending order. | [optional] 

### Return type

[**MspSalRateMonthlyPage**](MspSalRateMonthlyPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | SAL monthly event rates for enabled managed tenants |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_sal_event_usage**
> MspSalUsagePage get_msp_sal_event_usage(time_range, limit=limit, offset=offset, q=q, sort=sort)

Get MSP SAL Event Usage

Returns the amount of firewall and SD-WAN event data ingested by each managed tenant in the MSP portal that has Security Analytics and Logging (SAL) enabled. Each managed tenant includes totals and monthly averages for the requested time range, and a usage entry for each UTC calendar month in that range. Results can be filtered, sorted, and paginated.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_sal_usage_page import MspSalUsagePage
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
    api_instance = scc_firewall_manager_sdk.MSPEventsApi(api_client)
    time_range = 'time_range_example' # str | UTC calendar-month range to include.
    limit = '50' # str | Number of managed tenants to return. (optional) (default to '50')
    offset = '0' # str | Index of the first managed tenant to return. Must be 0 or a multiple of limit. (optional) (default to '0')
    q = 'managedTenantName:*soln23* AND managedTenantDisplayName:*auto* AND payType:INTERNAL' # str | Filter by managedTenantName, managedTenantDisplayName, or payType using Lucene syntax. Name fields use case-insensitive matching; use '*' for wildcard matching. payType uses an exact match. Only '*' is supported as a wildcard. (optional)
    sort = ['total.sdwanBytes:asc'] # List[str] | Sort by total.firewallBytes or total.sdwanBytes summed over the selected timeRange. Append :asc or :desc (uppercase directions are also accepted); a bare field sorts ascending. Repeat sort parameters to apply clauses in the supplied order. If sort is omitted, tenants are returned by UID in ascending order. (optional)

    try:
        # Get MSP SAL Event Usage
        api_response = api_instance.get_msp_sal_event_usage(time_range, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPEventsApi->get_msp_sal_event_usage:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPEventsApi->get_msp_sal_event_usage: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time_range** | **str**| UTC calendar-month range to include. | 
 **limit** | **str**| Number of managed tenants to return. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| Index of the first managed tenant to return. Must be 0 or a multiple of limit. | [optional] [default to &#39;0&#39;]
 **q** | **str**| Filter by managedTenantName, managedTenantDisplayName, or payType using Lucene syntax. Name fields use case-insensitive matching; use &#39;*&#39; for wildcard matching. payType uses an exact match. Only &#39;*&#39; is supported as a wildcard. | [optional] 
 **sort** | [**List[str]**](str.md)| Sort by total.firewallBytes or total.sdwanBytes summed over the selected timeRange. Append :asc or :desc (uppercase directions are also accepted); a bare field sorts ascending. Repeat sort parameters to apply clauses in the supplied order. If sort is omitted, tenants are returned by UID in ascending order. | [optional] 

### Return type

[**MspSalUsagePage**](MspSalUsagePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | SAL event usage for enabled managed tenants |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_sal_event_usage_daily**
> MspSalUsageDailyPage get_msp_sal_event_usage_daily(time_range, limit=limit, offset=offset, q=q, sort=sort)

Get MSP SAL Daily Event Usage

Returns the amount of firewall and SD-WAN event data ingested per day by each managed tenant in the MSP portal that has Security Analytics and Logging (SAL) enabled. Each managed tenant includes totals and daily averages for the requested time range, and a usage entry for each completed UTC day in that range, from the 1st of its first calendar month up to and including yesterday (UTC). Days without stored usage count as 0. Results can be filtered, sorted, and paginated.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_sal_usage_daily_page import MspSalUsageDailyPage
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
    api_instance = scc_firewall_manager_sdk.MSPEventsApi(api_client)
    time_range = 'time_range_example' # str | UTC calendar-month range to include.
    limit = '50' # str | Number of managed tenants to return. (optional) (default to '50')
    offset = '0' # str | Index of the first managed tenant to return. Must be 0 or a multiple of limit. (optional) (default to '0')
    q = 'managedTenantName:*soln23* AND managedTenantDisplayName:*auto* AND payType:INTERNAL' # str | Filter by managedTenantName, managedTenantDisplayName, or payType using Lucene syntax. Name fields use case-insensitive matching; use '*' for wildcard matching. payType uses an exact match. Only '*' is supported as a wildcard. (optional)
    sort = ['total.sdwanBytes:asc'] # List[str] | Sort by total.firewallBytes or total.sdwanBytes summed over the selected timeRange. Append :asc or :desc (uppercase directions are also accepted); a bare field sorts ascending. Repeat sort parameters to apply clauses in the supplied order. If sort is omitted, tenants are returned by UID in ascending order. (optional)

    try:
        # Get MSP SAL Daily Event Usage
        api_response = api_instance.get_msp_sal_event_usage_daily(time_range, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPEventsApi->get_msp_sal_event_usage_daily:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPEventsApi->get_msp_sal_event_usage_daily: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **time_range** | **str**| UTC calendar-month range to include. | 
 **limit** | **str**| Number of managed tenants to return. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| Index of the first managed tenant to return. Must be 0 or a multiple of limit. | [optional] [default to &#39;0&#39;]
 **q** | **str**| Filter by managedTenantName, managedTenantDisplayName, or payType using Lucene syntax. Name fields use case-insensitive matching; use &#39;*&#39; for wildcard matching. payType uses an exact match. Only &#39;*&#39; is supported as a wildcard. | [optional] 
 **sort** | [**List[str]**](str.md)| Sort by total.firewallBytes or total.sdwanBytes summed over the selected timeRange. Append :asc or :desc (uppercase directions are also accepted); a bare field sorts ascending. Repeat sort parameters to apply clauses in the supplied order. If sort is omitted, tenants are returned by UID in ascending order. | [optional] 

### Return type

[**MspSalUsageDailyPage**](MspSalUsageDailyPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Daily SAL event usage for enabled managed tenants |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

