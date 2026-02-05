# MspUpgradeAsaDevicesInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**asdm_version** | **str** | The target ASDM software version to upgrade the device to. | [optional] 
**device_uids** | **List[str]** |     &lt;p&gt;List of device identifiers to upgrade in Security Cloud Control.&lt;/p&gt;     &lt;p&gt;Each entry must meet all of the following:&lt;/p&gt;     &lt;ul&gt;       &lt;li&gt;The identifier is a UUID (unique identifier).&lt;/li&gt;       &lt;li&gt;The device is managed by the MSSP portal.&lt;/li&gt;       &lt;li&gt;The device type is &lt;code&gt;ASA&lt;/code&gt;.&lt;/li&gt;       &lt;li&gt;The connectivity state is &lt;code&gt;ONLINE&lt;/code&gt;.&lt;/li&gt;       &lt;li&gt;The device is compatible with the specified software version.&lt;/li&gt;     &lt;/ul&gt;  | 
**force_upgrade** | **bool** | A boolean value, indicating whether the upgrade should be forced. If this is set to true, the upgrade (staged or not) will be executed even if any staged upgrade exists on the device | [optional] [default to False]
**ignore_maintenance_window** | **bool** | A boolean flag that determines whether device maintenance windows should be ignored during a full upgrade | [optional] 
**name** | **str** | Specify a human-readable name for the upgrade run. | [optional] 
**software_version** | **str** | Specify the software version to which all devices will be upgraded. | [optional] 
**stage_upgrade_only** | **bool** | A boolean value, indicating whether the upgrade should be staged. If set to true, the image will be downloaded onto the devices and readiness checks will be performed. However, the upgrade will not be applied to the devices. | [optional] [default to False]

## Example

```python
from scc_firewall_manager_sdk.models.msp_upgrade_asa_devices_input import MspUpgradeAsaDevicesInput

# TODO update the JSON string below
json = "{}"
# create an instance of MspUpgradeAsaDevicesInput from a JSON string
msp_upgrade_asa_devices_input_instance = MspUpgradeAsaDevicesInput.from_json(json)
# print the JSON string representation of the object
print(MspUpgradeAsaDevicesInput.to_json())

# convert the object into a dict
msp_upgrade_asa_devices_input_dict = msp_upgrade_asa_devices_input_instance.to_dict()
# create an instance of MspUpgradeAsaDevicesInput from a dict
msp_upgrade_asa_devices_input_form_dict = msp_upgrade_asa_devices_input.from_dict(msp_upgrade_asa_devices_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


