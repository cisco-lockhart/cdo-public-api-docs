# DeviceManagerPatchInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_metadata** | [**DeviceMetadata**](DeviceMetadata.md) |  | [optional] 
**labels** | [**Labels**](Labels.md) |  | [optional] 
**name** | **str** | A human-readable name for the device manager. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device_manager_patch_input import DeviceManagerPatchInput

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceManagerPatchInput from a JSON string
device_manager_patch_input_instance = DeviceManagerPatchInput.from_json(json)
# print the JSON string representation of the object
print(DeviceManagerPatchInput.to_json())

# convert the object into a dict
device_manager_patch_input_dict = device_manager_patch_input_instance.to_dict()
# create an instance of DeviceManagerPatchInput from a dict
device_manager_patch_input_form_dict = device_manager_patch_input.from_dict(device_manager_patch_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


