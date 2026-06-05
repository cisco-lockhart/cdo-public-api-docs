# AsaInterface


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**description** | **str** | Interface description. | [optional] 
**device_uid** | **str** | Unique identifier (UUID) of the device associated with the interface. | [optional] 
**duplex_type** | [**DuplexType**](DuplexType.md) |  | [optional] 
**enabled** | **bool** | Indicates whether the interface is enabled. | [optional] 
**ether_channel_group_uid** | **str** | Unique identifier (UUID) of the EtherChannel group to which this interface belongs. | [optional] 
**ether_channel_id** | **int** | Identifier of the EtherChannel interface in the scope of the device. | [optional] 
**hardware_name** | **str** | Interface hardware name. It is usually structured from the type, speed, slot and port number. | [optional] 
**interface_type** | [**InterfaceType**](InterfaceType.md) |  | [optional] 
**ipv4** | [**InterfaceIpv4Address**](InterfaceIpv4Address.md) |  | [optional] 
**ipv6** | [**InterfaceIpv6Address**](InterfaceIpv6Address.md) |  | [optional] 
**lacp_mode** | **str** | The Link Aggregation Control Protocol (LACP) mode of the EtherChannel interface. | [optional] 
**mac_address** | **str** | Interface MAC address is a unique hardware identifier assigned to each physical interface and used for Layer 2 communication in a network segment. | [optional] 
**management_interface** | **bool** | Indicates whether the interface is used for device management. | [optional] 
**management_only** | **bool** | Indicates whether the interface is exclusively used for management traffic and not regular user or data traffic. | [optional] 
**member_interfaces** | **List[str]** | Set of unique identifiers (UUIDs) of interfaces that are part of an EtherChannel group. | [optional] 
**mode** | [**InterfaceMode**](InterfaceMode.md) |  | [optional] 
**monitor_interface** | **bool** | Indicates if the interface is actively monitored as part of an HA setup. The check determines if the interface is functioning correctly and whether a failover to a standby device should be triggered, if required. | [optional] 
**mtu** | **int** | Interface MTU, in bytes, of a packet or frame that can be sent over a network interface without needing to be fragmented. It defines the maximum payload size the interface can handle in a single transmission. Valid range is from 64 to 9198. | [optional] 
**name** | **str** | Logical name of the interface, unique in the scope of the device. | [optional] 
**parent_interface_uid** | **str** | Unique identifier (UUID) of the parent interface. | [optional] 
**runtime_data** | [**InterfaceRuntimeData**](InterfaceRuntimeData.md) |  | [optional] 
**security_level** | **int** | Indicates interface trust level, ranging from 0 (lowest) to 100 (highest). | [optional] 
**speed_type** | [**SpeedType**](SpeedType.md) |  | [optional] 
**standby_mac_address** | **str** | Standby interface MAC address for use in HA pairs. | [optional] 
**sub_interface_id** | **int** | Unique identifier of the subinterface, which is used to differentiate between the subinterfaces of the same parent interface. Valid range is from 1 to 4294967295. | [optional] 
**switch_port_config** | [**SwitchPortConfig**](SwitchPortConfig.md) |  | [optional] 
**uid** | **str** | Unique identifier (UUID) of the interface in Security Cloud Control. | [optional] 
**vlan_id** | **int** | Unique identifier of the VLAN associated with VLAN interface. Valid range is from 1 to 4070. | [optional] 
**vlan_primary_id** | **int** | Primary unique identifier of the VLAN associated with subinterface. Valid range is from 1 to 4094. | [optional] 
**vlan_secondary_ids** | [**List[VlanIdRange]**](VlanIdRange.md) | Secondary identifiers of the VLAN associated with subinterface. Each range is defined by a start and end value, which allows multiple VLANs to be associated with a subinterface. All ranges are inclusive. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.asa_interface import AsaInterface

# TODO update the JSON string below
json = "{}"
# create an instance of AsaInterface from a JSON string
asa_interface_instance = AsaInterface.from_json(json)
# print the JSON string representation of the object
print(AsaInterface.to_json())

# convert the object into a dict
asa_interface_dict = asa_interface_instance.to_dict()
# create an instance of AsaInterface from a dict
asa_interface_form_dict = asa_interface.from_dict(asa_interface_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


