# DeviceMaintenanceWindow


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**current_end** | **str** | End time of the current maintenance window in ISO 8601 format. Note: This field is only populated when isMaintenanceWindowOpen is true. | [optional] 
**current_start** | **str** | Start time of the current maintenance window in ISO 8601 format. Note: This field is only populated when isMaintenanceWindowOpen is true. | [optional] 
**duration_minutes** | **int** | Duration of the maintenance window in minutes. | [optional] 
**is_open** | **bool** | Indicates whether the maintenance window is currently open. | [optional] 
**next_end** | **str** | End time of the next scheduled maintenance window in ISO 8601 format. | [optional] 
**next_start** | **str** | Start time of the next scheduled maintenance window in ISO 8601 format. | [optional] 
**start_cron** | **str** | CRON expression defining the scheduled start time of the maintenance window in format: Minute, Hour, Day of Month, Month, Day of Week. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device_maintenance_window import DeviceMaintenanceWindow

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceMaintenanceWindow from a JSON string
device_maintenance_window_instance = DeviceMaintenanceWindow.from_json(json)
# print the JSON string representation of the object
print(DeviceMaintenanceWindow.to_json())

# convert the object into a dict
device_maintenance_window_dict = device_maintenance_window_instance.to_dict()
# create an instance of DeviceMaintenanceWindow from a dict
device_maintenance_window_form_dict = device_maintenance_window.from_dict(device_maintenance_window_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


