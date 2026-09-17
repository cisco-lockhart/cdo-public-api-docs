# TaskDtoResult


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **str** |  | 
**completed_acknowledged** | **int** |  | [optional] 
**completed_awaiting_acknowledgement** | **int** |  | [optional] 
**requested** | **int** |  | [optional] 
**error_msg** | **str** |  | [optional] 
**num_groups** | **int** |  | [optional] 
**num_users** | **int** |  | [optional] 
**device_address** | **str** |  | [optional] 
**log_path** | **str** |  | [optional] 
**progress** | **int** |  | [optional] 
**stage** | **str** |  | [optional] 
**version** | **str** |  | [optional] 
**devices_with_failing_validation_checks** | **List[str]** |  | [optional] 
**devices_with_passing_validation_checks** | **List[str]** |  | [optional] 
**check_identifier** | **str** |  | [optional] 
**download_urls** | **List[str]** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.task_dto_result import TaskDtoResult

# TODO update the JSON string below
json = "{}"
# create an instance of TaskDtoResult from a JSON string
task_dto_result_instance = TaskDtoResult.from_json(json)
# print the JSON string representation of the object
print(TaskDtoResult.to_json())

# convert the object into a dict
task_dto_result_dict = task_dto_result_instance.to_dict()
# create an instance of TaskDtoResult from a dict
task_dto_result_form_dict = task_dto_result.from_dict(task_dto_result_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


