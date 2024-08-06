# CpuHealthMetrics

The CPU health metrics for the device. This value will be available only if the health policy on the device has CPU monitoring enabled.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lina_usage_avg** | **float** | The average CPU usage of the Lina process on the device, expressed as a percentage value between 0 and 1. | [optional] 
**snort_usage_avg** | **float** | The average CPU usage of the Snort process on the device, expressed as a percentage value between 0 and 1. | [optional] 
**system_usage_avg** | **float** | The average CPU usage of all processes on the device, expressed as a percentage value between 0 and 1. | [optional] 

## Example

```python
from cdo_sdk_python.models.cpu_health_metrics import CpuHealthMetrics

# TODO update the JSON string below
json = "{}"
# create an instance of CpuHealthMetrics from a JSON string
cpu_health_metrics_instance = CpuHealthMetrics.from_json(json)
# print the JSON string representation of the object
print(CpuHealthMetrics.to_json())

# convert the object into a dict
cpu_health_metrics_dict = cpu_health_metrics_instance.to_dict()
# create an instance of CpuHealthMetrics from a dict
cpu_health_metrics_form_dict = cpu_health_metrics.from_dict(cpu_health_metrics_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


