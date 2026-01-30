# FmcTemplateVariable


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | Specify the name of the object to override. | 
**value** | **str** | Specify the value of the template variable. The value specified should be valid for the template variable, as defined in the template configuration. | 

## Example

```python
from scc_firewall_manager_sdk.models.fmc_template_variable import FmcTemplateVariable

# TODO update the JSON string below
json = "{}"
# create an instance of FmcTemplateVariable from a JSON string
fmc_template_variable_instance = FmcTemplateVariable.from_json(json)
# print the JSON string representation of the object
print(FmcTemplateVariable.to_json())

# convert the object into a dict
fmc_template_variable_dict = fmc_template_variable_instance.to_dict()
# create an instance of FmcTemplateVariable from a dict
fmc_template_variable_form_dict = fmc_template_variable.from_dict(fmc_template_variable_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


