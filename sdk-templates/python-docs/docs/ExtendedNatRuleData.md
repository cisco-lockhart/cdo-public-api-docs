# ExtendedNatRuleData


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
from scc_firewall_manager_sdk.models.extended_nat_rule_data import ExtendedNatRuleData

# TODO update the JSON string below
json = "{}"
# create an instance of ExtendedNatRuleData from a JSON string
extended_nat_rule_data_instance = ExtendedNatRuleData.from_json(json)
# print the JSON string representation of the object
print(ExtendedNatRuleData.to_json())

# convert the object into a dict
extended_nat_rule_data_dict = extended_nat_rule_data_instance.to_dict()
# create an instance of ExtendedNatRuleData from a dict
extended_nat_rule_data_form_dict = extended_nat_rule_data.from_dict(extended_nat_rule_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


