# RemediationActionsDto

The remediation plan for a NAT policy analysis.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**insight_uid** | **str** | The unique identifier of the NAT policy analysis. | [optional] 
**redundant_rules_actions** | [**List[ObservationRemediationActionDto]**](ObservationRemediationActionDto.md) | The remediation actions selected for redundant NAT rules. | [optional] 
**shadowed_rules_actions** | [**List[ObservationRemediationActionDto]**](ObservationRemediationActionDto.md) | The remediation actions selected for shadowed NAT rules. | [optional] 
**total_deleted_count** | **int** | The total number of unique NAT rules selected for deletion. | [optional] [readonly] 
**total_disabled_count** | **int** | The total number of unique NAT rules selected to be disabled. | [optional] [readonly] 

## Example

```python
from scc_firewall_manager_sdk.models.remediation_actions_dto import RemediationActionsDto

# TODO update the JSON string below
json = "{}"
# create an instance of RemediationActionsDto from a JSON string
remediation_actions_dto_instance = RemediationActionsDto.from_json(json)
# print the JSON string representation of the object
print(RemediationActionsDto.to_json())

# convert the object into a dict
remediation_actions_dto_dict = remediation_actions_dto_instance.to_dict()
# create an instance of RemediationActionsDto from a dict
remediation_actions_dto_form_dict = remediation_actions_dto.from_dict(remediation_actions_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


