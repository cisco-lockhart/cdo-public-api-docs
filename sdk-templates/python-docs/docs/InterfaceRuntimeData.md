# InterfaceRuntimeData


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**duplex_type** | **str** | Actual duplex mode of the physical interface. | [optional] 
**ip_address** | **str** | Actual IP address of the interface. | [optional] 
**link_enabled** | **bool** | Indicates whether interface link is enabled. | [optional] 
**link_state** | [**LinkState**](LinkState.md) |  | [optional] 
**mac_address** | **str** | Actual MAC address of the interface. | [optional] 
**mtu** | **int** | Actual MTU of the interface. | [optional] 
**speed_type** | **str** | Actual speed of the physical interface. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.interface_runtime_data import InterfaceRuntimeData

# TODO update the JSON string below
json = "{}"
# create an instance of InterfaceRuntimeData from a JSON string
interface_runtime_data_instance = InterfaceRuntimeData.from_json(json)
# print the JSON string representation of the object
print(InterfaceRuntimeData.to_json())

# convert the object into a dict
interface_runtime_data_dict = interface_runtime_data_instance.to_dict()
# create an instance of InterfaceRuntimeData from a dict
interface_runtime_data_form_dict = interface_runtime_data.from_dict(interface_runtime_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


