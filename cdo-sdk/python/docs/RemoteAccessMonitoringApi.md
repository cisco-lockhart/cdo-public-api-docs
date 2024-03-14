# openapi_client.RemoteAccessMonitoringApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_ra_vpn_session**](RemoteAccessMonitoringApi.md#get_ra_vpn_session) | **GET** /v1/vpnsessions/{raVpnSessionUid} | Fetch a RA VPN session by UID in the CDO tenant.
[**list_ra_vpn_sessions**](RemoteAccessMonitoringApi.md#list_ra_vpn_sessions) | **GET** /v1/vpnsessions | Fetch a list of RA VPN sessions.


# **get_ra_vpn_session**
> RaVpnSession get_ra_vpn_session(ra_vpn_session_uid)

Fetch a RA VPN session by UID in the CDO tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import openapi_client
from openapi_client.models.ra_vpn_session import RaVpnSession
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.RemoteAccessMonitoringApi(api_client)
    ra_vpn_session_uid = 'ra_vpn_session_uid_example' # str | The unique identifier of the RA VPN session in CDO.

    try:
        # Fetch a RA VPN session by UID in the CDO tenant.
        api_response = api_instance.get_ra_vpn_session(ra_vpn_session_uid)
        print("The response of RemoteAccessMonitoringApi->get_ra_vpn_session:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RemoteAccessMonitoringApi->get_ra_vpn_session: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ra_vpn_session_uid** | **str**| The unique identifier of the RA VPN session in CDO. | 

### Return type

[**RaVpnSession**](RaVpnSession.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | RA VPN Session object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_ra_vpn_sessions**
> RaVpnSessionPage list_ra_vpn_sessions(sort, limit=limit, offset=offset, q=q)

Fetch a list of RA VPN sessions.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import openapi_client
from openapi_client.models.ra_vpn_session_page import RaVpnSessionPage
from openapi_client.models.sort_criteria_param import SortCriteriaParam
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.RemoteAccessMonitoringApi(api_client)
    sort = openapi_client.SortCriteriaParam() # SortCriteriaParam | The fields to sort results by.
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | The offset of the results retrieved. The CDO Public API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'name:London-Office-ASA' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)

    try:
        # Fetch a list of RA VPN sessions.
        api_response = api_instance.list_ra_vpn_sessions(sort, limit=limit, offset=offset, q=q)
        print("The response of RemoteAccessMonitoringApi->list_ra_vpn_sessions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RemoteAccessMonitoringApi->list_ra_vpn_sessions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | [**SortCriteriaParam**](.md)| The fields to sort results by. | 
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| The offset of the results retrieved. The CDO Public API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 

### Return type

[**RaVpnSessionPage**](RaVpnSessionPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of RA VPN Sessions |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

