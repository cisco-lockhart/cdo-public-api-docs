# scc_firewall_manager_sdk.ASAAccessRulesApi

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
[**create_access_rule**](ASAAccessRulesApi.md#create_access_rule) | **POST** /v1/policies/asa/accessrules | Create ASA Access Rule
[**delete_access_rule**](ASAAccessRulesApi.md#delete_access_rule) | **DELETE** /v1/policies/asa/accessrules/{accessRuleUid} | Delete ASA Access Rule
[**fetch_access_rule**](ASAAccessRulesApi.md#fetch_access_rule) | **GET** /v1/policies/asa/accessrules/{accessRuleUid} | Get ASA Access Rule
[**list_access_rules**](ASAAccessRulesApi.md#list_access_rules) | **GET** /v1/policies/asa/accessrules | Get ASA Access Rules
[**modify_access_rule**](ASAAccessRulesApi.md#modify_access_rule) | **PATCH** /v1/policies/asa/accessrules/{accessRuleUid} | Modify ASA Access Rule


# **create_access_rule**
> AccessRule create_access_rule(access_rule_create_input)

Create ASA Access Rule

Create an Access Rule in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.access_rule import AccessRule
from scc_firewall_manager_sdk.models.access_rule_create_input import AccessRuleCreateInput
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
    api_instance = scc_firewall_manager_sdk.ASAAccessRulesApi(api_client)
    access_rule_create_input = scc_firewall_manager_sdk.AccessRuleCreateInput() # AccessRuleCreateInput | 

    try:
        # Create ASA Access Rule
        api_response = api_instance.create_access_rule(access_rule_create_input)
        print("The response of ASAAccessRulesApi->create_access_rule:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAAccessRulesApi->create_access_rule: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **access_rule_create_input** | [**AccessRuleCreateInput**](AccessRuleCreateInput.md)|  | 

### Return type

[**AccessRule**](AccessRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Security Cloud Control Access Rule. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_access_rule**
> delete_access_rule(access_rule_uid)

Delete ASA Access Rule

Delete Access Rule by UID in the Security Cloud Control tenant.

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
    api_instance = scc_firewall_manager_sdk.ASAAccessRulesApi(api_client)
    access_rule_uid = 'access_rule_uid_example' # str | The unique identifier, represented as a UUID, of the Access Rule in Security Cloud Control.

    try:
        # Delete ASA Access Rule
        api_instance.delete_access_rule(access_rule_uid)
    except Exception as e:
        print("Exception when calling ASAAccessRulesApi->delete_access_rule: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **access_rule_uid** | **str**| The unique identifier, represented as a UUID, of the Access Rule in Security Cloud Control. | 

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
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **fetch_access_rule**
> AccessRule fetch_access_rule(access_rule_uid)

Get ASA Access Rule

Get a single ASA Access Rule by UUID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.access_rule import AccessRule
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
    api_instance = scc_firewall_manager_sdk.ASAAccessRulesApi(api_client)
    access_rule_uid = 'access_rule_uid_example' # str | The unique identifier, represented as a UUID, of the Access Rule in Security Cloud Control.

    try:
        # Get ASA Access Rule
        api_response = api_instance.fetch_access_rule(access_rule_uid)
        print("The response of ASAAccessRulesApi->fetch_access_rule:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAAccessRulesApi->fetch_access_rule: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **access_rule_uid** | **str**| The unique identifier, represented as a UUID, of the Access Rule in Security Cloud Control. | 

### Return type

[**AccessRule**](AccessRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Access Rule object. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_access_rules**
> AccessRulePage list_access_rules(limit=limit, offset=offset, q=q, sort=sort)

Get ASA Access Rules

Get a list of ASA Access Rules.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.access_rule_page import AccessRulePage
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
    api_instance = scc_firewall_manager_sdk.ASAAccessRulesApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get ASA Access Rules
        api_response = api_instance.list_access_rules(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of ASAAccessRulesApi->list_access_rules:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAAccessRulesApi->list_access_rules: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**AccessRulePage**](AccessRulePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Access Rules. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_access_rule**
> AccessRule modify_access_rule(access_rule_uid, access_rule_update_input)

Modify ASA Access Rule

Modify an Access Rule in the Security Cloud Control tenant by UID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.access_rule import AccessRule
from scc_firewall_manager_sdk.models.access_rule_update_input import AccessRuleUpdateInput
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
    api_instance = scc_firewall_manager_sdk.ASAAccessRulesApi(api_client)
    access_rule_uid = 'access_rule_uid_example' # str | The unique identifier, represented as a UUID, of the Security Cloud Control Access Rule.
    access_rule_update_input = scc_firewall_manager_sdk.AccessRuleUpdateInput() # AccessRuleUpdateInput | 

    try:
        # Modify ASA Access Rule
        api_response = api_instance.modify_access_rule(access_rule_uid, access_rule_update_input)
        print("The response of ASAAccessRulesApi->modify_access_rule:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAAccessRulesApi->modify_access_rule: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **access_rule_uid** | **str**| The unique identifier, represented as a UUID, of the Security Cloud Control Access Rule. | 
 **access_rule_update_input** | [**AccessRuleUpdateInput**](AccessRuleUpdateInput.md)|  | 

### Return type

[**AccessRule**](AccessRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Security Cloud Control Access Rule |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

