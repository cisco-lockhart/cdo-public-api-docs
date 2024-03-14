# Override

The list of target overrides for the object. Each override the default content for its target.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**target_id** | **str** | The ID of the target. A target can be, for example, a device, service, or a shared policy (Ruleset). | [optional] 
**content** | [**ObjectContent**](ObjectContent.md) |  | 

## Example

```python
from openapi_client.models.override import Override

# TODO update the JSON string below
json = "{}"
# create an instance of Override from a JSON string
override_instance = Override.from_json(json)
# print the JSON string representation of the object
print(Override.to_json())

# convert the object into a dict
override_dict = override_instance.to_dict()
# create an instance of Override from a dict
override_form_dict = override.from_dict(override_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


