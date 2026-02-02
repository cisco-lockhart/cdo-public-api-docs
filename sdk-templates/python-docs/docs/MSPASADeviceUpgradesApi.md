# scc_firewall_manager_sdk.MSPASADeviceUpgradesApi

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
[**get_msp_asa_compatible_upgrade_versions**](MSPASADeviceUpgradesApi.md#get_msp_asa_compatible_upgrade_versions) | **GET** /v1/msp/inventory/devices/asas/upgrades/versions | Get compatible ASA upgrade versions


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
    api_instance = scc_firewall_manager_sdk.MSPASADeviceUpgradesApi(api_client)
    device_uids = ['device_uids_example'] # List[str] | A list of unique identifiers, represented as UUIDs, of the devices in Security Cloud Control. Note: All the specified devices must be in tenants managed by the MSP portal.

    try:
        # Get compatible ASA upgrade versions
        api_response = api_instance.get_msp_asa_compatible_upgrade_versions(device_uids)
        print("The response of MSPASADeviceUpgradesApi->get_msp_asa_compatible_upgrade_versions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MSPASADeviceUpgradesApi->get_msp_asa_compatible_upgrade_versions: %s\n" % e)
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

