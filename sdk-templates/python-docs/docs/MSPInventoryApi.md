# scc_firewall_manager_sdk.MSPInventoryApi

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
[**export_msp_managed_cloud_services**](MSPInventoryApi.md#export_msp_managed_cloud_services) | **POST** /v1/msp/inventory/services/export | Export MSP-managed Cloud Services
[**export_msp_managed_device_managers**](MSPInventoryApi.md#export_msp_managed_device_managers) | **POST** /v1/msp/inventory/managers/export | Export MSP-managed Device Managers
[**export_msp_managed_devices**](MSPInventoryApi.md#export_msp_managed_devices) | **POST** /v1/msp/inventory/devices/export | Export MSP-managed Devices
[**export_msp_managed_templates**](MSPInventoryApi.md#export_msp_managed_templates) | **POST** /v1/msp/inventory/templates/export | Export MSP-managed Templates
[**get_msp_managed_cloud_service**](MSPInventoryApi.md#get_msp_managed_cloud_service) | **GET** /v1/msp/inventory/services/{cloudServiceUid} | Get MSP-managed cloud service by UID
[**get_msp_managed_cloud_services**](MSPInventoryApi.md#get_msp_managed_cloud_services) | **GET** /v1/msp/inventory/services | Get MSP-managed cloud services
[**get_msp_managed_cloud_services_attribute_values**](MSPInventoryApi.md#get_msp_managed_cloud_services_attribute_values) | **GET** /v1/msp/inventory/services/attribute-values | Get distinct attribute values for MSP-managed cloud services
[**get_msp_managed_device**](MSPInventoryApi.md#get_msp_managed_device) | **GET** /v1/msp/inventory/devices/{deviceUid} | Get MSP-managed device by UID
[**get_msp_managed_device_manager**](MSPInventoryApi.md#get_msp_managed_device_manager) | **GET** /v1/msp/inventory/managers/{deviceManagerUid} | Get MSP-managed device manager by UID
[**get_msp_managed_device_managers**](MSPInventoryApi.md#get_msp_managed_device_managers) | **GET** /v1/msp/inventory/managers | Get MSP-managed device managers
[**get_msp_managed_device_managers_attribute_values**](MSPInventoryApi.md#get_msp_managed_device_managers_attribute_values) | **GET** /v1/msp/inventory/managers/attribute-values | Get distinct attribute values for MSP-managed device managers
[**get_msp_managed_devices**](MSPInventoryApi.md#get_msp_managed_devices) | **GET** /v1/msp/inventory/devices | Get MSP-managed devices
[**get_msp_managed_devices_attribute_values**](MSPInventoryApi.md#get_msp_managed_devices_attribute_values) | **GET** /v1/msp/inventory/devices/attribute-values | Get distinct attribute values for MSP-managed devices
[**get_msp_managed_template**](MSPInventoryApi.md#get_msp_managed_template) | **GET** /v1/msp/inventory/templates/{templateUid} | Get MSP-managed template by UID
[**get_msp_managed_template_attribute_values**](MSPInventoryApi.md#get_msp_managed_template_attribute_values) | **GET** /v1/msp/inventory/templates/attribute-values | Get distinct attribute values for MSP-managed templates
[**get_msp_managed_templates**](MSPInventoryApi.md#get_msp_managed_templates) | **GET** /v1/msp/inventory/templates | Get MSP-managed templates
[**modify_msp_managed_devices**](MSPInventoryApi.md#modify_msp_managed_devices) | **PATCH** /v1/msp/inventory/devices/bulk | Modify MSP-managed devices


# **export_msp_managed_cloud_services**
> CdoTransaction export_msp_managed_cloud_services(msp_export_input)

Export MSP-managed Cloud Services

This is an asynchronous operation to export MSP-managed cloud services in CSV format. Once complete, the file can be downloaded using a presigned AWS S3 URL specified in the entityUrl field of the transaction that expires in 1 hour.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.msp_export_input import MspExportInput
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)
    msp_export_input = scc_firewall_manager_sdk.MspExportInput() # MspExportInput | 

    try:
        # Export MSP-managed Cloud Services
        api_response = api_instance.export_msp_managed_cloud_services(msp_export_input)
        print("The response of MSPInventoryApi->export_msp_managed_cloud_services:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->export_msp_managed_cloud_services: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msp_export_input** | [**MspExportInput**](MspExportInput.md)|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Security Cloud Control Transaction object that can be used to track the status of the export. Once complete, the &lt;code&gt;entityUrl&lt;/code&gt; field of the transaction will contain a presigned AWS S3 URL, valid for 1 hour, to download the exported file. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **export_msp_managed_device_managers**
> CdoTransaction export_msp_managed_device_managers(msp_export_input)

Export MSP-managed Device Managers

This is an asynchronous operation to export MSP-managed device managers in CSV format. Once complete, the file can be downloaded using a presigned AWS S3 URL specified in the entityUrl field of the transaction that expires in 1 hour.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.msp_export_input import MspExportInput
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)
    msp_export_input = scc_firewall_manager_sdk.MspExportInput() # MspExportInput | 

    try:
        # Export MSP-managed Device Managers
        api_response = api_instance.export_msp_managed_device_managers(msp_export_input)
        print("The response of MSPInventoryApi->export_msp_managed_device_managers:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->export_msp_managed_device_managers: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msp_export_input** | [**MspExportInput**](MspExportInput.md)|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Security Cloud Control Transaction object that can be used to track the status of the export. Once complete, the &lt;code&gt;entityUrl&lt;/code&gt; field of the transaction will contain a presigned AWS S3 URL, valid for 1 hour, to download the exported file. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **export_msp_managed_devices**
> CdoTransaction export_msp_managed_devices(msp_export_input)

Export MSP-managed Devices

This is an asynchronous operation to export MSP-managed devices in CSV format. Once complete, the file can be downloaded using a presigned AWS S3 URL specified in the entityUrl field of the transaction that expires in 1 hour.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.msp_export_input import MspExportInput
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)
    msp_export_input = scc_firewall_manager_sdk.MspExportInput() # MspExportInput | 

    try:
        # Export MSP-managed Devices
        api_response = api_instance.export_msp_managed_devices(msp_export_input)
        print("The response of MSPInventoryApi->export_msp_managed_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->export_msp_managed_devices: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msp_export_input** | [**MspExportInput**](MspExportInput.md)|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Security Cloud Control Transaction object that can be used to track the status of the export. Once complete, the &lt;code&gt;entityUrl&lt;/code&gt; field of the transaction will contain a presigned AWS S3 URL, valid for 1 hour, to download the exported file. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **export_msp_managed_templates**
> CdoTransaction export_msp_managed_templates(msp_export_input)

Export MSP-managed Templates

This is an asynchronous operation to export MSP-managed templates in CSV format. Once complete, the file can be downloaded using a presigned AWS S3 URL specified in the entityUrl field of the transaction that expires in 1 hour.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.msp_export_input import MspExportInput
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)
    msp_export_input = scc_firewall_manager_sdk.MspExportInput() # MspExportInput | 

    try:
        # Export MSP-managed Templates
        api_response = api_instance.export_msp_managed_templates(msp_export_input)
        print("The response of MSPInventoryApi->export_msp_managed_templates:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->export_msp_managed_templates: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msp_export_input** | [**MspExportInput**](MspExportInput.md)|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Security Cloud Control Transaction object that can be used to track the status of the export. Once complete, the &lt;code&gt;entityUrl&lt;/code&gt; field of the transaction will contain a presigned AWS S3 URL, valid for 1 hour, to download the exported file. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_cloud_service**
> MspManagedCloudService get_msp_managed_cloud_service(cloud_service_uid)

Get MSP-managed cloud service by UID

Get the cloud service managed by the MSP portal using its UID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_managed_cloud_service import MspManagedCloudService
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)
    cloud_service_uid = 'cloud_service_uid_example' # str | The unique identifier of MSP-managed cloud service.

    try:
        # Get MSP-managed cloud service by UID
        api_response = api_instance.get_msp_managed_cloud_service(cloud_service_uid)
        print("The response of MSPInventoryApi->get_msp_managed_cloud_service:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->get_msp_managed_cloud_service: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud_service_uid** | **str**| The unique identifier of MSP-managed cloud service. | 

### Return type

[**MspManagedCloudService**](MspManagedCloudService.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | MSP-managed cloud service |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_cloud_services**
> MspManagedCloudServicePage get_msp_managed_cloud_services(limit=limit, offset=offset, q=q, sort=sort)

Get MSP-managed cloud services

Get cloud services, across all tenants, managed by the MSP portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_managed_cloud_service_page import MspManagedCloudServicePage
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get MSP-managed cloud services
        api_response = api_instance.get_msp_managed_cloud_services(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPInventoryApi->get_msp_managed_cloud_services:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->get_msp_managed_cloud_services: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**MspManagedCloudServicePage**](MspManagedCloudServicePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of MSP-managed cloud services |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_cloud_services_attribute_values**
> MspManagedCloudServiceDistinctAttributeValues get_msp_managed_cloud_services_attribute_values()

Get distinct attribute values for MSP-managed cloud services

Get distinct values for multiple fields in the MSP-managed cloud services across all tenants managed by the MSP portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_managed_cloud_service_distinct_attribute_values import MspManagedCloudServiceDistinctAttributeValues
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)

    try:
        # Get distinct attribute values for MSP-managed cloud services
        api_response = api_instance.get_msp_managed_cloud_services_attribute_values()
        print("The response of MSPInventoryApi->get_msp_managed_cloud_services_attribute_values:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->get_msp_managed_cloud_services_attribute_values: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**MspManagedCloudServiceDistinctAttributeValues**](MspManagedCloudServiceDistinctAttributeValues.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Distinct attribute values for MSP-managed cloud services |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_device**
> MspManagedDevice get_msp_managed_device(device_uid)

Get MSP-managed device by UID

Get the device managed by the MSP portal using its UID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_managed_device import MspManagedDevice
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)
    device_uid = 'device_uid_example' # str | The unique identifier of MSP-managed device.

    try:
        # Get MSP-managed device by UID
        api_response = api_instance.get_msp_managed_device(device_uid)
        print("The response of MSPInventoryApi->get_msp_managed_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->get_msp_managed_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier of MSP-managed device. | 

### Return type

[**MspManagedDevice**](MspManagedDevice.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | MSP-managed device |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_device_manager**
> MspManagedDeviceManager get_msp_managed_device_manager(device_manager_uid)

Get MSP-managed device manager by UID

Get the device manager managed by the MSP portal using its UID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_managed_device_manager import MspManagedDeviceManager
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)
    device_manager_uid = 'device_manager_uid_example' # str | The unique identifier of MSP-managed device manager.

    try:
        # Get MSP-managed device manager by UID
        api_response = api_instance.get_msp_managed_device_manager(device_manager_uid)
        print("The response of MSPInventoryApi->get_msp_managed_device_manager:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->get_msp_managed_device_manager: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_manager_uid** | **str**| The unique identifier of MSP-managed device manager. | 

### Return type

[**MspManagedDeviceManager**](MspManagedDeviceManager.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | MSP-managed device manager |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_device_managers**
> MspManagedDeviceManagersPage get_msp_managed_device_managers(limit=limit, offset=offset, q=q, sort=sort)

Get MSP-managed device managers

Get device managers, across all tenants, managed by the MSP portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_managed_device_managers_page import MspManagedDeviceManagersPage
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get MSP-managed device managers
        api_response = api_instance.get_msp_managed_device_managers(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPInventoryApi->get_msp_managed_device_managers:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->get_msp_managed_device_managers: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**MspManagedDeviceManagersPage**](MspManagedDeviceManagersPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of MSP-managed device managers |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_device_managers_attribute_values**
> MspManagedDeviceManagerDistinctAttributeValues get_msp_managed_device_managers_attribute_values()

Get distinct attribute values for MSP-managed device managers

Get distinct values for multiple fields in the MSP-managed device managers across all tenants managed by the MSP portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_managed_device_manager_distinct_attribute_values import MspManagedDeviceManagerDistinctAttributeValues
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)

    try:
        # Get distinct attribute values for MSP-managed device managers
        api_response = api_instance.get_msp_managed_device_managers_attribute_values()
        print("The response of MSPInventoryApi->get_msp_managed_device_managers_attribute_values:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->get_msp_managed_device_managers_attribute_values: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**MspManagedDeviceManagerDistinctAttributeValues**](MspManagedDeviceManagerDistinctAttributeValues.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Distinct attribute values for MSP-managed device managers |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_devices**
> MspManagedDevicePage get_msp_managed_devices(limit=limit, offset=offset, q=q, sort=sort)

Get MSP-managed devices

Get devices, across all tenants, managed by the MSP portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_managed_device_page import MspManagedDevicePage
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get MSP-managed devices
        api_response = api_instance.get_msp_managed_devices(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPInventoryApi->get_msp_managed_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->get_msp_managed_devices: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**MspManagedDevicePage**](MspManagedDevicePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of MSP-managed devices |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_devices_attribute_values**
> MspManagedDeviceDistinctAttributeValues get_msp_managed_devices_attribute_values()

Get distinct attribute values for MSP-managed devices

Get distinct values for multiple fields in the MSP-managed devices across all tenants managed by the MSP portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_managed_device_distinct_attribute_values import MspManagedDeviceDistinctAttributeValues
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)

    try:
        # Get distinct attribute values for MSP-managed devices
        api_response = api_instance.get_msp_managed_devices_attribute_values()
        print("The response of MSPInventoryApi->get_msp_managed_devices_attribute_values:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->get_msp_managed_devices_attribute_values: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**MspManagedDeviceDistinctAttributeValues**](MspManagedDeviceDistinctAttributeValues.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Distinct attribute values for MSP-managed devices |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_template**
> MspManagedTemplate get_msp_managed_template(template_uid)

Get MSP-managed template by UID

Get the template managed by the MSP portal using its UID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_managed_template import MspManagedTemplate
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)
    template_uid = 'template_uid_example' # str | The unique identifier of MSP-managed template.

    try:
        # Get MSP-managed template by UID
        api_response = api_instance.get_msp_managed_template(template_uid)
        print("The response of MSPInventoryApi->get_msp_managed_template:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->get_msp_managed_template: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **template_uid** | **str**| The unique identifier of MSP-managed template. | 

### Return type

[**MspManagedTemplate**](MspManagedTemplate.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | MSP-managed template |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_template_attribute_values**
> MspManagedTemplateDistinctAttributeValues get_msp_managed_template_attribute_values()

Get distinct attribute values for MSP-managed templates

Get distinct values for multiple fields in the MSP-managed templates across all tenants managed by the MSP portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_managed_template_distinct_attribute_values import MspManagedTemplateDistinctAttributeValues
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)

    try:
        # Get distinct attribute values for MSP-managed templates
        api_response = api_instance.get_msp_managed_template_attribute_values()
        print("The response of MSPInventoryApi->get_msp_managed_template_attribute_values:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->get_msp_managed_template_attribute_values: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**MspManagedTemplateDistinctAttributeValues**](MspManagedTemplateDistinctAttributeValues.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Distinct attribute values for MSP-managed templates |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_managed_templates**
> MspManagedTemplatesPage get_msp_managed_templates(limit=limit, offset=offset, q=q, sort=sort)

Get MSP-managed templates

Get templates, across all tenants, managed by the MSP portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_managed_templates_page import MspManagedTemplatesPage
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get MSP-managed templates
        api_response = api_instance.get_msp_managed_templates(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPInventoryApi->get_msp_managed_templates:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->get_msp_managed_templates: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**MspManagedTemplatesPage**](MspManagedTemplatesPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of MSP-managed templates |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_msp_managed_devices**
> CdoTransaction modify_msp_managed_devices(msp_bulk_device_update_input)

Modify MSP-managed devices

Modify a list of MSP-managed devices across multiple tenants, across multiple regions.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.msp_bulk_device_update_input import MspBulkDeviceUpdateInput
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
    api_instance = scc_firewall_manager_sdk.MSPInventoryApi(api_client)
    msp_bulk_device_update_input = scc_firewall_manager_sdk.MspBulkDeviceUpdateInput() # MspBulkDeviceUpdateInput | 

    try:
        # Modify MSP-managed devices
        api_response = api_instance.modify_msp_managed_devices(msp_bulk_device_update_input)
        print("The response of MSPInventoryApi->modify_msp_managed_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPInventoryApi->modify_msp_managed_devices: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msp_bulk_device_update_input** | [**MspBulkDeviceUpdateInput**](MspBulkDeviceUpdateInput.md)|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | Security Cloud Control Transaction object that can be used to track the status of the operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

