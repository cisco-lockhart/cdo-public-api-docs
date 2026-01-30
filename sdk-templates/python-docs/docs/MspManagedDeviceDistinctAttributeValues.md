# MspManagedDeviceDistinctAttributeValues


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config_states** | [**List[ConfigState]**](ConfigState.md) | The distinct config states for the devices managed by the MSP Portal. | [optional] 
**conflict_detection_states** | [**List[ConflictDetectionState]**](ConflictDetectionState.md) | The distinct conflict detection states for the devices managed by the MSP Portal. | [optional] 
**connectivity_states** | [**List[ConnectivityState]**](ConnectivityState.md) | The distinct connectivity states for the devices managed by the MSP Portal. | [optional] 
**device_types** | [**List[EntityType]**](EntityType.md) | The distinct device types for the devices managed by the MSP Portal. | [optional] 
**hardware_models** | **List[str]** | The distinct hardware models for the devices managed by the MSP Portal. | [optional] 
**managed_tenant_display_names** | **List[str]** | The display names of the tenants that have devices onboarded and are managed by the MSP Portal. | [optional] 
**managed_tenant_names** | **List[str]** | The names of the tenants that have devices onboarded and are managed by the MSP Portal. | [optional] 
**model_numbers** | **List[str]** | The distinct model numbers for the devices managed by the MSP Portal. | [optional] 
**redundancy_modes** | **List[str]** | The distinct redundancy modes for the devices managed by the MSP Portal. | [optional] 
**software_versions** | **Dict[str, List[str]]** | The software versions of the devices managed by the MSP Portal. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_managed_device_distinct_attribute_values import MspManagedDeviceDistinctAttributeValues

# TODO update the JSON string below
json = "{}"
# create an instance of MspManagedDeviceDistinctAttributeValues from a JSON string
msp_managed_device_distinct_attribute_values_instance = MspManagedDeviceDistinctAttributeValues.from_json(json)
# print the JSON string representation of the object
print(MspManagedDeviceDistinctAttributeValues.to_json())

# convert the object into a dict
msp_managed_device_distinct_attribute_values_dict = msp_managed_device_distinct_attribute_values_instance.to_dict()
# create an instance of MspManagedDeviceDistinctAttributeValues from a dict
msp_managed_device_distinct_attribute_values_form_dict = msp_managed_device_distinct_attribute_values.from_dict(msp_managed_device_distinct_attribute_values_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


