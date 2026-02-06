# DeviceUpgradeStatusDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**completion_statuses** | [**List[CompletionStatusDto]**](CompletionStatusDto.md) | The completion status of the upgrade on the device. For HA pairs or clusters, there will be multiple completion statuses; one per node of the pair or cluster. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the device in SCC Firewall Manager. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.device_upgrade_status_dto import DeviceUpgradeStatusDto

# TODO update the JSON string below
json = "{}"
# create an instance of DeviceUpgradeStatusDto from a JSON string
device_upgrade_status_dto_instance = DeviceUpgradeStatusDto.from_json(json)
# print the JSON string representation of the object
print(DeviceUpgradeStatusDto.to_json())

# convert the object into a dict
device_upgrade_status_dto_dict = device_upgrade_status_dto_instance.to_dict()
# create an instance of DeviceUpgradeStatusDto from a dict
device_upgrade_status_dto_form_dict = device_upgrade_status_dto.from_dict(device_upgrade_status_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


