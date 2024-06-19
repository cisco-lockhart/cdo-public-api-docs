# AccessGroupUpdateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A human-readable name for the Access Group. | [optional] 
**resources** | **List[Dict[str, object]]** | The set of of interface and direction pairs or global resource. | [optional] 

## Example

```python
from cdo_sdk_python.models.access_group_update_input import AccessGroupUpdateInput

# TODO update the JSON string below
json = "{}"
# create an instance of AccessGroupUpdateInput from a JSON string
access_group_update_input_instance = AccessGroupUpdateInput.from_json(json)
# print the JSON string representation of the object
print(AccessGroupUpdateInput.to_json())

# convert the object into a dict
access_group_update_input_dict = access_group_update_input_instance.to_dict()
# create an instance of AccessGroupUpdateInput from a dict
access_group_update_input_form_dict = access_group_update_input.from_dict(access_group_update_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


