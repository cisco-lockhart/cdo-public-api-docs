# MspUpgradeRunDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**asdm_version** | **str** | Target management version for the upgrade. | [optional] 
**error_msg** | **str** | Error message if the upgrade run failed. | [optional] 
**force_upgrade** | **bool** | Indicates if this upgrade run is forced on the devices. If set to true, the upgrade is forcefully applied no matter the current upgrades staged on the devices | [optional] 
**last_updated_time** | **datetime** | Time (UTC; represented using the RFC-3339 standard) at which the upgrade run was last updated. | [optional] 
**metadata** | [**UpgradeRunMetadataDto**](UpgradeRunMetadataDto.md) |  | [optional] 
**name** | **str** | Name of the upgrade run. Upgrade runs names are unique in a tenant in SCC Firewall Manager. | [optional] 
**software_version** | **str** | Target software version for the upgrade. | [optional] 
**stage_upgrade_only** | **bool** | Indicates if this upgrade run is to stage the upgrade on the device. If set to true, the upgrade is not applied on the device. | [optional] 
**submission_time** | **datetime** | Time (UTC; represented using the RFC-3339 standard) at which the upgrade run was triggered. | [optional] 
**transaction_uid** | **str** | Unique identifier, represented as a UUID, of the CDO transaction that tracks this upgrade run. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the upgrade run in SCC Firewall Manager. | [optional] 
**upgrade_run_status** | **str** | Current status of the upgrade run. | [optional] 
**upgrade_run_type** | **str** | Type of upgrade run | [optional] 
**username** | **str** | Username of the user who initiated the upgrade run. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_upgrade_run_dto import MspUpgradeRunDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspUpgradeRunDto from a JSON string
msp_upgrade_run_dto_instance = MspUpgradeRunDto.from_json(json)
# print the JSON string representation of the object
print(MspUpgradeRunDto.to_json())

# convert the object into a dict
msp_upgrade_run_dto_dict = msp_upgrade_run_dto_instance.to_dict()
# create an instance of MspUpgradeRunDto from a dict
msp_upgrade_run_dto_form_dict = msp_upgrade_run_dto.from_dict(msp_upgrade_run_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


