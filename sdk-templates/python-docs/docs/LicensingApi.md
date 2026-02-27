# scc_firewall_manager_sdk.LicensingApi

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
[**export_device_licenses**](LicensingApi.md#export_device_licenses) | **POST** /v1/licenses/devices/export | Export Per-Device Licenses
[**export_virtual_account_licenses**](LicensingApi.md#export_virtual_account_licenses) | **POST** /v1/licenses/smart-accounts/{smartAccountUid}/virtual-accounts/{virtualAccountUid}/licenses/export | Export Licenses for a Virtual Account
[**get_device_licenses**](LicensingApi.md#get_device_licenses) | **GET** /v1/licenses/devices | Get Device Licenses
[**get_device_licenses_by_uid**](LicensingApi.md#get_device_licenses_by_uid) | **GET** /v1/licenses/devices/{deviceUid} | Get Device Licensing Information by Device UID
[**get_smart_account_by_uid**](LicensingApi.md#get_smart_account_by_uid) | **GET** /v1/licenses/smart-accounts/{smartAccountUid} | Get Smart Account by UID
[**get_smart_accounts**](LicensingApi.md#get_smart_accounts) | **GET** /v1/licenses/smart-accounts | Get Smart Accounts used in this tenant.
[**get_virtual_account_by_uid**](LicensingApi.md#get_virtual_account_by_uid) | **GET** /v1/licenses/smart-accounts/{smartAccountUid}/virtual-accounts/{virtualAccountUid} | Get Virtual Account by UID
[**get_virtual_account_licenses**](LicensingApi.md#get_virtual_account_licenses) | **GET** /v1/licenses/smart-accounts/{smartAccountUid}/virtual-accounts/{virtualAccountUid}/licenses | Get Licenses for a Virtual Account
[**get_virtual_accounts**](LicensingApi.md#get_virtual_accounts) | **GET** /v1/licenses/smart-accounts/{smartAccountUid}/virtual-accounts | Get the Virtual Accounts for a specific Smart Account used in this tenant.


# **export_device_licenses**
> CdoTransaction export_device_licenses(export_input)

Export Per-Device Licenses

This is an asynchronous operation to export per-device licenses in CSV format. Once complete, the file can be downloaded using a presigned AWS S3 URL specified in the entityUrl field of the transaction that expires in 1 hour. Note: ⚠️ This endpoint is currently in limited availability. Please contact your Cisco account team or Cisco TAC to request access.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.export_input import ExportInput
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
    api_instance = scc_firewall_manager_sdk.LicensingApi(api_client)
    export_input = scc_firewall_manager_sdk.ExportInput() # ExportInput | 

    try:
        # Export Per-Device Licenses
        api_response = api_instance.export_device_licenses(export_input)
        print("The response of LicensingApi->export_device_licenses:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling LicensingApi->export_device_licenses: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **export_input** | [**ExportInput**](ExportInput.md)|  | 

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
**401** | Request not authorized. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **export_virtual_account_licenses**
> CdoTransaction export_virtual_account_licenses(smart_account_uid, virtual_account_uid, export_input)

Export Licenses for a Virtual Account

This endpoint exports all of the licenses used by devices across this tenant that are registered to Smart License using a token generated in this Virtual Account. Note: ⚠️ This endpoint is currently in limited availability. Please contact your Cisco account team or Cisco TAC to request access.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.export_input import ExportInput
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
    api_instance = scc_firewall_manager_sdk.LicensingApi(api_client)
    smart_account_uid = 'smart_account_uid_example' # str | The unique identifier, represented as a UUID, of the smart account
    virtual_account_uid = 'virtual_account_uid_example' # str | The unique identifier, represented as a UUID, of the virtual account
    export_input = scc_firewall_manager_sdk.ExportInput() # ExportInput | 

    try:
        # Export Licenses for a Virtual Account
        api_response = api_instance.export_virtual_account_licenses(smart_account_uid, virtual_account_uid, export_input)
        print("The response of LicensingApi->export_virtual_account_licenses:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling LicensingApi->export_virtual_account_licenses: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smart_account_uid** | **str**| The unique identifier, represented as a UUID, of the smart account | 
 **virtual_account_uid** | **str**| The unique identifier, represented as a UUID, of the virtual account | 
 **export_input** | [**ExportInput**](ExportInput.md)|  | 

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
**401** | Request not authorized. |  -  |
**404** | Virtual Account with UID not found within the given Smart Account. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device_licenses**
> DeviceLicensePage get_device_licenses(limit=limit, offset=offset, q=q, sort=sort)

Get Device Licenses

Get a paginated list of device licenses. Note: ⚠️ This endpoint is currently in limited availability. Please contact your Cisco account team or Cisco TAC to request access.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device_license_page import DeviceLicensePage
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
    api_instance = scc_firewall_manager_sdk.LicensingApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Device Licenses
        api_response = api_instance.get_device_licenses(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of LicensingApi->get_device_licenses:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling LicensingApi->get_device_licenses: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**DeviceLicensePage**](DeviceLicensePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of device licenses |  -  |
**401** | Request not authorized. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device_licenses_by_uid**
> DeviceLicenseDto get_device_licenses_by_uid(device_uid)

Get Device Licensing Information by Device UID

Retrieves the device licensing information for a specific device identified by UID. Note: ⚠️ This endpoint is currently in limited availability. Please contact your Cisco account team or Cisco TAC to request access.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device_license_dto import DeviceLicenseDto
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
    api_instance = scc_firewall_manager_sdk.LicensingApi(api_client)
    device_uid = 'device_uid_example' # str | The unique identifier (UID) of the device

    try:
        # Get Device Licensing Information by Device UID
        api_response = api_instance.get_device_licenses_by_uid(device_uid)
        print("The response of LicensingApi->get_device_licenses_by_uid:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling LicensingApi->get_device_licenses_by_uid: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier (UID) of the device | 

### Return type

[**DeviceLicenseDto**](DeviceLicenseDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Device licensing information for specified device UID |  -  |
**401** | Request not authorized. |  -  |
**404** | Device licensing information not found for the specified device UID |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_smart_account_by_uid**
> SmartAccountDto get_smart_account_by_uid(smart_account_uid)

Get Smart Account by UID

Get Smart Account identified by UID. Note: ⚠️ This endpoint is currently in limited availability. Please contact your Cisco account team or Cisco TAC to request access.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.smart_account_dto import SmartAccountDto
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
    api_instance = scc_firewall_manager_sdk.LicensingApi(api_client)
    smart_account_uid = 'smart_account_uid_example' # str | The unique identifier (UID) of the smart account

    try:
        # Get Smart Account by UID
        api_response = api_instance.get_smart_account_by_uid(smart_account_uid)
        print("The response of LicensingApi->get_smart_account_by_uid:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling LicensingApi->get_smart_account_by_uid: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smart_account_uid** | **str**| The unique identifier (UID) of the smart account | 

### Return type

[**SmartAccountDto**](SmartAccountDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Smart Account for specified UID |  -  |
**401** | Request not authorized. |  -  |
**404** | Smart Account not found for the specified UID |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_smart_accounts**
> SmartAccountsPage get_smart_accounts(limit=limit, offset=offset, q=q, sort=sort)

Get Smart Accounts used in this tenant.

Get the Smart Accounts and the number of Virtual Accounts used in this tenant. Note: This endpoint does not display all the Smart Accounts, or number of Virtual Accounts within a Smart Account, to which the customer has access. Only Smart Accounts — and number of Virtual Accounts therein — that have licenses used by devices in this tenant are displayed. Note: ⚠️ This endpoint is currently in limited availability. Please contact your Cisco account team or Cisco TAC to request access.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.smart_accounts_page import SmartAccountsPage
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
    api_instance = scc_firewall_manager_sdk.LicensingApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Smart Accounts used in this tenant.
        api_response = api_instance.get_smart_accounts(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of LicensingApi->get_smart_accounts:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling LicensingApi->get_smart_accounts: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**SmartAccountsPage**](SmartAccountsPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Smart Accounts |  -  |
**401** | Request not authorized. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_virtual_account_by_uid**
> VirtualAccountDto get_virtual_account_by_uid(smart_account_uid, virtual_account_uid)

Get Virtual Account by UID

Get Virtual Account identified by UID within a specific Smart Account. Note: ⚠️ This endpoint is currently in limited availability. Please contact your Cisco account team or Cisco TAC to request access.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.virtual_account_dto import VirtualAccountDto
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
    api_instance = scc_firewall_manager_sdk.LicensingApi(api_client)
    smart_account_uid = 'smart_account_uid_example' # str | The unique identifier (UID) of the smart account
    virtual_account_uid = 'virtual_account_uid_example' # str | The unique identifier (UID) of the virtual account

    try:
        # Get Virtual Account by UID
        api_response = api_instance.get_virtual_account_by_uid(smart_account_uid, virtual_account_uid)
        print("The response of LicensingApi->get_virtual_account_by_uid:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling LicensingApi->get_virtual_account_by_uid: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smart_account_uid** | **str**| The unique identifier (UID) of the smart account | 
 **virtual_account_uid** | **str**| The unique identifier (UID) of the virtual account | 

### Return type

[**VirtualAccountDto**](VirtualAccountDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Virtual Account for specified UID |  -  |
**401** | Request not authorized. |  -  |
**404** | Virtual Account not found for the specified UID |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_virtual_account_licenses**
> Page get_virtual_account_licenses(smart_account_uid, virtual_account_uid, limit=limit, offset=offset, q=q, sort=sort)

Get Licenses for a Virtual Account

This endpoint returns information on all of the licenses used by devices across this tenant that are registered to Smart License using a token generated in this Virtual Account. Note: ⚠️ This endpoint is currently in limited availability. Please contact your Cisco account team or Cisco TAC to request access.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.page import Page
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
    api_instance = scc_firewall_manager_sdk.LicensingApi(api_client)
    smart_account_uid = 'smart_account_uid_example' # str | The unique identifier (UID) of the smart account
    virtual_account_uid = 'virtual_account_uid_example' # str | The unique identifier (UID) of the virtual account
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Licenses for a Virtual Account
        api_response = api_instance.get_virtual_account_licenses(smart_account_uid, virtual_account_uid, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of LicensingApi->get_virtual_account_licenses:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling LicensingApi->get_virtual_account_licenses: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smart_account_uid** | **str**| The unique identifier (UID) of the smart account | 
 **virtual_account_uid** | **str**| The unique identifier (UID) of the virtual account | 
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**Page**](Page.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of licenses for the specified Virtual Account. |  -  |
**401** | Request not authorized. |  -  |
**404** | Virtual Account with UID not found within the given Smart Account. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_virtual_accounts**
> VirtualAccountsPage get_virtual_accounts(smart_account_uid, limit=limit, offset=offset, q=q, sort=sort)

Get the Virtual Accounts for a specific Smart Account used in this tenant.

Get the Virtual Accounts for a specific Smart Account used in this tenant. Note: This endpoint does not display all the Virtual Accounts within the Smart Account to which the customer has access. Only Virtual Accounts that have licenses used by devices in this tenant are displayed. Note: ⚠️ This endpoint is currently in limited availability. Please contact your Cisco account team or Cisco TAC to request access.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.virtual_accounts_page import VirtualAccountsPage
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
    api_instance = scc_firewall_manager_sdk.LicensingApi(api_client)
    smart_account_uid = 'smart_account_uid_example' # str | The unique identifier (UID) of the smart account
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get the Virtual Accounts for a specific Smart Account used in this tenant.
        api_response = api_instance.get_virtual_accounts(smart_account_uid, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of LicensingApi->get_virtual_accounts:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling LicensingApi->get_virtual_accounts: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smart_account_uid** | **str**| The unique identifier (UID) of the smart account | 
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**VirtualAccountsPage**](VirtualAccountsPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Virtual Accounts |  -  |
**401** | Request not authorized. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

