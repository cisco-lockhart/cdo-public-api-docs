# ConflictingRules

The Data of the conflicting rules.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**redundant_rules** | [**List[ExtendedNatRuleData]**](ExtendedNatRuleData.md) |  | [optional] 
**shadowed_rules** | [**List[ExtendedNatRuleData]**](ExtendedNatRuleData.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.conflicting_rules import ConflictingRules

# TODO update the JSON string below
json = "{}"
# create an instance of ConflictingRules from a JSON string
conflicting_rules_instance = ConflictingRules.from_json(json)
# print the JSON string representation of the object
print(ConflictingRules.to_json())

# convert the object into a dict
conflicting_rules_dict = conflicting_rules_instance.to_dict()
# create an instance of ConflictingRules from a dict
conflicting_rules_form_dict = conflicting_rules.from_dict(conflicting_rules_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


