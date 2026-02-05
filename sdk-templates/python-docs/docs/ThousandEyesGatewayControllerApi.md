# scc_firewall_manager_sdk.ThousandEyesGatewayControllerApi

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
[**get_ai_ops_status**](ThousandEyesGatewayControllerApi.md#get_ai_ops_status) | **GET** /aiops-status | 
[**get_applications**](ThousandEyesGatewayControllerApi.md#get_applications) | **GET** /applications | 
[**get_outages_by_date_range**](ThousandEyesGatewayControllerApi.md#get_outages_by_date_range) | **GET** /outages | 
[**handle_webhook**](ThousandEyesGatewayControllerApi.md#handle_webhook) | **POST** /webhook | 
[**update_ai_ops_status**](ThousandEyesGatewayControllerApi.md#update_ai_ops_status) | **PUT** /aiops-status | 
[**update_applications**](ThousandEyesGatewayControllerApi.md#update_applications) | **PUT** /applications | 
[**update_applications1**](ThousandEyesGatewayControllerApi.md#update_applications1) | **PATCH** /applications | 


# **get_ai_ops_status**
> ThousandEyesAiOpsStatusResponse get_ai_ops_status()



### Example


```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.thousand_eyes_ai_ops_status_response import ThousandEyesAiOpsStatusResponse
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
    api_instance = scc_firewall_manager_sdk.ThousandEyesGatewayControllerApi(api_client)

    try:
        api_response = api_instance.get_ai_ops_status()
        print("The response of ThousandEyesGatewayControllerApi->get_ai_ops_status:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThousandEyesGatewayControllerApi->get_ai_ops_status: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**ThousandEyesAiOpsStatusResponse**](ThousandEyesAiOpsStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_applications**
> ThousandEyesApplicationsResponse get_applications()



### Example


```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.thousand_eyes_applications_response import ThousandEyesApplicationsResponse
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
    api_instance = scc_firewall_manager_sdk.ThousandEyesGatewayControllerApi(api_client)

    try:
        api_response = api_instance.get_applications()
        print("The response of ThousandEyesGatewayControllerApi->get_applications:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThousandEyesGatewayControllerApi->get_applications: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**ThousandEyesApplicationsResponse**](ThousandEyesApplicationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_outages_by_date_range**
> ThousandEyesOutageQueryResponse get_outages_by_date_range(request)



### Example


```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.thousand_eyes_outage_query_request import ThousandEyesOutageQueryRequest
from scc_firewall_manager_sdk.models.thousand_eyes_outage_query_response import ThousandEyesOutageQueryResponse
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
    api_instance = scc_firewall_manager_sdk.ThousandEyesGatewayControllerApi(api_client)
    request = scc_firewall_manager_sdk.ThousandEyesOutageQueryRequest() # ThousandEyesOutageQueryRequest | 

    try:
        api_response = api_instance.get_outages_by_date_range(request)
        print("The response of ThousandEyesGatewayControllerApi->get_outages_by_date_range:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThousandEyesGatewayControllerApi->get_outages_by_date_range: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ThousandEyesOutageQueryRequest**](.md)|  | 

### Return type

[**ThousandEyesOutageQueryResponse**](ThousandEyesOutageQueryResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **handle_webhook**
> handle_webhook(body)



### Example


```python
import scc_firewall_manager_sdk
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
    api_instance = scc_firewall_manager_sdk.ThousandEyesGatewayControllerApi(api_client)
    body = 'body_example' # str | 

    try:
        api_instance.handle_webhook(body)
    except Exception as e:
        print("Exception when calling ThousandEyesGatewayControllerApi->handle_webhook: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **str**|  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_ai_ops_status**
> update_ai_ops_status(thousand_eyes_ai_ops_status_update_request)



### Example


```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.thousand_eyes_ai_ops_status_update_request import ThousandEyesAiOpsStatusUpdateRequest
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
    api_instance = scc_firewall_manager_sdk.ThousandEyesGatewayControllerApi(api_client)
    thousand_eyes_ai_ops_status_update_request = scc_firewall_manager_sdk.ThousandEyesAiOpsStatusUpdateRequest() # ThousandEyesAiOpsStatusUpdateRequest | 

    try:
        api_instance.update_ai_ops_status(thousand_eyes_ai_ops_status_update_request)
    except Exception as e:
        print("Exception when calling ThousandEyesGatewayControllerApi->update_ai_ops_status: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thousand_eyes_ai_ops_status_update_request** | [**ThousandEyesAiOpsStatusUpdateRequest**](ThousandEyesAiOpsStatusUpdateRequest.md)|  | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | No Content |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_applications**
> ThousandEyesApplicationsResponse update_applications(thousand_eyes_applications_update_request)



### Example


```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.thousand_eyes_applications_response import ThousandEyesApplicationsResponse
from scc_firewall_manager_sdk.models.thousand_eyes_applications_update_request import ThousandEyesApplicationsUpdateRequest
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
    api_instance = scc_firewall_manager_sdk.ThousandEyesGatewayControllerApi(api_client)
    thousand_eyes_applications_update_request = scc_firewall_manager_sdk.ThousandEyesApplicationsUpdateRequest() # ThousandEyesApplicationsUpdateRequest | 

    try:
        api_response = api_instance.update_applications(thousand_eyes_applications_update_request)
        print("The response of ThousandEyesGatewayControllerApi->update_applications:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThousandEyesGatewayControllerApi->update_applications: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thousand_eyes_applications_update_request** | [**ThousandEyesApplicationsUpdateRequest**](ThousandEyesApplicationsUpdateRequest.md)|  | 

### Return type

[**ThousandEyesApplicationsResponse**](ThousandEyesApplicationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_applications1**
> ThousandEyesApplicationsResponse update_applications1(thousand_eyes_applications_update_request)



### Example


```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.thousand_eyes_applications_response import ThousandEyesApplicationsResponse
from scc_firewall_manager_sdk.models.thousand_eyes_applications_update_request import ThousandEyesApplicationsUpdateRequest
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
    api_instance = scc_firewall_manager_sdk.ThousandEyesGatewayControllerApi(api_client)
    thousand_eyes_applications_update_request = scc_firewall_manager_sdk.ThousandEyesApplicationsUpdateRequest() # ThousandEyesApplicationsUpdateRequest | 

    try:
        api_response = api_instance.update_applications1(thousand_eyes_applications_update_request)
        print("The response of ThousandEyesGatewayControllerApi->update_applications1:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ThousandEyesGatewayControllerApi->update_applications1: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **thousand_eyes_applications_update_request** | [**ThousandEyesApplicationsUpdateRequest**](ThousandEyesApplicationsUpdateRequest.md)|  | 

### Return type

[**ThousandEyesApplicationsResponse**](ThousandEyesApplicationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

