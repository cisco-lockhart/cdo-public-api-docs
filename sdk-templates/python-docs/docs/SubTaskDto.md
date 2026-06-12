# SubTaskDto

The list of items retrieved.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** | The description of the subtask in Security Cloud Control | [optional] 
**status** | **str** | The status of the task in Security Cloud Control | [optional] 
**target_name** | **str** | The name of the target the subtask is running against. | [optional] 
**target_type** | **str** | The type of the target the subtask is running against. | [optional] 
**target_uid** | **str** | The unique identifier, represented as a UUID, of the target the subtask is running against. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the subtask in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.sub_task_dto import SubTaskDto

# TODO update the JSON string below
json = "{}"
# create an instance of SubTaskDto from a JSON string
sub_task_dto_instance = SubTaskDto.from_json(json)
# print the JSON string representation of the object
print(SubTaskDto.to_json())

# convert the object into a dict
sub_task_dto_dict = sub_task_dto_instance.to_dict()
# create an instance of SubTaskDto from a dict
sub_task_dto_form_dict = sub_task_dto.from_dict(sub_task_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


