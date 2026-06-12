# TaskDistinctAttributeValues


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**statuses** | **List[str]** | The unique status values for the tasks in the Security Cloud Control tenant | [optional] 
**types** | **List[str]** | The unique task type values for the tasks in the Security Cloud Control tenant | [optional] 
**usernames** | **List[str]** | The unique username values for who triggered tasks in the Security Cloud Control tenant | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.task_distinct_attribute_values import TaskDistinctAttributeValues

# TODO update the JSON string below
json = "{}"
# create an instance of TaskDistinctAttributeValues from a JSON string
task_distinct_attribute_values_instance = TaskDistinctAttributeValues.from_json(json)
# print the JSON string representation of the object
print(TaskDistinctAttributeValues.to_json())

# convert the object into a dict
task_distinct_attribute_values_dict = task_distinct_attribute_values_instance.to_dict()
# create an instance of TaskDistinctAttributeValues from a dict
task_distinct_attribute_values_form_dict = task_distinct_attribute_values.from_dict(task_distinct_attribute_values_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


