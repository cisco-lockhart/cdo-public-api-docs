# OS

The operating system of the client device.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**type** | **str** | The type of operating system running on the client device. | [optional] 
**version** | **str** | The version of the operating system running on the client device. | [optional] 

## Example

```python
from cdo_sdk_python.models.os import OS

# TODO update the JSON string below
json = "{}"
# create an instance of OS from a JSON string
os_instance = OS.from_json(json)
# print the JSON string representation of the object
print(OS.to_json())

# convert the object into a dict
os_dict = os_instance.to_dict()
# create an instance of OS from a dict
os_form_dict = os.from_dict(os_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


