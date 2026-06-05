# SwitchPortConfig


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**access_mode_vlan** | **str** | Unique identifier (UUID) of the access VLAN with which the interface is associated. | [optional] 
**protected_enabled** | **bool** | Indicates whether switchport interface is protected or not. Protection prevents Layer 2 communication (unicast, multicast, or broadcast) between other protected ports on the same VLAN. | [optional] 
**trunk_mode_allowed_vlans** | **List[str]** | Unique identifiers (UUIDs) of the VLAN interfaces that are permitted to send and receive traffic over trunk port. | [optional] 
**trunk_mode_native_vlan** | **str** | Unique identifier (UUID) of the trunk VLAN with which the interface is associated. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.switch_port_config import SwitchPortConfig

# TODO update the JSON string below
json = "{}"
# create an instance of SwitchPortConfig from a JSON string
switch_port_config_instance = SwitchPortConfig.from_json(json)
# print the JSON string representation of the object
print(SwitchPortConfig.to_json())

# convert the object into a dict
switch_port_config_dict = switch_port_config_instance.to_dict()
# create an instance of SwitchPortConfig from a dict
switch_port_config_form_dict = switch_port_config.from_dict(switch_port_config_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


