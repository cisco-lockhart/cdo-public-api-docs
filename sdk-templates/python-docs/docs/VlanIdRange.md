# VlanIdRange


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**end** | **int** |  | [optional] 
**start** | **int** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.vlan_id_range import VlanIdRange

# TODO update the JSON string below
json = "{}"
# create an instance of VlanIdRange from a JSON string
vlan_id_range_instance = VlanIdRange.from_json(json)
# print the JSON string representation of the object
print(VlanIdRange.to_json())

# convert the object into a dict
vlan_id_range_dict = vlan_id_range_instance.to_dict()
# create an instance of VlanIdRange from a dict
vlan_id_range_form_dict = vlan_id_range.from_dict(vlan_id_range_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


