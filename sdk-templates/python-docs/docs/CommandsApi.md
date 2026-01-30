# scc_firewall_manager_sdk.CommandsApi

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
[**execute_device_command**](CommandsApi.md#execute_device_command) | **POST** /v1/fmc/gateway/command | Send a direct request to a registered device in Security Cloud Control
[**get_device_command_response**](CommandsApi.md#get_device_command_response) | **GET** /v1/fmc/gateway/command/{requestId} | Retrieve the response for a previously sent command
[**get_device_commands**](CommandsApi.md#get_device_commands) | **GET** /v1/fmc/gateway/listcommands | Get supported endpoint patterns


# **execute_device_command**
> execute_device_command(device_gateway_api_request, var_async=var_async)

Send a direct request to a registered device in Security Cloud Control

Send a direct command to the specified device.

### Example


```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device_gateway_api_request import DeviceGatewayApiRequest
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
    api_instance = scc_firewall_manager_sdk.CommandsApi(api_client)
    device_gateway_api_request = scc_firewall_manager_sdk.DeviceGatewayApiRequest() # DeviceGatewayApiRequest | 
    var_async = 'var_async_example' # str | By default, the API waits up to 70 seconds for a device response. Append the 'async' flag to return immediately without waiting for the response. (optional)

    try:
        # Send a direct request to a registered device in Security Cloud Control
        api_instance.execute_device_command(device_gateway_api_request, var_async=var_async)
    except Exception as e:
        print("Exception when calling CommandsApi->execute_device_command: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_gateway_api_request** | [**DeviceGatewayApiRequest**](DeviceGatewayApiRequest.md)|  | 
 **var_async** | **str**| By default, the API waits up to 70 seconds for a device response. Append the &#39;async&#39; flag to return immediately without waiting for the response. | [optional] 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The request was successfully processed, and the device command was executed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device_command_response**
> get_device_command_response(request_id)

Retrieve the response for a previously sent command

Get the response from a device for a command previously issued.

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
    api_instance = scc_firewall_manager_sdk.CommandsApi(api_client)
    request_id = 'request_id_example' # str | The unique identifier of the request, represented as a UUID, returned from a previous POST request to the /command endpoint.

    try:
        # Retrieve the response for a previously sent command
        api_instance.get_device_command_response(request_id)
    except Exception as e:
        print("Exception when calling CommandsApi->get_device_command_response: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request_id** | **str**| The unique identifier of the request, represented as a UUID, returned from a previous POST request to the /command endpoint. | 

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
**200** | The device responded. The response may indicate success or contain an error returned by the device. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device_commands**
> CommandResponse get_device_commands()

Get supported endpoint patterns

Retrieve a list of supported device command endpoint patterns.

### Example


```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.command_response import CommandResponse
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
    api_instance = scc_firewall_manager_sdk.CommandsApi(api_client)

    try:
        # Get supported endpoint patterns
        api_response = api_instance.get_device_commands()
        print("The response of CommandsApi->get_device_commands:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CommandsApi->get_device_commands: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**CommandResponse**](CommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The list of supported command patterns. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

