# SecureClientUpgradeRunDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**devices** | [**List[UpgradedDeviceDto]**](UpgradedDeviceDto.md) | The devices included in this upgrade run. | [optional] 
**error_msg** | **str** | Error message if the upgrade run failed. | [optional] 
**last_updated_time** | **datetime** | Time (UTC; represented using the RFC-3339 standard) at which the upgrade run was last updated. | [optional] 
**name** | **str** | The name of the upgrade run. | [optional] 
**secure_client_version** | **str** | The Secure Client version being upgraded to. | [optional] 
**submission_time** | **datetime** | Time (UTC; represented using the RFC-3339 standard) at which the upgrade run was submitted. | [optional] 
**transaction_uid** | **str** | The unique identifier, represented as a UUID, of the CDO transaction that tracks this upgrade run. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the upgrade run. | [optional] 
**upgrade_run_status** | **str** | The current status of the upgrade run. One of: CHECKING_ASA_DEVICES_DNS, CHECKING_ASA_DEVICES_DNS_RESULTS, UPLOADING_IMAGES, CHECKING_IMAGES_UPLOAD_RESULTS, UPDATING_RA_VPN_CONFIG, DELETING_OLD_IMAGES, DONE, ERROR. | [optional] 
**upgrade_run_type** | **str** | The type of the Secure Client upgrade. One of: ASA_SECURE_CLIENT_UPGRADE, FTD_SECURE_CLIENT_UPGRADE. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.secure_client_upgrade_run_dto import SecureClientUpgradeRunDto

# TODO update the JSON string below
json = "{}"
# create an instance of SecureClientUpgradeRunDto from a JSON string
secure_client_upgrade_run_dto_instance = SecureClientUpgradeRunDto.from_json(json)
# print the JSON string representation of the object
print(SecureClientUpgradeRunDto.to_json())

# convert the object into a dict
secure_client_upgrade_run_dto_dict = secure_client_upgrade_run_dto_instance.to_dict()
# create an instance of SecureClientUpgradeRunDto from a dict
secure_client_upgrade_run_dto_form_dict = secure_client_upgrade_run_dto.from_dict(secure_client_upgrade_run_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


