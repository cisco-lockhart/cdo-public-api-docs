# UpgradeRunDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**device_uids** | **List[str]** | The set of device UIDs that are part of this upgrade run. | [optional] 
**device_upgrade_statuses** | [**List[DeviceUpgradeStatusDto]**](DeviceUpgradeStatusDto.md) | The current status of each device in the upgrade run. | [optional] 
**name** | **str** | The name of the upgrade run. Upgrade runs names are unique in a tenant in SCC Firewall Manager. | [optional] 
**stage_upgrade_only** | **bool** | Indicates if this upgrade run is to stage the upgrade on the device. If set to true, the upgrade is not applied on the device. | [optional] 
**transaction_uid** | **str** | The unique identifier of the CDO transaction that tracks this upgrade run. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the device upgrade run in SCC Firewall Manager. | [optional] 
**upgrade_package_uid** | **str** | (cdFMC-managed FTD device upgrades only) The unique identifier of the upgrade package being used. | [optional] 
**upgrade_run_status** | **str** | The current status of the upgrade run. | [optional] 
**upgrade_run_type** | **str** | The type of the upgrade run | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.upgrade_run_dto import UpgradeRunDto

# TODO update the JSON string below
json = "{}"
# create an instance of UpgradeRunDto from a JSON string
upgrade_run_dto_instance = UpgradeRunDto.from_json(json)
# print the JSON string representation of the object
print(UpgradeRunDto.to_json())

# convert the object into a dict
upgrade_run_dto_dict = upgrade_run_dto_instance.to_dict()
# create an instance of UpgradeRunDto from a dict
upgrade_run_dto_form_dict = upgrade_run_dto.from_dict(upgrade_run_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


