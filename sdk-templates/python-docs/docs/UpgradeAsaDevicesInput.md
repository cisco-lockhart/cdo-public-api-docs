# UpgradeAsaDevicesInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**asdm_version** | **str** | The target ASDM software version to upgrade the devices to. | [optional] 
**device_uids** | **List[str]** | The set of unique identifiers, represented as UUIDs, of the ASA devices to upgrade in Security Cloud Control. | 
**force_upgrade** | **bool** | A boolean value, indicating whether the upgrade should be forced. If this is set to true, the upgrade (staged or not) will be executed even if any staged upgrade exists on the device | [optional] [default to False]
**ignore_maintenance_window** | **bool** | A boolean value, indicating whether device maintenance window should be ignored. If this is set to true, upgrade will be allowed even if devices are outside maintenance window. | [optional] 
**name** | **str** | An optional name for the upgrade operation to help identify and track the upgrade. | [optional] 
**software_version** | **str** | The target ASA firmware version to upgrade the devices to. | [optional] 
**stage_upgrade** | **bool** | A boolean value, indicating whether the upgrade should be staged. If this is set to true, the image will be downloaded on to the device and readiness checks will be performed. However, the upgrade will not be applied to the device. | [optional] [default to False]

## Example

```python
from scc_firewall_manager_sdk.models.upgrade_asa_devices_input import UpgradeAsaDevicesInput

# TODO update the JSON string below
json = "{}"
# create an instance of UpgradeAsaDevicesInput from a JSON string
upgrade_asa_devices_input_instance = UpgradeAsaDevicesInput.from_json(json)
# print the JSON string representation of the object
print(UpgradeAsaDevicesInput.to_json())

# convert the object into a dict
upgrade_asa_devices_input_dict = upgrade_asa_devices_input_instance.to_dict()
# create an instance of UpgradeAsaDevicesInput from a dict
upgrade_asa_devices_input_form_dict = upgrade_asa_devices_input.from_dict(upgrade_asa_devices_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


