# AccessGroupCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A human-readable name for the Access Group. | [optional] 
**entity_uid** | **str** | The unique identifier of the device/manager associated with the Access Group. | 
**resources** | **List[Dict[str, object]]** | The set of of interface and direction pairs or global resource. | [optional] 

## Example

```python
from cdo_sdk_python.models.access_group_create_input import AccessGroupCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of AccessGroupCreateInput from a JSON string
access_group_create_input_instance = AccessGroupCreateInput.from_json(json)
# print the JSON string representation of the object
print(AccessGroupCreateInput.to_json())

# convert the object into a dict
access_group_create_input_dict = access_group_create_input_instance.to_dict()
# create an instance of AccessGroupCreateInput from a dict
access_group_create_input_form_dict = access_group_create_input.from_dict(access_group_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


