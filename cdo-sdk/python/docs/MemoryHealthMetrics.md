# MemoryHealthMetrics

The memory health metrics for the device.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lina_usage_avg** | **float** | The average memory usage of the Lina process on the device, expressed as a percentage value between 0 and 1. | [optional] 
**snort_usage_avg** | **float** | The average memory usage of the Snort process on the device, expressed as a percentage value between 0 and 1. | [optional] 
**system_usage_avg** | **float** | The average memory usage of all processes on the device, expressed as a percentage value between 0 and 1. | [optional] 

## Example

```python
from cdo_sdk_python.models.memory_health_metrics import MemoryHealthMetrics

# TODO update the JSON string below
json = "{}"
# create an instance of MemoryHealthMetrics from a JSON string
memory_health_metrics_instance = MemoryHealthMetrics.from_json(json)
# print the JSON string representation of the object
print(MemoryHealthMetrics.to_json())

# convert the object into a dict
memory_health_metrics_dict = memory_health_metrics_instance.to_dict()
# create an instance of MemoryHealthMetrics from a dict
memory_health_metrics_form_dict = memory_health_metrics.from_dict(memory_health_metrics_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


