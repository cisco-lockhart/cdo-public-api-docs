# cdo_sdk_python.MetaApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**get_jwks**](MetaApi.md#get_jwks) | **GET** /.well-known/jwks.json | Fetch JSON Web Key Set
[**get_meta**](MetaApi.md#get_meta) | **GET** /v1/meta | Get Meta information
[**get_regions**](MetaApi.md#get_regions) | **GET** /v1/regions | Get SCC Regions


# **get_jwks**
> JwkSet get_jwks()

Fetch JSON Web Key Set

Retrieves the JSON Web Key Set (JWKS), which is a set of keys containing the public keys used to verify any JSON Web Token (JWT) issued by the Authorization Server and signed using the RS256 signing algorithm. Note: Verification of the token using a JSON Web Key does not guarantee validity due to the possibility of revocation.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.jwk_set import JwkSet
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MetaApi(api_client)

    try:
        # Fetch JSON Web Key Set
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

Get Meta information

Get Meta information about CDO, including the IP addresses of SCC services.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.meta import Meta
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MetaApi(api_client)

    try:
        # Get Meta information
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
**200** | Meta information about SCC. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_regions**
> CdoRegionList get_regions()

Get SCC Regions

Get the list of regions that SCC is deployed in.

### Example


```python
import cdo_sdk_python
from cdo_sdk_python.models.cdo_region_list import CdoRegionList
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)


# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.MetaApi(api_client)

    try:
        # Get SCC Regions
        api_response = api_instance.get_regions()
        print("The response of MetaApi->get_regions:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling MetaApi->get_regions: %s\n" % e)
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
**200** | List of SCC regions |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

