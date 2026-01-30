# VpnHealthMetrics

The vpn health metrics for the device.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**active_ravpn_tunnels_avg** | **float** | The average number of active RA VPN tunnels. | [optional] 
**active_s2svpn_tunnels_avg** | **float** | The average number of active S2S VPN tunnels. | [optional] 
**inactive_ravpn_tunnels_avg** | **float** | The average number of inactive or down RA VPN tunnels. | [optional] 
**inactive_s2svpn_tunnels_avg** | **float** | The average number of inactive or down S2S VPN tunnels. | [optional] 
**peak_concur_ravpn_tunnels** | **float** | The peak concurrent RA VPN tunnels active since the last reset. | [optional] 
**peak_concur_s2svpn_tunnels** | **float** | The peak concurrent S2S VPN tunnels active since the last reset. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.vpn_health_metrics import VpnHealthMetrics

# TODO update the JSON string below
json = "{}"
# create an instance of VpnHealthMetrics from a JSON string
vpn_health_metrics_instance = VpnHealthMetrics.from_json(json)
# print the JSON string representation of the object
print(VpnHealthMetrics.to_json())

# convert the object into a dict
vpn_health_metrics_dict = vpn_health_metrics_instance.to_dict()
# create an instance of VpnHealthMetrics from a dict
vpn_health_metrics_from_dict = VpnHealthMetrics.from_dict(vpn_health_metrics_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


