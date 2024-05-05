# RaVpnDeviceInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_uids** | **List[str]** | List of UIDs of the devices to refresh RA VPN sessions for. Each of these devices has to be an RA VPN headend (this is indicated by the &#x60;deviceRole&#x60; field in the device object being set to &#x60;RA_VPN_HEADEND&#x60;). | 

## Example

```python
from cdo_sdk_python.models.ra_vpn_device_input import RaVpnDeviceInput

# TODO update the JSON string below
json = "{}"
# create an instance of RaVpnDeviceInput from a JSON string
ra_vpn_device_input_instance = RaVpnDeviceInput.from_json(json)
# print the JSON string representation of the object
print(RaVpnDeviceInput.to_json())

# convert the object into a dict
ra_vpn_device_input_dict = ra_vpn_device_input_instance.to_dict()
# create an instance of RaVpnDeviceInput from a dict
ra_vpn_device_input_form_dict = ra_vpn_device_input.from_dict(ra_vpn_device_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


