# MemoryHealthMetrics


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lina_usage_avg** | **float** | Average memory utilisation by the LINA process. Expressed as a percentage value between 0 and 100. | [optional] 
**snort_usage_avg** | **float** | Average memory usage by the Snort engine. Expressed as a percentage value between 0 and 100. | [optional] 
**system_usage_avg** | **float** | Provides an overview of the overall memory usage by the FTD system, useful for monitoring system health and detecting potential resource bottlenecks. Expressed as a percentage value between 0 and 100. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.memory_health_metrics import MemoryHealthMetrics

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


