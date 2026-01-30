# UpgradeRunMetadataDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**devices** | [**List[DeviceUpgradeStatusDto]**](DeviceUpgradeStatusDto.md) | Set of device UIDs that are part of this upgrade run. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.upgrade_run_metadata_dto import UpgradeRunMetadataDto

# TODO update the JSON string below
json = "{}"
# create an instance of UpgradeRunMetadataDto from a JSON string
upgrade_run_metadata_dto_instance = UpgradeRunMetadataDto.from_json(json)
# print the JSON string representation of the object
print(UpgradeRunMetadataDto.to_json())

# convert the object into a dict
upgrade_run_metadata_dto_dict = upgrade_run_metadata_dto_instance.to_dict()
# create an instance of UpgradeRunMetadataDto from a dict
upgrade_run_metadata_dto_form_dict = upgrade_run_metadata_dto.from_dict(upgrade_run_metadata_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


