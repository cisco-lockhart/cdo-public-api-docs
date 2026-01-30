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
[**bulk_accept_asa_certificates**](InventoryApi.md#bulk_accept_asa_certificates) | **POST** /v1/inventory/devices/asas/acceptCert | Accept certificates for multiple ASA devices
[**bulk_read_asa_device_configurations**](InventoryApi.md#bulk_read_asa_device_configurations) | **POST** /v1/inventory/devices/asas/read | Read configurations for multiple ASA devices
[**create_duo_admin_panel**](InventoryApi.md#create_duo_admin_panel) | **POST** /v1/inventory/devices/duoAdminPanels | Onboard Duo Admin Panel
[**create_ftd_device**](InventoryApi.md#create_ftd_device) | **POST** /v1/inventory/devices/ftds | Onboard FTD device
[**delete_cd_fmc_managed_ftd_device**](InventoryApi.md#delete_cd_fmc_managed_ftd_device) | **POST** /v1/inventory/devices/ftds/cdfmcManaged/{deviceUid}/delete | Delete cdFMC managed FTD device
[**delete_cloud_service**](InventoryApi.md#delete_cloud_service) | **DELETE** /v1/inventory/services/{cloudServiceUid} | Delete Cloud Service
[**delete_device**](InventoryApi.md#delete_device) | **DELETE** /v1/inventory/devices/{deviceUid} | Delete Device
[**delete_device_manager**](InventoryApi.md#delete_device_manager) | **DELETE** /v1/inventory/managers/{deviceManagerUid} | Delete Device Manager
[**delete_template_device**](InventoryApi.md#delete_template_device) | **DELETE** /v1/inventory/templates/{templateDeviceUid} | Delete Template Device
[**deploy_asa_device_changes**](InventoryApi.md#deploy_asa_device_changes) | **POST** /v1/inventory/devices/asas/{deviceUid}/deploy | Deploy ASA device changes
[**deploy_changes_to_multiple_ftd_devices**](InventoryApi.md#deploy_changes_to_multiple_ftd_devices) | **POST** /v1/inventory/devices/ftds/deploy | (cdFMC-managed FTDs only) Deploy changes to multiple FTD devices
[**deploy_ftd_device_changes**](InventoryApi.md#deploy_ftd_device_changes) | **POST** /v1/inventory/devices/ftds/{deviceUid}/deploy | (cdFMC-managed FTDs only) Deploy FTD device changes
[**enable_multicloud_defense**](InventoryApi.md#enable_multicloud_defense) | **POST** /v1/inventory/managers/mcd | Enable Multicloud Defense
[**finish_onboarding_ftd_device**](InventoryApi.md#finish_onboarding_ftd_device) | **POST** /v1/inventory/devices/ftds/register | Register FTD device to FMC
[**get_asa_configuration**](InventoryApi.md#get_asa_configuration) | **GET** /v1/inventory/devices/asas/{deviceUid}/configs | Get ASA configuration details
[**get_cloud_service**](InventoryApi.md#get_cloud_service) | **GET** /v1/inventory/services/{cloudServiceUid} | Get Cloud Service
[**get_cloud_services**](InventoryApi.md#get_cloud_services) | **GET** /v1/inventory/services | Get Cloud Services
[**get_device**](InventoryApi.md#get_device) | **GET** /v1/inventory/devices/{deviceUid} | Get Device
[**get_device_end_of_life_report**](InventoryApi.md#get_device_end_of_life_report) | **GET** /v1/inventory/devices/{deviceUid}/endoflifereports | Get Device End-Of-Life Report
[**get_device_end_of_life_reports**](InventoryApi.md#get_device_end_of_life_reports) | **GET** /v1/inventory/devices/endoflifereports | Get Device End-Of-Life Reports
[**get_device_manager**](InventoryApi.md#get_device_manager) | **GET** /v1/inventory/managers/{deviceManagerUid} | Get Device Manager
[**get_device_managers**](InventoryApi.md#get_device_managers) | **GET** /v1/inventory/managers | Get Device Managers
[**get_devices**](InventoryApi.md#get_devices) | **GET** /v1/inventory/devices | Get Devices
[**get_fmc_health**](InventoryApi.md#get_fmc_health) | **GET** /v1/inventory/managers/{fmcUid}/health/metrics | Get health metrics on devices managed by the FMC (cdFMC only)
[**get_template_device**](InventoryApi.md#get_template_device) | **GET** /v1/inventory/templates/{templateDeviceUid} | Get Template Device
[**get_template_devices**](InventoryApi.md#get_template_devices) | **GET** /v1/inventory/templates | Get Template Devices
[**modify_cloud_service**](InventoryApi.md#modify_cloud_service) | **PATCH** /v1/inventory/services/{cloudServiceUid} | Modify Cloud Service
[**modify_device**](InventoryApi.md#modify_device) | **PATCH** /v1/inventory/devices/{deviceUid} | Modify Device
[**modify_device_manager**](InventoryApi.md#modify_device_manager) | **PATCH** /v1/inventory/managers/{deviceManagerUid} | Modify Device Manager
[**modify_devices**](InventoryApi.md#modify_devices) | **PATCH** /v1/inventory/devices/bulk | Modify Devices
[**modify_template_device**](InventoryApi.md#modify_template_device) | **PATCH** /v1/inventory/templates/{templateDeviceUid} | Modify Template Device
[**onboard_asa_device**](InventoryApi.md#onboard_asa_device) | **POST** /v1/inventory/devices/asas | Onboard ASA device
[**onboard_ftd_device_using_ztp**](InventoryApi.md#onboard_ftd_device_using_ztp) | **POST** /v1/inventory/devices/ftds/ztp | Onboard FTD device using Zero-Touch Provisioning
[**onboard_ios_device**](InventoryApi.md#onboard_ios_device) | **POST** /v1/inventory/devices/ios | Onboard IOS Device
[**provision_cd_fmc**](InventoryApi.md#provision_cd_fmc) | **POST** /v1/inventory/managers/cdfmc | Provision cdFMC
[**read_asa_device_configuration**](InventoryApi.md#read_asa_device_configuration) | **POST** /v1/inventory/devices/asas/{deviceUid}/read | Read ASA device configuration
[**reconnect_asa_device**](InventoryApi.md#reconnect_asa_device) | **POST** /v1/inventory/devices/asas/{deviceUid}/reconnect | Reconnect ASA device


# **bulk_accept_asa_certificates**
> CdoTransaction bulk_accept_asa_certificates(bulk_operation_asa_device_request)

Accept certificates for multiple ASA devices

This is an asynchronous operation to accept certificates for multiple ASA devices. This operation returns a link to a transaction object that can be used to monitor the progress of the operation for all devices.<br/>**Warning:** This operation will accept the certificates for the ASA devices in the body without providing the user with the ability to view the certificates. It is recommended to use this operation only when you are sure that all certificates are valid.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.bulk_operation_asa_device_request import BulkOperationAsaDeviceRequest
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
    bulk_operation_asa_device_request = scc_firewall_manager_sdk.BulkOperationAsaDeviceRequest() # BulkOperationAsaDeviceRequest | 

    try:
        # Accept certificates for multiple ASA devices
        api_response = api_instance.bulk_accept_asa_certificates(bulk_operation_asa_device_request)
        print("The response of InventoryApi->bulk_accept_asa_certificates:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->bulk_accept_asa_certificates: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulk_operation_asa_device_request** | [**BulkOperationAsaDeviceRequest**](BulkOperationAsaDeviceRequest.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the accept certificates operation |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **bulk_read_asa_device_configurations**
> CdoTransaction bulk_read_asa_device_configurations(bulk_operation_asa_device_request)

Read configurations for multiple ASA devices

This is an asynchronous operation to read the latest configurations for multiple ASA devices into the Security Cloud Control tenant. This operation returns a link to a transaction object that can be used to monitor the progress of the operation for all devices.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.bulk_operation_asa_device_request import BulkOperationAsaDeviceRequest
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
    bulk_operation_asa_device_request = scc_firewall_manager_sdk.BulkOperationAsaDeviceRequest() # BulkOperationAsaDeviceRequest | 

    try:
        # Read configurations for multiple ASA devices
        api_response = api_instance.bulk_read_asa_device_configurations(bulk_operation_asa_device_request)
        print("The response of InventoryApi->bulk_read_asa_device_configurations:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->bulk_read_asa_device_configurations: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulk_operation_asa_device_request** | [**BulkOperationAsaDeviceRequest**](BulkOperationAsaDeviceRequest.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the read operation |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_duo_admin_panel**
> CdoTransaction create_duo_admin_panel(duo_admin_panel_create_or_update_input)

Onboard Duo Admin Panel

Onboard a Duo Admin Panel to the CDO tenant. The credentials to onboard the Duo Admin Panel to Security Cloud Control must be generated by creating an Admin API application on https://www.duo.com. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.duo_admin_panel_create_or_update_input import DuoAdminPanelCreateOrUpdateInput
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
    duo_admin_panel_create_or_update_input = scc_firewall_manager_sdk.DuoAdminPanelCreateOrUpdateInput() # DuoAdminPanelCreateOrUpdateInput | 

    try:
        # Onboard Duo Admin Panel
        api_response = api_instance.create_duo_admin_panel(duo_admin_panel_create_or_update_input)
        print("The response of InventoryApi->create_duo_admin_panel:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->create_duo_admin_panel: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **duo_admin_panel_create_or_update_input** | [**DuoAdminPanelCreateOrUpdateInput**](DuoAdminPanelCreateOrUpdateInput.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the onboarding operation |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_ftd_device**
> CdoTransaction create_ftd_device(ftd_create_or_update_input)

Onboard FTD device

This is an asynchronous operation to generate a registration key for a cdFMC managed FTD device in the CDO tenant. This operation returns a link to a transaction object that can be used to monitor the progress of the operation. Onboarding a cdFMC managed FTD device is a two-step process: the first step, handled by this operation, creates an FTD device with a configure manager string that must be pasted into the FTD device's Command-Line Interface. The FTD then uses this information to register itself with the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.ftd_create_or_update_input import FtdCreateOrUpdateInput
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
    ftd_create_or_update_input = scc_firewall_manager_sdk.FtdCreateOrUpdateInput() # FtdCreateOrUpdateInput | 

    try:
        # Onboard FTD device
        api_response = api_instance.create_ftd_device(ftd_create_or_update_input)
        print("The response of InventoryApi->create_ftd_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->create_ftd_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ftd_create_or_update_input** | [**FtdCreateOrUpdateInput**](FtdCreateOrUpdateInput.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the creation operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_cd_fmc_managed_ftd_device**
> CdoTransaction delete_cd_fmc_managed_ftd_device(device_uid)

Delete cdFMC managed FTD device

This is an asynchronous operation to delete cdFMC managed FTD device in the CDO tenant. This operation returns a link to a transaction object that can be used to monitor the progress of the operation. The reason this operation is asynchronous is because the device is first removed from the cdFMC, following which it is deleted from the Security Cloud Control tenant.

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
        # Delete cdFMC managed FTD device
        api_response = api_instance.delete_cd_fmc_managed_ftd_device(device_uid)
        print("The response of InventoryApi->delete_cd_fmc_managed_ftd_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->delete_cd_fmc_managed_ftd_device: %s\n" % e)
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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the deletion operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_cloud_service**
> delete_cloud_service(cloud_service_uid)

Delete Cloud Service

Delete a Cloud Service by UID in the Security Cloud Control tenant.

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
    api_instance = scc_firewall_manager_sdk.InventoryApi(api_client)
    cloud_service_uid = 'cloud_service_uid_example' # str | The unique identifier, represented as a UUID, of the cloud service in Security Cloud Control.

    try:
        # Delete Cloud Service
        api_instance.delete_cloud_service(cloud_service_uid)
    except Exception as e:
        print("Exception when calling InventoryApi->delete_cloud_service: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud_service_uid** | **str**| The unique identifier, represented as a UUID, of the cloud service in Security Cloud Control. | 

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

# **delete_device**
> delete_device(device_uid)

Delete Device

Delete a device by UID in the Security Cloud Control tenant. On-prem FMCs and cloud-delivered FMCs cannot be deleted using this endpoint.

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
    api_instance = scc_firewall_manager_sdk.InventoryApi(api_client)
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the device in Security Cloud Control.

    try:
        # Delete Device
        api_instance.delete_device(device_uid)
    except Exception as e:
        print("Exception when calling InventoryApi->delete_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the device in Security Cloud Control. | 

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
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_device_manager**
> delete_device_manager(device_manager_uid)

Delete Device Manager

Delete a Device Manager by UID in the Security Cloud Control tenant.

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
    api_instance = scc_firewall_manager_sdk.InventoryApi(api_client)
    device_manager_uid = 'device_manager_uid_example' # str | The unique identifier, represented as a UUID, of the device manager in Security Cloud Control.

    try:
        # Delete Device Manager
        api_instance.delete_device_manager(device_manager_uid)
    except Exception as e:
        print("Exception when calling InventoryApi->delete_device_manager: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_manager_uid** | **str**| The unique identifier, represented as a UUID, of the device manager in Security Cloud Control. | 

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
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_template_device**
> delete_template_device(template_device_uid)

Delete Template Device

Delete a template device by UID in the Security Cloud Control tenant.

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
    api_instance = scc_firewall_manager_sdk.InventoryApi(api_client)
    template_device_uid = 'template_device_uid_example' # str | The unique identifier, represented as a UUID, of the template device in Security Cloud Control.

    try:
        # Delete Template Device
        api_instance.delete_template_device(template_device_uid)
    except Exception as e:
        print("Exception when calling InventoryApi->delete_template_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **template_device_uid** | **str**| The unique identifier, represented as a UUID, of the template device in Security Cloud Control. | 

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

# **deploy_asa_device_changes**
> CdoTransaction deploy_asa_device_changes(device_uid)

Deploy ASA device changes

This is an asynchronous operation to deploy changes made to an ASA device's configuration on Security Cloud Control to the device. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.

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
    device_uid = 'device_uid_example' # str | 

    try:
        # Deploy ASA device changes
        api_response = api_instance.deploy_asa_device_changes(device_uid)
        print("The response of InventoryApi->deploy_asa_device_changes:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->deploy_asa_device_changes: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the ASA deploy operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deploy_changes_to_multiple_ftd_devices**
> CdoTransaction deploy_changes_to_multiple_ftd_devices(ftd_multi_device_deployment_input)

(cdFMC-managed FTDs only) Deploy changes to multiple FTD devices

This is an asynchronous operation to deploy changes made to the configurations of multiple cdFMC-managed FTD devices on Security Cloud Control. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.  Notes:  - This operation is only supported for cdFMC-managed FTD devices.  - This operation will only deploy changes to the device if there are pending changes to deploy.   

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.ftd_multi_device_deployment_input import FtdMultiDeviceDeploymentInput
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
    ftd_multi_device_deployment_input = scc_firewall_manager_sdk.FtdMultiDeviceDeploymentInput() # FtdMultiDeviceDeploymentInput | 

    try:
        # (cdFMC-managed FTDs only) Deploy changes to multiple FTD devices
        api_response = api_instance.deploy_changes_to_multiple_ftd_devices(ftd_multi_device_deployment_input)
        print("The response of InventoryApi->deploy_changes_to_multiple_ftd_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->deploy_changes_to_multiple_ftd_devices: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ftd_multi_device_deployment_input** | [**FtdMultiDeviceDeploymentInput**](FtdMultiDeviceDeploymentInput.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the creation operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **deploy_ftd_device_changes**
> CdoTransaction deploy_ftd_device_changes(device_uid, ftd_deployment_input)

(cdFMC-managed FTDs only) Deploy FTD device changes

This is an asynchronous operation to deploy changes made to a cdFMC-managed FTD device's configuration on Security Cloud Control to the device. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.  Notes:  - This operation is only supported for cdFMC-managed FTD devices.  - This operation will only deploy changes to the device if there are pending changes to deploy.  - Once this operation is finished, it can take up to 10 minutes for the [device](https://developer.cisco.com/docs/cisco-security-cloud-control/device/) [configState](https://developer.cisco.com/docs/cisco-security-cloud-control/configstate/) to be updated from `SYNCED` to `NOT_SYNCED` on Security Cloud Control. 

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.ftd_deployment_input import FtdDeploymentInput
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
    device_uid = 'device_uid_example' # str | 
    ftd_deployment_input = scc_firewall_manager_sdk.FtdDeploymentInput() # FtdDeploymentInput | 

    try:
        # (cdFMC-managed FTDs only) Deploy FTD device changes
        api_response = api_instance.deploy_ftd_device_changes(device_uid, ftd_deployment_input)
        print("The response of InventoryApi->deploy_ftd_device_changes:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->deploy_ftd_device_changes: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**|  | 
 **ftd_deployment_input** | [**FtdDeploymentInput**](FtdDeploymentInput.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the creation operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **enable_multicloud_defense**
> CdoTransaction enable_multicloud_defense()

Enable Multicloud Defense

This is an asynchronous operation to Enable Multicloud Defense for the Security Cloud Control tenant. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.

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

    try:
        # Enable Multicloud Defense
        api_response = api_instance.enable_multicloud_defense()
        print("The response of InventoryApi->enable_multicloud_defense:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->enable_multicloud_defense: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

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
**202** | Security Cloud Control Transaction object that can be used to track the status of the operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**409** | Conflict. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **finish_onboarding_ftd_device**
> CdoTransaction finish_onboarding_ftd_device(ftd_registration_input)

Register FTD device to FMC

Complete registration of an FTD device managed by an FMC to the Security Cloud Control tenant. Call this API endpoint after you have created an FTD and pasted the generated CLI output in the FTD. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.ftd_registration_input import FtdRegistrationInput
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
    ftd_registration_input = scc_firewall_manager_sdk.FtdRegistrationInput() # FtdRegistrationInput | 

    try:
        # Register FTD device to FMC
        api_response = api_instance.finish_onboarding_ftd_device(ftd_registration_input)
        print("The response of InventoryApi->finish_onboarding_ftd_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->finish_onboarding_ftd_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ftd_registration_input** | [**FtdRegistrationInput**](FtdRegistrationInput.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the creation operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_configuration**
> AsaConfig get_asa_configuration(device_uid)

Get ASA configuration details

Fetches the ASA configuration for a specified device by its unique identifier, represented as a UUID. This endpoint returns both the current configuration from the device ('configOnDevice') and the configuration stored in Security Cloud Control ('configOnCloud').

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_config import AsaConfig
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
    device_uid = 'device_uid_example' # str | 

    try:
        # Get ASA configuration details
        api_response = api_instance.get_asa_configuration(device_uid)
        print("The response of InventoryApi->get_asa_configuration:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->get_asa_configuration: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**|  | 

### Return type

[**AsaConfig**](AsaConfig.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Successfully retrieved ASA configuration details from both device and Security Cloud Control. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cloud_service**
> Device get_cloud_service(cloud_service_uid)

Get Cloud Service

Get a Cloud Service by UID in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device import Device
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
    cloud_service_uid = 'cloud_service_uid_example' # str | The unique identifier, represented as a UUID, of the cloud service in Security Cloud Control.

    try:
        # Get Cloud Service
        api_response = api_instance.get_cloud_service(cloud_service_uid)
        print("The response of InventoryApi->get_cloud_service:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->get_cloud_service: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud_service_uid** | **str**| The unique identifier, represented as a UUID, of the cloud service in Security Cloud Control. | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Cloud Service |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cloud_services**
> DevicePage get_cloud_services(limit=limit, offset=offset, q=q, sort=sort)

Get Cloud Services

Get a list of Cloud Services in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device_page import DevicePage
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
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Cloud Services
        api_response = api_instance.get_cloud_services(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of InventoryApi->get_cloud_services:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->get_cloud_services: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**DevicePage**](DevicePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Cloud Services |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device**
> Device get_device(device_uid)

Get Device

Get a device by UID in the Security Cloud Control tenant

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device import Device
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
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the device in Security Cloud Control.

    try:
        # Get Device
        api_response = api_instance.get_device(device_uid)
        print("The response of InventoryApi->get_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->get_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the device in Security Cloud Control. | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Device object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device_end_of_life_report**
> DeviceEndOfLifeReport get_device_end_of_life_report(device_uid)

Get Device End-Of-Life Report

Get the Device End-of-Life (EOL) report that details the EOL dates for the device hardware along with recommended hardware upgrade options.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device_end_of_life_report import DeviceEndOfLifeReport
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
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the Device.

    try:
        # Get Device End-Of-Life Report
        api_response = api_instance.get_device_end_of_life_report(device_uid)
        print("The response of InventoryApi->get_device_end_of_life_report:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->get_device_end_of_life_report: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the Device. | 

### Return type

[**DeviceEndOfLifeReport**](DeviceEndOfLifeReport.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Device End-of-Life (EOL) report detailing the EOL dates for the device hardware along with recommended hardware upgrade options. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device_end_of_life_reports**
> List[DeviceEndOfLifeReport] get_device_end_of_life_reports(q=q)

Get Device End-Of-Life Reports

The reports provide information on devices approaching their End-Of-Life (EOL) status, indicating that they will cease to receive vendor support. For each hardware, the report outlines key EOL dates, offers recommendations for appropriate hardware upgrades, and includes an inventory of affected devices.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device_end_of_life_report import DeviceEndOfLifeReport
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
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Get Device End-Of-Life Reports
        api_response = api_instance.get_device_end_of_life_reports(q=q)
        print("The response of InventoryApi->get_device_end_of_life_reports:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->get_device_end_of_life_reports: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 

### Return type

[**List[DeviceEndOfLifeReport]**](DeviceEndOfLifeReport.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of device End-Of-Life reports. For each hardware, the report outlines key EOL dates, offers recommendations for appropriate hardware upgrades, and includes an inventory of affected devices. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device_manager**
> Device get_device_manager(device_manager_uid)

Get Device Manager

Get a Device Manager by UID in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device import Device
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
    device_manager_uid = 'device_manager_uid_example' # str | The unique identifier, represented as a UUID, of the device manager in Security Cloud Control.

    try:
        # Get Device Manager
        api_response = api_instance.get_device_manager(device_manager_uid)
        print("The response of InventoryApi->get_device_manager:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->get_device_manager: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_manager_uid** | **str**| The unique identifier, represented as a UUID, of the device manager in Security Cloud Control. | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Device Manager |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_device_managers**
> DevicePage get_device_managers(limit=limit, offset=offset, q=q, sort=sort)

Get Device Managers

Fetch a list of Device Managers (on-prem FMCs and cloud-delivered FMCs) in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device_page import DevicePage
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
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Device Managers
        api_response = api_instance.get_device_managers(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of InventoryApi->get_device_managers:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->get_device_managers: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**DevicePage**](DevicePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Device Manager objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_devices**
> DevicePage get_devices(limit=limit, offset=offset, q=q, sort=sort)

Get Devices

Get a list of devices in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device_page import DevicePage
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
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Devices
        api_response = api_instance.get_devices(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of InventoryApi->get_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->get_devices: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**DevicePage**](DevicePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Device objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_fmc_health**
> List[FmcHealthMetrics] get_fmc_health(fmc_uid, time_range=time_range)

Get health metrics on devices managed by the FMC (cdFMC only)

Get metrics that indicate the current health of all devices managed by the cdFMC. Note: For specific health metrics to be available for a given device under management of the cdFMC, the health policy for that device should be configured to collect those metrics. For example, CPU metrics will be unavailable for a device if the health policy applied to that device has CPU metric collection turned off. Note: This endpoint can only be queried twice every minute.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.fmc_health_metrics import FmcHealthMetrics
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
    fmc_uid = 'fmc_uid_example' # str | The unique identifier, represented as a UUID, of the FMC in Security Cloud Control.
    time_range = 'time_range_example' # str | The time range for which results should be retrieved. (optional)

    try:
        # Get health metrics on devices managed by the FMC (cdFMC only)
        api_response = api_instance.get_fmc_health(fmc_uid, time_range=time_range)
        print("The response of InventoryApi->get_fmc_health:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->get_fmc_health: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fmc_uid** | **str**| The unique identifier, represented as a UUID, of the FMC in Security Cloud Control. | 
 **time_range** | **str**| The time range for which results should be retrieved. | [optional] 

### Return type

[**List[FmcHealthMetrics]**](FmcHealthMetrics.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | FMC health metrics |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_template_device**
> Device get_template_device(template_device_uid)

Get Template Device

Get a template device by UID in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device import Device
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
    template_device_uid = 'template_device_uid_example' # str | The unique identifier, represented as a UUID, of the template device in Security Cloud Control.

    try:
        # Get Template Device
        api_response = api_instance.get_template_device(template_device_uid)
        print("The response of InventoryApi->get_template_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->get_template_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **template_device_uid** | **str**| The unique identifier, represented as a UUID, of the template device in Security Cloud Control. | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Template Device object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_template_devices**
> DevicePage get_template_devices(limit=limit, offset=offset, q=q, sort=sort)

Get Template Devices

Get a list of template devices in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device_page import DevicePage
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
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Template Devices
        api_response = api_instance.get_template_devices(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of InventoryApi->get_template_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->get_template_devices: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**DevicePage**](DevicePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Template Devices |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_cloud_service**
> Device modify_cloud_service(cloud_service_uid, device_patch_input)

Modify Cloud Service

Modify a Cloud Service by UID in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device import Device
from scc_firewall_manager_sdk.models.device_patch_input import DevicePatchInput
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
    cloud_service_uid = 'cloud_service_uid_example' # str | The unique identifier, represented as a UUID, of the cloud service in Security Cloud Control.
    device_patch_input = scc_firewall_manager_sdk.DevicePatchInput() # DevicePatchInput | 

    try:
        # Modify Cloud Service
        api_response = api_instance.modify_cloud_service(cloud_service_uid, device_patch_input)
        print("The response of InventoryApi->modify_cloud_service:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->modify_cloud_service: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloud_service_uid** | **str**| The unique identifier, represented as a UUID, of the cloud service in Security Cloud Control. | 
 **device_patch_input** | [**DevicePatchInput**](DevicePatchInput.md)|  | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Cloud Service |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_device**
> Device modify_device(device_uid, device_patch_input)

Modify Device

Modify a device in the Security Cloud Control tenant

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device import Device
from scc_firewall_manager_sdk.models.device_patch_input import DevicePatchInput
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
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the device in Security Cloud Control.
    device_patch_input = scc_firewall_manager_sdk.DevicePatchInput() # DevicePatchInput | 

    try:
        # Modify Device
        api_response = api_instance.modify_device(device_uid, device_patch_input)
        print("The response of InventoryApi->modify_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->modify_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the device in Security Cloud Control. | 
 **device_patch_input** | [**DevicePatchInput**](DevicePatchInput.md)|  | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Device object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_device_manager**
> Device modify_device_manager(device_manager_uid, device_manager_patch_input)

Modify Device Manager

Modify a device manager by UID in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device import Device
from scc_firewall_manager_sdk.models.device_manager_patch_input import DeviceManagerPatchInput
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
    device_manager_uid = 'device_manager_uid_example' # str | The unique identifier, represented as a UUID, of the device manager in Security Cloud Control.
    device_manager_patch_input = scc_firewall_manager_sdk.DeviceManagerPatchInput() # DeviceManagerPatchInput | 

    try:
        # Modify Device Manager
        api_response = api_instance.modify_device_manager(device_manager_uid, device_manager_patch_input)
        print("The response of InventoryApi->modify_device_manager:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->modify_device_manager: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_manager_uid** | **str**| The unique identifier, represented as a UUID, of the device manager in Security Cloud Control. | 
 **device_manager_patch_input** | [**DeviceManagerPatchInput**](DeviceManagerPatchInput.md)|  | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Device manager |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_devices**
> CdoTransaction modify_devices(devices_patch_input)

Modify Devices

Modify a list of devices in the Security Cloud Control tenant

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.devices_patch_input import DevicesPatchInput
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
    devices_patch_input = scc_firewall_manager_sdk.DevicesPatchInput() # DevicesPatchInput | 

    try:
        # Modify Devices
        api_response = api_instance.modify_devices(devices_patch_input)
        print("The response of InventoryApi->modify_devices:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->modify_devices: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **devices_patch_input** | [**DevicesPatchInput**](DevicesPatchInput.md)|  | 

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

# **modify_template_device**
> Device modify_template_device(template_device_uid, device_patch_input)

Modify Template Device

Modify a template device in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.device import Device
from scc_firewall_manager_sdk.models.device_patch_input import DevicePatchInput
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
    template_device_uid = 'template_device_uid_example' # str | The unique identifier, represented as a UUID, of the template device in Security Cloud Control.
    device_patch_input = scc_firewall_manager_sdk.DevicePatchInput() # DevicePatchInput | 

    try:
        # Modify Template Device
        api_response = api_instance.modify_template_device(template_device_uid, device_patch_input)
        print("The response of InventoryApi->modify_template_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->modify_template_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **template_device_uid** | **str**| The unique identifier, represented as a UUID, of the template device in Security Cloud Control. | 
 **device_patch_input** | [**DevicePatchInput**](DevicePatchInput.md)|  | 

### Return type

[**Device**](Device.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Template Device object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **onboard_asa_device**
> CdoTransaction onboard_asa_device(asa_create_or_update_input)

Onboard ASA device

This is an asynchronous operation to onboard an ASA to a Security Cloud Control tenant. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_create_or_update_input import AsaCreateOrUpdateInput
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
    asa_create_or_update_input = scc_firewall_manager_sdk.AsaCreateOrUpdateInput() # AsaCreateOrUpdateInput | 

    try:
        # Onboard ASA device
        api_response = api_instance.onboard_asa_device(asa_create_or_update_input)
        print("The response of InventoryApi->onboard_asa_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->onboard_asa_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asa_create_or_update_input** | [**AsaCreateOrUpdateInput**](AsaCreateOrUpdateInput.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the onboarding operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **onboard_ftd_device_using_ztp**
> CdoTransaction onboard_ftd_device_using_ztp(ztp_onboarding_input)

Onboard FTD device using Zero-Touch Provisioning

This is an asynchronous operation to onboard a cdFMC managed FTD using Zero-Touch Provisioning. The operation returns a transaction object that can be used to track the progress of the onboarding operation. Note: Zero-Touch Onboarding can be done with Secure Firewall 1xxx, 2xxx, and 3xxx Series devices. This operation will be marked as complete once CDO is ready to handle a response from the device once it is plugged in and connected to the Internet; it does not wait for the device to communicate back to Security Cloud Control.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.ztp_onboarding_input import ZtpOnboardingInput
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
    ztp_onboarding_input = scc_firewall_manager_sdk.ZtpOnboardingInput() # ZtpOnboardingInput | 

    try:
        # Onboard FTD device using Zero-Touch Provisioning
        api_response = api_instance.onboard_ftd_device_using_ztp(ztp_onboarding_input)
        print("The response of InventoryApi->onboard_ftd_device_using_ztp:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->onboard_ftd_device_using_ztp: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ztp_onboarding_input** | [**ZtpOnboardingInput**](ZtpOnboardingInput.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the creation operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **onboard_ios_device**
> CdoTransaction onboard_ios_device(ios_create_or_update_input)

Onboard IOS Device

Onboard a IOS device in the Security Cloud Control tenant. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.ios_create_or_update_input import IosCreateOrUpdateInput
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
    ios_create_or_update_input = scc_firewall_manager_sdk.IosCreateOrUpdateInput() # IosCreateOrUpdateInput | 

    try:
        # Onboard IOS Device
        api_response = api_instance.onboard_ios_device(ios_create_or_update_input)
        print("The response of InventoryApi->onboard_ios_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->onboard_ios_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ios_create_or_update_input** | [**IosCreateOrUpdateInput**](IosCreateOrUpdateInput.md)|  | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the onboarding operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **provision_cd_fmc**
> CdoTransaction provision_cd_fmc()

Provision cdFMC

This is an asynchronous operation to provision a cloud-delivered FMC in a tenant. This operation can only be performed as a super-admin user.

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

    try:
        # Provision cdFMC
        api_response = api_instance.provision_cd_fmc()
        print("The response of InventoryApi->provision_cd_fmc:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->provision_cd_fmc: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

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
**202** | Security Cloud Control Transaction object that can be used to track the status of the operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**409** | Conflict. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **read_asa_device_configuration**
> CdoTransaction read_asa_device_configuration(device_uid)

Read ASA device configuration

This is an asynchronous operation to read the latest configuration on an ASA device in to the Security Cloud Control tenant. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.

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
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the ASA device in Security Cloud Control.

    try:
        # Read ASA device configuration
        api_response = api_instance.read_asa_device_configuration(device_uid)
        print("The response of InventoryApi->read_asa_device_configuration:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->read_asa_device_configuration: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the ASA device in Security Cloud Control. | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the read operation |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **reconnect_asa_device**
> CdoTransaction reconnect_asa_device(device_uid)

Reconnect ASA device

This is an asynchronous operation to re-establish the connection between an ASA and the Security Cloud Control cloud. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.

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
    device_uid = 'device_uid_example' # str | The unique identifier, represented as a UUID, of the ASA device in Security Cloud Control.

    try:
        # Reconnect ASA device
        api_response = api_instance.reconnect_asa_device(device_uid)
        print("The response of InventoryApi->reconnect_asa_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling InventoryApi->reconnect_asa_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| The unique identifier, represented as a UUID, of the ASA device in Security Cloud Control. | 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the reconnecting operation. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

