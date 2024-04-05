# cdo_sdk_python.ChangelogsApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_changelog**](ChangelogsApi.md#get_changelog) | **GET** /v1/changelogs/{changelogUid} | Get Change Log
[**get_changelogs**](ChangelogsApi.md#get_changelogs) | **GET** /v1/changelogs | Get Change Logs


# **get_changelog**
> Changelog get_changelog(changelog_uid)

Get Change Log

Get a specific Change Log object by UID in the CDO tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.changelog import Changelog
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
    api_instance = cdo_sdk_python.ChangelogsApi(api_client)
    changelog_uid = 'changelog_uid_example' # str | The unique identifier of the changelog in CDO.

    try:
        # Get Change Log
        api_response = api_instance.get_changelog(changelog_uid)
        print("The response of ChangelogsApi->get_changelog:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChangelogsApi->get_changelog: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changelog_uid** | **str**| The unique identifier of the changelog in CDO. | 

### Return type

[**Changelog**](Changelog.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Change Log objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_changelogs**
> ChangelogPage get_changelogs(limit=limit, offset=offset, search_text=search_text, q=q)

Get Change Logs

Get a list of Change Logs in the CDO tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.changelog_page import ChangelogPage
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
    api_instance = cdo_sdk_python.ChangelogsApi(api_client)
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    search_text = 'search_text_example' # str | The searchText parameter serves as a flexible search option that allows for text-based filtering across all fields of the Change Log object. This parameter can be used independently to search for entries containing the specified text, or in combination with the q query parameter for more targeted results. When used with q, the search conditions of searchText are logically ANDed with the q parameter's criteria, ensuring that the returned entries satisfy both sets of conditions. (optional)
    q = 'name:London-Office-ASA' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Get Change Logs
        api_response = api_instance.get_changelogs(limit=limit, offset=offset, search_text=search_text, q=q)
        print("The response of ChangelogsApi->get_changelogs:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChangelogsApi->get_changelogs: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **search_text** | **str**| The searchText parameter serves as a flexible search option that allows for text-based filtering across all fields of the Change Log object. This parameter can be used independently to search for entries containing the specified text, or in combination with the q query parameter for more targeted results. When used with q, the search conditions of searchText are logically ANDed with the q parameter&#39;s criteria, ensuring that the returned entries satisfy both sets of conditions. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 

### Return type

[**ChangelogPage**](ChangelogPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Change Log objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

