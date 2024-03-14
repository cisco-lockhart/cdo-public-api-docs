# UpdateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of the object to update | [optional] 
**description** | **str** | The description of the object | [optional] 
**value** | [**SharedObjectValue**](SharedObjectValue.md) |  | [optional] 
**tags** | **Dict[str, List[str]]** | The tags for the object | [optional] 
**labels** | **List[str]** | The labels for the object | [optional] 

## Example

```python
from openapi_client.models.update_request import UpdateRequest

# TODO update the JSON string below
json = "{}"
# create an instance of UpdateRequest from a JSON string
update_request_instance = UpdateRequest.from_json(json)
# print the JSON string representation of the object
print(UpdateRequest.to_json())

# convert the object into a dict
update_request_dict = update_request_instance.to_dict()
# create an instance of UpdateRequest from a dict
update_request_form_dict = update_request.from_dict(update_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


