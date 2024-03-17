# DevicePatchInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A human-readable name for the device. | [optional] 
**labels** | [**Labels**](Labels.md) |  | [optional] 

## Example

```python
from cdo-python-sdk.models.device_patch_input import DevicePatchInput

# TODO update the JSON string below
json = "{}"
# create an instance of DevicePatchInput from a JSON string
device_patch_input_instance = DevicePatchInput.from_json(json)
# print the JSON string representation of the object
print(DevicePatchInput.to_json())

# convert the object into a dict
device_patch_input_dict = device_patch_input_instance.to_dict()
# create an instance of DevicePatchInput from a dict
device_patch_input_form_dict = device_patch_input.from_dict(device_patch_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


