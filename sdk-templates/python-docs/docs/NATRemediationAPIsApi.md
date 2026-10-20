# scc_firewall_manager_sdk.NATRemediationAPIsApi

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
[**apply_nat_remediation_actions**](NATRemediationAPIsApi.md#apply_nat_remediation_actions) | **POST** /v1/agenticops/policies/nat/analysis/insights/{insightUid}/remediation-plan-executions | Apply NAT policy remediation plan
[**get_nat_remediation_actions**](NATRemediationAPIsApi.md#get_nat_remediation_actions) | **GET** /v1/agenticops/policies/nat/analysis/insights/{insightUid}/remediation-plans | Get NAT policy remediation plan
[**modify_nat_remediation_actions**](NATRemediationAPIsApi.md#modify_nat_remediation_actions) | **PUT** /v1/agenticops/policies/nat/analysis/insights/{insightUid}/remediation-plans | Save NAT policy remediation plan


# **apply_nat_remediation_actions**
> CdoTransaction apply_nat_remediation_actions(insight_uid)

Apply NAT policy remediation plan

Starts an asynchronous operation to apply the saved remediation plan for the specified NAT policy analysis. The operation returns a Security Cloud Control Transaction object that can be used to track remediation progress.

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
    api_instance = scc_firewall_manager_sdk.NATRemediationAPIsApi(api_client)
    insight_uid = 'a276b855-feb7-4774-863f-569f2826f3f5' # str | The unique identifier of the NAT policy analysis.

    try:
        # Apply NAT policy remediation plan
        api_response = api_instance.apply_nat_remediation_actions(insight_uid)
        print("The response of NATRemediationAPIsApi->apply_nat_remediation_actions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NATRemediationAPIsApi->apply_nat_remediation_actions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insight_uid** | **str**| The unique identifier of the NAT policy analysis. | 

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
**202** | Security Cloud Control Transaction object that can be used to track remediation progress |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**409** | Conflict. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_nat_remediation_actions**
> RemediationActionsDto get_nat_remediation_actions(insight_uid)

Get NAT policy remediation plan

Gets the saved remediation plan for the specified NAT policy analysis. The plan includes the selected actions for shadowed and redundant rules and the total number of rules selected for deletion or disablement. If no plan has been saved, an empty plan is returned.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.remediation_actions_dto import RemediationActionsDto
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
    api_instance = scc_firewall_manager_sdk.NATRemediationAPIsApi(api_client)
    insight_uid = 'a276b855-feb7-4774-863f-569f2826f3f5' # str | The unique identifier of the NAT policy analysis.

    try:
        # Get NAT policy remediation plan
        api_response = api_instance.get_nat_remediation_actions(insight_uid)
        print("The response of NATRemediationAPIsApi->get_nat_remediation_actions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NATRemediationAPIsApi->get_nat_remediation_actions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insight_uid** | **str**| The unique identifier of the NAT policy analysis. | 

### Return type

[**RemediationActionsDto**](RemediationActionsDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | NAT policy remediation plan |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_nat_remediation_actions**
> RemediationActionsDto modify_nat_remediation_actions(insight_uid, remediation_actions_dto)

Save NAT policy remediation plan

Saves the remediation plan for the specified NAT policy analysis. The request replaces the saved selections for shadowed and redundant rules. Omitted selections are removed. The `totalDeletedCount` and `totalDisabledCount` fields are calculated from the selected rules and are ignored in the request.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.remediation_actions_dto import RemediationActionsDto
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
    api_instance = scc_firewall_manager_sdk.NATRemediationAPIsApi(api_client)
    insight_uid = 'a276b855-feb7-4774-863f-569f2826f3f5' # str | The unique identifier of the NAT policy analysis.
    remediation_actions_dto = scc_firewall_manager_sdk.RemediationActionsDto() # RemediationActionsDto | 

    try:
        # Save NAT policy remediation plan
        api_response = api_instance.modify_nat_remediation_actions(insight_uid, remediation_actions_dto)
        print("The response of NATRemediationAPIsApi->modify_nat_remediation_actions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NATRemediationAPIsApi->modify_nat_remediation_actions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insight_uid** | **str**| The unique identifier of the NAT policy analysis. | 
 **remediation_actions_dto** | [**RemediationActionsDto**](RemediationActionsDto.md)|  | 

### Return type

[**RemediationActionsDto**](RemediationActionsDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Saved NAT policy remediation plan |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**409** | Conflict. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

