# InterfaceHealthMetrics

The interface health metrics for the device.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**status** | **str** |  | [optional] 
**buffer_underruns_avg** | **float** |  | [optional] 
**buffer_overruns_avg** | **float** |  | [optional] 
**interface** | **str** |  | [optional] 
**interface_name** | **str** |  | [optional] 

## Example

```python
from cdo_sdk_python.models.interface_health_metrics import InterfaceHealthMetrics

# TODO update the JSON string below
json = "{}"
# create an instance of InterfaceHealthMetrics from a JSON string
interface_health_metrics_instance = InterfaceHealthMetrics.from_json(json)
# print the JSON string representation of the object
print(InterfaceHealthMetrics.to_json())

# convert the object into a dict
interface_health_metrics_dict = interface_health_metrics_instance.to_dict()
# create an instance of InterfaceHealthMetrics from a dict
interface_health_metrics_form_dict = interface_health_metrics.from_dict(interface_health_metrics_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

