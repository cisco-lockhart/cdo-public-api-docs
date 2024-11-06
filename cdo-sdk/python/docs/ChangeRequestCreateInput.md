# ChangeRequestCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of the change request in Security Cloud Control. | 
**description** | **str** | The description of the change request in Security Cloud Control. | [optional] 

## Example

```python
from cdo_sdk_python.models.change_request_create_input import ChangeRequestCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of ChangeRequestCreateInput from a JSON string
change_request_create_input_instance = ChangeRequestCreateInput.from_json(json)
# print the JSON string representation of the object
print(ChangeRequestCreateInput.to_json())

# convert the object into a dict
change_request_create_input_dict = change_request_create_input_instance.to_dict()
# create an instance of ChangeRequestCreateInput from a dict
change_request_create_input_form_dict = change_request_create_input.from_dict(change_request_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


