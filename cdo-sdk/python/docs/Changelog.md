# Changelog


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**uid** | **str** | The unique identifier of the Change Log. | 
**status** | **str** | The status of the Change Log. | [optional] 
**entity_uid** | **str** | The uid of the device/manager/service the Change Log refers to. | [optional] 
**events** | [**List[Event]**](Event.md) | The events recorded in this Change Log. | [optional] 

## Example

```python
from cdo_python_sdk.models.changelog import Changelog

# TODO update the JSON string below
json = "{}"
# create an instance of Changelog from a JSON string
changelog_instance = Changelog.from_json(json)
# print the JSON string representation of the object
print(Changelog.to_json())

# convert the object into a dict
changelog_dict = changelog_instance.to_dict()
# create an instance of Changelog from a dict
changelog_form_dict = changelog.from_dict(changelog_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


