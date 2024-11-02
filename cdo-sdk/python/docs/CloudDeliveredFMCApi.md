# cdo_sdk_python.CloudDeliveredFMCApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**proxy_fmc_request**](CloudDeliveredFMCApi.md#proxy_fmc_request) | **GET** /v1/cdfmc/** | Proxy Request To Cloud-Delivered FMC
[**proxy_fmc_request1**](CloudDeliveredFMCApi.md#proxy_fmc_request1) | **POST** /v1/cdfmc/** | Proxy Request To Cloud-Delivered FMC
[**proxy_fmc_request2**](CloudDeliveredFMCApi.md#proxy_fmc_request2) | **PUT** /v1/cdfmc/** | Proxy Request To Cloud-Delivered FMC
[**proxy_fmc_request3**](CloudDeliveredFMCApi.md#proxy_fmc_request3) | **PATCH** /v1/cdfmc/** | Proxy Request To Cloud-Delivered FMC
[**proxy_fmc_request4**](CloudDeliveredFMCApi.md#proxy_fmc_request4) | **DELETE** /v1/cdfmc/** | Proxy Request To Cloud-Delivered FMC


# **proxy_fmc_request**
> proxy_fmc_request()

Proxy Request To Cloud-Delivered FMC

This endpoint proxies the request to the Cloud-Delivered FMC (cdFMC). Refer to the <a href=\"https://www.cisco.com/c/en/us/td/docs/security/firepower/730/Rapid-Release/API/CDO/cloud_delivered_firewall_management_center_rest_api_quick_start_guide/About_The_API_Explorer.html\">cdFMC API Explorer in CDO</a> for the list of available endpoints. You can append the relative API paths provided in the cdFMC API documentation to the `/v1/cdfmc` URL to make requests to the cdFMC in your CDO tenant. <b>Note:</b> These endpoints will return 404 if the SCC tenant does not have a cdFMC provisioned.

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
    api_instance = cdo_sdk_python.CloudDeliveredFMCApi(api_client)

    try:
        # Proxy Request To Cloud-Delivered FMC
        api_instance.proxy_fmc_request()
    except Exception as e:
        print("Exception when calling CloudDeliveredFMCApi->proxy_fmc_request: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

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
**401** | Request not authorized. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **proxy_fmc_request1**
> proxy_fmc_request1()

Proxy Request To Cloud-Delivered FMC

This endpoint proxies the request to the Cloud-Delivered FMC (cdFMC). Refer to the <a href=\"https://www.cisco.com/c/en/us/td/docs/security/firepower/730/Rapid-Release/API/CDO/cloud_delivered_firewall_management_center_rest_api_quick_start_guide/About_The_API_Explorer.html\">cdFMC API Explorer in CDO</a> for the list of available endpoints. You can append the relative API paths provided in the cdFMC API documentation to the `/v1/cdfmc` URL to make requests to the cdFMC in your CDO tenant. <b>Note:</b> These endpoints will return 404 if the SCC tenant does not have a cdFMC provisioned.

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
    api_instance = cdo_sdk_python.CloudDeliveredFMCApi(api_client)

    try:
        # Proxy Request To Cloud-Delivered FMC
        api_instance.proxy_fmc_request1()
    except Exception as e:
        print("Exception when calling CloudDeliveredFMCApi->proxy_fmc_request1: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

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
**401** | Request not authorized. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **proxy_fmc_request2**
> proxy_fmc_request2()

Proxy Request To Cloud-Delivered FMC

This endpoint proxies the request to the Cloud-Delivered FMC (cdFMC). Refer to the <a href=\"https://www.cisco.com/c/en/us/td/docs/security/firepower/730/Rapid-Release/API/CDO/cloud_delivered_firewall_management_center_rest_api_quick_start_guide/About_The_API_Explorer.html\">cdFMC API Explorer in CDO</a> for the list of available endpoints. You can append the relative API paths provided in the cdFMC API documentation to the `/v1/cdfmc` URL to make requests to the cdFMC in your CDO tenant. <b>Note:</b> These endpoints will return 404 if the SCC tenant does not have a cdFMC provisioned.

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
    api_instance = cdo_sdk_python.CloudDeliveredFMCApi(api_client)

    try:
        # Proxy Request To Cloud-Delivered FMC
        api_instance.proxy_fmc_request2()
    except Exception as e:
        print("Exception when calling CloudDeliveredFMCApi->proxy_fmc_request2: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

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
**401** | Request not authorized. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **proxy_fmc_request3**
> proxy_fmc_request3()

Proxy Request To Cloud-Delivered FMC

This endpoint proxies the request to the Cloud-Delivered FMC (cdFMC). Refer to the <a href=\"https://www.cisco.com/c/en/us/td/docs/security/firepower/730/Rapid-Release/API/CDO/cloud_delivered_firewall_management_center_rest_api_quick_start_guide/About_The_API_Explorer.html\">cdFMC API Explorer in CDO</a> for the list of available endpoints. You can append the relative API paths provided in the cdFMC API documentation to the `/v1/cdfmc` URL to make requests to the cdFMC in your CDO tenant. <b>Note:</b> These endpoints will return 404 if the SCC tenant does not have a cdFMC provisioned.

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
    api_instance = cdo_sdk_python.CloudDeliveredFMCApi(api_client)

    try:
        # Proxy Request To Cloud-Delivered FMC
        api_instance.proxy_fmc_request3()
    except Exception as e:
        print("Exception when calling CloudDeliveredFMCApi->proxy_fmc_request3: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

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
**401** | Request not authorized. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **proxy_fmc_request4**
> proxy_fmc_request4()

Proxy Request To Cloud-Delivered FMC

This endpoint proxies the request to the Cloud-Delivered FMC (cdFMC). Refer to the <a href=\"https://www.cisco.com/c/en/us/td/docs/security/firepower/730/Rapid-Release/API/CDO/cloud_delivered_firewall_management_center_rest_api_quick_start_guide/About_The_API_Explorer.html\">cdFMC API Explorer in CDO</a> for the list of available endpoints. You can append the relative API paths provided in the cdFMC API documentation to the `/v1/cdfmc` URL to make requests to the cdFMC in your CDO tenant. <b>Note:</b> These endpoints will return 404 if the SCC tenant does not have a cdFMC provisioned.

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
    api_instance = cdo_sdk_python.CloudDeliveredFMCApi(api_client)

    try:
        # Proxy Request To Cloud-Delivered FMC
        api_instance.proxy_fmc_request4()
    except Exception as e:
        print("Exception when calling CloudDeliveredFMCApi->proxy_fmc_request4: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

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
**401** | Request not authorized. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

