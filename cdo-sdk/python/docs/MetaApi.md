# openapi_client.MetaApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_jwks**](MetaApi.md#get_jwks) | **GET** /.well-known/jwks.json | Fetch The JSON Web Key Set
[**get_meta**](MetaApi.md#get_meta) | **GET** /v1/meta | Get Meta information about CDO, including the IP addresses of CDO services
[**list_regions**](MetaApi.md#list_regions) | **GET** /v1/regions | Fetch a list of CDO regions.


# **get_jwks**
> JwkSet get_jwks()

Fetch The JSON Web Key Set

Retrieves the JSON Web Key Set (JWKS), which is a set of keys containing the public keys used to verify any JSON Web Token (JWT) issued by the Authorization Server and signed using the RS256 signing algorithm. Note: Verification of the token using a JSON Web Key does not guarantee validity due to the possibility of revocation.

### Example


```python
import openapi_client
from openapi_client.models.jwk_set import JwkSet
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.MetaApi(api_client)

    try:
        # Fetch The JSON Web Key Set
        api_response = api_instance.get_jwks()
        print("The response of MetaApi->get_jwks:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MetaApi->get_jwks: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**JwkSet**](JwkSet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | The JSON Web Key Set. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_meta**
> Meta get_meta()

Get Meta information about CDO, including the IP addresses of CDO services

### Example


```python
import openapi_client
from openapi_client.models.meta import Meta
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.MetaApi(api_client)

    try:
        # Get Meta information about CDO, including the IP addresses of CDO services
        api_response = api_instance.get_meta()
        print("The response of MetaApi->get_meta:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MetaApi->get_meta: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**Meta**](Meta.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Meta information about CDO. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **list_regions**
> CdoRegionList list_regions()

Fetch a list of CDO regions.

### Example


```python
import openapi_client
from openapi_client.models.cdo_region_list import CdoRegionList
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.MetaApi(api_client)

    try:
        # Fetch a list of CDO regions.
        api_response = api_instance.list_regions()
        print("The response of MetaApi->list_regions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MetaApi->list_regions: %s\n" % e)
```



### Parameters

This endpoint does not need any parameter.

### Return type

[**CdoRegionList**](CdoRegionList.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of CDO regions |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

