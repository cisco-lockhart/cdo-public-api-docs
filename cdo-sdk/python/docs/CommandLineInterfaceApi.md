# cdo_sdk_python.CommandLineInterfaceApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**create_cli_macro**](CommandLineInterfaceApi.md#create_cli_macro) | **POST** /v1/cli/macros | Create CLI Macro
[**delete_cli_macro**](CommandLineInterfaceApi.md#delete_cli_macro) | **DELETE** /v1/cli/macros/{macroUid} | Delete CLI Macro
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
import cdo_sdk_python
from cdo_sdk_python.models.cdo_cli_macro import CdoCliMacro
from cdo_sdk_python.models.cli_macro_create_input import CliMacroCreateInput
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.CommandLineInterfaceApi(api_client)
    cli_macro_create_input = cdo_sdk_python.CliMacroCreateInput() # CliMacroCreateInput | 

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
**200** | CDO CLI macro |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_cli_macro**
> delete_cli_macro(macro_uid)

Delete CLI Macro

Delete a CDO CLI Macro by UID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.CommandLineInterfaceApi(api_client)
    macro_uid = 'macro_uid_example' # str | The unique identifier, represented as a UUID, of the CDO CLI macro in CDO.

    try:
        # Delete CLI Macro
        api_instance.delete_cli_macro(macro_uid)
    except Exception as e:
        print("Exception when calling CommandLineInterfaceApi->delete_cli_macro: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **macro_uid** | **str**| The unique identifier, represented as a UUID, of the CDO CLI macro in CDO. | 

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
**204** | CDO CLI macro |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cli_macro**
> CdoCliMacro get_cli_macro(macro_uid)

Get CLI Macro

Get a CDO CLI Macro by UID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_cli_macro import CdoCliMacro
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.CommandLineInterfaceApi(api_client)
    macro_uid = 'macro_uid_example' # str | The unique identifier, represented as a UUID, of the CDO CLI macro in CDO.

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
 **macro_uid** | **str**| The unique identifier, represented as a UUID, of the CDO CLI macro in CDO. | 

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
**200** | CDO CLI macro |  -  |
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
import cdo_sdk_python
from cdo_sdk_python.models.cdo_cli_macro_page import CdoCliMacroPage
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.CommandLineInterfaceApi(api_client)
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
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
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
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
**200** | List of CDO CLI macros |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_cli_result**
> CdoCliResult get_cli_result(cli_result_uid)

Get CLI Result

Get the result of a command executed using the CDO Command Line Interface (CLI) by UID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_cli_result import CdoCliResult
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.CommandLineInterfaceApi(api_client)
    cli_result_uid = 'cli_result_uid_example' # str | The unique identifier, represented as a UUID, of the CDO CLI result in CDO.

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
 **cli_result_uid** | **str**| The unique identifier, represented as a UUID, of the CDO CLI result in CDO. | 

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
**200** | CDO CLI result |  -  |
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

Get a list of results from command lines executed using the [CDO Command Line Interface (CLI)](https://docs.defenseorchestrator.com/c-using-the-cdo-command-line-interface.html). These commands can be executed using either the CDO UI or the API.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_cli_result_page import CdoCliResultPage
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.CommandLineInterfaceApi(api_client)
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
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
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
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
**200** | List of CDO CLI execution results |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_cli_macro**
> CdoCliMacro modify_cli_macro(macro_uid, cli_macro_patch_input)

Modify CLI Macro

Modify a CDO CLI Macro by UID.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_cli_macro import CdoCliMacro
from cdo_sdk_python.models.cli_macro_patch_input import CliMacroPatchInput
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.CommandLineInterfaceApi(api_client)
    macro_uid = 'macro_uid_example' # str | The unique identifier, represented as a UUID, of the CDO CLI macro in CDO.
    cli_macro_patch_input = cdo_sdk_python.CliMacroPatchInput() # CliMacroPatchInput | 

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
 **macro_uid** | **str**| The unique identifier, represented as a UUID, of the CDO CLI macro in CDO. | 
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
**200** | CDO CLI macro |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

