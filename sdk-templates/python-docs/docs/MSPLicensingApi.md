# scc_firewall_manager_sdk.MSPLicensingApi

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
[**export_msp_device_licenses**](MSPLicensingApi.md#export_msp_device_licenses) | **POST** /v1/msp/licenses/devices/export | Export licenses for MSP-managed devices
[**export_msp_virtual_account_licenses**](MSPLicensingApi.md#export_msp_virtual_account_licenses) | **POST** /v1/msp/licenses/smart-accounts/{smartAccountUid}/virtual-accounts/{virtualAccountUid}/licenses/export | Export Licenses for a MSP-managed Virtual Account
[**get_msp_device_licenses**](MSPLicensingApi.md#get_msp_device_licenses) | **GET** /v1/msp/licenses/devices | Get licenses for MSP-managed devices
[**get_msp_device_licenses_by_uid**](MSPLicensingApi.md#get_msp_device_licenses_by_uid) | **GET** /v1/msp/licenses/devices/{mspManagedDeviceUid} | Get licenses for MSP-managed device
[**get_msp_smart_account_by_uid**](MSPLicensingApi.md#get_msp_smart_account_by_uid) | **GET** /v1/msp/licenses/smart-accounts/{smartAccountUid} | Get Smart Account by UID
[**get_msp_smart_accounts**](MSPLicensingApi.md#get_msp_smart_accounts) | **GET** /v1/msp/licenses/smart-accounts | Get Smart Accounts used by MSP-managed tenants.
[**get_msp_virtual_account_by_uid**](MSPLicensingApi.md#get_msp_virtual_account_by_uid) | **GET** /v1/msp/licenses/smart-accounts/{smartAccountUid}/virtual-accounts/{virtualAccountUid} | Get Virtual Account by UID
[**get_msp_virtual_account_licenses**](MSPLicensingApi.md#get_msp_virtual_account_licenses) | **GET** /v1/msp/licenses/smart-accounts/{smartAccountUid}/virtual-accounts/{virtualAccountUid}/licenses | Get Licenses for a Virtual Account
[**get_msp_virtual_accounts**](MSPLicensingApi.md#get_msp_virtual_accounts) | **GET** /v1/msp/licenses/smart-accounts/{smartAccountUid}/virtual-accounts | Get the Virtual Accounts for a specific Smart Account used by MSP-managed tenants.


# **export_msp_device_licenses**
> CdoTransaction export_msp_device_licenses(msp_export_input)

Export licenses for MSP-managed devices

This is an asynchronous operation to export device licenses, across all tenants, managed by the MSP portal, in CSV format. Once complete, the file can be downloaded using a presigned AWS S3 URL specified in the entityUrl field of the transaction that expires in 1 hour.

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
    api_instance = scc_firewall_manager_sdk.MSPLicensingApi(api_client)
    msp_export_input = scc_firewall_manager_sdk.MspExportInput() # MspExportInput | 

    try:
        # Export licenses for MSP-managed devices
        api_response = api_instance.export_msp_device_licenses(msp_export_input)
        print("The response of MSPLicensingApi->export_msp_device_licenses:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPLicensingApi->export_msp_device_licenses: %s\n" % e)
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

# **export_msp_virtual_account_licenses**
> CdoTransaction export_msp_virtual_account_licenses(smart_account_uid, virtual_account_uid, msp_export_input)

Export Licenses for a MSP-managed Virtual Account

This endpoint exports all of the licenses used by devices across all tenants, managed by the MSP portal, that are registered to Smart License using a token generated in this Virtual Account, in CSV format. Once complete, the file can be downloaded using a presigned AWS S3 URL specified in the entityUrl field of the transaction that expires in 1 hour.

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
    api_instance = scc_firewall_manager_sdk.MSPLicensingApi(api_client)
    smart_account_uid = 'smart_account_uid_example' # str | The unique identifier, represented as a UUID, of the smart account
    virtual_account_uid = 'virtual_account_uid_example' # str | The unique identifier, represented as a UUID, of the virtual account
    msp_export_input = scc_firewall_manager_sdk.MspExportInput() # MspExportInput | 

    try:
        # Export Licenses for a MSP-managed Virtual Account
        api_response = api_instance.export_msp_virtual_account_licenses(smart_account_uid, virtual_account_uid, msp_export_input)
        print("The response of MSPLicensingApi->export_msp_virtual_account_licenses:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPLicensingApi->export_msp_virtual_account_licenses: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smart_account_uid** | **str**| The unique identifier, represented as a UUID, of the smart account | 
 **virtual_account_uid** | **str**| The unique identifier, represented as a UUID, of the virtual account | 
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
**404** | Virtual Account with UID not found within the given Smart Account. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_device_licenses**
> MspDeviceLicensePage get_msp_device_licenses(limit=limit, offset=offset, q=q, sort=sort)

Get licenses for MSP-managed devices

Get device licenses, across all tenants, managed by the MSP portal.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_device_license_page import MspDeviceLicensePage
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
    api_instance = scc_firewall_manager_sdk.MSPLicensingApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get licenses for MSP-managed devices
        api_response = api_instance.get_msp_device_licenses(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPLicensingApi->get_msp_device_licenses:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPLicensingApi->get_msp_device_licenses: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**MspDeviceLicensePage**](MspDeviceLicensePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of licenses for MSP-managed devices |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_device_licenses_by_uid**
> MspDeviceLicenseDto get_msp_device_licenses_by_uid(msp_managed_device_uid)

Get licenses for MSP-managed device

Get device licenses for device managed by the MSP portal using its UID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_device_license_dto import MspDeviceLicenseDto
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
    api_instance = scc_firewall_manager_sdk.MSPLicensingApi(api_client)
    msp_managed_device_uid = 'msp_managed_device_uid_example' # str | The unique identifier of MSP-managed device.

    try:
        # Get licenses for MSP-managed device
        api_response = api_instance.get_msp_device_licenses_by_uid(msp_managed_device_uid)
        print("The response of MSPLicensingApi->get_msp_device_licenses_by_uid:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPLicensingApi->get_msp_device_licenses_by_uid: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **msp_managed_device_uid** | **str**| The unique identifier of MSP-managed device. | 

### Return type

[**MspDeviceLicenseDto**](MspDeviceLicenseDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Get device licenses for MSP-managed device |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_smart_account_by_uid**
> MspSmartAccountDto get_msp_smart_account_by_uid(smart_account_uid)

Get Smart Account by UID

Get Smart Account identified by UID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_smart_account_dto import MspSmartAccountDto
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
    api_instance = scc_firewall_manager_sdk.MSPLicensingApi(api_client)
    smart_account_uid = 'smart_account_uid_example' # str | The unique identifier (UID) of the smart account

    try:
        # Get Smart Account by UID
        api_response = api_instance.get_msp_smart_account_by_uid(smart_account_uid)
        print("The response of MSPLicensingApi->get_msp_smart_account_by_uid:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPLicensingApi->get_msp_smart_account_by_uid: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smart_account_uid** | **str**| The unique identifier (UID) of the smart account | 

### Return type

[**MspSmartAccountDto**](MspSmartAccountDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Smart Account for specified UID |  -  |
**404** | Smart Account not found for the specified UID |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_smart_accounts**
> MspSmartAccountsPage get_msp_smart_accounts(limit=limit, offset=offset, q=q, sort=sort)

Get Smart Accounts used by MSP-managed tenants.

Get the Smart Accounts and the number of Virtual Accounts used by MSP-managed tenants. Note: This endpoint does not display all the Smart Accounts, or number of Virtual Accounts within a Smart Account, to which the customer has access. Only Smart Accounts — and number of Virtual Accounts therein — that have licenses used by devices in MSP-managed tenants are displayed.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_smart_accounts_page import MspSmartAccountsPage
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
    api_instance = scc_firewall_manager_sdk.MSPLicensingApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Smart Accounts used by MSP-managed tenants.
        api_response = api_instance.get_msp_smart_accounts(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPLicensingApi->get_msp_smart_accounts:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPLicensingApi->get_msp_smart_accounts: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**MspSmartAccountsPage**](MspSmartAccountsPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Smart Accounts |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_virtual_account_by_uid**
> MspVirtualAccountDto get_msp_virtual_account_by_uid(smart_account_uid, virtual_account_uid)

Get Virtual Account by UID

Get Virtual Account identified by UID within a specific Smart Account.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_virtual_account_dto import MspVirtualAccountDto
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
    api_instance = scc_firewall_manager_sdk.MSPLicensingApi(api_client)
    smart_account_uid = 'smart_account_uid_example' # str | The unique identifier (UID) of the smart account
    virtual_account_uid = 'virtual_account_uid_example' # str | The unique identifier (UID) of the virtual account

    try:
        # Get Virtual Account by UID
        api_response = api_instance.get_msp_virtual_account_by_uid(smart_account_uid, virtual_account_uid)
        print("The response of MSPLicensingApi->get_msp_virtual_account_by_uid:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPLicensingApi->get_msp_virtual_account_by_uid: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smart_account_uid** | **str**| The unique identifier (UID) of the smart account | 
 **virtual_account_uid** | **str**| The unique identifier (UID) of the virtual account | 

### Return type

[**MspVirtualAccountDto**](MspVirtualAccountDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Virtual Account for specified UID |  -  |
**404** | Virtual Account not found for the specified UID |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_virtual_account_licenses**
> MspLicensesPage get_msp_virtual_account_licenses(smart_account_uid, virtual_account_uid, limit=limit, offset=offset, q=q, sort=sort)

Get Licenses for a Virtual Account

This endpoint returns information on all of the licenses used by devices across all MSP-managed tenants that are registered to Smart License using a token generated in this Virtual Account.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_licenses_page import MspLicensesPage
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
    api_instance = scc_firewall_manager_sdk.MSPLicensingApi(api_client)
    smart_account_uid = 'smart_account_uid_example' # str | The unique identifier (UID) of the smart account
    virtual_account_uid = 'virtual_account_uid_example' # str | The unique identifier (UID) of the virtual account
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Licenses for a Virtual Account
        api_response = api_instance.get_msp_virtual_account_licenses(smart_account_uid, virtual_account_uid, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPLicensingApi->get_msp_virtual_account_licenses:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPLicensingApi->get_msp_virtual_account_licenses: %s\n" % e)
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

[**MspLicensesPage**](MspLicensesPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of licenses for the specified Virtual Account. |  -  |
**404** | Virtual Account with UID not found within the given Smart Account. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_msp_virtual_accounts**
> MspVirtualAccountsPage get_msp_virtual_accounts(smart_account_uid, limit=limit, offset=offset, q=q, sort=sort)

Get the Virtual Accounts for a specific Smart Account used by MSP-managed tenants.

Get the Virtual Accounts for a specific Smart Account used by MSP-managed tenants. Note: This endpoint does not display all the Virtual Accounts within the Smart Account to which the customer has access. Only Virtual Accounts that have licenses used by devices in MSP-managed tenants are displayed.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.msp_virtual_accounts_page import MspVirtualAccountsPage
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
    api_instance = scc_firewall_manager_sdk.MSPLicensingApi(api_client)
    smart_account_uid = 'smart_account_uid_example' # str | The unique identifier (UID) of the smart account
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get the Virtual Accounts for a specific Smart Account used by MSP-managed tenants.
        api_response = api_instance.get_msp_virtual_accounts(smart_account_uid, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of MSPLicensingApi->get_msp_virtual_accounts:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPLicensingApi->get_msp_virtual_accounts: %s\n" % e)
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

[**MspVirtualAccountsPage**](MspVirtualAccountsPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Virtual Accounts |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

