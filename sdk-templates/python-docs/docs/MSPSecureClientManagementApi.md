# scc_firewall_manager_sdk.MSPSecureClientManagementApi

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
[**calculate_msp_compatible_secure_client_versions**](MSPSecureClientManagementApi.md#calculate_msp_compatible_secure_client_versions) | **POST** /v1/msp/inventory/devices/asas/secure-client/upgrades/versions | Calculate compatible Secure Client upgrade versions across managed tenants
[**get_msp_compatible_secure_client_versions**](MSPSecureClientManagementApi.md#get_msp_compatible_secure_client_versions) | **GET** /v1/msp/inventory/devices/secure-client/upgrades/versions/{compatibilityInfoUid} | Get the result of a Secure Client compatibility calculation
[**get_msp_device_secure_client_packages**](MSPSecureClientManagementApi.md#get_msp_device_secure_client_packages) | **GET** /v1/msp/inventory/devices/{deviceUid}/packages | Get Secure Client packages installed on an MSP-managed device
[**get_msp_secure_client_upgrade_run**](MSPSecureClientManagementApi.md#get_msp_secure_client_upgrade_run) | **GET** /v1/msp/inventory/devices/secure-client/upgrades/runs/{upgradeRunUid} | Get MSP Secure Client Upgrade Run
[**get_msp_secure_client_upgrade_runs**](MSPSecureClientManagementApi.md#get_msp_secure_client_upgrade_runs) | **GET** /v1/msp/inventory/devices/secure-client/upgrades/runs | Get MSP Secure Client Upgrade Runs
[**get_msp_secure_client_upgrade_runs_attribute_values**](MSPSecureClientManagementApi.md#get_msp_secure_client_upgrade_runs_attribute_values) | **GET** /v1/msp/inventory/devices/secure-client/upgrades/runs/attribute-values | Get distinct attribute values for MSP Secure Client upgrade runs
[**upgrade_msp_managed_asa_secure_client_packages**](MSPSecureClientManagementApi.md#upgrade_msp_managed_asa_secure_client_packages) | **POST** /v1/msp/inventory/devices/asas/secure-client/upgrades | Upgrade Secure Client packages on ASAs across multiple tenants


# **calculate_msp_compatible_secure_client_versions**
> CdoTransaction calculate_msp_compatible_secure_client_versions(msp_calculate_compatible_secure_client_versions_input)

Calculate compatible Secure Client upgrade versions across managed tenants

An asynchronous operation to calculate the list of Secure Client versions compatible with the specified ASA devices across multiple managed tenants. The set returned is the intersection across tenants: a version is included only when every managed tenant offers it, and the platforms reported for that version are likewise the platforms offered by every managed tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.msp_calculate_compatible_secure_client_versions_input import MspCalculateCompatibleSecureClientVersionsInput
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
    api_instance = scc_firewall_manager_sdk.MSPSecureClientManagementApi(api_client)
    msp_calculate_compatible_secure_client_versions_input = scc_firewall_manager_sdk.MspCalculateCompatibleSecureClientVersionsInput() # MspCalculateCompatibleSecureClientVersionsInput | 

    try:
        # Calculate compatible Secure Client upgrade versions across managed tenants
        api_response = api_instance.calculate_msp_compatible_secure_client_versions(msp_calculate_compatible_secure_client_versions_input)
        print("The response of MSPSecureClientManagementApi->calculate_msp_compatible_secure_client_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPSecureClientManagementApi->calculate_msp_compatible_secure_client_versions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msp_calculate_compatible_secure_client_versions_input** | [**MspCalculateCompatibleSecureClientVersionsInput**](MspCalculateCompatibleSecureClientVersionsInput.md)|  | 

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
**202** | Returns a Transaction object to track the progress of the calculation. The Transaction&#39;s entity URL points at the GET endpoint that returns the aggregated result. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**422** | Unprocessable entity. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_compatible_secure_client_versions**
> MspSecureClientUpgradeCompatibilityInfoDto get_msp_compatible_secure_client_versions(compatibility_info_uid)

Get the result of a Secure Client compatibility calculation

Get the aggregated compatible Secure Client versions for an MSP cross-tenant calculation, identified by the compatibility-info UID returned in the entity URL of the calculation's Transaction.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_secure_client_upgrade_compatibility_info_dto import MspSecureClientUpgradeCompatibilityInfoDto
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
    api_instance = scc_firewall_manager_sdk.MSPSecureClientManagementApi(api_client)
    compatibility_info_uid = 'compatibility_info_uid_example' # str | Unique identifier, represented as a UUID, of the Secure Client compatibility-info object returned by the calculation endpoint.

    try:
        # Get the result of a Secure Client compatibility calculation
        api_response = api_instance.get_msp_compatible_secure_client_versions(compatibility_info_uid)
        print("The response of MSPSecureClientManagementApi->get_msp_compatible_secure_client_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPSecureClientManagementApi->get_msp_compatible_secure_client_versions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **compatibility_info_uid** | **str**| Unique identifier, represented as a UUID, of the Secure Client compatibility-info object returned by the calculation endpoint. | 

### Return type

[**MspSecureClientUpgradeCompatibilityInfoDto**](MspSecureClientUpgradeCompatibilityInfoDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Compatible Secure Client upgrade versions for the requested MSP-managed ASA devices, aggregated across managed tenants. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_device_secure_client_packages**
> MspSecureClientPackagesPage get_msp_device_secure_client_packages(device_uid, limit=limit, offset=offset, q=q, sort=sort)

Get Secure Client packages installed on an MSP-managed device

Get the Secure Client webdeploy packages installed on a device managed by the MSP portal, including version, operating system, and CPU architecture.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_secure_client_packages_page import MspSecureClientPackagesPage
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
    api_instance = scc_firewall_manager_sdk.MSPSecureClientManagementApi(api_client)
    device_uid = 'device_uid_example' # str | The unique identifier of the MSP-managed device.
    limit = '50' # str | Number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Secure Client packages installed on an MSP-managed device
        api_response = api_instance.get_msp_device_secure_client_packages(device_uid, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPSecureClientManagementApi->get_msp_device_secure_client_packages:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPSecureClientManagementApi->get_msp_device_secure_client_packages: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier of the MSP-managed device. | 
 **limit** | **str**| Number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**MspSecureClientPackagesPage**](MspSecureClientPackagesPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Secure Client packages installed on the MSP-managed device |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_secure_client_upgrade_run**
> MspSecureClientUpgradeRunDto get_msp_secure_client_upgrade_run(upgrade_run_uid)

Get MSP Secure Client Upgrade Run

Get an MSP Secure Client Upgrade Run by UID in the SCC Firewall Manager Tenant. Each upgrade run represents a group of devices across managed tenants having their Secure Client packages upgraded together.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_secure_client_upgrade_run_dto import MspSecureClientUpgradeRunDto
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
    api_instance = scc_firewall_manager_sdk.MSPSecureClientManagementApi(api_client)
    upgrade_run_uid = 'upgrade_run_uid_example' # str | The unique identifier, represented as a UUID, of the MSP Secure Client Upgrade Run in SCC Firewall Manager.

    try:
        # Get MSP Secure Client Upgrade Run
        api_response = api_instance.get_msp_secure_client_upgrade_run(upgrade_run_uid)
        print("The response of MSPSecureClientManagementApi->get_msp_secure_client_upgrade_run:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPSecureClientManagementApi->get_msp_secure_client_upgrade_run: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upgrade_run_uid** | **str**| The unique identifier, represented as a UUID, of the MSP Secure Client Upgrade Run in SCC Firewall Manager. | 

### Return type

[**MspSecureClientUpgradeRunDto**](MspSecureClientUpgradeRunDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The MSP Secure Client Upgrade Run object |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_secure_client_upgrade_runs**
> MspSecureClientUpgradeRunPage get_msp_secure_client_upgrade_runs(limit=limit, offset=offset, q=q, sort=sort)

Get MSP Secure Client Upgrade Runs

Get a list of MSP Secure Client upgrade runs in the SCC Firewall Manager Tenant. Each upgrade run represents a group of devices across managed tenants having their Secure Client packages upgraded together.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_secure_client_upgrade_run_page import MspSecureClientUpgradeRunPage
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
    api_instance = scc_firewall_manager_sdk.MSPSecureClientManagementApi(api_client)
    limit = '50' # str | Number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get MSP Secure Client Upgrade Runs
        api_response = api_instance.get_msp_secure_client_upgrade_runs(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPSecureClientManagementApi->get_msp_secure_client_upgrade_runs:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPSecureClientManagementApi->get_msp_secure_client_upgrade_runs: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**MspSecureClientUpgradeRunPage**](MspSecureClientUpgradeRunPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of MSP Secure Client Upgrade Run objects |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_secure_client_upgrade_runs_attribute_values**
> MspSecureClientUpgradeRunsAttributeValues get_msp_secure_client_upgrade_runs_attribute_values()

Get distinct attribute values for MSP Secure Client upgrade runs

Get distinct attribute values for filterable fields on MSP Secure Client upgrade runs.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_secure_client_upgrade_runs_attribute_values import MspSecureClientUpgradeRunsAttributeValues
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
    api_instance = scc_firewall_manager_sdk.MSPSecureClientManagementApi(api_client)

    try:
        # Get distinct attribute values for MSP Secure Client upgrade runs
        api_response = api_instance.get_msp_secure_client_upgrade_runs_attribute_values()
        print("The response of MSPSecureClientManagementApi->get_msp_secure_client_upgrade_runs_attribute_values:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPSecureClientManagementApi->get_msp_secure_client_upgrade_runs_attribute_values: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**MspSecureClientUpgradeRunsAttributeValues**](MspSecureClientUpgradeRunsAttributeValues.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Distinct attribute values for MSP Secure Client upgrade runs. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **upgrade_msp_managed_asa_secure_client_packages**
> CdoTransaction upgrade_msp_managed_asa_secure_client_packages(msp_upgrade_secure_client_packages_input)

Upgrade Secure Client packages on ASAs across multiple tenants

<p>Asynchronous operation to upgrade the Secure Client packages on ASA devices across managed tenants.</p> <p><strong>Response:</strong></p> <ul> <li>Returns a transaction object to track the upgrade operation progress</li> <li>Use the transaction ID or the entity UID to monitor the status of the upgrade</li> </ul>

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.msp_upgrade_secure_client_packages_input import MspUpgradeSecureClientPackagesInput
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
    api_instance = scc_firewall_manager_sdk.MSPSecureClientManagementApi(api_client)
    msp_upgrade_secure_client_packages_input = scc_firewall_manager_sdk.MspUpgradeSecureClientPackagesInput() # MspUpgradeSecureClientPackagesInput | 

    try:
        # Upgrade Secure Client packages on ASAs across multiple tenants
        api_response = api_instance.upgrade_msp_managed_asa_secure_client_packages(msp_upgrade_secure_client_packages_input)
        print("The response of MSPSecureClientManagementApi->upgrade_msp_managed_asa_secure_client_packages:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPSecureClientManagementApi->upgrade_msp_managed_asa_secure_client_packages: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msp_upgrade_secure_client_packages_input** | [**MspUpgradeSecureClientPackagesInput**](MspUpgradeSecureClientPackagesInput.md)|  | 

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
**202** | Returns a Transaction object that can be used to track the progress of the multi-tenant Secure Client upgrade operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

