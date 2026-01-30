# UpgradeAsaDeviceInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**asdm_version** | **str** | The target ASDM software version to upgrade the device to. | [optional] 
**ignore_maintenance_window** | **bool** | A boolean value, indicating whether device maintenance window should be ignored. If this is set to true, upgrade will be allowed even if device is outside maintenance window. | [optional] 
**software_version** | **str** | The target ASA firmware version to upgrade the device to. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.upgrade_asa_device_input import UpgradeAsaDeviceInput

# TODO update the JSON string below
json = "{}"
# create an instance of UpgradeAsaDeviceInput from a JSON string
upgrade_asa_device_input_instance = UpgradeAsaDeviceInput.from_json(json)
# print the JSON string representation of the object
print(UpgradeAsaDeviceInput.to_json())

# convert the object into a dict
upgrade_asa_device_input_dict = upgrade_asa_device_input_instance.to_dict()
# create an instance of UpgradeAsaDeviceInput from a dict
upgrade_asa_device_input_form_dict = upgrade_asa_device_input.from_dict(upgrade_asa_device_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


