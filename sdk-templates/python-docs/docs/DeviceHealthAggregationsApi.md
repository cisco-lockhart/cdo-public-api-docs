# scc_firewall_manager_sdk.DeviceHealthAggregationsApi

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
[**get_device_health_metric_aggregation_list**](DeviceHealthAggregationsApi.md#get_device_health_metric_aggregation_list) | **GET** /v1/inventory/devices/health/metrics/aggregations/details | Get device detailed list for aggregations
[**get_device_health_metric_aggregations**](DeviceHealthAggregationsApi.md#get_device_health_metric_aggregations) | **GET** /v1/inventory/devices/health/metrics/aggregations | Get device health metric aggregations


# **get_device_health_metric_aggregation_list**
> MetricAggregationListResponse get_device_health_metric_aggregation_list(metric, threshold, aggregation_period=aggregation_period, managed_tenant_uid=managed_tenant_uid, q=q)

Get device detailed list for aggregations

Retrieve filtered device detailed list corresponding to aggregation metrics.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.metric_aggregation_list_response import MetricAggregationListResponse
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
    api_instance = scc_firewall_manager_sdk.DeviceHealthAggregationsApi(api_client)
    metric = 'cpu' # str | The metric for the device health aggregation detailed list.
    threshold = 'CRITICAL' # str | The metric threshold (e.g. CRITICAL).
    aggregation_period = '24h' # str | The aggregation period of the metrics returned. (optional)
    managed_tenant_uid = '7131daad-e813-4b8f-8f42-be1e241e8cdb' # str | A managed tenant UUID to filter list, if applicable. (optional)
    q = 'name:my-devi*' # str | Lucene-style query filter. Only prefix queries (e.g. field:value*) and wildcard queries (e.g. field:*value or field:*value*) are supported for filtering. (optional)

    try:
        # Get device detailed list for aggregations
        api_response = api_instance.get_device_health_metric_aggregation_list(metric, threshold, aggregation_period=aggregation_period, managed_tenant_uid=managed_tenant_uid, q=q)
        print("The response of DeviceHealthAggregationsApi->get_device_health_metric_aggregation_list:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceHealthAggregationsApi->get_device_health_metric_aggregation_list: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metric** | **str**| The metric for the device health aggregation detailed list. | 
 **threshold** | **str**| The metric threshold (e.g. CRITICAL). | 
 **aggregation_period** | **str**| The aggregation period of the metrics returned. | [optional] 
 **managed_tenant_uid** | **str**| A managed tenant UUID to filter list, if applicable. | [optional] 
 **q** | **str**| Lucene-style query filter. Only prefix queries (e.g. field:value*) and wildcard queries (e.g. field:*value or field:*value*) are supported for filtering. | [optional] 

### Return type

[**MetricAggregationListResponse**](MetricAggregationListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successfully retrieved device metric aggregation detailed list |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device_health_metric_aggregations**
> MetricAggregationResponse get_device_health_metric_aggregations(metrics=metrics, aggregation_period=aggregation_period)

Get device health metric aggregations

Retrieve aggregation device health metrics for all managed devices.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.metric_aggregation_response import MetricAggregationResponse
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
    api_instance = scc_firewall_manager_sdk.DeviceHealthAggregationsApi(api_client)
    metrics = 'cpu,mem,disk' # str | Comma-separated list of metrics to return (e.g. cpu,mem). Returns all if omitted. (optional)
    aggregation_period = '24h' # str | The aggregation period of the metrics returned (optional)

    try:
        # Get device health metric aggregations
        api_response = api_instance.get_device_health_metric_aggregations(metrics=metrics, aggregation_period=aggregation_period)
        print("The response of DeviceHealthAggregationsApi->get_device_health_metric_aggregations:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceHealthAggregationsApi->get_device_health_metric_aggregations: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metrics** | **str**| Comma-separated list of metrics to return (e.g. cpu,mem). Returns all if omitted. | [optional] 
 **aggregation_period** | **str**| The aggregation period of the metrics returned | [optional] 

### Return type

[**MetricAggregationResponse**](MetricAggregationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successfully retrieved device metric aggregations. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

