# OriginalRule

The Data of the Original rule.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**destination_interface** | [**ObjectData**](ObjectData.md) |  | [optional] 
**direction** | **str** |  | [optional] 
**enabled** | **bool** |  | [optional] 
**name** | **str** |  | [optional] 
**original_destination** | [**ObjectData**](ObjectData.md) |  | [optional] 
**original_destination_port** | [**ObjectData**](ObjectData.md) |  | [optional] 
**original_source** | [**ObjectData**](ObjectData.md) |  | [optional] 
**original_source_port** | [**ObjectData**](ObjectData.md) |  | [optional] 
**rule_index** | **int** |  | [optional] 
**rule_name** | **str** |  | [optional] 
**rule_section** | **str** |  | [optional] 
**rule_type** | **str** |  | [optional] 
**source_interface** | [**ObjectData**](ObjectData.md) |  | [optional] 
**translated_destination** | [**ObjectData**](ObjectData.md) |  | [optional] 
**translated_destination_port** | [**ObjectData**](ObjectData.md) |  | [optional] 
**translated_source** | [**ObjectData**](ObjectData.md) |  | [optional] 
**translated_source_port** | [**ObjectData**](ObjectData.md) |  | [optional] 
**type** | **str** | Type of the rule. | [optional] 
**uid** | **str** | UID of the rule. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.original_rule import OriginalRule

# TODO update the JSON string below
json = "{}"
# create an instance of OriginalRule from a JSON string
original_rule_instance = OriginalRule.from_json(json)
# print the JSON string representation of the object
print(OriginalRule.to_json())

# convert the object into a dict
original_rule_dict = original_rule_instance.to_dict()
# create an instance of OriginalRule from a dict
original_rule_form_dict = original_rule.from_dict(original_rule_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


