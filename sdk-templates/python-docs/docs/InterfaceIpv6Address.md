# InterfaceIpv6Address


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auto_config** | **bool** | Indicates whether address is automatically configured. | [optional] 
**dad_attempts** | **int** | Indicates the number of times a device performs Duplicate Address Detection (DAD) to verify the uniqueness of an address before assigning the IPv6 address. Valid range is from 0 to 600. | [optional] 
**enabled** | **bool** | Indicates whether IPv6 is enabled in interface. | [optional] 
**ip_addresses** | [**List[InterfaceIpAddress]**](InterfaceIpAddress.md) | List of IPv6 addresses assigned to interface. | [optional] 
**link_local_address** | [**InterfaceIpAddress**](InterfaceIpAddress.md) |  | [optional] 
**suppress_ra** | **bool** | Indicates whether router advertisements should be suppressed. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.interface_ipv6_address import InterfaceIpv6Address

# TODO update the JSON string below
json = "{}"
# create an instance of InterfaceIpv6Address from a JSON string
interface_ipv6_address_instance = InterfaceIpv6Address.from_json(json)
# print the JSON string representation of the object
print(InterfaceIpv6Address.to_json())

# convert the object into a dict
interface_ipv6_address_dict = interface_ipv6_address_instance.to_dict()
# create an instance of InterfaceIpv6Address from a dict
interface_ipv6_address_form_dict = interface_ipv6_address.from_dict(interface_ipv6_address_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


