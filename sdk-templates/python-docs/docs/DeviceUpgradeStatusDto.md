# DeviceUpgradeStatusDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**completion_statuses** | [**List[CompletionStatusDto]**](CompletionStatusDto.md) | The completion status of the upgrade on the device. | [optional] 
**hardware_model** | **str** | Hardware model of the devices being upgraded. | [optional] 
**managed_tenant_display_name** | **str** | Display name of the managed tenant in CDO. | [optional] 
**managed_tenant_name** | **str** | Name of the managed tenant in CDO. | [optional] 
**managed_tenant_uid** | **str** | Unique identifier, represented as a UUID, of the managed tenant in Security Cloud Control that this device belongs to. | [optional] 
**name** | **str** | Name of the device being upgraded. | [optional] 
**software_version_before_upgrade** | **str** | Software version of the device before the upgrade. | [optional] 
**transaction_uid** | **str** | Unique identifier, represented as a UUID, of the CDO transaction on the managed tenant that tracks this upgrade run. | [optional] 
**uid** | **str** | Unique identifier, represented as a UUID, of the device being upgraded. | [optional] 
**upgrade_run_status** | **str** | Current status of the upgrade run. | [optional] 

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


