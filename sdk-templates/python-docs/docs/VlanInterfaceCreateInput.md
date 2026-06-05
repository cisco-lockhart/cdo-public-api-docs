# VlanInterfaceCreateInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** | Interface description. | [optional] 
**enabled** | **bool** | Indicates whether the interface is enabled. | [optional] 
**ipv4** | [**InterfaceIpv4Address**](InterfaceIpv4Address.md) |  | [optional] 
**ipv6** | [**InterfaceIpv6Address**](InterfaceIpv6Address.md) |  | [optional] 
**mac_address** | **str** | Interface MAC address is a unique hardware identifier assigned to each physical interface and used for Layer 2 communication in a network segment. | [optional] 
**management_only** | **bool** | Indicates if the interface is used exclusively for management traffic and not regular user or data traffic. | [optional] 
**member_interfaces** | **List[str]** | Set of unique identifiers (UUIDs) of the member interfaces that are a part of the VLAN group. | [optional] 
**monitor_interface** | **bool** | Indicates if the interface is actively monitored as part of an HA setup. The check determines if the interface is functioning correctly and whether a failover to a standby device should be triggered, if required. | [optional] 
**mtu** | **int** | Interface MTU, in bytes, of a packet or frame that can be sent over a network interface without needing to be fragmented. MTU defines the maximum payload size the interface can handle in a single transmission. . Valid range is from 64 to 9198. | [optional] 
**name** | **str** | Logical name of the interface, which must be unique in the scope of the device. | [optional] 
**security_level** | **int** | Indicates the interface trust level, ranging from 0 (lowest) to 100 (highest). | [optional] 
**standby_mac_address** | **str** | The standby interface Media Access Control (MAC) address, for use in a high-availability (HA) pair. | [optional] 
**vlan_id** | **int** | Unique identifier of the VLAN associated with the VLAN interface. Valid range is from 1 to 4070. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.vlan_interface_create_input import VlanInterfaceCreateInput

# TODO update the JSON string below
json = "{}"
# create an instance of VlanInterfaceCreateInput from a JSON string
vlan_interface_create_input_instance = VlanInterfaceCreateInput.from_json(json)
# print the JSON string representation of the object
print(VlanInterfaceCreateInput.to_json())

# convert the object into a dict
vlan_interface_create_input_dict = vlan_interface_create_input_instance.to_dict()
# create an instance of VlanInterfaceCreateInput from a dict
vlan_interface_create_input_form_dict = vlan_interface_create_input.from_dict(vlan_interface_create_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


