# InterfaceIpv4Address


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ip_address** | [**InterfaceIpAddress**](InterfaceIpAddress.md) |  | [optional] 
**ip_type** | [**IpType**](IpType.md) |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.interface_ipv4_address import InterfaceIpv4Address

# TODO update the JSON string below
json = "{}"
# create an instance of InterfaceIpv4Address from a JSON string
interface_ipv4_address_instance = InterfaceIpv4Address.from_json(json)
# print the JSON string representation of the object
print(InterfaceIpv4Address.to_json())

# convert the object into a dict
interface_ipv4_address_dict = interface_ipv4_address_instance.to_dict()
# create an instance of InterfaceIpv4Address from a dict
interface_ipv4_address_form_dict = interface_ipv4_address.from_dict(interface_ipv4_address_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


