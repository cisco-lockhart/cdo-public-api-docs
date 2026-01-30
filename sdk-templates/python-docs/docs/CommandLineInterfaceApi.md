# scc_firewall_manager_sdk.CommandLineInterfaceApi

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
[**create_cli_macro**](CommandLineInterfaceApi.md#create_cli_macro) | **POST** /v1/cli/macros | Create CLI Macro
[**delete_cli_macro**](CommandLineInterfaceApi.md#delete_cli_macro) | **DELETE** /v1/cli/macros/{macroUid} | Delete CLI Macro
[**execute_cli_command**](CommandLineInterfaceApi.md#execute_cli_command) | **POST** /v1/inventory/devices/asas/cli/execute | Execute CLI Command
[**execute_cli_macro**](CommandLineInterfaceApi.md#execute_cli_macro) | **POST** /v1/inventory/devices/asas/cli/executeMacro | Execute CLI Macro Command
[**get_cli_macro**](CommandLineInterfaceApi.md#get_cli_macro) | **GET** /v1/cli/macros/{macroUid} | Get CLI Macro
[**get_cli_macros**](CommandLineInterfaceApi.md#get_cli_macros) | **GET** /v1/cli/macros | Get CLI Macros
[**get_cli_result**](CommandLineInterfaceApi.md#get_cli_result) | **GET** /v1/cli/results/{cliResultUid} | Get CLI Result
[**get_cli_results**](CommandLineInterfaceApi.md#get_cli_results) | **GET** /v1/cli/results | Get CLI Results
[**modify_cli_macro**](CommandLineInterfaceApi.md#modify_cli_macro) | **PATCH** /v1/cli/macros/{macroUid} | Modify CLI Macro


# **create_cli_macro**
> CdoCliMacro create_cli_macro(cli_macro_create_input)

Create CLI Macro

Create a CLI macro, which is a fully-formed CLI command ready to use or a template of a CLI command you can modify before you run it.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_cli_macro import CdoCliMacro
from scc_firewall_manager_sdk.models.cli_macro_create_input import CliMacroCreateInput
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
    api_instance = scc_firewall_manager_sdk.CommandLineInterfaceApi(api_client)
    cli_macro_create_input = scc_firewall_manager_sdk.CliMacroCreateInput() # CliMacroCreateInput | 

    try:
        # Create CLI Macro
        api_response = api_instance.create_cli_macro(cli_macro_create_input)
        print("The response of CommandLineInterfaceApi->create_cli_macro:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CommandLineInterfaceApi->create_cli_macro: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cli_macro_create_input** | [**CliMacroCreateInput**](CliMacroCreateInput.md)|  | 

### Return type

[**CdoCliMacro**](CdoCliMacro.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Security Cloud Control CLI macro |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_cli_macro**
> delete_cli_macro(macro_uid)

Delete CLI Macro

Delete a Security Cloud Control CLI Macro by UID.

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
    api_instance = scc_firewall_manager_sdk.CommandLineInterfaceApi(api_client)
    macro_uid = 'macro_uid_example' # str | The unique identifier, represented as a UUID, of the CDO CLI macro in Security Cloud Control.

    try:
        # Delete CLI Macro
        api_instance.delete_cli_macro(macro_uid)
    except Exception as e:
        print("Exception when calling CommandLineInterfaceApi->delete_cli_macro: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **macro_uid** | **str**| The unique identifier, represented as a UUID, of the CDO CLI macro in Security Cloud Control. | 

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
**204** | Security Cloud Control CLI macro |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **execute_cli_command**
> CdoTransaction execute_cli_command(cli_command_input=cli_command_input)

Execute CLI Command

This is an asynchronous operation to execute CLI commands on an ASA device in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.cli_command_input import CliCommandInput
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
    api_instance = scc_firewall_manager_sdk.CommandLineInterfaceApi(api_client)
    cli_command_input = scc_firewall_manager_sdk.CliCommandInput() # CliCommandInput |  (optional)

    try:
        # Execute CLI Command
        api_response = api_instance.execute_cli_command(cli_command_input=cli_command_input)
        print("The response of CommandLineInterfaceApi->execute_cli_command:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CommandLineInterfaceApi->execute_cli_command: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cli_command_input** | [**CliCommandInput**](CliCommandInput.md)|  | [optional] 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the execute CLI operation |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **execute_cli_macro**
> CdoTransaction execute_cli_macro(cli_macro_execute_input=cli_macro_execute_input)

Execute CLI Macro Command

This is an asynchronous operation to execute an CLI macro on an ASA device in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_transaction import CdoTransaction
from scc_firewall_manager_sdk.models.cli_macro_execute_input import CliMacroExecuteInput
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
    api_instance = scc_firewall_manager_sdk.CommandLineInterfaceApi(api_client)
    cli_macro_execute_input = scc_firewall_manager_sdk.CliMacroExecuteInput() # CliMacroExecuteInput |  (optional)

    try:
        # Execute CLI Macro Command
        api_response = api_instance.execute_cli_macro(cli_macro_execute_input=cli_macro_execute_input)
        print("The response of CommandLineInterfaceApi->execute_cli_macro:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CommandLineInterfaceApi->execute_cli_macro: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cli_macro_execute_input** | [**CliMacroExecuteInput**](CliMacroExecuteInput.md)|  | [optional] 

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
**202** | Security Cloud Control Transaction object that can be used to track the progress of the execute CLI operation |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cli_macro**
> CdoCliMacro get_cli_macro(macro_uid)

Get CLI Macro

Get a Security Cloud Control CLI Macro by UID. Note: This operation only returns the SCC CLI Macro. For executing the SCC CLI Macro, refer to the [Execute CLI Macro](https://developer.cisco.com/docs/cisco-security-cloud-control/execute-cli-macro-command/) command.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_cli_macro import CdoCliMacro
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
    api_instance = scc_firewall_manager_sdk.CommandLineInterfaceApi(api_client)
    macro_uid = 'macro_uid_example' # str | The unique identifier, represented as a UUID, of the CDO CLI macro in Security Cloud Control.

    try:
        # Get CLI Macro
        api_response = api_instance.get_cli_macro(macro_uid)
        print("The response of CommandLineInterfaceApi->get_cli_macro:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CommandLineInterfaceApi->get_cli_macro: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **macro_uid** | **str**| The unique identifier, represented as a UUID, of the CDO CLI macro in Security Cloud Control. | 

### Return type

[**CdoCliMacro**](CdoCliMacro.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Security Cloud Control CLI macro |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cli_macros**
> CdoCliMacroPage get_cli_macros(limit=limit, offset=offset, q=q, sort=sort)

Get CLI Macros

Get a list of CLI macros. A CLI macro is a fully-formed CLI command ready to use, or a template of a CLI command you can modify before you run it.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_cli_macro_page import CdoCliMacroPage
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
    api_instance = scc_firewall_manager_sdk.CommandLineInterfaceApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get CLI Macros
        api_response = api_instance.get_cli_macros(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of CommandLineInterfaceApi->get_cli_macros:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CommandLineInterfaceApi->get_cli_macros: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**CdoCliMacroPage**](CdoCliMacroPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Security Cloud Control CLI macros |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cli_result**
> CdoCliResult get_cli_result(cli_result_uid)

Get CLI Result

Get the result of a command executed using the Security Cloud Control Command Line Interface (CLI) by UID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_cli_result import CdoCliResult
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
    api_instance = scc_firewall_manager_sdk.CommandLineInterfaceApi(api_client)
    cli_result_uid = 'cli_result_uid_example' # str | The unique identifier, represented as a UUID, of the CDO CLI result in Security Cloud Control.

    try:
        # Get CLI Result
        api_response = api_instance.get_cli_result(cli_result_uid)
        print("The response of CommandLineInterfaceApi->get_cli_result:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CommandLineInterfaceApi->get_cli_result: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cli_result_uid** | **str**| The unique identifier, represented as a UUID, of the CDO CLI result in Security Cloud Control. | 

### Return type

[**CdoCliResult**](CdoCliResult.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Security Cloud Control CLI result |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cli_results**
> CdoCliResultPage get_cli_results(limit=limit, offset=offset, q=q, sort=sort)

Get CLI Results

Get a list of results from command lines executed using the [CDO Command Line Interface (CLI)](https://docs.defenseorchestrator.com/c-using-the-cdo-command-line-interface.html). These commands can be executed using either the Security Cloud Control UI or the API.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_cli_result_page import CdoCliResultPage
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
    api_instance = scc_firewall_manager_sdk.CommandLineInterfaceApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get CLI Results
        api_response = api_instance.get_cli_results(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of CommandLineInterfaceApi->get_cli_results:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CommandLineInterfaceApi->get_cli_results: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**CdoCliResultPage**](CdoCliResultPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Security Cloud Control CLI execution results |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_cli_macro**
> CdoCliMacro modify_cli_macro(macro_uid, cli_macro_patch_input)

Modify CLI Macro

Modify a Security Cloud Control CLI Macro by UID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.cdo_cli_macro import CdoCliMacro
from scc_firewall_manager_sdk.models.cli_macro_patch_input import CliMacroPatchInput
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
    api_instance = scc_firewall_manager_sdk.CommandLineInterfaceApi(api_client)
    macro_uid = 'macro_uid_example' # str | The unique identifier, represented as a UUID, of the CDO CLI macro in Security Cloud Control.
    cli_macro_patch_input = scc_firewall_manager_sdk.CliMacroPatchInput() # CliMacroPatchInput | 

    try:
        # Modify CLI Macro
        api_response = api_instance.modify_cli_macro(macro_uid, cli_macro_patch_input)
        print("The response of CommandLineInterfaceApi->modify_cli_macro:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CommandLineInterfaceApi->modify_cli_macro: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **macro_uid** | **str**| The unique identifier, represented as a UUID, of the CDO CLI macro in Security Cloud Control. | 
 **cli_macro_patch_input** | [**CliMacroPatchInput**](CliMacroPatchInput.md)|  | 

### Return type

[**CdoCliMacro**](CdoCliMacro.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Security Cloud Control CLI macro |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

