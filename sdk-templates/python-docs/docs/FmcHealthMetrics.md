# FmcHealthMetrics


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**chassis_stats_health_metrics** | [**ChassisStatsHealthMetrics**](ChassisStatsHealthMetrics.md) |  | [optional] 
**cpu_health_metrics** | [**CpuHealthMetrics**](CpuHealthMetrics.md) |  | [optional] 
**device_name** | **str** | The name of the device in Security Cloud Control. | [optional] 
**device_uid** | **str** | The unique identifier, represented as a UUID, of the device in Security Cloud Control. | [optional] 
**disk_health_metrics** | [**DiskHealthMetrics**](DiskHealthMetrics.md) |  | [optional] 
**end_time** | **datetime** | The end of the time period (inclusive) for which the health metrics were retrieved for this device. | [optional] 
**ha_health_metrics** | [**HaHealthMetrics**](HaHealthMetrics.md) |  | [optional] 
**interface_health_metrics** | [**List[InterfaceHealthMetrics]**](InterfaceHealthMetrics.md) | The interface health metrics for the device. | [optional] 
**memory_health_metrics** | [**MemoryHealthMetrics**](MemoryHealthMetrics.md) |  | [optional] 
**ra_vpn_session_health_metrics** | [**RaVpnSessionHealthMetrics**](RaVpnSessionHealthMetrics.md) |  | [optional] 
**s2s_vpn_tunnel_health_metrics** | [**List[S2sVpnTunnelHealthMetrics]**](S2sVpnTunnelHealthMetrics.md) | The S2S VPN tunnel health metrics for the device. The maximum number of tunnels retrieved is 1000. | [optional] 
**start_time** | **datetime** | The start of the time period (inclusive) for which the health metrics were retrieved for this device. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.fmc_health_metrics import FmcHealthMetrics

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


