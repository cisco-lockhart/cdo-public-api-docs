# AccessRuleUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier, represented as a UUID, of the Access Rule. | 
**index** | **int** | Access rule index position in Access Group ordered rule list. | [optional] 
**rule_action** | **str** | The rule&#39;s action: PERMIT or DENY. | [optional] 
**protocol** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**source_port** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**destination_port** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**source_network** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**destination_network** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**source_dynamic_object** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**destination_dynamic_object** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**log_settings** | [**LogSettings**](LogSettings.md) |  | [optional] 
**rule_time_range** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**remark** | **str** | A remark. | [optional] 
**is_active_rule** | **bool** | Is active. True by default | [optional] 

## Example

```python
from cdo_sdk_python.models.access_rule_update_input import AccessRuleUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of AccessRuleUpdateInput from a JSON string
access_rule_update_input_instance = AccessRuleUpdateInput.from_json(json)
# print the JSON string representation of the object
print(AccessRuleUpdateInput.to_json())

# convert the object into a dict
access_rule_update_input_dict = access_rule_update_input_instance.to_dict()
# create an instance of AccessRuleUpdateInput from a dict
access_rule_update_input_form_dict = access_rule_update_input.from_dict(access_rule_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


