# DiskHealthMetrics

The disk health metrics for the device.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**total_disk_usage_avg** | **float** | The average utilisation of disk space, important for monitoring log storage, reporting, and archiving activities which are critical for audit and diagnostics. Expressed as a percentage value between 0 and 100. | [optional] 

## Example

```python
from cdo_sdk_python.models.disk_health_metrics import DiskHealthMetrics

# TODO update the JSON string below
json = "{}"
# create an instance of DiskHealthMetrics from a JSON string
disk_health_metrics_instance = DiskHealthMetrics.from_json(json)
# print the JSON string representation of the object
print(DiskHealthMetrics.to_json())

# convert the object into a dict
disk_health_metrics_dict = disk_health_metrics_instance.to_dict()
# create an instance of DiskHealthMetrics from a dict
disk_health_metrics_form_dict = disk_health_metrics.from_dict(disk_health_metrics_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


