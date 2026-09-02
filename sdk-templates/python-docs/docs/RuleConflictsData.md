# RuleConflictsData

Rule Conflicts Data holding original rule and conflicting rules.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**conflicting_rules** | [**ConflictingRules**](ConflictingRules.md) |  | [optional] 
**is_redundant** | **bool** | Are the conflicting rules redundant rules. | [optional] 
**is_shadowed** | **bool** | Are the conflicting rules shadowed rules. | [optional] 
**observation_id** | **str** | Rule conflict observation identifier. | [optional] 
**original_rule** | [**OriginalRule**](OriginalRule.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.rule_conflicts_data import RuleConflictsData

# TODO update the JSON string below
json = "{}"
# create an instance of RuleConflictsData from a JSON string
rule_conflicts_data_instance = RuleConflictsData.from_json(json)
# print the JSON string representation of the object
print(RuleConflictsData.to_json())

# convert the object into a dict
rule_conflicts_data_dict = rule_conflicts_data_instance.to_dict()
# create an instance of RuleConflictsData from a dict
rule_conflicts_data_form_dict = rule_conflicts_data.from_dict(rule_conflicts_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


