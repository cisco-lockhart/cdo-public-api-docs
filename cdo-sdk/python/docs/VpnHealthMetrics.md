# VpnHealthMetrics

The vpn health metrics for the device.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**active_ravpn_tunnels_avg** | **float** | Measures the average number of active RAVPN tunnels. | [optional] 
**cumulative_ravpn_tunnels_avg** | **float** | Measures the average number of cumulative RAVPN tunnels since the last system reset or statistics clearing. | [optional] 
**inactive_ravpn_tunnels_avg** | **float** | Measures the average number of inactive of down RAVPN tunnels. | [optional] 
**peak_concur_ravpn_tunnels_avg** | **float** | Measures the average number of the highest concurrent RAVPN tunnels active since the last system reset or statistics clearing. | [optional] 

## Example

```python
from cdo_sdk_python.models.vpn_health_metrics import VpnHealthMetrics

# TODO update the JSON string below
json = "{}"
# create an instance of VpnHealthMetrics from a JSON string
vpn_health_metrics_instance = VpnHealthMetrics.from_json(json)
# print the JSON string representation of the object
print(VpnHealthMetrics.to_json())

# convert the object into a dict
vpn_health_metrics_dict = vpn_health_metrics_instance.to_dict()
# create an instance of VpnHealthMetrics from a dict
vpn_health_metrics_form_dict = vpn_health_metrics.from_dict(vpn_health_metrics_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


