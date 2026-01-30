# scc_firewall_manager_sdk.ChangelogsApi

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
[**get_changelog**](ChangelogsApi.md#get_changelog) | **GET** /v1/changelogs/{changelogUid} | Get Change Log
[**get_changelogs**](ChangelogsApi.md#get_changelogs) | **GET** /v1/changelogs | Get Change Logs


# **get_changelog**
> Changelog get_changelog(changelog_uid)

Get Change Log

Get a specific Change Log object by UID in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.changelog import Changelog
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
    api_instance = scc_firewall_manager_sdk.ChangelogsApi(api_client)
    changelog_uid = 'changelog_uid_example' # str | The unique identifier, represented as a UUID, of the changelog in Security Cloud Control.

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
 **changelog_uid** | **str**| The unique identifier, represented as a UUID, of the changelog in Security Cloud Control. | 

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
> ChangelogPage get_changelogs(limit=limit, offset=offset, search_text=search_text, time_range=time_range, q=q)

Get Change Logs

Get a list of Change Logs in the Security Cloud Control tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.changelog_page import ChangelogPage
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
    api_instance = scc_firewall_manager_sdk.ChangelogsApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    search_text = 'search_text_example' # str | The searchText parameter serves as a flexible search option that allows for text-based filtering across all fields of the Change Log object. This parameter can be used independently to search for entries containing the specified text, or in combination with the q query parameter for more targeted results. When used with q, the search conditions of searchText are logically ANDed with the q parameter's criteria, ensuring that the returned entries satisfy both sets of conditions. (optional)
    time_range = 'time_range_example' # str | The time range for which to retrieve Change Logs. This parameter cannot be used in conjunction with a query on the lastEventDate field. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Get Change Logs
        api_response = api_instance.get_changelogs(limit=limit, offset=offset, search_text=search_text, time_range=time_range, q=q)
        print("The response of ChangelogsApi->get_changelogs:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ChangelogsApi->get_changelogs: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **search_text** | **str**| The searchText parameter serves as a flexible search option that allows for text-based filtering across all fields of the Change Log object. This parameter can be used independently to search for entries containing the specified text, or in combination with the q query parameter for more targeted results. When used with q, the search conditions of searchText are logically ANDed with the q parameter&#39;s criteria, ensuring that the returned entries satisfy both sets of conditions. | [optional] 
 **time_range** | **str**| The time range for which to retrieve Change Logs. This parameter cannot be used in conjunction with a query on the lastEventDate field. | [optional] 
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

