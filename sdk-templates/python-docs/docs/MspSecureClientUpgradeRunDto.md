# MspSecureClientUpgradeRunDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**devices** | [**List[UpgradedDeviceDto]**](UpgradedDeviceDto.md) | The devices included in this upgrade run. | [optional] 
**error_msg** | **str** | Error message if the upgrade run failed. | [optional] 
**last_updated_time** | **datetime** | Time (UTC; represented using the RFC-3339 standard) at which the upgrade run was last updated. | [optional] 
**name** | **str** | The name of the upgrade run. | [optional] 
**percentage_complete** | **float** | Percentage completion of the current step (0.0 to 1.0). Null when there is no active progress tracking. | [optional] 
**secure_client_version** | **str** | The Secure Client version being upgraded to. | [optional] 
**submission_time** | **datetime** | Time (UTC; represented using the RFC-3339 standard) at which the upgrade run was submitted. | [optional] 
**transaction_uid** | **str** | The unique identifier, represented as a UUID, of the SCC Firewall Manager transaction that tracks this upgrade run. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the upgrade run. | [optional] 
**upgrade_run_status** | **str** | The current status of the upgrade run. One of: VALIDATING_INPUTS, UPLOADING_IMAGES, DONE, DONE_WITH_DEVICE_WARNINGS, DONE_WITH_DEVICE_ERRORS, ERROR. | [optional] 
**upgrade_run_type** | **str** | The type of the Secure Client upgrade. One of: ASA_SECURE_CLIENT_UPGRADE, FTD_SECURE_CLIENT_UPGRADE. | [optional] 
**username** | **str** | Username of the user who initiated the upgrade run. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_secure_client_upgrade_run_dto import MspSecureClientUpgradeRunDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspSecureClientUpgradeRunDto from a JSON string
msp_secure_client_upgrade_run_dto_instance = MspSecureClientUpgradeRunDto.from_json(json)
# print the JSON string representation of the object
print(MspSecureClientUpgradeRunDto.to_json())

# convert the object into a dict
msp_secure_client_upgrade_run_dto_dict = msp_secure_client_upgrade_run_dto_instance.to_dict()
# create an instance of MspSecureClientUpgradeRunDto from a dict
msp_secure_client_upgrade_run_dto_form_dict = msp_secure_client_upgrade_run_dto.from_dict(msp_secure_client_upgrade_run_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


