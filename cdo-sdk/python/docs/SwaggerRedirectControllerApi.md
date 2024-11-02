# cdo_sdk_python.SwaggerRedirectControllerApi

All URIs are relative to *https://edge.us.SCC.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**redirect_with_using_redirect_view**](SwaggerRedirectControllerApi.md#redirect_with_using_redirect_view) | **GET** /swagger-ui/** | 


# **redirect_with_using_redirect_view**
> RedirectView redirect_with_using_redirect_view()



### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.redirect_view import RedirectView
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.SCC.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.SCC.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.SwaggerRedirectControllerApi(api_client)

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

