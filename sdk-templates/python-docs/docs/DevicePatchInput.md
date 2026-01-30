# DevicePatchInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_metadata** | [**DeviceMetadata**](DeviceMetadata.md) |  | [optional] 
**labels** | [**Labels**](Labels.md) |  | [optional] 
**maintenance_window_duration_minutes** | **int** | Duration of the maintenance window in minutes. | [optional] 
**maintenance_window_start_cron** | **str** | CRON expression defining the scheduled start time of the maintenance window in format: Minute, Hour, Day of Month, Month, Day of Week. | [optional] 
**name** | **str** | A human-readable name for the device. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device_patch_input import DevicePatchInput

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


