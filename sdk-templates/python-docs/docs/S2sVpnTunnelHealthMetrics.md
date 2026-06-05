# S2sVpnTunnelHealthMetrics


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tunnel_id** | **str** | The unique identifier of the S2S VPN tunnel. | [optional] 
**tunnel_name** | **str** | The name assigned to the S2S VPN tunnel. | [optional] 
**tunnel_state** | **str** | The state of the S2S VPN tunnel. The possible values are TUNNEL_UP, TUNNEL_DOWN, and UNKNOWN. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.s2s_vpn_tunnel_health_metrics import S2sVpnTunnelHealthMetrics

# TODO update the JSON string below
json = "{}"
# create an instance of S2sVpnTunnelHealthMetrics from a JSON string
s2s_vpn_tunnel_health_metrics_instance = S2sVpnTunnelHealthMetrics.from_json(json)
# print the JSON string representation of the object
print(S2sVpnTunnelHealthMetrics.to_json())

# convert the object into a dict
s2s_vpn_tunnel_health_metrics_dict = s2s_vpn_tunnel_health_metrics_instance.to_dict()
# create an instance of S2sVpnTunnelHealthMetrics from a dict
s2s_vpn_tunnel_health_metrics_form_dict = s2s_vpn_tunnel_health_metrics.from_dict(s2s_vpn_tunnel_health_metrics_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


