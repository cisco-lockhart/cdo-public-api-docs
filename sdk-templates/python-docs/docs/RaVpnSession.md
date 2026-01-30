# RaVpnSession


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**assigned_ip_v4** | **str** | The IPv4 address assigned to the RA VPN session. | [optional] 
**assigned_ip_v6** | **str** | The IPv6 address assigned to the RA VPN session. | [optional] 
**audit_session_id** | **str** | The audit session ID associated with this RA VPN session. | [optional] 
**bytes_rx** | **int** | The number of bytes received during the RA VPN session. | [optional] 
**bytes_tx** | **int** | The number of bytes transmitted during the RA VPN session. | [optional] 
**device_uid** | **str** | The unique identifier, represented as a UUID, of the device associated with the RA VPN session. | 
**last_active_time** | **datetime** | The time (in UTC) at which the user was last active in the RA VPN session, represented using the RFC-3339 standard. | [optional] 
**location** | [**Location**](Location.md) |  | [optional] 
**login_time** | **datetime** | The time (in UTC) at which the user logged in to the RA VPN session, represented using the RFC-3339 standard. | [optional] 
**os** | [**OS**](OS.md) |  | [optional] 
**public_ip** | **str** | The public IP address of the client that has established this RA VPN session. | [optional] 
**status** | **str** | The status of the RA VPN session. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the VPN session. | 
**username** | **str** | The name of the user associated with the RA VPN session. | 

## Example

```python
from scc_firewall_manager_sdk.models.ra_vpn_session import RaVpnSession

# TODO update the JSON string below
json = "{}"
# create an instance of RaVpnSession from a JSON string
ra_vpn_session_instance = RaVpnSession.from_json(json)
# print the JSON string representation of the object
print(RaVpnSession.to_json())

# convert the object into a dict
ra_vpn_session_dict = ra_vpn_session_instance.to_dict()
# create an instance of RaVpnSession from a dict
ra_vpn_session_form_dict = ra_vpn_session.from_dict(ra_vpn_session_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


