# scc_firewall_manager_sdk.SwaggerRedirectControllerApi

All URIs are relative to *https://api.us.security.cisco.com/firewall*

Method | HTTP request | Description
------------- | ------------- | -------------
[**redirect_with_using_redirect_view**](SwaggerRedirectControllerApi.md#redirect_with_using_redirect_view) | **GET** /swagger-ui/** | 


# **redirect_with_using_redirect_view**
> RedirectView redirect_with_using_redirect_view()



### Example


```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.redirect_view import RedirectView
from scc_firewall_manager_sdk.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.us.security.cisco.com/firewall
# See configuration.py for a list of all supported configuration parameters.
configuration = scc_firewall_manager_sdk.Configuration(
    host = "https://api.us.security.cisco.com/firewall"
)


# Enter a context with an instance of the API client
with scc_firewall_manager_sdk.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = scc_firewall_manager_sdk.SwaggerRedirectControllerApi(api_client)

    try:
        api_response = api_instance.redirect_with_using_redirect_view()
        print("The response of SwaggerRedirectControllerApi->redirect_with_using_redirect_view:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling SwaggerRedirectControllerApi->redirect_with_using_redirect_view: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**RedirectView**](RedirectView.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | OK |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

