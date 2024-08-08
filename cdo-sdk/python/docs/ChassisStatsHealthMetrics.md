# ChassisStatsHealthMetrics

The chassis health metrics for the device.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**fan1_rpm_avg** | **float** | The average speed of fan 1, if present, in RPM. | [optional] 
**fan2_rpm_avg** | **float** | The average speed of fan 2, if present, in RPM. | [optional] 
**fan3_rpm_avg** | **float** | The average speed of fan 3, if present, in RPM. | [optional] 
**fan4_rpm_avg** | **float** | The average speed of fan 4, if present, in RPM. | [optional] 

## Example

```python
from cdo_sdk_python.models.chassis_stats_health_metrics import ChassisStatsHealthMetrics

# TODO update the JSON string below
json = "{}"
# create an instance of ChassisStatsHealthMetrics from a JSON string
chassis_stats_health_metrics_instance = ChassisStatsHealthMetrics.from_json(json)
# print the JSON string representation of the object
print(ChassisStatsHealthMetrics.to_json())

# convert the object into a dict
chassis_stats_health_metrics_dict = chassis_stats_health_metrics_instance.to_dict()
# create an instance of ChassisStatsHealthMetrics from a dict
chassis_stats_health_metrics_form_dict = chassis_stats_health_metrics.from_dict(chassis_stats_health_metrics_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


