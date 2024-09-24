# HaHealthMetrics

The HA health metrics for the device. This value will be available only if the device is part of an HA pair.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**node_type** | **str** | Indicates whether the device is a primary or secondary node in an HA pair. | [optional] 
**node_status** | **str** | The status of the HA node. | [optional] 

## Example

```python
from cdo_sdk_python.models.ha_health_metrics import HaHealthMetrics

# TODO update the JSON string below
json = "{}"
# create an instance of HaHealthMetrics from a JSON string
ha_health_metrics_instance = HaHealthMetrics.from_json(json)
# print the JSON string representation of the object
print(HaHealthMetrics.to_json())

# convert the object into a dict
ha_health_metrics_dict = ha_health_metrics_instance.to_dict()
# create an instance of HaHealthMetrics from a dict
ha_health_metrics_form_dict = ha_health_metrics.from_dict(ha_health_metrics_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


