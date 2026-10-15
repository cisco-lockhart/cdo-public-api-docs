# ObservationRemediationActionDto

Action selection for one rule-conflict observation.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**action** | **str** | Remediation action type to be applied for observation level remediation. | [optional] 
**observation_id** | **str** | Rule conflict observation identifier. | 
**rule_actions** | [**List[RuleRemediationActionDto]**](RuleRemediationActionDto.md) | Optional rule-level remediation actions. When non-empty, these actions are authoritative. | [optional] 

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


