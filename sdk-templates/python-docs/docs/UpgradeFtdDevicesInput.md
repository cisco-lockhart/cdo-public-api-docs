# UpgradeFtdDevicesInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_uids** | **List[str]** | The set of unique identifiers, represented as UUIDs, of the devices to upgrade in Security Cloud Control. All of the devices in the list have to be compatible with the upgrade package. | 
**ignore_maintenance_window** | **bool** | A boolean value, indicating whether devices maintenance window should be ignored. If this is set to true, upgrade will be allowed even if there are devices outside their maintenance window. | [optional] 
**name** | **str** | An optional name for the upgrade operation to help identify and track the upgrade. | [optional] 
**stage_upgrade** | **bool** | A boolean value, indicating whether the upgrade should be staged. If this is set to true, the image will be downloaded on to the device and readiness checks will be performed. However, the upgrade will not be applied to the device. | [optional] [default to False]
**upgrade_package_uid** | **str** | The unique identifier, represented as a UUID, of the upgrade package to be applied to the devices. | 

## Example

```python
from scc_firewall_manager_sdk.models.upgrade_ftd_devices_input import UpgradeFtdDevicesInput

# TODO update the JSON string below
json = "{}"
# create an instance of UpgradeFtdDevicesInput from a JSON string
upgrade_ftd_devices_input_instance = UpgradeFtdDevicesInput.from_json(json)
# print the JSON string representation of the object
print(UpgradeFtdDevicesInput.to_json())

# convert the object into a dict
upgrade_ftd_devices_input_dict = upgrade_ftd_devices_input_instance.to_dict()
# create an instance of UpgradeFtdDevicesInput from a dict
upgrade_ftd_devices_input_form_dict = upgrade_ftd_devices_input.from_dict(upgrade_ftd_devices_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


