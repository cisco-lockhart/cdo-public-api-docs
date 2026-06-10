# DeviceDistinctAttributeValues


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config_states** | [**List[ConfigState]**](ConfigState.md) | The distinct config states for the devices in the Security Cloud Control tenant. | [optional] 
**conflict_detection_states** | [**List[ConflictDetectionState]**](ConflictDetectionState.md) | The distinct conflict detection states for the devices in the Security Cloud Control tenant. | [optional] 
**connectivity_states** | [**List[ConnectivityState]**](ConnectivityState.md) | The distinct connectivity states for the devices in the Security Cloud Control tenant. | [optional] 
**device_roles** | [**List[DeviceRole]**](DeviceRole.md) | The distinct device roles for the devices in the Security Cloud Control tenant. | [optional] 
**device_types** | [**List[EntityType]**](EntityType.md) | The distinct device types for the devices in the Security Cloud Control tenant. | [optional] 
**hardware_models** | **List[str]** | The distinct hardware models for the devices in the Security Cloud Control tenant. | [optional] 
**model_numbers** | **List[str]** | The distinct model numbers for the devices in the Security Cloud Control tenant. | [optional] 
**redundancy_modes** | **List[str]** | The distinct redundancy modes for the devices in the Security Cloud Control tenant. | [optional] 
**software_versions** | **Dict[str, List[str]]** | The software versions of the devices in the Security Cloud Control tenant, grouped by device type. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device_distinct_attribute_values import DeviceDistinctAttributeValues

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceDistinctAttributeValues from a JSON string
device_distinct_attribute_values_instance = DeviceDistinctAttributeValues.from_json(json)
# print the JSON string representation of the object
print(DeviceDistinctAttributeValues.to_json())

# convert the object into a dict
device_distinct_attribute_values_dict = device_distinct_attribute_values_instance.to_dict()
# create an instance of DeviceDistinctAttributeValues from a dict
device_distinct_attribute_values_form_dict = device_distinct_attribute_values.from_dict(device_distinct_attribute_values_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


