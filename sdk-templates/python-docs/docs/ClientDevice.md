# ClientDevice


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**browser** | [**Browser**](Browser.md) |  | [optional] 
**encrypted** | **bool** | Indicates whether encryption is enabled on the client device. | [optional] 
**firewalled** | **bool** | Indicates whether a firewall is enabled on the client device. | [optional] 
**ip_address** | **str** | The IP address of the client device that has triggered this MFA event. | [optional] 
**location** | [**Location**](Location.md) |  | [optional] 
**os** | [**OS**](OS.md) |  | [optional] 
**password_set** | **bool** | Indicates whether a password is set on the client device. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the device. | 

## Example

```python
from scc_firewall_manager_sdk.models.client_device import ClientDevice

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


