# LabelDefinitionReference

Source-owned label assignments for display

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** | The label UUID supplied by the owning source | 
**name** | **str** | The label display name, when known | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.label_definition_reference import LabelDefinitionReference

# TODO update the JSON string below
json = "{}"
# create an instance of LabelDefinitionReference from a JSON string
label_definition_reference_instance = LabelDefinitionReference.from_json(json)
# print the JSON string representation of the object
print(LabelDefinitionReference.to_json())

# convert the object into a dict
label_definition_reference_dict = label_definition_reference_instance.to_dict()
# create an instance of LabelDefinitionReference from a dict
label_definition_reference_form_dict = label_definition_reference.from_dict(label_definition_reference_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


