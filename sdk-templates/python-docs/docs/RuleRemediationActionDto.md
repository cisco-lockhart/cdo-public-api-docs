# RuleRemediationActionDto

The remediation action selected for a NAT rule.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**action** | **str** | The remediation action to apply to the NAT rule. | 
**rule_id** | **str** | The unique identifier of the NAT rule. | 

## Example

```python
from scc_firewall_manager_sdk.models.rule_remediation_action_dto import RuleRemediationActionDto

# TODO update the JSON string below
json = "{}"
# create an instance of RuleRemediationActionDto from a JSON string
rule_remediation_action_dto_instance = RuleRemediationActionDto.from_json(json)
# print the JSON string representation of the object
print(RuleRemediationActionDto.to_json())

# convert the object into a dict
rule_remediation_action_dto_dict = rule_remediation_action_dto_instance.to_dict()
# create an instance of RuleRemediationActionDto from a dict
rule_remediation_action_dto_form_dict = rule_remediation_action_dto.from_dict(rule_remediation_action_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


