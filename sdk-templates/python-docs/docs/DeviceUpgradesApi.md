# scc_firewall_manager_sdk.DeviceUpgradesApi

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
[**delete_device_upgrade_run**](DeviceUpgradesApi.md#delete_device_upgrade_run) | **DELETE** /v1/inventory/devices/upgrades/runs/{uid} | Delete Device Upgrade Run
[**get_asa_upgrade_versions**](DeviceUpgradesApi.md#get_asa_upgrade_versions) | **GET** /v1/inventory/devices/asas/{deviceUid}/upgrades/versions | Get compatible upgrade versions for an ASA
[**get_compatible_asa_versions_for_multiple_asas**](DeviceUpgradesApi.md#get_compatible_asa_versions_for_multiple_asas) | **GET** /v1/inventory/devices/asas/upgrades/versions | Get upgrade versions compatible with multiple ASAs
[**get_compatible_ftd_versions**](DeviceUpgradesApi.md#get_compatible_ftd_versions) | **GET** /v1/inventory/devices/ftds/{deviceUid}/upgrades/versions | Get compatible upgrade versions for an FTD
[**get_compatible_ftd_versions_for_multiple_ftds**](DeviceUpgradesApi.md#get_compatible_ftd_versions_for_multiple_ftds) | **GET** /v1/inventory/devices/ftds/upgrades/versions | Get upgrade versions compatible with multiple FTDs
[**get_device_upgrade_run**](DeviceUpgradesApi.md#get_device_upgrade_run) | **GET** /v1/inventory/devices/upgrades/runs/{upgradeRunUid} | Get Device Upgrade Run
[**get_device_upgrade_runs**](DeviceUpgradesApi.md#get_device_upgrade_runs) | **GET** /v1/inventory/devices/upgrades/runs | Get Device Upgrade Runs
[**modify_device_upgrade_run**](DeviceUpgradesApi.md#modify_device_upgrade_run) | **PATCH** /v1/inventory/devices/upgrades/runs/{upgradeRunUid} | Modify Device Upgrade Run
[**update_ftd_upgrade_packages_cache**](DeviceUpgradesApi.md#update_ftd_upgrade_packages_cache) | **PUT** /v1/inventory/devices/ftds/upgrades/packages/build-cache | Update cache of compatible upgrade packages for all FTDs
[**upgrade_asa_device**](DeviceUpgradesApi.md#upgrade_asa_device) | **POST** /v1/inventory/devices/asas/{deviceUid}/upgrades/trigger | Upgrade ASA device
[**upgrade_asa_devices**](DeviceUpgradesApi.md#upgrade_asa_devices) | **POST** /v1/inventory/devices/asas/upgrades/trigger | Upgrade multiple ASA devices
[**upgrade_ftd_device**](DeviceUpgradesApi.md#upgrade_ftd_device) | **POST** /v1/inventory/devices/ftds/{deviceUid}/upgrades/trigger | Upgrade FTD device
[**upgrade_ftd_devices**](DeviceUpgradesApi.md#upgrade_ftd_devices) | **POST** /v1/inventory/devices/ftds/upgrades/trigger | Upgrade multiple FTD devices


# **delete_device_upgrade_run**
> delete_device_upgrade_run(uid)

Delete Device Upgrade Run

Delete a Device Upgrade Run by UID in the SCC Firewall Manager Tenant. Warning: if you delete an upgrade run currently in progress, you may end up allowing multiple upgrades to be triggered on devices.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
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
    api_instance = scc_firewall_manager_sdk.DeviceUpgradesApi(api_client)
    uid = 'uid_example' # str | The unique identifier, represented as a UUID, of the FTD Device Upgrade Run in SCC Firewall Manager.

    try:
        # Delete Device Upgrade Run
        api_instance.delete_device_upgrade_run(uid)
    except Exception as e:
        print("Exception when calling DeviceUpgradesApi->delete_device_upgrade_run: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **str**| The unique identifier, represented as a UUID, of the FTD Device Upgrade Run in SCC Firewall Manager. | 

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
**204** | No Content |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_upgrade_versions**
> AsaCompatibleVersionsResponse get_asa_upgrade_versions(device_uid)

Get compatible upgrade versions for an ASA

Get a list of compatible upgrade versions for a given ASA device.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_compatible_versions_response import AsaCompatibleVersionsResponse
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
    api_instance = scc_firewall_manager_sdk.DeviceUpgradesApi(api_client)
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the device in Security Cloud Control.

    try:
        # Get compatible upgrade versions for an ASA
        api_response = api_instance.get_asa_upgrade_versions(device_uid)
        print("The response of DeviceUpgradesApi->get_asa_upgrade_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceUpgradesApi->get_asa_upgrade_versions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the device in Security Cloud Control. | 

### Return type

[**AsaCompatibleVersionsResponse**](AsaCompatibleVersionsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of compatible upgrade versions |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_compatible_asa_versions_for_multiple_asas**
> AsaCompatibilityVersion get_compatible_asa_versions_for_multiple_asas(device_uids)

Get upgrade versions compatible with multiple ASAs

Get a list of compatible upgrade versions for the specified ASA devices.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_compatibility_version import AsaCompatibilityVersion
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
    api_instance = scc_firewall_manager_sdk.DeviceUpgradesApi(api_client)
    device_uids = ['device_uids_example'] # List[str] | A list of unique identifiers, represented as UUIDs, of the devices in Security Cloud Control.

    try:
        # Get upgrade versions compatible with multiple ASAs
        api_response = api_instance.get_compatible_asa_versions_for_multiple_asas(device_uids)
        print("The response of DeviceUpgradesApi->get_compatible_asa_versions_for_multiple_asas:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceUpgradesApi->get_compatible_asa_versions_for_multiple_asas: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uids** | [**List[str]**](str.md)| A list of unique identifiers, represented as UUIDs, of the devices in Security Cloud Control. | 

### Return type

[**AsaCompatibilityVersion**](AsaCompatibilityVersion.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of compatible upgrade versions |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_compatible_ftd_versions**
> FtdVersionsResponse get_compatible_ftd_versions(device_uid)

Get compatible upgrade versions for an FTD

Get a list of compatible upgrade versions for a given FTD device. Note: this endpoint will only return versions that are directly downloadable to the FTD from the Cisco support site.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.ftd_versions_response import FtdVersionsResponse
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
    api_instance = scc_firewall_manager_sdk.DeviceUpgradesApi(api_client)
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the device in Security Cloud Control.

    try:
        # Get compatible upgrade versions for an FTD
        api_response = api_instance.get_compatible_ftd_versions(device_uid)
        print("The response of DeviceUpgradesApi->get_compatible_ftd_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceUpgradesApi->get_compatible_ftd_versions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the device in Security Cloud Control. | 

### Return type

[**FtdVersionsResponse**](FtdVersionsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of compatible upgrade versions |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**422** | Unprocessable entity. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_compatible_ftd_versions_for_multiple_ftds**
> FtdCompatibilityVersion get_compatible_ftd_versions_for_multiple_ftds(device_uids, x_calling_service=x_calling_service, x_request_metadata=x_request_metadata)

Get upgrade versions compatible with multiple FTDs

Get a list of compatible upgrade versions for the specified FTD devices. Note 1: this endpoint will only return versions that are directly downloadable to the FTD from the Cisco support site. Note 2: If compatible versions need computation, returns 202 with a transaction containing the entityUrl for tracking progress. Otherwise, returns 200 with the complete result.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.ftd_compatibility_version import FtdCompatibilityVersion
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
    api_instance = scc_firewall_manager_sdk.DeviceUpgradesApi(api_client)
    device_uids = ['device_uids_example'] # List[str] | A list of unique identifiers, represented as UUIDs, of the devices in Security Cloud Control.
    x_calling_service = 'x_calling_service_example' # str |  (optional)
    x_request_metadata = 'x_request_metadata_example' # str | Custom metadata to identify the request (optional)

    try:
        # Get upgrade versions compatible with multiple FTDs
        api_response = api_instance.get_compatible_ftd_versions_for_multiple_ftds(device_uids, x_calling_service=x_calling_service, x_request_metadata=x_request_metadata)
        print("The response of DeviceUpgradesApi->get_compatible_ftd_versions_for_multiple_ftds:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceUpgradesApi->get_compatible_ftd_versions_for_multiple_ftds: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uids** | [**List[str]**](str.md)| A list of unique identifiers, represented as UUIDs, of the devices in Security Cloud Control. | 
 **x_calling_service** | **str**|  | [optional] 
 **x_request_metadata** | **str**| Custom metadata to identify the request | [optional] 

### Return type

[**FtdCompatibilityVersion**](FtdCompatibilityVersion.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of compatible upgrade versions |  -  |
**202** | Returns a Security Cloud Control returns a Transaction object that can be used to track the progress of the operation to get the set of compatibility versions. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**422** | Unprocessable entity. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device_upgrade_run**
> UpgradeRunDto get_device_upgrade_run(upgrade_run_uid)

Get Device Upgrade Run

Get a Device Upgrade Run by UID in the SCC Firewall Manager Tenant. Each upgrade run represents a group of devices being upgraded, or staged for upgrades, together.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.upgrade_run_dto import UpgradeRunDto
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
    api_instance = scc_firewall_manager_sdk.DeviceUpgradesApi(api_client)
    upgrade_run_uid = 'upgrade_run_uid_example' # str | The unique identifier, represented as a UUID, of the FTD Device Upgrade Run in SCC Firewall Manager.

    try:
        # Get Device Upgrade Run
        api_response = api_instance.get_device_upgrade_run(upgrade_run_uid)
        print("The response of DeviceUpgradesApi->get_device_upgrade_run:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceUpgradesApi->get_device_upgrade_run: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upgrade_run_uid** | **str**| The unique identifier, represented as a UUID, of the FTD Device Upgrade Run in SCC Firewall Manager. | 

### Return type

[**UpgradeRunDto**](UpgradeRunDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The Device Upgrade Run object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device_upgrade_runs**
> FtdUpgradeRunDtoPage get_device_upgrade_runs(limit=limit, offset=offset, q=q, sort=sort)

Get Device Upgrade Runs

Get a list of device upgrade runs in the SCC Firewall Manager Tenant. Each upgrade run represents a group of devices being upgraded, or staged for upgrades, together.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.ftd_upgrade_run_dto_page import FtdUpgradeRunDtoPage
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
    api_instance = scc_firewall_manager_sdk.DeviceUpgradesApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Device Upgrade Runs
        api_response = api_instance.get_device_upgrade_runs(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of DeviceUpgradesApi->get_device_upgrade_runs:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceUpgradesApi->get_device_upgrade_runs: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**FtdUpgradeRunDtoPage**](FtdUpgradeRunDtoPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Device Upgrade Run objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_device_upgrade_run**
> FtdUpgradeRunDtoPage modify_device_upgrade_run(upgrade_run_uid, upgrade_run_modify_input)

Modify Device Upgrade Run

Modify a Device Upgrade Run by UID in the SCC Firewall Manager Tenant. Each upgrade run represents a group of devices being upgraded, or staged for upgrades, together.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.ftd_upgrade_run_dto_page import FtdUpgradeRunDtoPage
from scc_firewall_manager_sdk.models.upgrade_run_modify_input import UpgradeRunModifyInput
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
    api_instance = scc_firewall_manager_sdk.DeviceUpgradesApi(api_client)
    upgrade_run_uid = 'upgrade_run_uid_example' # str | The unique identifier, represented as a UUID, of the FTD Device Upgrade Run in SCC Firewall Manager.
    upgrade_run_modify_input = scc_firewall_manager_sdk.UpgradeRunModifyInput() # UpgradeRunModifyInput | 

    try:
        # Modify Device Upgrade Run
        api_response = api_instance.modify_device_upgrade_run(upgrade_run_uid, upgrade_run_modify_input)
        print("The response of DeviceUpgradesApi->modify_device_upgrade_run:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceUpgradesApi->modify_device_upgrade_run: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upgrade_run_uid** | **str**| The unique identifier, represented as a UUID, of the FTD Device Upgrade Run in SCC Firewall Manager. | 
 **upgrade_run_modify_input** | [**UpgradeRunModifyInput**](UpgradeRunModifyInput.md)|  | 

### Return type

[**FtdUpgradeRunDtoPage**](FtdUpgradeRunDtoPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The Device Upgrade Run object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **update_ftd_upgrade_packages_cache**
> CdoTransaction update_ftd_upgrade_packages_cache(x_calling_service=x_calling_service, x_request_metadata=x_request_metadata)

Update cache of compatible upgrade packages for all FTDs

Update cache of compatible upgrade packages for a all FTD devices.

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
    api_instance = scc_firewall_manager_sdk.DeviceUpgradesApi(api_client)
    x_calling_service = 'x_calling_service_example' # str |  (optional)
    x_request_metadata = 'x_request_metadata_example' # str | Custom metadata to identify the request (optional)

    try:
        # Update cache of compatible upgrade packages for all FTDs
        api_response = api_instance.update_ftd_upgrade_packages_cache(x_calling_service=x_calling_service, x_request_metadata=x_request_metadata)
        print("The response of DeviceUpgradesApi->update_ftd_upgrade_packages_cache:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceUpgradesApi->update_ftd_upgrade_packages_cache: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **x_calling_service** | **str**|  | [optional] 
 **x_request_metadata** | **str**| Custom metadata to identify the request | [optional] 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the cache building operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **upgrade_asa_device**
> CdoTransaction upgrade_asa_device(device_uid, upgrade_asa_device_input)

Upgrade ASA device

This asynchronous operation upgrades the ASA firmware and ASDM software versions on the device. Note: Newly detected or modified certificates will be automatically approved.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.upgrade_asa_device_input import UpgradeAsaDeviceInput
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
    api_instance = scc_firewall_manager_sdk.DeviceUpgradesApi(api_client)
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the device in Security Cloud Control.
    upgrade_asa_device_input = scc_firewall_manager_sdk.UpgradeAsaDeviceInput() # UpgradeAsaDeviceInput | 

    try:
        # Upgrade ASA device
        api_response = api_instance.upgrade_asa_device(device_uid, upgrade_asa_device_input)
        print("The response of DeviceUpgradesApi->upgrade_asa_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceUpgradesApi->upgrade_asa_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the device in Security Cloud Control. | 
 **upgrade_asa_device_input** | [**UpgradeAsaDeviceInput**](UpgradeAsaDeviceInput.md)|  | 

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

# **upgrade_asa_devices**
> CdoTransaction upgrade_asa_devices(upgrade_asa_devices_input)

Upgrade multiple ASA devices

Upgrade up to 50 ASA devices to the requested software and ASDM versions in a single request.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.upgrade_asa_devices_input import UpgradeAsaDevicesInput
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
    api_instance = scc_firewall_manager_sdk.DeviceUpgradesApi(api_client)
    upgrade_asa_devices_input = scc_firewall_manager_sdk.UpgradeAsaDevicesInput() # UpgradeAsaDevicesInput | 

    try:
        # Upgrade multiple ASA devices
        api_response = api_instance.upgrade_asa_devices(upgrade_asa_devices_input)
        print("The response of DeviceUpgradesApi->upgrade_asa_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceUpgradesApi->upgrade_asa_devices: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upgrade_asa_devices_input** | [**UpgradeAsaDevicesInput**](UpgradeAsaDevicesInput.md)|  | 

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

# **upgrade_ftd_device**
> CdoTransaction upgrade_ftd_device(device_uid, upgrade_ftd_device_input)

Upgrade FTD device

Upgrade FTD device using a specified upgrade package.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.upgrade_ftd_device_input import UpgradeFtdDeviceInput
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
    api_instance = scc_firewall_manager_sdk.DeviceUpgradesApi(api_client)
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the device in Security Cloud Control.
    upgrade_ftd_device_input = scc_firewall_manager_sdk.UpgradeFtdDeviceInput() # UpgradeFtdDeviceInput | 

    try:
        # Upgrade FTD device
        api_response = api_instance.upgrade_ftd_device(device_uid, upgrade_ftd_device_input)
        print("The response of DeviceUpgradesApi->upgrade_ftd_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceUpgradesApi->upgrade_ftd_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the device in Security Cloud Control. | 
 **upgrade_ftd_device_input** | [**UpgradeFtdDeviceInput**](UpgradeFtdDeviceInput.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the upgrade operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**422** | Unprocessable entity. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **upgrade_ftd_devices**
> CdoTransaction upgrade_ftd_devices(upgrade_ftd_devices_input, x_calling_service=x_calling_service, x_request_metadata=x_request_metadata)

Upgrade multiple FTD devices

Upgrade up to 50 FTD devices using a specified upgrade package. All of the FTDs in the list have to be compatible with the upgrade package.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.upgrade_ftd_devices_input import UpgradeFtdDevicesInput
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
    api_instance = scc_firewall_manager_sdk.DeviceUpgradesApi(api_client)
    upgrade_ftd_devices_input = scc_firewall_manager_sdk.UpgradeFtdDevicesInput() # UpgradeFtdDevicesInput | 
    x_calling_service = 'x_calling_service_example' # str |  (optional)
    x_request_metadata = 'x_request_metadata_example' # str | Custom metadata to identify the request (optional)

    try:
        # Upgrade multiple FTD devices
        api_response = api_instance.upgrade_ftd_devices(upgrade_ftd_devices_input, x_calling_service=x_calling_service, x_request_metadata=x_request_metadata)
        print("The response of DeviceUpgradesApi->upgrade_ftd_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling DeviceUpgradesApi->upgrade_ftd_devices: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upgrade_ftd_devices_input** | [**UpgradeFtdDevicesInput**](UpgradeFtdDevicesInput.md)|  | 
 **x_calling_service** | **str**|  | [optional] 
 **x_request_metadata** | **str**| Custom metadata to identify the request | [optional] 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the upgrade operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**409** | Conflict. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

