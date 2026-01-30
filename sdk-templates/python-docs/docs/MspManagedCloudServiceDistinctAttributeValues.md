# MspManagedCloudServiceDistinctAttributeValues


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cloud_service_types** | [**List[EntityType]**](EntityType.md) | The distinct device types for the cloud services managed by the MSP Portal. | [optional] 
**config_states** | [**List[ConfigState]**](ConfigState.md) | The distinct config states for the cloud services managed by the MSP Portal. | [optional] 
**conflict_detection_states** | [**List[ConflictDetectionState]**](ConflictDetectionState.md) | The distinct conflict detection states for the cloud services managed by the MSP Portal. | [optional] 
**connectivity_states** | [**List[ConnectivityState]**](ConnectivityState.md) | The distinct connectivity states for the cloud services managed by the MSP Portal. | [optional] 
**managed_tenant_display_names** | **List[str]** | The display names of the tenants that have cloud services onboarded and are managed by the MSP Portal. | [optional] 
**managed_tenant_names** | **List[str]** | The names of the tenants that have cloud services onboarded and are managed by the MSP Portal. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_managed_cloud_service_distinct_attribute_values import MspManagedCloudServiceDistinctAttributeValues

# TODO update the JSON string below
json = "{}"
# create an instance of MspManagedCloudServiceDistinctAttributeValues from a JSON string
msp_managed_cloud_service_distinct_attribute_values_instance = MspManagedCloudServiceDistinctAttributeValues.from_json(json)
# print the JSON string representation of the object
print(MspManagedCloudServiceDistinctAttributeValues.to_json())

# convert the object into a dict
msp_managed_cloud_service_distinct_attribute_values_dict = msp_managed_cloud_service_distinct_attribute_values_instance.to_dict()
# create an instance of MspManagedCloudServiceDistinctAttributeValues from a dict
msp_managed_cloud_service_distinct_attribute_values_form_dict = msp_managed_cloud_service_distinct_attribute_values.from_dict(msp_managed_cloud_service_distinct_attribute_values_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


