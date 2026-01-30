# MspManagedTemplateDistinctAttributeValues


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**config_states** | [**List[ConfigState]**](ConfigState.md) | The distinct config states for the templates managed by the MSP Portal. | [optional] 
**managed_tenant_display_names** | **List[str]** | The display names of the tenants that have templates managed by the MSP Portal. | [optional] 
**managed_tenant_names** | **List[str]** | The names of the tenants that have templates managed by the MSP Portal. | [optional] 
**software_versions** | **Dict[str, List[str]]** | The software versions of the templates managed by the MSP Portal. | [optional] 
**template_types** | [**List[EntityType]**](EntityType.md) | The distinct templates types for the templates managed by the MSP Portal. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_managed_template_distinct_attribute_values import MspManagedTemplateDistinctAttributeValues

# TODO update the JSON string below
json = "{}"
# create an instance of MspManagedTemplateDistinctAttributeValues from a JSON string
msp_managed_template_distinct_attribute_values_instance = MspManagedTemplateDistinctAttributeValues.from_json(json)
# print the JSON string representation of the object
print(MspManagedTemplateDistinctAttributeValues.to_json())

# convert the object into a dict
msp_managed_template_distinct_attribute_values_dict = msp_managed_template_distinct_attribute_values_instance.to_dict()
# create an instance of MspManagedTemplateDistinctAttributeValues from a dict
msp_managed_template_distinct_attribute_values_form_dict = msp_managed_template_distinct_attribute_values.from_dict(msp_managed_template_distinct_attribute_values_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


