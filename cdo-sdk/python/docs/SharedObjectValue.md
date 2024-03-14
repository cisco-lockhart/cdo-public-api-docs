# SharedObjectValue

The value of the object

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**object_type** | **str** | The type of object | 
**default_content** | [**ObjectContent**](ObjectContent.md) |  | 
**overrides** | [**List[Override]**](Override.md) | The list of target overrides for the object. Each override the default content for its target. | [optional] 

## Example

```python
from openapi_client.models.shared_object_value import SharedObjectValue

# TODO update the JSON string below
json = "{}"
# create an instance of SharedObjectValue from a JSON string
shared_object_value_instance = SharedObjectValue.from_json(json)
# print the JSON string representation of the object
print(SharedObjectValue.to_json())

# convert the object into a dict
shared_object_value_dict = shared_object_value_instance.to_dict()
# create an instance of SharedObjectValue from a dict
shared_object_value_form_dict = shared_object_value.from_dict(shared_object_value_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


