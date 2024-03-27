# ClientDevice

The client device that triggered this MFA event.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier of the device. | 
**location** | [**Location**](Location.md) |  | [optional] 
**ip_address** | **str** | The IP address of the client device that has triggered this MFA event. | [optional] 
**password_set** | **bool** | Indicates whether a password is set on the client device. | [optional] 
**encrypted** | **bool** | Indicates whether encryption is enabled on the client device. | [optional] 
**firewalled** | **bool** | Indicates whether a firewall is enabled on the client device. | [optional] 
**os** | [**OS**](OS.md) |  | [optional] 
**browser** | [**Browser**](Browser.md) |  | [optional] 

## Example

```python
from cdo_sdk_python.models.client_device import ClientDevice

# TODO update the JSON string below
json = "{}"
# create an instance of ClientDevice from a JSON string
client_device_instance = ClientDevice.from_json(json)
# print the JSON string representation of the object
print(ClientDevice.to_json())

# convert the object into a dict
client_device_dict = client_device_instance.to_dict()
# create an instance of ClientDevice from a dict
client_device_form_dict = client_device.from_dict(client_device_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


