# FmcHealthMetrics


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_uid** | **str** | The unique identifier, represented as a UUID, of the device in SCC. | [optional] 
**device_name** | **str** | The name of the device in SCC. | [optional] 
**start_time** | **datetime** | The start of the time period (inclusive) for which the health metrics were retrieved for this device. | [optional] 
**end_time** | **datetime** | The end of the time period (inclusive) for which the health metrics were retrieved for this device. | [optional] 
**cpu_health_metrics** | [**CpuHealthMetrics**](CpuHealthMetrics.md) |  | [optional] 
**memory_health_metrics** | [**MemoryHealthMetrics**](MemoryHealthMetrics.md) |  | [optional] 
**disk_health_metrics** | [**DiskHealthMetrics**](DiskHealthMetrics.md) |  | [optional] 
**chassis_stats_health_metrics** | [**ChassisStatsHealthMetrics**](ChassisStatsHealthMetrics.md) |  | [optional] 
**interface_health_metrics** | [**List[InterfaceHealthMetrics]**](InterfaceHealthMetrics.md) | The interface health metrics for the device. | [optional] 
**vpn_health_metrics** | [**VpnHealthMetrics**](VpnHealthMetrics.md) |  | [optional] 
**ha_health_metrics** | [**HaHealthMetrics**](HaHealthMetrics.md) |  | [optional] 

## Example

```python
from cdo_sdk_python.models.fmc_health_metrics import FmcHealthMetrics

# TODO update the JSON string below
json = "{}"
# create an instance of FmcHealthMetrics from a JSON string
fmc_health_metrics_instance = FmcHealthMetrics.from_json(json)
# print the JSON string representation of the object
print(FmcHealthMetrics.to_json())

# convert the object into a dict
fmc_health_metrics_dict = fmc_health_metrics_instance.to_dict()
# create an instance of FmcHealthMetrics from a dict
fmc_health_metrics_form_dict = fmc_health_metrics.from_dict(fmc_health_metrics_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


