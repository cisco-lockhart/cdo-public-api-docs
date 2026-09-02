# NatPolicySyncRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data_source_uids** | **List[str]** | the unique identifiers, represented as UUIDs, of the Data sources to sync. A data source is the device manager that owns the NAT policies on this tenant, so these are the unique identifiers of the FMC device managers onboarded to it, as returned by the &lt;a href&#x3D;\&quot;https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-device-managers/\&quot;&gt;Device Managers&lt;/a&gt; endpoint. If omitted or empty, all eligible NAT PAO FMC data sources are synced. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.nat_policy_sync_request import NatPolicySyncRequest

# TODO update the JSON string below
json = "{}"
# create an instance of NatPolicySyncRequest from a JSON string
nat_policy_sync_request_instance = NatPolicySyncRequest.from_json(json)
# print the JSON string representation of the object
print(NatPolicySyncRequest.to_json())

# convert the object into a dict
nat_policy_sync_request_dict = nat_policy_sync_request_instance.to_dict()
# create an instance of NatPolicySyncRequest from a dict
nat_policy_sync_request_form_dict = nat_policy_sync_request.from_dict(nat_policy_sync_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


