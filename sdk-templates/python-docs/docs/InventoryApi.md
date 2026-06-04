# scc_firewall_manager_sdk.InventoryApi

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
[**calculate_ftd_device_pending_changes**](InventoryApi.md#calculate_ftd_device_pending_changes) | **POST** /v1/inventory/devices/ftds/{deviceUid}/changes/pending | Calculate pending changes on a cdFMC-managed FTD
[**get_ftd_device_pending_changes**](InventoryApi.md#get_ftd_device_pending_changes) | **GET** /v1/inventory/devices/ftds/{deviceUid}/changes/pending | Get pending changes on a cdFMC-managed FTD


# **calculate_ftd_device_pending_changes**
> CdoTransaction calculate_ftd_device_pending_changes(device_uid)

Calculate pending changes on a cdFMC-managed FTD

This is an asynchronous operation to calculate the pending changes on a cdFMC-managed FTD. Note: if there is no deployment baseline available to compare the current state of the device against, an empty list will be returned.

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
    api_instance = scc_firewall_manager_sdk.InventoryApi(api_client)
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the cdFMC managed FTD device in Security Cloud Control.

    try:
        # Calculate pending changes on a cdFMC-managed FTD
        api_response = api_instance.calculate_ftd_device_pending_changes(device_uid)
        print("The response of InventoryApi->calculate_ftd_device_pending_changes:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->calculate_ftd_device_pending_changes: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the cdFMC managed FTD device in Security Cloud Control. | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the calculation operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_ftd_device_pending_changes**
> FtdChangeItemDto get_ftd_device_pending_changes(device_uid)

Get pending changes on a cdFMC-managed FTD

Get the pending changes on a cdFMC-managed FTD. Note 1: if there is no deployment baseline available to compare the current state of the device against, an empty list will be returned. Note 2: This is not a paginated endpoint.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.ftd_change_item_dto import FtdChangeItemDto
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
    api_instance = scc_firewall_manager_sdk.InventoryApi(api_client)
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the cdFMC managed FTD device in Security Cloud Control.

    try:
        # Get pending changes on a cdFMC-managed FTD
        api_response = api_instance.get_ftd_device_pending_changes(device_uid)
        print("The response of InventoryApi->get_ftd_device_pending_changes:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->get_ftd_device_pending_changes: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the cdFMC managed FTD device in Security Cloud Control. | 

### Return type

[**FtdChangeItemDto**](FtdChangeItemDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of pending changes on a cdFMC-managed FTD |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

