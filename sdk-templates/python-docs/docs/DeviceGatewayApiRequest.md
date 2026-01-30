# DeviceGatewayApiRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_uid** | **str** | The unique identifier, represented as a UUID, of the target device for the command. Currently, only FMC devices are supported. | 
**request** | [**FmcRequest**](FmcRequest.md) |  | 

## Example

```python
from scc_firewall_manager_sdk.models.device_gateway_api_request import DeviceGatewayApiRequest

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceGatewayApiRequest from a JSON string
device_gateway_api_request_instance = DeviceGatewayApiRequest.from_json(json)
# print the JSON string representation of the object
print(DeviceGatewayApiRequest.to_json())

# convert the object into a dict
device_gateway_api_request_dict = device_gateway_api_request_instance.to_dict()
# create an instance of DeviceGatewayApiRequest from a dict
device_gateway_api_request_form_dict = device_gateway_api_request.from_dict(device_gateway_api_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


