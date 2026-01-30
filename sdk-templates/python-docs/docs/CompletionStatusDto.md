# CompletionStatusDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**message** | **str** | A message, indicating current progress, returned from the device. | [optional] 
**name** | **str** | The name of the device. Note: this will be the name of the node on the cdFMC if the device is part of an HA pair or cluster. | [optional] 
**node_type** | **str** | (HA Pairs and clusters only) The node type of the device. | [optional] 
**percentage_complete** | **float** | The percentage completion of the current task being performed | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.completion_status_dto import CompletionStatusDto

# TODO update the JSON string below
json = "{}"
# create an instance of CompletionStatusDto from a JSON string
completion_status_dto_instance = CompletionStatusDto.from_json(json)
# print the JSON string representation of the object
print(CompletionStatusDto.to_json())

# convert the object into a dict
completion_status_dto_dict = completion_status_dto_instance.to_dict()
# create an instance of CompletionStatusDto from a dict
completion_status_dto_form_dict = completion_status_dto.from_dict(completion_status_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


