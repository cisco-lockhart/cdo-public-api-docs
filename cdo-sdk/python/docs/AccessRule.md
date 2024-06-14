# AccessRule


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier of Access Rule in CDO. | 
**access_group_uid** | **str** | The unique identifier of the Access Group associated with the Access Rule. | 
**shared_access_group_uid** | **str** | Optional unique identifier for the shared Access Group associated with a shared Access Rule. | [optional] 
**entity_uid** | **str** | The unique identifier of the device/manager associated with the Access Rule. Points to shared Access Group in case of shared Rule | 
**index** | **int** | Access rule index position in Access Group ordered rule list. | 
**rule_type** | **str** | The L3 level rule type. L3, L7 or CONTENT_FILTERING. Defaults to L3. | [optional] 
**rule_action** | **str** | The rule&#39;s action: PERMIT or DENY. | 
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
**issue** | **str** | Issues. SHADOWED or null. | [optional] 
**is_active_rule** | **bool** | Is active. True by default | [optional] 
**created_date** | **datetime** | The time (in UTC) at which Access Rule was created, represented using the RFC-3339 standard. | [optional] 
**updated_date** | **datetime** | The time (in UTC) at which Access Rule was updated, represented using the RFC-3339 standard. | [optional] 

## Example

```python
from cdo_sdk_python.models.access_rule import AccessRule

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


