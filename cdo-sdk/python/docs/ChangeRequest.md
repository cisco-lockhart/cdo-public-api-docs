# ChangeRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier of the Change Request. | 
**name** | **str** | The name of the Change Request. | 
**description** | **str** | The description of the Change Request. | [optional] 

## Example

```python
from cdo-python-sdk.models.change_request import ChangeRequest

# TODO update the JSON string below
json = "{}"
# create an instance of ChangeRequest from a JSON string
change_request_instance = ChangeRequest.from_json(json)
# print the JSON string representation of the object
print(ChangeRequest.to_json())

# convert the object into a dict
change_request_dict = change_request_instance.to_dict()
# create an instance of ChangeRequest from a dict
change_request_form_dict = change_request.from_dict(change_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


