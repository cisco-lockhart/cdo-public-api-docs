# LogSettings

The set of of interface and direction pairs or global resource.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**level** | **str** | The level | 
**interval** | **int** | The interval | [optional] 

## Example

```python
from cdo_sdk_python.models.log_settings import LogSettings

# TODO update the JSON string below
json = "{}"
# create an instance of LogSettings from a JSON string
log_settings_instance = LogSettings.from_json(json)
# print the JSON string representation of the object
print(LogSettings.to_json())

# convert the object into a dict
log_settings_dict = log_settings_instance.to_dict()
# create an instance of LogSettings from a dict
log_settings_form_dict = log_settings.from_dict(log_settings_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


