# RaVpnSessionHealthMetrics


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**active_ravpn_sessions_avg** | **float** | The average number of active RA VPN sessions. | [optional] 
**inactive_ravpn_sessions_avg** | **float** | The average number of inactive or down RA VPN sessions. | [optional] 
**peak_concur_ravpn_sessions** | **float** | The peak concurrent RA VPN sessions active since the last reset. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.ra_vpn_session_health_metrics import RaVpnSessionHealthMetrics

# TODO update the JSON string below
json = "{}"
# create an instance of RaVpnSessionHealthMetrics from a JSON string
ra_vpn_session_health_metrics_instance = RaVpnSessionHealthMetrics.from_json(json)
# print the JSON string representation of the object
print(RaVpnSessionHealthMetrics.to_json())

# convert the object into a dict
ra_vpn_session_health_metrics_dict = ra_vpn_session_health_metrics_instance.to_dict()
# create an instance of RaVpnSessionHealthMetrics from a dict
ra_vpn_session_health_metrics_form_dict = ra_vpn_session_health_metrics.from_dict(ra_vpn_session_health_metrics_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


