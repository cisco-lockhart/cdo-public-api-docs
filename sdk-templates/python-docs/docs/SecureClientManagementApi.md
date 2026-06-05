# scc_firewall_manager_sdk.SecureClientManagementApi

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
[**get_asa_compatible_secure_client_versions**](SecureClientManagementApi.md#get_asa_compatible_secure_client_versions) | **GET** /v1/inventory/devices/asas/secure-client/upgrades/versions | List Secure Client versions compatible with the specified ASA devices
[**get_asa_secure_client_package_by_uid**](SecureClientManagementApi.md#get_asa_secure_client_package_by_uid) | **GET** /v1/inventory/devices/asas/{deviceUid}/secure-client/packages/{packageUid} | Get a Secure Client package installed on an ASA device by UID
[**get_asa_secure_client_packages**](SecureClientManagementApi.md#get_asa_secure_client_packages) | **GET** /v1/inventory/devices/asas/{deviceUid}/secure-client/packages | List Secure Client packages installed on an ASA device
[**get_secure_client_upgrade_run**](SecureClientManagementApi.md#get_secure_client_upgrade_run) | **GET** /v1/inventory/devices/secure-client/upgrades/runs/{upgradeRunUid} | Get Secure Client Upgrade Run
[**get_secure_client_upgrade_runs**](SecureClientManagementApi.md#get_secure_client_upgrade_runs) | **GET** /v1/inventory/devices/secure-client/upgrades/runs | Get Secure Client Upgrade Runs
[**get_secure_client_upgrade_runs_attribute_values**](SecureClientManagementApi.md#get_secure_client_upgrade_runs_attribute_values) | **GET** /v1/inventory/devices/secure-client/upgrades/runs/attribute-values | Get distinct attribute values for Secure Client upgrade runs
[**trigger_secure_client_upgrade**](SecureClientManagementApi.md#trigger_secure_client_upgrade) | **POST** /v1/inventory/devices/asas/secure-client/upgrades | Trigger Secure Client upgrade on ASA devices


# **get_asa_compatible_secure_client_versions**
> SecureClientUpgradeVersionsPage get_asa_compatible_secure_client_versions(device_uids, limit=limit, offset=offset)

List Secure Client versions compatible with the specified ASA devices

Get a list of compatible Secure Client upgrade versions for the specified ASA devices.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.secure_client_upgrade_versions_page import SecureClientUpgradeVersionsPage
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
    api_instance = scc_firewall_manager_sdk.SecureClientManagementApi(api_client)
    device_uids = 'device_uids_example' # str | The unique identifiers, represented as UUIDs, of the ASA devices in Security Cloud Control.
    limit = '50' # str | Number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')

    try:
        # List Secure Client versions compatible with the specified ASA devices
        api_response = api_instance.get_asa_compatible_secure_client_versions(device_uids, limit=limit, offset=offset)
        print("The response of SecureClientManagementApi->get_asa_compatible_secure_client_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SecureClientManagementApi->get_asa_compatible_secure_client_versions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uids** | **str**| The unique identifiers, represented as UUIDs, of the ASA devices in Security Cloud Control. | 
 **limit** | **str**| Number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]

### Return type

[**SecureClientUpgradeVersionsPage**](SecureClientUpgradeVersionsPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | A paginated list of compatible Secure Client versions |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |
**503** | Service temporarily unavailable |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_secure_client_package_by_uid**
> SecureClientPackageDto get_asa_secure_client_package_by_uid(device_uid, package_uid)

Get a Secure Client package installed on an ASA device by UID

Returns a single Secure Client webdeploy package installed on an ASA device, including  version, operating system, and CPU architecture.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.secure_client_package_dto import SecureClientPackageDto
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
    api_instance = scc_firewall_manager_sdk.SecureClientManagementApi(api_client)
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the ASA device in Security Cloud Control.
    package_uid = 'package_uid_example' # str | The unique identifier, represented as a UUID, of the Secure Client package.

    try:
        # Get a Secure Client package installed on an ASA device by UID
        api_response = api_instance.get_asa_secure_client_package_by_uid(device_uid, package_uid)
        print("The response of SecureClientManagementApi->get_asa_secure_client_package_by_uid:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SecureClientManagementApi->get_asa_secure_client_package_by_uid: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the ASA device in Security Cloud Control. | 
 **package_uid** | **str**| The unique identifier, represented as a UUID, of the Secure Client package. | 

### Return type

[**SecureClientPackageDto**](SecureClientPackageDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Secure Client package retrieved successfully |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_secure_client_packages**
> SecureClientPackagesPage get_asa_secure_client_packages(device_uid, limit=limit, offset=offset)

List Secure Client packages installed on an ASA device

Returns Secure Client webdeploy packages installed on an ASA device, including  version, operating system, and CPU architecture.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.secure_client_packages_page import SecureClientPackagesPage
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
    api_instance = scc_firewall_manager_sdk.SecureClientManagementApi(api_client)
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the ASA device in Security Cloud Control.
    limit = '50' # str | Number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')

    try:
        # List Secure Client packages installed on an ASA device
        api_response = api_instance.get_asa_secure_client_packages(device_uid, limit=limit, offset=offset)
        print("The response of SecureClientManagementApi->get_asa_secure_client_packages:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SecureClientManagementApi->get_asa_secure_client_packages: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the ASA device in Security Cloud Control. | 
 **limit** | **str**| Number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]

### Return type

[**SecureClientPackagesPage**](SecureClientPackagesPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Secure Client packages retrieved successfully |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_secure_client_upgrade_run**
> SecureClientUpgradeRunDto get_secure_client_upgrade_run(upgrade_run_uid)

Get Secure Client Upgrade Run

Get a Secure Client Upgrade Run by UID in the Security Cloud Control tenant. Each upgrade run represents a group of devices having their Secure Client packages upgraded together.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.secure_client_upgrade_run_dto import SecureClientUpgradeRunDto
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
    api_instance = scc_firewall_manager_sdk.SecureClientManagementApi(api_client)
    upgrade_run_uid = 'upgrade_run_uid_example' # str | The unique identifier, represented as a UUID, of the Secure Client Upgrade Run in Security Cloud Control.

    try:
        # Get Secure Client Upgrade Run
        api_response = api_instance.get_secure_client_upgrade_run(upgrade_run_uid)
        print("The response of SecureClientManagementApi->get_secure_client_upgrade_run:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SecureClientManagementApi->get_secure_client_upgrade_run: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **upgrade_run_uid** | **str**| The unique identifier, represented as a UUID, of the Secure Client Upgrade Run in Security Cloud Control. | 

### Return type

[**SecureClientUpgradeRunDto**](SecureClientUpgradeRunDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The Secure Client Upgrade Run object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_secure_client_upgrade_runs**
> SecureClientUpgradeRunPage get_secure_client_upgrade_runs(limit=limit, offset=offset, q=q, sort=sort)

Get Secure Client Upgrade Runs

Get a list of Secure Client upgrade runs in the Security Cloud Control tenant. Each upgrade run represents a group of devices having their Secure Client packages upgraded together.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.secure_client_upgrade_run_page import SecureClientUpgradeRunPage
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
    api_instance = scc_firewall_manager_sdk.SecureClientManagementApi(api_client)
    limit = '50' # str | Number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = 'name:DESC' # str | The fields to sort results by. (optional)

    try:
        # Get Secure Client Upgrade Runs
        api_response = api_instance.get_secure_client_upgrade_runs(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of SecureClientManagementApi->get_secure_client_upgrade_runs:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SecureClientManagementApi->get_secure_client_upgrade_runs: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | **str**| The fields to sort results by. | [optional] 

### Return type

[**SecureClientUpgradeRunPage**](SecureClientUpgradeRunPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Secure Client Upgrade Run objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_secure_client_upgrade_runs_attribute_values**
> SecureClientUpgradeRunAttributeValues get_secure_client_upgrade_runs_attribute_values()

Get distinct attribute values for Secure Client upgrade runs

Get distinct attribute values for filterable fields on Secure Client upgrade runs in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.secure_client_upgrade_run_attribute_values import SecureClientUpgradeRunAttributeValues
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
    api_instance = scc_firewall_manager_sdk.SecureClientManagementApi(api_client)

    try:
        # Get distinct attribute values for Secure Client upgrade runs
        api_response = api_instance.get_secure_client_upgrade_runs_attribute_values()
        print("The response of SecureClientManagementApi->get_secure_client_upgrade_runs_attribute_values:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SecureClientManagementApi->get_secure_client_upgrade_runs_attribute_values: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**SecureClientUpgradeRunAttributeValues**](SecureClientUpgradeRunAttributeValues.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Distinct attribute values for Secure Client upgrade runs. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **trigger_secure_client_upgrade**
> CdoTransaction trigger_secure_client_upgrade(trigger_secure_client_upgrade_input)

Trigger Secure Client upgrade on ASA devices

Initiate a Secure Client package upgrade on one or more ASA devices. The endpoint validates that all target devices are eligible, creates a CDO transaction and a SecureClientUpgradeRun tracking object, and returns 202 Accepted with the transaction.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.trigger_secure_client_upgrade_input import TriggerSecureClientUpgradeInput
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
    api_instance = scc_firewall_manager_sdk.SecureClientManagementApi(api_client)
    trigger_secure_client_upgrade_input = scc_firewall_manager_sdk.TriggerSecureClientUpgradeInput() # TriggerSecureClientUpgradeInput | 

    try:
        # Trigger Secure Client upgrade on ASA devices
        api_response = api_instance.trigger_secure_client_upgrade(trigger_secure_client_upgrade_input)
        print("The response of SecureClientManagementApi->trigger_secure_client_upgrade:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SecureClientManagementApi->trigger_secure_client_upgrade: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trigger_secure_client_upgrade_input** | [**TriggerSecureClientUpgradeInput**](TriggerSecureClientUpgradeInput.md)|  | 

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
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

