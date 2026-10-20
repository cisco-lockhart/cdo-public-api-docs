# ObservationRemediationActionDto

The remediation actions selected for a rule conflict observation.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**action** | **str** | The remediation action to apply to all shadowed or redundant rules in this observation if &#x60;ruleActions&#x60; is omitted, null, or empty. | [optional] 
**observation_id** | **str** | The unique identifier of the rule conflict observation. | 
**rule_actions** | [**List[RuleRemediationActionDto]**](RuleRemediationActionDto.md) | The remediation actions selected for individual NAT rules. If this list is not empty, these actions are used instead of the observation action. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.observation_remediation_action_dto import ObservationRemediationActionDto

# TODO update the JSON string below
json = "{}"
# create an instance of ObservationRemediationActionDto from a JSON string
observation_remediation_action_dto_instance = ObservationRemediationActionDto.from_json(json)
# print the JSON string representation of the object
print(ObservationRemediationActionDto.to_json())

# convert the object into a dict
observation_remediation_action_dto_dict = observation_remediation_action_dto_instance.to_dict()
# create an instance of ObservationRemediationActionDto from a dict
observation_remediation_action_dto_form_dict = observation_remediation_action_dto.from_dict(observation_remediation_action_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


