# UsageReference


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | The name of an entity that uses this object | [optional] 
**object_type** | **str** | The type of an entity that uses this object | [optional] 
**uid** | **str** | The unique identifier of an entity that uses this object | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.usage_reference import UsageReference

# TODO update the JSON string below
json = "{}"
# create an instance of UsageReference from a JSON string
usage_reference_instance = UsageReference.from_json(json)
# print the JSON string representation of the object
print(UsageReference.to_json())

# convert the object into a dict
usage_reference_dict = usage_reference_instance.to_dict()
# create an instance of UsageReference from a dict
usage_reference_form_dict = usage_reference.from_dict(usage_reference_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


