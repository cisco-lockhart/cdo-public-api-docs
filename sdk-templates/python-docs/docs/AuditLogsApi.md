# scc_firewall_manager_sdk.AuditLogsApi

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
[**get_audit_logs**](AuditLogsApi.md#get_audit_logs) | **GET** /v1/auditlogs | Get Audit Logs


# **get_audit_logs**
> AuditLogPage get_audit_logs(limit=limit, offset=offset, search_text=search_text, time_range=time_range, q=q, sort=sort)

Get Audit Logs

Get a list of Audit Logs.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.audit_log_page import AuditLogPage
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
    api_instance = scc_firewall_manager_sdk.AuditLogsApi(api_client)
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    search_text = 'search_text_example' # str | The searchText parameter serves as a flexible search option that allows for text-based filtering across the username fields of the Audit Log object. This parameter can be used independently to search for entries containing the specified text, or in combination with the q query parameter for more targeted results. When used with q, the search conditions of searchText are logically ANDed with the q parameter's criteria, ensuring that the returned entries satisfy both sets of conditions. (optional)
    time_range = 'time_range_example' # str | The time range for which to retrieve Audit Logs. This parameter cannot be used in conjunction with a query on the eventTime field. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Audit Logs
        api_response = api_instance.get_audit_logs(limit=limit, offset=offset, search_text=search_text, time_range=time_range, q=q, sort=sort)
        print("The response of AuditLogsApi->get_audit_logs:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AuditLogsApi->get_audit_logs: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **search_text** | **str**| The searchText parameter serves as a flexible search option that allows for text-based filtering across the username fields of the Audit Log object. This parameter can be used independently to search for entries containing the specified text, or in combination with the q query parameter for more targeted results. When used with q, the search conditions of searchText are logically ANDed with the q parameter&#39;s criteria, ensuring that the returned entries satisfy both sets of conditions. | [optional] 
 **time_range** | **str**| The time range for which to retrieve Audit Logs. This parameter cannot be used in conjunction with a query on the eventTime field. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**AuditLogPage**](AuditLogPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Audit Logs objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

