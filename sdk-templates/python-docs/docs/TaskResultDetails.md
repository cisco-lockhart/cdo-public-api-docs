# TaskResultDetails

Type-specific result produced by a task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **str** |  | 
**num_groups** | **int** |  | [optional] 
**num_users** | **int** |  | [optional] 
**download_urls** | **List[str]** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.task_result_details import TaskResultDetails

# TODO update the JSON string below
json = "{}"
# create an instance of TaskResultDetails from a JSON string
task_result_details_instance = TaskResultDetails.from_json(json)
# print the JSON string representation of the object
print(TaskResultDetails.to_json())

# convert the object into a dict
task_result_details_dict = task_result_details_instance.to_dict()
# create an instance of TaskResultDetails from a dict
task_result_details_form_dict = task_result_details.from_dict(task_result_details_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


