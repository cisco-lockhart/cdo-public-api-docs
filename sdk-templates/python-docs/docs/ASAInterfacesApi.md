# scc_firewall_manager_sdk.ASAInterfacesApi

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
[**create_asa_ether_channel_interface**](ASAInterfacesApi.md#create_asa_ether_channel_interface) | **POST** /v1/inventory/devices/asas/{deviceUid}/etherchannelinterfaces | Create ASA EtherChannel interface
[**create_asa_sub_interface**](ASAInterfacesApi.md#create_asa_sub_interface) | **POST** /v1/inventory/devices/asas/{deviceUid}/interfaces/{interfaceUid}/subinterfaces | Create ASA subinterface
[**create_asa_vlan_interface**](ASAInterfacesApi.md#create_asa_vlan_interface) | **POST** /v1/inventory/devices/asas/{deviceUid}/vlaninterfaces | Create ASA VLAN interface
[**delete_asa_ether_channel_interface**](ASAInterfacesApi.md#delete_asa_ether_channel_interface) | **DELETE** /v1/inventory/devices/asas/{deviceUid}/etherchannelinterfaces/{interfaceUid} | Delete ASA EtherChannel interface
[**delete_asa_sub_interface**](ASAInterfacesApi.md#delete_asa_sub_interface) | **DELETE** /v1/inventory/devices/asas/{deviceUid}/subinterfaces/{interfaceUid} | Delete ASA subinterface
[**delete_asa_vlan_interface**](ASAInterfacesApi.md#delete_asa_vlan_interface) | **DELETE** /v1/inventory/devices/asas/{deviceUid}/vlaninterfaces/{interfaceUid} | Delete ASA VLAN interface
[**get_asa_ether_channel_interface**](ASAInterfacesApi.md#get_asa_ether_channel_interface) | **GET** /v1/inventory/devices/asas/{deviceUid}/etherchannelinterfaces/{interfaceUid} | Get ASA EtherChannel interface
[**get_asa_ether_channel_interfaces**](ASAInterfacesApi.md#get_asa_ether_channel_interfaces) | **GET** /v1/inventory/devices/asas/{deviceUid}/etherchannelinterfaces | Get ASA EtherChannel interfaces
[**get_asa_interface_sub_interfaces**](ASAInterfacesApi.md#get_asa_interface_sub_interfaces) | **GET** /v1/inventory/devices/asas/{deviceUid}/interfaces/{interfaceUid}/subinterfaces | Get an ASA interface’s subinterfaces
[**get_asa_physical_interface**](ASAInterfacesApi.md#get_asa_physical_interface) | **GET** /v1/inventory/devices/asas/{deviceUid}/physicalinterfaces/{interfaceUid} | Get an ASA physical interface
[**get_asa_physical_interfaces**](ASAInterfacesApi.md#get_asa_physical_interfaces) | **GET** /v1/inventory/devices/asas/{deviceUid}/physicalinterfaces | Get ASA physical interfaces
[**get_asa_sub_interface**](ASAInterfacesApi.md#get_asa_sub_interface) | **GET** /v1/inventory/devices/asas/{deviceUid}/subinterfaces/{interfaceUid} | Get an ASA subinterface
[**get_asa_sub_interfaces**](ASAInterfacesApi.md#get_asa_sub_interfaces) | **GET** /v1/inventory/devices/asas/{deviceUid}/subinterfaces | Get ASA subinterfaces
[**get_asa_virtual_tunnel_interface**](ASAInterfacesApi.md#get_asa_virtual_tunnel_interface) | **GET** /v1/inventory/devices/asas/{deviceUid}/virtualtunnelinterfaces/{interfaceUid} | Get an ASA virtual tunnel interface
[**get_asa_virtual_tunnel_interfaces**](ASAInterfacesApi.md#get_asa_virtual_tunnel_interfaces) | **GET** /v1/inventory/devices/asas/{deviceUid}/virtualtunnelinterfaces | Get ASA virtual tunnel interfaces
[**get_asa_vlan_interface**](ASAInterfacesApi.md#get_asa_vlan_interface) | **GET** /v1/inventory/devices/asas/{deviceUid}/vlaninterfaces/{interfaceUid} | Get an ASA VLAN interface
[**get_asa_vlan_interfaces**](ASAInterfacesApi.md#get_asa_vlan_interfaces) | **GET** /v1/inventory/devices/asas/{deviceUid}/vlaninterfaces | Get ASA VLAN interfaces
[**modify_asa_ether_channel_interface**](ASAInterfacesApi.md#modify_asa_ether_channel_interface) | **PATCH** /v1/inventory/devices/asas/{deviceUid}/etherchannelinterfaces/{interfaceUid} | Modify ASA EtherChannel interface
[**modify_asa_physical_interface**](ASAInterfacesApi.md#modify_asa_physical_interface) | **PATCH** /v1/inventory/devices/asas/{deviceUid}/physicalinterfaces/{interfaceUid} | Modify ASA Physical interface
[**modify_asa_sub_interface**](ASAInterfacesApi.md#modify_asa_sub_interface) | **PATCH** /v1/inventory/devices/asas/{deviceUid}/subinterfaces/{interfaceUid} | Modify ASA subinterface
[**modify_asa_vlan_interface**](ASAInterfacesApi.md#modify_asa_vlan_interface) | **PATCH** /v1/inventory/devices/asas/{deviceUid}/vlaninterfaces/{interfaceUid} | Modify ASA VLAN interface


# **create_asa_ether_channel_interface**
> AsaInterface create_asa_ether_channel_interface(device_uid, ether_channel_interface_create_input)

Create ASA EtherChannel interface

Create an ASA EtherChannel interface. An EtherChannel interface is a logical port-channel interface that aggregates multiple physical Ethernet links into a single logical interface to increase bandwidth and provide redundancy. It is configured and used like a physical interface for interface-related features, and supports up to 48 EtherChannels, depending on the ASA model.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface import AsaInterface
from scc_firewall_manager_sdk.models.ether_channel_interface_create_input import EtherChannelInterfaceCreateInput
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    ether_channel_interface_create_input = scc_firewall_manager_sdk.EtherChannelInterfaceCreateInput() # EtherChannelInterfaceCreateInput | 

    try:
        # Create ASA EtherChannel interface
        api_response = api_instance.create_asa_ether_channel_interface(device_uid, ether_channel_interface_create_input)
        print("The response of ASAInterfacesApi->create_asa_ether_channel_interface:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->create_asa_ether_channel_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **ether_channel_interface_create_input** | [**EtherChannelInterfaceCreateInput**](EtherChannelInterfaceCreateInput.md)|  | 

### Return type

[**AsaInterface**](AsaInterface.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | AsaInterface object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_asa_sub_interface**
> AsaInterface create_asa_sub_interface(device_uid, interface_uid, sub_interface_create_input)

Create ASA subinterface

Create an ASA subinterface. A subinterface is a logical interface created in a physical interface that allows the ASA to support multiple VLANs by assigning each subinterface to a different VLAN, enabling segmented network traffic over a single physical port.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface import AsaInterface
from scc_firewall_manager_sdk.models.sub_interface_create_input import SubInterfaceCreateInput
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    interface_uid = 'interface_uid_example' # str | Unique identifier (UUID) of the ASA parent interface.
    sub_interface_create_input = scc_firewall_manager_sdk.SubInterfaceCreateInput() # SubInterfaceCreateInput | 

    try:
        # Create ASA subinterface
        api_response = api_instance.create_asa_sub_interface(device_uid, interface_uid, sub_interface_create_input)
        print("The response of ASAInterfacesApi->create_asa_sub_interface:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->create_asa_sub_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **interface_uid** | **str**| Unique identifier (UUID) of the ASA parent interface. | 
 **sub_interface_create_input** | [**SubInterfaceCreateInput**](SubInterfaceCreateInput.md)|  | 

### Return type

[**AsaInterface**](AsaInterface.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | AsaInterface object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **create_asa_vlan_interface**
> AsaInterface create_asa_vlan_interface(device_uid, vlan_interface_create_input)

Create ASA VLAN interface

Create an ASA VLAN interface. A VLAN interface is a logical interface that represents a virtual LAN in a physical interface, allowing multiple VLANs to be configured and segmented in a single physical interface for traffic separation and security purposes.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface import AsaInterface
from scc_firewall_manager_sdk.models.vlan_interface_create_input import VlanInterfaceCreateInput
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    vlan_interface_create_input = scc_firewall_manager_sdk.VlanInterfaceCreateInput() # VlanInterfaceCreateInput | 

    try:
        # Create ASA VLAN interface
        api_response = api_instance.create_asa_vlan_interface(device_uid, vlan_interface_create_input)
        print("The response of ASAInterfacesApi->create_asa_vlan_interface:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->create_asa_vlan_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **vlan_interface_create_input** | [**VlanInterfaceCreateInput**](VlanInterfaceCreateInput.md)|  | 

### Return type

[**AsaInterface**](AsaInterface.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | AsaInterface object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_asa_ether_channel_interface**
> delete_asa_ether_channel_interface(device_uid, interface_uid)

Delete ASA EtherChannel interface

Delete the ASA EtherChannel interface. An EtherChannel interface is a logical port-channel interface that aggregates multiple physical Ethernet links into a single logical interface to increase bandwidth and provide redundancy. It is configured and used like a physical interface for interface-related features and supports up to 48 EtherChannels, depending on the ASA model.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    interface_uid = 'interface_uid_example' # str | Unique identifier (UUID) of the ASA EtherChannel interface.

    try:
        # Delete ASA EtherChannel interface
        api_instance.delete_asa_ether_channel_interface(device_uid, interface_uid)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->delete_asa_ether_channel_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **interface_uid** | **str**| Unique identifier (UUID) of the ASA EtherChannel interface. | 

### Return type

void (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | No Content |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_asa_sub_interface**
> delete_asa_sub_interface(device_uid, interface_uid)

Delete ASA subinterface

Delete the ASA subinterface. A subinterface is a logical interface created in a physical interface allowing the ASA to support multiple VLANs by assigning each subinterface to a different VLAN, enabling segmented network traffic over a single physical port.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    interface_uid = 'interface_uid_example' # str | Unique identifier (UUID) of the ASA subinterface.

    try:
        # Delete ASA subinterface
        api_instance.delete_asa_sub_interface(device_uid, interface_uid)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->delete_asa_sub_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **interface_uid** | **str**| Unique identifier (UUID) of the ASA subinterface. | 

### Return type

void (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | No Content |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **delete_asa_vlan_interface**
> delete_asa_vlan_interface(device_uid, interface_uid)

Delete ASA VLAN interface

Delete the ASA VLAN interface. A VLAN interface is a logical interface that represents a virtual LAN in a physical interface, allowing multiple VLANs to be configured and segmented in a single physical interface for traffic separation and security.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    interface_uid = 'interface_uid_example' # str | Unique identifier (UUID) of the ASA VLAN interface.

    try:
        # Delete ASA VLAN interface
        api_instance.delete_asa_vlan_interface(device_uid, interface_uid)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->delete_asa_vlan_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **interface_uid** | **str**| Unique identifier (UUID) of the ASA VLAN interface. | 

### Return type

void (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**204** | No Content |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_ether_channel_interface**
> AsaInterface get_asa_ether_channel_interface(device_uid, interface_uid)

Get ASA EtherChannel interface

Get an ASA EtherChannel interface using its UUID. An EtherChannel interface is a logical port-channel interface that aggregates multiple physical Ethernet links into a single logical interface to increase bandwidth and provide redundancy. It is configured and used like a physical interface for interface-related features, and supports up to 48 EtherChannels, depending on the ASA model.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface import AsaInterface
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    interface_uid = 'interface_uid_example' # str | Unique identifier (UUID) of the ASA EtherChannel interface.

    try:
        # Get ASA EtherChannel interface
        api_response = api_instance.get_asa_ether_channel_interface(device_uid, interface_uid)
        print("The response of ASAInterfacesApi->get_asa_ether_channel_interface:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->get_asa_ether_channel_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **interface_uid** | **str**| Unique identifier (UUID) of the ASA EtherChannel interface. | 

### Return type

[**AsaInterface**](AsaInterface.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | AsaInterface object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_ether_channel_interfaces**
> AsaInterfacePage get_asa_ether_channel_interfaces(device_uid, limit=limit, offset=offset, q=q, sort=sort)

Get ASA EtherChannel interfaces

Get a list of ASA EtherChannel interfaces. An EtherChannel interface is a logical port-channel interface that aggregates multiple physical Ethernet links into a single logical interface to increase bandwidth and provide redundancy. It is configured and used like a physical interface for interface-related features, and supports up to 48 EtherChannels, depending on the ASA model.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface_page import AsaInterfacePage
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get ASA EtherChannel interfaces
        api_response = api_instance.get_asa_ether_channel_interfaces(device_uid, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of ASAInterfacesApi->get_asa_ether_channel_interfaces:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->get_asa_ether_channel_interfaces: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**AsaInterfacePage**](AsaInterfacePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of AsaInterface objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_interface_sub_interfaces**
> AsaInterfacePage get_asa_interface_sub_interfaces(device_uid, interface_uid, limit=limit, offset=offset, q=q, sort=sort)

Get an ASA interface’s subinterfaces

Get a list of an ASA interface’s subinterfaces. A subinterface is a logical interface created in a physical interface that allows the ASA to support multiple VLANs by assigning each subinterface to a different VLAN, enabling segmented network traffic over a single physical port.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface_page import AsaInterfacePage
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    interface_uid = 'interface_uid_example' # str | Unique identifier (UUID) of the ASA parent interface.
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get an ASA interface’s subinterfaces
        api_response = api_instance.get_asa_interface_sub_interfaces(device_uid, interface_uid, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of ASAInterfacesApi->get_asa_interface_sub_interfaces:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->get_asa_interface_sub_interfaces: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **interface_uid** | **str**| Unique identifier (UUID) of the ASA parent interface. | 
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**AsaInterfacePage**](AsaInterfacePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of AsaInterface objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_physical_interface**
> AsaInterface get_asa_physical_interface(device_uid, interface_uid)

Get an ASA physical interface

Get an ASA physical interface using its UUID. A physical interface is a hardware port in an ASA device that can be enabled and configured to connect the firewall to a network.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface import AsaInterface
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    interface_uid = 'interface_uid_example' # str | Unique identifier (UUID) of the ASA physical interface.

    try:
        # Get an ASA physical interface
        api_response = api_instance.get_asa_physical_interface(device_uid, interface_uid)
        print("The response of ASAInterfacesApi->get_asa_physical_interface:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->get_asa_physical_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **interface_uid** | **str**| Unique identifier (UUID) of the ASA physical interface. | 

### Return type

[**AsaInterface**](AsaInterface.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | AsaInterface object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_physical_interfaces**
> AsaInterfacePage get_asa_physical_interfaces(device_uid, limit=limit, offset=offset, q=q, sort=sort)

Get ASA physical interfaces

Get a list of ASA physical interfaces. A physical interface is a hardware port in an ASA device, which can be enabled and configured to connect the firewall to a network.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface_page import AsaInterfacePage
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get ASA physical interfaces
        api_response = api_instance.get_asa_physical_interfaces(device_uid, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of ASAInterfacesApi->get_asa_physical_interfaces:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->get_asa_physical_interfaces: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**AsaInterfacePage**](AsaInterfacePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of AsaInterface objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_sub_interface**
> AsaInterface get_asa_sub_interface(device_uid, interface_uid)

Get an ASA subinterface

Get an ASA subinterface using its UUID. A subinterface is a logical interface created in a physical interface that allows the ASA to support multiple VLANs by assigning each subinterface to a different VLAN, enabling segmented network traffic over a single physical port.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface import AsaInterface
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    interface_uid = 'interface_uid_example' # str | Unique identifier (UUID) of the ASA subinterface.

    try:
        # Get an ASA subinterface
        api_response = api_instance.get_asa_sub_interface(device_uid, interface_uid)
        print("The response of ASAInterfacesApi->get_asa_sub_interface:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->get_asa_sub_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **interface_uid** | **str**| Unique identifier (UUID) of the ASA subinterface. | 

### Return type

[**AsaInterface**](AsaInterface.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | AsaInterface object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_sub_interfaces**
> AsaInterfacePage get_asa_sub_interfaces(device_uid, limit=limit, offset=offset, q=q, sort=sort)

Get ASA subinterfaces

Get a list of the ASA subinterfaces. A subinterface is a logical interface created in a physical interface that allows the ASA to support multiple VLANs by assigning each subinterface to a different VLAN, enabling segmented network traffic over a single physical port.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface_page import AsaInterfacePage
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get ASA subinterfaces
        api_response = api_instance.get_asa_sub_interfaces(device_uid, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of ASAInterfacesApi->get_asa_sub_interfaces:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->get_asa_sub_interfaces: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**AsaInterfacePage**](AsaInterfacePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of AsaInterface objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_virtual_tunnel_interface**
> AsaInterface get_asa_virtual_tunnel_interface(device_uid, interface_uid)

Get an ASA virtual tunnel interface

Get an ASA virtual tunnel interface using its UUID. A virtual tunnel interface (VTI) is a logical interface that provides a point-to-point GRE or IPsec tunnel endpoint, enabling simplified VPN configuration and routing integration in the ASA device.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface import AsaInterface
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    interface_uid = 'interface_uid_example' # str | Unique identifier (UUID) of the ASA virtual tunnel interface.

    try:
        # Get an ASA virtual tunnel interface
        api_response = api_instance.get_asa_virtual_tunnel_interface(device_uid, interface_uid)
        print("The response of ASAInterfacesApi->get_asa_virtual_tunnel_interface:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->get_asa_virtual_tunnel_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **interface_uid** | **str**| Unique identifier (UUID) of the ASA virtual tunnel interface. | 

### Return type

[**AsaInterface**](AsaInterface.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | AsaInterface object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_virtual_tunnel_interfaces**
> AsaInterfacePage get_asa_virtual_tunnel_interfaces(device_uid, limit=limit, offset=offset, q=q, sort=sort)

Get ASA virtual tunnel interfaces

Get a list of ASA virtual tunnel interfaces. A virtual tunnel interface (VTI) is a logical interface that provides a point-to-point GRE or IPsec tunnel endpoint, enabling simplified VPN configuration and routing integration in the ASA device.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface_page import AsaInterfacePage
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get ASA virtual tunnel interfaces
        api_response = api_instance.get_asa_virtual_tunnel_interfaces(device_uid, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of ASAInterfacesApi->get_asa_virtual_tunnel_interfaces:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->get_asa_virtual_tunnel_interfaces: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**AsaInterfacePage**](AsaInterfacePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of AsaInterface objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_vlan_interface**
> AsaInterface get_asa_vlan_interface(device_uid, interface_uid)

Get an ASA VLAN interface

Get an ASA VLAN interface using its UUID. A VLAN interface is a logical interface that represents a virtual LAN in a physical interface, allowing multiple VLANs to be configured and segmented in a single physical interface for traffic separation and security purposes.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface import AsaInterface
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    interface_uid = 'interface_uid_example' # str | Unique identifier (UUID) of the ASA VLAN interface.

    try:
        # Get an ASA VLAN interface
        api_response = api_instance.get_asa_vlan_interface(device_uid, interface_uid)
        print("The response of ASAInterfacesApi->get_asa_vlan_interface:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->get_asa_vlan_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **interface_uid** | **str**| Unique identifier (UUID) of the ASA VLAN interface. | 

### Return type

[**AsaInterface**](AsaInterface.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | AsaInterface object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_asa_vlan_interfaces**
> AsaInterfacePage get_asa_vlan_interfaces(device_uid, limit=limit, offset=offset, q=q, sort=sort)

Get ASA VLAN interfaces

Get a list of ASA VLAN interfaces. A VLAN interface is a logical interface that represents a virtual LAN in a physical interface, allowing multiple VLANs to be configured and segmented in a single physical interface for traffic separation and security purposes.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface_page import AsaInterfacePage
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    limit = 'limit_example' # str | Number of results to retrieve. (optional)
    offset = 'offset_example' # str | Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional)
    q = 'fieldName:fieldValue' # str | The query to execute. Use the Lucene Query Syntax to construct your query. (optional)
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get ASA VLAN interfaces
        api_response = api_instance.get_asa_vlan_interfaces(device_uid, limit=limit, offset=offset, q=q, sort=sort)
        print("The response of ASAInterfacesApi->get_asa_vlan_interfaces:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->get_asa_vlan_interfaces: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **limit** | **str**| Number of results to retrieve. | [optional] 
 **offset** | **str**| Offset of the results retrieved. The Security Cloud Control APIs use the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 
 **q** | **str**| The query to execute. Use the Lucene Query Syntax to construct your query. | [optional] 
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**AsaInterfacePage**](AsaInterfacePage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of AsaInterface objects |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_asa_ether_channel_interface**
> AsaInterface modify_asa_ether_channel_interface(device_uid, interface_uid, ether_channel_interface_patch_input)

Modify ASA EtherChannel interface

Modify the ASA EtherChannel interface. An EtherChannel interface is a logical port-channel interface that aggregates multiple physical Ethernet links into a single logical interface to increase bandwidth and provide redundancy. It is configured and used like a physical interface for interface-related features, and supports up to 48 EtherChannels, depending on the ASA model.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface import AsaInterface
from scc_firewall_manager_sdk.models.ether_channel_interface_patch_input import EtherChannelInterfacePatchInput
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    interface_uid = 'interface_uid_example' # str | Unique identifier (UUID) of the ASA EtherChannel interface.
    ether_channel_interface_patch_input = scc_firewall_manager_sdk.EtherChannelInterfacePatchInput() # EtherChannelInterfacePatchInput | 

    try:
        # Modify ASA EtherChannel interface
        api_response = api_instance.modify_asa_ether_channel_interface(device_uid, interface_uid, ether_channel_interface_patch_input)
        print("The response of ASAInterfacesApi->modify_asa_ether_channel_interface:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->modify_asa_ether_channel_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **interface_uid** | **str**| Unique identifier (UUID) of the ASA EtherChannel interface. | 
 **ether_channel_interface_patch_input** | [**EtherChannelInterfacePatchInput**](EtherChannelInterfacePatchInput.md)|  | 

### Return type

[**AsaInterface**](AsaInterface.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | AsaInterface object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_asa_physical_interface**
> AsaInterface modify_asa_physical_interface(device_uid, interface_uid, physical_interface_patch_input)

Modify ASA Physical interface

Modify the ASA physical interface. A physical interface is a hardware port in an ASA device that can be enabled and configured to connect the firewall to a network.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface import AsaInterface
from scc_firewall_manager_sdk.models.physical_interface_patch_input import PhysicalInterfacePatchInput
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    interface_uid = 'interface_uid_example' # str | Unique identifier (UUID) of the ASA physical interface.
    physical_interface_patch_input = scc_firewall_manager_sdk.PhysicalInterfacePatchInput() # PhysicalInterfacePatchInput | 

    try:
        # Modify ASA Physical interface
        api_response = api_instance.modify_asa_physical_interface(device_uid, interface_uid, physical_interface_patch_input)
        print("The response of ASAInterfacesApi->modify_asa_physical_interface:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->modify_asa_physical_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **interface_uid** | **str**| Unique identifier (UUID) of the ASA physical interface. | 
 **physical_interface_patch_input** | [**PhysicalInterfacePatchInput**](PhysicalInterfacePatchInput.md)|  | 

### Return type

[**AsaInterface**](AsaInterface.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | AsaInterface object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_asa_sub_interface**
> AsaInterface modify_asa_sub_interface(device_uid, interface_uid, sub_interface_patch_input)

Modify ASA subinterface

Modify the ASA subinterface. A subinterface is a logical interface created in a physical interface that allows the ASA to support multiple VLANs by assigning each subinterface to a different VLAN, enabling segmented network traffic over a single physical port.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface import AsaInterface
from scc_firewall_manager_sdk.models.sub_interface_patch_input import SubInterfacePatchInput
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    interface_uid = 'interface_uid_example' # str | Unique identifier (UUID) of the ASA subinterface.
    sub_interface_patch_input = scc_firewall_manager_sdk.SubInterfacePatchInput() # SubInterfacePatchInput | 

    try:
        # Modify ASA subinterface
        api_response = api_instance.modify_asa_sub_interface(device_uid, interface_uid, sub_interface_patch_input)
        print("The response of ASAInterfacesApi->modify_asa_sub_interface:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->modify_asa_sub_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **interface_uid** | **str**| Unique identifier (UUID) of the ASA subinterface. | 
 **sub_interface_patch_input** | [**SubInterfacePatchInput**](SubInterfacePatchInput.md)|  | 

### Return type

[**AsaInterface**](AsaInterface.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | AsaInterface object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **modify_asa_vlan_interface**
> AsaInterface modify_asa_vlan_interface(device_uid, interface_uid, vlan_interface_patch_input)

Modify ASA VLAN interface

Modify the ASA VLAN interface. A VLAN interface is a logical interface that represents a virtual LAN in a physical interface, allowing multiple VLANs to be configured and segmented in a single physical interface for traffic separation and security.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import scc_firewall_manager_sdk
from scc_firewall_manager_sdk.models.asa_interface import AsaInterface
from scc_firewall_manager_sdk.models.vlan_interface_patch_input import VlanInterfacePatchInput
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
    api_instance = scc_firewall_manager_sdk.ASAInterfacesApi(api_client)
    device_uid = 'device_uid_example' # str | Unique identifier (UUID) of the device.
    interface_uid = 'interface_uid_example' # str | Unique identifier (UUID) of the ASA VLAN interface.
    vlan_interface_patch_input = scc_firewall_manager_sdk.VlanInterfacePatchInput() # VlanInterfacePatchInput | 

    try:
        # Modify ASA VLAN interface
        api_response = api_instance.modify_asa_vlan_interface(device_uid, interface_uid, vlan_interface_patch_input)
        print("The response of ASAInterfacesApi->modify_asa_vlan_interface:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling ASAInterfacesApi->modify_asa_vlan_interface: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device_uid** | **str**| Unique identifier (UUID) of the device. | 
 **interface_uid** | **str**| Unique identifier (UUID) of the ASA VLAN interface. | 
 **vlan_interface_patch_input** | [**VlanInterfacePatchInput**](VlanInterfacePatchInput.md)|  | 

### Return type

[**AsaInterface**](AsaInterface.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | AsaInterface object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

