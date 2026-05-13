# TriggerSecureClientUpgradeInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_uids** | **List[str]** | The unique identifiers of the ASA devices to upgrade in Security Cloud Control. | 
**name** | **str** | An optional name for the upgrade operation to help identify and track the upgrade. | [optional] 
**platforms** | [**List[SecureClientPlatformDto]**](SecureClientPlatformDto.md) | The OS and CPU architecture combinations to upgrade. If omitted, all platforms available for the specified version are uploaded to each device. | [optional] 
**version** | **str** | The Secure Client version to upgrade to. Must be a version present in the Secure Client package catalog. | 

## Example

```python
from scc_firewall_manager_sdk.models.trigger_secure_client_upgrade_input import TriggerSecureClientUpgradeInput

# TODO update the JSON string below
json = "{}"
# create an instance of TriggerSecureClientUpgradeInput from a JSON string
trigger_secure_client_upgrade_input_instance = TriggerSecureClientUpgradeInput.from_json(json)
# print the JSON string representation of the object
print(TriggerSecureClientUpgradeInput.to_json())

# convert the object into a dict
trigger_secure_client_upgrade_input_dict = trigger_secure_client_upgrade_input_instance.to_dict()
# create an instance of TriggerSecureClientUpgradeInput from a dict
trigger_secure_client_upgrade_input_form_dict = trigger_secure_client_upgrade_input.from_dict(trigger_secure_client_upgrade_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


