# Target

Set of targets that contain the object. A target can be, for example, a device, service, or a shared policy (Ruleset).

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **str** | The ID of the target with which the object is associated. A target can be, for example, a device, service, or a shared policy (Ruleset). | [optional] 
**display_name** | **str** | The display name of the target | [optional] 
**type** | **str** | The target type | [optional] 

## Example

```python
from cdo-python-sdk.models.target import Target

# TODO update the JSON string below
json = "{}"
# create an instance of Target from a JSON string
target_instance = Target.from_json(json)
# print the JSON string representation of the object
print(Target.to_json())

# convert the object into a dict
target_dict = target_instance.to_dict()
# create an instance of Target from a dict
target_form_dict = target.from_dict(target_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


