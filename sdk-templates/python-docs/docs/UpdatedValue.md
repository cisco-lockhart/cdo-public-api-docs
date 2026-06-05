# UpdatedValue

The list of values updated on this entity as part of the change.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**field_name** | **str** |  | [optional] 
**new_value** | **str** |  | [optional] 
**old_value** | **str** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.updated_value import UpdatedValue

# TODO update the JSON string below
json = "{}"
# create an instance of UpdatedValue from a JSON string
updated_value_instance = UpdatedValue.from_json(json)
# print the JSON string representation of the object
print(UpdatedValue.to_json())

# convert the object into a dict
updated_value_dict = updated_value_instance.to_dict()
# create an instance of UpdatedValue from a dict
updated_value_form_dict = updated_value.from_dict(updated_value_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


