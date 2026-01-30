# InterfaceIpAddress


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ip_address** | **str** | Unique network address assigned to a physical or logical port, which enables the port to communicate on a specific network segment and enforce security policies for traffic traversing that interface. | [optional] 
**netmask** | **str** | A 32-bit number that defines the network portion of an IP address, differentiating it from the host portion. This number determines the size of the local network segment directly connected to the interface, allowing it to identify which traffic belongs to its local network and which needs to be routed. | [optional] 
**setroute** | **bool** | Indicates whether a default route using the gateway information provided by the DHCP server is created automatically. The DHCP-assigned gateway is used as the next-hop for routing traffic. | [optional] 
**standby_ip_address** | **str** | Secondary IP address configured in an interface, specifically for use in an HA pair. This address stays inactive in the standby unit until a failover occurs, at which point, the standby assumes ownership of both the primary and secondary IP addresses to ensure network connectivity for devices. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.interface_ip_address import InterfaceIpAddress

# TODO update the JSON string below
json = "{}"
# create an instance of InterfaceIpAddress from a JSON string
interface_ip_address_instance = InterfaceIpAddress.from_json(json)
# print the JSON string representation of the object
print(InterfaceIpAddress.to_json())

# convert the object into a dict
interface_ip_address_dict = interface_ip_address_instance.to_dict()
# create an instance of InterfaceIpAddress from a dict
interface_ip_address_form_dict = interface_ip_address.from_dict(interface_ip_address_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


