# PageTaskDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**count** | **int** | The total number of results available. | [optional] 
**items** | [**List[TaskDto]**](TaskDto.md) | The list of items retrieved. | [optional] 
**limit** | **int** | The number of results retrieved. | [optional] 
**offset** | **int** | The offset of the results retrieved. The Security Cloud Control API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.page_task_dto import PageTaskDto

# TODO update the JSON string below
json = "{}"
# create an instance of PageTaskDto from a JSON string
page_task_dto_instance = PageTaskDto.from_json(json)
# print the JSON string representation of the object
print(PageTaskDto.to_json())

# convert the object into a dict
page_task_dto_dict = page_task_dto_instance.to_dict()
# create an instance of PageTaskDto from a dict
page_task_dto_form_dict = page_task_dto.from_dict(page_task_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


