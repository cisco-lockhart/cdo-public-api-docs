# MspCalculateCompatibleUpgradeVersionsInput

List of unique identifiers, represented as UUIDs, of the ASA devices in Security Cloud Control. All specified devices must be in tenants managed by the MSP portal, must be ASA devices in connectivity state ONLINE, and must be in configuration state SYNCED with the ASA_RA_VPN_HEADEND role and a known software version. Per-device validation is performed against each managed tenant; any device that fails validation in its tenant causes the calculation to fail with the offending devices reported in the error message.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_uids** | **List[str]** | The list of devices to calculate compatible upgrade versions for. All of the devices must exist, be associated with the managed tenant, and be in the connectivity state &#39;ONLINE&#39;. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_calculate_compatible_upgrade_versions_input import MspCalculateCompatibleUpgradeVersionsInput

# TODO update the JSON string below
json = "{}"
# create an instance of MspCalculateCompatibleUpgradeVersionsInput from a JSON string
msp_calculate_compatible_upgrade_versions_input_instance = MspCalculateCompatibleUpgradeVersionsInput.from_json(json)
# print the JSON string representation of the object
print(MspCalculateCompatibleUpgradeVersionsInput.to_json())

# convert the object into a dict
msp_calculate_compatible_upgrade_versions_input_dict = msp_calculate_compatible_upgrade_versions_input_instance.to_dict()
# create an instance of MspCalculateCompatibleUpgradeVersionsInput from a dict
msp_calculate_compatible_upgrade_versions_input_form_dict = msp_calculate_compatible_upgrade_versions_input.from_dict(msp_calculate_compatible_upgrade_versions_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


