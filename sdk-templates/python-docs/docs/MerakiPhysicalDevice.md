# MerakiPhysicalDevice

Physical device information for a Meraki MX appliance.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**lan_ip** | **str** | The LAN IP address assigned to the Meraki Security Appliance. | [optional] 
**latitude** | **float** | The latitude of the location of the Meraki Security Appliance. | [optional] 
**longitude** | **float** | The longitude of the location of the Meraki Security Appliance. | [optional] 
**model** | **str** | The model of the Meraki Security Appliance. | [optional] 
**notes** | **str** | Freeform text notes describing the Meraki Security Appliance. | [optional] 
**serial** | **str** | The serial number of the Meraki Security Appliance. | [optional] 
**wan1_ip** | **str** | The IP address assigned to the first WAN link of the Meraki Security Appliance. | [optional] 
**wan2_ip** | **str** | (Only dual-WAN capable and configured devices) The IP address assigned to the second WAN link of the Meraki Security Appliance. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.meraki_physical_device import MerakiPhysicalDevice

# TODO update the JSON string below
json = "{}"
# create an instance of MerakiPhysicalDevice from a JSON string
meraki_physical_device_instance = MerakiPhysicalDevice.from_json(json)
# print the JSON string representation of the object
print(MerakiPhysicalDevice.to_json())

# convert the object into a dict
meraki_physical_device_dict = meraki_physical_device_instance.to_dict()
# create an instance of MerakiPhysicalDevice from a dict
meraki_physical_device_form_dict = meraki_physical_device.from_dict(meraki_physical_device_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


