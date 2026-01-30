# scc_firewall_manager_sdk.DeviceHealthApi

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
[**get_asa_health_metrics**](DeviceHealthApi.md#get_asa_health_metrics) | **GET** /v1/inventory/devices/asas/health/metrics | Get time-series health metrics for one or more ASA devices
[**get_asa_interface_health_metrics**](DeviceHealthApi.md#get_asa_interface_health_metrics) | **GET** /v1/inventory/devices/asas/health/{deviceUid}/interfaces | Get time-series interface metrics for an ASA device
[**opt_in_to_asa_health_metrics**](DeviceHealthApi.md#opt_in_to_asa_health_metrics) | **POST** /v1/inventory/devices/asas/health/metrics | Opt in to ASA Health Metrics
[**opt_out_of_asa_health_metrics**](DeviceHealthApi.md#opt_out_of_asa_health_metrics) | **DELETE** /v1/inventory/devices/asas/health/metrics | Opt out of ASA Health Metrics


# **get_asa_health_metrics**
> MetricsResponse get_asa_health_metrics(start=start, end=end, time_range=time_range, limit=limit, offset=offset, device_uids=device_uids, metrics=metrics)

Get time-series health metrics for one or more ASA devices

Returns time-series metrics such as CPU, memory, disk, connections, and environment data over a specified time range. Supports filtering by device and metric. Pagination applies only when device UIDs are not specified.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.metrics_response import MetricsResponse
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
    api_instance = scc_firewall_manager_sdk.DeviceHealthApi(api_client)
    start = '2025-04-05T00:00:00Z' # str | Start of the time range (ISO 8601 format). (optional)
    end = '2025-04-05T04:00:00Z' # str | End of the time range (ISO 8601 format). (optional)
    time_range = '10m' # str | Relative time range (mutually exclusive with start/end). Data is collected every 10 minutes, so shorter time ranges may return fewer data points. (optional)
    limit = 'limit_example' # str | Maximum number of device records to return (used only when no deviceUids are provided). (optional)
    offset = 'offset_example' # str | Offset for pagination (used only when no deviceUids are provided). (optional)
    device_uids = '256461f6-bd60-11ef-8beb-6cf1610cf55d,41a1d57b-ffc2-49aa-933b-440cdd76b2fc' # str | Comma-separated list of device UIDs to query. Max 50. If omitted, results are paginated. (optional)
    metrics = 'cpu,mem' # str | Comma-separated list of metrics to return (e.g. cpu, mem). Returns all if omitted. (optional)

    try:
        # Get time-series health metrics for one or more ASA devices
        api_response = api_instance.get_asa_health_metrics(start=start, end=end, time_range=time_range, limit=limit, offset=offset, device_uids=device_uids, metrics=metrics)
        print("The response of DeviceHealthApi->get_asa_health_metrics:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceHealthApi->get_asa_health_metrics: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **str**| Start of the time range (ISO 8601 format). | [optional] 
 **end** | **str**| End of the time range (ISO 8601 format). | [optional] 
 **time_range** | **str**| Relative time range (mutually exclusive with start/end). Data is collected every 10 minutes, so shorter time ranges may return fewer data points. | [optional] 
 **limit** | **str**| Maximum number of device records to return (used only when no deviceUids are provided). | [optional] 
 **offset** | **str**| Offset for pagination (used only when no deviceUids are provided). | [optional] 
 **device_uids** | **str**| Comma-separated list of device UIDs to query. Max 50. If omitted, results are paginated. | [optional] 
 **metrics** | **str**| Comma-separated list of metrics to return (e.g. cpu, mem). Returns all if omitted. | [optional] 

### Return type

[**MetricsResponse**](MetricsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successfully retrieved time-series metrics for one or more ASA devices. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_interface_health_metrics**
> MetricsResponse get_asa_interface_health_metrics(device_uid, start=start, end=end, time_range=time_range, limit=limit, offset=offset, interface_uids=interface_uids, metrics=metrics, q=q)

Get time-series interface metrics for an ASA device

Returns time-series interface metrics - including link status, overruns, and other key performance indicators - over a specified time range. Supports filtering by interface and metric. Pagination applies only when interface UIDs are not specified.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.metrics_response import MetricsResponse
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
    api_instance = scc_firewall_manager_sdk.DeviceHealthApi(api_client)
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the ASA device in Security Cloud Control.
    start = '2025-04-05T00:00:00Z' # str | Start of the time range (ISO 8601 format). (optional)
    end = '2025-04-05T04:00:00Z' # str | End of the time range (ISO 8601 format). (optional)
    time_range = '10m' # str | Relative time range (mutually exclusive with start/end). Data is collected every 10 minutes, so shorter time ranges may return fewer data points. (optional)
    limit = 'limit_example' # str | Maximum number of device records to return (used only when no deviceUids are provided). (optional)
    offset = 'offset_example' # str | Offset for pagination (used only when no deviceUids are provided). (optional)
    interface_uids = 'def6c34a-a5d4-45c1-8d96-820fb00de723,4bc125a0-cf08-4dcc-bfac-45669ba898b0' # str | Comma-separated list of interface UIDs to query. Max 50. If omitted, results are paginated. (optional)
    metrics = 'linkStatus,overruns' # str | Comma-separated list of metrics to return (e.g. linkStatus, overruns). Returns all if omitted. (optional)
    q = 'name:GigabitEthernet* OR linkStatus:1' # str | Lucene-style query filter. Supports logical operators AND, OR, and NOT.Applicable only to the fields name and linkStatus. (optional)

    try:
        # Get time-series interface metrics for an ASA device
        api_response = api_instance.get_asa_interface_health_metrics(device_uid, start=start, end=end, time_range=time_range, limit=limit, offset=offset, interface_uids=interface_uids, metrics=metrics, q=q)
        print("The response of DeviceHealthApi->get_asa_interface_health_metrics:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceHealthApi->get_asa_interface_health_metrics: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the ASA device in Security Cloud Control. | 
 **start** | **str**| Start of the time range (ISO 8601 format). | [optional] 
 **end** | **str**| End of the time range (ISO 8601 format). | [optional] 
 **time_range** | **str**| Relative time range (mutually exclusive with start/end). Data is collected every 10 minutes, so shorter time ranges may return fewer data points. | [optional] 
 **limit** | **str**| Maximum number of device records to return (used only when no deviceUids are provided). | [optional] 
 **offset** | **str**| Offset for pagination (used only when no deviceUids are provided). | [optional] 
 **interface_uids** | **str**| Comma-separated list of interface UIDs to query. Max 50. If omitted, results are paginated. | [optional] 
 **metrics** | **str**| Comma-separated list of metrics to return (e.g. linkStatus, overruns). Returns all if omitted. | [optional] 
 **q** | **str**| Lucene-style query filter. Supports logical operators AND, OR, and NOT.Applicable only to the fields name and linkStatus. | [optional] 

### Return type

[**MetricsResponse**](MetricsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successfully retrieved time-series metrics for one or more interfaces. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opt_in_to_asa_health_metrics**
> CdoTransaction opt_in_to_asa_health_metrics()

Opt in to ASA Health Metrics

Allow Security Cloud Control Firewall Manager to collect health metrics for ASAs in this tenant. By default, metrics are collected from all Cloud Device Gateway–managed ASAs. To let Security Cloud Control Firewall Manager collect metrics from Secure Device Connector–managed devices, set the <code>optedInToAsaHealthMetrics</code> property on each device using the <a href=\"https://developer.cisco.com/docs/cisco-security-cloud-control/modify-device/\">Modify Device</a> endpoint. <b>Note:</b> A single Secure Device Connector can support metrics collection for up to 50 ASAs.

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
    api_instance = scc_firewall_manager_sdk.DeviceHealthApi(api_client)

    try:
        # Opt in to ASA Health Metrics
        api_response = api_instance.opt_in_to_asa_health_metrics()
        print("The response of DeviceHealthApi->opt_in_to_asa_health_metrics:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceHealthApi->opt_in_to_asa_health_metrics: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Security Cloud Control Transaction object that can be used to track the status of the operation. |  -  |
**204** | No Content - the tenant was already in the desired state. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **opt_out_of_asa_health_metrics**
> CdoTransaction opt_out_of_asa_health_metrics()

Opt out of ASA Health Metrics

Allows a tenant to opt out of receiving health metrics for their ASAs.

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
    api_instance = scc_firewall_manager_sdk.DeviceHealthApi(api_client)

    try:
        # Opt out of ASA Health Metrics
        api_response = api_instance.opt_out_of_asa_health_metrics()
        print("The response of DeviceHealthApi->opt_out_of_asa_health_metrics:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceHealthApi->opt_out_of_asa_health_metrics: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Security Cloud Control Transaction object that can be used to track the status of the operation. |  -  |
**204** | No Content - the tenant was already in the desired state. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

