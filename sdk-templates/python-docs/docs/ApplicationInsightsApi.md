# scc_firewall_manager_sdk.ApplicationInsightsApi

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
[**get_application_outages**](ApplicationInsightsApi.md#get_application_outages) | **GET** /v1/application-insights/outages | Get Application Outages


# **get_application_outages**
> PublicThousandEyesOutageDto get_application_outages(start=start, end=end, time_range=time_range, application_name=application_name)

Get Application Outages

Get a list of application outages that started in the specified time period

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.public_thousand_eyes_outage_dto import PublicThousandEyesOutageDto
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
    api_instance = scc_firewall_manager_sdk.ApplicationInsightsApi(api_client)
    start = '2025-01-01T00:00:00Z' # datetime | Start time for outage (UTC; represented using the RFC-3339 standard, inclusive) (optional)
    end = '2025-01-31T23:59:59Z' # datetime | End time for outage (UTC; represented using the RFC-3339 standard, exclusive) (optional)
    time_range = '24h' # str | Relative time range (mutually exclusive with start/end). (optional)
    application_name = 'Webex' # str | Filter by application name (optional)

    try:
        # Get Application Outages
        api_response = api_instance.get_application_outages(start=start, end=end, time_range=time_range, application_name=application_name)
        print("The response of ApplicationInsightsApi->get_application_outages:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ApplicationInsightsApi->get_application_outages: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **datetime**| Start time for outage (UTC; represented using the RFC-3339 standard, inclusive) | [optional] 
 **end** | **datetime**| End time for outage (UTC; represented using the RFC-3339 standard, exclusive) | [optional] 
 **time_range** | **str**| Relative time range (mutually exclusive with start/end). | [optional] 
 **application_name** | **str**| Filter by application name | [optional] 

### Return type

[**PublicThousandEyesOutageDto**](PublicThousandEyesOutageDto.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of application outages that started in the specified time period |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**405** | Method not allowed. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

