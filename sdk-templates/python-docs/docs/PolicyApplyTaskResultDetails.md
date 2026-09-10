# PolicyApplyTaskResultDetails


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**devices_with_failing_validation_checks** | **List[str]** |  | [optional] 
**devices_with_passing_validation_checks** | **List[str]** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.policy_apply_task_result_details import PolicyApplyTaskResultDetails

# TODO update the JSON string below
json = "{}"
# create an instance of PolicyApplyTaskResultDetails from a JSON string
policy_apply_task_result_details_instance = PolicyApplyTaskResultDetails.from_json(json)
# print the JSON string representation of the object
print(PolicyApplyTaskResultDetails.to_json())

# convert the object into a dict
policy_apply_task_result_details_dict = policy_apply_task_result_details_instance.to_dict()
# create an instance of PolicyApplyTaskResultDetails from a dict
policy_apply_task_result_details_form_dict = policy_apply_task_result_details.from_dict(policy_apply_task_result_details_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


