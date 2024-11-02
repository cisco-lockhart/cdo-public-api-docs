# cdo_sdk_python.RemoteAccessMonitoringApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_mfa_event**](RemoteAccessMonitoringApi.md#get_mfa_event) | **GET** /v1/mfaevents/{mfaEventUid} | Get MFA Event
[**get_mfa_events**](RemoteAccessMonitoringApi.md#get_mfa_events) | **GET** /v1/mfaevents | Get MFA Events
[**get_ra_vpn_session**](RemoteAccessMonitoringApi.md#get_ra_vpn_session) | **GET** /v1/vpnsessions/{raVpnSessionUid} | Get RA VPN Session
[**get_ra_vpn_sessions**](RemoteAccessMonitoringApi.md#get_ra_vpn_sessions) | **GET** /v1/vpnsessions | Get RA VPN Sessions
[**refresh_ra_vpn_sessions_by_device**](RemoteAccessMonitoringApi.md#refresh_ra_vpn_sessions_by_device) | **POST** /v1/vpnsessions/refresh | Refresh RA VPN Sessions
[**terminate_ra_vpn_sessions_by_device**](RemoteAccessMonitoringApi.md#terminate_ra_vpn_sessions_by_device) | **POST** /v1/vpnsessions/{deviceUid}/terminate | Terminate RA VPN Sessions
[**terminate_ra_vpn_sessions_by_device_and_user_name**](RemoteAccessMonitoringApi.md#terminate_ra_vpn_sessions_by_device_and_user_name) | **POST** /v1/vpnsessions/{deviceUid}/terminate/{userName} | Terminate User&#39;s RA VPN Sessions


# **get_mfa_event**
> MfaEvent get_mfa_event(mfa_event_uid)

Get MFA Event

Get a MFA event by UID in the SCC tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.mfa_event import MfaEvent
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
    api_instance = cdo_sdk_python.RemoteAccessMonitoringApi(api_client)
    mfa_event_uid = 'mfa_event_uid_example' # str | The unique identifier, represented as a UUID, of the MFA event in SCC.

    try:
        # Get MFA Event
        api_response = api_instance.get_mfa_event(mfa_event_uid)
        print("The response of RemoteAccessMonitoringApi->get_mfa_event:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RemoteAccessMonitoringApi->get_mfa_event: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mfa_event_uid** | **str**| The unique identifier, represented as a UUID, of the MFA event in SCC. | 

### Return type

[**MfaEvent**](MfaEvent.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | MFA Event object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_mfa_events**
> MfaEventPage get_mfa_events(limit=limit, offset=offset, q=q, sort=sort)

Get MFA Events

Get a list of MFA events.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.mfa_event_page import MfaEventPage
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
    api_instance = cdo_sdk_python.RemoteAccessMonitoringApi(api_client)
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | The offset of the results retrieved. The SCC API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get MFA Events
        api_response = api_instance.get_mfa_events(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of RemoteAccessMonitoringApi->get_mfa_events:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RemoteAccessMonitoringApi->get_mfa_events: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| The offset of the results retrieved. The SCC API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**MfaEventPage**](MfaEventPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of MFA events |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_ra_vpn_session**
> RaVpnSession get_ra_vpn_session(ra_vpn_session_uid)

Get RA VPN Session

Get a RA VPN session by UID in the SCC tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.ra_vpn_session import RaVpnSession
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
    api_instance = cdo_sdk_python.RemoteAccessMonitoringApi(api_client)
    ra_vpn_session_uid = 'ra_vpn_session_uid_example' # str | The unique identifier, represented as a UUID, of the RA VPN session in SCC.

    try:
        # Get RA VPN Session
        api_response = api_instance.get_ra_vpn_session(ra_vpn_session_uid)
        print("The response of RemoteAccessMonitoringApi->get_ra_vpn_session:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RemoteAccessMonitoringApi->get_ra_vpn_session: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ra_vpn_session_uid** | **str**| The unique identifier, represented as a UUID, of the RA VPN session in SCC. | 

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

# **get_ra_vpn_sessions**
> RaVpnSessionPage get_ra_vpn_sessions(limit=limit, offset=offset, q=q, sort=sort)

Get RA VPN Sessions

Get a list of RA VPN sessions.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.ra_vpn_session_page import RaVpnSessionPage
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
    api_instance = cdo_sdk_python.RemoteAccessMonitoringApi(api_client)
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | The offset of the results retrieved. The SCC API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get RA VPN Sessions
        api_response = api_instance.get_ra_vpn_sessions(limit=limit, offset=offset, q=q, sort=sort)
        print("The response of RemoteAccessMonitoringApi->get_ra_vpn_sessions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RemoteAccessMonitoringApi->get_ra_vpn_sessions: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| The offset of the results retrieved. The SCC API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

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

# **refresh_ra_vpn_sessions_by_device**
> CdoTransaction refresh_ra_vpn_sessions_by_device(ra_vpn_device_input=ra_vpn_device_input)

Refresh RA VPN Sessions

This is an asynchronous operation to refresh RA VPN sessions for all devices in the SCC tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
from cdo_sdk_python.models.ra_vpn_device_input import RaVpnDeviceInput
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
    api_instance = cdo_sdk_python.RemoteAccessMonitoringApi(api_client)
    ra_vpn_device_input = cdo_sdk_python.RaVpnDeviceInput() # RaVpnDeviceInput |  (optional)

    try:
        # Refresh RA VPN Sessions
        api_response = api_instance.refresh_ra_vpn_sessions_by_device(ra_vpn_device_input=ra_vpn_device_input)
        print("The response of RemoteAccessMonitoringApi->refresh_ra_vpn_sessions_by_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RemoteAccessMonitoringApi->refresh_ra_vpn_sessions_by_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ra_vpn_device_input** | [**RaVpnDeviceInput**](RaVpnDeviceInput.md)|  | [optional] 

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
**202** | SCC Transaction object that can be used to track the progress of the refresh operation |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **terminate_ra_vpn_sessions_by_device**
> CdoTransaction terminate_ra_vpn_sessions_by_device(device_uid)

Terminate RA VPN Sessions

This is an asynchronous operation to terminate all RA VPN sessions on a device in the SCC tenant. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
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
    api_instance = cdo_sdk_python.RemoteAccessMonitoringApi(api_client)
    device_uid = 'device_uid_example' # str | 

    try:
        # Terminate RA VPN Sessions
        api_response = api_instance.terminate_ra_vpn_sessions_by_device(device_uid)
        print("The response of RemoteAccessMonitoringApi->terminate_ra_vpn_sessions_by_device:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RemoteAccessMonitoringApi->terminate_ra_vpn_sessions_by_device: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**|  | 

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
**202** | SCC Transaction object that can be used to track the progress of the termination operation |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **terminate_ra_vpn_sessions_by_device_and_user_name**
> CdoTransaction terminate_ra_vpn_sessions_by_device_and_user_name(device_uid, user_name)

Terminate User's RA VPN Sessions

This is an asynchronous operation to terminate all of a user's RA VPN sessions on a device in the SCC tenant. This operation returns a link to a transaction object that can be used to monitor the progress of the operation.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
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
    api_instance = cdo_sdk_python.RemoteAccessMonitoringApi(api_client)
    device_uid = 'device_uid_example' # str | 
    user_name = 'user_name_example' # str | 

    try:
        # Terminate User's RA VPN Sessions
        api_response = api_instance.terminate_ra_vpn_sessions_by_device_and_user_name(device_uid, user_name)
        print("The response of RemoteAccessMonitoringApi->terminate_ra_vpn_sessions_by_device_and_user_name:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling RemoteAccessMonitoringApi->terminate_ra_vpn_sessions_by_device_and_user_name: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**|  | 
 **user_name** | **str**|  | 

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
**202** | SCC Transaction object that can be used to track the progress of the termination operation |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

