# AccessRule


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**access_group_uid** | **str** | The unique identifier, represented as a UUID, of the Access Group associated with the Access Rule. | 
**created_date** | **datetime** | The time (in UTC) at which Access Rule was created, represented using the RFC-3339 standard. | [optional] 
**destination_dynamic_object** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**destination_network** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**destination_port** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**entity_uid** | **str** | The unique identifier, represented as a UUID, of the device/manager associated with the Access Rule. Points to shared Access Group in case of shared Rule | 
**index** | **int** | Access rule index position in Access Group ordered rule list. | 
**is_active_rule** | **bool** | Indicates whether this rule is active. | [optional] 
**issues** | [**List[AccessRuleIssue]**](AccessRuleIssue.md) | Indicates if rule has SHADOWED, SHARED or DUPLICATE (remark) issue. | [optional] 
**log_settings** | [**LogSettings**](LogSettings.md) |  | [optional] 
**protocol** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**remark** | **str** | A remark. | [optional] 
**rule_action** | **str** | Representation of the rule action. | [optional] 
**rule_configuration_line_number** | **int** | Representation of the rule&#39;s line number in configuration. | [optional] 
**rule_configuration_text** | **str** | Representation of the rule in configuration. | [optional] 
**rule_parsing_error** | **str** | The error, if any, that was encountered when parsing the rule. | [optional] 
**rule_time_range** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**rule_type** | **str** | The L3 level rule type. L3, L7 or CONTENT_FILTERING. Defaults to L3. | [optional] 
**shared_access_group_uid** | **str** | Optional unique identifier for the shared Access Group associated with a shared Access Rule. | [optional] 
**source_dynamic_object** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**source_network** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**source_port** | [**AccessRuleDetailsContent**](AccessRuleDetailsContent.md) |  | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of Access Rule in Security Cloud Control. | 
**updated_date** | **datetime** | The time (in UTC) at which Access Rule was updated, represented using the RFC-3339 standard. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.access_rule import AccessRule

# TODO update the JSON string below
json = "{}"
# create an instance of AccessRule from a JSON string
access_rule_instance = AccessRule.from_json(json)
# print the JSON string representation of the object
print(AccessRule.to_json())

# convert the object into a dict
access_rule_dict = access_rule_instance.to_dict()
# create an instance of AccessRule from a dict
access_rule_form_dict = access_rule.from_dict(access_rule_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


