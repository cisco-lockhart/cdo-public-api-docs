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
[**apply_nat_remediation_actions**](NATRemediationAPIsApi.md#apply_nat_remediation_actions) | **POST** /v1/agenticops/policies/nat/analysis/insights/{insightUid}/remediation-plan-executions | Create NAT Remediation Plan Execution
[**get_nat_remediation_actions**](NATRemediationAPIsApi.md#get_nat_remediation_actions) | **GET** /v1/agenticops/policies/nat/analysis/insights/{insightUid}/remediation-plans | Get NAT Remediation Plan
[**modify_nat_remediation_actions**](NATRemediationAPIsApi.md#modify_nat_remediation_actions) | **PUT** /v1/agenticops/policies/nat/analysis/insights/{insightUid}/remediation-plans | Modify NAT Remediation Plan


# **apply_nat_remediation_actions**
> CdoTransaction apply_nat_remediation_actions(insight_uid)

Create NAT Remediation Plan Execution

This is an asynchronous operation to apply the saved NAT remediation plan for one analysis insight. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.

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
    insight_uid = 'insight_uid_example' # str | 

    try:
        # Create NAT Remediation Plan Execution
        api_response = api_instance.apply_nat_remediation_actions(insight_uid)
        print("The response of NATRemediationAPIsApi->apply_nat_remediation_actions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NATRemediationAPIsApi->apply_nat_remediation_actions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insight_uid** | **str**|  | 

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
**202** | Transaction details to poll for remediation completion |  -  |
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

Get NAT Remediation Plan

Returns the saved NAT remediation plan for one analysis insight.

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
    insight_uid = 'insight_uid_example' # str | 

    try:
        # Get NAT Remediation Plan
        api_response = api_instance.get_nat_remediation_actions(insight_uid)
        print("The response of NATRemediationAPIsApi->get_nat_remediation_actions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NATRemediationAPIsApi->get_nat_remediation_actions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insight_uid** | **str**|  | 

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
**200** | Successfully retrieved remediation plan |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_nat_remediation_actions**
> RemediationActionsDto modify_nat_remediation_actions(insight_uid, remediation_actions_dto)

Modify NAT Remediation Plan

Creates or replaces a NAT remediation plan for one analysis insight. Computed count fields in the request are ignored and returned from the saved rule-level actions.

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
    insight_uid = 'insight_uid_example' # str | 
    remediation_actions_dto = scc_firewall_manager_sdk.RemediationActionsDto() # RemediationActionsDto | 

    try:
        # Modify NAT Remediation Plan
        api_response = api_instance.modify_nat_remediation_actions(insight_uid, remediation_actions_dto)
        print("The response of NATRemediationAPIsApi->modify_nat_remediation_actions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling NATRemediationAPIsApi->modify_nat_remediation_actions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **insight_uid** | **str**|  | 
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
**200** | Successfully saved remediation plan |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**409** | Conflict. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

