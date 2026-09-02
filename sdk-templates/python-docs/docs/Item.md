# Item

Policy analysis requests grouped by data source

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**data_source_uid** | **str** | The unique identifier, represented as a UUID, of the data source whose policies are analyzed. A data source is the device manager that owns the NAT policies on this tenant, so this is the unique identifier of an FMC device manager onboarded to it, as returned by the &lt;a href&#x3D;\&quot;https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-device-managers/\&quot;&gt;Device Managers&lt;/a&gt; endpoint. | 
**policy_uids** | **List[str]** | The unique identifiers, represented as UUIDs, of the NAT policies to analyze. Use the &lt;a href&#x3D;\&quot;https://developer.cisco.com/docs/cisco-security-cloud-control-firewall-manager/get-all-ftd-nat-policy/\&quot;&gt;Get cdFMC NAT Policies&lt;/a&gt; endpoint, or the equivalent NAT policy retrieval APIs for on-prem FMCs, to get these identifiers. | 

## Example

```python
from scc_firewall_manager_sdk.models.item import Item

# TODO update the JSON string below
json = "{}"
# create an instance of Item from a JSON string
item_instance = Item.from_json(json)
# print the JSON string representation of the object
print(Item.to_json())

# convert the object into a dict
item_dict = item_instance.to_dict()
# create an instance of Item from a dict
item_form_dict = item.from_dict(item_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


