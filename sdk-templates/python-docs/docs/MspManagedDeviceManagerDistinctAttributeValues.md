# MspManagedDeviceManagerDistinctAttributeValues


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**connectivity_states** | [**List[ConnectivityState]**](ConnectivityState.md) | The distinct connectivity states for the device managers managed by the MSP Portal. | [optional] 
**device_manager_types** | [**List[EntityType]**](EntityType.md) | The distinct device manager types for the device managers managed by the MSP Portal. | [optional] 
**managed_tenant_display_names** | **List[str]** | The display names of the tenants that have device managers onboarded and are managed by the MSP Portal. | [optional] 
**managed_tenant_names** | **List[str]** | The names of the tenants that have device managers onboarded and are managed by the MSP Portal. | [optional] 
**software_versions** | **Dict[str, List[str]]** | The software versions of the device managers managed by the MSP Portal. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_managed_device_manager_distinct_attribute_values import MspManagedDeviceManagerDistinctAttributeValues

# TODO update the JSON string below
json = "{}"
# create an instance of MspManagedDeviceManagerDistinctAttributeValues from a JSON string
msp_managed_device_manager_distinct_attribute_values_instance = MspManagedDeviceManagerDistinctAttributeValues.from_json(json)
# print the JSON string representation of the object
print(MspManagedDeviceManagerDistinctAttributeValues.to_json())

# convert the object into a dict
msp_managed_device_manager_distinct_attribute_values_dict = msp_managed_device_manager_distinct_attribute_values_instance.to_dict()
# create an instance of MspManagedDeviceManagerDistinctAttributeValues from a dict
msp_managed_device_manager_distinct_attribute_values_form_dict = msp_managed_device_manager_distinct_attribute_values.from_dict(msp_managed_device_manager_distinct_attribute_values_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


