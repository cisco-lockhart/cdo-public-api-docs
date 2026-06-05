# CpuHealthMetrics


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lina_usage_avg** | **float** | Measures the average CPU utilisation by the LINA (Cisco&#39;s ASA software running natively). Expressed as a percentage value between 0 and 100. | [optional] 
**snort_usage_avg** | **float** | Indicates the average CPU usage by the Snort process, responsible for threat detection, including intrusion prevention and advanced malware protection. Expressed as a percentage value between 0 and 100. | [optional] 
**system_usage_avg** | **float** | Represents the total average CPU load utilised by the FTD system, including both firewall and threat defense mechanisms. Expressed as a percentage value between 0 and 100. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.cpu_health_metrics import CpuHealthMetrics

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


