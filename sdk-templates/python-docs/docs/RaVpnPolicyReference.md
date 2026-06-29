# RaVpnPolicyReference

(Remote Access VPN headend FMC-managed FTDs only) The Remote Access VPN Policy applied to the device.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**link** | **str** | The endpoint to access this resource from. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the Remote Access VPN Policy in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.ra_vpn_policy_reference import RaVpnPolicyReference

# TODO update the JSON string below
json = "{}"
# create an instance of RaVpnPolicyReference from a JSON string
ra_vpn_policy_reference_instance = RaVpnPolicyReference.from_json(json)
# print the JSON string representation of the object
print(RaVpnPolicyReference.to_json())

# convert the object into a dict
ra_vpn_policy_reference_dict = ra_vpn_policy_reference_instance.to_dict()
# create an instance of RaVpnPolicyReference from a dict
ra_vpn_policy_reference_form_dict = ra_vpn_policy_reference.from_dict(ra_vpn_policy_reference_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


