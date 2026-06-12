# TaskDto

The list of items retrieved.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** | The description of the task in Security Cloud Control | [optional] 
**last_updated_time** | **datetime** | Time (UTC; represented using the RFC-3339 standard) at which the task was last updated. | [optional] 
**num_sub_tasks** | **int** | The number of subtasks associated with this task. | [optional] 
**status** | **str** | The status of the task in Security Cloud Control | [optional] 
**status_details** | **str** | Human-readable details on the status of the task | [optional] 
**submission_time** | **datetime** | Time (UTC; represented using the RFC-3339 standard) at which the task was triggered. | [optional] 
**type** | **str** | The type of the task in Security Cloud Control | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the task in Security Cloud Control. | [optional] 
**username** | **str** | The last user to interact with (or trigger) this task. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.task_dto import TaskDto

# TODO update the JSON string below
json = "{}"
# create an instance of TaskDto from a JSON string
task_dto_instance = TaskDto.from_json(json)
# print the JSON string representation of the object
print(TaskDto.to_json())

# convert the object into a dict
task_dto_dict = task_dto_instance.to_dict()
# create an instance of TaskDto from a dict
task_dto_form_dict = task_dto.from_dict(task_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


