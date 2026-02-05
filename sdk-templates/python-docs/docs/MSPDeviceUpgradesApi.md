# scc_firewall_manager_sdk.MSPDeviceUpgradesApi

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
[**calculate_msp_ftd_compatible_upgrade_versions**](MSPDeviceUpgradesApi.md#calculate_msp_ftd_compatible_upgrade_versions) | **POST** /v1/msp/inventory/devices/ftds/upgrades/versions | Calculate compatible FTD upgrade versions
[**delete_msp_device_upgrade_run**](MSPDeviceUpgradesApi.md#delete_msp_device_upgrade_run) | **DELETE** /v1/msp/inventory/devices/upgrades/runs/{uid} | Delete MSP Device Upgrade Run
[**delete_msp_ftd_device_upgrade_run**](MSPDeviceUpgradesApi.md#delete_msp_ftd_device_upgrade_run) | **DELETE** /v1/msp/inventory/devices/ftds/upgrades/runs/{uid} | Delete MSP FTD Device Upgrade Run
[**get_msp_asa_compatible_upgrade_versions**](MSPDeviceUpgradesApi.md#get_msp_asa_compatible_upgrade_versions) | **GET** /v1/msp/inventory/devices/asas/upgrades/versions | Get compatible ASA upgrade versions
[**get_msp_device_upgrade_run**](MSPDeviceUpgradesApi.md#get_msp_device_upgrade_run) | **GET** /v1/msp/inventory/devices/upgrades/runs/{uid} | Get MSP Device Upgrade Run
[**get_msp_device_upgrade_runs**](MSPDeviceUpgradesApi.md#get_msp_device_upgrade_runs) | **GET** /v1/msp/inventory/devices/upgrades/runs | Get MSP Device Upgrade Runs
[**get_msp_device_upgrade_runs_attribute_values**](MSPDeviceUpgradesApi.md#get_msp_device_upgrade_runs_attribute_values) | **GET** /v1/msp/inventory/devices/upgrades/runs/attribute-values | Get distinct attribute values for MSP upgrade runs
[**get_msp_ftd_compatible_upgrade_versions**](MSPDeviceUpgradesApi.md#get_msp_ftd_compatible_upgrade_versions) | **GET** /v1/msp/inventory/devices/ftds/upgrades/versions/{upgradeVersionUid} | Get compatible FTD upgrade versions
[**get_msp_ftd_device_upgrade_run**](MSPDeviceUpgradesApi.md#get_msp_ftd_device_upgrade_run) | **GET** /v1/msp/inventory/devices/ftds/upgrades/runs/{mspUpgradeRunUid} | Get MSP FTD Device Upgrade Run
[**get_msp_ftd_device_upgrade_runs**](MSPDeviceUpgradesApi.md#get_msp_ftd_device_upgrade_runs) | **GET** /v1/msp/inventory/devices/ftds/upgrades/runs | Get MSP FTD Device Upgrade Runs
[**get_msp_ftd_upgrade_runs_attribute_values**](MSPDeviceUpgradesApi.md#get_msp_ftd_upgrade_runs_attribute_values) | **GET** /v1/msp/inventory/devices/ftds/upgrades/runs/attribute-values | Get distinct attribute values for MSP upgrade runs
[**upgrade_msp_managed_asa_devices**](MSPDeviceUpgradesApi.md#upgrade_msp_managed_asa_devices) | **POST** /v1/msp/inventory/devices/asas/upgrades/trigger | Upgrade multiple ASAs across multiple tenants
[**upgrade_msp_managed_ftd_devices**](MSPDeviceUpgradesApi.md#upgrade_msp_managed_ftd_devices) | **POST** /v1/msp/inventory/devices/ftds/upgrades/trigger | Upgrade multiple FTDs across multiple tenants


# **calculate_msp_ftd_compatible_upgrade_versions**
> CdoTransaction calculate_msp_ftd_compatible_upgrade_versions(msp_calculate_compatible_upgrade_versions_input)

Calculate compatible FTD upgrade versions

An asynchronous operation to calculate a list of compatible upgrade versions for a list of CdFMC-managed FTD devices across multiple managed tenants.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.msp_calculate_compatible_upgrade_versions_input import MspCalculateCompatibleUpgradeVersionsInput
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
    api_instance = scc_firewall_manager_sdk.MSPDeviceUpgradesApi(api_client)
    msp_calculate_compatible_upgrade_versions_input = scc_firewall_manager_sdk.MspCalculateCompatibleUpgradeVersionsInput() # MspCalculateCompatibleUpgradeVersionsInput | 

    try:
        # Calculate compatible FTD upgrade versions
        api_response = api_instance.calculate_msp_ftd_compatible_upgrade_versions(msp_calculate_compatible_upgrade_versions_input)
        print("The response of MSPDeviceUpgradesApi->calculate_msp_ftd_compatible_upgrade_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPDeviceUpgradesApi->calculate_msp_ftd_compatible_upgrade_versions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msp_calculate_compatible_upgrade_versions_input** | [**MspCalculateCompatibleUpgradeVersionsInput**](MspCalculateCompatibleUpgradeVersionsInput.md)|  | 

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
**202** | Returns a Transaction object to track the progress of the operation to get the set of compatibility versions. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**422** | Unprocessable entity. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_msp_device_upgrade_run**
> delete_msp_device_upgrade_run(uid)

Delete MSP Device Upgrade Run

Delete a MSP Device Upgrade Run, by UID, in the SCC Firewall Manager tenant. Caution: If you delete an upgrade run currently in progress, you may end up allowing multiple upgrades to be triggered on devices.

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
    api_instance = scc_firewall_manager_sdk.MSPDeviceUpgradesApi(api_client)
    uid = 'uid_example' # str | The unique identifier, represented as a UUID, of the Device Upgrade Run in SCC Firewall Manager.

    try:
        # Delete MSP Device Upgrade Run
        api_instance.delete_msp_device_upgrade_run(uid)
    except Exception as e:
        print("Exception when calling MSPDeviceUpgradesApi->delete_msp_device_upgrade_run: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **str**| The unique identifier, represented as a UUID, of the Device Upgrade Run in SCC Firewall Manager. | 

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
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_msp_ftd_device_upgrade_run**
> delete_msp_ftd_device_upgrade_run(uid)

Delete MSP FTD Device Upgrade Run

<strong>Deprecated</strong> - Use /v1/msp/inventory/devices/upgrades/runs/{uid} instead.  Delete a MSP Device Upgrade Run, by UID, in the SCC Firewall Manager tenant. Caution: If you delete an upgrade run currently in progress, you may end up allowing multiple upgrades to be triggered on devices. 

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
    api_instance = scc_firewall_manager_sdk.MSPDeviceUpgradesApi(api_client)
    uid = 'uid_example' # str | The unique identifier, represented as a UUID, of the FTD Device Upgrade Run in SCC Firewall Manager.

    try:
        # Delete MSP FTD Device Upgrade Run
        api_instance.delete_msp_ftd_device_upgrade_run(uid)
    except Exception as e:
        print("Exception when calling MSPDeviceUpgradesApi->delete_msp_ftd_device_upgrade_run: %s\n" % e)
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
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_asa_compatible_upgrade_versions**
> MspAsaCompatibleVersionsDto get_msp_asa_compatible_upgrade_versions(device_uids)

Get compatible ASA upgrade versions

Get a list of compatible upgrade versions for a list of ASA devices across multiple managed tenants.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_asa_compatible_versions_dto import MspAsaCompatibleVersionsDto
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
    api_instance = scc_firewall_manager_sdk.MSPDeviceUpgradesApi(api_client)
    device_uids = ['device_uids_example'] # List[str] | A list of unique identifiers, represented as UUIDs, of the devices in Security Cloud Control. Note: All the specified devices must be in tenants managed by the MSP portal.

    try:
        # Get compatible ASA upgrade versions
        api_response = api_instance.get_msp_asa_compatible_upgrade_versions(device_uids)
        print("The response of MSPDeviceUpgradesApi->get_msp_asa_compatible_upgrade_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPDeviceUpgradesApi->get_msp_asa_compatible_upgrade_versions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uids** | [**List[str]**](str.md)| A list of unique identifiers, represented as UUIDs, of the devices in Security Cloud Control. Note: All the specified devices must be in tenants managed by the MSP portal. | 

### Return type

[**MspAsaCompatibleVersionsDto**](MspAsaCompatibleVersionsDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Compatible upgrade versions for MSP managed ASA devices across multiple tenants. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**422** | Unprocessable entity. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_device_upgrade_run**
> MspUpgradeRunDto get_msp_device_upgrade_run(uid)

Get MSP Device Upgrade Run

Get a MSP Device Upgrade Run by UID in the SCC Firewall Manager Tenant. Each upgrade run represents a group of devices across multiple managed tenants being upgraded, or staged for upgrades, together.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_upgrade_run_dto import MspUpgradeRunDto
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
    api_instance = scc_firewall_manager_sdk.MSPDeviceUpgradesApi(api_client)
    uid = 'uid_example' # str | The unique identifier, represented as a UUID, of the Device Upgrade Run in SCC Firewall Manager.

    try:
        # Get MSP Device Upgrade Run
        api_response = api_instance.get_msp_device_upgrade_run(uid)
        print("The response of MSPDeviceUpgradesApi->get_msp_device_upgrade_run:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPDeviceUpgradesApi->get_msp_device_upgrade_run: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uid** | **str**| The unique identifier, represented as a UUID, of the Device Upgrade Run in SCC Firewall Manager. | 

### Return type

[**MspUpgradeRunDto**](MspUpgradeRunDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The MSP Device Upgrade Run object |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_device_upgrade_runs**
> MspUpgradeRunDtoPage get_msp_device_upgrade_runs(limit=limit, offset=offset, q=q, sort=sort)

Get MSP Device Upgrade Runs

Get a list of MSP device upgrade runs in the SCC Firewall Manager Tenant. Each upgrade run represents a group of devices across multiple managed tenants being upgraded, or staged for upgrades, together.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_upgrade_run_dto_page import MspUpgradeRunDtoPage
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
    api_instance = scc_firewall_manager_sdk.MSPDeviceUpgradesApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get MSP Device Upgrade Runs
        api_response = api_instance.get_msp_device_upgrade_runs(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPDeviceUpgradesApi->get_msp_device_upgrade_runs:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPDeviceUpgradesApi->get_msp_device_upgrade_runs: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**MspUpgradeRunDtoPage**](MspUpgradeRunDtoPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of MSP Device Upgrade Run objects |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_device_upgrade_runs_attribute_values**
> MspUpgradeRunsAttributeValues get_msp_device_upgrade_runs_attribute_values()

Get distinct attribute values for MSP upgrade runs

Get distinct attribute values for fields in the MSP upgrade runs.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_upgrade_runs_attribute_values import MspUpgradeRunsAttributeValues
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
    api_instance = scc_firewall_manager_sdk.MSPDeviceUpgradesApi(api_client)

    try:
        # Get distinct attribute values for MSP upgrade runs
        api_response = api_instance.get_msp_device_upgrade_runs_attribute_values()
        print("The response of MSPDeviceUpgradesApi->get_msp_device_upgrade_runs_attribute_values:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPDeviceUpgradesApi->get_msp_device_upgrade_runs_attribute_values: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**MspUpgradeRunsAttributeValues**](MspUpgradeRunsAttributeValues.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Distinct attribute values for MSP upgrade runs. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_ftd_compatible_upgrade_versions**
> UpgradeCompatibilityInfoDto get_msp_ftd_compatible_upgrade_versions(upgrade_version_uid)

Get compatible FTD upgrade versions

Get a Device Upgrade Run by UID in the SCC Firewall Manager Tenant. Each upgrade run represents a group of devices being upgraded, or staged for upgrades, together.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.upgrade_compatibility_info_dto import UpgradeCompatibilityInfoDto
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
    api_instance = scc_firewall_manager_sdk.MSPDeviceUpgradesApi(api_client)
    upgrade_version_uid = 'upgrade_version_uid_example' # str | Unique identifier, represented as a UUID, of the Device Upgrade Version object in SCC Firewall Manager.

    try:
        # Get compatible FTD upgrade versions
        api_response = api_instance.get_msp_ftd_compatible_upgrade_versions(upgrade_version_uid)
        print("The response of MSPDeviceUpgradesApi->get_msp_ftd_compatible_upgrade_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPDeviceUpgradesApi->get_msp_ftd_compatible_upgrade_versions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upgrade_version_uid** | **str**| Unique identifier, represented as a UUID, of the Device Upgrade Version object in SCC Firewall Manager. | 

### Return type

[**UpgradeCompatibilityInfoDto**](UpgradeCompatibilityInfoDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Compatible upgrade versions for MSP managed devices across multiple tenants. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_ftd_device_upgrade_run**
> MspUpgradeRunDto get_msp_ftd_device_upgrade_run(msp_upgrade_run_uid)

Get MSP FTD Device Upgrade Run

<strong>Deprecated</strong> - Use /v1/msp/inventory/devices/upgrades/runs/{mspUpgradeRunUid} instead.  Get a MSP Device Upgrade Run by UID in the SCC Firewall Manager Tenant. Each upgrade run represents a group of devices across multiple managed tenants being upgraded, or staged for upgrades, together. 

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_upgrade_run_dto import MspUpgradeRunDto
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
    api_instance = scc_firewall_manager_sdk.MSPDeviceUpgradesApi(api_client)
    msp_upgrade_run_uid = 'msp_upgrade_run_uid_example' # str | The unique identifier, represented as a UUID, of the FTD Device Upgrade Run in SCC Firewall Manager.

    try:
        # Get MSP FTD Device Upgrade Run
        api_response = api_instance.get_msp_ftd_device_upgrade_run(msp_upgrade_run_uid)
        print("The response of MSPDeviceUpgradesApi->get_msp_ftd_device_upgrade_run:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPDeviceUpgradesApi->get_msp_ftd_device_upgrade_run: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msp_upgrade_run_uid** | **str**| The unique identifier, represented as a UUID, of the FTD Device Upgrade Run in SCC Firewall Manager. | 

### Return type

[**MspUpgradeRunDto**](MspUpgradeRunDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The MSP Device Upgrade Run object |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_ftd_device_upgrade_runs**
> MspUpgradeRunDtoPage get_msp_ftd_device_upgrade_runs(limit=limit, offset=offset, q=q, sort=sort)

Get MSP FTD Device Upgrade Runs

<strong>Deprecated</strong> - Use /v1/msp/inventory/devices/upgrades/runs instead.  Get a list of MSP device upgrade runs in the SCC Firewall Manager Tenant. Each upgrade run represents a group of devices across multiple managed tenants being upgraded, or staged for upgrades, together. 

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_upgrade_run_dto_page import MspUpgradeRunDtoPage
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
    api_instance = scc_firewall_manager_sdk.MSPDeviceUpgradesApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get MSP FTD Device Upgrade Runs
        api_response = api_instance.get_msp_ftd_device_upgrade_runs(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPDeviceUpgradesApi->get_msp_ftd_device_upgrade_runs:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPDeviceUpgradesApi->get_msp_ftd_device_upgrade_runs: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**MspUpgradeRunDtoPage**](MspUpgradeRunDtoPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of MSP Device Upgrade Run objects |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_ftd_upgrade_runs_attribute_values**
> MspUpgradeRunsAttributeValues get_msp_ftd_upgrade_runs_attribute_values()

Get distinct attribute values for MSP upgrade runs

<strong>Deprecated</strong> - Use /v1/msp/inventory/devices/upgrades/runs/attribute-values instead.  Get distinct attribute values for fields in the MSP upgrade runs. 

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_upgrade_runs_attribute_values import MspUpgradeRunsAttributeValues
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
    api_instance = scc_firewall_manager_sdk.MSPDeviceUpgradesApi(api_client)

    try:
        # Get distinct attribute values for MSP upgrade runs
        api_response = api_instance.get_msp_ftd_upgrade_runs_attribute_values()
        print("The response of MSPDeviceUpgradesApi->get_msp_ftd_upgrade_runs_attribute_values:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPDeviceUpgradesApi->get_msp_ftd_upgrade_runs_attribute_values: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**MspUpgradeRunsAttributeValues**](MspUpgradeRunsAttributeValues.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Distinct attribute values for MSP upgrade runs. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **upgrade_msp_managed_asa_devices**
> CdoTransaction upgrade_msp_managed_asa_devices(msp_upgrade_asa_devices_input)

Upgrade multiple ASAs across multiple tenants

<p>Asynchronous operation to upgrade multiple ASA devices across managed tenants.</p> <p><strong>Notes:</strong></p> <ul> <li>Maximum of 50 ASA devices per request</li> </ul> <p><strong>Response:</strong></p> <ul> <li>Returns a transaction object to track upgrade operation progress</li> <li>Use the transaction ID or the entity UID to monitor the status of the upgrade</li> </ul>

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.msp_upgrade_asa_devices_input import MspUpgradeAsaDevicesInput
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
    api_instance = scc_firewall_manager_sdk.MSPDeviceUpgradesApi(api_client)
    msp_upgrade_asa_devices_input = scc_firewall_manager_sdk.MspUpgradeAsaDevicesInput() # MspUpgradeAsaDevicesInput | 

    try:
        # Upgrade multiple ASAs across multiple tenants
        api_response = api_instance.upgrade_msp_managed_asa_devices(msp_upgrade_asa_devices_input)
        print("The response of MSPDeviceUpgradesApi->upgrade_msp_managed_asa_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPDeviceUpgradesApi->upgrade_msp_managed_asa_devices: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msp_upgrade_asa_devices_input** | [**MspUpgradeAsaDevicesInput**](MspUpgradeAsaDevicesInput.md)|  | 

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
**202** | Returns a Transaction object that can be used to track the progress of the multi-tenant device upgrade operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **upgrade_msp_managed_ftd_devices**
> CdoTransaction upgrade_msp_managed_ftd_devices(msp_upgrade_ftd_devices_input)

Upgrade multiple FTDs across multiple tenants

<p>Asynchronous operation to upgrade multiple FTD devices across managed tenants.</p> <p><strong>Notes:</strong></p> <ul> <li>Maximum of 50 FTD devices per hardware model per managed tenant per request</li> </ul> <p><strong>Response:</strong></p> <ul> <li>Returns a transaction object to track upgrade operation progress</li> <li>Use the transaction ID or the entity UID to monitor the status of the upgrade</li> </ul>

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.msp_upgrade_ftd_devices_input import MspUpgradeFtdDevicesInput
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
    api_instance = scc_firewall_manager_sdk.MSPDeviceUpgradesApi(api_client)
    msp_upgrade_ftd_devices_input = scc_firewall_manager_sdk.MspUpgradeFtdDevicesInput() # MspUpgradeFtdDevicesInput | 

    try:
        # Upgrade multiple FTDs across multiple tenants
        api_response = api_instance.upgrade_msp_managed_ftd_devices(msp_upgrade_ftd_devices_input)
        print("The response of MSPDeviceUpgradesApi->upgrade_msp_managed_ftd_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPDeviceUpgradesApi->upgrade_msp_managed_ftd_devices: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msp_upgrade_ftd_devices_input** | [**MspUpgradeFtdDevicesInput**](MspUpgradeFtdDevicesInput.md)|  | 

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
**202** | Returns a Transaction object that can be used to track the progress of the multi-tenant device upgrade operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

